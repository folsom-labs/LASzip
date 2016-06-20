/*
===============================================================================

  FILE:  laszip_dll.h

  CONTENTS:

    A simple DLL interface to LASzip

  PROGRAMMERS:

    martin.isenburg@rapidlasso.com  -  http://rapidlasso.com

  COPYRIGHT:

    (c) 2007-2015, martin isenburg, rapidlasso - fast tools to catch reality

    This is free software; you can redistribute and/or modify it under the
    terms of the GNU Lesser General Licence as published by the Free Software
    Foundation. See the COPYING file for more information.

    This software is distributed WITHOUT ANY WARRANTY and without even the
    implied warranty of MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.

  CHANGE HISTORY:

    23 September 2015 -- correct update of bounding box and counters from inventory on closing
    22 September 2015 -- bug fix for not overwriting description of pre-existing "extra bytes"
    5 September 2015 -- "LAS 1.4 compatibility mode" now allows pre-existing "extra bytes"
    3 August 2015 -- incompatible DLL change for QSI-sponsored "LAS 1.4 compatibility mode"
    8 July 2015 -- adding support for NOAA-sponsored "LAS 1.4 compatibility mode"
    1 April 2015 -- adding exploitation and creation of spatial indexing information
    8 August 2013 -- added laszip_get_coordinates() and laszip_set_coordinates()
    6 August 2013 -- added laszip_auto_offset() and laszip_check_for_integer_overflow()
    31 July 2013 -- added laszip_get_point_count() for FUSION integration
    29 July 2013 -- reorganized to create an easy to use LASzip DLL

===============================================================================
*/

#ifndef LASZIP_DLL_H
#define LASZIP_DLL_H

