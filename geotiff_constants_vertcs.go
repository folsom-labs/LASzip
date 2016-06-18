package main

import "fmt"

// https://github.com/smanders/libgeotiff/blob/master/epsg_vertcs.inc
const (
	VertCS_Airy_1830_ellipsoid = 5001
	VertCS_Airy_Modified_1849_ellipsoid = 5002
	VertCS_ANS_ellipsoid = 5003
	VertCS_Bessel_1841_ellipsoid = 5004
	VertCS_Bessel_Modified_ellipsoid = 5005
	VertCS_Bessel_Namibia_ellipsoid = 5006
	VertCS_Clarke_1858_ellipsoid = 5007
	VertCS_Clarke_1866_ellipsoid = 5008
	VertCS_Clarke_1880_Benoit_ellipsoid = 5010
	VertCS_Clarke_1880_IGN_ellipsoid = 5011
	VertCS_Clarke_1880_RGS_ellipsoid = 5012
	VertCS_Clarke_1880_Arc_ellipsoid = 5013
	VertCS_Clarke_1880_SGA_1922_ellipsoid = 5014
	VertCS_Everest_1830_1937_Adjustment_ellipsoid = 5015
	VertCS_Everest_1830_1967_Definition_ellipsoid = 5016
	VertCS_Everest_1830_1975_Definition_ellipsoid = 5017
	VertCS_Everest_1830_Modified_ellipsoid = 5018
	VertCS_GRS_1980_ellipsoid = 5019
	VertCS_Helmert_1906_ellipsoid = 5020
	VertCS_INS_ellipsoid = 5021
	VertCS_International_1924_ellipsoid = 5022
	VertCS_International_1967_ellipsoid = 5023
	VertCS_Krassowsky_1940_ellipsoid = 5024
	VertCS_NWL_9D_ellipsoid = 5025
	VertCS_NWL_10D_ellipsoid = 5026
	VertCS_Plessis_1817_ellipsoid = 5027
	VertCS_Struve_1860_ellipsoid = 5028
	VertCS_War_Office_ellipsoid = 5029
	VertCS_WGS_84_ellipsoid = 5030
	VertCS_GEM_10C_ellipsoid = 5031
	VertCS_OSU86F_ellipsoid = 5032
	VertCS_OSU91A_ellipsoid = 5033
	VertCS_Newlyn = 5101
	VertCS_North_American_Vertical_Datum_1929 = 5102
	VertCS_North_American_Vertical_Datum_1988 = 5103
	VertCS_Yellow_Sea_1956 = 5104
	VertCS_Baltic_Sea = 5105
	VertCS_Caspian_Sea = 5106
)


func VcsName(v int) string {
  switch v {
	case VertCS_Airy_1830_ellipsoid:
		return "VertCS_Airy_1830_ellipsoid"
	case VertCS_Airy_Modified_1849_ellipsoid:
		return "VertCS_Airy_Modified_1849_ellipsoid"
	case VertCS_ANS_ellipsoid:
		return "VertCS_ANS_ellipsoid"
	case VertCS_Bessel_1841_ellipsoid:
		return "VertCS_Bessel_1841_ellipsoid"
	case VertCS_Bessel_Modified_ellipsoid:
		return "VertCS_Bessel_Modified_ellipsoid"
	case VertCS_Bessel_Namibia_ellipsoid:
		return "VertCS_Bessel_Namibia_ellipsoid"
	case VertCS_Clarke_1858_ellipsoid:
		return "VertCS_Clarke_1858_ellipsoid"
	case VertCS_Clarke_1866_ellipsoid:
		return "VertCS_Clarke_1866_ellipsoid"
	case VertCS_Clarke_1880_Benoit_ellipsoid:
		return "VertCS_Clarke_1880_Benoit_ellipsoid"
	case VertCS_Clarke_1880_IGN_ellipsoid:
		return "VertCS_Clarke_1880_IGN_ellipsoid"
	case VertCS_Clarke_1880_RGS_ellipsoid:
		return "VertCS_Clarke_1880_RGS_ellipsoid"
	case VertCS_Clarke_1880_Arc_ellipsoid:
		return "VertCS_Clarke_1880_Arc_ellipsoid"
	case VertCS_Clarke_1880_SGA_1922_ellipsoid:
		return "VertCS_Clarke_1880_SGA_1922_ellipsoid"
	case VertCS_Everest_1830_1937_Adjustment_ellipsoid:
		return "VertCS_Everest_1830_1937_Adjustment_ellipsoid"
	case VertCS_Everest_1830_1967_Definition_ellipsoid:
		return "VertCS_Everest_1830_1967_Definition_ellipsoid"
	case VertCS_Everest_1830_1975_Definition_ellipsoid:
		return "VertCS_Everest_1830_1975_Definition_ellipsoid"
	case VertCS_Everest_1830_Modified_ellipsoid:
		return "VertCS_Everest_1830_Modified_ellipsoid"
	case VertCS_GRS_1980_ellipsoid:
		return "VertCS_GRS_1980_ellipsoid"
	case VertCS_Helmert_1906_ellipsoid:
		return "VertCS_Helmert_1906_ellipsoid"
	case VertCS_INS_ellipsoid:
		return "VertCS_INS_ellipsoid"
	case VertCS_International_1924_ellipsoid:
		return "VertCS_International_1924_ellipsoid"
	case VertCS_International_1967_ellipsoid:
		return "VertCS_International_1967_ellipsoid"
	case VertCS_Krassowsky_1940_ellipsoid:
		return "VertCS_Krassowsky_1940_ellipsoid"
	case VertCS_NWL_9D_ellipsoid:
		return "VertCS_NWL_9D_ellipsoid"
	case VertCS_NWL_10D_ellipsoid:
		return "VertCS_NWL_10D_ellipsoid"
	case VertCS_Plessis_1817_ellipsoid:
		return "VertCS_Plessis_1817_ellipsoid"
	case VertCS_Struve_1860_ellipsoid:
		return "VertCS_Struve_1860_ellipsoid"
	case VertCS_War_Office_ellipsoid:
		return "VertCS_War_Office_ellipsoid"
	case VertCS_WGS_84_ellipsoid:
		return "VertCS_WGS_84_ellipsoid"
	case VertCS_GEM_10C_ellipsoid:
		return "VertCS_GEM_10C_ellipsoid"
	case VertCS_OSU86F_ellipsoid:
		return "VertCS_OSU86F_ellipsoid"
	case VertCS_OSU91A_ellipsoid:
		return "VertCS_OSU91A_ellipsoid"
	case VertCS_Newlyn:
		return "VertCS_Newlyn"
	case VertCS_North_American_Vertical_Datum_1929:
		return "VertCS_North_American_Vertical_Datum_1929"
	case VertCS_North_American_Vertical_Datum_1988:
		return "VertCS_North_American_Vertical_Datum_1988"
	case VertCS_Yellow_Sea_1956:
		return "VertCS_Yellow_Sea_1956"
	case VertCS_Baltic_Sea:
		return "VertCS_Baltic_Sea"
	case VertCS_Caspian_Sea:
		return "VertCS_Caspian_Sea"
  }
  return fmt.Sprintf("Unkown-%d", v)
}
