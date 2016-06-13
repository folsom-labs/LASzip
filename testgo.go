package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"io"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
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

// we support 1.0, 1.1, 1.2, 1.3, 1.4
func validVersion(major, minor byte) bool {
	if major != 1 || minor > 4 {
		return false
	}
	return true
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
	if !validVersion(hdr.VersionMajor, hdr.VersionMinor) {
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

func formatPointsByReturn(d [5]uint32) string {
	s := ""
	for i := 0; i < 5; i++ {
		s += fmt.Sprintf("%d ", d[i])
	}
	return s[:len(s)-1]
}

// for easy testing, dump header like lassinfo tool (http://www.liblas.org/utilities/lasinfo.html
/* It looks like:
---------------------------------------------------------
  Header Summary
---------------------------------------------------------

  Version:                     1.2
  Source ID:                   0
  Reserved:                    0
  Project ID/GUID:             '00000000-0000-0000-0000-000000000000'
  System ID:                   ''
  Generating Software:         'TerraScan'
  File Creation Day/Year:      0/0
  Header Byte Size             227
  Data Offset:                 1220
  Header Padding:              0
  Number Var. Length Records:  3
  Point Data Format:           1
  Number of Point Records:     10653
  Compressed:                  False
  Number of Points by Return:  9079 1244 288 42 0
  Scale Factor X Y Z:          0.01000000000000 0.01000000000000 0.01000000000000
  Offset X Y Z:                -0.00 -0.00 -0.00
  Min X Y Z:                   635589.01 848886.45 406.59
  Max X Y Z:                   638994.75 853535.43 593.73
*/
func dumpHeaderLikeLasInfo(hdr *LasPublicHeader, w io.Writer) {
	fmt.Fprintf(w, `---------------------------------------------------------
  Header Summary
---------------------------------------------------------

`)
	fmt.Fprintf(w, "  %-28s %d.%d\n", "Version:", hdr.VersionMajor, hdr.VersionMinor)
	fmt.Fprintf(w, "  %-28s %d\n", "Source ID:", hdr.FileSourceID)
	fmt.Fprintf(w, "  %-28s %d\n", "Reserved:", hdr.GlobalEncoding)
	// TODO: read and format GUID
	fmt.Fprintf(w, "  %-28s '%s'\n", "Project ID/GUID:", "00000000-0000-0000-0000-000000000000")
	fmt.Fprintf(w, "  %-28s '%s'\n", "System ID:", hdr.SystemIdentifier)
	fmt.Fprintf(w, "  %-28s '%s'\n", "Generating Software:", hdr.GeneratingSoftware)
	fmt.Fprintf(w, "  %-28s %d/%d\n", "File Creation Day/Year:", hdr.FileCreationDayOfYear, hdr.FileCreationYear)
	fmt.Fprintf(w, "  %-28s %d\n", "Header Byte Size", hdr.HeaderSize)
	fmt.Fprintf(w, "  %-28s %d\n", "Data Offset:", hdr.OffsetToPointData)
	// TODO: what is this?
	fmt.Fprintf(w, "  %-28s %d\n", "Header Padding:", 0)
	fmt.Fprintf(w, "  %-28s %d\n", "Number Var. Length Records:", hdr.NumberOfVariableLengthRecords)
	fmt.Fprintf(w, "  %-28s %d\n", "Point Data Format:", hdr.PointDataFormatID)

	fmt.Fprintf(w, "  %-28s %d\n", "Number of Point Records:", hdr.NumberOfPointRecords)
	fmt.Fprintf(w, "  %-28s %s\n", "Compressed:", "False")
	fmt.Fprintf(w, "  %-28s %s\n", "Number of Points by Return:", formatPointsByReturn(hdr.NumberOfPointsByReturn))

	fmt.Fprintf(w, "  %-28s %.14f %.14f %.14f\n", "Scale Factor X Y Z:", hdr.XScaleFactor, hdr.YScaleFactor, hdr.ZScaleFactor)
	fmt.Fprintf(w, "  %-28s %.2f %.2f %.2f\n", "Offset X Y Z:", hdr.XOffset, hdr.YOffset, hdr.ZOffset)
	fmt.Fprintf(w, "  %-28s %.2f %.2f %.2f\n", "Min X Y Z:", hdr.MinX, hdr.MinY, hdr.MinZ)
	fmt.Fprintf(w, "  %-28s %.2f %.2f %.2f\n", "Max X Y Z:", hdr.MaxX, hdr.MaxY, hdr.MaxZ)
}

// BinaryReader is a helper for reading binary data
type BinaryReader struct {
	r             io.Reader
	BytesConsumed int
	Error         error
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
	br := NewBinaryReader(r.r)
	r.Header, r.Error = ReadLasPublicHeader(br)
	fmt.Printf("bytes consumed: %d\n", br.BytesConsumed)
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
		res = strings.TrimRight(res, "\000")
	}
	r.Error = err
	r.BytesConsumed += nChars
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
	r.BytesConsumed++
	return res
}

// ReadUint16 reads uint16
func (r *BinaryReader) ReadUint16() uint16 {
	var res uint16
	if r.Error != nil {
		return res
	}
	r.Error = binary.Read(r.r, binary.LittleEndian, &res)
	r.BytesConsumed += 2
	return res
}

// ReadUint32 reads uint32
func (r *BinaryReader) ReadUint32() uint32 {
	var res uint32
	if r.Error != nil {
		return res
	}
	r.Error = binary.Read(r.r, binary.LittleEndian, &res)
	r.BytesConsumed += 4
	return res
}

// ReadFloat64 reads float64
func (r *BinaryReader) ReadFloat64() float64 {
	var res float64
	if r.Error != nil {
		return res
	}
	r.Error = binary.Read(r.r, binary.LittleEndian, &res)
	r.BytesConsumed += 8
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

func runLasInfo(path string) string {
	// docs: http://www.liblas.org/utilities/lasinfo.html
	cmd := exec.Command("lasinfo", "-i", path)
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
	dumpHeaderLikeLasInfo(hdr, os.Stdout)
}

func dumpHexLine(s string) {
	n := len(s)
	if n == 0 {
		return
	}
	line := ""
	for i := 0; i < n; i++ {
		line += fmt.Sprintf("%02x ", s[i])
	}
	fmt.Printf("%s\n", line)
}

func dumpHex(s string, nPerLine int) {
	for len(s) > 0 {
		n := len(s)
		if n > nPerLine {
			n = nPerLine
		}
		dumpHexLine(s[:n])
		s = s[n:]
	}
}

func trimEmptyStringsRight(a []string) []string {
	for {
		n := len(a) - 1
		if n < 0 || len(a[n]) > 0 {
			return a
		}
		a = a[:n]
	}
}

func compareLasInfo(sLasInfo, sMe string) {
	linesLasInfo := trimEmptyStringsRight(strings.Split(sLasInfo, "\n"))
	linesMe := trimEmptyStringsRight(strings.Split(sMe, "\n"))

	n := len(linesLasInfo)
	if len(linesMe) < n {
		n = len(linesMe)
	}
	for i := 0; i < n; i++ {
		lineLasInfo := linesLasInfo[i]
		lineMe := linesMe[i]
		lineLasInfoStripped := strings.TrimSpace(lineLasInfo)
		lineMeStripped := strings.TrimSpace(lineMe)

		if lineLasInfoStripped != lineMeStripped {
			fmt.Printf("lines %d are different\n", i+1)
			//fmt.Printf("%s: lassinfo\n", lineLasInfo)
			//fmt.Printf("%s: me\n", lineMe)
			//dumpHex(lineLasInfo, 8)
			//dumpHex(lineMe, 8)
			fmt.Printf("%s: lassinfo stripped\n", lineLasInfoStripped)
			fmt.Printf("%s: me stripped\n", lineMeStripped)
		} else {
			fmt.Printf("lines %d are same: '%s'\n", i+1, lineMe)
		}
	}
}

func getLasInfoCompatibleOutput(path string) string {
	f, err := os.Open(path)
	fatalIfErr(err)
	defer f.Close()
	r := NewLasReader(f)
	hdr, err := r.ReadHeader()
	fatalIfErr(err)
	var buf bytes.Buffer
	dumpHeaderLikeLasInfo(hdr, &buf)
	return string(buf.Bytes())
}

func compareLassInfoOutput(path string) {
	lasInfoOut := runLasInfo(path)
	meOut := getLasInfoCompatibleOutput(path)
	compareLasInfo(lasInfoOut, meOut)
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
	compareLassInfoOutput(path)
	//las2txtOut := runLas2Txt(path)
	//fmt.Printf("%s", las2txtOut)
	//readLasFile(path)
}
