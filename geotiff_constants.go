package main

import "fmt"

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

// https://github.com/smanders/libgeotiff/blob/master/epsg_units.inc#L9
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

// https://github.com/smanders/libgeotiff/blob/master/epsg_gcs.inc#L13

// https://github.com/smanders/libgeotiff/blob/master/epsg_datum.inc#L13

const (
	Datum_Dealul_Piscului_1970 = 6317

	/* Datums for which only the ellipsoid is known */
	DatumE_Airy1830                   = 6001
	DatumE_AiryModified1849           = 6002
	DatumE_AustralianNationalSpheroid = 6003
	DatumE_Bessel1841                 = 6004
	DatumE_BesselModified             = 6005
	DatumE_BesselNamibia              = 6006
	DatumE_Clarke1858                 = 6007
	DatumE_Clarke1866                 = 6008
	DatumE_Clarke1866Michigan         = 6009
	DatumE_Clarke1880_Benoit          = 6010
	DatumE_Clarke1880_IGN             = 6011
	DatumE_Clarke1880_RGS             = 6012
	DatumE_Clarke1880_Arc             = 6013
	DatumE_Clarke1880_SGA1922         = 6014
	DatumE_Everest1830_1937Adjustment = 6015
	DatumE_Everest1830_1967Definition = 6016
	DatumE_Everest1830_1975Definition = 6017
	DatumE_Everest1830Modified        = 6018
	DatumE_GRS1980                    = 6019
	DatumE_Helmert1906                = 6020
	DatumE_IndonesianNationalSpheroid = 6021
	DatumE_International1924          = 6022
	DatumE_International1967          = 6023
	DatumE_Krassowsky1960             = 6024
	DatumE_NWL9D                      = 6025
	DatumE_NWL10D                     = 6026
	DatumE_Plessis1817                = 6027
	DatumE_Struve1860                 = 6028
	DatumE_WarOffice                  = 6029
	DatumE_WGS84                      = 6030
	DatumE_GEM10C                     = 6031
	DatumE_OSU86F                     = 6032
	DatumE_OSU91A                     = 6033
	DatumE_Clarke1880                 = 6034
	DatumE_Sphere                     = 6035

	/* standard datums */
	Datum_Adindan                            = 6201
	Datum_Australian_Geodetic_Datum_1966     = 6202
	Datum_Australian_Geodetic_Datum_1984     = 6203
	Datum_Ain_el_Abd_1970                    = 6204
	Datum_Afgooye                            = 6205
	Datum_Agadez                             = 6206
	Datum_Lisbon                             = 6207
	Datum_Aratu                              = 6208
	Datum_Arc_1950                           = 6209
	Datum_Arc_1960                           = 6210
	Datum_Batavia                            = 6211
	Datum_Barbados                           = 6212
	Datum_Beduaram                           = 6213
	Datum_Beijing_1954                       = 6214
	Datum_Reseau_National_Belge_1950         = 6215
	Datum_Bermuda_1957                       = 6216
	Datum_Bern_1898                          = 6217
	Datum_Bogota                             = 6218
	Datum_Bukit_Rimpah                       = 6219
	Datum_Camacupa                           = 6220
	Datum_Campo_Inchauspe                    = 6221
	Datum_Cape                               = 6222
	Datum_Carthage                           = 6223
	Datum_Chua                               = 6224
	Datum_Corrego_Alegre                     = 6225
	Datum_Cote_d_Ivoire                      = 6226
	Datum_Deir_ez_Zor                        = 6227
	Datum_Douala                             = 6228
	Datum_Egypt_1907                         = 6229
	Datum_European_Datum_1950                = 6230
	Datum_European_Datum_1987                = 6231
	Datum_Fahud                              = 6232
	Datum_Gandajika_1970                     = 6233
	Datum_Garoua                             = 6234
	Datum_Guyane_Francaise                   = 6235
	Datum_Hu_Tzu_Shan                        = 6236
	Datum_Hungarian_Datum_1972               = 6237
	Datum_Indonesian_Datum_1974              = 6238
	Datum_Indian_1954                        = 6239
	Datum_Indian_1975                        = 6240
	Datum_Jamaica_1875                       = 6241
	Datum_Jamaica_1969                       = 6242
	Datum_Kalianpur                          = 6243
	Datum_Kandawala                          = 6244
	Datum_Kertau                             = 6245
	Datum_Kuwait_Oil_Company                 = 6246
	Datum_La_Canoa                           = 6247
	Datum_Provisional_S_American_Datum_1956  = 6248
	Datum_Lake                               = 6249
	Datum_Leigon                             = 6250
	Datum_Liberia_1964                       = 6251
	Datum_Lome                               = 6252
	Datum_Luzon_1911                         = 6253
	Datum_Hito_XVIII_1963                    = 6254
	Datum_Herat_North                        = 6255
	Datum_Mahe_1971                          = 6256
	Datum_Makassar                           = 6257
	Datum_European_Reference_System_1989     = 6258
	Datum_Malongo_1987                       = 6259
	Datum_Manoca                             = 6260
	Datum_Merchich                           = 6261
	Datum_Massawa                            = 6262
	Datum_Minna                              = 6263
	Datum_Mhast                              = 6264
	Datum_Monte_Mario                        = 6265
	Datum_M_poraloko                         = 6266
	Datum_North_American_Datum_1927          = 6267
	Datum_NAD_Michigan                       = 6268
	Datum_North_American_Datum_1983          = 6269
	Datum_Nahrwan_1967                       = 6270
	Datum_Naparima_1972                      = 6271
	Datum_New_Zealand_Geodetic_Datum_1949    = 6272
	Datum_NGO_1948                           = 6273
	Datum_Datum_73                           = 6274
	Datum_Nouvelle_Triangulation_Francaise   = 6275
	Datum_NSWC_9Z_2                          = 6276
	Datum_OSGB_1936                          = 6277
	Datum_OSGB_1970_SN                       = 6278
	Datum_OS_SN_1980                         = 6279
	Datum_Padang_1884                        = 6280
	Datum_Palestine_1923                     = 6281
	Datum_Pointe_Noire                       = 6282
	Datum_Geocentric_Datum_of_Australia_1994 = 6283
	Datum_Pulkovo_1942                       = 6284
	Datum_Qatar                              = 6285
	Datum_Qatar_1948                         = 6286
	Datum_Qornoq                             = 6287
	Datum_Loma_Quintana                      = 6288
	Datum_Amersfoort                         = 6289
	Datum_RT38                               = 6290
	Datum_South_American_Datum_1969          = 6291
	Datum_Sapper_Hill_1943                   = 6292
	Datum_Schwarzeck                         = 6293
	Datum_Segora                             = 6294
	Datum_Serindung                          = 6295
	Datum_Sudan                              = 6296
	Datum_Tananarive_1925                    = 6297
	Datum_Timbalai_1948                      = 6298
	Datum_TM65                               = 6299
	Datum_TM75                               = 6300
	Datum_Tokyo                              = 6301
	Datum_Trinidad_1903                      = 6302
	Datum_Trucial_Coast_1948                 = 6303
	Datum_Voirol_1875                        = 6304
	Datum_Voirol_Unifie_1960                 = 6305
	Datum_Bern_1938                          = 6306
	Datum_Nord_Sahara_1959                   = 6307
	Datum_Stockholm_1938                     = 6308
	Datum_Yacare                             = 6309
	Datum_Yoff                               = 6310
	Datum_Zanderij                           = 6311
	Datum_Militar_Geographische_Institut     = 6312
	Datum_Reseau_National_Belge_1972         = 6313
	Datum_Deutsche_Hauptdreiecksnetz         = 6314
	Datum_Conakry_1905                       = 6315
	Datum_WGS72                              = 6322
	Datum_WGS72_Transit_Broadcast_Ephemeris  = 6324
	Datum_WGS84                              = 6326
	Datum_Ancienne_Triangulation_Francaise   = 6901
	Datum_Nord_de_Guerre                     = 6902
)

