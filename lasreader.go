package main

import (
	"fmt"
	"io"
)

// LasPublicHeader describes a public header in .las file
type LasPublicHeader struct {
	FileSourceID                  uint16
	GlobalEncoding                uint16
	VersionMajor                  byte
	VersionMinor                  byte
	SystemIdentifier              string
	GeneratingSoftware            string
	FileCreationDayOfYear         uint16
	FileCreationYear              uint16
	HeaderSize                    uint16
	OffsetToPointData             uint32
	NumberOfVariableLengthRecords uint32
	PointDataFormatID             byte
	PointDataRecordLength         uint16
	NumberOfPointRecords          uint32
	NumberOfPointsByReturn        [5]uint32
	XScaleFactor                  float64
	YScaleFactor                  float64
	ZScaleFactor                  float64
	XOffset                       float64
	YOffset                       float64
	ZOffset                       float64
	MaxX                          float64
	MinX                          float64
	MaxY                          float64
	MinY                          float64
	MaxZ                          float64
	MinZ                          float64

	// calculated fields
	// corresponds to bit 0 of GlobalEncoding. If true, GPS Time in Point Records
	// is standard GPS TIME (satellite GPS Time) minus 10e9.
	IsGPSTimeStandard bool
}

// VariableLengthRecord describes Variable Length Record Header and its data
// It's placed after LasPublicHeader
type VariableLengthRecord struct {
	reserved                uint16
	UserID                  string
	RecordID                uint16
	RecordLengthAfterHeader uint16
	Description             string

	Data []byte
}

// LasReader is a reader for .las files
type LasReader struct {
	r                     io.Reader
	Header                *LasPublicHeader
	VariableLengthRecords []*VariableLengthRecord
	Error                 error
}

// we support 1.0, 1.1, 1.2, 1.3, 1.4
func isValidVersion(major, minor byte) bool {
	if major != 1 || minor > 4 {
		return false
	}
	return true
}

func uint16IsBitSet(v uint16, bitNo int) bool {
	var mask uint16 = 1 << uint16(bitNo)
	return v&mask == 0
}

// NewLasReader creates a LasReader
func NewLasReader(r io.Reader) *LasReader {
	return &LasReader{
		r: r,
	}
}

// ReadLasPublicHeader reads LasPublicHeader from a reader
func ReadLasPublicHeader(r *BinaryReader) (*LasPublicHeader, error) {
	var hdr LasPublicHeader
	sig := r.ReadFixedString(4)
	if sig != "LASF" {
		return nil, fmt.Errorf("Invalid signature, expected 'LASF', got '%s'", sig)
	}
	hdr.FileSourceID = r.ReadUint16()
	hdr.GlobalEncoding = r.ReadUint16()
	// project id - guid data 1 : 4 bytes
	// project id - guid data 2 : 2 bytes
	// project id - guid data 3 : 2 bytes
	// project id - guid data 4 : 8 bytes
	r.Skip(16)
	hdr.VersionMajor = r.ReadUint8()
	hdr.VersionMinor = r.ReadUint8()
	if !isValidVersion(hdr.VersionMajor, hdr.VersionMinor) {
		return nil, fmt.Errorf("Invalid version: %d.%d (we understand 1.0-1.4)", hdr.VersionMajor, hdr.VersionMinor)
	}
	hdr.SystemIdentifier = r.ReadFixedString(32)
	hdr.GeneratingSoftware = r.ReadFixedString(32)
	hdr.FileCreationDayOfYear = r.ReadUint16()
	hdr.FileCreationYear = r.ReadUint16()
	hdr.HeaderSize = r.ReadUint16()
	hdr.OffsetToPointData = r.ReadUint32()
	hdr.NumberOfVariableLengthRecords = r.ReadUint32()
	hdr.PointDataFormatID = r.ReadUint8()
	hdr.PointDataRecordLength = r.ReadUint16()
	hdr.NumberOfPointRecords = r.ReadUint32()
	for i := 0; i < 5; i++ {
		hdr.NumberOfPointsByReturn[i] = r.ReadUint32()
	}

	hdr.XScaleFactor = r.ReadFloat64()
	hdr.YScaleFactor = r.ReadFloat64()
	hdr.ZScaleFactor = r.ReadFloat64()

	hdr.XOffset = r.ReadFloat64()
	hdr.YOffset = r.ReadFloat64()
	hdr.ZOffset = r.ReadFloat64()

	hdr.MaxX = r.ReadFloat64()
	hdr.MinX = r.ReadFloat64()

	hdr.MaxY = r.ReadFloat64()
	hdr.MinY = r.ReadFloat64()

	hdr.MaxZ = r.ReadFloat64()
	hdr.MinZ = r.ReadFloat64()

	// TODO: read more fields if v1.3 or v1.4

	// set calculated fields
	hdr.IsGPSTimeStandard = uint16IsBitSet(hdr.GlobalEncoding, 0)

	// skip fields at the end of the header we don't yet understand
	gap := int(hdr.HeaderSize) - r.BytesConsumed
	fatalIf(gap < 0)
	r.Skip(gap)

	return &hdr, r.Error
}

// ReadVariableLengthRecord reads VariableLengthRecord
func ReadVariableLengthRecord(r *BinaryReader) (*VariableLengthRecord, error) {
	var hdr VariableLengthRecord
	hdr.reserved = r.ReadUint16()
	hdr.UserID = r.ReadFixedString(16)
	hdr.RecordID = r.ReadUint16()
	hdr.RecordLengthAfterHeader = r.ReadUint16()
	hdr.Description = r.ReadFixedString(32)

	hdr.Data = r.ReadBytes(int(hdr.RecordLengthAfterHeader))
	return &hdr, r.Error
}

// ReadHeaders reads public headers and Variable Length Records
func (r *LasReader) ReadHeaders() error {
	var err error
	r.Header, err = r.ReadHeader()
	if err != nil {
		return err
	}
	br := NewBinaryReader(r.r)
	n := int(r.Header.NumberOfVariableLengthRecords)
	var records []*VariableLengthRecord
	for i := 0; i < n; i++ {
		r, err := ReadVariableLengthRecord(br)
		if err != nil {
			return err
		}
		records = append(records, r)
	}
	r.VariableLengthRecords = records
	return nil
}

// ReadHeader reads public header
func (r *LasReader) ReadHeader() (*LasPublicHeader, error) {
	br := NewBinaryReader(r.r)
	r.Header, r.Error = ReadLasPublicHeader(br)
	//fmt.Printf("bytes consumed: %d\n", br.BytesConsumed)
	return r.Header, r.Error
}
