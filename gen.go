package main

import (
	"fmt"
	"strings"
)

// https://raw.githubusercontent.com/smanders/libgeotiff/master/epsg_gcs.inc
var gcs_codes = `
/*
 *  EPSG/POSC GCS Codes -- GeoTIFF Rev. 0.2
 */

/* C database for Geotiff include files.   */
/* the macro ValuePair() must be defined   */
/* by the enclosing include file           */

#ifdef INCLUDE_OLD_CODES
#include old_gcs.inc
#endif /* OLD Codes */

/* Unspecified GCS based on ellipsoid */
ValuePair(GCSE_Airy1830,	4001)
ValuePair(GCSE_AiryModified1849,	4002)
ValuePair(GCSE_AustralianNationalSpheroid,	4003)
ValuePair(GCSE_Bessel1841,	4004)
ValuePair(GCSE_BesselModified,	4005)
ValuePair(GCSE_BesselNamibia,	4006)
ValuePair(GCSE_Clarke1858,	4007)
ValuePair(GCSE_Clarke1866,	4008)
ValuePair(GCSE_Clarke1866Michigan,	4009)
ValuePair(GCSE_Clarke1880_Benoit,	4010)
ValuePair(GCSE_Clarke1880_IGN,	4011)
ValuePair(GCSE_Clarke1880_RGS,	4012)
ValuePair(GCSE_Clarke1880_Arc,	4013)
ValuePair(GCSE_Clarke1880_SGA1922,	4014)
ValuePair(GCSE_Everest1830_1937Adjustment,	4015)
ValuePair(GCSE_Everest1830_1967Definition,	4016)
ValuePair(GCSE_Everest1830_1975Definition,	4017)
ValuePair(GCSE_Everest1830Modified,	4018)
ValuePair(GCSE_GRS1980,	4019)
ValuePair(GCSE_Helmert1906,	4020)
ValuePair(GCSE_IndonesianNationalSpheroid,	4021)
ValuePair(GCSE_International1924,	4022)
ValuePair(GCSE_International1967,	4023)
ValuePair(GCSE_Krassowsky1940,	4024)
ValuePair(GCSE_NWL9D,	4025)
ValuePair(GCSE_NWL10D,	4026)
ValuePair(GCSE_Plessis1817,	4027)
ValuePair(GCSE_Struve1860,	4028)
ValuePair(GCSE_WarOffice,	4029)
ValuePair(GCSE_WGS84,	4030)
ValuePair(GCSE_GEM10C,	4031)
ValuePair(GCSE_OSU86F,	4032)
ValuePair(GCSE_OSU91A,	4033)
ValuePair(GCSE_Clarke1880,	4034)
ValuePair(GCSE_Sphere,	4035)

/* New GCS */
ValuePair(GCS_Greek,4120)
ValuePair(GCS_GGRS87,4121)
ValuePair(GCS_KKJ,4123)
ValuePair(GCS_RT90,4124)
ValuePair(GCS_EST92,4133)
ValuePair(GCS_Dealul_Piscului_1970,4317)
ValuePair(GCS_Greek_Athens,4815)

/* Standard GCS */
ValuePair(GCS_Adindan,	4201)
ValuePair(GCS_AGD66,	4202)
ValuePair(GCS_AGD84,	4203)
ValuePair(GCS_Ain_el_Abd,	4204)
ValuePair(GCS_Afgooye,	4205)
ValuePair(GCS_Agadez,	4206)
ValuePair(GCS_Lisbon,	4207)
ValuePair(GCS_Aratu,	4208)
ValuePair(GCS_Arc_1950,	4209)
ValuePair(GCS_Arc_1960,	4210)
ValuePair(GCS_Batavia,	4211)
ValuePair(GCS_Barbados,	4212)
ValuePair(GCS_Beduaram,	4213)
ValuePair(GCS_Beijing_1954,	4214)
ValuePair(GCS_Belge_1950,	4215)
ValuePair(GCS_Bermuda_1957,	4216)
ValuePair(GCS_Bern_1898,	4217)
ValuePair(GCS_Bogota,	4218)
ValuePair(GCS_Bukit_Rimpah,	4219)
ValuePair(GCS_Camacupa,	4220)
ValuePair(GCS_Campo_Inchauspe,	4221)
ValuePair(GCS_Cape,	4222)
ValuePair(GCS_Carthage,	4223)
ValuePair(GCS_Chua,	4224)
ValuePair(GCS_Corrego_Alegre,	4225)
ValuePair(GCS_Cote_d_Ivoire,	4226)
ValuePair(GCS_Deir_ez_Zor,	4227)
ValuePair(GCS_Douala,	4228)
ValuePair(GCS_Egypt_1907,	4229)
ValuePair(GCS_ED50,	4230)
ValuePair(GCS_ED87,	4231)
ValuePair(GCS_Fahud,	4232)
ValuePair(GCS_Gandajika_1970,	4233)
ValuePair(GCS_Garoua,	4234)
ValuePair(GCS_Guyane_Francaise,	4235)
ValuePair(GCS_Hu_Tzu_Shan,	4236)
ValuePair(GCS_HD72,	4237)
ValuePair(GCS_ID74,	4238)
ValuePair(GCS_Indian_1954,	4239)
ValuePair(GCS_Indian_1975,	4240)
ValuePair(GCS_Jamaica_1875,	4241)
ValuePair(GCS_JAD69,	4242)
ValuePair(GCS_Kalianpur,	4243)
ValuePair(GCS_Kandawala,	4244)
ValuePair(GCS_Kertau,	4245)
ValuePair(GCS_KOC,	4246)
ValuePair(GCS_La_Canoa,	4247)
ValuePair(GCS_PSAD56,	4248)
ValuePair(GCS_Lake,	4249)
ValuePair(GCS_Leigon,	4250)
ValuePair(GCS_Liberia_1964,	4251)
ValuePair(GCS_Lome,	4252)
ValuePair(GCS_Luzon_1911,	4253)
ValuePair(GCS_Hito_XVIII_1963,	4254)
ValuePair(GCS_Herat_North,	4255)
ValuePair(GCS_Mahe_1971,	4256)
ValuePair(GCS_Makassar,	4257)
ValuePair(GCS_EUREF89,	4258)
ValuePair(GCS_Malongo_1987,	4259)
ValuePair(GCS_Manoca,	4260)
ValuePair(GCS_Merchich,	4261)
ValuePair(GCS_Massawa,	4262)
ValuePair(GCS_Minna,	4263)
ValuePair(GCS_Mhast,	4264)
ValuePair(GCS_Monte_Mario,	4265)
ValuePair(GCS_M_poraloko,	4266)
ValuePair(GCS_NAD27,	4267)
ValuePair(GCS_NAD_Michigan,	4268)
ValuePair(GCS_NAD83,	4269)
ValuePair(GCS_Nahrwan_1967,	4270)
ValuePair(GCS_Naparima_1972,	4271)
ValuePair(GCS_GD49,	4272)
ValuePair(GCS_NGO_1948,	4273)
ValuePair(GCS_Datum_73,	4274)
ValuePair(GCS_NTF,	4275)
ValuePair(GCS_NSWC_9Z_2,	4276)
ValuePair(GCS_OSGB_1936,	4277)
ValuePair(GCS_OSGB70,	4278)
ValuePair(GCS_OS_SN80,	4279)
ValuePair(GCS_Padang,	4280)
ValuePair(GCS_Palestine_1923,	4281)
ValuePair(GCS_Pointe_Noire,	4282)
ValuePair(GCS_GDA94,	4283)
ValuePair(GCS_Pulkovo_1942,	4284)
ValuePair(GCS_Qatar,	4285)
ValuePair(GCS_Qatar_1948,	4286)
ValuePair(GCS_Qornoq,	4287)
ValuePair(GCS_Loma_Quintana,	4288)
ValuePair(GCS_Amersfoort,	4289)
ValuePair(GCS_RT38,	4290)
ValuePair(GCS_SAD69,	4291)
ValuePair(GCS_Sapper_Hill_1943,	4292)
ValuePair(GCS_Schwarzeck,	4293)
ValuePair(GCS_Segora,	4294)
ValuePair(GCS_Serindung,	4295)
ValuePair(GCS_Sudan,	4296)
ValuePair(GCS_Tananarive,	4297)
ValuePair(GCS_Timbalai_1948,	4298)
ValuePair(GCS_TM65,	4299)
ValuePair(GCS_TM75,	4300)
ValuePair(GCS_Tokyo,	4301)
ValuePair(GCS_Trinidad_1903,	4302)
ValuePair(GCS_TC_1948,	4303)
ValuePair(GCS_Voirol_1875,	4304)
ValuePair(GCS_Voirol_Unifie,	4305)
ValuePair(GCS_Bern_1938,	4306)
ValuePair(GCS_Nord_Sahara_1959,	4307)
ValuePair(GCS_Stockholm_1938,	4308)
ValuePair(GCS_Yacare,	4309)
ValuePair(GCS_Yoff,	4310)
ValuePair(GCS_Zanderij,	4311)
ValuePair(GCS_MGI,	4312)
ValuePair(GCS_Belge_1972,	4313)
ValuePair(GCS_DHDN,	4314)
ValuePair(GCS_Conakry_1905,	4315)
ValuePair(GCS_WGS_72,	4322)
ValuePair(GCS_WGS_72BE,	4324)
ValuePair(GCS_WGS_84,	4326)
ValuePair(GCS_Bern_1898_Bern,	4801)
ValuePair(GCS_Bogota_Bogota,	4802)
ValuePair(GCS_Lisbon_Lisbon,	4803)
ValuePair(GCS_Makassar_Jakarta,	4804)
ValuePair(GCS_MGI_Ferro,	4805)
ValuePair(GCS_Monte_Mario_Rome,	4806)
ValuePair(GCS_NTF_Paris,	4807)
ValuePair(GCS_Padang_Jakarta,	4808)
ValuePair(GCS_Belge_1950_Brussels,	4809)
ValuePair(GCS_Tananarive_Paris,	4810)
ValuePair(GCS_Voirol_1875_Paris,	4811)
ValuePair(GCS_Voirol_Unifie_Paris,	4812)
ValuePair(GCS_Batavia_Jakarta,	4813)
ValuePair(GCS_ATF_Paris,	4901)
ValuePair(GCS_NDG_Paris,	4902)
/* End of list */
`