// DatumName returns a name for a given datum id
func DatumName(v int) string {
	switch v {
	case Datum_Dealul_Piscului_1970:
		return "Datum_Dealul_Piscului_1970"
	case DatumE_Airy1830:
		return "DatumE_Airy1830"
	case DatumE_AiryModified1849:
		return "DatumE_AiryModified1849"
	case DatumE_AustralianNationalSpheroid:
		return "DatumE_AustralianNationalSpheroid"
	case DatumE_Bessel1841:
		return "DatumE_Bessel1841"
	case DatumE_BesselModified:
		return "DatumE_BesselModified"
	case DatumE_BesselNamibia:
		return "DatumE_BesselNamibia"
	case DatumE_Clarke1858:
		return "DatumE_Clarke1858"
	case DatumE_Clarke1866:
		return "DatumE_Clarke1866"
	case DatumE_Clarke1866Michigan:
		return "DatumE_Clarke1866Michigan"
	case DatumE_Clarke1880_Benoit:
		return "DatumE_Clarke1880_Benoit"
	case DatumE_Clarke1880_IGN:
		return "DatumE_Clarke1880_IGN"
	case DatumE_Clarke1880_RGS:
		return "DatumE_Clarke1880_RGS"
	case DatumE_Clarke1880_Arc:
		return "DatumE_Clarke1880_Arc"
	case DatumE_Clarke1880_SGA1922:
		return "DatumE_Clarke1880_SGA1922"
	case DatumE_Everest1830_1937Adjustment:
		return "DatumE_Everest1830_1937Adjustment"
	case DatumE_Everest1830_1967Definition:
		return "DatumE_Everest1830_1967Definition"
	case DatumE_Everest1830_1975Definition:
		return "DatumE_Everest1830_1975Definition"
	case DatumE_Everest1830Modified:
		return "DatumE_Everest1830Modified"
	case DatumE_GRS1980:
		return "DatumE_GRS1980"
	case DatumE_Helmert1906:
		return "DatumE_Helmert1906"
	case DatumE_IndonesianNationalSpheroid:
		return "DatumE_IndonesianNationalSpheroid"
	case DatumE_International1924:
		return "DatumE_International1924"
	case DatumE_International1967:
		return "DatumE_International1967"
	case DatumE_Krassowsky1960:
		return "DatumE_Krassowsky1960"
	case DatumE_NWL9D:
		return "DatumE_NWL9D"
	case DatumE_NWL10D:
		return "DatumE_NWL10D"
	case DatumE_Plessis1817:
		return "DatumE_Plessis1817"
	case DatumE_Struve1860:
		return "DatumE_Struve1860"
	case DatumE_WarOffice:
		return "DatumE_WarOffice"
	case DatumE_WGS84:
		return "DatumE_WGS84"
	case DatumE_GEM10C:
		return "DatumE_GEM10C"
	case DatumE_OSU86F:
		return "DatumE_OSU86F"
	case DatumE_OSU91A:
		return "DatumE_OSU91A"
	case DatumE_Clarke1880:
		return "DatumE_Clarke1880"
	case DatumE_Sphere:
		return "DatumE_Sphere"
	case Datum_Adindan:
		return "Datum_Adindan"
	case Datum_Australian_Geodetic_Datum_1966:
		return "Datum_Australian_Geodetic_Datum_1966"
	case Datum_Australian_Geodetic_Datum_1984:
		return "Datum_Australian_Geodetic_Datum_1984"
	case Datum_Ain_el_Abd_1970:
		return "Datum_Ain_el_Abd_1970"
	case Datum_Afgooye:
		return "Datum_Afgooye"
	case Datum_Agadez:
		return "Datum_Agadez"
	case Datum_Lisbon:
		return "Datum_Lisbon"
	case Datum_Aratu:
		return "Datum_Aratu"
	case Datum_Arc_1950:
		return "Datum_Arc_1950"
	case Datum_Arc_1960:
		return "Datum_Arc_1960"
	case Datum_Batavia:
		return "Datum_Batavia"
	case Datum_Barbados:
		return "Datum_Barbados"
	case Datum_Beduaram:
		return "Datum_Beduaram"
	case Datum_Beijing_1954:
		return "Datum_Beijing_1954"
	case Datum_Reseau_National_Belge_1950:
		return "Datum_Reseau_National_Belge_1950"
	case Datum_Bermuda_1957:
		return "Datum_Bermuda_1957"
	case Datum_Bern_1898:
		return "Datum_Bern_1898"
	case Datum_Bogota:
		return "Datum_Bogota"
	case Datum_Bukit_Rimpah:
		return "Datum_Bukit_Rimpah"
	case Datum_Camacupa:
		return "Datum_Camacupa"
	case Datum_Campo_Inchauspe:
		return "Datum_Campo_Inchauspe"
	case Datum_Cape:
		return "Datum_Cape"
	case Datum_Carthage:
		return "Datum_Carthage"
	case Datum_Chua:
		return "Datum_Chua"
	case Datum_Corrego_Alegre:
		return "Datum_Corrego_Alegre"
	case Datum_Cote_d_Ivoire:
		return "Datum_Cote_d_Ivoire"
	case Datum_Deir_ez_Zor:
		return "Datum_Deir_ez_Zor"
	case Datum_Douala:
		return "Datum_Douala"
	case Datum_Egypt_1907:
		return "Datum_Egypt_1907"
	case Datum_European_Datum_1950:
		return "Datum_European_Datum_1950"
	case Datum_European_Datum_1987:
		return "Datum_European_Datum_1987"
	case Datum_Fahud:
		return "Datum_Fahud"
	case Datum_Gandajika_1970:
		return "Datum_Gandajika_1970"
	case Datum_Garoua:
		return "Datum_Garoua"
	case Datum_Guyane_Francaise:
		return "Datum_Guyane_Francaise"
	case Datum_Hu_Tzu_Shan:
		return "Datum_Hu_Tzu_Shan"
	case Datum_Hungarian_Datum_1972:
		return "Datum_Hungarian_Datum_1972"
	case Datum_Indonesian_Datum_1974:
		return "Datum_Indonesian_Datum_1974"
	case Datum_Indian_1954:
		return "Datum_Indian_1954"
	case Datum_Indian_1975:
		return "Datum_Indian_1975"
	case Datum_Jamaica_1875:
		return "Datum_Jamaica_1875"
	case Datum_Jamaica_1969:
		return "Datum_Jamaica_1969"
	case Datum_Kalianpur:
		return "Datum_Kalianpur"
	case Datum_Kandawala:
		return "Datum_Kandawala"
	case Datum_Kertau:
		return "Datum_Kertau"
	case Datum_Kuwait_Oil_Company:
		return "Datum_Kuwait_Oil_Company"
	case Datum_La_Canoa:
		return "Datum_La_Canoa"
	case Datum_Provisional_S_American_Datum_1956:
		return "Datum_Provisional_S_American_Datum_1956"
	case Datum_Lake:
		return "Datum_Lake"
	case Datum_Leigon:
		return "Datum_Leigon"
	case Datum_Liberia_1964:
		return "Datum_Liberia_1964"
	case Datum_Lome:
		return "Datum_Lome"
	case Datum_Luzon_1911:
		return "Datum_Luzon_1911"
	case Datum_Hito_XVIII_1963:
		return "Datum_Hito_XVIII_1963"
	case Datum_Herat_North:
		return "Datum_Herat_North"
	case Datum_Mahe_1971:
		return "Datum_Mahe_1971"
	case Datum_Makassar:
		return "Datum_Makassar"
	case Datum_European_Reference_System_1989:
		return "Datum_European_Reference_System_1989"
	case Datum_Malongo_1987:
		return "Datum_Malongo_1987"
	case Datum_Manoca:
		return "Datum_Manoca"
	case Datum_Merchich:
		return "Datum_Merchich"
	case Datum_Massawa:
		return "Datum_Massawa"
	case Datum_Minna:
		return "Datum_Minna"
	case Datum_Mhast:
		return "Datum_Mhast"
	case Datum_Monte_Mario:
		return "Datum_Monte_Mario"
	case Datum_M_poraloko:
		return "Datum_M_poraloko"
	case Datum_North_American_Datum_1927:
		return "Datum_North_American_Datum_1927"
	case Datum_NAD_Michigan:
		return "Datum_NAD_Michigan"
	case Datum_North_American_Datum_1983:
		return "Datum_North_American_Datum_1983"
	case Datum_Nahrwan_1967:
		return "Datum_Nahrwan_1967"
	case Datum_Naparima_1972:
		return "Datum_Naparima_1972"
	case Datum_New_Zealand_Geodetic_Datum_1949:
		return "Datum_New_Zealand_Geodetic_Datum_1949"
	case Datum_NGO_1948:
		return "Datum_NGO_1948"
	case Datum_Datum_73:
		return "Datum_Datum_73"
	case Datum_Nouvelle_Triangulation_Francaise:
		return "Datum_Nouvelle_Triangulation_Francaise"
	case Datum_NSWC_9Z_2:
		return "Datum_NSWC_9Z_2"
	case Datum_OSGB_1936:
		return "Datum_OSGB_1936"
	case Datum_OSGB_1970_SN:
		return "Datum_OSGB_1970_SN"
	case Datum_OS_SN_1980:
		return "Datum_OS_SN_1980"
	case Datum_Padang_1884:
		return "Datum_Padang_1884"
	case Datum_Palestine_1923:
		return "Datum_Palestine_1923"
	case Datum_Pointe_Noire:
		return "Datum_Pointe_Noire"
	case Datum_Geocentric_Datum_of_Australia_1994:
		return "Datum_Geocentric_Datum_of_Australia_1994"
	case Datum_Pulkovo_1942:
		return "Datum_Pulkovo_1942"
	case Datum_Qatar:
		return "Datum_Qatar"
	case Datum_Qatar_1948:
		return "Datum_Qatar_1948"
	case Datum_Qornoq:
		return "Datum_Qornoq"
	case Datum_Loma_Quintana:
		return "Datum_Loma_Quintana"
	case Datum_Amersfoort:
		return "Datum_Amersfoort"
	case Datum_RT38:
		return "Datum_RT38"
	case Datum_South_American_Datum_1969:
		return "Datum_South_American_Datum_1969"
	case Datum_Sapper_Hill_1943:
		return "Datum_Sapper_Hill_1943"
	case Datum_Schwarzeck:
		return "Datum_Schwarzeck"
	case Datum_Segora:
		return "Datum_Segora"
	case Datum_Serindung:
		return "Datum_Serindung"
	case Datum_Sudan:
		return "Datum_Sudan"
	case Datum_Tananarive_1925:
		return "Datum_Tananarive_1925"
	case Datum_Timbalai_1948:
		return "Datum_Timbalai_1948"
	case Datum_TM65:
		return "Datum_TM65"
	case Datum_TM75:
		return "Datum_TM75"
	case Datum_Tokyo:
		return "Datum_Tokyo"
	case Datum_Trinidad_1903:
		return "Datum_Trinidad_1903"
	case Datum_Trucial_Coast_1948:
		return "Datum_Trucial_Coast_1948"
	case Datum_Voirol_1875:
		return "Datum_Voirol_1875"
	case Datum_Voirol_Unifie_1960:
		return "Datum_Voirol_Unifie_1960"
	case Datum_Bern_1938:
		return "Datum_Bern_1938"
	case Datum_Nord_Sahara_1959:
		return "Datum_Nord_Sahara_1959"
	case Datum_Stockholm_1938:
		return "Datum_Stockholm_1938"
	case Datum_Yacare:
		return "Datum_Yacare"
	case Datum_Yoff:
		return "Datum_Yoff"
	case Datum_Zanderij:
		return "Datum_Zanderij"
	case Datum_Militar_Geographische_Institut:
		return "Datum_Militar_Geographische_Institut"
	case Datum_Reseau_National_Belge_1972:
		return "Datum_Reseau_National_Belge_1972"
	case Datum_Deutsche_Hauptdreiecksnetz:
		return "Datum_Deutsche_Hauptdreiecksnetz"
	case Datum_Conakry_1905:
		return "Datum_Conakry_1905"
	case Datum_WGS72:
		return "Datum_WGS72"
	case Datum_WGS72_Transit_Broadcast_Ephemeris:
		return "Datum_WGS72_Transit_Broadcast_Ephemeris"
	case Datum_WGS84:
		return "Datum_WGS84"
	case Datum_Ancienne_Triangulation_Francaise:
		return "Datum_Ancienne_Triangulation_Francaise"
	case Datum_Nord_de_Guerre:
		return "Datum_Nord_de_Guerre"
	}
	return fmt.Sprintf("Unknown dataum (%d)", v)
}

