package main

import (
	"errors"
	"fmt"
	"strings"
)

/*
This doesn't actually implement GeoTIFF format but .las lidar files
use the same tags as GeoTIFF spec.

Tags have unique numeric ids. Their values can be: short (uint16),
double (float64) or string.

For convenience, we also put tag name as a string.

Spec: http://www.remotesensing.org/geotiff/spec/contents.html
*/

// GeoTagShort describes a tag whose value is short (uint16)
type GeoTagShort struct {
	TagID int
	Name  string
	Value uint16
}

// GeoTagDouble describes a tag whose value is double (float64)
type GeoTagDouble struct {
	TagID int
	Name  string
	Value float64
}

// GeoTagString describes a tag whose value is string
type GeoTagString struct {
	TagID int
	Name  string
	Value string
}

// GeoTags represents decoded GeoKey data
type GeoTags struct {
	TagsShort  []*GeoTagShort
	TagsDouble []*GeoTagDouble
	TagsString []*GeoTagString
	Tags       []interface{}
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
	Count         uint16
	ValueOrOffset uint16
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

// GeoKeyInfo describes all info needed to
type GeoKeyInfo struct {
	Directory    *GeoKeyDirectory
	DoubleParams []float64
	// array of ascii data. null-separated strings, referenced  by
	// GeoKey
	ASCIIParams []byte
}

// GeoKeyKnownValueName returns name of the value for known
// key / vallue pairs
func GeoKeyKnownValueName(geoKeyID int, val uint16) string {
	switch geoKeyID {
	case GTModelTypeGeoKey:
		return ModelTypeName(ModelType(val))
	case GTRasterTypeGeoKey:
		return RasterTypeName(RasterType(val))
	case GeogAngularUnitsGeoKey:
		return AngularUnitName(AngularUnit(val))
	case ProjLinearUnitsGeoKey:
		return LinearUnitName(LinearUnit(val))
	case GeogGeodeticDatumGeoKey:
		return DatumName(int(val))
	case GeographicTypeGeoKey:
		return GcsName(int(val))
	case GeogPrimeMeridianGeoKey:
		return PrimeMeridianName(int(val))
	case GeogEllipsoidGeoKey:
		return EllipseName(int(val))
	case ProjectedCSTypeGeoKey:
		return PcsName(int(val))
	}
	return fmt.Sprintf("Unknown-%d", val)
}

// DecodeGeoKeyInfo decodes geo information
func DecodeGeoKeyInfo(geoInfo *GeoKeyInfo) (*GeoTags, error) {
	geoDir := geoInfo.Directory
	if geoDir == nil {
		return nil, errors.New("missing GeoKeyDirectoryTag record")
	}
	if geoDir.KeyDirectoryVersion != 1 {
		return nil, fmt.Errorf("GeoKeyDirectory.KeyDirectoryVersion is %d, expected 1", geoDir.KeyDirectoryVersion)
	}
	if geoDir.KeyRevision != 1 {
		return nil, fmt.Errorf("GeoKeyDirectory.KeyRevision is %d, expected 1", geoDir.KeyRevision)
	}
	if geoDir.MinorRevision != 0 {
		return nil, fmt.Errorf("GeoKeyDirectory.MinorRevision is %d, expected 0", geoDir.MinorRevision)
	}

	var res GeoTags
	for _, key := range geoDir.KeysRaw {
		loc := key.TIFFTagLocation
		switch loc {
		case 0:
			var v GeoTagShort
			v.TagID = int(key.KeyID)
			v.Name = TagIDToName(v.TagID)
			v.Value = key.ValueOrOffset
			res.TagsShort = append(res.TagsShort, &v)
			res.Tags = append(res.Tags, &v)
		case GeoDoubleParamsTag:
			if geoInfo.DoubleParams == nil {
				return nil, fmt.Errorf("key location in double params but GeoDoubleParamsTag record not present")
			}
			idx := int(key.ValueOrOffset)
			if idx >= len(geoInfo.DoubleParams) {
				return nil, fmt.Errorf("idx %d outside of len(geInfo.DoubleParams) (%d)", idx, len(geoInfo.DoubleParams))
			}
			var v GeoTagDouble
			v.TagID = int(key.KeyID)
			v.Name = TagIDToName(v.TagID)
			v.Value = geoInfo.DoubleParams[idx]
			res.TagsDouble = append(res.TagsDouble, &v)
			res.Tags = append(res.Tags, &v)
		case GeoASCIIParamsTag:
			if geoInfo.ASCIIParams == nil {
				return nil, fmt.Errorf("key location in ASCII params but GeoASCIIParamsTag record not present")
			}
			idx := int(key.ValueOrOffset)
			n := int(key.Count)
			if idx+n > len(geoInfo.ASCIIParams) {
				return nil, fmt.Errorf("idx+len (%d+%d=%d) outside of len(geInfo.ASCIIParams) (%d)", idx, n, idx+n, len(geoInfo.ASCIIParams))
			}
			var v GeoTagString
			v.TagID = int(key.KeyID)
			v.Name = TagIDToName(v.TagID)
			s := string(geoInfo.ASCIIParams[idx : idx+n])
			// geotiff replaces terminating 0 with |
			v.Value = strings.TrimRight(s, "|")
			res.TagsString = append(res.TagsString, &v)
			res.Tags = append(res.Tags, &v)
		}
	}

	return &res, nil
}

// FindGeoTagShort returns GeoTagShort with a given tagID
func (tags *GeoTags) FindGeoTagShort(tagID int) *GeoTagShort {
	for _, t := range tags.TagsShort {
		if t.TagID == tagID {
			return t
		}
	}
	return nil
}
