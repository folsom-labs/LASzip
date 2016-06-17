package main

import (
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

// Classification describes elements of LASF_Spec record
type Classification struct {
	ClassNumber uint8
	Description string
}

// PointDataRecord describes Point Data Record, all formats
type PointDataRecord struct {
	FormatID int

	// data available in all formats
	X int32
	Y int32
	Z int32
	// pulse return magnitude
	Intensity      uint16
	flags          uint8
	Classification uint8
	// -90 .. +90, 0 is nadir, -90 is to the left side of the aircraft in the
	// direction of flight
	ScanAngleRank int8
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

	// data available only in some formats
	GPSTime float64 // 1, 3

	Red   uint16 // 2, 3
	Green uint16 // 2, 3
	Blue  uint16 // 2, 3
}

// ClassificationType defines ASPRS LIDAR point classification
type ClassificationType int

const (
	Created ClassificationType = iota
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

// GetClassificationName returns name of a given ClassificationType
func GetClassificationName(c ClassificationType) string {
	switch c {
	case Unclassified:
		return "Unclassified"
	case Ground:
		return "Ground"
	case LowVegetation:
		return "LowVegetation"
	case MediumVegetation:
		return "MediumVegetation"
	case HighVegetation:
		return "HighVegetation"
	case BuildingLowPoint:
		return "BuildingLowPoint"
	case ModelKeyPoint:
		return "ModelKeyPoint"
	case Water:
		return "Water"
	case OverlapPoints:
		return "OverlapPoints"
	default:
		return fmt.Sprintf("Unknown classification (%d)", int(c))
	}
}

// LasReader is a reader for .las files
type LasReader struct {
	r                     io.ReadSeeker
	Header                *LasPublicHeader
	VariableLengthRecords []*VariableLengthRecord
	GeoKeyInfo            GeoKeyInfo
	GeoTags               *GeoTags // decoded version of GeoKeyInfo
	Classifications       []Classification
	Error                 error

	// for optimizing sequential point reading, we remember what
	// point was read last time
	lastPointRead int
}

func uint16IsBitSet(v uint16, bitNo int) bool {
	var mask uint16 = 1 << uint16(bitNo)
	return v&mask != 0
}

func uint8IsBitSet(v uint8, bitNo int) bool {
	var mask uint8 = 1 << uint8(bitNo)
	return v&mask != 0
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
func (p *PointDataRecord) GetClassification() ClassificationType {
	n, _ := eatBits(p.Classification, 5)
	return ClassificationType(n)
}

// IsSynthetic returns true if not from LIDAR
func (p *PointDataRecord) IsSynthetic() bool {
	return uint8IsBitSet(p.Classification, 5)
}

// IsKeyPoint returns true if considered a model key-point and thus should not
// be witheld in a thinning algorithm
func (p *PointDataRecord) IsKeyPoint() bool {
	return uint8IsBitSet(p.Classification, 6)
}

// IsWithheld returns true if point should not be included in processing
// (synonymous with IsDeleted)
func (p *PointDataRecord) IsWithheld() bool {
	return uint8IsBitSet(p.Classification, 7)
}

func intIntArray(n int, a []int) bool {
	for _, v := range a {
		if n == v {
			return true
		}
	}
	return false
}

func hasGPSTime(formatID int) bool {
	switch formatID {
	case 1, 3:
		return true
	}
	return false
}

// GetGPSTime returns GPS time and a bool indicating if GPS time is available
// for this point
func (p *PointDataRecord) GetGPSTime() (float64, bool) {
	return p.GPSTime, hasGPSTime(p.FormatID)
}

// NewLasReader creates a LasReader
func NewLasReader(r io.ReadSeeker) *LasReader {
	return &LasReader{
		r:             r,
		lastPointRead: -1,
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
	if hdr.PointDataFormatID > 3 {
		return nil, fmt.Errorf("Unsupported point format id: %d (we understand 0-3)", hdr.PointDataFormatID)
	}
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
	var params []float64
	for i := 0; i < n; i++ {
		param := br.ReadFloat64()
		params = append(params, param)
	}
	r.GeoKeyInfo.DoubleParams = params
	return br.Error
}

// ReadGeoASCIIParams reads GeoASCIIParams record
func (r *LasReader) ReadGeoASCIIParams(br *BinaryReader, vlr *VariableLengthRecord) error {
	r.GeoKeyInfo.ASCIIParams = br.ReadBytes(int(vlr.RecordLengthAfterHeader))
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
		key.ValueOrOffset = br.ReadUint16()
		rec.KeysRaw = append(rec.KeysRaw, key)
	}
	r.GeoKeyInfo.Directory = &rec
	return br.Error
}

// ReadPointDataRecord reads format records
func ReadPointDataRecord(r *BinaryReader, formatID int) (*PointDataRecord, error) {

	// read common data (basically format 0)
	var p PointDataRecord
	p.FormatID = formatID
	p.X = r.ReadInt32()
	p.Y = r.ReadInt32()
	p.Z = r.ReadInt32()
	p.Intensity = r.ReadUint16()
	p.flags = r.ReadUint8()
	p.Classification = r.ReadUint8()
	p.ScanAngleRank = r.ReadInt8()
	p.UserData = r.ReadUint8()
	p.PointSourceID = r.ReadUint16()

	b := p.flags
	p.ReturnNumber, b = eatBits(b, 3)
	p.NumberOfReturns, b = eatBits(b, 3)
	n, b := eatBits(b, 1)
	p.ScanDirectionFlag = (n == 1)
	n, b = eatBits(b, 1)
	p.EdgeOfFlightLine = (n == 1)

	if formatID == 1 {
		p.GPSTime = r.ReadFloat64()
	} else if formatID == 2 {
		p.Red = r.ReadUint16()
		p.Green = r.ReadUint16()
		p.Blue = r.ReadUint16()
	} else if formatID == 3 {
		p.GPSTime = r.ReadFloat64()
		p.Red = r.ReadUint16()
		p.Green = r.ReadUint16()
		p.Blue = r.ReadUint16()
	} else {
		return nil, fmt.Errorf("%d is unsupported PointDataFormatID", formatID)
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
	r.GeoTags, err = DecodeGeoKeyInfo(&r.GeoKeyInfo)
	return err
}

// ReadHeader reads public header
func (r *LasReader) ReadHeader() (*LasPublicHeader, error) {
	br := NewBinaryReader(r.r)
	r.Header, r.Error = ReadLasPublicHeader(br)
	//fmt.Printf("bytes consumed: %d\n", br.BytesConsumed)
	return r.Header, r.Error
}

// ReadPoint reads a point 0..r.Header.NumberOfPointRecords-1
func (r *LasReader) ReadPoint(n int) (*PointDataRecord, error) {
	if n < 0 || n >= int(r.Header.NumberOfPointRecords) {
		return nil, fmt.Errorf("%d is invalid point number, must be >= 0 and < %d", n, r.Header.NumberOfPointRecords)
	}
	if r.lastPointRead != n-1 {
		offset := int64(r.Header.OffsetToPointData) + (int64(n) * int64(r.Header.PointDataRecordLength))
		newOffset, err := r.r.Seek(offset, 0)
		if err != nil {
			return nil, err
		}
		fatalIf(newOffset != offset)
	}
	br := NewBinaryReader(r.r)
	res, err := ReadPointDataRecord(br, int(r.Header.PointDataFormatID))
	if err != nil {
		return nil, err
	}
	toSkip := int(r.Header.PointDataRecordLength) - br.BytesConsumed
	fatalIf(toSkip < 0)
	br.Skip(toSkip)
	r.lastPointRead = n
	return res, br.Error
}

// TransformPoints converts points as found in PointDataRecord to real value
func (r *LasReader) TransformPoints(x, y, z int32) (float64, float64, float64) {
	resX := r.Header.XOffset + (float64(x) * r.Header.XScaleFactor)
	resY := r.Header.YOffset + (float64(y) * r.Header.YScaleFactor)
	resZ := r.Header.ZOffset + (float64(z) * r.Header.ZScaleFactor)
	return resX, resY, resZ
}

// GetModelType returns value of GTModelTypeGeoKey GeoTiff key and
// a bool that indicates if value is present and valid
func (r *LasReader) GetModelType() (ModelType, bool) {
	tagShort := r.GeoTags.FindGeoTagShort(GTModelTypeGeoKey)
	if tagShort != nil {
		v := tagShort.Value
		if v >= 1 && v <= 3 {
			return ModelType(v), true
		}
	}
	return ModelType(0), false
}

// GetRasterType returns value GTRasterTypeGeoKey
func (r *LasReader) GetRasterType() (RasterType, bool) {
	tagShort := r.GeoTags.FindGeoTagShort(GTRasterTypeGeoKey)
	if tagShort != nil {
		v := tagShort.Value
		if v >= 1 && v <= 2 {
			return RasterType(v), true
		}
	}
	return RasterType(0), false
}