const (
	GCSE_Airy1830                   = 4001
	GCSE_AiryModified1849           = 4002
	GCSE_AustralianNationalSpheroid = 4003
	GCSE_Bessel1841                 = 4004
	GCSE_BesselModified             = 4005
	GCSE_BesselNamibia              = 4006
	GCSE_Clarke1858                 = 4007
	GCSE_Clarke1866                 = 4008
	GCSE_Clarke1866Michigan         = 4009
	GCSE_Clarke1880_Benoit          = 4010
	GCSE_Clarke1880_IGN             = 4011
	GCSE_Clarke1880_RGS             = 4012
	GCSE_Clarke1880_Arc             = 4013
	GCSE_Clarke1880_SGA1922         = 4014
	GCSE_Everest1830_1937Adjustment = 4015
	GCSE_Everest1830_1967Definition = 4016
	GCSE_Everest1830_1975Definition = 4017
	GCSE_Everest1830Modified        = 4018
	GCSE_GRS1980                    = 4019
	GCSE_Helmert1906                = 4020
	GCSE_IndonesianNationalSpheroid = 4021
	GCSE_International1924          = 4022
	GCSE_International1967          = 4023
	GCSE_Krassowsky1940             = 4024
	GCSE_NWL9D                      = 4025
	GCSE_NWL10D                     = 4026
	GCSE_Plessis1817                = 4027
	GCSE_Struve1860                 = 4028
	GCSE_WarOffice                  = 4029
	GCSE_WGS84                      = 4030
	GCSE_GEM10C                     = 4031
	GCSE_OSU86F                     = 4032
	GCSE_OSU91A                     = 4033
	GCSE_Clarke1880                 = 4034
	GCSE_Sphere                     = 4035
	GCS_Greek                       = 4120
	GCS_GGRS87                      = 4121
	GCS_KKJ                         = 4123
	GCS_RT90                        = 4124
	GCS_EST92                       = 4133
	GCS_Dealul_Piscului_1970        = 4317
	GCS_Greek_Athens                = 4815
	GCS_Adindan                     = 4201
	GCS_AGD66                       = 4202
	GCS_AGD84                       = 4203
	GCS_Ain_el_Abd                  = 4204
	GCS_Afgooye                     = 4205
	GCS_Agadez                      = 4206
	GCS_Lisbon                      = 4207
	GCS_Aratu                       = 4208
	GCS_Arc_1950                    = 4209
	GCS_Arc_1960                    = 4210
	GCS_Batavia                     = 4211
	GCS_Barbados                    = 4212
	GCS_Beduaram                    = 4213
	GCS_Beijing_1954                = 4214
	GCS_Belge_1950                  = 4215
	GCS_Bermuda_1957                = 4216
	GCS_Bern_1898                   = 4217
	GCS_Bogota                      = 4218
	GCS_Bukit_Rimpah                = 4219
	GCS_Camacupa                    = 4220
	GCS_Campo_Inchauspe             = 4221
	GCS_Cape                        = 4222
	GCS_Carthage                    = 4223
	GCS_Chua                        = 4224
	GCS_Corrego_Alegre              = 4225
	GCS_Cote_d_Ivoire               = 4226
	GCS_Deir_ez_Zor                 = 4227
	GCS_Douala                      = 4228
	GCS_Egypt_1907                  = 4229
	GCS_ED50                        = 4230
	GCS_ED87                        = 4231
	GCS_Fahud                       = 4232
	GCS_Gandajika_1970              = 4233
	GCS_Garoua                      = 4234
	GCS_Guyane_Francaise            = 4235
	GCS_Hu_Tzu_Shan                 = 4236
	GCS_HD72                        = 4237
	GCS_ID74                        = 4238
	GCS_Indian_1954                 = 4239
	GCS_Indian_1975                 = 4240
	GCS_Jamaica_1875                = 4241
	GCS_JAD69                       = 4242
	GCS_Kalianpur                   = 4243
	GCS_Kandawala                   = 4244
	GCS_Kertau                      = 4245
	GCS_KOC                         = 4246
	GCS_La_Canoa                    = 4247
	GCS_PSAD56                      = 4248
	GCS_Lake                        = 4249
	GCS_Leigon                      = 4250
	GCS_Liberia_1964                = 4251
	GCS_Lome                        = 4252
	GCS_Luzon_1911                  = 4253
	GCS_Hito_XVIII_1963             = 4254
	GCS_Herat_North                 = 4255
	GCS_Mahe_1971                   = 4256
	GCS_Makassar                    = 4257
	GCS_EUREF89                     = 4258
	GCS_Malongo_1987                = 4259
	GCS_Manoca                      = 4260
	GCS_Merchich                    = 4261
	GCS_Massawa                     = 4262
	GCS_Minna                       = 4263
	GCS_Mhast                       = 4264
	GCS_Monte_Mario                 = 4265
	GCS_M_poraloko                  = 4266
	GCS_NAD27                       = 4267
	GCS_NAD_Michigan                = 4268
	GCS_NAD83                       = 4269
	GCS_Nahrwan_1967                = 4270
	GCS_Naparima_1972               = 4271
	GCS_GD49                        = 4272
	GCS_NGO_1948                    = 4273
	GCS_Datum_73                    = 4274
	GCS_NTF                         = 4275
	GCS_NSWC_9Z_2                   = 4276
	GCS_OSGB_1936                   = 4277
	GCS_OSGB70                      = 4278
	GCS_OS_SN80                     = 4279
	GCS_Padang                      = 4280
	GCS_Palestine_1923              = 4281
	GCS_Pointe_Noire                = 4282
	GCS_GDA94                       = 4283
	GCS_Pulkovo_1942                = 4284
	GCS_Qatar                       = 4285
	GCS_Qatar_1948                  = 4286
	GCS_Qornoq                      = 4287
	GCS_Loma_Quintana               = 4288
	GCS_Amersfoort                  = 4289
	GCS_RT38                        = 4290
	GCS_SAD69                       = 4291
	GCS_Sapper_Hill_1943            = 4292
	GCS_Schwarzeck                  = 4293
	GCS_Segora                      = 4294
	GCS_Serindung                   = 4295
	GCS_Sudan                       = 4296
	GCS_Tananarive                  = 4297
	GCS_Timbalai_1948               = 4298
	GCS_TM65                        = 4299
	GCS_TM75                        = 4300
	GCS_Tokyo                       = 4301
	GCS_Trinidad_1903               = 4302
	GCS_TC_1948                     = 4303
	GCS_Voirol_1875                 = 4304
	GCS_Voirol_Unifie               = 4305
	GCS_Bern_1938                   = 4306
	GCS_Nord_Sahara_1959            = 4307
	GCS_Stockholm_1938              = 4308
	GCS_Yacare                      = 4309
	GCS_Yoff                        = 4310
	GCS_Zanderij                    = 4311
	GCS_MGI                         = 4312
	GCS_Belge_1972                  = 4313
	GCS_DHDN                        = 4314
	GCS_Conakry_1905                = 4315
	GCS_WGS_72                      = 4322
	GCS_WGS_72BE                    = 4324
	GCS_WGS_84                      = 4326
	GCS_Bern_1898_Bern              = 4801
	GCS_Bogota_Bogota               = 4802
	GCS_Lisbon_Lisbon               = 4803
	GCS_Makassar_Jakarta            = 4804
	GCS_MGI_Ferro                   = 4805
	GCS_Monte_Mario_Rome            = 4806
	GCS_NTF_Paris                   = 4807
	GCS_Padang_Jakarta              = 4808
	GCS_Belge_1950_Brussels         = 4809
	GCS_Tananarive_Paris            = 4810
	GCS_Voirol_1875_Paris           = 4811
	GCS_Voirol_Unifie_Paris         = 4812
	GCS_Batavia_Jakarta             = 4813
	GCS_ATF_Paris                   = 4901
	GCS_NDG_Paris                   = 4902
)

