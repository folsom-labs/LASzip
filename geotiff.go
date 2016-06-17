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

//
// http://www.remotesensing.org/geotiff/spec/geotiff2.7.html#2.7
// https://github.com/smanders/libgeotiff/blob/sdl_1_2_4/geokeys.inc
const (
	/* 6.2.1 GeoTIFF Configuration Keys */
	GTModelTypeGeoKey  = 1024 /* Section 6.3.1.1 Codes       */
	GTRasterTypeGeoKey = 1025 /* Section 6.3.1.2 Codes       */
	GTCitationGeoKey   = 1026 /* documentation */

	/* 6.2.2 Geographic CS Parameter Keys */
	GeographicTypeGeoKey        = 2048 /* Section 6.3.2.1 Codes     */
	GeogCitationGeoKey          = 2049 /* documentation             */
	GeogGeodeticDatumGeoKey     = 2050 /* Section 6.3.2.2 Codes     */
	GeogPrimeMeridianGeoKey     = 2051 /* Section 6.3.2.4 codes     */
	GeogLinearUnitsGeoKey       = 2052 /* Section 6.3.1.3 Codes     */
	GeogLinearUnitSizeGeoKey    = 2053 /* meters                    */
	GeogAngularUnitsGeoKey      = 2054 /* Section 6.3.1.4 Codes     */
	GeogAngularUnitSizeGeoKey   = 2055 /* radians                   */
	GeogEllipsoidGeoKey         = 2056 /* Section 6.3.2.3 Codes     */
	GeogSemiMajorAxisGeoKey     = 2057 /* GeogLinearUnits           */
	GeogSemiMinorAxisGeoKey     = 2058 /* GeogLinearUnits           */
	GeogInvFlatteningGeoKey     = 2059 /* ratio                     */
	GeogAzimuthUnitsGeoKey      = 2060 /* Section 6.3.1.4 Codes     */
	GeogPrimeMeridianLongGeoKey = 2061 /* GeoAngularUnit            */

	/* 6.2.3 Projected CS Parameter Keys */
	/*    Several keys have been renamed,*/
	/*    and the deprecated names aliased for backward compatibility */
	ProjectedCSTypeGeoKey          = 3072                       /* Section 6.3.3.1 codes */
	PCSCitationGeoKey              = 3073                       /* documentation */
	ProjectionGeoKey               = 3074                       /* Section 6.3.3.2 codes */
	ProjCoordTransGeoKey           = 3075                       /* Section 6.3.3.3 codes */
	ProjLinearUnitsGeoKey          = 3076                       /* Section 6.3.1.3 codes */
	ProjLinearUnitSizeGeoKey       = 3077                       /* meters */
	ProjStdParallel1GeoKey         = 3078                       /* GeogAngularUnit */
	ProjStdParallelGeoKey          = ProjStdParallel1GeoKey     /* ** alias ** */
	ProjStdParallel2GeoKey         = 3079                       /* GeogAngularUnit */
	ProjNatOriginLongGeoKey        = 3080                       /* GeogAngularUnit */
	ProjOriginLongGeoKey           = ProjNatOriginLongGeoKey    /* ** alias ** */
	ProjNatOriginLatGeoKey         = 3081                       /* GeogAngularUnit */
	ProjOriginLatGeoKey            = ProjNatOriginLatGeoKey     /* ** alias ** */
	ProjFalseEastingGeoKey         = 3082                       /* ProjLinearUnits */
	ProjFalseNorthingGeoKey        = 3083                       /* ProjLinearUnits */
	ProjFalseOriginLongGeoKey      = 3084                       /* GeogAngularUnit */
	ProjFalseOriginLatGeoKey       = 3085                       /* GeogAngularUnit */
	ProjFalseOriginEastingGeoKey   = 3086                       /* ProjLinearUnits */
	ProjFalseOriginNorthingGeoKey  = 3087                       /* ProjLinearUnits */
	ProjCenterLongGeoKey           = 3088                       /* GeogAngularUnit */
	ProjCenterLatGeoKey            = 3089                       /* GeogAngularUnit */
	ProjCenterEastingGeoKey        = 3090                       /* ProjLinearUnits */
	ProjCenterNorthingGeoKey       = 3091                       /* ProjLinearUnits */
	ProjScaleAtNatOriginGeoKey     = 3092                       /* ratio */
	ProjScaleAtOriginGeoKey        = ProjScaleAtNatOriginGeoKey /* ** alias ** */
	ProjScaleAtCenterGeoKey        = 3093                       /* ratio */
	ProjAzimuthAngleGeoKey         = 3094                       /* GeogAzimuthUnit */
	ProjStraightVertPoleLongGeoKey = 3095                       /* GeogAngularUnit */
	ProjRectifiedGridAngleGeoKey   = 3096                       /* GeogAngularUnit */

	/* 6.2.4 Vertical CS Keys */
	VerticalCSTypeGeoKey   = 4096 /* Section 6.3.4.1 codes   */
	VerticalCitationGeoKey = 4097 /* documentation */
	VerticalDatumGeoKey    = 4098 /* Section 6.3.4.2 codes   */
	VerticalUnitsGeoKey    = 4099 /* Section 6.3.1 (.x) codes   */
)

