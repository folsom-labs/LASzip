package main

import "fmt"

// https://github.com/smanders/libgeotiff/blob/master/epsg_pcs.inc
const (
	PCS_Hjorsey_1955_Lambert               = 3053
	PCS_ISN93_Lambert_1993                 = 3057
	PCS_ETRS89_Poland_CS2000_zone_5        = 2176
	PCS_ETRS89_Poland_CS2000_zone_6        = 2177
	PCS_ETRS89_Poland_CS2000_zone_7        = 2177
	PCS_ETRS89_Poland_CS2000_zone_8        = 2178
	PCS_ETRS89_Poland_CS92                 = 2180
	PCS_GGRS87_Greek_Grid                  = 2100
	PCS_KKJ_Finland_zone_1                 = 2391
	PCS_KKJ_Finland_zone_2                 = 2392
	PCS_KKJ_Finland_zone_3                 = 2393
	PCS_KKJ_Finland_zone_4                 = 2394
	PCS_RT90_2_5_gon_W                     = 2400
	PCS_Lietuvos_Koordinoei_Sistema_1994   = 2600
	PCS_Estonian_Coordinate_System_of_1992 = 3300
	PCS_HD72_EOV                           = 23700
	PCS_Dealul_Piscului_1970_Stereo_70     = 31700
	PCS_Adindan_UTM_zone_37N               = 20137
	PCS_Adindan_UTM_zone_38N               = 20138
	PCS_AGD66_AMG_zone_48                  = 20248
	PCS_AGD66_AMG_zone_49                  = 20249
	PCS_AGD66_AMG_zone_50                  = 20250
	PCS_AGD66_AMG_zone_51                  = 20251
	PCS_AGD66_AMG_zone_52                  = 20252
	PCS_AGD66_AMG_zone_53                  = 20253
	PCS_AGD66_AMG_zone_54                  = 20254
	PCS_AGD66_AMG_zone_55                  = 20255
	PCS_AGD66_AMG_zone_56                  = 20256
	PCS_AGD66_AMG_zone_57                  = 20257
	PCS_AGD66_AMG_zone_58                  = 20258
	PCS_AGD84_AMG_zone_48                  = 20348
	PCS_AGD84_AMG_zone_49                  = 20349
	PCS_AGD84_AMG_zone_50                  = 20350
	PCS_AGD84_AMG_zone_51                  = 20351
	PCS_AGD84_AMG_zone_52                  = 20352
	PCS_AGD84_AMG_zone_53                  = 20353
	PCS_AGD84_AMG_zone_54                  = 20354
	PCS_AGD84_AMG_zone_55                  = 20355
	PCS_AGD84_AMG_zone_56                  = 20356
	PCS_AGD84_AMG_zone_57                  = 20357
	PCS_AGD84_AMG_zone_58                  = 20358
	PCS_Ain_el_Abd_UTM_zone_37N            = 20437
	PCS_Ain_el_Abd_UTM_zone_38N            = 20438
	PCS_Ain_el_Abd_UTM_zone_39N            = 20439
	PCS_Ain_el_Abd_Bahrain_Grid            = 20499
	PCS_Afgooye_UTM_zone_38N               = 20538
	PCS_Afgooye_UTM_zone_39N               = 20539
	PCS_Lisbon_Portugese_Grid              = 20700
	PCS_Aratu_UTM_zone_22S                 = 20822
	PCS_Aratu_UTM_zone_23S                 = 20823
	PCS_Aratu_UTM_zone_24S                 = 20824
	PCS_Arc_1950_Lo13                      = 20973
	PCS_Arc_1950_Lo15                      = 20975
	PCS_Arc_1950_Lo17                      = 20977
	PCS_Arc_1950_Lo19                      = 20979
	PCS_Arc_1950_Lo21                      = 20981
	PCS_Arc_1950_Lo23                      = 20983
	PCS_Arc_1950_Lo25                      = 20985
	PCS_Arc_1950_Lo27                      = 20987
	PCS_Arc_1950_Lo29                      = 20989
	PCS_Arc_1950_Lo31                      = 20991
	PCS_Arc_1950_Lo33                      = 20993
	PCS_Arc_1950_Lo35                      = 20995
	PCS_Batavia_NEIEZ                      = 21100
	PCS_Batavia_UTM_zone_48S               = 21148
	PCS_Batavia_UTM_zone_49S               = 21149
	PCS_Batavia_UTM_zone_50S               = 21150
	PCS_Beijing_Gauss_zone_13              = 21413
	PCS_Beijing_Gauss_zone_14              = 21414
	PCS_Beijing_Gauss_zone_15              = 21415
	PCS_Beijing_Gauss_zone_16              = 21416
	PCS_Beijing_Gauss_zone_17              = 21417
	PCS_Beijing_Gauss_zone_18              = 21418
	PCS_Beijing_Gauss_zone_19              = 21419
	PCS_Beijing_Gauss_zone_20              = 21420
	PCS_Beijing_Gauss_zone_21              = 21421
	PCS_Beijing_Gauss_zone_22              = 21422
	PCS_Beijing_Gauss_zone_23              = 21423
	PCS_Beijing_Gauss_13N                  = 21473
	PCS_Beijing_Gauss_14N                  = 21474
	PCS_Beijing_Gauss_15N                  = 21475
	PCS_Beijing_Gauss_16N                  = 21476
	PCS_Beijing_Gauss_17N                  = 21477
	PCS_Beijing_Gauss_18N                  = 21478
	PCS_Beijing_Gauss_19N                  = 21479
	PCS_Beijing_Gauss_20N                  = 21480
	PCS_Beijing_Gauss_21N                  = 21481
	PCS_Beijing_Gauss_22N                  = 21482
	PCS_Beijing_Gauss_23N                  = 21483
	PCS_Belge_Lambert_50                   = 21500
	PCS_Bern_1898_Swiss_Old                = 21790
	PCS_Bogota_UTM_zone_17N                = 21817
	PCS_Bogota_UTM_zone_18N                = 21818
	PCS_Bogota_Colombia_3W                 = 21891
	PCS_Bogota_Colombia_Bogota             = 21892
	PCS_Bogota_Colombia_3E                 = 21893
	PCS_Bogota_Colombia_6E                 = 21894
	PCS_Camacupa_UTM_32S                   = 22032
	PCS_Camacupa_UTM_33S                   = 22033
	PCS_C_Inchauspe_Argentina_1            = 22191
	PCS_C_Inchauspe_Argentina_2            = 22192
	PCS_C_Inchauspe_Argentina_3            = 22193
	PCS_C_Inchauspe_Argentina_4            = 22194
	PCS_C_Inchauspe_Argentina_5            = 22195
	PCS_C_Inchauspe_Argentina_6            = 22196
	PCS_C_Inchauspe_Argentina_7            = 22197
	PCS_Carthage_UTM_zone_32N              = 22332
	PCS_Carthage_Nord_Tunisie              = 22391
	PCS_Carthage_Sud_Tunisie               = 22392
	PCS_Corrego_Alegre_UTM_23S             = 22523
	PCS_Corrego_Alegre_UTM_24S             = 22524
	PCS_Douala_UTM_zone_32N                = 22832
	PCS_Egypt_1907_Red_Belt                = 22992
	PCS_Egypt_1907_Purple_Belt             = 22993
	PCS_Egypt_1907_Ext_Purple              = 22994
	PCS_ED50_UTM_zone_28N                  = 23028
	PCS_ED50_UTM_zone_29N                  = 23029
	PCS_ED50_UTM_zone_30N                  = 23030
	PCS_ED50_UTM_zone_31N                  = 23031
	PCS_ED50_UTM_zone_32N                  = 23032
	PCS_ED50_UTM_zone_33N                  = 23033
	PCS_ED50_UTM_zone_34N                  = 23034
	PCS_ED50_UTM_zone_35N                  = 23035
	PCS_ED50_UTM_zone_36N                  = 23036
	PCS_ED50_UTM_zone_37N                  = 23037
	PCS_ED50_UTM_zone_38N                  = 23038
	PCS_Fahud_UTM_zone_39N                 = 23239
	PCS_Fahud_UTM_zone_40N                 = 23240
	PCS_Garoua_UTM_zone_33N                = 23433
	PCS_ID74_UTM_zone_46N                  = 23846
	PCS_ID74_UTM_zone_47N                  = 23847
	PCS_ID74_UTM_zone_48N                  = 23848
	PCS_ID74_UTM_zone_49N                  = 23849
	PCS_ID74_UTM_zone_50N                  = 23850
	PCS_ID74_UTM_zone_51N                  = 23851
	PCS_ID74_UTM_zone_52N                  = 23852
	PCS_ID74_UTM_zone_53N                  = 23853
	PCS_ID74_UTM_zone_46S                  = 23886
	PCS_ID74_UTM_zone_47S                  = 23887
	PCS_ID74_UTM_zone_48S                  = 23888
	PCS_ID74_UTM_zone_49S                  = 23889
	PCS_ID74_UTM_zone_50S                  = 23890
	PCS_ID74_UTM_zone_51S                  = 23891
	PCS_ID74_UTM_zone_52S                  = 23892
	PCS_ID74_UTM_zone_53S                  = 23893
	PCS_ID74_UTM_zone_54S                  = 23894
	PCS_Indian_1954_UTM_47N                = 23947
	PCS_Indian_1954_UTM_48N                = 23948
	PCS_Indian_1975_UTM_47N                = 24047
	PCS_Indian_1975_UTM_48N                = 24048
	PCS_Jamaica_1875_Old_Grid              = 24100
	PCS_JAD69_Jamaica_Grid                 = 24200
	PCS_Kalianpur_India_0                  = 24370
	PCS_Kalianpur_India_I                  = 24371
	PCS_Kalianpur_India_IIa                = 24372
	PCS_Kalianpur_India_IIIa               = 24373
	PCS_Kalianpur_India_IVa                = 24374
	PCS_Kalianpur_India_IIb                = 24382
	PCS_Kalianpur_India_IIIb               = 24383
	PCS_Kalianpur_India_IVb                = 24384
	PCS_Kertau_Singapore_Grid              = 24500
	PCS_Kertau_UTM_zone_47N                = 24547
	PCS_Kertau_UTM_zone_48N                = 24548
	PCS_La_Canoa_UTM_zone_20N              = 24720
	PCS_La_Canoa_UTM_zone_21N              = 24721
	PCS_PSAD56_UTM_zone_18N                = 24818
	PCS_PSAD56_UTM_zone_19N                = 24819
	PCS_PSAD56_UTM_zone_20N                = 24820
	PCS_PSAD56_UTM_zone_21N                = 24821
	PCS_PSAD56_UTM_zone_17S                = 24877
	PCS_PSAD56_UTM_zone_18S                = 24878
	PCS_PSAD56_UTM_zone_19S                = 24879
	PCS_PSAD56_UTM_zone_20S                = 24880
	PCS_PSAD56_Peru_west_zone              = 24891
	PCS_PSAD56_Peru_central                = 24892
	PCS_PSAD56_Peru_east_zone              = 24893
	PCS_Leigon_Ghana_Grid                  = 25000
	PCS_Lome_UTM_zone_31N                  = 25231
	PCS_Luzon_Philippines_I                = 25391
	PCS_Luzon_Philippines_II               = 25392
	PCS_Luzon_Philippines_III              = 25393
	PCS_Luzon_Philippines_IV               = 25394
	PCS_Luzon_Philippines_V                = 25395
	PCS_Makassar_NEIEZ                     = 25700
	PCS_Malongo_1987_UTM_32S               = 25932
	PCS_Merchich_Nord_Maroc                = 26191
	PCS_Merchich_Sud_Maroc                 = 26192
	PCS_Merchich_Sahara                    = 26193
	PCS_Massawa_UTM_zone_37N               = 26237
	PCS_Minna_UTM_zone_31N                 = 26331
	PCS_Minna_UTM_zone_32N                 = 26332
	PCS_Minna_Nigeria_West                 = 26391
	PCS_Minna_Nigeria_Mid_Belt             = 26392
	PCS_Minna_Nigeria_East                 = 26393
	PCS_Mhast_UTM_zone_32S                 = 26432
	PCS_Monte_Mario_Italy_1                = 26591
	PCS_Monte_Mario_Italy_2                = 26592
	PCS_M_poraloko_UTM_32N                 = 26632
	PCS_M_poraloko_UTM_32S                 = 26692
	PCS_NAD27_UTM_zone_3N                  = 26703
	PCS_NAD27_UTM_zone_4N                  = 26704
	PCS_NAD27_UTM_zone_5N                  = 26705
	PCS_NAD27_UTM_zone_6N                  = 26706
	PCS_NAD27_UTM_zone_7N                  = 26707
	PCS_NAD27_UTM_zone_8N                  = 26708
	PCS_NAD27_UTM_zone_9N                  = 26709
	PCS_NAD27_UTM_zone_10N                 = 26710
	PCS_NAD27_UTM_zone_11N                 = 26711
	PCS_NAD27_UTM_zone_12N                 = 26712
	PCS_NAD27_UTM_zone_13N                 = 26713
	PCS_NAD27_UTM_zone_14N                 = 26714
	PCS_NAD27_UTM_zone_15N                 = 26715
	PCS_NAD27_UTM_zone_16N                 = 26716
	PCS_NAD27_UTM_zone_17N                 = 26717
	PCS_NAD27_UTM_zone_18N                 = 26718
	PCS_NAD27_UTM_zone_19N                 = 26719
	PCS_NAD27_UTM_zone_20N                 = 26720
	PCS_NAD27_UTM_zone_21N                 = 26721
	PCS_NAD27_UTM_zone_22N                 = 26722
	PCS_NAD27_Alabama_East                 = 26729
	PCS_NAD27_Alabama_West                 = 26730
	PCS_NAD27_Alaska_zone_1                = 26731
	PCS_NAD27_Alaska_zone_2                = 26732
	PCS_NAD27_Alaska_zone_3                = 26733
	PCS_NAD27_Alaska_zone_4                = 26734
	PCS_NAD27_Alaska_zone_5                = 26735
	PCS_NAD27_Alaska_zone_6                = 26736
	PCS_NAD27_Alaska_zone_7                = 26737
	PCS_NAD27_Alaska_zone_8                = 26738
	PCS_NAD27_Alaska_zone_9                = 26739
	PCS_NAD27_Alaska_zone_10               = 26740
	PCS_NAD27_California_I                 = 26741
	PCS_NAD27_California_II                = 26742
	PCS_NAD27_California_III               = 26743
	PCS_NAD27_California_IV                = 26744
	PCS_NAD27_California_V                 = 26745
	PCS_NAD27_California_VI                = 26746
	PCS_NAD27_California_VII               = 26747
	PCS_NAD27_Arizona_East                 = 26748
	PCS_NAD27_Arizona_Central              = 26749
	PCS_NAD27_Arizona_West                 = 26750
	PCS_NAD27_Arkansas_North               = 26751
	PCS_NAD27_Arkansas_South               = 26752
	PCS_NAD27_Colorado_North               = 26753
	PCS_NAD27_Colorado_Central             = 26754
	PCS_NAD27_Colorado_South               = 26755
	PCS_NAD27_Connecticut                  = 26756
	PCS_NAD27_Delaware                     = 26757
	PCS_NAD27_Florida_East                 = 26758
	PCS_NAD27_Florida_West                 = 26759
	PCS_NAD27_Florida_North                = 26760
	PCS_NAD27_Hawaii_zone_1                = 26761
	PCS_NAD27_Hawaii_zone_2                = 26762
	PCS_NAD27_Hawaii_zone_3                = 26763
	PCS_NAD27_Hawaii_zone_4                = 26764
	PCS_NAD27_Hawaii_zone_5                = 26765
	PCS_NAD27_Georgia_East                 = 26766
	PCS_NAD27_Georgia_West                 = 26767
	PCS_NAD27_Idaho_East                   = 26768
	PCS_NAD27_Idaho_Central                = 26769
	PCS_NAD27_Idaho_West                   = 26770
	PCS_NAD27_Illinois_East                = 26771
	PCS_NAD27_Illinois_West                = 26772
	PCS_NAD27_Indiana_East                 = 26773
	PCS_NAD27_BLM_14N_feet                 = 26774
	PCS_NAD27_Indiana_West                 = 26774
	PCS_NAD27_BLM_15N_feet                 = 26775
	PCS_NAD27_Iowa_North                   = 26775
	PCS_NAD27_BLM_16N_feet                 = 26776
	PCS_NAD27_Iowa_South                   = 26776
	PCS_NAD27_BLM_17N_feet                 = 26777
	PCS_NAD27_Kansas_North                 = 26777
	PCS_NAD27_Kansas_South                 = 26778
	PCS_NAD27_Kentucky_North               = 26779
	PCS_NAD27_Kentucky_South               = 26780
	PCS_NAD27_Louisiana_North              = 26781
	PCS_NAD27_Louisiana_South              = 26782
	PCS_NAD27_Maine_East                   = 26783
	PCS_NAD27_Maine_West                   = 26784
	PCS_NAD27_Maryland                     = 26785
	PCS_NAD27_Massachusetts                = 26786
	PCS_NAD27_Massachusetts_Is             = 26787
	PCS_NAD27_Michigan_North               = 26788
	PCS_NAD27_Michigan_Central             = 26789
	PCS_NAD27_Michigan_South               = 26790
	PCS_NAD27_Minnesota_North              = 26791
	PCS_NAD27_Minnesota_Cent               = 26792
	PCS_NAD27_Minnesota_South              = 26793
	PCS_NAD27_Mississippi_East             = 26794
	PCS_NAD27_Mississippi_West             = 26795
	PCS_NAD27_Missouri_East                = 26796
	PCS_NAD27_Missouri_Central             = 26797
	PCS_NAD27_Missouri_West                = 26798
	PCS_NAD_Michigan_Michigan_East         = 26801
	PCS_NAD_Michigan_Michigan_Old_Central  = 26802
	PCS_NAD_Michigan_Michigan_West         = 26803
	PCS_NAD83_UTM_zone_3N                  = 26903
	PCS_NAD83_UTM_zone_4N                  = 26904
	PCS_NAD83_UTM_zone_5N                  = 26905
	PCS_NAD83_UTM_zone_6N                  = 26906
	PCS_NAD83_UTM_zone_7N                  = 26907
	PCS_NAD83_UTM_zone_8N                  = 26908
	PCS_NAD83_UTM_zone_9N                  = 26909
	PCS_NAD83_UTM_zone_10N                 = 26910
	PCS_NAD83_UTM_zone_11N                 = 26911
	PCS_NAD83_UTM_zone_12N                 = 26912
	PCS_NAD83_UTM_zone_13N                 = 26913
	PCS_NAD83_UTM_zone_14N                 = 26914
	PCS_NAD83_UTM_zone_15N                 = 26915
	PCS_NAD83_UTM_zone_16N                 = 26916
	PCS_NAD83_UTM_zone_17N                 = 26917
	PCS_NAD83_UTM_zone_18N                 = 26918
	PCS_NAD83_UTM_zone_19N                 = 26919
	PCS_NAD83_UTM_zone_20N                 = 26920
	PCS_NAD83_UTM_zone_21N                 = 26921
	PCS_NAD83_UTM_zone_22N                 = 26922
	PCS_NAD83_UTM_zone_23N                 = 26923
	PCS_NAD83_Alabama_East                 = 26929
	PCS_NAD83_Alabama_West                 = 26930
	PCS_NAD83_Alaska_zone_1                = 26931
	PCS_NAD83_Alaska_zone_2                = 26932
	PCS_NAD83_Alaska_zone_3                = 26933
	PCS_NAD83_Alaska_zone_4                = 26934
	PCS_NAD83_Alaska_zone_5                = 26935
	PCS_NAD83_Alaska_zone_6                = 26936
	PCS_NAD83_Alaska_zone_7                = 26937
	PCS_NAD83_Alaska_zone_8                = 26938
	PCS_NAD83_Alaska_zone_9                = 26939
	PCS_NAD83_Alaska_zone_10               = 26940
	PCS_NAD83_California_1                 = 26941
	PCS_NAD83_California_2                 = 26942
	PCS_NAD83_California_3                 = 26943
	PCS_NAD83_California_4                 = 26944
	PCS_NAD83_California_5                 = 26945
	PCS_NAD83_California_6                 = 26946
	PCS_NAD83_Arizona_East                 = 26948
	PCS_NAD83_Arizona_Central              = 26949
	PCS_NAD83_Arizona_West                 = 26950
	PCS_NAD83_Arkansas_North               = 26951
	PCS_NAD83_Arkansas_South               = 26952
	PCS_NAD83_Colorado_North               = 26953
	PCS_NAD83_Colorado_Central             = 26954
	PCS_NAD83_Colorado_South               = 26955
	PCS_NAD83_Connecticut                  = 26956
	PCS_NAD83_Delaware                     = 26957
	PCS_NAD83_Florida_East                 = 26958
	PCS_NAD83_Florida_West                 = 26959
	PCS_NAD83_Florida_North                = 26960
	PCS_NAD83_Hawaii_zone_1                = 26961
	PCS_NAD83_Hawaii_zone_2                = 26962
	PCS_NAD83_Hawaii_zone_3                = 26963
	PCS_NAD83_Hawaii_zone_4                = 26964
	PCS_NAD83_Hawaii_zone_5                = 26965
	PCS_NAD83_Georgia_East                 = 26966
	PCS_NAD83_Georgia_West                 = 26967
	PCS_NAD83_Idaho_East                   = 26968
	PCS_NAD83_Idaho_Central                = 26969
	PCS_NAD83_Idaho_West                   = 26970
	PCS_NAD83_Illinois_East                = 26971
	PCS_NAD83_Illinois_West                = 26972
	PCS_NAD83_Indiana_East                 = 26973
	PCS_NAD83_Indiana_West                 = 26974
	PCS_NAD83_Iowa_North                   = 26975
	PCS_NAD83_Iowa_South                   = 26976
	PCS_NAD83_Kansas_North                 = 26977
	PCS_NAD83_Kansas_South                 = 26978
	PCS_NAD83_Kentucky_North               = 2205
	PCS_NAD83_Kentucky_South               = 26980
	PCS_NAD83_Louisiana_North              = 26981
	PCS_NAD83_Louisiana_South              = 26982
	PCS_NAD83_Maine_East                   = 26983
	PCS_NAD83_Maine_West                   = 26984
	PCS_NAD83_Maryland                     = 26985
	PCS_NAD83_Massachusetts                = 26986
	PCS_NAD83_Massachusetts_Is             = 26987
	PCS_NAD83_Michigan_North               = 26988
	PCS_NAD83_Michigan_Central             = 26989
	PCS_NAD83_Michigan_South               = 26990
	PCS_NAD83_Minnesota_North              = 26991
	PCS_NAD83_Minnesota_Cent               = 26992
	PCS_NAD83_Minnesota_South              = 26993
	PCS_NAD83_Mississippi_East             = 26994
	PCS_NAD83_Mississippi_West             = 26995
	PCS_NAD83_Missouri_East                = 26996
	PCS_NAD83_Missouri_Central             = 26997
	PCS_NAD83_Missouri_West                = 26998
	PCS_Nahrwan_1967_UTM_38N               = 27038
	PCS_Nahrwan_1967_UTM_39N               = 27039
	PCS_Nahrwan_1967_UTM_40N               = 27040
	PCS_Naparima_UTM_20N                   = 27120
	PCS_GD49_NZ_Map_Grid                   = 27200
	PCS_GD49_North_Island_Grid             = 27291
	PCS_GD49_South_Island_Grid             = 27292
	PCS_Datum_73_UTM_zone_29N              = 27429
	PCS_ATF_Nord_de_Guerre                 = 27500
	PCS_NTF_France_I                       = 27581
	PCS_NTF_France_II                      = 27582
	PCS_NTF_France_III                     = 27583
	PCS_NTF_Nord_France                    = 27591
	PCS_NTF_Centre_France                  = 27592
	PCS_NTF_Sud_France                     = 27593
	PCS_British_National_Grid              = 27700
	PCS_Point_Noire_UTM_32S                = 28232
	PCS_GDA94_MGA_zone_48                  = 28348
	PCS_GDA94_MGA_zone_49                  = 28349
	PCS_GDA94_MGA_zone_50                  = 28350
	PCS_GDA94_MGA_zone_51                  = 28351
	PCS_GDA94_MGA_zone_52                  = 28352
	PCS_GDA94_MGA_zone_53                  = 28353
	PCS_GDA94_MGA_zone_54                  = 28354
	PCS_GDA94_MGA_zone_55                  = 28355
	PCS_GDA94_MGA_zone_56                  = 28356
	PCS_GDA94_MGA_zone_57                  = 28357
	PCS_GDA94_MGA_zone_58                  = 28358
	PCS_Pulkovo_Gauss_zone_4               = 28404
	PCS_Pulkovo_Gauss_zone_5               = 28405
	PCS_Pulkovo_Gauss_zone_6               = 28406
	PCS_Pulkovo_Gauss_zone_7               = 28407
	PCS_Pulkovo_Gauss_zone_8               = 28408
	PCS_Pulkovo_Gauss_zone_9               = 28409
	PCS_Pulkovo_Gauss_zone_10              = 28410
	PCS_Pulkovo_Gauss_zone_11              = 28411
	PCS_Pulkovo_Gauss_zone_12              = 28412
	PCS_Pulkovo_Gauss_zone_13              = 28413
	PCS_Pulkovo_Gauss_zone_14              = 28414
	PCS_Pulkovo_Gauss_zone_15              = 28415
	PCS_Pulkovo_Gauss_zone_16              = 28416
	PCS_Pulkovo_Gauss_zone_17              = 28417
	PCS_Pulkovo_Gauss_zone_18              = 28418
	PCS_Pulkovo_Gauss_zone_19              = 28419
	PCS_Pulkovo_Gauss_zone_20              = 28420
	PCS_Pulkovo_Gauss_zone_21              = 28421
	PCS_Pulkovo_Gauss_zone_22              = 28422
	PCS_Pulkovo_Gauss_zone_23              = 28423
	PCS_Pulkovo_Gauss_zone_24              = 28424
	PCS_Pulkovo_Gauss_zone_25              = 28425
	PCS_Pulkovo_Gauss_zone_26              = 28426
	PCS_Pulkovo_Gauss_zone_27              = 28427
	PCS_Pulkovo_Gauss_zone_28              = 28428
	PCS_Pulkovo_Gauss_zone_29              = 28429
	PCS_Pulkovo_Gauss_zone_30              = 28430
	PCS_Pulkovo_Gauss_zone_31              = 28431
	PCS_Pulkovo_Gauss_zone_32              = 28432
	PCS_Pulkovo_Gauss_4N                   = 28464
	PCS_Pulkovo_Gauss_5N                   = 28465
	PCS_Pulkovo_Gauss_6N                   = 28466
	PCS_Pulkovo_Gauss_7N                   = 28467
	PCS_Pulkovo_Gauss_8N                   = 28468
	PCS_Pulkovo_Gauss_9N                   = 28469
	PCS_Pulkovo_Gauss_10N                  = 28470
	PCS_Pulkovo_Gauss_11N                  = 28471
	PCS_Pulkovo_Gauss_12N                  = 28472
	PCS_Pulkovo_Gauss_13N                  = 28473
	PCS_Pulkovo_Gauss_14N                  = 28474
	PCS_Pulkovo_Gauss_15N                  = 28475
	PCS_Pulkovo_Gauss_16N                  = 28476
	PCS_Pulkovo_Gauss_17N                  = 28477
	PCS_Pulkovo_Gauss_18N                  = 28478
	PCS_Pulkovo_Gauss_19N                  = 28479
	PCS_Pulkovo_Gauss_20N                  = 28480
	PCS_Pulkovo_Gauss_21N                  = 28481
	PCS_Pulkovo_Gauss_22N                  = 28482
	PCS_Pulkovo_Gauss_23N                  = 28483
	PCS_Pulkovo_Gauss_24N                  = 28484
	PCS_Pulkovo_Gauss_25N                  = 28485
	PCS_Pulkovo_Gauss_26N                  = 28486
	PCS_Pulkovo_Gauss_27N                  = 28487
	PCS_Pulkovo_Gauss_28N                  = 28488
	PCS_Pulkovo_Gauss_29N                  = 28489
	PCS_Pulkovo_Gauss_30N                  = 28490
	PCS_Pulkovo_Gauss_31N                  = 28491
	PCS_Pulkovo_Gauss_32N                  = 28492
	PCS_Qatar_National_Grid                = 28600
	PCS_RD_Netherlands_Old                 = 28991
	PCS_RD_Netherlands_New                 = 28992
	PCS_SAD69_UTM_zone_18N                 = 29118
	PCS_SAD69_UTM_zone_19N                 = 29119
	PCS_SAD69_UTM_zone_20N                 = 29120
	PCS_SAD69_UTM_zone_21N                 = 29121
	PCS_SAD69_UTM_zone_22N                 = 29122
	PCS_SAD69_UTM_zone_17S                 = 29177
	PCS_SAD69_UTM_zone_18S                 = 29178
	PCS_SAD69_UTM_zone_19S                 = 29179
	PCS_SAD69_UTM_zone_20S                 = 29180
	PCS_SAD69_UTM_zone_21S                 = 29181
	PCS_SAD69_UTM_zone_22S                 = 29182
	PCS_SAD69_UTM_zone_23S                 = 29183
	PCS_SAD69_UTM_zone_24S                 = 29184
	PCS_SAD69_UTM_zone_25S                 = 29185
	PCS_Sapper_Hill_UTM_20S                = 29220
	PCS_Sapper_Hill_UTM_21S                = 29221
	PCS_Schwarzeck_UTM_33S                 = 29333
	PCS_Sudan_UTM_zone_35N                 = 29635
	PCS_Sudan_UTM_zone_36N                 = 29636
	PCS_Tananarive_Laborde                 = 29700
	PCS_Tananarive_UTM_38S                 = 29738
	PCS_Tananarive_UTM_39S                 = 29739
	PCS_Timbalai_1948_Borneo               = 29800
	PCS_Timbalai_1948_UTM_49N              = 29849
	PCS_Timbalai_1948_UTM_50N              = 29850
	PCS_TM65_Irish_Nat_Grid                = 29900
	PCS_Trinidad_1903_Trinidad             = 30200
	PCS_TC_1948_UTM_zone_39N               = 30339
	PCS_TC_1948_UTM_zone_40N               = 30340
	PCS_Voirol_N_Algerie_ancien            = 30491
	PCS_Voirol_S_Algerie_ancien            = 30492
	PCS_Voirol_Unifie_N_Algerie            = 30591
	PCS_Voirol_Unifie_S_Algerie            = 30592
	PCS_Bern_1938_Swiss_New                = 30600
	PCS_Nord_Sahara_UTM_29N                = 30729
	PCS_Nord_Sahara_UTM_30N                = 30730
	PCS_Nord_Sahara_UTM_31N                = 30731
	PCS_Nord_Sahara_UTM_32N                = 30732
	PCS_Yoff_UTM_zone_28N                  = 31028
	PCS_Zanderij_UTM_zone_21N              = 31121
	PCS_MGI_Austria_West                   = 31291
	PCS_MGI_Austria_Central                = 31292
	PCS_MGI_Austria_East                   = 31293
	PCS_Belge_Lambert_72                   = 31300
	PCS_DHDN_Germany_zone_1                = 31491
	PCS_DHDN_Germany_zone_2                = 31492
	PCS_DHDN_Germany_zone_3                = 31493
	PCS_DHDN_Germany_zone_4                = 31494
	PCS_DHDN_Germany_zone_5                = 31495
	PCS_NAD27_Montana_North                = 32001
	PCS_NAD27_Montana_Central              = 32002
	PCS_NAD27_Montana_South                = 32003
	PCS_NAD27_Nebraska_North               = 32005
	PCS_NAD27_Nebraska_South               = 32006
	PCS_NAD27_Nevada_East                  = 32007
	PCS_NAD27_Nevada_Central               = 32008
	PCS_NAD27_Nevada_West                  = 32009
	PCS_NAD27_New_Hampshire                = 32010
	PCS_NAD27_New_Jersey                   = 32011
	PCS_NAD27_New_Mexico_East              = 32012
	PCS_NAD27_New_Mexico_Cent              = 32013
	PCS_NAD27_New_Mexico_West              = 32014
	PCS_NAD27_New_York_East                = 32015
	PCS_NAD27_New_York_Central             = 32016
	PCS_NAD27_New_York_West                = 32017
	PCS_NAD27_New_York_Long_Is             = 32018
	PCS_NAD27_North_Carolina               = 32019
	PCS_NAD27_North_Dakota_N               = 32020
	PCS_NAD27_North_Dakota_S               = 32021
	PCS_NAD27_Ohio_North                   = 32022
	PCS_NAD27_Ohio_South                   = 32023
	PCS_NAD27_Oklahoma_North               = 32024
	PCS_NAD27_Oklahoma_South               = 32025
	PCS_NAD27_Oregon_North                 = 32026
	PCS_NAD27_Oregon_South                 = 32027
	PCS_NAD27_Pennsylvania_N               = 32028
	PCS_NAD27_Pennsylvania_S               = 32029
	PCS_NAD27_Rhode_Island                 = 32030
	PCS_NAD27_South_Carolina_N             = 32031
	PCS_NAD27_South_Carolina_S             = 32033
	PCS_NAD27_South_Dakota_N               = 32034
	PCS_NAD27_South_Dakota_S               = 32035
	PCS_NAD27_Tennessee                    = 2204
	PCS_NAD27_Texas_North                  = 32037
	PCS_NAD27_Texas_North_Cen              = 32038
	PCS_NAD27_Texas_Central                = 32039
	PCS_NAD27_Texas_South_Cen              = 32040
	PCS_NAD27_Texas_South                  = 32041
	PCS_NAD27_Utah_North                   = 32042
	PCS_NAD27_Utah_Central                 = 32043
	PCS_NAD27_Utah_South                   = 32044
	PCS_NAD27_Vermont                      = 32045
	PCS_NAD27_Virginia_North               = 32046
	PCS_NAD27_Virginia_South               = 32047
	PCS_NAD27_Washington_North             = 32048
	PCS_NAD27_Washington_South             = 32049
	PCS_NAD27_West_Virginia_N              = 32050
	PCS_NAD27_West_Virginia_S              = 32051
	PCS_NAD27_Wisconsin_North              = 32052
	PCS_NAD27_Wisconsin_Cen                = 32053
	PCS_NAD27_Wisconsin_South              = 32054
	PCS_NAD27_Wyoming_East                 = 32055
	PCS_NAD27_Wyoming_E_Cen                = 32056
	PCS_NAD27_Wyoming_W_Cen                = 32057
	PCS_NAD27_Wyoming_West                 = 32058
	PCS_NAD27_Puerto_Rico                  = 32059
	PCS_NAD27_St_Croix                     = 32060
	PCS_NAD83_Montana                      = 32100
	PCS_NAD83_Nebraska                     = 32104
	PCS_NAD83_Nevada_East                  = 32107
	PCS_NAD83_Nevada_Central               = 32108
	PCS_NAD83_Nevada_West                  = 32109
	PCS_NAD83_New_Hampshire                = 32110
	PCS_NAD83_New_Jersey                   = 32111
	PCS_NAD83_New_Mexico_East              = 32112
	PCS_NAD83_New_Mexico_Cent              = 32113
	PCS_NAD83_New_Mexico_West              = 32114
	PCS_NAD83_New_York_East                = 32115
	PCS_NAD83_New_York_Central             = 32116
	PCS_NAD83_New_York_West                = 32117
	PCS_NAD83_New_York_Long_Is             = 32118
	PCS_NAD83_North_Carolina               = 32119
	PCS_NAD83_North_Dakota_N               = 32120
	PCS_NAD83_North_Dakota_S               = 32121
	PCS_NAD83_Ohio_North                   = 32122
	PCS_NAD83_Ohio_South                   = 32123
	PCS_NAD83_Oklahoma_North               = 32124
	PCS_NAD83_Oklahoma_South               = 32125
	PCS_NAD83_Oregon_North                 = 32126
	PCS_NAD83_Oregon_South                 = 32127
	PCS_NAD83_Pennsylvania_N               = 32128
	PCS_NAD83_Pennsylvania_S               = 32129
	PCS_NAD83_Rhode_Island                 = 32130
	PCS_NAD83_South_Carolina               = 32133
	PCS_NAD83_South_Dakota_N               = 32134
	PCS_NAD83_South_Dakota_S               = 32135
	PCS_NAD83_Tennessee                    = 32136
	PCS_NAD83_Texas_North                  = 32137
	PCS_NAD83_Texas_North_Cen              = 32138
	PCS_NAD83_Texas_Central                = 32139
	PCS_NAD83_Texas_South_Cen              = 32140
	PCS_NAD83_Texas_South                  = 32141
	PCS_NAD83_Utah_North                   = 32142
	PCS_NAD83_Utah_Central                 = 32143
	PCS_NAD83_Utah_South                   = 32144
	PCS_NAD83_Vermont                      = 32145
	PCS_NAD83_Virginia_North               = 32146
	PCS_NAD83_Virginia_South               = 32147
	PCS_NAD83_Washington_North             = 32148
	PCS_NAD83_Washington_South             = 32149
	PCS_NAD83_West_Virginia_N              = 32150
	PCS_NAD83_West_Virginia_S              = 32151
	PCS_NAD83_Wisconsin_North              = 32152
	PCS_NAD83_Wisconsin_Cen                = 32153
	PCS_NAD83_Wisconsin_South              = 32154
	PCS_NAD83_Wyoming_East                 = 32155
	PCS_NAD83_Wyoming_E_Cen                = 32156
	PCS_NAD83_Wyoming_W_Cen                = 32157
	PCS_NAD83_Wyoming_West                 = 32158
	PCS_NAD83_Puerto_Rico_Virgin_Is        = 32161
	PCS_WGS72_UTM_zone_1N                  = 32201
	PCS_WGS72_UTM_zone_2N                  = 32202
	PCS_WGS72_UTM_zone_3N                  = 32203
	PCS_WGS72_UTM_zone_4N                  = 32204
	PCS_WGS72_UTM_zone_5N                  = 32205
	PCS_WGS72_UTM_zone_6N                  = 32206
	PCS_WGS72_UTM_zone_7N                  = 32207
	PCS_WGS72_UTM_zone_8N                  = 32208
	PCS_WGS72_UTM_zone_9N                  = 32209
	PCS_WGS72_UTM_zone_10N                 = 32210
	PCS_WGS72_UTM_zone_11N                 = 32211
	PCS_WGS72_UTM_zone_12N                 = 32212
	PCS_WGS72_UTM_zone_13N                 = 32213
	PCS_WGS72_UTM_zone_14N                 = 32214
	PCS_WGS72_UTM_zone_15N                 = 32215
	PCS_WGS72_UTM_zone_16N                 = 32216
	PCS_WGS72_UTM_zone_17N                 = 32217
	PCS_WGS72_UTM_zone_18N                 = 32218
	PCS_WGS72_UTM_zone_19N                 = 32219
	PCS_WGS72_UTM_zone_20N                 = 32220
	PCS_WGS72_UTM_zone_21N                 = 32221
	PCS_WGS72_UTM_zone_22N                 = 32222
	PCS_WGS72_UTM_zone_23N                 = 32223
	PCS_WGS72_UTM_zone_24N                 = 32224
	PCS_WGS72_UTM_zone_25N                 = 32225
	PCS_WGS72_UTM_zone_26N                 = 32226
	PCS_WGS72_UTM_zone_27N                 = 32227
	PCS_WGS72_UTM_zone_28N                 = 32228
	PCS_WGS72_UTM_zone_29N                 = 32229
	PCS_WGS72_UTM_zone_30N                 = 32230
	PCS_WGS72_UTM_zone_31N                 = 32231
	PCS_WGS72_UTM_zone_32N                 = 32232
	PCS_WGS72_UTM_zone_33N                 = 32233
	PCS_WGS72_UTM_zone_34N                 = 32234
	PCS_WGS72_UTM_zone_35N                 = 32235
	PCS_WGS72_UTM_zone_36N                 = 32236
	PCS_WGS72_UTM_zone_37N                 = 32237
	PCS_WGS72_UTM_zone_38N                 = 32238
	PCS_WGS72_UTM_zone_39N                 = 32239
	PCS_WGS72_UTM_zone_40N                 = 32240
	PCS_WGS72_UTM_zone_41N                 = 32241
	PCS_WGS72_UTM_zone_42N                 = 32242
	PCS_WGS72_UTM_zone_43N                 = 32243
	PCS_WGS72_UTM_zone_44N                 = 32244
	PCS_WGS72_UTM_zone_45N                 = 32245
	PCS_WGS72_UTM_zone_46N                 = 32246
	PCS_WGS72_UTM_zone_47N                 = 32247
	PCS_WGS72_UTM_zone_48N                 = 32248
	PCS_WGS72_UTM_zone_49N                 = 32249
	PCS_WGS72_UTM_zone_50N                 = 32250
	PCS_WGS72_UTM_zone_51N                 = 32251
	PCS_WGS72_UTM_zone_52N                 = 32252
	PCS_WGS72_UTM_zone_53N                 = 32253
	PCS_WGS72_UTM_zone_54N                 = 32254
	PCS_WGS72_UTM_zone_55N                 = 32255
	PCS_WGS72_UTM_zone_56N                 = 32256
	PCS_WGS72_UTM_zone_57N                 = 32257
	PCS_WGS72_UTM_zone_58N                 = 32258
	PCS_WGS72_UTM_zone_59N                 = 32259
	PCS_WGS72_UTM_zone_60N                 = 32260
	PCS_WGS72_UTM_zone_1S                  = 32301
	PCS_WGS72_UTM_zone_2S                  = 32302
	PCS_WGS72_UTM_zone_3S                  = 32303
	PCS_WGS72_UTM_zone_4S                  = 32304
	PCS_WGS72_UTM_zone_5S                  = 32305
	PCS_WGS72_UTM_zone_6S                  = 32306
	PCS_WGS72_UTM_zone_7S                  = 32307
	PCS_WGS72_UTM_zone_8S                  = 32308
	PCS_WGS72_UTM_zone_9S                  = 32309
	PCS_WGS72_UTM_zone_10S                 = 32310
	PCS_WGS72_UTM_zone_11S                 = 32311
	PCS_WGS72_UTM_zone_12S                 = 32312
	PCS_WGS72_UTM_zone_13S                 = 32313
	PCS_WGS72_UTM_zone_14S                 = 32314
	PCS_WGS72_UTM_zone_15S                 = 32315
	PCS_WGS72_UTM_zone_16S                 = 32316
	PCS_WGS72_UTM_zone_17S                 = 32317
	PCS_WGS72_UTM_zone_18S                 = 32318
	PCS_WGS72_UTM_zone_19S                 = 32319
	PCS_WGS72_UTM_zone_20S                 = 32320
	PCS_WGS72_UTM_zone_21S                 = 32321
	PCS_WGS72_UTM_zone_22S                 = 32322
	PCS_WGS72_UTM_zone_23S                 = 32323
	PCS_WGS72_UTM_zone_24S                 = 32324
	PCS_WGS72_UTM_zone_25S                 = 32325
	PCS_WGS72_UTM_zone_26S                 = 32326
	PCS_WGS72_UTM_zone_27S                 = 32327
	PCS_WGS72_UTM_zone_28S                 = 32328
	PCS_WGS72_UTM_zone_29S                 = 32329
	PCS_WGS72_UTM_zone_30S                 = 32330
	PCS_WGS72_UTM_zone_31S                 = 32331
	PCS_WGS72_UTM_zone_32S                 = 32332
	PCS_WGS72_UTM_zone_33S                 = 32333
	PCS_WGS72_UTM_zone_34S                 = 32334
	PCS_WGS72_UTM_zone_35S                 = 32335
	PCS_WGS72_UTM_zone_36S                 = 32336
	PCS_WGS72_UTM_zone_37S                 = 32337
	PCS_WGS72_UTM_zone_38S                 = 32338
	PCS_WGS72_UTM_zone_39S                 = 32339
	PCS_WGS72_UTM_zone_40S                 = 32340
	PCS_WGS72_UTM_zone_41S                 = 32341
	PCS_WGS72_UTM_zone_42S                 = 32342
	PCS_WGS72_UTM_zone_43S                 = 32343
	PCS_WGS72_UTM_zone_44S                 = 32344
	PCS_WGS72_UTM_zone_45S                 = 32345
	PCS_WGS72_UTM_zone_46S                 = 32346
	PCS_WGS72_UTM_zone_47S                 = 32347
	PCS_WGS72_UTM_zone_48S                 = 32348
	PCS_WGS72_UTM_zone_49S                 = 32349
	PCS_WGS72_UTM_zone_50S                 = 32350
	PCS_WGS72_UTM_zone_51S                 = 32351
	PCS_WGS72_UTM_zone_52S                 = 32352
	PCS_WGS72_UTM_zone_53S                 = 32353
	PCS_WGS72_UTM_zone_54S                 = 32354
	PCS_WGS72_UTM_zone_55S                 = 32355
	PCS_WGS72_UTM_zone_56S                 = 32356
	PCS_WGS72_UTM_zone_57S                 = 32357
	PCS_WGS72_UTM_zone_58S                 = 32358
	PCS_WGS72_UTM_zone_59S                 = 32359
	PCS_WGS72_UTM_zone_60S                 = 32360
	PCS_WGS72BE_UTM_zone_1N                = 32401
	PCS_WGS72BE_UTM_zone_2N                = 32402
	PCS_WGS72BE_UTM_zone_3N                = 32403
	PCS_WGS72BE_UTM_zone_4N                = 32404
	PCS_WGS72BE_UTM_zone_5N                = 32405
	PCS_WGS72BE_UTM_zone_6N                = 32406
	PCS_WGS72BE_UTM_zone_7N                = 32407
	PCS_WGS72BE_UTM_zone_8N                = 32408
	PCS_WGS72BE_UTM_zone_9N                = 32409
	PCS_WGS72BE_UTM_zone_10N               = 32410
	PCS_WGS72BE_UTM_zone_11N               = 32411
	PCS_WGS72BE_UTM_zone_12N               = 32412
	PCS_WGS72BE_UTM_zone_13N               = 32413
	PCS_WGS72BE_UTM_zone_14N               = 32414
	PCS_WGS72BE_UTM_zone_15N               = 32415
	PCS_WGS72BE_UTM_zone_16N               = 32416
	PCS_WGS72BE_UTM_zone_17N               = 32417
	PCS_WGS72BE_UTM_zone_18N               = 32418
	PCS_WGS72BE_UTM_zone_19N               = 32419
	PCS_WGS72BE_UTM_zone_20N               = 32420
	PCS_WGS72BE_UTM_zone_21N               = 32421
	PCS_WGS72BE_UTM_zone_22N               = 32422
	PCS_WGS72BE_UTM_zone_23N               = 32423
	PCS_WGS72BE_UTM_zone_24N               = 32424
	PCS_WGS72BE_UTM_zone_25N               = 32425
	PCS_WGS72BE_UTM_zone_26N               = 32426
	PCS_WGS72BE_UTM_zone_27N               = 32427
	PCS_WGS72BE_UTM_zone_28N               = 32428
	PCS_WGS72BE_UTM_zone_29N               = 32429
	PCS_WGS72BE_UTM_zone_30N               = 32430
	PCS_WGS72BE_UTM_zone_31N               = 32431
	PCS_WGS72BE_UTM_zone_32N               = 32432
	PCS_WGS72BE_UTM_zone_33N               = 32433
	PCS_WGS72BE_UTM_zone_34N               = 32434
	PCS_WGS72BE_UTM_zone_35N               = 32435
	PCS_WGS72BE_UTM_zone_36N               = 32436
	PCS_WGS72BE_UTM_zone_37N               = 32437
	PCS_WGS72BE_UTM_zone_38N               = 32438
	PCS_WGS72BE_UTM_zone_39N               = 32439
	PCS_WGS72BE_UTM_zone_40N               = 32440
	PCS_WGS72BE_UTM_zone_41N               = 32441
	PCS_WGS72BE_UTM_zone_42N               = 32442
	PCS_WGS72BE_UTM_zone_43N               = 32443
	PCS_WGS72BE_UTM_zone_44N               = 32444
	PCS_WGS72BE_UTM_zone_45N               = 32445
	PCS_WGS72BE_UTM_zone_46N               = 32446
	PCS_WGS72BE_UTM_zone_47N               = 32447
	PCS_WGS72BE_UTM_zone_48N               = 32448
	PCS_WGS72BE_UTM_zone_49N               = 32449
	PCS_WGS72BE_UTM_zone_50N               = 32450
	PCS_WGS72BE_UTM_zone_51N               = 32451
	PCS_WGS72BE_UTM_zone_52N               = 32452
	PCS_WGS72BE_UTM_zone_53N               = 32453
	PCS_WGS72BE_UTM_zone_54N               = 32454
	PCS_WGS72BE_UTM_zone_55N               = 32455
	PCS_WGS72BE_UTM_zone_56N               = 32456
	PCS_WGS72BE_UTM_zone_57N               = 32457
	PCS_WGS72BE_UTM_zone_58N               = 32458
	PCS_WGS72BE_UTM_zone_59N               = 32459
	PCS_WGS72BE_UTM_zone_60N               = 32460
	PCS_WGS72BE_UTM_zone_1S                = 32501
	PCS_WGS72BE_UTM_zone_2S                = 32502
	PCS_WGS72BE_UTM_zone_3S                = 32503
	PCS_WGS72BE_UTM_zone_4S                = 32504
	PCS_WGS72BE_UTM_zone_5S                = 32505
	PCS_WGS72BE_UTM_zone_6S                = 32506
	PCS_WGS72BE_UTM_zone_7S                = 32507
	PCS_WGS72BE_UTM_zone_8S                = 32508
	PCS_WGS72BE_UTM_zone_9S                = 32509
	PCS_WGS72BE_UTM_zone_10S               = 32510
	PCS_WGS72BE_UTM_zone_11S               = 32511
	PCS_WGS72BE_UTM_zone_12S               = 32512
	PCS_WGS72BE_UTM_zone_13S               = 32513
	PCS_WGS72BE_UTM_zone_14S               = 32514
	PCS_WGS72BE_UTM_zone_15S               = 32515
	PCS_WGS72BE_UTM_zone_16S               = 32516
	PCS_WGS72BE_UTM_zone_17S               = 32517
	PCS_WGS72BE_UTM_zone_18S               = 32518
	PCS_WGS72BE_UTM_zone_19S               = 32519
	PCS_WGS72BE_UTM_zone_20S               = 32520
	PCS_WGS72BE_UTM_zone_21S               = 32521
	PCS_WGS72BE_UTM_zone_22S               = 32522
	PCS_WGS72BE_UTM_zone_23S               = 32523
	PCS_WGS72BE_UTM_zone_24S               = 32524
	PCS_WGS72BE_UTM_zone_25S               = 32525
	PCS_WGS72BE_UTM_zone_26S               = 32526
	PCS_WGS72BE_UTM_zone_27S               = 32527
	PCS_WGS72BE_UTM_zone_28S               = 32528
	PCS_WGS72BE_UTM_zone_29S               = 32529
	PCS_WGS72BE_UTM_zone_30S               = 32530
	PCS_WGS72BE_UTM_zone_31S               = 32531
	PCS_WGS72BE_UTM_zone_32S               = 32532
	PCS_WGS72BE_UTM_zone_33S               = 32533
	PCS_WGS72BE_UTM_zone_34S               = 32534
	PCS_WGS72BE_UTM_zone_35S               = 32535
	PCS_WGS72BE_UTM_zone_36S               = 32536
	PCS_WGS72BE_UTM_zone_37S               = 32537
	PCS_WGS72BE_UTM_zone_38S               = 32538
	PCS_WGS72BE_UTM_zone_39S               = 32539
	PCS_WGS72BE_UTM_zone_40S               = 32540
	PCS_WGS72BE_UTM_zone_41S               = 32541
	PCS_WGS72BE_UTM_zone_42S               = 32542
	PCS_WGS72BE_UTM_zone_43S               = 32543
	PCS_WGS72BE_UTM_zone_44S               = 32544
	PCS_WGS72BE_UTM_zone_45S               = 32545
	PCS_WGS72BE_UTM_zone_46S               = 32546
	PCS_WGS72BE_UTM_zone_47S               = 32547
	PCS_WGS72BE_UTM_zone_48S               = 32548
	PCS_WGS72BE_UTM_zone_49S               = 32549
	PCS_WGS72BE_UTM_zone_50S               = 32550
	PCS_WGS72BE_UTM_zone_51S               = 32551
	PCS_WGS72BE_UTM_zone_52S               = 32552
	PCS_WGS72BE_UTM_zone_53S               = 32553
	PCS_WGS72BE_UTM_zone_54S               = 32554
	PCS_WGS72BE_UTM_zone_55S               = 32555
	PCS_WGS72BE_UTM_zone_56S               = 32556
	PCS_WGS72BE_UTM_zone_57S               = 32557
	PCS_WGS72BE_UTM_zone_58S               = 32558
	PCS_WGS72BE_UTM_zone_59S               = 32559
	PCS_WGS72BE_UTM_zone_60S               = 32560
	PCS_WGS84_UTM_zone_1N                  = 32601
	PCS_WGS84_UTM_zone_2N                  = 32602
	PCS_WGS84_UTM_zone_3N                  = 32603
	PCS_WGS84_UTM_zone_4N                  = 32604
	PCS_WGS84_UTM_zone_5N                  = 32605
	PCS_WGS84_UTM_zone_6N                  = 32606
	PCS_WGS84_UTM_zone_7N                  = 32607
	PCS_WGS84_UTM_zone_8N                  = 32608
	PCS_WGS84_UTM_zone_9N                  = 32609
	PCS_WGS84_UTM_zone_10N                 = 32610
	PCS_WGS84_UTM_zone_11N                 = 32611
	PCS_WGS84_UTM_zone_12N                 = 32612
	PCS_WGS84_UTM_zone_13N                 = 32613
	PCS_WGS84_UTM_zone_14N                 = 32614
	PCS_WGS84_UTM_zone_15N                 = 32615
	PCS_WGS84_UTM_zone_16N                 = 32616
	PCS_WGS84_UTM_zone_17N                 = 32617
	PCS_WGS84_UTM_zone_18N                 = 32618
	PCS_WGS84_UTM_zone_19N                 = 32619
	PCS_WGS84_UTM_zone_20N                 = 32620
	PCS_WGS84_UTM_zone_21N                 = 32621
	PCS_WGS84_UTM_zone_22N                 = 32622
	PCS_WGS84_UTM_zone_23N                 = 32623
	PCS_WGS84_UTM_zone_24N                 = 32624
	PCS_WGS84_UTM_zone_25N                 = 32625
	PCS_WGS84_UTM_zone_26N                 = 32626
	PCS_WGS84_UTM_zone_27N                 = 32627
	PCS_WGS84_UTM_zone_28N                 = 32628
	PCS_WGS84_UTM_zone_29N                 = 32629
	PCS_WGS84_UTM_zone_30N                 = 32630
	PCS_WGS84_UTM_zone_31N                 = 32631
	PCS_WGS84_UTM_zone_32N                 = 32632
	PCS_WGS84_UTM_zone_33N                 = 32633
	PCS_WGS84_UTM_zone_34N                 = 32634
	PCS_WGS84_UTM_zone_35N                 = 32635
	PCS_WGS84_UTM_zone_36N                 = 32636
	PCS_WGS84_UTM_zone_37N                 = 32637
	PCS_WGS84_UTM_zone_38N                 = 32638
	PCS_WGS84_UTM_zone_39N                 = 32639
	PCS_WGS84_UTM_zone_40N                 = 32640
	PCS_WGS84_UTM_zone_41N                 = 32641
	PCS_WGS84_UTM_zone_42N                 = 32642
	PCS_WGS84_UTM_zone_43N                 = 32643
	PCS_WGS84_UTM_zone_44N                 = 32644
	PCS_WGS84_UTM_zone_45N                 = 32645
	PCS_WGS84_UTM_zone_46N                 = 32646
	PCS_WGS84_UTM_zone_47N                 = 32647
	PCS_WGS84_UTM_zone_48N                 = 32648
	PCS_WGS84_UTM_zone_49N                 = 32649
	PCS_WGS84_UTM_zone_50N                 = 32650
	PCS_WGS84_UTM_zone_51N                 = 32651
	PCS_WGS84_UTM_zone_52N                 = 32652
	PCS_WGS84_UTM_zone_53N                 = 32653
	PCS_WGS84_UTM_zone_54N                 = 32654
	PCS_WGS84_UTM_zone_55N                 = 32655
	PCS_WGS84_UTM_zone_56N                 = 32656
	PCS_WGS84_UTM_zone_57N                 = 32657
	PCS_WGS84_UTM_zone_58N                 = 32658
	PCS_WGS84_UTM_zone_59N                 = 32659
	PCS_WGS84_UTM_zone_60N                 = 32660
	PCS_WGS84_UTM_zone_1S                  = 32701
	PCS_WGS84_UTM_zone_2S                  = 32702
	PCS_WGS84_UTM_zone_3S                  = 32703
	PCS_WGS84_UTM_zone_4S                  = 32704
	PCS_WGS84_UTM_zone_5S                  = 32705
	PCS_WGS84_UTM_zone_6S                  = 32706
	PCS_WGS84_UTM_zone_7S                  = 32707
	PCS_WGS84_UTM_zone_8S                  = 32708
	PCS_WGS84_UTM_zone_9S                  = 32709
	PCS_WGS84_UTM_zone_10S                 = 32710
	PCS_WGS84_UTM_zone_11S                 = 32711
	PCS_WGS84_UTM_zone_12S                 = 32712
	PCS_WGS84_UTM_zone_13S                 = 32713
	PCS_WGS84_UTM_zone_14S                 = 32714
	PCS_WGS84_UTM_zone_15S                 = 32715
	PCS_WGS84_UTM_zone_16S                 = 32716
	PCS_WGS84_UTM_zone_17S                 = 32717
	PCS_WGS84_UTM_zone_18S                 = 32718
	PCS_WGS84_UTM_zone_19S                 = 32719
	PCS_WGS84_UTM_zone_20S                 = 32720
	PCS_WGS84_UTM_zone_21S                 = 32721
	PCS_WGS84_UTM_zone_22S                 = 32722
	PCS_WGS84_UTM_zone_23S                 = 32723
	PCS_WGS84_UTM_zone_24S                 = 32724
	PCS_WGS84_UTM_zone_25S                 = 32725
	PCS_WGS84_UTM_zone_26S                 = 32726
	PCS_WGS84_UTM_zone_27S                 = 32727
	PCS_WGS84_UTM_zone_28S                 = 32728
	PCS_WGS84_UTM_zone_29S                 = 32729
	PCS_WGS84_UTM_zone_30S                 = 32730
	PCS_WGS84_UTM_zone_31S                 = 32731
	PCS_WGS84_UTM_zone_32S                 = 32732
	PCS_WGS84_UTM_zone_33S                 = 32733
	PCS_WGS84_UTM_zone_34S                 = 32734
	PCS_WGS84_UTM_zone_35S                 = 32735
	PCS_WGS84_UTM_zone_36S                 = 32736
	PCS_WGS84_UTM_zone_37S                 = 32737
	PCS_WGS84_UTM_zone_38S                 = 32738
	PCS_WGS84_UTM_zone_39S                 = 32739
	PCS_WGS84_UTM_zone_40S                 = 32740
	PCS_WGS84_UTM_zone_41S                 = 32741
	PCS_WGS84_UTM_zone_42S                 = 32742
	PCS_WGS84_UTM_zone_43S                 = 32743
	PCS_WGS84_UTM_zone_44S                 = 32744
	PCS_WGS84_UTM_zone_45S                 = 32745
	PCS_WGS84_UTM_zone_46S                 = 32746
	PCS_WGS84_UTM_zone_47S                 = 32747
	PCS_WGS84_UTM_zone_48S                 = 32748
	PCS_WGS84_UTM_zone_49S                 = 32749
	PCS_WGS84_UTM_zone_50S                 = 32750
	PCS_WGS84_UTM_zone_51S                 = 32751
	PCS_WGS84_UTM_zone_52S                 = 32752
	PCS_WGS84_UTM_zone_53S                 = 32753
	PCS_WGS84_UTM_zone_54S                 = 32754
	PCS_WGS84_UTM_zone_55S                 = 32755
	PCS_WGS84_UTM_zone_56S                 = 32756
	PCS_WGS84_UTM_zone_57S                 = 32757
	PCS_WGS84_UTM_zone_58S                 = 32758
	PCS_WGS84_UTM_zone_59S                 = 32759
	PCS_WGS84_UTM_zone_60S                 = 32760
)