func GcsName(v int) string {
	switch v {
	case GCSE_Airy1830:
		return "GCSE_Airy1830"
	case GCSE_AiryModified1849:
		return "GCSE_AiryModified1849"
	case GCSE_AustralianNationalSpheroid:
		return "GCSE_AustralianNationalSpheroid"
	case GCSE_Bessel1841:
		return "GCSE_Bessel1841"
	case GCSE_BesselModified:
		return "GCSE_BesselModified"
	case GCSE_BesselNamibia:
		return "GCSE_BesselNamibia"
	case GCSE_Clarke1858:
		return "GCSE_Clarke1858"
	case GCSE_Clarke1866:
		return "GCSE_Clarke1866"
	case GCSE_Clarke1866Michigan:
		return "GCSE_Clarke1866Michigan"
	case GCSE_Clarke1880_Benoit:
		return "GCSE_Clarke1880_Benoit"
	case GCSE_Clarke1880_IGN:
		return "GCSE_Clarke1880_IGN"
	case GCSE_Clarke1880_RGS:
		return "GCSE_Clarke1880_RGS"
	case GCSE_Clarke1880_Arc:
		return "GCSE_Clarke1880_Arc"
	case GCSE_Clarke1880_SGA1922:
		return "GCSE_Clarke1880_SGA1922"
	case GCSE_Everest1830_1937Adjustment:
		return "GCSE_Everest1830_1937Adjustment"
	case GCSE_Everest1830_1967Definition:
		return "GCSE_Everest1830_1967Definition"
	case GCSE_Everest1830_1975Definition:
		return "GCSE_Everest1830_1975Definition"
	case GCSE_Everest1830Modified:
		return "GCSE_Everest1830Modified"
	case GCSE_GRS1980:
		return "GCSE_GRS1980"
	case GCSE_Helmert1906:
		return "GCSE_Helmert1906"
	case GCSE_IndonesianNationalSpheroid:
		return "GCSE_IndonesianNationalSpheroid"
	case GCSE_International1924:
		return "GCSE_International1924"
	case GCSE_International1967:
		return "GCSE_International1967"
	case GCSE_Krassowsky1940:
		return "GCSE_Krassowsky1940"
	case GCSE_NWL9D:
		return "GCSE_NWL9D"
	case GCSE_NWL10D:
		return "GCSE_NWL10D"
	case GCSE_Plessis1817:
		return "GCSE_Plessis1817"
	case GCSE_Struve1860:
		return "GCSE_Struve1860"
	case GCSE_WarOffice:
		return "GCSE_WarOffice"
	case GCSE_WGS84:
		return "GCSE_WGS84"
	case GCSE_GEM10C:
		return "GCSE_GEM10C"
	case GCSE_OSU86F:
		return "GCSE_OSU86F"
	case GCSE_OSU91A:
		return "GCSE_OSU91A"
	case GCSE_Clarke1880:
		return "GCSE_Clarke1880"
	case GCSE_Sphere:
		return "GCSE_Sphere"
	case GCS_Greek:
		return "GCS_Greek"
	case GCS_GGRS87:
		return "GCS_GGRS87"
	case GCS_KKJ:
		return "GCS_KKJ"
	case GCS_RT90:
		return "GCS_RT90"
	case GCS_EST92:
		return "GCS_EST92"
	case GCS_Dealul_Piscului_1970:
		return "GCS_Dealul_Piscului_1970"
	case GCS_Greek_Athens:
		return "GCS_Greek_Athens"
	case GCS_Adindan:
		return "GCS_Adindan"
	case GCS_AGD66:
		return "GCS_AGD66"
	case GCS_AGD84:
		return "GCS_AGD84"
	case GCS_Ain_el_Abd:
		return "GCS_Ain_el_Abd"
	case GCS_Afgooye:
		return "GCS_Afgooye"
	case GCS_Agadez:
		return "GCS_Agadez"
	case GCS_Lisbon:
		return "GCS_Lisbon"
	case GCS_Aratu:
		return "GCS_Aratu"
	case GCS_Arc_1950:
		return "GCS_Arc_1950"
	case GCS_Arc_1960:
		return "GCS_Arc_1960"
	case GCS_Batavia:
		return "GCS_Batavia"
	case GCS_Barbados:
		return "GCS_Barbados"
	case GCS_Beduaram:
		return "GCS_Beduaram"
	case GCS_Beijing_1954:
		return "GCS_Beijing_1954"
	case GCS_Belge_1950:
		return "GCS_Belge_1950"
	case GCS_Bermuda_1957:
		return "GCS_Bermuda_1957"
	case GCS_Bern_1898:
		return "GCS_Bern_1898"
	case GCS_Bogota:
		return "GCS_Bogota"
	case GCS_Bukit_Rimpah:
		return "GCS_Bukit_Rimpah"
	case GCS_Camacupa:
		return "GCS_Camacupa"
	case GCS_Campo_Inchauspe:
		return "GCS_Campo_Inchauspe"
	case GCS_Cape:
		return "GCS_Cape"
	case GCS_Carthage:
		return "GCS_Carthage"
	case GCS_Chua:
		return "GCS_Chua"
	case GCS_Corrego_Alegre:
		return "GCS_Corrego_Alegre"
	case GCS_Cote_d_Ivoire:
		return "GCS_Cote_d_Ivoire"
	case GCS_Deir_ez_Zor:
		return "GCS_Deir_ez_Zor"
	case GCS_Douala:
		return "GCS_Douala"
	case GCS_Egypt_1907:
		return "GCS_Egypt_1907"
	case GCS_ED50:
		return "GCS_ED50"
	case GCS_ED87:
		return "GCS_ED87"
	case GCS_Fahud:
		return "GCS_Fahud"
	case GCS_Gandajika_1970:
		return "GCS_Gandajika_1970"
	case GCS_Garoua:
		return "GCS_Garoua"
	case GCS_Guyane_Francaise:
		return "GCS_Guyane_Francaise"
	case GCS_Hu_Tzu_Shan:
		return "GCS_Hu_Tzu_Shan"
	case GCS_HD72:
		return "GCS_HD72"
	case GCS_ID74:
		return "GCS_ID74"
	case GCS_Indian_1954:
		return "GCS_Indian_1954"
	case GCS_Indian_1975:
		return "GCS_Indian_1975"
	case GCS_Jamaica_1875:
		return "GCS_Jamaica_1875"
	case GCS_JAD69:
		return "GCS_JAD69"
	case GCS_Kalianpur:
		return "GCS_Kalianpur"
	case GCS_Kandawala:
		return "GCS_Kandawala"
	case GCS_Kertau:
		return "GCS_Kertau"
	case GCS_KOC:
		return "GCS_KOC"
	case GCS_La_Canoa:
		return "GCS_La_Canoa"
	case GCS_PSAD56:
		return "GCS_PSAD56"
	case GCS_Lake:
		return "GCS_Lake"
	case GCS_Leigon:
		return "GCS_Leigon"
	case GCS_Liberia_1964:
		return "GCS_Liberia_1964"
	case GCS_Lome:
		return "GCS_Lome"
	case GCS_Luzon_1911:
		return "GCS_Luzon_1911"
	case GCS_Hito_XVIII_1963:
		return "GCS_Hito_XVIII_1963"
	case GCS_Herat_North:
		return "GCS_Herat_North"
	case GCS_Mahe_1971:
		return "GCS_Mahe_1971"
	case GCS_Makassar:
		return "GCS_Makassar"
	case GCS_EUREF89:
		return "GCS_EUREF89"
	case GCS_Malongo_1987:
		return "GCS_Malongo_1987"
	case GCS_Manoca:
		return "GCS_Manoca"
	case GCS_Merchich:
		return "GCS_Merchich"
	case GCS_Massawa:
		return "GCS_Massawa"
	case GCS_Minna:
		return "GCS_Minna"
	case GCS_Mhast:
		return "GCS_Mhast"
	case GCS_Monte_Mario:
		return "GCS_Monte_Mario"
	case GCS_M_poraloko:
		return "GCS_M_poraloko"
	case GCS_NAD27:
		return "GCS_NAD27"
	case GCS_NAD_Michigan:
		return "GCS_NAD_Michigan"
	case GCS_NAD83:
		return "GCS_NAD83"
	case GCS_Nahrwan_1967:
		return "GCS_Nahrwan_1967"
	case GCS_Naparima_1972:
		return "GCS_Naparima_1972"
	case GCS_GD49:
		return "GCS_GD49"
	case GCS_NGO_1948:
		return "GCS_NGO_1948"
	case GCS_Datum_73:
		return "GCS_Datum_73"
	case GCS_NTF:
		return "GCS_NTF"
	case GCS_NSWC_9Z_2:
		return "GCS_NSWC_9Z_2"
	case GCS_OSGB_1936:
		return "GCS_OSGB_1936"
	case GCS_OSGB70:
		return "GCS_OSGB70"
	case GCS_OS_SN80:
		return "GCS_OS_SN80"
	case GCS_Padang:
		return "GCS_Padang"
	case GCS_Palestine_1923:
		return "GCS_Palestine_1923"
	case GCS_Pointe_Noire:
		return "GCS_Pointe_Noire"
	case GCS_GDA94:
		return "GCS_GDA94"
	case GCS_Pulkovo_1942:
		return "GCS_Pulkovo_1942"
	case GCS_Qatar:
		return "GCS_Qatar"
	case GCS_Qatar_1948:
		return "GCS_Qatar_1948"
	case GCS_Qornoq:
		return "GCS_Qornoq"
	case GCS_Loma_Quintana:
		return "GCS_Loma_Quintana"
	case GCS_Amersfoort:
		return "GCS_Amersfoort"
	case GCS_RT38:
		return "GCS_RT38"
	case GCS_SAD69:
		return "GCS_SAD69"
	case GCS_Sapper_Hill_1943:
		return "GCS_Sapper_Hill_1943"
	case GCS_Schwarzeck:
		return "GCS_Schwarzeck"
	case GCS_Segora:
		return "GCS_Segora"
	case GCS_Serindung:
		return "GCS_Serindung"
	case GCS_Sudan:
		return "GCS_Sudan"
	case GCS_Tananarive:
		return "GCS_Tananarive"
	case GCS_Timbalai_1948:
		return "GCS_Timbalai_1948"
	case GCS_TM65:
		return "GCS_TM65"
	case GCS_TM75:
		return "GCS_TM75"
	case GCS_Tokyo:
		return "GCS_Tokyo"
	case GCS_Trinidad_1903:
		return "GCS_Trinidad_1903"
	case GCS_TC_1948:
		return "GCS_TC_1948"
	case GCS_Voirol_1875:
		return "GCS_Voirol_1875"
	case GCS_Voirol_Unifie:
		return "GCS_Voirol_Unifie"
	case GCS_Bern_1938:
		return "GCS_Bern_1938"
	case GCS_Nord_Sahara_1959:
		return "GCS_Nord_Sahara_1959"
	case GCS_Stockholm_1938:
		return "GCS_Stockholm_1938"
	case GCS_Yacare:
		return "GCS_Yacare"
	case GCS_Yoff:
		return "GCS_Yoff"
	case GCS_Zanderij:
		return "GCS_Zanderij"
	case GCS_MGI:
		return "GCS_MGI"
	case GCS_Belge_1972:
		return "GCS_Belge_1972"
	case GCS_DHDN:
		return "GCS_DHDN"
	case GCS_Conakry_1905:
		return "GCS_Conakry_1905"
	case GCS_WGS_72:
		return "GCS_WGS_72"
	case GCS_WGS_72BE:
		return "GCS_WGS_72BE"
	case GCS_WGS_84:
		return "GCS_WGS_84"
	case GCS_Bern_1898_Bern:
		return "GCS_Bern_1898_Bern"
	case GCS_Bogota_Bogota:
		return "GCS_Bogota_Bogota"
	case GCS_Lisbon_Lisbon:
		return "GCS_Lisbon_Lisbon"
	case GCS_Makassar_Jakarta:
		return "GCS_Makassar_Jakarta"
	case GCS_MGI_Ferro:
		return "GCS_MGI_Ferro"
	case GCS_Monte_Mario_Rome:
		return "GCS_Monte_Mario_Rome"
	case GCS_NTF_Paris:
		return "GCS_NTF_Paris"
	case GCS_Padang_Jakarta:
		return "GCS_Padang_Jakarta"
	case GCS_Belge_1950_Brussels:
		return "GCS_Belge_1950_Brussels"
	case GCS_Tananarive_Paris:
		return "GCS_Tananarive_Paris"
	case GCS_Voirol_1875_Paris:
		return "GCS_Voirol_1875_Paris"
	case GCS_Voirol_Unifie_Paris:
		return "GCS_Voirol_Unifie_Paris"
	case GCS_Batavia_Jakarta:
		return "GCS_Batavia_Jakarta"
	case GCS_ATF_Paris:
		return "GCS_ATF_Paris"
	case GCS_NDG_Paris:
		return "GCS_NDG_Paris"
	}
	return fmt.Sprintf("Unkown GCS (%d)", v)
}

