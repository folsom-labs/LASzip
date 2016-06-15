package main

import (
	"errors"
	"fmt"
	"io"
)

// VariableLengthRecord.RecordID when VariableLengthRecord.UserID is 'LASF_Projection'
const (
	GeoKeyDirectoryTag = 34735 // required
	GeoDoubleParamsTag = 34736 // optional or not (depending on GeoKeyDirectoryTag)
	GeoASCIIParamsTag  = 34737 // optional or not (depending on GeoKeyDirectoryTag)

	// optional
	ModelTiePointTag       = 33922
	ModelPixelScaleTag     = 33550
	ModelTransformationTag = 34264

	LASF_Projection = "LASF_Projection"
	LASF_Spec       = "LASF_Spec"
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

// GeoKeyRaw describes a key
type GeoKeyRaw struct {
	// value in GeoTIFF spec
	KeyID uint16
	// 0 - data is in ValueOffset as unsigned short
	// 34736 - data is at ValueOffset index of GeoDoubleParams record
	// 34767 - data is at ValueOffset index of GetAsciiParams record
	TIFFTagLocation uint16
	// number of characters in string values of GeoAsciiParam record
	Count       uint16
	ValueOffset uint16
}

// GeoKeyDirectory describes VariableLengthRecord.Data when
// UserID == LasfProjection and RecordID == GeoKeyDirectoryTag
type GeoKeyDirectory struct {
	KeyDirectoryVersion uint16 // must be 1
	KeyRevision         uint16 // must be 1
	MinorRevision       uint16 // must be 0
	NumberOfKeys        uint16
	KeysRaw             []GeoKeyRaw
}

// GeoDoubleParams has data from GeoDoubleParamsTag record
type GeoDoubleParams struct {
	Values []float64
}

// GeoASCIIParams has data from GeoASCIIParamsTag record
type GeoASCIIParams struct {
	// array of ascii data. null-separated strings, referenced  by
	// GeoKey
	Data []byte
}

// Classification describes elements of LASF_Spec record
type Classification struct {
	ClassNumber uint8
	Description string
}

// PointDataRecord0 describes Point Data Record format 0
type PointDataRecord0 struct {
	X int32
	Y int32
	Z int32
	// pulse return magnitude
	Intensity      uint16
	flags          uint8
	Classification uint8
	// -90 .. +90, 0 is nadir, -90 is to the left side of the aircraft in the
	// direction of flight
	ScanAngleRank uint8
	UserData      uint8
	// Where did point come from, 0 means this file, otherwise should be
	// FileSourceID
	PointSourceID uint16

	// those are calculated from flags
	// pulse return number for a given output pulse
	ReturnNumber int // bits 0,1,2
	// pulse return number for a given output pulse
	// values: 1, 2...
	NumberOfReturns int // bits 3, 4, 5; given pulse
	// true for positive scan direction (scan left to right side of in-track direction)
	ScanDirectionFlag bool
	// true if point at the end of the scan (last point in a given scan line
	// before scan changes direction)
	EdgeOfFlightLine bool
}

// PointDataRecord1 describes Point Data Record format 1
type PointDataRecord1 struct {
	PointDataRecord0
	GPSTime float64
}

// PointDataRecord2 describes Point Data Record format 2
type PointDataRecord2 struct {
	PointDataRecord0
	Red   uint16
	Green uint16
	Blue  uint16
}

// PointDataRecord3 describes Point Data Record format 3
type PointDataRecord3 struct {
	PointDataRecord0
	GPSTime float64
	Red     uint16
	Green   uint16
	Blue    uint16
}

// ClassificationType defines ASPRS LIDAR point classification
type ClassificationType int

const (
	Created ClassificationType = 0
	Unclassified
	Ground
	LowVegetation
	MediumVegetation
	HighVegetation
	BuildingLowPoint
	ModelKeyPoint
	Water
	Reserved1
	Reserved2
	OverlapPoints
)

// LasReader is a reader for .las files
type LasReader struct {
	r                     io.Reader
	Header                *LasPublicHeader
	VariableLengthRecords []*VariableLengthRecord
	GeoKeyDirectory       *GeoKeyDirectory
	GeoDoubleParams       []float64
	GeoASCIIParams        *GeoASCIIParams
	Classifications       []Classification
	Error                 error
}

func uint16IsBitSet(v uint16, bitNo int) bool {
	var mask uint16 = 1 << uint16(bitNo)
	return v&mask == 0
}

func uint8IsBitSet(v uint8, bitNo int) bool {
	var mask uint8 = 1 << uint8(bitNo)
	return v&mask == 0
}

// return first nBits from b as int and the remaining bits shifted
func eatBits(b uint8, nBits uint) (int, uint8) {
	mask := uint8(1<<nBits) - 1
	res := int(b & mask)
	//fmt.Printf("b: 0x%x, nBits: %d, mask: 0x%x, res: %d", b, nBits, mask, res)
	b = b >> uint(nBits)
	//fmt.Printf(", b after: 0x%b\n", b)
	return res, b
}

// we support 1.0, 1.1, 1.2, 1.3, 1.4
func isValidVersion(major, minor byte) bool {
	if major != 1 || minor > 4 {
		return false
	}
	return true
}

// GetClassification returns ASPRS ClassificationType, bits 0..4 of Classification
func (p *PointDataRecord0) GetClassification() ClassificationType {
	n, _ := eatBits(p.Classification, 5)
	return ClassificationType(n)
}

// IsSynthetic returns true if not from LIDAR
func (p *PointDataRecord0) IsSynthetic() bool {
	return uint8IsBitSet(p.Classification, 5)
}

// IsKeyPoint returns true if considered a model key-point and thus should not
// be witheld in a thinning algorithm
func (p *PointDataRecord0) IsKeyPoint() bool {
	return uint8IsBitSet(p.Classification, 6)
}

// IsWithheld returns true if point should not be included in processing
// (synonymous with IsDeleted)
func (p *PointDataRecord0) IsWithheld() bool {
	return uint8IsBitSet(p.Classification, 7)
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
func (r *LasReader) ReadVariableLengthRecord(br *BinaryReader) error {
	var vlr VariableLengthRecord
	vlr.reserved = br.ReadUint16()
	vlr.UserID = br.ReadFixedString(16)
	vlr.RecordID = br.ReadUint16()
	vlr.RecordLengthAfterHeader = br.ReadUint16()
	vlr.Description = br.ReadFixedString(32)

	if br.Error != nil {
		return br.Error
	}
	r.VariableLengthRecords = append(r.VariableLengthRecords, &vlr)

	if vlr.UserID == LASF_Spec {
		if vlr.RecordID == 0 {
			return r.ReadClassificationLookup(br, &vlr)
		}
	}

	if vlr.UserID == LASF_Projection {
		if vlr.RecordID == GeoKeyDirectoryTag {
			return r.ReadGeoKeyDirectory(br)
		}

		if vlr.RecordID == GeoDoubleParamsTag {
			return r.ReadGeoDoubleParams(br, &vlr)
		}

		if vlr.RecordID == GeoASCIIParamsTag {
			return r.ReadGeoASCIIParams(br, &vlr)
		}
	}

	vlr.Data = br.ReadBytes(int(vlr.RecordLengthAfterHeader))
	return r.Error
}

// ReadGeoDoubleParams reads GeoDoubleParams record
func (r *LasReader) ReadGeoDoubleParams(br *BinaryReader, vlr *VariableLengthRecord) error {
	if vlr.RecordLengthAfterHeader%8 != 0 {
		return fmt.Errorf("Unexpected size of GeoDoubleParamsTag record. Is %d and expected to be multiple of 8", vlr.RecordLengthAfterHeader)
	}
	n := int(vlr.RecordLengthAfterHeader) / 8
	for i := 0; i < n; i++ {
		param := br.ReadFloat64()
		r.GeoDoubleParams = append(r.GeoDoubleParams, param)
	}
	return br.Error
}

// ReadGeoASCIIParams reads GeoASCIIParams record
func (r *LasReader) ReadGeoASCIIParams(br *BinaryReader, vlr *VariableLengthRecord) error {
	var rec GeoASCIIParams
	rec.Data = br.ReadBytes(int(vlr.RecordLengthAfterHeader))
	r.GeoASCIIParams = &rec
	return br.Error
}

// ReadClassificationLookup reads Classification lookup
func (r *LasReader) ReadClassificationLookup(br *BinaryReader, vlr *VariableLengthRecord) error {
	if vlr.RecordLengthAfterHeader != 255*16 {
		return fmt.Errorf("Unexpected size of Classification Lookup record. Is %d and should be %d (255*16)", vlr.RecordLengthAfterHeader, 16*255)
	}
	for i := 0; i < 255; i++ {
		var rec Classification
		rec.ClassNumber = br.ReadUint8()
		rec.Description = br.ReadFixedString(15)
		r.Classifications = append(r.Classifications, rec)
	}
	return r.Error
}

// ReadGeoKeyDirectory reads GeoKeyDirectory record
func (r *LasReader) ReadGeoKeyDirectory(br *BinaryReader) error {
	var rec GeoKeyDirectory
	rec.KeyDirectoryVersion = br.ReadUint16()
	rec.KeyRevision = br.ReadUint16()
	rec.MinorRevision = br.ReadUint16()
	rec.NumberOfKeys = br.ReadUint16()
	for i := 0; i < int(rec.NumberOfKeys); i++ {
		var key GeoKeyRaw
		key.KeyID = br.ReadUint16()
		key.TIFFTagLocation = br.ReadUint16()
		key.Count = br.ReadUint16()
		key.ValueOffset = br.ReadUint16()
		rec.KeysRaw = append(rec.KeysRaw, key)
	}
	r.GeoKeyDirectory = &rec
	return br.Error
}

// ReadPointDataRecord0 reads Point Data Record Format 0
func ReadPointDataRecord0(r *BinaryReader) (*PointDataRecord0, error) {
	var p PointDataRecord0

	p.X = r.ReadInt32()
	p.Y = r.ReadInt32()
	p.Z = r.ReadInt32()
	p.Intensity = r.ReadUint16()
	p.flags = r.ReadUint8()
	p.Classification = r.ReadUint8()
	p.ScanAngleRank = r.ReadUint8()
	p.UserData = r.ReadUint8()
	p.PointSourceID = r.ReadUint16()

	b := p.flags
	p.ReturnNumber, b = eatBits(b, 3)
	p.NumberOfReturns, b = eatBits(b, 3)
	n, b := eatBits(b, 1)
	p.ScanDirectionFlag = (n == 1)
	n, b = eatBits(b, 1)
	p.EdgeOfFlightLine = (n == 1)

	return &p, r.Error
}

// ReadPointDataRecord1 reads Point Data Record Format 1
func ReadPointDataRecord1(r *BinaryReader) (*PointDataRecord1, error) {
	p0, err := ReadPointDataRecord0(r)
	p := PointDataRecord1{
		PointDataRecord0: *p0,
		GPSTime:          r.ReadFloat64(),
	}
	if err != nil {
		return &p, err
	}
	return &p, r.Error
}

// ReadPointDataRecord2 reads Point Data Record Format 2
func ReadPointDataRecord2(r *BinaryReader) (*PointDataRecord2, error) {
	p0, err := ReadPointDataRecord0(r)
	p := PointDataRecord2{
		PointDataRecord0: *p0,
		Red:              r.ReadUint16(),
		Green:            r.ReadUint16(),
		Blue:             r.ReadUint16(),
	}
	if err != nil {
		return &p, err
	}
	return &p, r.Error
}

// ReadPointDataRecord3 reads Point Data Record Format 3
func ReadPointDataRecord3(r *BinaryReader) (*PointDataRecord3, error) {
	p0, err := ReadPointDataRecord0(r)
	p := PointDataRecord3{
		PointDataRecord0: *p0,
		GPSTime:          r.ReadFloat64(),
		Red:              r.ReadUint16(),
		Green:            r.ReadUint16(),
		Blue:             r.ReadUint16(),
	}
	if err != nil {
		return &p, err
	}
	return &p, r.Error
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
	for i := 0; i < n; i++ {
		err := r.ReadVariableLengthRecord(br)
		if err != nil {
			return err
		}
	}
	geoDir := r.GeoKeyDirectory
	if geoDir == nil {
		return errors.New("missing GeoKeyDirectoryTag record")
	}
	if geoDir.KeyDirectoryVersion != 1 {
		return fmt.Errorf("GeoKeyDirectory.KeyDirectoryVersion is %d, expected 1", geoDir.KeyDirectoryVersion)
	}
	if geoDir.KeyRevision != 1 {
		return fmt.Errorf("GeoKeyDirectory.KeyRevision is %d, expected 1", geoDir.KeyRevision)
	}
	if geoDir.MinorRevision != 0 {
		return fmt.Errorf("GeoKeyDirectory.MinorRevision is %d, expected 0", geoDir.MinorRevision)
	}
	for _, key := range geoDir.KeysRaw {
		loc := key.TIFFTagLocation
		switch loc {
		case 0:
			// do nothing
		case GeoDoubleParamsTag:
			if r.GeoDoubleParams == nil {
				return fmt.Errorf("key location in double params but GeoDoubleParamsTag record not present")
				// TODO: update key with converted value ?
			}
		case GeoASCIIParamsTag:
			// Note: 1.2 spec uses 34767 and not 34737, which I assume is a typo
			if r.GeoASCIIParams == nil {
				return fmt.Errorf("key location in ASCII params but GeoASCIIParamsTag record not present")
				// TODO: update key with converted value ?
			}
		}
	}
	return nil
}

// ReadHeader reads public header
func (r *LasReader) ReadHeader() (*LasPublicHeader, error) {
	br := NewBinaryReader(r.r)
	r.Header, r.Error = ReadLasPublicHeader(br)
	//fmt.Printf("bytes consumed: %d\n", br.BytesConsumed)
	return r.Header, r.Error
}
