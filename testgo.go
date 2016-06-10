package main

import (
	"encoding/binary"
	"fmt"
	"io"
	"os"
	"os/exec"
	"path/filepath"
)

func fatalIfErr(err error) {
	if err != nil {
		fmt.Printf("error: '%s'\n", err)
		os.Exit(1)
	}
}

func verifyFileExists(path string) {
	fi, err := os.Stat(path)
	if err != nil {
		fmt.Printf("file '%s' doesn't exist, err: '%s'\n", path, err)
		os.Exit(1)
	}
	if !fi.Mode().IsRegular() {
		fmt.Printf("'%s' is not a file\n", path)
		os.Exit(1)
	}
}

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
}

// ReadLasPublicHeader reads LasPublicHeader from a reader
func ReadLasPublicHeader(rIn io.Reader) (*LasPublicHeader, error) {
	var hdr LasPublicHeader
	r := NewBinaryReader(rIn)
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

	return &hdr, r.Error
}

func dumpHeader(hdr *LasPublicHeader, w io.Writer) {
	fmt.Fprintf(w, "Version: %d.%d\n", hdr.VersionMajor, hdr.VersionMinor)
	fmt.Fprintf(w, "SystemIdentifier: %s\n", hdr.SystemIdentifier)
	fmt.Fprintf(w, "GeneratingSoftare: %s\n", hdr.GeneratingSoftware)
	fmt.Fprintf(w, "HeaderSize: %d\n", hdr.HeaderSize)
	fmt.Fprintf(w, "OffsetToPointData: %d\n", hdr.OffsetToPointData)
	fmt.Fprintf(w, "NumberOfVariableLengthRecords: %d\n", hdr.NumberOfVariableLengthRecords)
}

// BinaryReader is a helper for reading binary data
type BinaryReader struct {
	r     io.Reader
	Error error
}

// LasReader is a reader for .las files
type LasReader struct {
	r      io.Reader
	Header *LasPublicHeader
	Error  error
}

// NewLasReader creates a LasReader
func NewLasReader(r io.Reader) *LasReader {
	return &LasReader{
		r: r,
	}
}

// ReadHeader reads public header
func (r *LasReader) ReadHeader() (*LasPublicHeader, error) {
	r.Header, r.Error = ReadLasPublicHeader(r.r)
	return r.Header, r.Error
}

// NewBinaryReader creates a new binary reader
func NewBinaryReader(r io.Reader) *BinaryReader {
	return &BinaryReader{
		r: r,
	}
}

// ReadFixedString reads a fixed string of nChars characters
func (r *BinaryReader) ReadFixedString(nChars int) string {
	var res string
	if r.Error != nil {
		return res
	}
	data := make([]byte, nChars, nChars)
	n, err := r.r.Read(data[:])
	if err == nil && n != nChars {
		err = fmt.Errorf("ReadFixedString: wanted to read %d bytes, only read %d", nChars, n)
	}
	if err == nil {
		res = string(data)
	}
	r.Error = err
	return res
}

// Skip skips n bytes
func (r *BinaryReader) Skip(nBytes int) {
	r.ReadFixedString(nBytes)
}

// ReadUint8 reads a byte
func (r *BinaryReader) ReadUint8() byte {
	var res byte
	if r.Error != nil {
		return res
	}
	r.Error = binary.Read(r.r, binary.LittleEndian, &res)
	return res
}

// ReadUint16 reads uint16
func (r *BinaryReader) ReadUint16() uint16 {
	var res uint16
	if r.Error != nil {
		return res
	}
	r.Error = binary.Read(r.r, binary.LittleEndian, &res)
	return res
}

// ReadUint32 reads uint32
func (r *BinaryReader) ReadUint32() uint32 {
	var res uint32
	if r.Error != nil {
		return res
	}
	r.Error = binary.Read(r.r, binary.LittleEndian, &res)
	return res
}

// ReadFloat64 reads float64
func (r *BinaryReader) ReadFloat64() float64 {
	var res float64
	if r.Error != nil {
		return res
	}
	r.Error = binary.Read(r.r, binary.LittleEndian, &res)
	return res
}

func runLas2Txt(path string) string {
	// it seems by default las2txt does: --parse xyz
	// docs: http://www.liblas.org/utilities/las2txt.html
	cmd := exec.Command("las2txt", "-i", path, "--stdout")
	d, err := cmd.CombinedOutput()
	fatalIfErr(err)
	return string(d)
}

func readLasFile(path string) {
	f, err := os.Open(path)
	fatalIfErr(err)
	defer f.Close()
	r := NewLasReader(f)
	hdr, err := r.ReadHeader()
	fatalIfErr(err)
	dumpHeader(hdr, os.Stdout)
}

// Testing decoding of .las files
// We run las2txt on .las file and compare the results with our rendering
// To install last2txt on mac: brew install liblas
// more info on liblas: http://www.liblas.org/
func main() {
	args := os.Args[1:]
	if len(args) != 1 {
		exePath := os.Args[0]
		name := filepath.Base(exePath)
		fmt.Printf("usage: %s <file.las>\n", name)
		os.Exit(1)
	}
	path := args[0]
	verifyFileExists(path)
	//las2txtOut := runLas2Txt(path)
	//fmt.Printf("%s", las2txtOut)
	readLasFile(path)
}
