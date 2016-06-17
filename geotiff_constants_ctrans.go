package main

import "fmt"

// https://github.com/smanders/libgeotiff/blob/master/geo_ctrans.inc
const (
	CT_TransverseMercator             = 1
	CT_TransvMercator_Modified_Alaska = 2
	CT_ObliqueMercator                = 3
	CT_ObliqueMercator_Laborde        = 4
	CT_ObliqueMercator_Rosenmund      = 5
	CT_Mercator                       = 7
	CT_LambertConfConic_2SP           = 8
	CT_LambertConfConic_1SP           = 9
	CT_LambertAzimEqualArea           = 10
	CT_AlbersEqualArea                = 11
	CT_AzimuthalEquidistant           = 12
	CT_EquidistantConic               = 13
	CT_Stereographic                  = 14
	CT_PolarStereographic             = 15
	CT_Equirectangular                = 17
	CT_CassiniSoldner                 = 18
	CT_Gnomonic                       = 19
	CT_MillerCylindrical              = 20
	CT_Orthographic                   = 21
	CT_Polyconic                      = 22
	CT_Robinson                       = 23
	CT_Sinusoidal                     = 24
	CT_VanDerGrinten                  = 25
	CT_NewZealandMapGrid              = 26
	CT_TransvMercator_SouthOrientated = 27
	CT_CylindricalEqualArea           = 28
)

func CoordTransName(v int) string {
	switch v {
	case CT_TransverseMercator:
		return "CT_TransverseMercator"
	case CT_TransvMercator_Modified_Alaska:
		return "CT_TransvMercator_Modified_Alaska"
	case CT_ObliqueMercator:
		return "CT_ObliqueMercator"
	case CT_ObliqueMercator_Laborde:
		return "CT_ObliqueMercator_Laborde"
	case CT_ObliqueMercator_Rosenmund:
		return "CT_ObliqueMercator_Rosenmund"
	case CT_Mercator:
		return "CT_Mercator"
	case CT_LambertConfConic_2SP:
		return "CT_LambertConfConic_2SP"
	case CT_LambertConfConic_1SP:
		return "CT_LambertConfConic_1SP"
	case CT_LambertAzimEqualArea:
		return "CT_LambertAzimEqualArea"
	case CT_AlbersEqualArea:
		return "CT_AlbersEqualArea"
	case CT_AzimuthalEquidistant:
		return "CT_AzimuthalEquidistant"
	case CT_EquidistantConic:
		return "CT_EquidistantConic"
	case CT_Stereographic:
		return "CT_Stereographic"
	case CT_PolarStereographic:
		return "CT_PolarStereographic"
	case CT_Equirectangular:
		return "CT_Equirectangular"
	case CT_CassiniSoldner:
		return "CT_CassiniSoldner"
	case CT_Gnomonic:
		return "CT_Gnomonic"
	case CT_MillerCylindrical:
		return "CT_MillerCylindrical"
	case CT_Orthographic:
		return "CT_Orthographic"
	case CT_Polyconic:
		return "CT_Polyconic"
	case CT_Robinson:
		return "CT_Robinson"
	case CT_Sinusoidal:
		return "CT_Sinusoidal"
	case CT_VanDerGrinten:
		return "CT_VanDerGrinten"
	case CT_NewZealandMapGrid:
		return "CT_NewZealandMapGrid"
	case CT_TransvMercator_SouthOrientated:
		return "CT_TransvMercator_SouthOrientated"
	case CT_CylindricalEqualArea:
		return "CT_CylindricalEqualArea"
	}
	return fmt.Sprintf("Unkown CoordTrans (%d)", v)
}