#ifdef __cplusplus
extern "C"
{
#endif

/*---------------------------------------------------------------------------*/
/*--------------- DLL variables to pass data to/from LASzip -----------------*/
/*---------------------------------------------------------------------------*/
#include <stdint.h>

typedef int                laszip_BOOL;
typedef uint8_t            laszip_U8;
typedef uint16_t           laszip_U16;
typedef uint32_t           laszip_U32;
typedef uint64_t           laszip_U64;
typedef int8_t             laszip_I8;
typedef int16_t            laszip_I16;
typedef int32_t            laszip_I32;
typedef int64_t            laszip_I64;

typedef struct laszip_geokey
{
  laszip_U16 key_id;
  laszip_U16 tiff_tag_location;
  laszip_U16 count;
  laszip_U16 value_offset;
} laszip_geokey_struct;

typedef struct laszip_vlr
{
  laszip_U16 reserved;
  char user_id[16];
  laszip_U16 record_id;
  laszip_U16 record_length_after_header;
  char description[32];
  laszip_U8* data;
} laszip_vlr_struct;

typedef struct laszip_header
{
  laszip_U16 file_source_ID;
  laszip_U16 global_encoding;
  laszip_U32 project_ID_GUID_data_1;
  laszip_U16 project_ID_GUID_data_2;
  laszip_U16 project_ID_GUID_data_3;
  char project_ID_GUID_data_4[8];
  laszip_U8 version_major;
  laszip_U8 version_minor;
  char system_identifier[32];
  char generating_software[32];
  laszip_U16 file_creation_day;
  laszip_U16 file_creation_year;
  laszip_U16 header_size;
  laszip_U32 offset_to_point_data;
  laszip_U32 number_of_variable_length_records;
  laszip_U8 point_data_format;
  laszip_U16 point_data_record_length;
  laszip_U32 number_of_point_records;
  laszip_U32 number_of_points_by_return[5];
  double x_scale_factor;
  double y_scale_factor;
  double z_scale_factor;
  double x_offset;
  double y_offset;
  double z_offset;
  double max_x;
  double min_x;
  double max_y;
  double min_y;
  double max_z;
  double min_z;

  // LAS 1.3 and higher only
  laszip_U64 start_of_waveform_data_packet_record;

  // LAS 1.4 and higher only
  laszip_U64 start_of_first_extended_variable_length_record;
  laszip_U32 number_of_extended_variable_length_records;
  laszip_U64 extended_number_of_point_records;
  laszip_U64 extended_number_of_points_by_return[15];

  // optional
  laszip_U32 user_data_in_header_size;
  laszip_U8* user_data_in_header;

  // optional VLRs
  laszip_vlr_struct* vlrs;

  // optional
  laszip_U32 user_data_after_header_size;
  laszip_U8* user_data_after_header;

} laszip_header_struct;

typedef struct laszip_point
{
  laszip_I32 X;
  laszip_I32 Y;
  laszip_I32 Z;
  laszip_U16 intensity;
  laszip_U8 return_number : 3;
  laszip_U8 number_of_returns : 3;
  laszip_U8 scan_direction_flag : 1;
  laszip_U8 edge_of_flight_line : 1;
  laszip_U8 classification : 5;
  laszip_U8 synthetic_flag : 1;
  laszip_U8 keypoint_flag  : 1;
  laszip_U8 withheld_flag  : 1;
  laszip_I8 scan_angle_rank;
  laszip_U8 user_data;
  laszip_U16 point_source_ID;

  // LAS 1.4 only
  laszip_I16 extended_scan_angle;
  laszip_U8 extended_point_type : 2;
  laszip_U8 extended_scanner_channel : 2;
  laszip_U8 extended_classification_flags : 4;
  laszip_U8 extended_classification;
  laszip_U8 extended_return_number : 4;
  laszip_U8 extended_number_of_returns : 4;

  // for 8 byte alignment of the GPS time
  laszip_U8 dummy[7];

  double gps_time;
  laszip_U16 rgb[4];
  laszip_U8 wave_packet[29];

  laszip_I32 num_extra_bytes;
  laszip_U8* extra_bytes;

} laszip_point_struct;

/*---------------------------------------------------------------------------*/
/*---------------- DLL functions to manage the LASzip DLL -------------------*/
/*---------------------------------------------------------------------------*/

#define LASZIP_API

/*---------------------------------------------------------------------------*/
LASZIP_API laszip_I32
laszip_get_version
(
    laszip_U8*                         version_major
    , laszip_U8*                       version_minor
    , laszip_U16*                      version_revision
    , laszip_U32*                      version_build
);

/*---------------------------------------------------------------------------*/
LASZIP_API laszip_I32
laszip_create(
    void **                    pointer
);

/*---------------------------------------------------------------------------*/
LASZIP_API laszip_I32
laszip_get_error
(
    void *                     pointer
    , char**                    error
);

/*---------------------------------------------------------------------------*/
LASZIP_API laszip_I32
laszip_get_warning
(
    void *                     pointer
    , char**                    warning
);

/*---------------------------------------------------------------------------*/
LASZIP_API laszip_I32
laszip_clean(
    void *                     pointer
);

/*---------------------------------------------------------------------------*/
LASZIP_API laszip_I32
laszip_destroy(
    void *                     pointer
);

/*---------------------------------------------------------------------------*/
/*---------- DLL functions to write and read LAS and LAZ files --------------*/
/*---------------------------------------------------------------------------*/

/*---------------------------------------------------------------------------*/
LASZIP_API laszip_I32
laszip_get_header_pointer(
    void *                     pointer
    , laszip_header_struct**           header_pointer
);

/*---------------------------------------------------------------------------*/
LASZIP_API laszip_I32
laszip_get_point_pointer(
    void *                     pointer
    , laszip_point_struct**            point_pointer
);

/*---------------------------------------------------------------------------*/
LASZIP_API laszip_I32
laszip_get_point_count(
    void *                     pointer
    , laszip_I64*                      count
);

/*---------------------------------------------------------------------------*/
LASZIP_API laszip_I32
laszip_set_header(
    void *                     pointer
    , const laszip_header_struct*      header
);

/*---------------------------------------------------------------------------*/
LASZIP_API laszip_I32
laszip_check_for_integer_overflow(
    void *                     pointer
);

/*---------------------------------------------------------------------------*/
LASZIP_API laszip_I32
laszip_auto_offset(
    void *                     pointer
);

/*---------------------------------------------------------------------------*/
LASZIP_API laszip_I32
laszip_set_point(
    void *                     pointer
    , const laszip_point_struct*       point
);

/*---------------------------------------------------------------------------*/
LASZIP_API laszip_I32
laszip_set_coordinates(
    void *                     pointer
    , const double*                coordinates
);

/*---------------------------------------------------------------------------*/
LASZIP_API laszip_I32
laszip_get_coordinates(
    void *                     pointer
    , double*                      coordinates
);

/*---------------------------------------------------------------------------*/
LASZIP_API laszip_I32
laszip_set_geokeys(
    void *                     pointer
    , laszip_U32                       number
    , const laszip_geokey_struct*      key_entries
);

/*---------------------------------------------------------------------------*/
LASZIP_API laszip_I32
laszip_set_geodouble_params(
    void *                     pointer
    , laszip_U32                       number
    , const double*                geodouble_params
);

/*---------------------------------------------------------------------------*/
LASZIP_API laszip_I32
laszip_set_geoascii_params(
    void *                     pointer
    , laszip_U32                       number
    , const char*               geoascii_params
);

/*---------------------------------------------------------------------------*/
LASZIP_API laszip_I32
laszip_add_attribute(
    void *                     pointer
    , laszip_U32                       type
    , const char*               name
    , const char*               description
    , double                       scale
    , double                       offset
);

/*---------------------------------------------------------------------------*/
LASZIP_API laszip_I32
laszip_add_vlr(
    void *                     pointer
    , const char*               user_id
    , laszip_U16                       record_id
    , laszip_U16                       record_length_after_header
    , const char*               description
    , const laszip_U8*                 data
);

/*---------------------------------------------------------------------------*/
LASZIP_API laszip_I32
laszip_remove_vlr(
    void *                     pointer
    , const char*               user_id
    , laszip_U16                       record_id
);

/*---------------------------------------------------------------------------*/
LASZIP_API laszip_I32
laszip_create_spatial_index(
    void *                     pointer
    , const laszip_BOOL                create
    , const laszip_BOOL                append
);

/*---------------------------------------------------------------------------*/
LASZIP_API laszip_I32
laszip_preserve_generating_software(
    void *                     pointer
    , const laszip_BOOL                preserve
);

/*---------------------------------------------------------------------------*/
LASZIP_API laszip_I32
laszip_request_compatibility_mode(
    void *                     pointer
    , const laszip_BOOL                request
);

/*---------------------------------------------------------------------------*/
LASZIP_API laszip_I32
laszip_open_writer(
    void *                     pointer
    , const char*               file_name
    , laszip_BOOL                      compress
);

/*---------------------------------------------------------------------------*/
LASZIP_API laszip_I32
laszip_write_point(
    void *                     pointer
);

/*---------------------------------------------------------------------------*/
LASZIP_API laszip_I32
laszip_write_indexed_point(
    void *                     pointer
);

/*---------------------------------------------------------------------------*/
LASZIP_API laszip_I32
laszip_update_inventory(
    void *                     pointer
);

/*---------------------------------------------------------------------------*/
LASZIP_API laszip_I32
laszip_close_writer(
    void *                     pointer
);

/*---------------------------------------------------------------------------*/
LASZIP_API laszip_I32
laszip_exploit_spatial_index(
    void *                     pointer
    , const laszip_BOOL                exploit
);

/*---------------------------------------------------------------------------*/
LASZIP_API laszip_I32
laszip_open_reader(
    void *                     pointer
    , const char*               file_name
    , laszip_BOOL*                     is_compressed
);

/*---------------------------------------------------------------------------*/
LASZIP_API laszip_I32
laszip_has_spatial_index(
    void *                     pointer
    , laszip_BOOL*                     is_indexed
    , laszip_BOOL*                     is_appended
);

/*---------------------------------------------------------------------------*/
LASZIP_API laszip_I32
laszip_inside_rectangle(
    void *                     pointer
    , double                       min_x
    , double                       min_y
    , double                       max_x
    , double                       max_y
    , laszip_BOOL*                     is_empty
);

/*---------------------------------------------------------------------------*/
LASZIP_API laszip_I32
laszip_seek_point(
    void *                     pointer
    , laszip_I64                       index
);

/*---------------------------------------------------------------------------*/
LASZIP_API laszip_I32
laszip_read_point(
    void *                     pointer
);

/*---------------------------------------------------------------------------*/
LASZIP_API laszip_I32
laszip_read_inside_point(
    void *                     pointer
    , laszip_BOOL*                     is_done
);

/*---------------------------------------------------------------------------*/
LASZIP_API laszip_I32
laszip_close_reader(
    void *                     pointer
);

/*---------------------------------------------------------------------------*/
/*---------------- DLL functions to load and unload LASzip ------------------*/
/*---------------------------------------------------------------------------*/

/*---------------------------------------------------------------------------*/
LASZIP_API laszip_I32
laszip_load_dll
(
);

/*---------------------------------------------------------------------------*/
LASZIP_API laszip_I32
laszip_unload_dll
(
);

#ifdef __cplusplus
}
#endif

#endif /* LASZIP_DLL_H */