const (
	PM_Greenwich = 8901
	PM_Lisbon    = 8902
	PM_Paris     = 8903
	PM_Bogota    = 8904
	PM_Madrid    = 8905
	PM_Rome      = 8906
	PM_Bern      = 8907
	PM_Jakarta   = 8908
	PM_Ferro     = 8909
	PM_Brussels  = 8910
	PM_Stockholm = 8911
)

func PrimeMeridianName(v int) string {
	switch v {
	case PM_Greenwich:
		return "PM_Greenwich"
	case PM_Lisbon:
		return "PM_Lisbon"
	case PM_Paris:
		return "PM_Paris"
	case PM_Bogota:
		return "PM_Bogota"
	case PM_Madrid:
		return "PM_Madrid"
	case PM_Rome:
		return "PM_Rome"
	case PM_Bern:
		return "PM_Bern"
	case PM_Jakarta:
		return "PM_Jakarta"
	case PM_Ferro:
		return "PM_Ferro"
	case PM_Brussels:
		return "PM_Brussels"
	case PM_Stockholm:
		return "PM_Stockholm"
	}
	return fmt.Sprintf("Unkown PrimeMeridion (%d)", v)
}

// https://github.com/smanders/libgeotiff/blob/master/epsg_ellipse.inc
const (
	Ellipse_Airy_1830 = 7001
	Ellipse_Airy_Modified_1849 = 7002
	Ellipse_Australian_National_Spheroid = 7003
	Ellipse_Bessel_1841 = 7004
	Ellipse_Bessel_Modified = 7005
	Ellipse_Bessel_Namibia = 7006
	Ellipse_Clarke_1858 = 7007
	Ellipse_Clarke_1866 = 7008
	Ellipse_Clarke_1866_Michigan = 7009
	Ellipse_Clarke_1880_Benoit = 7010
	Ellipse_Clarke_1880_IGN = 7011
	Ellipse_Clarke_1880_RGS = 7012
	Ellipse_Clarke_1880_Arc = 7013
	Ellipse_Clarke_1880_SGA_1922 = 7014
	Ellipse_Everest_1830_1937_Adjustment = 7015
	Ellipse_Everest_1830_1967_Definition = 7016
	Ellipse_Everest_1830_1975_Definition = 7017
	Ellipse_Everest_1830_Modified = 7018
	Ellipse_GRS_1980 = 7019
	Ellipse_Helmert_1906 = 7020
	Ellipse_Indonesian_National_Spheroid = 7021
	Ellipse_International_1924 = 7022
	Ellipse_International_1967 = 7023
	Ellipse_Krassowsky_1940 = 7024
	Ellipse_NWL_9D = 7025
	Ellipse_NWL_10D = 7026
	Ellipse_Plessis_1817 = 7027
	Ellipse_Struve_1860 = 7028
	Ellipse_War_Office = 7029
	Ellipse_WGS_84 = 7030
	Ellipse_GEM_10C = 7031
	Ellipse_OSU86F = 7032
	Ellipse_OSU91A = 7033
	Ellipse_Clarke_1880 = 7034
	Ellipse_Sphere = 7035
)