// TagIDToName returns a string version of tagID value
func TagIDToName(tagID int) string {

	switch tagID {

	case GTModelTypeGeoKey:
		return "GTModelTypeGeoKey"
	case GTRasterTypeGeoKey:
		return "GTRasterTypeGeoKey"
	case GTCitationGeoKey:
		return "GTCitationGeoKey"
	case GeographicTypeGeoKey:
		return "GeographicTypeGeoKey"
	case GeogCitationGeoKey:
		return "GeogCitationGeoKey"
	case GeogGeodeticDatumGeoKey:
		return "GeogGeodeticDatumGeoKey"
	case GeogPrimeMeridianGeoKey:
		return "GeogPrimeMeridianGeoKey"
	case GeogLinearUnitsGeoKey:
		return "GeogLinearUnitsGeoKey"
	case GeogLinearUnitSizeGeoKey:
		return "GeogLinearUnitSizeGeoKey"
	case GeogAngularUnitsGeoKey:
		return "GeogAngularUnitsGeoKey"
	case GeogAngularUnitSizeGeoKey:
		return "GeogAngularUnitSizeGeoKey"
	case GeogEllipsoidGeoKey:
		return "GeogEllipsoidGeoKey"
	case GeogSemiMajorAxisGeoKey:
		return "GeogSemiMajorAxisGeoKey"
	case GeogSemiMinorAxisGeoKey:
		return "GeogSemiMinorAxisGeoKey"
	case GeogInvFlatteningGeoKey:
		return "GeogInvFlatteningGeoKey"
	case GeogAzimuthUnitsGeoKey:
		return "GeogAzimuthUnitsGeoKey"
	case GeogPrimeMeridianLongGeoKey:
		return "GeogPrimeMeridianLongGeoKey"
	case ProjectedCSTypeGeoKey:
		return "ProjectedCSTypeGeoKey"
	case PCSCitationGeoKey:
		return "PCSCitationGeoKey"
	case ProjectionGeoKey:
		return "ProjectionGeoKey"
	case ProjCoordTransGeoKey:
		return "ProjCoordTransGeoKey"
	case ProjLinearUnitsGeoKey:
		return "ProjLinearUnitsGeoKey"
	case ProjLinearUnitSizeGeoKey:
		return "ProjLinearUnitSizeGeoKey"
	case ProjStdParallel1GeoKey:
		return "ProjStdParallel1GeoKey"
	case ProjStdParallel2GeoKey:
		return "ProjStdParallel2GeoKey"
	case ProjNatOriginLongGeoKey:
		return "ProjNatOriginLongGeoKey"
	case ProjNatOriginLatGeoKey:
		return "ProjNatOriginLatGeoKey"
	case ProjFalseEastingGeoKey:
		return "ProjFalseEastingGeoKey"
	case ProjFalseNorthingGeoKey:
		return "ProjFalseNorthingGeoKey"
	case ProjFalseOriginLongGeoKey:
		return "ProjFalseOriginLongGeoKey"
	case ProjFalseOriginLatGeoKey:
		return "ProjFalseOriginLatGeoKey"
	case ProjFalseOriginEastingGeoKey:
		return "ProjFalseOriginEastingGeoKey"
	case ProjFalseOriginNorthingGeoKey:
		return "ProjFalseOriginNorthingGeoKey"
	case ProjCenterLongGeoKey:
		return "ProjCenterLongGeoKey"
	case ProjCenterLatGeoKey:
		return "ProjCenterLatGeoKey"
	case ProjCenterEastingGeoKey:
		return "ProjCenterEastingGeoKey"
	case ProjCenterNorthingGeoKey:
		return "ProjCenterNorthingGeoKey"
	case ProjScaleAtNatOriginGeoKey:
		return "ProjScaleAtNatOriginGeoKey"
	case ProjScaleAtCenterGeoKey:
		return "ProjScaleAtCenterGeoKey"
	case ProjAzimuthAngleGeoKey:
		return "ProjAzimuthAngleGeoKey"
	case ProjStraightVertPoleLongGeoKey:
		return "ProjStraightVertPoleLongGeoKey"
	case ProjRectifiedGridAngleGeoKey:
		return "ProjRectifiedGridAngleGeoKey"
	case VerticalCSTypeGeoKey:
		return "VerticalCSTypeGeoKey"
	case VerticalCitationGeoKey:
		return "VerticalCitationGeoKey"
	case VerticalDatumGeoKey:
		return "VerticalDatumGeoKey"
	case VerticalUnitsGeoKey:
		return "VerticalUnitsGeoKey"
	}
	return fmt.Sprintf("unknown tagID (%d)", tagID)
}

// 6.3.1.1 Model Type Codes for GTModelTypeGeoKey
// https://github.com/smanders/libgeotiff/blob/master/geovalues.h#L54
type ModelType int

