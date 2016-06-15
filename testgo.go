package main

import (
	"bytes"
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

func fatalIf(cond bool) {
	if cond {
		panic("condition failed")
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

func dumpHeader(w io.Writer, hdr *LasPublicHeader) {
	fmt.Fprintf(w, "Version: %d.%d\n", hdr.VersionMajor, hdr.VersionMinor)
	fmt.Fprintf(w, "FileSourceID: %d\n", hdr.FileSourceID)
	fmt.Fprintf(w, "SystemIdentifier: %s\n", hdr.SystemIdentifier)
	fmt.Fprintf(w, "GeneratingSoftare: %s\n", hdr.GeneratingSoftware)
	fmt.Fprintf(w, "HeaderSize: %d\n", hdr.HeaderSize)
	fmt.Fprintf(w, "OffsetToPointData: %d\n", hdr.OffsetToPointData)
	fmt.Fprintf(w, "NumberOfVariableLengthRecords: %d\n", hdr.NumberOfVariableLengthRecords)

	fmt.Fprintf(w, "PointDataFormatID: %d\n", hdr.PointDataFormatID)

	fmt.Fprintf(w, "NumberOfPointRecords: %d\n", hdr.NumberOfPointRecords)
	fmt.Fprintf(w, "NumerOfPointsByReturn: %s\n", formatPointsByReturn(hdr.NumberOfPointsByReturn))
	fmt.Fprintf(w, "Scale factor X Y Z: %.14f %.14f %.14f\n", hdr.XScaleFactor, hdr.YScaleFactor, hdr.ZScaleFactor)
	fmt.Fprintf(w, "Offset X Y Z: %.2f %.2f %.2f\n", hdr.XOffset, hdr.YOffset, hdr.ZOffset)
	fmt.Fprintf(w, "Min X Y Z: %.2f %.2f %.2f\n", hdr.MinX, hdr.MinY, hdr.MinZ)
	fmt.Fprintf(w, "Max X Y Z: %.2f %.2f %.2f\n", hdr.MaxX, hdr.MaxY, hdr.MaxZ)
}

func dumpVariableLengthHeader(w io.Writer, hdr *VariableLengthRecord) {
	fmt.Fprintf(w, "UserID: %s\n", hdr.UserID)
	fmt.Fprintf(w, "RecordID: %d\n", hdr.RecordID)
	fmt.Fprintf(w, "RecordLengthAfterHeader: %d\n", hdr.RecordLengthAfterHeader)
	fmt.Fprintf(w, "Description: %s\n", hdr.Description)
}

func dumpGeoKeyDirectory(w io.Writer, d *GeoKeyDirectory) {
	fmt.Fprintf(w, "Number of keys: %d\n", d.NumberOfKeys)
	for _, k := range d.KeysRaw {
		fmt.Fprintf(w, "id: %d, tiff loc: %d, count: %d offset: %d\n", k.KeyID, k.TIFFTagLocation, k.Count, k.ValueOrOffset)
	}
}

func dumpGeoTags(w io.Writer, tags *GeoTags) {
	fmt.Fprintf(w, "%d short keys:\n", len(tags.TagsShort))
	for _, tag := range tags.TagsShort {
		fmt.Fprintf(w, "%d %s: %d\n", tag.TagID, tag.Name, tag.Value)
	}
	fmt.Fprintf(w, "%d double keys:\n", len(tags.TagsDouble))
	for _, tag := range tags.TagsDouble {
		fmt.Fprintf(w, "%d %s: %.4f\n", tag.TagID, tag.Name, tag.Value)
	}
	fmt.Fprintf(w, "%d string keys:\n", len(tags.TagsString))
	for _, tag := range tags.TagsString {
		fmt.Fprintf(w, "%d %s: %s\n", tag.TagID, tag.Name, tag.Value)
	}
}

func dumpASCIIParams(w io.Writer, params []byte) {
	strings := bytes.Split(params, []byte{0})
	for _, s := range strings {
		fmt.Fprintf(w, "%s\n", string(s))
	}
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
	err = r.ReadHeaders()
	fatalIfErr(err)
	dumpHeaderLikeLasInfo(r.Header, os.Stdout)
}

func readLasFile2(path string) {
	f, err := os.Open(path)
	fatalIfErr(err)
	defer f.Close()
	r := NewLasReader(f)
	err = r.ReadHeaders()
	fatalIfErr(err)

	w := os.Stdout
	dumpHeader(w, r.Header)
	for _, record := range r.VariableLengthRecords {
		fmt.Fprint(w, "\n")
		dumpVariableLengthHeader(w, record)
	}
	if r.GeoKeyInfo.ASCIIParams != nil {
		fmt.Fprint(w, "\nGeoASCIIParams:\n")
		dumpASCIIParams(w, r.GeoKeyInfo.ASCIIParams)
	}

	if r.GeoKeyInfo.Directory != nil {
		fmt.Fprint(w, "\nGeoKeyDirectory:\n")
		dumpGeoKeyDirectory(w, r.GeoKeyInfo.Directory)
	}

	fmt.Fprint(w, "\nGeoTags:\n")
	dumpGeoTags(w, r.GeoTags)
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
	err = r.ReadHeaders()
	fatalIfErr(err)
	var buf bytes.Buffer
	dumpHeaderLikeLasInfo(r.Header, &buf)
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
	//compareLassInfoOutput(path)
	readLasFile2(path)

	//las2txtOut := runLas2Txt(path)
	//fmt.Printf("%s", las2txtOut)
}