func EllipseName(v int) string {
     switch v {
	case Ellipse_Airy_1830:
		return "Ellipse_Airy_1830"
	case Ellipse_Airy_Modified_1849:
		return "Ellipse_Airy_Modified_1849"
	case Ellipse_Australian_National_Spheroid:
		return "Ellipse_Australian_National_Spheroid"
	case Ellipse_Bessel_1841:
		return "Ellipse_Bessel_1841"
	case Ellipse_Bessel_Modified:
		return "Ellipse_Bessel_Modified"
	case Ellipse_Bessel_Namibia:
		return "Ellipse_Bessel_Namibia"
	case Ellipse_Clarke_1858:
		return "Ellipse_Clarke_1858"
	case Ellipse_Clarke_1866:
		return "Ellipse_Clarke_1866"
	case Ellipse_Clarke_1866_Michigan:
		return "Ellipse_Clarke_1866_Michigan"
	case Ellipse_Clarke_1880_Benoit:
		return "Ellipse_Clarke_1880_Benoit"
	case Ellipse_Clarke_1880_IGN:
		return "Ellipse_Clarke_1880_IGN"
	case Ellipse_Clarke_1880_RGS:
		return "Ellipse_Clarke_1880_RGS"
	case Ellipse_Clarke_1880_Arc:
		return "Ellipse_Clarke_1880_Arc"
	case Ellipse_Clarke_1880_SGA_1922:
		return "Ellipse_Clarke_1880_SGA_1922"
	case Ellipse_Everest_1830_1937_Adjustment:
		return "Ellipse_Everest_1830_1937_Adjustment"
	case Ellipse_Everest_1830_1967_Definition:
		return "Ellipse_Everest_1830_1967_Definition"
	case Ellipse_Everest_1830_1975_Definition:
		return "Ellipse_Everest_1830_1975_Definition"
	case Ellipse_Everest_1830_Modified:
		return "Ellipse_Everest_1830_Modified"
	case Ellipse_GRS_1980:
		return "Ellipse_GRS_1980"
	case Ellipse_Helmert_1906:
		return "Ellipse_Helmert_1906"
	case Ellipse_Indonesian_National_Spheroid:
		return "Ellipse_Indonesian_National_Spheroid"
	case Ellipse_International_1924:
		return "Ellipse_International_1924"
	case Ellipse_International_1967:
		return "Ellipse_International_1967"
	case Ellipse_Krassowsky_1940:
		return "Ellipse_Krassowsky_1940"
	case Ellipse_NWL_9D:
		return "Ellipse_NWL_9D"
	case Ellipse_NWL_10D:
		return "Ellipse_NWL_10D"
	case Ellipse_Plessis_1817:
		return "Ellipse_Plessis_1817"
	case Ellipse_Struve_1860:
		return "Ellipse_Struve_1860"
	case Ellipse_War_Office:
		return "Ellipse_War_Office"
	case Ellipse_WGS_84:
		return "Ellipse_WGS_84"
	case Ellipse_GEM_10C:
		return "Ellipse_GEM_10C"
	case Ellipse_OSU86F:
		return "Ellipse_OSU86F"
	case Ellipse_OSU91A:
		return "Ellipse_OSU91A"
	case Ellipse_Clarke_1880:
		return "Ellipse_Clarke_1880"
	case Ellipse_Sphere:
		return "Ellipse_Sphere"
  }
  return fmt.Sprintf("Unkown ellipse (%d)", v)
}