const (
	ModelTypeProjected  ModelType = 1
	ModelTypeGeographic ModelType = 2
	ModelTypeGeocentric ModelType = 3
)

// ModelTypeName returns a name for a given ModelType constant
func ModelTypeName(v ModelType) string {
	switch v {
	case ModelTypeProjected:
		return "ModelTypeProjected"
	case ModelTypeGeographic:
		return "ModelTypeGeographic"
	case ModelTypeGeocentric:
		return "ModelTypeGeocentric"
	}
	return fmt.Sprintf("Unknown ModelType value (%d)", int(v))
}

// 6.3.1.2 Raster Type Codes for GTRasterTypeGeoKey
// https://github.com/smanders/libgeotiff/blob/master/geovalues.h#L64
type RasterType int

const (
	RasterPixelIsArea  RasterType = 1
	RasterPixelIsPoint RasterType = 2
)

// RasterTypeName returns a name for a given RasterType constant
func RasterTypeName(v RasterType) string {
	switch v {
	case RasterPixelIsArea:
		return "RasterPixelIsArea"
	case RasterPixelIsPoint:
		return "RasterPixelIsPoint"
	}
	return fmt.Sprintf("Unkown RasterType value (%d)", int(v))
}

// https://github.com/smanders/libgeotiff/blob/master/epsg_units.inc#L27
type AngularUnit int

const (
	Angular_Radian         AngularUnit = 9101
	Angular_Degree         AngularUnit = 9102
	Angular_Arc_Minute     AngularUnit = 9103
	Angular_Arc_Second     AngularUnit = 9104
	Angular_Grad           AngularUnit = 9105
	Angular_Gon            AngularUnit = 9106
	Angular_DMS            AngularUnit = 9107
	Angular_DMS_Hemisphere AngularUnit = 9108
)

// AngularUnitName retuns a name for a given AngularUnit constant
func AngularUnitName(v AngularUnit) string {
	switch v {
	case Angular_Radian:
		return "Angular_Radian"
	case Angular_Degree:
		return "Angular_Degree"
	case Angular_Arc_Minute:
		return "Angular_Arc_Minute"
	case Angular_Arc_Second:
		return "Angular_Arc_Second"
	case Angular_Grad:
		return "Angular_Grad"
	case Angular_Gon:
		return "Angular_Gon"
	case Angular_DMS:
		return "Angular_DMS"
	case Angular_DMS_Hemisphere:
		return "Angular_DMS_Hemisphere"
	}
	return fmt.Sprintf("Unkown AngularUnit value (%d)", int(v))
}

//https://github.com/smanders/libgeotiff/blob/master/epsg_units.inc#L9
type LinearUnit int

const (
	Linear_Meter                       LinearUnit = 9001
	Linear_Foot                        LinearUnit = 9002
	Linear_Foot_US_Survey              LinearUnit = 9003
	Linear_Foot_Modified_American      LinearUnit = 9004
	Linear_Foot_Clarke                 LinearUnit = 9005
	Linear_Foot_Indian                 LinearUnit = 9006
	Linear_Link                        LinearUnit = 9007
	Linear_Link_Benoit                 LinearUnit = 9008
	Linear_Link_Sears                  LinearUnit = 9009
	Linear_Chain_Benoit                LinearUnit = 9010
	Linear_Chain_Sears                 LinearUnit = 9011
	Linear_Yard_Sears                  LinearUnit = 9012
	Linear_Yard_Indian                 LinearUnit = 9013
	Linear_Fathom                      LinearUnit = 9014
	Linear_Mile_International_Nautical LinearUnit = 9015
)

// LinearUnitName retuns a name for a given LinearUnit constant
func LinearUnitName(v LinearUnit) string {
	switch v {
	case Linear_Meter:
		return "Linear_Meter"
	case Linear_Foot:
		return "Linear_Foot"
	case Linear_Foot_US_Survey:
		return "Linear_Foot_US_Survey"
	case Linear_Foot_Modified_American:
		return "Linear_Foot_Modified_American"
	case Linear_Foot_Clarke:
		return "Linear_Foot_Clarke"
	case Linear_Foot_Indian:
		return "Linear_Foot_Indian"
	case Linear_Link:
		return "Linear_Link"
	case Linear_Link_Benoit:
		return "Linear_Link_Benoit"
	case Linear_Link_Sears:
		return "Linear_Link_Sears"
	case Linear_Chain_Benoit:
		return "Linear_Chain_Benoit"
	case Linear_Chain_Sears:
		return "Linear_Chain_Sears"
	case Linear_Yard_Sears:
		return "Linear_Yard_Sears"
	case Linear_Yard_Indian:
		return "Linear_Yard_Indian"
	case Linear_Fathom:
		return "Linear_Fathom"
	case Linear_Mile_International_Nautical:
		return "Linear_Mile_International_Nautical"
	}
	return fmt.Sprintf("Unknown-%d", v)
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