// NameVal describes name and value
type NameVal struct {
	Name  string
	Value string
}

func toNameVals(src string) []NameVal {
	var res []NameVal
	lines := strings.Split(src, "\n")
	for _, s := range lines {
		if !strings.HasPrefix(s, "ValuePair") {
			continue
		}

		s = strings.TrimPrefix(s, "ValuePair(")
		s = strings.TrimSuffix(s, ")")
		s = strings.TrimSpace(s)
		parts := strings.Split(s, ",")
		if false && len(parts) != 2 {
			continue
		}
		name := strings.TrimSpace(parts[0])
		val := strings.TrimSpace(parts[1])
		res = append(res, NameVal{name, val})
	}
	return res
}

func genGo(src string) {
	nameVals := toNameVals(src)
	fmt.Printf(`
const (
  `)
	for _, nm := range nameVals {
		fmt.Printf("\t%s = %s\n", nm.Name, nm.Value)
	}
	fmt.Printf(")\n\n")

	fmt.Printf(`
func *Name(v int) string {
  switch v {
`)
	for _, nm := range nameVals {
		fmt.Printf(`	case %s:
		return "%s"
`, nm.Name, nm.Name)
	}
	fmt.Printf(`
  }
  return fmt.Sprintf("Unkown (%%d)", v)
}
`)
}

func main() {
	genGo(gcs_codes)
}
