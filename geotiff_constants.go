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
