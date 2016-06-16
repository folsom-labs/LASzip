package main

import (
	"bytes"
	"flag"
	"fmt"
	"io"
	"os"
	"os/exec"
	"strings"
)

/*
usage: testgo [-show-header] [-show-points] [-compare-points] <file.las>

-show-header
		print header information to stdout
show-lasinfo-header
		print header information to stdout in the same format as lasinfo
-show-points
		print x,y,z points to stdout
-compare-with-las2txt
		compare output of points with result of running las2txt --stdout
-compare-with-lasinfo
		compare output with result of running lasinfo

If no flag given, defaults to -show-header.

To install last2txt on mac: brew install liblas
More info on liblas: http://www.liblas.org/
*/

var (
	flgShowHeader         bool
	flgShowLasInfoHeader  bool
	flgShowPoints         bool
	flgCompareWithLas2Txt bool
	flgCompareWithLasInfo bool
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

func trimEmptyStringsRight(a []string) []string {
	for {
		n := len(a) - 1
		if n < 0 || len(a[n]) > 0 {
			return a
		}
		a = a[:n]
	}
}

func splitStringIntoLines(s string) []string {
	return trimEmptyStringsRight(strings.Split(s, "\n"))
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

/*
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
func dumpLasHeaderSummary(r *LasReader, w io.Writer) {
	hdr := r.Header
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
	// TODO: what is it? true for .laz files?
	fmt.Fprintf(w, "  %-28s %s\n", "Compressed:", "False")
	fmt.Fprintf(w, "  %-28s %s\n", "Number of Points by Return:", formatPointsByReturn(hdr.NumberOfPointsByReturn))

	fmt.Fprintf(w, "  %-28s %.14f %.14f %.14f\n", "Scale Factor X Y Z:", hdr.XScaleFactor, hdr.YScaleFactor, hdr.ZScaleFactor)
	fmt.Fprintf(w, "  %-28s %.2f %.2f %.2f\n", "Offset X Y Z:", hdr.XOffset, hdr.YOffset, hdr.ZOffset)
	fmt.Fprintf(w, "  %-28s %.2f %.2f %.2f\n", "Min X Y Z:", hdr.MinX, hdr.MinY, hdr.MinZ)
	fmt.Fprintf(w, "  %-28s %.2f %.2f %.2f\n", "Max X Y Z:", hdr.MaxX, hdr.MaxY, hdr.MaxZ)
}

/*
---------------------------------------------------------
	VLR Summary
---------------------------------------------------------
		User: 'LASF_Projection' - Description: 'GeoTIFF GeoKeyDirectoryTag'
		ID: 34735 Length: 64 Total Size: 118
		User: 'LASF_Projection' - Description: 'GeoTIFF GeoAsciiParamsTag'
		ID: 34737 Length: 47 Total Size: 101
		User: 'liblas' - Description: 'OGR variant of OpenGIS WKT SRS'
		ID: 2112 Length: 720 Total Size: 774
*/
func dumpLasVLRSummary(r *LasReader, w io.Writer) {
	fmt.Fprintf(w, `
---------------------------------------------------------
  VLR Summary
---------------------------------------------------------
`)
	for _, vlr := range r.VariableLengthRecords {
		fmt.Fprintf(w, "    User: '%s' - Description: '%s'\n", vlr.UserID, vlr.Description)
		fmt.Fprintf(w, "    ID: %d Length: %d Total Size: %d\n", vlr.RecordID, vlr.RecordLengthAfterHeader, vlr.RecordLengthAfterHeader+54)
	}
}

func dumpLasDimensions(r *LasReader, w io.Writer) {
	hdr := r.Header
	if hdr.PointDataFormatID == 1 {
		fmt.Fprint(w, `
  Dimensions
---------------------------------------------------------
  'X'                            --  size: 32 offset: 0
  'Y'                            --  size: 32 offset: 4
  'Z'                            --  size: 32 offset: 8
  'Intensity'                    --  size: 16 offset: 12
  'Return Number'                --  size: 3 offset: 14
  'Number of Returns'            --  size: 3 offset: 14
  'Scan Direction'               --  size: 1 offset: 14
  'Flightline Edge'              --  size: 1 offset: 14
  'Classification'               --  size: 8 offset: 15
  'Scan Angle Rank'              --  size: 8 offset: 16
  'User Data'                    --  size: 8 offset: 17
  'Point Source ID'              --  size: 16 offset: 18
  'Time'                         --  size: 64 offset: 20

`)
	}
}

// for easy testing, dump header like lassinfo tool (http://www.liblas.org/utilities/lasinfo.html
func dumpLikeLasInfo(r *LasReader, w io.Writer) {
	dumpLasHeaderSummary(r, w)
	dumpLasVLRSummary(r, w)
	dumpLasDimensions(r, w)
}

func runLas2Txt(path string) []string {
	// it seems by default las2txt does: --parse xyz
	// docs: http://www.liblas.org/utilities/las2txt.html
	cmd := exec.Command("las2txt", "-i", path, "--stdout")
	d, err := cmd.CombinedOutput()
	fatalIfErr(err)
	return splitStringIntoLines(string(d))
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
	dumpLikeLasInfo(r, os.Stdout)
}

func showLasInfo(path string, showHeader, showPoints bool) {
	f, err := os.Open(path)
	fatalIfErr(err)
	defer f.Close()
	r := NewLasReader(f)
	err = r.ReadHeaders()
	fatalIfErr(err)

	w := os.Stdout
	if showHeader {
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

	if showPoints {
		nPoints := int(r.Header.NumberOfPointRecords)
		if false && nPoints > 10 {
			nPoints = 10
		}
		for i := 0; i < nPoints; i++ {
			p, err := r.ReadPoint0(i)
			fatalIfErr(err)
			x, y, z := r.TransformPoints(p.X, p.Y, p.Z)
			fmt.Fprintf(w, "%.2f,%.2f,%.2f\n", x, y, z)
		}
	}
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

func getLasInfoCompatibleOutput(path string) string {
	f, err := os.Open(path)
	fatalIfErr(err)
	defer f.Close()
	r := NewLasReader(f)
	err = r.ReadHeaders()
	fatalIfErr(err)
	var buf bytes.Buffer
	dumpLikeLasInfo(r, &buf)
	return string(buf.Bytes())
}

func strTrimmedArrayRemove(a []string, s string) ([]string, bool) {
	for i, s2 := range a {
		s2 = strings.TrimSpace(s2)
		if s == s2 {
			// remove element at index i
			a = append(a[:i], a[i+1:]...)
			return a, true
		}
	}
	return a, false
}

func compareWithLasInfo(path string) {
	lasInfoOut := runLasInfo(path)
	meOut := getLasInfoCompatibleOutput(path)

	linesLasInfo := strings.Split(lasInfoOut, "\n")
	linesMe := strings.Split(meOut, "\n")

	var commonLines, myUniqueLines []string

	var removed bool
	for _, s := range linesMe {
		sTrimmed := strings.TrimSpace(s)
		if len(sTrimmed) == 0 {
			continue
		}
		linesLasInfo, removed = strTrimmedArrayRemove(linesLasInfo, sTrimmed)

		if removed {
			commonLines = append(commonLines, s)
		} else {
			myUniqueLines = append(myUniqueLines, s)
		}
	}

	fmt.Printf("%d lines are the same\n", len(commonLines))

	// what's left in linesLasInfo are lines unique to it
	n := len(linesLasInfo)
	if n > 0 {
		fmt.Printf("\n%d lines unique to lasinfo:\n", n)
		for _, s := range linesLasInfo {
			fmt.Printf("%s\n", s)
		}
	}

	n = len(myUniqueLines)
	if n > 0 {
		fmt.Printf("\n%d lines unique to me:\n", n)
		for _, s := range myUniqueLines {
			fmt.Printf("%s\n", s)
		}
	}
}

func boolToNumStr(b bool) string {
	if b {
		return "1"
	}
	return "0"
}

func getPointsMe(path string) []string {
	f, err := os.Open(path)
	fatalIfErr(err)
	defer f.Close()
	r := NewLasReader(f)
	err = r.ReadHeaders()
	fatalIfErr(err)

	var res []string
	nPoints := int(r.Header.NumberOfPointRecords)
	for i := 0; i < nPoints; i++ {
		p, err := r.ReadPoint0(i)
		fatalIfErr(err)
		x, y, z := r.TransformPoints(p.X, p.Y, p.Z)
		a := p.ScanAngleRank
		i := p.Intensity
		n := p.NumberOfReturns
		r := p.ReturnNumber
		ps := p.PointSourceID
		e := boolToNumStr(p.EdgeOfFlightLine)
		d := boolToNumStr(p.ScanDirectionFlag)
		c := p.GetClassification()
		C := GetClassificationName(c)
		u := p.UserData
		s := fmt.Sprintf("%.2f,%.2f,%.2f,%d,%d,%d,%d,%d,%s,%s,%d,%s,%d", x, y, z, a, i, n, r, ps, e, d, c, C, u)
		res = append(res, s)
	}
	return res
}

/*
--parse arg format:
	x - x coordinate as a double
	y - y coordinate as a double
	z - z coordinate as a double
	X - x coordinate as unscaled integer
	Y - y coordinate as unscaled integer
	Z - z coordinate as unscaled integer
	a - scan angle
	i - intensity
	n - number of returns for given pulse
	r - number of this return
	p - point source ID
	e - edge of flight line
	d - direction of scan flag
	c - classification number
	C - classification name
	u - user data
	R - red channel of RGB color
	G - green channel of RGB color
	B - blue channel of RGB color
	M - vertex index number
*/
func compareWithLas2Txt(path string) {
	// docs: http://www.liblas.org/utilities/las2txt.html
	cmd := exec.Command("las2txt", "-i", path, "--stdout", "--parse", "xyzainrpedcCu")
	d, err := cmd.CombinedOutput()
	fatalIfErr(err)
	lasLines := splitStringIntoLines(string(d))
	meLines := getPointsMe(path)

	n := len(lasLines)
	if n != len(meLines) {
		fmt.Print("error: mismatched number of points\n")
		fmt.Printf("me     : %d\n", len(meLines))
		fmt.Printf("las2txt: %d\n", len(lasLines))
		os.Exit(1)
	}
	for i := 0; i < n; i++ {
		if meLines[i] != lasLines[i] {
			fmt.Printf("error: different result for point %d\n", i)
			fmt.Printf("me     : '%s'\n", meLines[i])
			fmt.Printf("las2txt: '%s'\n", lasLines[i])
			os.Exit(1)
		}
	}
	fmt.Printf("All %d points seem to be ok!\n", n)
}

func parseFlags() {
	flag.BoolVar(&flgShowHeader, "show-header", false, "print header information to stdout")
	flag.BoolVar(&flgShowLasInfoHeader, "show-lasinfo-header", false, "print header information in the same format as lasinfo")
	flag.BoolVar(&flgShowPoints, "show-points", false, "print x, y, z points to stdout")
	flag.BoolVar(&flgCompareWithLas2Txt, "compare-with-las2txt", false, "compare our output with las2txt")
	flag.BoolVar(&flgCompareWithLasInfo, "compare-with-lasinfo", false, "compare our output with lasinfo")
	flag.Parse()

	// default to -show-header if nothing else given
	if !flgShowPoints && !flgCompareWithLas2Txt && !flgCompareWithLasInfo {
		flgShowHeader = true
	}
}

func main() {
	parseFlags()
	args := flag.Args()
	if len(args) != 1 {
		flag.Usage()
		os.Exit(1)
	}
	path := args[0]
	verifyFileExists(path)

	if flgCompareWithLas2Txt {
		compareWithLas2Txt(path)
		// ignore other flags in this case
		return
	}

	if flgCompareWithLasInfo {
		compareWithLasInfo(path)
		// ignore other flags in this case
		return
	}

	showLasInfo(path, flgShowHeader, flgShowPoints)
}