func PcsName(v int) string {
	switch v {
	case PCS_Hjorsey_1955_Lambert:
		return "PCS_Hjorsey_1955_Lambert"
	case PCS_ISN93_Lambert_1993:
		return "PCS_ISN93_Lambert_1993"
	case PCS_ETRS89_Poland_CS2000_zone_5:
		return "PCS_ETRS89_Poland_CS2000_zone_5"
	case PCS_ETRS89_Poland_CS2000_zone_8:
		return "PCS_ETRS89_Poland_CS2000_zone_8"
	case PCS_ETRS89_Poland_CS92:
		return "PCS_ETRS89_Poland_CS92"
	case PCS_GGRS87_Greek_Grid:
		return "PCS_GGRS87_Greek_Grid"
	case PCS_KKJ_Finland_zone_1:
		return "PCS_KKJ_Finland_zone_1"
	case PCS_KKJ_Finland_zone_2:
		return "PCS_KKJ_Finland_zone_2"
	case PCS_KKJ_Finland_zone_3:
		return "PCS_KKJ_Finland_zone_3"
	case PCS_KKJ_Finland_zone_4:
		return "PCS_KKJ_Finland_zone_4"
	case PCS_RT90_2_5_gon_W:
		return "PCS_RT90_2_5_gon_W"
	case PCS_Lietuvos_Koordinoei_Sistema_1994:
		return "PCS_Lietuvos_Koordinoei_Sistema_1994"
	case PCS_Estonian_Coordinate_System_of_1992:
		return "PCS_Estonian_Coordinate_System_of_1992"
	case PCS_HD72_EOV:
		return "PCS_HD72_EOV"
	case PCS_Dealul_Piscului_1970_Stereo_70:
		return "PCS_Dealul_Piscului_1970_Stereo_70"
	case PCS_Adindan_UTM_zone_37N:
		return "PCS_Adindan_UTM_zone_37N"
	case PCS_Adindan_UTM_zone_38N:
		return "PCS_Adindan_UTM_zone_38N"
	case PCS_AGD66_AMG_zone_48:
		return "PCS_AGD66_AMG_zone_48"
	case PCS_AGD66_AMG_zone_49:
		return "PCS_AGD66_AMG_zone_49"
	case PCS_AGD66_AMG_zone_50:
		return "PCS_AGD66_AMG_zone_50"
	case PCS_AGD66_AMG_zone_51:
		return "PCS_AGD66_AMG_zone_51"
	case PCS_AGD66_AMG_zone_52:
		return "PCS_AGD66_AMG_zone_52"
	case PCS_AGD66_AMG_zone_53:
		return "PCS_AGD66_AMG_zone_53"
	case PCS_AGD66_AMG_zone_54:
		return "PCS_AGD66_AMG_zone_54"
	case PCS_AGD66_AMG_zone_55:
		return "PCS_AGD66_AMG_zone_55"
	case PCS_AGD66_AMG_zone_56:
		return "PCS_AGD66_AMG_zone_56"
	case PCS_AGD66_AMG_zone_57:
		return "PCS_AGD66_AMG_zone_57"
	case PCS_AGD66_AMG_zone_58:
		return "PCS_AGD66_AMG_zone_58"
	case PCS_AGD84_AMG_zone_48:
		return "PCS_AGD84_AMG_zone_48"
	case PCS_AGD84_AMG_zone_49:
		return "PCS_AGD84_AMG_zone_49"
	case PCS_AGD84_AMG_zone_50:
		return "PCS_AGD84_AMG_zone_50"
	case PCS_AGD84_AMG_zone_51:
		return "PCS_AGD84_AMG_zone_51"
	case PCS_AGD84_AMG_zone_52:
		return "PCS_AGD84_AMG_zone_52"
	case PCS_AGD84_AMG_zone_53:
		return "PCS_AGD84_AMG_zone_53"
	case PCS_AGD84_AMG_zone_54:
		return "PCS_AGD84_AMG_zone_54"
	case PCS_AGD84_AMG_zone_55:
		return "PCS_AGD84_AMG_zone_55"
	case PCS_AGD84_AMG_zone_56:
		return "PCS_AGD84_AMG_zone_56"
	case PCS_AGD84_AMG_zone_57:
		return "PCS_AGD84_AMG_zone_57"
	case PCS_AGD84_AMG_zone_58:
		return "PCS_AGD84_AMG_zone_58"
	case PCS_Ain_el_Abd_UTM_zone_37N:
		return "PCS_Ain_el_Abd_UTM_zone_37N"
	case PCS_Ain_el_Abd_UTM_zone_38N:
		return "PCS_Ain_el_Abd_UTM_zone_38N"
	case PCS_Ain_el_Abd_UTM_zone_39N:
		return "PCS_Ain_el_Abd_UTM_zone_39N"
	case PCS_Ain_el_Abd_Bahrain_Grid:
		return "PCS_Ain_el_Abd_Bahrain_Grid"
	case PCS_Afgooye_UTM_zone_38N:
		return "PCS_Afgooye_UTM_zone_38N"
	case PCS_Afgooye_UTM_zone_39N:
		return "PCS_Afgooye_UTM_zone_39N"
	case PCS_Lisbon_Portugese_Grid:
		return "PCS_Lisbon_Portugese_Grid"
	case PCS_Aratu_UTM_zone_22S:
		return "PCS_Aratu_UTM_zone_22S"
	case PCS_Aratu_UTM_zone_23S:
		return "PCS_Aratu_UTM_zone_23S"
	case PCS_Aratu_UTM_zone_24S:
		return "PCS_Aratu_UTM_zone_24S"
	case PCS_Arc_1950_Lo13:
		return "PCS_Arc_1950_Lo13"
	case PCS_Arc_1950_Lo15:
		return "PCS_Arc_1950_Lo15"
	case PCS_Arc_1950_Lo17:
		return "PCS_Arc_1950_Lo17"
	case PCS_Arc_1950_Lo19:
		return "PCS_Arc_1950_Lo19"
	case PCS_Arc_1950_Lo21:
		return "PCS_Arc_1950_Lo21"
	case PCS_Arc_1950_Lo23:
		return "PCS_Arc_1950_Lo23"
	case PCS_Arc_1950_Lo25:
		return "PCS_Arc_1950_Lo25"
	case PCS_Arc_1950_Lo27:
		return "PCS_Arc_1950_Lo27"
	case PCS_Arc_1950_Lo29:
		return "PCS_Arc_1950_Lo29"
	case PCS_Arc_1950_Lo31:
		return "PCS_Arc_1950_Lo31"
	case PCS_Arc_1950_Lo33:
		return "PCS_Arc_1950_Lo33"
	case PCS_Arc_1950_Lo35:
		return "PCS_Arc_1950_Lo35"
	case PCS_Batavia_NEIEZ:
		return "PCS_Batavia_NEIEZ"
	case PCS_Batavia_UTM_zone_48S:
		return "PCS_Batavia_UTM_zone_48S"
	case PCS_Batavia_UTM_zone_49S:
		return "PCS_Batavia_UTM_zone_49S"
	case PCS_Batavia_UTM_zone_50S:
		return "PCS_Batavia_UTM_zone_50S"
	case PCS_Beijing_Gauss_zone_13:
		return "PCS_Beijing_Gauss_zone_13"
	case PCS_Beijing_Gauss_zone_14:
		return "PCS_Beijing_Gauss_zone_14"
	case PCS_Beijing_Gauss_zone_15:
		return "PCS_Beijing_Gauss_zone_15"
	case PCS_Beijing_Gauss_zone_16:
		return "PCS_Beijing_Gauss_zone_16"
	case PCS_Beijing_Gauss_zone_17:
		return "PCS_Beijing_Gauss_zone_17"
	case PCS_Beijing_Gauss_zone_18:
		return "PCS_Beijing_Gauss_zone_18"
	case PCS_Beijing_Gauss_zone_19:
		return "PCS_Beijing_Gauss_zone_19"
	case PCS_Beijing_Gauss_zone_20:
		return "PCS_Beijing_Gauss_zone_20"
	case PCS_Beijing_Gauss_zone_21:
		return "PCS_Beijing_Gauss_zone_21"
	case PCS_Beijing_Gauss_zone_22:
		return "PCS_Beijing_Gauss_zone_22"
	case PCS_Beijing_Gauss_zone_23:
		return "PCS_Beijing_Gauss_zone_23"
	case PCS_Beijing_Gauss_13N:
		return "PCS_Beijing_Gauss_13N"
	case PCS_Beijing_Gauss_14N:
		return "PCS_Beijing_Gauss_14N"
	case PCS_Beijing_Gauss_15N:
		return "PCS_Beijing_Gauss_15N"
	case PCS_Beijing_Gauss_16N:
		return "PCS_Beijing_Gauss_16N"
	case PCS_Beijing_Gauss_17N:
		return "PCS_Beijing_Gauss_17N"
	case PCS_Beijing_Gauss_18N:
		return "PCS_Beijing_Gauss_18N"
	case PCS_Beijing_Gauss_19N:
		return "PCS_Beijing_Gauss_19N"
	case PCS_Beijing_Gauss_20N:
		return "PCS_Beijing_Gauss_20N"
	case PCS_Beijing_Gauss_21N:
		return "PCS_Beijing_Gauss_21N"
	case PCS_Beijing_Gauss_22N:
		return "PCS_Beijing_Gauss_22N"
	case PCS_Beijing_Gauss_23N:
		return "PCS_Beijing_Gauss_23N"
	case PCS_Belge_Lambert_50:
		return "PCS_Belge_Lambert_50"
	case PCS_Bern_1898_Swiss_Old:
		return "PCS_Bern_1898_Swiss_Old"
	case PCS_Bogota_UTM_zone_17N:
		return "PCS_Bogota_UTM_zone_17N"
	case PCS_Bogota_UTM_zone_18N:
		return "PCS_Bogota_UTM_zone_18N"
	case PCS_Bogota_Colombia_3W:
		return "PCS_Bogota_Colombia_3W"
	case PCS_Bogota_Colombia_Bogota:
		return "PCS_Bogota_Colombia_Bogota"
	case PCS_Bogota_Colombia_3E:
		return "PCS_Bogota_Colombia_3E"
	case PCS_Bogota_Colombia_6E:
		return "PCS_Bogota_Colombia_6E"
	case PCS_Camacupa_UTM_32S:
		return "PCS_Camacupa_UTM_32S"
	case PCS_Camacupa_UTM_33S:
		return "PCS_Camacupa_UTM_33S"
	case PCS_C_Inchauspe_Argentina_1:
		return "PCS_C_Inchauspe_Argentina_1"
	case PCS_C_Inchauspe_Argentina_2:
		return "PCS_C_Inchauspe_Argentina_2"
	case PCS_C_Inchauspe_Argentina_3:
		return "PCS_C_Inchauspe_Argentina_3"
	case PCS_C_Inchauspe_Argentina_4:
		return "PCS_C_Inchauspe_Argentina_4"
	case PCS_C_Inchauspe_Argentina_5:
		return "PCS_C_Inchauspe_Argentina_5"
	case PCS_C_Inchauspe_Argentina_6:
		return "PCS_C_Inchauspe_Argentina_6"
	case PCS_C_Inchauspe_Argentina_7:
		return "PCS_C_Inchauspe_Argentina_7"
	case PCS_Carthage_UTM_zone_32N:
		return "PCS_Carthage_UTM_zone_32N"
	case PCS_Carthage_Nord_Tunisie:
		return "PCS_Carthage_Nord_Tunisie"
	case PCS_Carthage_Sud_Tunisie:
		return "PCS_Carthage_Sud_Tunisie"
	case PCS_Corrego_Alegre_UTM_23S:
		return "PCS_Corrego_Alegre_UTM_23S"
	case PCS_Corrego_Alegre_UTM_24S:
		return "PCS_Corrego_Alegre_UTM_24S"
	case PCS_Douala_UTM_zone_32N:
		return "PCS_Douala_UTM_zone_32N"
	case PCS_Egypt_1907_Red_Belt:
		return "PCS_Egypt_1907_Red_Belt"
	case PCS_Egypt_1907_Purple_Belt:
		return "PCS_Egypt_1907_Purple_Belt"
	case PCS_Egypt_1907_Ext_Purple:
		return "PCS_Egypt_1907_Ext_Purple"
	case PCS_ED50_UTM_zone_28N:
		return "PCS_ED50_UTM_zone_28N"
	case PCS_ED50_UTM_zone_29N:
		return "PCS_ED50_UTM_zone_29N"
	case PCS_ED50_UTM_zone_30N:
		return "PCS_ED50_UTM_zone_30N"
	case PCS_ED50_UTM_zone_31N:
		return "PCS_ED50_UTM_zone_31N"
	case PCS_ED50_UTM_zone_32N:
		return "PCS_ED50_UTM_zone_32N"
	case PCS_ED50_UTM_zone_33N:
		return "PCS_ED50_UTM_zone_33N"
	case PCS_ED50_UTM_zone_34N:
		return "PCS_ED50_UTM_zone_34N"
	case PCS_ED50_UTM_zone_35N:
		return "PCS_ED50_UTM_zone_35N"
	case PCS_ED50_UTM_zone_36N:
		return "PCS_ED50_UTM_zone_36N"
	case PCS_ED50_UTM_zone_37N:
		return "PCS_ED50_UTM_zone_37N"
	case PCS_ED50_UTM_zone_38N:
		return "PCS_ED50_UTM_zone_38N"
	case PCS_Fahud_UTM_zone_39N:
		return "PCS_Fahud_UTM_zone_39N"
	case PCS_Fahud_UTM_zone_40N:
		return "PCS_Fahud_UTM_zone_40N"
	case PCS_Garoua_UTM_zone_33N:
		return "PCS_Garoua_UTM_zone_33N"
	case PCS_ID74_UTM_zone_46N:
		return "PCS_ID74_UTM_zone_46N"
	case PCS_ID74_UTM_zone_47N:
		return "PCS_ID74_UTM_zone_47N"
	case PCS_ID74_UTM_zone_48N:
		return "PCS_ID74_UTM_zone_48N"
	case PCS_ID74_UTM_zone_49N:
		return "PCS_ID74_UTM_zone_49N"
	case PCS_ID74_UTM_zone_50N:
		return "PCS_ID74_UTM_zone_50N"
	case PCS_ID74_UTM_zone_51N:
		return "PCS_ID74_UTM_zone_51N"
	case PCS_ID74_UTM_zone_52N:
		return "PCS_ID74_UTM_zone_52N"
	case PCS_ID74_UTM_zone_53N:
		return "PCS_ID74_UTM_zone_53N"
	case PCS_ID74_UTM_zone_46S:
		return "PCS_ID74_UTM_zone_46S"
	case PCS_ID74_UTM_zone_47S:
		return "PCS_ID74_UTM_zone_47S"
	case PCS_ID74_UTM_zone_48S:
		return "PCS_ID74_UTM_zone_48S"
	case PCS_ID74_UTM_zone_49S:
		return "PCS_ID74_UTM_zone_49S"
	case PCS_ID74_UTM_zone_50S:
		return "PCS_ID74_UTM_zone_50S"
	case PCS_ID74_UTM_zone_51S:
		return "PCS_ID74_UTM_zone_51S"
	case PCS_ID74_UTM_zone_52S:
		return "PCS_ID74_UTM_zone_52S"
	case PCS_ID74_UTM_zone_53S:
		return "PCS_ID74_UTM_zone_53S"
	case PCS_ID74_UTM_zone_54S:
		return "PCS_ID74_UTM_zone_54S"
	case PCS_Indian_1954_UTM_47N:
		return "PCS_Indian_1954_UTM_47N"
	case PCS_Indian_1954_UTM_48N:
		return "PCS_Indian_1954_UTM_48N"
	case PCS_Indian_1975_UTM_47N:
		return "PCS_Indian_1975_UTM_47N"
	case PCS_Indian_1975_UTM_48N:
		return "PCS_Indian_1975_UTM_48N"
	case PCS_Jamaica_1875_Old_Grid:
		return "PCS_Jamaica_1875_Old_Grid"
	case PCS_JAD69_Jamaica_Grid:
		return "PCS_JAD69_Jamaica_Grid"
	case PCS_Kalianpur_India_0:
		return "PCS_Kalianpur_India_0"
	case PCS_Kalianpur_India_I:
		return "PCS_Kalianpur_India_I"
	case PCS_Kalianpur_India_IIa:
		return "PCS_Kalianpur_India_IIa"
	case PCS_Kalianpur_India_IIIa:
		return "PCS_Kalianpur_India_IIIa"
	case PCS_Kalianpur_India_IVa:
		return "PCS_Kalianpur_India_IVa"
	case PCS_Kalianpur_India_IIb:
		return "PCS_Kalianpur_India_IIb"
	case PCS_Kalianpur_India_IIIb:
		return "PCS_Kalianpur_India_IIIb"
	case PCS_Kalianpur_India_IVb:
		return "PCS_Kalianpur_India_IVb"
	case PCS_Kertau_Singapore_Grid:
		return "PCS_Kertau_Singapore_Grid"
	case PCS_Kertau_UTM_zone_47N:
		return "PCS_Kertau_UTM_zone_47N"
	case PCS_Kertau_UTM_zone_48N:
		return "PCS_Kertau_UTM_zone_48N"
	case PCS_La_Canoa_UTM_zone_20N:
		return "PCS_La_Canoa_UTM_zone_20N"
	case PCS_La_Canoa_UTM_zone_21N:
		return "PCS_La_Canoa_UTM_zone_21N"
	case PCS_PSAD56_UTM_zone_18N:
		return "PCS_PSAD56_UTM_zone_18N"
	case PCS_PSAD56_UTM_zone_19N:
		return "PCS_PSAD56_UTM_zone_19N"
	case PCS_PSAD56_UTM_zone_20N:
		return "PCS_PSAD56_UTM_zone_20N"
	case PCS_PSAD56_UTM_zone_21N:
		return "PCS_PSAD56_UTM_zone_21N"
	case PCS_PSAD56_UTM_zone_17S:
		return "PCS_PSAD56_UTM_zone_17S"
	case PCS_PSAD56_UTM_zone_18S:
		return "PCS_PSAD56_UTM_zone_18S"
	case PCS_PSAD56_UTM_zone_19S:
		return "PCS_PSAD56_UTM_zone_19S"
	case PCS_PSAD56_UTM_zone_20S:
		return "PCS_PSAD56_UTM_zone_20S"
	case PCS_PSAD56_Peru_west_zone:
		return "PCS_PSAD56_Peru_west_zone"
	case PCS_PSAD56_Peru_central:
		return "PCS_PSAD56_Peru_central"
	case PCS_PSAD56_Peru_east_zone:
		return "PCS_PSAD56_Peru_east_zone"
	case PCS_Leigon_Ghana_Grid:
		return "PCS_Leigon_Ghana_Grid"
	case PCS_Lome_UTM_zone_31N:
		return "PCS_Lome_UTM_zone_31N"
	case PCS_Luzon_Philippines_I:
		return "PCS_Luzon_Philippines_I"
	case PCS_Luzon_Philippines_II:
		return "PCS_Luzon_Philippines_II"
	case PCS_Luzon_Philippines_III:
		return "PCS_Luzon_Philippines_III"
	case PCS_Luzon_Philippines_IV:
		return "PCS_Luzon_Philippines_IV"
	case PCS_Luzon_Philippines_V:
		return "PCS_Luzon_Philippines_V"
	case PCS_Makassar_NEIEZ:
		return "PCS_Makassar_NEIEZ"
	case PCS_Malongo_1987_UTM_32S:
		return "PCS_Malongo_1987_UTM_32S"
	case PCS_Merchich_Nord_Maroc:
		return "PCS_Merchich_Nord_Maroc"
	case PCS_Merchich_Sud_Maroc:
		return "PCS_Merchich_Sud_Maroc"
	case PCS_Merchich_Sahara:
		return "PCS_Merchich_Sahara"
	case PCS_Massawa_UTM_zone_37N:
		return "PCS_Massawa_UTM_zone_37N"
	case PCS_Minna_UTM_zone_31N:
		return "PCS_Minna_UTM_zone_31N"
	case PCS_Minna_UTM_zone_32N:
		return "PCS_Minna_UTM_zone_32N"
	case PCS_Minna_Nigeria_West:
		return "PCS_Minna_Nigeria_West"
	case PCS_Minna_Nigeria_Mid_Belt:
		return "PCS_Minna_Nigeria_Mid_Belt"
	case PCS_Minna_Nigeria_East:
		return "PCS_Minna_Nigeria_East"
	case PCS_Mhast_UTM_zone_32S:
		return "PCS_Mhast_UTM_zone_32S"
	case PCS_Monte_Mario_Italy_1:
		return "PCS_Monte_Mario_Italy_1"
	case PCS_Monte_Mario_Italy_2:
		return "PCS_Monte_Mario_Italy_2"
	case PCS_M_poraloko_UTM_32N:
		return "PCS_M_poraloko_UTM_32N"
	case PCS_M_poraloko_UTM_32S:
		return "PCS_M_poraloko_UTM_32S"
	case PCS_NAD27_UTM_zone_3N:
		return "PCS_NAD27_UTM_zone_3N"
	case PCS_NAD27_UTM_zone_4N:
		return "PCS_NAD27_UTM_zone_4N"
	case PCS_NAD27_UTM_zone_5N:
		return "PCS_NAD27_UTM_zone_5N"
	case PCS_NAD27_UTM_zone_6N:
		return "PCS_NAD27_UTM_zone_6N"
	case PCS_NAD27_UTM_zone_7N:
		return "PCS_NAD27_UTM_zone_7N"
	case PCS_NAD27_UTM_zone_8N:
		return "PCS_NAD27_UTM_zone_8N"
	case PCS_NAD27_UTM_zone_9N:
		return "PCS_NAD27_UTM_zone_9N"
	case PCS_NAD27_UTM_zone_10N:
		return "PCS_NAD27_UTM_zone_10N"
	case PCS_NAD27_UTM_zone_11N:
		return "PCS_NAD27_UTM_zone_11N"
	case PCS_NAD27_UTM_zone_12N:
		return "PCS_NAD27_UTM_zone_12N"
	case PCS_NAD27_UTM_zone_13N:
		return "PCS_NAD27_UTM_zone_13N"
	case PCS_NAD27_UTM_zone_14N:
		return "PCS_NAD27_UTM_zone_14N"
	case PCS_NAD27_UTM_zone_15N:
		return "PCS_NAD27_UTM_zone_15N"
	case PCS_NAD27_UTM_zone_16N:
		return "PCS_NAD27_UTM_zone_16N"
	case PCS_NAD27_UTM_zone_17N:
		return "PCS_NAD27_UTM_zone_17N"
	case PCS_NAD27_UTM_zone_18N:
		return "PCS_NAD27_UTM_zone_18N"
	case PCS_NAD27_UTM_zone_19N:
		return "PCS_NAD27_UTM_zone_19N"
	case PCS_NAD27_UTM_zone_20N:
		return "PCS_NAD27_UTM_zone_20N"
	case PCS_NAD27_UTM_zone_21N:
		return "PCS_NAD27_UTM_zone_21N"
	case PCS_NAD27_UTM_zone_22N:
		return "PCS_NAD27_UTM_zone_22N"
	case PCS_NAD27_Alabama_East:
		return "PCS_NAD27_Alabama_East"
	case PCS_NAD27_Alabama_West:
		return "PCS_NAD27_Alabama_West"
	case PCS_NAD27_Alaska_zone_1:
		return "PCS_NAD27_Alaska_zone_1"
	case PCS_NAD27_Alaska_zone_2:
		return "PCS_NAD27_Alaska_zone_2"
	case PCS_NAD27_Alaska_zone_3:
		return "PCS_NAD27_Alaska_zone_3"
	case PCS_NAD27_Alaska_zone_4:
		return "PCS_NAD27_Alaska_zone_4"
	case PCS_NAD27_Alaska_zone_5:
		return "PCS_NAD27_Alaska_zone_5"
	case PCS_NAD27_Alaska_zone_6:
		return "PCS_NAD27_Alaska_zone_6"
	case PCS_NAD27_Alaska_zone_7:
		return "PCS_NAD27_Alaska_zone_7"
	case PCS_NAD27_Alaska_zone_8:
		return "PCS_NAD27_Alaska_zone_8"
	case PCS_NAD27_Alaska_zone_9:
		return "PCS_NAD27_Alaska_zone_9"
	case PCS_NAD27_Alaska_zone_10:
		return "PCS_NAD27_Alaska_zone_10"
	case PCS_NAD27_California_I:
		return "PCS_NAD27_California_I"
	case PCS_NAD27_California_II:
		return "PCS_NAD27_California_II"
	case PCS_NAD27_California_III:
		return "PCS_NAD27_California_III"
	case PCS_NAD27_California_IV:
		return "PCS_NAD27_California_IV"
	case PCS_NAD27_California_V:
		return "PCS_NAD27_California_V"
	case PCS_NAD27_California_VI:
		return "PCS_NAD27_California_VI"
	case PCS_NAD27_California_VII:
		return "PCS_NAD27_California_VII"
	case PCS_NAD27_Arizona_East:
		return "PCS_NAD27_Arizona_East"
	case PCS_NAD27_Arizona_Central:
		return "PCS_NAD27_Arizona_Central"
	case PCS_NAD27_Arizona_West:
		return "PCS_NAD27_Arizona_West"
	case PCS_NAD27_Arkansas_North:
		return "PCS_NAD27_Arkansas_North"
	case PCS_NAD27_Arkansas_South:
		return "PCS_NAD27_Arkansas_South"
	case PCS_NAD27_Colorado_North:
		return "PCS_NAD27_Colorado_North"
	case PCS_NAD27_Colorado_Central:
		return "PCS_NAD27_Colorado_Central"
	case PCS_NAD27_Colorado_South:
		return "PCS_NAD27_Colorado_South"
	case PCS_NAD27_Connecticut:
		return "PCS_NAD27_Connecticut"
	case PCS_NAD27_Delaware:
		return "PCS_NAD27_Delaware"
	case PCS_NAD27_Florida_East:
		return "PCS_NAD27_Florida_East"
	case PCS_NAD27_Florida_West:
		return "PCS_NAD27_Florida_West"
	case PCS_NAD27_Florida_North:
		return "PCS_NAD27_Florida_North"
	case PCS_NAD27_Hawaii_zone_1:
		return "PCS_NAD27_Hawaii_zone_1"
	case PCS_NAD27_Hawaii_zone_2:
		return "PCS_NAD27_Hawaii_zone_2"
	case PCS_NAD27_Hawaii_zone_3:
		return "PCS_NAD27_Hawaii_zone_3"
	case PCS_NAD27_Hawaii_zone_4:
		return "PCS_NAD27_Hawaii_zone_4"
	case PCS_NAD27_Hawaii_zone_5:
		return "PCS_NAD27_Hawaii_zone_5"
	case PCS_NAD27_Georgia_East:
		return "PCS_NAD27_Georgia_East"
	case PCS_NAD27_Georgia_West:
		return "PCS_NAD27_Georgia_West"
	case PCS_NAD27_Idaho_East:
		return "PCS_NAD27_Idaho_East"
	case PCS_NAD27_Idaho_Central:
		return "PCS_NAD27_Idaho_Central"
	case PCS_NAD27_Idaho_West:
		return "PCS_NAD27_Idaho_West"
	case PCS_NAD27_Illinois_East:
		return "PCS_NAD27_Illinois_East"
	case PCS_NAD27_Illinois_West:
		return "PCS_NAD27_Illinois_West"
	case PCS_NAD27_Indiana_East:
		return "PCS_NAD27_Indiana_East"
	case PCS_NAD27_Kansas_South:
		return "PCS_NAD27_Kansas_South"
	case PCS_NAD27_Kentucky_North:
		return "PCS_NAD27_Kentucky_North"
	case PCS_NAD27_Kentucky_South:
		return "PCS_NAD27_Kentucky_South"
	case PCS_NAD27_Louisiana_North:
		return "PCS_NAD27_Louisiana_North"
	case PCS_NAD27_Louisiana_South:
		return "PCS_NAD27_Louisiana_South"
	case PCS_NAD27_Maine_East:
		return "PCS_NAD27_Maine_East"
	case PCS_NAD27_Maine_West:
		return "PCS_NAD27_Maine_West"
	case PCS_NAD27_Maryland:
		return "PCS_NAD27_Maryland"
	case PCS_NAD27_Massachusetts:
		return "PCS_NAD27_Massachusetts"
	case PCS_NAD27_Massachusetts_Is:
		return "PCS_NAD27_Massachusetts_Is"
	case PCS_NAD27_Michigan_North:
		return "PCS_NAD27_Michigan_North"
	case PCS_NAD27_Michigan_Central:
		return "PCS_NAD27_Michigan_Central"
	case PCS_NAD27_Michigan_South:
		return "PCS_NAD27_Michigan_South"
	case PCS_NAD27_Minnesota_North:
		return "PCS_NAD27_Minnesota_North"
	case PCS_NAD27_Minnesota_Cent:
		return "PCS_NAD27_Minnesota_Cent"
	case PCS_NAD27_Minnesota_South:
		return "PCS_NAD27_Minnesota_South"
	case PCS_NAD27_Mississippi_East:
		return "PCS_NAD27_Mississippi_East"
	case PCS_NAD27_Mississippi_West:
		return "PCS_NAD27_Mississippi_West"
	case PCS_NAD27_Missouri_East:
		return "PCS_NAD27_Missouri_East"
	case PCS_NAD27_Missouri_Central:
		return "PCS_NAD27_Missouri_Central"
	case PCS_NAD27_Missouri_West:
		return "PCS_NAD27_Missouri_West"
	case PCS_NAD_Michigan_Michigan_East:
		return "PCS_NAD_Michigan_Michigan_East"
	case PCS_NAD_Michigan_Michigan_Old_Central:
		return "PCS_NAD_Michigan_Michigan_Old_Central"
	case PCS_NAD_Michigan_Michigan_West:
		return "PCS_NAD_Michigan_Michigan_West"
	case PCS_NAD83_UTM_zone_3N:
		return "PCS_NAD83_UTM_zone_3N"
	case PCS_NAD83_UTM_zone_4N:
		return "PCS_NAD83_UTM_zone_4N"
	case PCS_NAD83_UTM_zone_5N:
		return "PCS_NAD83_UTM_zone_5N"
	case PCS_NAD83_UTM_zone_6N:
		return "PCS_NAD83_UTM_zone_6N"
	case PCS_NAD83_UTM_zone_7N:
		return "PCS_NAD83_UTM_zone_7N"
	case PCS_NAD83_UTM_zone_8N:
		return "PCS_NAD83_UTM_zone_8N"
	case PCS_NAD83_UTM_zone_9N:
		return "PCS_NAD83_UTM_zone_9N"
	case PCS_NAD83_UTM_zone_10N:
		return "PCS_NAD83_UTM_zone_10N"
	case PCS_NAD83_UTM_zone_11N:
		return "PCS_NAD83_UTM_zone_11N"
	case PCS_NAD83_UTM_zone_12N:
		return "PCS_NAD83_UTM_zone_12N"
	case PCS_NAD83_UTM_zone_13N:
		return "PCS_NAD83_UTM_zone_13N"
	case PCS_NAD83_UTM_zone_14N:
		return "PCS_NAD83_UTM_zone_14N"
	case PCS_NAD83_UTM_zone_15N:
		return "PCS_NAD83_UTM_zone_15N"
	case PCS_NAD83_UTM_zone_16N:
		return "PCS_NAD83_UTM_zone_16N"
	case PCS_NAD83_UTM_zone_17N:
		return "PCS_NAD83_UTM_zone_17N"
	case PCS_NAD83_UTM_zone_18N:
		return "PCS_NAD83_UTM_zone_18N"
	case PCS_NAD83_UTM_zone_19N:
		return "PCS_NAD83_UTM_zone_19N"
	case PCS_NAD83_UTM_zone_20N:
		return "PCS_NAD83_UTM_zone_20N"
	case PCS_NAD83_UTM_zone_21N:
		return "PCS_NAD83_UTM_zone_21N"
	case PCS_NAD83_UTM_zone_22N:
		return "PCS_NAD83_UTM_zone_22N"
	case PCS_NAD83_UTM_zone_23N:
		return "PCS_NAD83_UTM_zone_23N"
	case PCS_NAD83_Alabama_East:
		return "PCS_NAD83_Alabama_East"
	case PCS_NAD83_Alabama_West:
		return "PCS_NAD83_Alabama_West"
	case PCS_NAD83_Alaska_zone_1:
		return "PCS_NAD83_Alaska_zone_1"
	case PCS_NAD83_Alaska_zone_2:
		return "PCS_NAD83_Alaska_zone_2"
	case PCS_NAD83_Alaska_zone_3:
		return "PCS_NAD83_Alaska_zone_3"
	case PCS_NAD83_Alaska_zone_4:
		return "PCS_NAD83_Alaska_zone_4"
	case PCS_NAD83_Alaska_zone_5:
		return "PCS_NAD83_Alaska_zone_5"
	case PCS_NAD83_Alaska_zone_6:
		return "PCS_NAD83_Alaska_zone_6"
	case PCS_NAD83_Alaska_zone_7:
		return "PCS_NAD83_Alaska_zone_7"
	case PCS_NAD83_Alaska_zone_8:
		return "PCS_NAD83_Alaska_zone_8"
	case PCS_NAD83_Alaska_zone_9:
		return "PCS_NAD83_Alaska_zone_9"
	case PCS_NAD83_Alaska_zone_10:
		return "PCS_NAD83_Alaska_zone_10"
	case PCS_NAD83_California_1:
		return "PCS_NAD83_California_1"
	case PCS_NAD83_California_2:
		return "PCS_NAD83_California_2"
	case PCS_NAD83_California_3:
		return "PCS_NAD83_California_3"
	case PCS_NAD83_California_4:
		return "PCS_NAD83_California_4"
	case PCS_NAD83_California_5:
		return "PCS_NAD83_California_5"
	case PCS_NAD83_California_6:
		return "PCS_NAD83_California_6"
	case PCS_NAD83_Arizona_East:
		return "PCS_NAD83_Arizona_East"
	case PCS_NAD83_Arizona_Central:
		return "PCS_NAD83_Arizona_Central"
	case PCS_NAD83_Arizona_West:
		return "PCS_NAD83_Arizona_West"
	case PCS_NAD83_Arkansas_North:
		return "PCS_NAD83_Arkansas_North"
	case PCS_NAD83_Arkansas_South:
		return "PCS_NAD83_Arkansas_South"
	case PCS_NAD83_Colorado_North:
		return "PCS_NAD83_Colorado_North"
	case PCS_NAD83_Colorado_Central:
		return "PCS_NAD83_Colorado_Central"
	case PCS_NAD83_Colorado_South:
		return "PCS_NAD83_Colorado_South"
	case PCS_NAD83_Connecticut:
		return "PCS_NAD83_Connecticut"
	case PCS_NAD83_Delaware:
		return "PCS_NAD83_Delaware"
	case PCS_NAD83_Florida_East:
		return "PCS_NAD83_Florida_East"
	case PCS_NAD83_Florida_West:
		return "PCS_NAD83_Florida_West"
	case PCS_NAD83_Florida_North:
		return "PCS_NAD83_Florida_North"
	case PCS_NAD83_Hawaii_zone_1:
		return "PCS_NAD83_Hawaii_zone_1"
	case PCS_NAD83_Hawaii_zone_2:
		return "PCS_NAD83_Hawaii_zone_2"
	case PCS_NAD83_Hawaii_zone_3:
		return "PCS_NAD83_Hawaii_zone_3"
	case PCS_NAD83_Hawaii_zone_4:
		return "PCS_NAD83_Hawaii_zone_4"
	case PCS_NAD83_Hawaii_zone_5:
		return "PCS_NAD83_Hawaii_zone_5"
	case PCS_NAD83_Georgia_East:
		return "PCS_NAD83_Georgia_East"
	case PCS_NAD83_Georgia_West:
		return "PCS_NAD83_Georgia_West"
	case PCS_NAD83_Idaho_East:
		return "PCS_NAD83_Idaho_East"
	case PCS_NAD83_Idaho_Central:
		return "PCS_NAD83_Idaho_Central"
	case PCS_NAD83_Idaho_West:
		return "PCS_NAD83_Idaho_West"
	case PCS_NAD83_Illinois_East:
		return "PCS_NAD83_Illinois_East"
	case PCS_NAD83_Illinois_West:
		return "PCS_NAD83_Illinois_West"
	case PCS_NAD83_Indiana_East:
		return "PCS_NAD83_Indiana_East"
	case PCS_NAD83_Indiana_West:
		return "PCS_NAD83_Indiana_West"
	case PCS_NAD83_Iowa_North:
		return "PCS_NAD83_Iowa_North"
	case PCS_NAD83_Iowa_South:
		return "PCS_NAD83_Iowa_South"
	case PCS_NAD83_Kansas_North:
		return "PCS_NAD83_Kansas_North"
	case PCS_NAD83_Kansas_South:
		return "PCS_NAD83_Kansas_South"
	case PCS_NAD83_Kentucky_North:
		return "PCS_NAD83_Kentucky_North"
	case PCS_NAD83_Kentucky_South:
		return "PCS_NAD83_Kentucky_South"
	case PCS_NAD83_Louisiana_North:
		return "PCS_NAD83_Louisiana_North"
	case PCS_NAD83_Louisiana_South:
		return "PCS_NAD83_Louisiana_South"
	case PCS_NAD83_Maine_East:
		return "PCS_NAD83_Maine_East"
	case PCS_NAD83_Maine_West:
		return "PCS_NAD83_Maine_West"
	case PCS_NAD83_Maryland:
		return "PCS_NAD83_Maryland"
	case PCS_NAD83_Massachusetts:
		return "PCS_NAD83_Massachusetts"
	case PCS_NAD83_Massachusetts_Is:
		return "PCS_NAD83_Massachusetts_Is"
	case PCS_NAD83_Michigan_North:
		return "PCS_NAD83_Michigan_North"
	case PCS_NAD83_Michigan_Central:
		return "PCS_NAD83_Michigan_Central"
	case PCS_NAD83_Michigan_South:
		return "PCS_NAD83_Michigan_South"
	case PCS_NAD83_Minnesota_North:
		return "PCS_NAD83_Minnesota_North"
	case PCS_NAD83_Minnesota_Cent:
		return "PCS_NAD83_Minnesota_Cent"
	case PCS_NAD83_Minnesota_South:
		return "PCS_NAD83_Minnesota_South"
	case PCS_NAD83_Mississippi_East:
		return "PCS_NAD83_Mississippi_East"
	case PCS_NAD83_Mississippi_West:
		return "PCS_NAD83_Mississippi_West"
	case PCS_NAD83_Missouri_East:
		return "PCS_NAD83_Missouri_East"
	case PCS_NAD83_Missouri_Central:
		return "PCS_NAD83_Missouri_Central"
	case PCS_NAD83_Missouri_West:
		return "PCS_NAD83_Missouri_West"
	case PCS_Nahrwan_1967_UTM_38N:
		return "PCS_Nahrwan_1967_UTM_38N"
	case PCS_Nahrwan_1967_UTM_39N:
		return "PCS_Nahrwan_1967_UTM_39N"
	case PCS_Nahrwan_1967_UTM_40N:
		return "PCS_Nahrwan_1967_UTM_40N"
	case PCS_Naparima_UTM_20N:
		return "PCS_Naparima_UTM_20N"
	case PCS_GD49_NZ_Map_Grid:
		return "PCS_GD49_NZ_Map_Grid"
	case PCS_GD49_North_Island_Grid:
		return "PCS_GD49_North_Island_Grid"
	case PCS_GD49_South_Island_Grid:
		return "PCS_GD49_South_Island_Grid"
	case PCS_Datum_73_UTM_zone_29N:
		return "PCS_Datum_73_UTM_zone_29N"
	case PCS_ATF_Nord_de_Guerre:
		return "PCS_ATF_Nord_de_Guerre"
	case PCS_NTF_France_I:
		return "PCS_NTF_France_I"
	case PCS_NTF_France_II:
		return "PCS_NTF_France_II"
	case PCS_NTF_France_III:
		return "PCS_NTF_France_III"
	case PCS_NTF_Nord_France:
		return "PCS_NTF_Nord_France"
	case PCS_NTF_Centre_France:
		return "PCS_NTF_Centre_France"
	case PCS_NTF_Sud_France:
		return "PCS_NTF_Sud_France"
	case PCS_British_National_Grid:
		return "PCS_British_National_Grid"
	case PCS_Point_Noire_UTM_32S:
		return "PCS_Point_Noire_UTM_32S"
	case PCS_GDA94_MGA_zone_48:
		return "PCS_GDA94_MGA_zone_48"
	case PCS_GDA94_MGA_zone_49:
		return "PCS_GDA94_MGA_zone_49"
	case PCS_GDA94_MGA_zone_50:
		return "PCS_GDA94_MGA_zone_50"
	case PCS_GDA94_MGA_zone_51:
		return "PCS_GDA94_MGA_zone_51"
	case PCS_GDA94_MGA_zone_52:
		return "PCS_GDA94_MGA_zone_52"
	case PCS_GDA94_MGA_zone_53:
		return "PCS_GDA94_MGA_zone_53"
	case PCS_GDA94_MGA_zone_54:
		return "PCS_GDA94_MGA_zone_54"
	case PCS_GDA94_MGA_zone_55:
		return "PCS_GDA94_MGA_zone_55"
	case PCS_GDA94_MGA_zone_56:
		return "PCS_GDA94_MGA_zone_56"
	case PCS_GDA94_MGA_zone_57:
		return "PCS_GDA94_MGA_zone_57"
	case PCS_GDA94_MGA_zone_58:
		return "PCS_GDA94_MGA_zone_58"
	case PCS_Pulkovo_Gauss_zone_4:
		return "PCS_Pulkovo_Gauss_zone_4"
	case PCS_Pulkovo_Gauss_zone_5:
		return "PCS_Pulkovo_Gauss_zone_5"
	case PCS_Pulkovo_Gauss_zone_6:
		return "PCS_Pulkovo_Gauss_zone_6"
	case PCS_Pulkovo_Gauss_zone_7:
		return "PCS_Pulkovo_Gauss_zone_7"
	case PCS_Pulkovo_Gauss_zone_8:
		return "PCS_Pulkovo_Gauss_zone_8"
	case PCS_Pulkovo_Gauss_zone_9:
		return "PCS_Pulkovo_Gauss_zone_9"
	case PCS_Pulkovo_Gauss_zone_10:
		return "PCS_Pulkovo_Gauss_zone_10"
	case PCS_Pulkovo_Gauss_zone_11:
		return "PCS_Pulkovo_Gauss_zone_11"
	case PCS_Pulkovo_Gauss_zone_12:
		return "PCS_Pulkovo_Gauss_zone_12"
	case PCS_Pulkovo_Gauss_zone_13:
		return "PCS_Pulkovo_Gauss_zone_13"
	case PCS_Pulkovo_Gauss_zone_14:
		return "PCS_Pulkovo_Gauss_zone_14"
	case PCS_Pulkovo_Gauss_zone_15:
		return "PCS_Pulkovo_Gauss_zone_15"
	case PCS_Pulkovo_Gauss_zone_16:
		return "PCS_Pulkovo_Gauss_zone_16"
	case PCS_Pulkovo_Gauss_zone_17:
		return "PCS_Pulkovo_Gauss_zone_17"
	case PCS_Pulkovo_Gauss_zone_18:
		return "PCS_Pulkovo_Gauss_zone_18"
	case PCS_Pulkovo_Gauss_zone_19:
		return "PCS_Pulkovo_Gauss_zone_19"
	case PCS_Pulkovo_Gauss_zone_20:
		return "PCS_Pulkovo_Gauss_zone_20"
	case PCS_Pulkovo_Gauss_zone_21:
		return "PCS_Pulkovo_Gauss_zone_21"
	case PCS_Pulkovo_Gauss_zone_22:
		return "PCS_Pulkovo_Gauss_zone_22"
	case PCS_Pulkovo_Gauss_zone_23:
		return "PCS_Pulkovo_Gauss_zone_23"
	case PCS_Pulkovo_Gauss_zone_24:
		return "PCS_Pulkovo_Gauss_zone_24"
	case PCS_Pulkovo_Gauss_zone_25:
		return "PCS_Pulkovo_Gauss_zone_25"
	case PCS_Pulkovo_Gauss_zone_26:
		return "PCS_Pulkovo_Gauss_zone_26"
	case PCS_Pulkovo_Gauss_zone_27:
		return "PCS_Pulkovo_Gauss_zone_27"
	case PCS_Pulkovo_Gauss_zone_28:
		return "PCS_Pulkovo_Gauss_zone_28"
	case PCS_Pulkovo_Gauss_zone_29:
		return "PCS_Pulkovo_Gauss_zone_29"
	case PCS_Pulkovo_Gauss_zone_30:
		return "PCS_Pulkovo_Gauss_zone_30"
	case PCS_Pulkovo_Gauss_zone_31:
		return "PCS_Pulkovo_Gauss_zone_31"
	case PCS_Pulkovo_Gauss_zone_32:
		return "PCS_Pulkovo_Gauss_zone_32"
	case PCS_Pulkovo_Gauss_4N:
		return "PCS_Pulkovo_Gauss_4N"
	case PCS_Pulkovo_Gauss_5N:
		return "PCS_Pulkovo_Gauss_5N"
	case PCS_Pulkovo_Gauss_6N:
		return "PCS_Pulkovo_Gauss_6N"
	case PCS_Pulkovo_Gauss_7N:
		return "PCS_Pulkovo_Gauss_7N"
	case PCS_Pulkovo_Gauss_8N:
		return "PCS_Pulkovo_Gauss_8N"
	case PCS_Pulkovo_Gauss_9N:
		return "PCS_Pulkovo_Gauss_9N"
	case PCS_Pulkovo_Gauss_10N:
		return "PCS_Pulkovo_Gauss_10N"
	case PCS_Pulkovo_Gauss_11N:
		return "PCS_Pulkovo_Gauss_11N"
	case PCS_Pulkovo_Gauss_12N:
		return "PCS_Pulkovo_Gauss_12N"
	case PCS_Pulkovo_Gauss_13N:
		return "PCS_Pulkovo_Gauss_13N"
	case PCS_Pulkovo_Gauss_14N:
		return "PCS_Pulkovo_Gauss_14N"
	case PCS_Pulkovo_Gauss_15N:
		return "PCS_Pulkovo_Gauss_15N"
	case PCS_Pulkovo_Gauss_16N:
		return "PCS_Pulkovo_Gauss_16N"
	case PCS_Pulkovo_Gauss_17N:
		return "PCS_Pulkovo_Gauss_17N"
	case PCS_Pulkovo_Gauss_18N:
		return "PCS_Pulkovo_Gauss_18N"
	case PCS_Pulkovo_Gauss_19N:
		return "PCS_Pulkovo_Gauss_19N"
	case PCS_Pulkovo_Gauss_20N:
		return "PCS_Pulkovo_Gauss_20N"
	case PCS_Pulkovo_Gauss_21N:
		return "PCS_Pulkovo_Gauss_21N"
	case PCS_Pulkovo_Gauss_22N:
		return "PCS_Pulkovo_Gauss_22N"
	case PCS_Pulkovo_Gauss_23N:
		return "PCS_Pulkovo_Gauss_23N"
	case PCS_Pulkovo_Gauss_24N:
		return "PCS_Pulkovo_Gauss_24N"
	case PCS_Pulkovo_Gauss_25N:
		return "PCS_Pulkovo_Gauss_25N"
	case PCS_Pulkovo_Gauss_26N:
		return "PCS_Pulkovo_Gauss_26N"
	case PCS_Pulkovo_Gauss_27N:
		return "PCS_Pulkovo_Gauss_27N"
	case PCS_Pulkovo_Gauss_28N:
		return "PCS_Pulkovo_Gauss_28N"
	case PCS_Pulkovo_Gauss_29N:
		return "PCS_Pulkovo_Gauss_29N"
	case PCS_Pulkovo_Gauss_30N:
		return "PCS_Pulkovo_Gauss_30N"
	case PCS_Pulkovo_Gauss_31N:
		return "PCS_Pulkovo_Gauss_31N"
	case PCS_Pulkovo_Gauss_32N:
		return "PCS_Pulkovo_Gauss_32N"
	case PCS_Qatar_National_Grid:
		return "PCS_Qatar_National_Grid"
	case PCS_RD_Netherlands_Old:
		return "PCS_RD_Netherlands_Old"
	case PCS_RD_Netherlands_New:
		return "PCS_RD_Netherlands_New"
	case PCS_SAD69_UTM_zone_18N:
		return "PCS_SAD69_UTM_zone_18N"
	case PCS_SAD69_UTM_zone_19N:
		return "PCS_SAD69_UTM_zone_19N"
	case PCS_SAD69_UTM_zone_20N:
		return "PCS_SAD69_UTM_zone_20N"
	case PCS_SAD69_UTM_zone_21N:
		return "PCS_SAD69_UTM_zone_21N"
	case PCS_SAD69_UTM_zone_22N:
		return "PCS_SAD69_UTM_zone_22N"
	case PCS_SAD69_UTM_zone_17S:
		return "PCS_SAD69_UTM_zone_17S"
	case PCS_SAD69_UTM_zone_18S:
		return "PCS_SAD69_UTM_zone_18S"
	case PCS_SAD69_UTM_zone_19S:
		return "PCS_SAD69_UTM_zone_19S"
	case PCS_SAD69_UTM_zone_20S:
		return "PCS_SAD69_UTM_zone_20S"
	case PCS_SAD69_UTM_zone_21S:
		return "PCS_SAD69_UTM_zone_21S"
	case PCS_SAD69_UTM_zone_22S:
		return "PCS_SAD69_UTM_zone_22S"
	case PCS_SAD69_UTM_zone_23S:
		return "PCS_SAD69_UTM_zone_23S"
	case PCS_SAD69_UTM_zone_24S:
		return "PCS_SAD69_UTM_zone_24S"
	case PCS_SAD69_UTM_zone_25S:
		return "PCS_SAD69_UTM_zone_25S"
	case PCS_Sapper_Hill_UTM_20S:
		return "PCS_Sapper_Hill_UTM_20S"
	case PCS_Sapper_Hill_UTM_21S:
		return "PCS_Sapper_Hill_UTM_21S"
	case PCS_Schwarzeck_UTM_33S:
		return "PCS_Schwarzeck_UTM_33S"
	case PCS_Sudan_UTM_zone_35N:
		return "PCS_Sudan_UTM_zone_35N"
	case PCS_Sudan_UTM_zone_36N:
		return "PCS_Sudan_UTM_zone_36N"
	case PCS_Tananarive_Laborde:
		return "PCS_Tananarive_Laborde"
	case PCS_Tananarive_UTM_38S:
		return "PCS_Tananarive_UTM_38S"
	case PCS_Tananarive_UTM_39S:
		return "PCS_Tananarive_UTM_39S"
	case PCS_Timbalai_1948_Borneo:
		return "PCS_Timbalai_1948_Borneo"
	case PCS_Timbalai_1948_UTM_49N:
		return "PCS_Timbalai_1948_UTM_49N"
	case PCS_Timbalai_1948_UTM_50N:
		return "PCS_Timbalai_1948_UTM_50N"
	case PCS_TM65_Irish_Nat_Grid:
		return "PCS_TM65_Irish_Nat_Grid"
	case PCS_Trinidad_1903_Trinidad:
		return "PCS_Trinidad_1903_Trinidad"
	case PCS_TC_1948_UTM_zone_39N:
		return "PCS_TC_1948_UTM_zone_39N"
	case PCS_TC_1948_UTM_zone_40N:
		return "PCS_TC_1948_UTM_zone_40N"
	case PCS_Voirol_N_Algerie_ancien:
		return "PCS_Voirol_N_Algerie_ancien"
	case PCS_Voirol_S_Algerie_ancien:
		return "PCS_Voirol_S_Algerie_ancien"
	case PCS_Voirol_Unifie_N_Algerie:
		return "PCS_Voirol_Unifie_N_Algerie"
	case PCS_Voirol_Unifie_S_Algerie:
		return "PCS_Voirol_Unifie_S_Algerie"
	case PCS_Bern_1938_Swiss_New:
		return "PCS_Bern_1938_Swiss_New"
	case PCS_Nord_Sahara_UTM_29N:
		return "PCS_Nord_Sahara_UTM_29N"
	case PCS_Nord_Sahara_UTM_30N:
		return "PCS_Nord_Sahara_UTM_30N"
	case PCS_Nord_Sahara_UTM_31N:
		return "PCS_Nord_Sahara_UTM_31N"
	case PCS_Nord_Sahara_UTM_32N:
		return "PCS_Nord_Sahara_UTM_32N"
	case PCS_Yoff_UTM_zone_28N:
		return "PCS_Yoff_UTM_zone_28N"
	case PCS_Zanderij_UTM_zone_21N:
		return "PCS_Zanderij_UTM_zone_21N"
	case PCS_MGI_Austria_West:
		return "PCS_MGI_Austria_West"
	case PCS_MGI_Austria_Central:
		return "PCS_MGI_Austria_Central"
	case PCS_MGI_Austria_East:
		return "PCS_MGI_Austria_East"
	case PCS_Belge_Lambert_72:
		return "PCS_Belge_Lambert_72"
	case PCS_DHDN_Germany_zone_1:
		return "PCS_DHDN_Germany_zone_1"
	case PCS_DHDN_Germany_zone_2:
		return "PCS_DHDN_Germany_zone_2"
	case PCS_DHDN_Germany_zone_3:
		return "PCS_DHDN_Germany_zone_3"
	case PCS_DHDN_Germany_zone_4:
		return "PCS_DHDN_Germany_zone_4"
	case PCS_DHDN_Germany_zone_5:
		return "PCS_DHDN_Germany_zone_5"
	case PCS_NAD27_Montana_North:
		return "PCS_NAD27_Montana_North"
	case PCS_NAD27_Montana_Central:
		return "PCS_NAD27_Montana_Central"
	case PCS_NAD27_Montana_South:
		return "PCS_NAD27_Montana_South"
	case PCS_NAD27_Nebraska_North:
		return "PCS_NAD27_Nebraska_North"
	case PCS_NAD27_Nebraska_South:
		return "PCS_NAD27_Nebraska_South"
	case PCS_NAD27_Nevada_East:
		return "PCS_NAD27_Nevada_East"
	case PCS_NAD27_Nevada_Central:
		return "PCS_NAD27_Nevada_Central"
	case PCS_NAD27_Nevada_West:
		return "PCS_NAD27_Nevada_West"
	case PCS_NAD27_New_Hampshire:
		return "PCS_NAD27_New_Hampshire"
	case PCS_NAD27_New_Jersey:
		return "PCS_NAD27_New_Jersey"
	case PCS_NAD27_New_Mexico_East:
		return "PCS_NAD27_New_Mexico_East"
	case PCS_NAD27_New_Mexico_Cent:
		return "PCS_NAD27_New_Mexico_Cent"
	case PCS_NAD27_New_Mexico_West:
		return "PCS_NAD27_New_Mexico_West"
	case PCS_NAD27_New_York_East:
		return "PCS_NAD27_New_York_East"
	case PCS_NAD27_New_York_Central:
		return "PCS_NAD27_New_York_Central"
	case PCS_NAD27_New_York_West:
		return "PCS_NAD27_New_York_West"
	case PCS_NAD27_New_York_Long_Is:
		return "PCS_NAD27_New_York_Long_Is"
	case PCS_NAD27_North_Carolina:
		return "PCS_NAD27_North_Carolina"
	case PCS_NAD27_North_Dakota_N:
		return "PCS_NAD27_North_Dakota_N"
	case PCS_NAD27_North_Dakota_S:
		return "PCS_NAD27_North_Dakota_S"
	case PCS_NAD27_Ohio_North:
		return "PCS_NAD27_Ohio_North"
	case PCS_NAD27_Ohio_South:
		return "PCS_NAD27_Ohio_South"
	case PCS_NAD27_Oklahoma_North:
		return "PCS_NAD27_Oklahoma_North"
	case PCS_NAD27_Oklahoma_South:
		return "PCS_NAD27_Oklahoma_South"
	case PCS_NAD27_Oregon_North:
		return "PCS_NAD27_Oregon_North"
	case PCS_NAD27_Oregon_South:
		return "PCS_NAD27_Oregon_South"
	case PCS_NAD27_Pennsylvania_N:
		return "PCS_NAD27_Pennsylvania_N"
	case PCS_NAD27_Pennsylvania_S:
		return "PCS_NAD27_Pennsylvania_S"
	case PCS_NAD27_Rhode_Island:
		return "PCS_NAD27_Rhode_Island"
	case PCS_NAD27_South_Carolina_N:
		return "PCS_NAD27_South_Carolina_N"
	case PCS_NAD27_South_Carolina_S:
		return "PCS_NAD27_South_Carolina_S"
	case PCS_NAD27_South_Dakota_N:
		return "PCS_NAD27_South_Dakota_N"
	case PCS_NAD27_South_Dakota_S:
		return "PCS_NAD27_South_Dakota_S"
	case PCS_NAD27_Tennessee:
		return "PCS_NAD27_Tennessee"
	case PCS_NAD27_Texas_North:
		return "PCS_NAD27_Texas_North"
	case PCS_NAD27_Texas_North_Cen:
		return "PCS_NAD27_Texas_North_Cen"
	case PCS_NAD27_Texas_Central:
		return "PCS_NAD27_Texas_Central"
	case PCS_NAD27_Texas_South_Cen:
		return "PCS_NAD27_Texas_South_Cen"
	case PCS_NAD27_Texas_South:
		return "PCS_NAD27_Texas_South"
	case PCS_NAD27_Utah_North:
		return "PCS_NAD27_Utah_North"
	case PCS_NAD27_Utah_Central:
		return "PCS_NAD27_Utah_Central"
	case PCS_NAD27_Utah_South:
		return "PCS_NAD27_Utah_South"
	case PCS_NAD27_Vermont:
		return "PCS_NAD27_Vermont"
	case PCS_NAD27_Virginia_North:
		return "PCS_NAD27_Virginia_North"
	case PCS_NAD27_Virginia_South:
		return "PCS_NAD27_Virginia_South"
	case PCS_NAD27_Washington_North:
		return "PCS_NAD27_Washington_North"
	case PCS_NAD27_Washington_South:
		return "PCS_NAD27_Washington_South"
	case PCS_NAD27_West_Virginia_N:
		return "PCS_NAD27_West_Virginia_N"
	case PCS_NAD27_West_Virginia_S:
		return "PCS_NAD27_West_Virginia_S"
	case PCS_NAD27_Wisconsin_North:
		return "PCS_NAD27_Wisconsin_North"
	case PCS_NAD27_Wisconsin_Cen:
		return "PCS_NAD27_Wisconsin_Cen"
	case PCS_NAD27_Wisconsin_South:
		return "PCS_NAD27_Wisconsin_South"
	case PCS_NAD27_Wyoming_East:
		return "PCS_NAD27_Wyoming_East"
	case PCS_NAD27_Wyoming_E_Cen:
		return "PCS_NAD27_Wyoming_E_Cen"
	case PCS_NAD27_Wyoming_W_Cen:
		return "PCS_NAD27_Wyoming_W_Cen"
	case PCS_NAD27_Wyoming_West:
		return "PCS_NAD27_Wyoming_West"
	case PCS_NAD27_Puerto_Rico:
		return "PCS_NAD27_Puerto_Rico"
	case PCS_NAD27_St_Croix:
		return "PCS_NAD27_St_Croix"
	case PCS_NAD83_Montana:
		return "PCS_NAD83_Montana"
	case PCS_NAD83_Nebraska:
		return "PCS_NAD83_Nebraska"
	case PCS_NAD83_Nevada_East:
		return "PCS_NAD83_Nevada_East"
	case PCS_NAD83_Nevada_Central:
		return "PCS_NAD83_Nevada_Central"
	case PCS_NAD83_Nevada_West:
		return "PCS_NAD83_Nevada_West"
	case PCS_NAD83_New_Hampshire:
		return "PCS_NAD83_New_Hampshire"
	case PCS_NAD83_New_Jersey:
		return "PCS_NAD83_New_Jersey"
	case PCS_NAD83_New_Mexico_East:
		return "PCS_NAD83_New_Mexico_East"
	case PCS_NAD83_New_Mexico_Cent:
		return "PCS_NAD83_New_Mexico_Cent"
	case PCS_NAD83_New_Mexico_West:
		return "PCS_NAD83_New_Mexico_West"
	case PCS_NAD83_New_York_East:
		return "PCS_NAD83_New_York_East"
	case PCS_NAD83_New_York_Central:
		return "PCS_NAD83_New_York_Central"
	case PCS_NAD83_New_York_West:
		return "PCS_NAD83_New_York_West"
	case PCS_NAD83_New_York_Long_Is:
		return "PCS_NAD83_New_York_Long_Is"
	case PCS_NAD83_North_Carolina:
		return "PCS_NAD83_North_Carolina"
	case PCS_NAD83_North_Dakota_N:
		return "PCS_NAD83_North_Dakota_N"
	case PCS_NAD83_North_Dakota_S:
		return "PCS_NAD83_North_Dakota_S"
	case PCS_NAD83_Ohio_North:
		return "PCS_NAD83_Ohio_North"
	case PCS_NAD83_Ohio_South:
		return "PCS_NAD83_Ohio_South"
	case PCS_NAD83_Oklahoma_North:
		return "PCS_NAD83_Oklahoma_North"
	case PCS_NAD83_Oklahoma_South:
		return "PCS_NAD83_Oklahoma_South"
	case PCS_NAD83_Oregon_North:
		return "PCS_NAD83_Oregon_North"
	case PCS_NAD83_Oregon_South:
		return "PCS_NAD83_Oregon_South"
	case PCS_NAD83_Pennsylvania_N:
		return "PCS_NAD83_Pennsylvania_N"
	case PCS_NAD83_Pennsylvania_S:
		return "PCS_NAD83_Pennsylvania_S"
	case PCS_NAD83_Rhode_Island:
		return "PCS_NAD83_Rhode_Island"
	case PCS_NAD83_South_Carolina:
		return "PCS_NAD83_South_Carolina"
	case PCS_NAD83_South_Dakota_N:
		return "PCS_NAD83_South_Dakota_N"
	case PCS_NAD83_South_Dakota_S:
		return "PCS_NAD83_South_Dakota_S"
	case PCS_NAD83_Tennessee:
		return "PCS_NAD83_Tennessee"
	case PCS_NAD83_Texas_North:
		return "PCS_NAD83_Texas_North"
	case PCS_NAD83_Texas_North_Cen:
		return "PCS_NAD83_Texas_North_Cen"
	case PCS_NAD83_Texas_Central:
		return "PCS_NAD83_Texas_Central"
	case PCS_NAD83_Texas_South_Cen:
		return "PCS_NAD83_Texas_South_Cen"
	case PCS_NAD83_Texas_South:
		return "PCS_NAD83_Texas_South"
	case PCS_NAD83_Utah_North:
		return "PCS_NAD83_Utah_North"
	case PCS_NAD83_Utah_Central:
		return "PCS_NAD83_Utah_Central"
	case PCS_NAD83_Utah_South:
		return "PCS_NAD83_Utah_South"
	case PCS_NAD83_Vermont:
		return "PCS_NAD83_Vermont"
	case PCS_NAD83_Virginia_North:
		return "PCS_NAD83_Virginia_North"
	case PCS_NAD83_Virginia_South:
		return "PCS_NAD83_Virginia_South"
	case PCS_NAD83_Washington_North:
		return "PCS_NAD83_Washington_North"
	case PCS_NAD83_Washington_South:
		return "PCS_NAD83_Washington_South"
	case PCS_NAD83_West_Virginia_N:
		return "PCS_NAD83_West_Virginia_N"
	case PCS_NAD83_West_Virginia_S:
		return "PCS_NAD83_West_Virginia_S"
	case PCS_NAD83_Wisconsin_North:
		return "PCS_NAD83_Wisconsin_North"
	case PCS_NAD83_Wisconsin_Cen:
		return "PCS_NAD83_Wisconsin_Cen"
	case PCS_NAD83_Wisconsin_South:
		return "PCS_NAD83_Wisconsin_South"
	case PCS_NAD83_Wyoming_East:
		return "PCS_NAD83_Wyoming_East"
	case PCS_NAD83_Wyoming_E_Cen:
		return "PCS_NAD83_Wyoming_E_Cen"
	case PCS_NAD83_Wyoming_W_Cen:
		return "PCS_NAD83_Wyoming_W_Cen"
	case PCS_NAD83_Wyoming_West:
		return "PCS_NAD83_Wyoming_West"
	case PCS_NAD83_Puerto_Rico_Virgin_Is:
		return "PCS_NAD83_Puerto_Rico_Virgin_Is"
	case PCS_WGS72_UTM_zone_1N:
		return "PCS_WGS72_UTM_zone_1N"
	case PCS_WGS72_UTM_zone_2N:
		return "PCS_WGS72_UTM_zone_2N"
	case PCS_WGS72_UTM_zone_3N:
		return "PCS_WGS72_UTM_zone_3N"
	case PCS_WGS72_UTM_zone_4N:
		return "PCS_WGS72_UTM_zone_4N"
	case PCS_WGS72_UTM_zone_5N:
		return "PCS_WGS72_UTM_zone_5N"
	case PCS_WGS72_UTM_zone_6N:
		return "PCS_WGS72_UTM_zone_6N"
	case PCS_WGS72_UTM_zone_7N:
		return "PCS_WGS72_UTM_zone_7N"
	case PCS_WGS72_UTM_zone_8N:
		return "PCS_WGS72_UTM_zone_8N"
	case PCS_WGS72_UTM_zone_9N:
		return "PCS_WGS72_UTM_zone_9N"
	case PCS_WGS72_UTM_zone_10N:
		return "PCS_WGS72_UTM_zone_10N"
	case PCS_WGS72_UTM_zone_11N:
		return "PCS_WGS72_UTM_zone_11N"
	case PCS_WGS72_UTM_zone_12N:
		return "PCS_WGS72_UTM_zone_12N"
	case PCS_WGS72_UTM_zone_13N:
		return "PCS_WGS72_UTM_zone_13N"
	case PCS_WGS72_UTM_zone_14N:
		return "PCS_WGS72_UTM_zone_14N"
	case PCS_WGS72_UTM_zone_15N:
		return "PCS_WGS72_UTM_zone_15N"
	case PCS_WGS72_UTM_zone_16N:
		return "PCS_WGS72_UTM_zone_16N"
	case PCS_WGS72_UTM_zone_17N:
		return "PCS_WGS72_UTM_zone_17N"
	case PCS_WGS72_UTM_zone_18N:
		return "PCS_WGS72_UTM_zone_18N"
	case PCS_WGS72_UTM_zone_19N:
		return "PCS_WGS72_UTM_zone_19N"
	case PCS_WGS72_UTM_zone_20N:
		return "PCS_WGS72_UTM_zone_20N"
	case PCS_WGS72_UTM_zone_21N:
		return "PCS_WGS72_UTM_zone_21N"
	case PCS_WGS72_UTM_zone_22N:
		return "PCS_WGS72_UTM_zone_22N"
	case PCS_WGS72_UTM_zone_23N:
		return "PCS_WGS72_UTM_zone_23N"
	case PCS_WGS72_UTM_zone_24N:
		return "PCS_WGS72_UTM_zone_24N"
	case PCS_WGS72_UTM_zone_25N:
		return "PCS_WGS72_UTM_zone_25N"
	case PCS_WGS72_UTM_zone_26N:
		return "PCS_WGS72_UTM_zone_26N"
	case PCS_WGS72_UTM_zone_27N:
		return "PCS_WGS72_UTM_zone_27N"
	case PCS_WGS72_UTM_zone_28N:
		return "PCS_WGS72_UTM_zone_28N"
	case PCS_WGS72_UTM_zone_29N:
		return "PCS_WGS72_UTM_zone_29N"
	case PCS_WGS72_UTM_zone_30N:
		return "PCS_WGS72_UTM_zone_30N"
	case PCS_WGS72_UTM_zone_31N:
		return "PCS_WGS72_UTM_zone_31N"
	case PCS_WGS72_UTM_zone_32N:
		return "PCS_WGS72_UTM_zone_32N"
	case PCS_WGS72_UTM_zone_33N:
		return "PCS_WGS72_UTM_zone_33N"
	case PCS_WGS72_UTM_zone_34N:
		return "PCS_WGS72_UTM_zone_34N"
	case PCS_WGS72_UTM_zone_35N:
		return "PCS_WGS72_UTM_zone_35N"
	case PCS_WGS72_UTM_zone_36N:
		return "PCS_WGS72_UTM_zone_36N"
	case PCS_WGS72_UTM_zone_37N:
		return "PCS_WGS72_UTM_zone_37N"
	case PCS_WGS72_UTM_zone_38N:
		return "PCS_WGS72_UTM_zone_38N"
	case PCS_WGS72_UTM_zone_39N:
		return "PCS_WGS72_UTM_zone_39N"
	case PCS_WGS72_UTM_zone_40N:
		return "PCS_WGS72_UTM_zone_40N"
	case PCS_WGS72_UTM_zone_41N:
		return "PCS_WGS72_UTM_zone_41N"
	case PCS_WGS72_UTM_zone_42N:
		return "PCS_WGS72_UTM_zone_42N"
	case PCS_WGS72_UTM_zone_43N:
		return "PCS_WGS72_UTM_zone_43N"
	case PCS_WGS72_UTM_zone_44N:
		return "PCS_WGS72_UTM_zone_44N"
	case PCS_WGS72_UTM_zone_45N:
		return "PCS_WGS72_UTM_zone_45N"
	case PCS_WGS72_UTM_zone_46N:
		return "PCS_WGS72_UTM_zone_46N"
	case PCS_WGS72_UTM_zone_47N:
		return "PCS_WGS72_UTM_zone_47N"
	case PCS_WGS72_UTM_zone_48N:
		return "PCS_WGS72_UTM_zone_48N"
	case PCS_WGS72_UTM_zone_49N:
		return "PCS_WGS72_UTM_zone_49N"
	case PCS_WGS72_UTM_zone_50N:
		return "PCS_WGS72_UTM_zone_50N"
	case PCS_WGS72_UTM_zone_51N:
		return "PCS_WGS72_UTM_zone_51N"
	case PCS_WGS72_UTM_zone_52N:
		return "PCS_WGS72_UTM_zone_52N"
	case PCS_WGS72_UTM_zone_53N:
		return "PCS_WGS72_UTM_zone_53N"
	case PCS_WGS72_UTM_zone_54N:
		return "PCS_WGS72_UTM_zone_54N"
	case PCS_WGS72_UTM_zone_55N:
		return "PCS_WGS72_UTM_zone_55N"
	case PCS_WGS72_UTM_zone_56N:
		return "PCS_WGS72_UTM_zone_56N"
	case PCS_WGS72_UTM_zone_57N:
		return "PCS_WGS72_UTM_zone_57N"
	case PCS_WGS72_UTM_zone_58N:
		return "PCS_WGS72_UTM_zone_58N"
	case PCS_WGS72_UTM_zone_59N:
		return "PCS_WGS72_UTM_zone_59N"
	case PCS_WGS72_UTM_zone_60N:
		return "PCS_WGS72_UTM_zone_60N"
	case PCS_WGS72_UTM_zone_1S:
		return "PCS_WGS72_UTM_zone_1S"
	case PCS_WGS72_UTM_zone_2S:
		return "PCS_WGS72_UTM_zone_2S"
	case PCS_WGS72_UTM_zone_3S:
		return "PCS_WGS72_UTM_zone_3S"
	case PCS_WGS72_UTM_zone_4S:
		return "PCS_WGS72_UTM_zone_4S"
	case PCS_WGS72_UTM_zone_5S:
		return "PCS_WGS72_UTM_zone_5S"
	case PCS_WGS72_UTM_zone_6S:
		return "PCS_WGS72_UTM_zone_6S"
	case PCS_WGS72_UTM_zone_7S:
		return "PCS_WGS72_UTM_zone_7S"
	case PCS_WGS72_UTM_zone_8S:
		return "PCS_WGS72_UTM_zone_8S"
	case PCS_WGS72_UTM_zone_9S:
		return "PCS_WGS72_UTM_zone_9S"
	case PCS_WGS72_UTM_zone_10S:
		return "PCS_WGS72_UTM_zone_10S"
	case PCS_WGS72_UTM_zone_11S:
		return "PCS_WGS72_UTM_zone_11S"
	case PCS_WGS72_UTM_zone_12S:
		return "PCS_WGS72_UTM_zone_12S"
	case PCS_WGS72_UTM_zone_13S:
		return "PCS_WGS72_UTM_zone_13S"
	case PCS_WGS72_UTM_zone_14S:
		return "PCS_WGS72_UTM_zone_14S"
	case PCS_WGS72_UTM_zone_15S:
		return "PCS_WGS72_UTM_zone_15S"
	case PCS_WGS72_UTM_zone_16S:
		return "PCS_WGS72_UTM_zone_16S"
	case PCS_WGS72_UTM_zone_17S:
		return "PCS_WGS72_UTM_zone_17S"
	case PCS_WGS72_UTM_zone_18S:
		return "PCS_WGS72_UTM_zone_18S"
	case PCS_WGS72_UTM_zone_19S:
		return "PCS_WGS72_UTM_zone_19S"
	case PCS_WGS72_UTM_zone_20S:
		return "PCS_WGS72_UTM_zone_20S"
	case PCS_WGS72_UTM_zone_21S:
		return "PCS_WGS72_UTM_zone_21S"
	case PCS_WGS72_UTM_zone_22S:
		return "PCS_WGS72_UTM_zone_22S"
	case PCS_WGS72_UTM_zone_23S:
		return "PCS_WGS72_UTM_zone_23S"
	case PCS_WGS72_UTM_zone_24S:
		return "PCS_WGS72_UTM_zone_24S"
	case PCS_WGS72_UTM_zone_25S:
		return "PCS_WGS72_UTM_zone_25S"
	case PCS_WGS72_UTM_zone_26S:
		return "PCS_WGS72_UTM_zone_26S"
	case PCS_WGS72_UTM_zone_27S:
		return "PCS_WGS72_UTM_zone_27S"
	case PCS_WGS72_UTM_zone_28S:
		return "PCS_WGS72_UTM_zone_28S"
	case PCS_WGS72_UTM_zone_29S:
		return "PCS_WGS72_UTM_zone_29S"
	case PCS_WGS72_UTM_zone_30S:
		return "PCS_WGS72_UTM_zone_30S"
	case PCS_WGS72_UTM_zone_31S:
		return "PCS_WGS72_UTM_zone_31S"
	case PCS_WGS72_UTM_zone_32S:
		return "PCS_WGS72_UTM_zone_32S"
	case PCS_WGS72_UTM_zone_33S:
		return "PCS_WGS72_UTM_zone_33S"
	case PCS_WGS72_UTM_zone_34S:
		return "PCS_WGS72_UTM_zone_34S"
	case PCS_WGS72_UTM_zone_35S:
		return "PCS_WGS72_UTM_zone_35S"
	case PCS_WGS72_UTM_zone_36S:
		return "PCS_WGS72_UTM_zone_36S"
	case PCS_WGS72_UTM_zone_37S:
		return "PCS_WGS72_UTM_zone_37S"
	case PCS_WGS72_UTM_zone_38S:
		return "PCS_WGS72_UTM_zone_38S"
	case PCS_WGS72_UTM_zone_39S:
		return "PCS_WGS72_UTM_zone_39S"
	case PCS_WGS72_UTM_zone_40S:
		return "PCS_WGS72_UTM_zone_40S"
	case PCS_WGS72_UTM_zone_41S:
		return "PCS_WGS72_UTM_zone_41S"
	case PCS_WGS72_UTM_zone_42S:
		return "PCS_WGS72_UTM_zone_42S"
	case PCS_WGS72_UTM_zone_43S:
		return "PCS_WGS72_UTM_zone_43S"
	case PCS_WGS72_UTM_zone_44S:
		return "PCS_WGS72_UTM_zone_44S"
	case PCS_WGS72_UTM_zone_45S:
		return "PCS_WGS72_UTM_zone_45S"
	case PCS_WGS72_UTM_zone_46S:
		return "PCS_WGS72_UTM_zone_46S"
	case PCS_WGS72_UTM_zone_47S:
		return "PCS_WGS72_UTM_zone_47S"
	case PCS_WGS72_UTM_zone_48S:
		return "PCS_WGS72_UTM_zone_48S"
	case PCS_WGS72_UTM_zone_49S:
		return "PCS_WGS72_UTM_zone_49S"
	case PCS_WGS72_UTM_zone_50S:
		return "PCS_WGS72_UTM_zone_50S"
	case PCS_WGS72_UTM_zone_51S:
		return "PCS_WGS72_UTM_zone_51S"
	case PCS_WGS72_UTM_zone_52S:
		return "PCS_WGS72_UTM_zone_52S"
	case PCS_WGS72_UTM_zone_53S:
		return "PCS_WGS72_UTM_zone_53S"
	case PCS_WGS72_UTM_zone_54S:
		return "PCS_WGS72_UTM_zone_54S"
	case PCS_WGS72_UTM_zone_55S:
		return "PCS_WGS72_UTM_zone_55S"
	case PCS_WGS72_UTM_zone_56S:
		return "PCS_WGS72_UTM_zone_56S"
	case PCS_WGS72_UTM_zone_57S:
		return "PCS_WGS72_UTM_zone_57S"
	case PCS_WGS72_UTM_zone_58S:
		return "PCS_WGS72_UTM_zone_58S"
	case PCS_WGS72_UTM_zone_59S:
		return "PCS_WGS72_UTM_zone_59S"
	case PCS_WGS72_UTM_zone_60S:
		return "PCS_WGS72_UTM_zone_60S"
	case PCS_WGS72BE_UTM_zone_1N:
		return "PCS_WGS72BE_UTM_zone_1N"
	case PCS_WGS72BE_UTM_zone_2N:
		return "PCS_WGS72BE_UTM_zone_2N"
	case PCS_WGS72BE_UTM_zone_3N:
		return "PCS_WGS72BE_UTM_zone_3N"
	case PCS_WGS72BE_UTM_zone_4N:
		return "PCS_WGS72BE_UTM_zone_4N"
	case PCS_WGS72BE_UTM_zone_5N:
		return "PCS_WGS72BE_UTM_zone_5N"
	case PCS_WGS72BE_UTM_zone_6N:
		return "PCS_WGS72BE_UTM_zone_6N"
	case PCS_WGS72BE_UTM_zone_7N:
		return "PCS_WGS72BE_UTM_zone_7N"
	case PCS_WGS72BE_UTM_zone_8N:
		return "PCS_WGS72BE_UTM_zone_8N"
	case PCS_WGS72BE_UTM_zone_9N:
		return "PCS_WGS72BE_UTM_zone_9N"
	case PCS_WGS72BE_UTM_zone_10N:
		return "PCS_WGS72BE_UTM_zone_10N"
	case PCS_WGS72BE_UTM_zone_11N:
		return "PCS_WGS72BE_UTM_zone_11N"
	case PCS_WGS72BE_UTM_zone_12N:
		return "PCS_WGS72BE_UTM_zone_12N"
	case PCS_WGS72BE_UTM_zone_13N:
		return "PCS_WGS72BE_UTM_zone_13N"
	case PCS_WGS72BE_UTM_zone_14N:
		return "PCS_WGS72BE_UTM_zone_14N"
	case PCS_WGS72BE_UTM_zone_15N:
		return "PCS_WGS72BE_UTM_zone_15N"
	case PCS_WGS72BE_UTM_zone_16N:
		return "PCS_WGS72BE_UTM_zone_16N"
	case PCS_WGS72BE_UTM_zone_17N:
		return "PCS_WGS72BE_UTM_zone_17N"
	case PCS_WGS72BE_UTM_zone_18N:
		return "PCS_WGS72BE_UTM_zone_18N"
	case PCS_WGS72BE_UTM_zone_19N:
		return "PCS_WGS72BE_UTM_zone_19N"
	case PCS_WGS72BE_UTM_zone_20N:
		return "PCS_WGS72BE_UTM_zone_20N"
	case PCS_WGS72BE_UTM_zone_21N:
		return "PCS_WGS72BE_UTM_zone_21N"
	case PCS_WGS72BE_UTM_zone_22N:
		return "PCS_WGS72BE_UTM_zone_22N"
	case PCS_WGS72BE_UTM_zone_23N:
		return "PCS_WGS72BE_UTM_zone_23N"
	case PCS_WGS72BE_UTM_zone_24N:
		return "PCS_WGS72BE_UTM_zone_24N"
	case PCS_WGS72BE_UTM_zone_25N:
		return "PCS_WGS72BE_UTM_zone_25N"
	case PCS_WGS72BE_UTM_zone_26N:
		return "PCS_WGS72BE_UTM_zone_26N"
	case PCS_WGS72BE_UTM_zone_27N:
		return "PCS_WGS72BE_UTM_zone_27N"
	case PCS_WGS72BE_UTM_zone_28N:
		return "PCS_WGS72BE_UTM_zone_28N"
	case PCS_WGS72BE_UTM_zone_29N:
		return "PCS_WGS72BE_UTM_zone_29N"
	case PCS_WGS72BE_UTM_zone_30N:
		return "PCS_WGS72BE_UTM_zone_30N"
	case PCS_WGS72BE_UTM_zone_31N:
		return "PCS_WGS72BE_UTM_zone_31N"
	case PCS_WGS72BE_UTM_zone_32N:
		return "PCS_WGS72BE_UTM_zone_32N"
	case PCS_WGS72BE_UTM_zone_33N:
		return "PCS_WGS72BE_UTM_zone_33N"
	case PCS_WGS72BE_UTM_zone_34N:
		return "PCS_WGS72BE_UTM_zone_34N"
	case PCS_WGS72BE_UTM_zone_35N:
		return "PCS_WGS72BE_UTM_zone_35N"
	case PCS_WGS72BE_UTM_zone_36N:
		return "PCS_WGS72BE_UTM_zone_36N"
	case PCS_WGS72BE_UTM_zone_37N:
		return "PCS_WGS72BE_UTM_zone_37N"
	case PCS_WGS72BE_UTM_zone_38N:
		return "PCS_WGS72BE_UTM_zone_38N"
	case PCS_WGS72BE_UTM_zone_39N:
		return "PCS_WGS72BE_UTM_zone_39N"
	case PCS_WGS72BE_UTM_zone_40N:
		return "PCS_WGS72BE_UTM_zone_40N"
	case PCS_WGS72BE_UTM_zone_41N:
		return "PCS_WGS72BE_UTM_zone_41N"
	case PCS_WGS72BE_UTM_zone_42N:
		return "PCS_WGS72BE_UTM_zone_42N"
	case PCS_WGS72BE_UTM_zone_43N:
		return "PCS_WGS72BE_UTM_zone_43N"
	case PCS_WGS72BE_UTM_zone_44N:
		return "PCS_WGS72BE_UTM_zone_44N"
	case PCS_WGS72BE_UTM_zone_45N:
		return "PCS_WGS72BE_UTM_zone_45N"
	case PCS_WGS72BE_UTM_zone_46N:
		return "PCS_WGS72BE_UTM_zone_46N"
	case PCS_WGS72BE_UTM_zone_47N:
		return "PCS_WGS72BE_UTM_zone_47N"
	case PCS_WGS72BE_UTM_zone_48N:
		return "PCS_WGS72BE_UTM_zone_48N"
	case PCS_WGS72BE_UTM_zone_49N:
		return "PCS_WGS72BE_UTM_zone_49N"
	case PCS_WGS72BE_UTM_zone_50N:
		return "PCS_WGS72BE_UTM_zone_50N"
	case PCS_WGS72BE_UTM_zone_51N:
		return "PCS_WGS72BE_UTM_zone_51N"
	case PCS_WGS72BE_UTM_zone_52N:
		return "PCS_WGS72BE_UTM_zone_52N"
	case PCS_WGS72BE_UTM_zone_53N:
		return "PCS_WGS72BE_UTM_zone_53N"
	case PCS_WGS72BE_UTM_zone_54N:
		return "PCS_WGS72BE_UTM_zone_54N"
	case PCS_WGS72BE_UTM_zone_55N:
		return "PCS_WGS72BE_UTM_zone_55N"
	case PCS_WGS72BE_UTM_zone_56N:
		return "PCS_WGS72BE_UTM_zone_56N"
	case PCS_WGS72BE_UTM_zone_57N:
		return "PCS_WGS72BE_UTM_zone_57N"
	case PCS_WGS72BE_UTM_zone_58N:
		return "PCS_WGS72BE_UTM_zone_58N"
	case PCS_WGS72BE_UTM_zone_59N:
		return "PCS_WGS72BE_UTM_zone_59N"
	case PCS_WGS72BE_UTM_zone_60N:
		return "PCS_WGS72BE_UTM_zone_60N"
	case PCS_WGS72BE_UTM_zone_1S:
		return "PCS_WGS72BE_UTM_zone_1S"
	case PCS_WGS72BE_UTM_zone_2S:
		return "PCS_WGS72BE_UTM_zone_2S"
	case PCS_WGS72BE_UTM_zone_3S:
		return "PCS_WGS72BE_UTM_zone_3S"
	case PCS_WGS72BE_UTM_zone_4S:
		return "PCS_WGS72BE_UTM_zone_4S"
	case PCS_WGS72BE_UTM_zone_5S:
		return "PCS_WGS72BE_UTM_zone_5S"
	case PCS_WGS72BE_UTM_zone_6S:
		return "PCS_WGS72BE_UTM_zone_6S"
	case PCS_WGS72BE_UTM_zone_7S:
		return "PCS_WGS72BE_UTM_zone_7S"
	case PCS_WGS72BE_UTM_zone_8S:
		return "PCS_WGS72BE_UTM_zone_8S"
	case PCS_WGS72BE_UTM_zone_9S:
		return "PCS_WGS72BE_UTM_zone_9S"
	case PCS_WGS72BE_UTM_zone_10S:
		return "PCS_WGS72BE_UTM_zone_10S"
	case PCS_WGS72BE_UTM_zone_11S:
		return "PCS_WGS72BE_UTM_zone_11S"
	case PCS_WGS72BE_UTM_zone_12S:
		return "PCS_WGS72BE_UTM_zone_12S"
	case PCS_WGS72BE_UTM_zone_13S:
		return "PCS_WGS72BE_UTM_zone_13S"
	case PCS_WGS72BE_UTM_zone_14S:
		return "PCS_WGS72BE_UTM_zone_14S"
	case PCS_WGS72BE_UTM_zone_15S:
		return "PCS_WGS72BE_UTM_zone_15S"
	case PCS_WGS72BE_UTM_zone_16S:
		return "PCS_WGS72BE_UTM_zone_16S"
	case PCS_WGS72BE_UTM_zone_17S:
		return "PCS_WGS72BE_UTM_zone_17S"
	case PCS_WGS72BE_UTM_zone_18S:
		return "PCS_WGS72BE_UTM_zone_18S"
	case PCS_WGS72BE_UTM_zone_19S:
		return "PCS_WGS72BE_UTM_zone_19S"
	case PCS_WGS72BE_UTM_zone_20S:
		return "PCS_WGS72BE_UTM_zone_20S"
	case PCS_WGS72BE_UTM_zone_21S:
		return "PCS_WGS72BE_UTM_zone_21S"
	case PCS_WGS72BE_UTM_zone_22S:
		return "PCS_WGS72BE_UTM_zone_22S"
	case PCS_WGS72BE_UTM_zone_23S:
		return "PCS_WGS72BE_UTM_zone_23S"
	case PCS_WGS72BE_UTM_zone_24S:
		return "PCS_WGS72BE_UTM_zone_24S"
	case PCS_WGS72BE_UTM_zone_25S:
		return "PCS_WGS72BE_UTM_zone_25S"
	case PCS_WGS72BE_UTM_zone_26S:
		return "PCS_WGS72BE_UTM_zone_26S"
	case PCS_WGS72BE_UTM_zone_27S:
		return "PCS_WGS72BE_UTM_zone_27S"
	case PCS_WGS72BE_UTM_zone_28S:
		return "PCS_WGS72BE_UTM_zone_28S"
	case PCS_WGS72BE_UTM_zone_29S:
		return "PCS_WGS72BE_UTM_zone_29S"
	case PCS_WGS72BE_UTM_zone_30S:
		return "PCS_WGS72BE_UTM_zone_30S"
	case PCS_WGS72BE_UTM_zone_31S:
		return "PCS_WGS72BE_UTM_zone_31S"
	case PCS_WGS72BE_UTM_zone_32S:
		return "PCS_WGS72BE_UTM_zone_32S"
	case PCS_WGS72BE_UTM_zone_33S:
		return "PCS_WGS72BE_UTM_zone_33S"
	case PCS_WGS72BE_UTM_zone_34S:
		return "PCS_WGS72BE_UTM_zone_34S"
	case PCS_WGS72BE_UTM_zone_35S:
		return "PCS_WGS72BE_UTM_zone_35S"
	case PCS_WGS72BE_UTM_zone_36S:
		return "PCS_WGS72BE_UTM_zone_36S"
	case PCS_WGS72BE_UTM_zone_37S:
		return "PCS_WGS72BE_UTM_zone_37S"
	case PCS_WGS72BE_UTM_zone_38S:
		return "PCS_WGS72BE_UTM_zone_38S"
	case PCS_WGS72BE_UTM_zone_39S:
		return "PCS_WGS72BE_UTM_zone_39S"
	case PCS_WGS72BE_UTM_zone_40S:
		return "PCS_WGS72BE_UTM_zone_40S"
	case PCS_WGS72BE_UTM_zone_41S:
		return "PCS_WGS72BE_UTM_zone_41S"
	case PCS_WGS72BE_UTM_zone_42S:
		return "PCS_WGS72BE_UTM_zone_42S"
	case PCS_WGS72BE_UTM_zone_43S:
		return "PCS_WGS72BE_UTM_zone_43S"
	case PCS_WGS72BE_UTM_zone_44S:
		return "PCS_WGS72BE_UTM_zone_44S"
	case PCS_WGS72BE_UTM_zone_45S:
		return "PCS_WGS72BE_UTM_zone_45S"
	case PCS_WGS72BE_UTM_zone_46S:
		return "PCS_WGS72BE_UTM_zone_46S"
	case PCS_WGS72BE_UTM_zone_47S:
		return "PCS_WGS72BE_UTM_zone_47S"
	case PCS_WGS72BE_UTM_zone_48S:
		return "PCS_WGS72BE_UTM_zone_48S"
	case PCS_WGS72BE_UTM_zone_49S:
		return "PCS_WGS72BE_UTM_zone_49S"
	case PCS_WGS72BE_UTM_zone_50S:
		return "PCS_WGS72BE_UTM_zone_50S"
	case PCS_WGS72BE_UTM_zone_51S:
		return "PCS_WGS72BE_UTM_zone_51S"
	case PCS_WGS72BE_UTM_zone_52S:
		return "PCS_WGS72BE_UTM_zone_52S"
	case PCS_WGS72BE_UTM_zone_53S:
		return "PCS_WGS72BE_UTM_zone_53S"
	case PCS_WGS72BE_UTM_zone_54S:
		return "PCS_WGS72BE_UTM_zone_54S"
	case PCS_WGS72BE_UTM_zone_55S:
		return "PCS_WGS72BE_UTM_zone_55S"
	case PCS_WGS72BE_UTM_zone_56S:
		return "PCS_WGS72BE_UTM_zone_56S"
	case PCS_WGS72BE_UTM_zone_57S:
		return "PCS_WGS72BE_UTM_zone_57S"
	case PCS_WGS72BE_UTM_zone_58S:
		return "PCS_WGS72BE_UTM_zone_58S"
	case PCS_WGS72BE_UTM_zone_59S:
		return "PCS_WGS72BE_UTM_zone_59S"
	case PCS_WGS72BE_UTM_zone_60S:
		return "PCS_WGS72BE_UTM_zone_60S"
	case PCS_WGS84_UTM_zone_1N:
		return "PCS_WGS84_UTM_zone_1N"
	case PCS_WGS84_UTM_zone_2N:
		return "PCS_WGS84_UTM_zone_2N"
	case PCS_WGS84_UTM_zone_3N:
		return "PCS_WGS84_UTM_zone_3N"
	case PCS_WGS84_UTM_zone_4N:
		return "PCS_WGS84_UTM_zone_4N"
	case PCS_WGS84_UTM_zone_5N:
		return "PCS_WGS84_UTM_zone_5N"
	case PCS_WGS84_UTM_zone_6N:
		return "PCS_WGS84_UTM_zone_6N"
	case PCS_WGS84_UTM_zone_7N:
		return "PCS_WGS84_UTM_zone_7N"
	case PCS_WGS84_UTM_zone_8N:
		return "PCS_WGS84_UTM_zone_8N"
	case PCS_WGS84_UTM_zone_9N:
		return "PCS_WGS84_UTM_zone_9N"
	case PCS_WGS84_UTM_zone_10N:
		return "PCS_WGS84_UTM_zone_10N"
	case PCS_WGS84_UTM_zone_11N:
		return "PCS_WGS84_UTM_zone_11N"
	case PCS_WGS84_UTM_zone_12N:
		return "PCS_WGS84_UTM_zone_12N"
	case PCS_WGS84_UTM_zone_13N:
		return "PCS_WGS84_UTM_zone_13N"
	case PCS_WGS84_UTM_zone_14N:
		return "PCS_WGS84_UTM_zone_14N"
	case PCS_WGS84_UTM_zone_15N:
		return "PCS_WGS84_UTM_zone_15N"
	case PCS_WGS84_UTM_zone_16N:
		return "PCS_WGS84_UTM_zone_16N"
	case PCS_WGS84_UTM_zone_17N:
		return "PCS_WGS84_UTM_zone_17N"
	case PCS_WGS84_UTM_zone_18N:
		return "PCS_WGS84_UTM_zone_18N"
	case PCS_WGS84_UTM_zone_19N:
		return "PCS_WGS84_UTM_zone_19N"
	case PCS_WGS84_UTM_zone_20N:
		return "PCS_WGS84_UTM_zone_20N"
	case PCS_WGS84_UTM_zone_21N:
		return "PCS_WGS84_UTM_zone_21N"
	case PCS_WGS84_UTM_zone_22N:
		return "PCS_WGS84_UTM_zone_22N"
	case PCS_WGS84_UTM_zone_23N:
		return "PCS_WGS84_UTM_zone_23N"
	case PCS_WGS84_UTM_zone_24N:
		return "PCS_WGS84_UTM_zone_24N"
	case PCS_WGS84_UTM_zone_25N:
		return "PCS_WGS84_UTM_zone_25N"
	case PCS_WGS84_UTM_zone_26N:
		return "PCS_WGS84_UTM_zone_26N"
	case PCS_WGS84_UTM_zone_27N:
		return "PCS_WGS84_UTM_zone_27N"
	case PCS_WGS84_UTM_zone_28N:
		return "PCS_WGS84_UTM_zone_28N"
	case PCS_WGS84_UTM_zone_29N:
		return "PCS_WGS84_UTM_zone_29N"
	case PCS_WGS84_UTM_zone_30N:
		return "PCS_WGS84_UTM_zone_30N"
	case PCS_WGS84_UTM_zone_31N:
		return "PCS_WGS84_UTM_zone_31N"
	case PCS_WGS84_UTM_zone_32N:
		return "PCS_WGS84_UTM_zone_32N"
	case PCS_WGS84_UTM_zone_33N:
		return "PCS_WGS84_UTM_zone_33N"
	case PCS_WGS84_UTM_zone_34N:
		return "PCS_WGS84_UTM_zone_34N"
	case PCS_WGS84_UTM_zone_35N:
		return "PCS_WGS84_UTM_zone_35N"
	case PCS_WGS84_UTM_zone_36N:
		return "PCS_WGS84_UTM_zone_36N"
	case PCS_WGS84_UTM_zone_37N:
		return "PCS_WGS84_UTM_zone_37N"
	case PCS_WGS84_UTM_zone_38N:
		return "PCS_WGS84_UTM_zone_38N"
	case PCS_WGS84_UTM_zone_39N:
		return "PCS_WGS84_UTM_zone_39N"
	case PCS_WGS84_UTM_zone_40N:
		return "PCS_WGS84_UTM_zone_40N"
	case PCS_WGS84_UTM_zone_41N:
		return "PCS_WGS84_UTM_zone_41N"
	case PCS_WGS84_UTM_zone_42N:
		return "PCS_WGS84_UTM_zone_42N"
	case PCS_WGS84_UTM_zone_43N:
		return "PCS_WGS84_UTM_zone_43N"
	case PCS_WGS84_UTM_zone_44N:
		return "PCS_WGS84_UTM_zone_44N"
	case PCS_WGS84_UTM_zone_45N:
		return "PCS_WGS84_UTM_zone_45N"
	case PCS_WGS84_UTM_zone_46N:
		return "PCS_WGS84_UTM_zone_46N"
	case PCS_WGS84_UTM_zone_47N:
		return "PCS_WGS84_UTM_zone_47N"
	case PCS_WGS84_UTM_zone_48N:
		return "PCS_WGS84_UTM_zone_48N"
	case PCS_WGS84_UTM_zone_49N:
		return "PCS_WGS84_UTM_zone_49N"
	case PCS_WGS84_UTM_zone_50N:
		return "PCS_WGS84_UTM_zone_50N"
	case PCS_WGS84_UTM_zone_51N:
		return "PCS_WGS84_UTM_zone_51N"
	case PCS_WGS84_UTM_zone_52N:
		return "PCS_WGS84_UTM_zone_52N"
	case PCS_WGS84_UTM_zone_53N:
		return "PCS_WGS84_UTM_zone_53N"
	case PCS_WGS84_UTM_zone_54N:
		return "PCS_WGS84_UTM_zone_54N"
	case PCS_WGS84_UTM_zone_55N:
		return "PCS_WGS84_UTM_zone_55N"
	case PCS_WGS84_UTM_zone_56N:
		return "PCS_WGS84_UTM_zone_56N"
	case PCS_WGS84_UTM_zone_57N:
		return "PCS_WGS84_UTM_zone_57N"
	case PCS_WGS84_UTM_zone_58N:
		return "PCS_WGS84_UTM_zone_58N"
	case PCS_WGS84_UTM_zone_59N:
		return "PCS_WGS84_UTM_zone_59N"
	case PCS_WGS84_UTM_zone_60N:
		return "PCS_WGS84_UTM_zone_60N"
	case PCS_WGS84_UTM_zone_1S:
		return "PCS_WGS84_UTM_zone_1S"
	case PCS_WGS84_UTM_zone_2S:
		return "PCS_WGS84_UTM_zone_2S"
	case PCS_WGS84_UTM_zone_3S:
		return "PCS_WGS84_UTM_zone_3S"
	case PCS_WGS84_UTM_zone_4S:
		return "PCS_WGS84_UTM_zone_4S"
	case PCS_WGS84_UTM_zone_5S:
		return "PCS_WGS84_UTM_zone_5S"
	case PCS_WGS84_UTM_zone_6S:
		return "PCS_WGS84_UTM_zone_6S"
	case PCS_WGS84_UTM_zone_7S:
		return "PCS_WGS84_UTM_zone_7S"
	case PCS_WGS84_UTM_zone_8S:
		return "PCS_WGS84_UTM_zone_8S"
	case PCS_WGS84_UTM_zone_9S:
		return "PCS_WGS84_UTM_zone_9S"
	case PCS_WGS84_UTM_zone_10S:
		return "PCS_WGS84_UTM_zone_10S"
	case PCS_WGS84_UTM_zone_11S:
		return "PCS_WGS84_UTM_zone_11S"
	case PCS_WGS84_UTM_zone_12S:
		return "PCS_WGS84_UTM_zone_12S"
	case PCS_WGS84_UTM_zone_13S:
		return "PCS_WGS84_UTM_zone_13S"
	case PCS_WGS84_UTM_zone_14S:
		return "PCS_WGS84_UTM_zone_14S"
	case PCS_WGS84_UTM_zone_15S:
		return "PCS_WGS84_UTM_zone_15S"
	case PCS_WGS84_UTM_zone_16S:
		return "PCS_WGS84_UTM_zone_16S"
	case PCS_WGS84_UTM_zone_17S:
		return "PCS_WGS84_UTM_zone_17S"
	case PCS_WGS84_UTM_zone_18S:
		return "PCS_WGS84_UTM_zone_18S"
	case PCS_WGS84_UTM_zone_19S:
		return "PCS_WGS84_UTM_zone_19S"
	case PCS_WGS84_UTM_zone_20S:
		return "PCS_WGS84_UTM_zone_20S"
	case PCS_WGS84_UTM_zone_21S:
		return "PCS_WGS84_UTM_zone_21S"
	case PCS_WGS84_UTM_zone_22S:
		return "PCS_WGS84_UTM_zone_22S"
	case PCS_WGS84_UTM_zone_23S:
		return "PCS_WGS84_UTM_zone_23S"
	case PCS_WGS84_UTM_zone_24S:
		return "PCS_WGS84_UTM_zone_24S"
	case PCS_WGS84_UTM_zone_25S:
		return "PCS_WGS84_UTM_zone_25S"
	case PCS_WGS84_UTM_zone_26S:
		return "PCS_WGS84_UTM_zone_26S"
	case PCS_WGS84_UTM_zone_27S:
		return "PCS_WGS84_UTM_zone_27S"
	case PCS_WGS84_UTM_zone_28S:
		return "PCS_WGS84_UTM_zone_28S"
	case PCS_WGS84_UTM_zone_29S:
		return "PCS_WGS84_UTM_zone_29S"
	case PCS_WGS84_UTM_zone_30S:
		return "PCS_WGS84_UTM_zone_30S"
	case PCS_WGS84_UTM_zone_31S:
		return "PCS_WGS84_UTM_zone_31S"
	case PCS_WGS84_UTM_zone_32S:
		return "PCS_WGS84_UTM_zone_32S"
	case PCS_WGS84_UTM_zone_33S:
		return "PCS_WGS84_UTM_zone_33S"
	case PCS_WGS84_UTM_zone_34S:
		return "PCS_WGS84_UTM_zone_34S"
	case PCS_WGS84_UTM_zone_35S:
		return "PCS_WGS84_UTM_zone_35S"
	case PCS_WGS84_UTM_zone_36S:
		return "PCS_WGS84_UTM_zone_36S"
	case PCS_WGS84_UTM_zone_37S:
		return "PCS_WGS84_UTM_zone_37S"
	case PCS_WGS84_UTM_zone_38S:
		return "PCS_WGS84_UTM_zone_38S"
	case PCS_WGS84_UTM_zone_39S:
		return "PCS_WGS84_UTM_zone_39S"
	case PCS_WGS84_UTM_zone_40S:
		return "PCS_WGS84_UTM_zone_40S"
	case PCS_WGS84_UTM_zone_41S:
		return "PCS_WGS84_UTM_zone_41S"
	case PCS_WGS84_UTM_zone_42S:
		return "PCS_WGS84_UTM_zone_42S"
	case PCS_WGS84_UTM_zone_43S:
		return "PCS_WGS84_UTM_zone_43S"
	case PCS_WGS84_UTM_zone_44S:
		return "PCS_WGS84_UTM_zone_44S"
	case PCS_WGS84_UTM_zone_45S:
		return "PCS_WGS84_UTM_zone_45S"
	case PCS_WGS84_UTM_zone_46S:
		return "PCS_WGS84_UTM_zone_46S"
	case PCS_WGS84_UTM_zone_47S:
		return "PCS_WGS84_UTM_zone_47S"
	case PCS_WGS84_UTM_zone_48S:
		return "PCS_WGS84_UTM_zone_48S"
	case PCS_WGS84_UTM_zone_49S:
		return "PCS_WGS84_UTM_zone_49S"
	case PCS_WGS84_UTM_zone_50S:
		return "PCS_WGS84_UTM_zone_50S"
	case PCS_WGS84_UTM_zone_51S:
		return "PCS_WGS84_UTM_zone_51S"
	case PCS_WGS84_UTM_zone_52S:
		return "PCS_WGS84_UTM_zone_52S"
	case PCS_WGS84_UTM_zone_53S:
		return "PCS_WGS84_UTM_zone_53S"
	case PCS_WGS84_UTM_zone_54S:
		return "PCS_WGS84_UTM_zone_54S"
	case PCS_WGS84_UTM_zone_55S:
		return "PCS_WGS84_UTM_zone_55S"
	case PCS_WGS84_UTM_zone_56S:
		return "PCS_WGS84_UTM_zone_56S"
	case PCS_WGS84_UTM_zone_57S:
		return "PCS_WGS84_UTM_zone_57S"
	case PCS_WGS84_UTM_zone_58S:
		return "PCS_WGS84_UTM_zone_58S"
	case PCS_WGS84_UTM_zone_59S:
		return "PCS_WGS84_UTM_zone_59S"
	case PCS_WGS84_UTM_zone_60S:
		return "PCS_WGS84_UTM_zone_60S"
	}
	return fmt.Sprintf("Unkown pcs (%d)", v)
}
