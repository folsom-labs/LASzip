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

typedef struct laszip_geokey
{
  uint16_t key_id;
  uint16_t tiff_tag_location;
  uint16_t count;
  uint16_t value_offset;
} laszip_geokey_struct;

typedef struct laszip_vlr
{
  uint16_t reserved;
  char user_id[16];
  uint16_t record_id;
  uint16_t record_length_after_header;
  char description[32];
  uint8_t* data;
} laszip_vlr_struct;

typedef struct laszip_header
{
  uint16_t file_source_ID;
  uint16_t global_encoding;
  uint32_t project_ID_GUID_data_1;
  uint16_t project_ID_GUID_data_2;
  uint16_t project_ID_GUID_data_3;
  char project_ID_GUID_data_4[8];
  uint8_t version_major;
  uint8_t version_minor;
  char system_identifier[32];
  char generating_software[32];
  uint16_t file_creation_day;
  uint16_t file_creation_year;
  uint16_t header_size;
  uint32_t offset_to_point_data;
  uint32_t number_of_variable_length_records;
  uint8_t point_data_format;
  uint16_t point_data_record_length;
  uint32_t number_of_point_records;
  uint32_t number_of_points_by_return[5];
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
  uint64_t start_of_waveform_data_packet_record;

  // LAS 1.4 and higher only
  uint64_t start_of_first_extended_variable_length_record;
  uint32_t number_of_extended_variable_length_records;
  uint64_t extended_number_of_point_records;
  uint64_t extended_number_of_points_by_return[15];

  // optional
  uint32_t user_data_in_header_size;
  uint8_t* user_data_in_header;

  // optional VLRs
  laszip_vlr_struct* vlrs;

  // optional
  uint32_t user_data_after_header_size;
  uint8_t* user_data_after_header;

} laszip_header_struct;

typedef struct laszip_point
{
  int32_t X;
  int32_t Y;
  int32_t Z;
  uint16_t intensity;
  uint8_t return_number : 3;
  uint8_t number_of_returns : 3;
  uint8_t scan_direction_flag : 1;
  uint8_t edge_of_flight_line : 1;
  uint8_t classification : 5;
  uint8_t synthetic_flag : 1;
  uint8_t keypoint_flag  : 1;
  uint8_t withheld_flag  : 1;
  int8_t scan_angle_rank;
  uint8_t user_data;
  uint16_t point_source_ID;

  // LAS 1.4 only
  int16_t extended_scan_angle;
  uint8_t extended_point_type : 2;
  uint8_t extended_scanner_channel : 2;
  uint8_t extended_classification_flags : 4;
  uint8_t extended_classification;
  uint8_t extended_return_number : 4;
  uint8_t extended_number_of_returns : 4;

  // for 8 byte alignment of the GPS time
  uint8_t dummy[7];

  double gps_time;
  uint16_t rgb[4];
  uint8_t wave_packet[29];

  int32_t num_extra_bytes;
  uint8_t* extra_bytes;

} laszip_point_struct;

/*---------------------------------------------------------------------------*/
/*---------------- DLL functions to manage the LASzip DLL -------------------*/
/*---------------------------------------------------------------------------*/

#define LASZIP_API

/*---------------------------------------------------------------------------*/
LASZIP_API int32_t
laszip_create(
    void **                    pointer
);

/*---------------------------------------------------------------------------*/
LASZIP_API int32_t
laszip_get_error
(
    void *                     pointer
    , char**                    error
);

/*---------------------------------------------------------------------------*/
LASZIP_API int32_t
laszip_get_warning
(
    void *                     pointer
    , char**                    warning
);

/*---------------------------------------------------------------------------*/
LASZIP_API int32_t
laszip_clean(
    void *                     pointer
);

/*---------------------------------------------------------------------------*/
LASZIP_API int32_t
laszip_destroy(
    void *                     pointer
);

/*---------------------------------------------------------------------------*/
/*---------- DLL functions to write and read LAS and LAZ files --------------*/
/*---------------------------------------------------------------------------*/

/*---------------------------------------------------------------------------*/
LASZIP_API int32_t
laszip_get_header_pointer(
    void *                     pointer
    , laszip_header_struct**           header_pointer
);

/*---------------------------------------------------------------------------*/
LASZIP_API int32_t
laszip_get_point_pointer(
    void *                     pointer
    , laszip_point_struct**            point_pointer
);

/*---------------------------------------------------------------------------*/
LASZIP_API int32_t
laszip_get_point_count(
    void *                     pointer
    , uint64_t*                      count
);

/*---------------------------------------------------------------------------*/
LASZIP_API int32_t
laszip_check_for_integer_overflow(
    void *                     pointer
);

/*---------------------------------------------------------------------------*/
LASZIP_API int32_t
laszip_auto_offset(
    void *                     pointer
);

/*---------------------------------------------------------------------------*/
LASZIP_API int32_t
laszip_get_coordinates(
    void *                     pointer
    , double*                      coordinates
);

/*---------------------------------------------------------------------------*/
LASZIP_API int32_t
laszip_add_vlr(
    void *                     pointer
    , const char*               user_id
    , uint16_t                       record_id
    , uint16_t                       record_length_after_header
    , const char*               description
    , const uint8_t*                 data
);

/*---------------------------------------------------------------------------*/
LASZIP_API int32_t
laszip_remove_vlr(
    void *                     pointer
    , const char*               user_id
    , uint16_t                       record_id
);

/*---------------------------------------------------------------------------*/
LASZIP_API int32_t
laszip_create_spatial_index(
    void *                     pointer
    , const laszip_BOOL                create
    , const laszip_BOOL                append
);

/*---------------------------------------------------------------------------*/
LASZIP_API int32_t
laszip_preserve_generating_software(
    void *                     pointer
    , const laszip_BOOL                preserve
);

/*---------------------------------------------------------------------------*/
LASZIP_API int32_t
laszip_request_compatibility_mode(
    void *                     pointer
    , const laszip_BOOL                request
);

/*---------------------------------------------------------------------------*/
LASZIP_API int32_t
laszip_open_writer(
    void *                     pointer
    , const char*               file_name
    , laszip_BOOL                      compress
);

/*---------------------------------------------------------------------------*/
LASZIP_API int32_t
laszip_write_point(
    void *                     pointer
);

/*---------------------------------------------------------------------------*/
LASZIP_API int32_t
laszip_write_indexed_point(
    void *                     pointer
);

/*---------------------------------------------------------------------------*/
LASZIP_API int32_t
laszip_update_inventory(
    void *                     pointer
);

/*---------------------------------------------------------------------------*/
LASZIP_API int32_t
laszip_close_writer(
    void *                     pointer
);

/*---------------------------------------------------------------------------*/
LASZIP_API int32_t
laszip_exploit_spatial_index(
    void *                     pointer
    , const laszip_BOOL                exploit
);

/*---------------------------------------------------------------------------*/
LASZIP_API int32_t
laszip_open_reader(
    void *                     pointer
    , const char*               file_name
    , laszip_BOOL*                     is_compressed
);

/*---------------------------------------------------------------------------*/
LASZIP_API int32_t
laszip_has_spatial_index(
    void *                     pointer
    , laszip_BOOL*                     is_indexed
    , laszip_BOOL*                     is_appended
);

/*---------------------------------------------------------------------------*/
LASZIP_API int32_t
laszip_inside_rectangle(
    void *                     pointer
    , double                       min_x
    , double                       min_y
    , double                       max_x
    , double                       max_y
    , laszip_BOOL*                     is_empty
);

/*---------------------------------------------------------------------------*/
LASZIP_API int32_t
laszip_seek_point(
    void *                     pointer
    , uint64_t                       index
);

/*---------------------------------------------------------------------------*/
LASZIP_API int32_t
laszip_read_point(
    void *                     pointer
);

/*---------------------------------------------------------------------------*/
LASZIP_API int32_t
laszip_read_inside_point(
    void *                     pointer
    , laszip_BOOL*                     is_done
);

/*---------------------------------------------------------------------------*/
LASZIP_API int32_t
laszip_close_reader(
    void *                     pointer
);

/*---------------------------------------------------------------------------*/
/*---------------- DLL functions to load and unload LASzip ------------------*/
/*---------------------------------------------------------------------------*/

/*---------------------------------------------------------------------------*/
LASZIP_API int32_t
laszip_load_dll
(
);

/*---------------------------------------------------------------------------*/
LASZIP_API int32_t
laszip_unload_dll
(
);

#ifdef __cplusplus
}
#endif

#endif /* LASZIP_DLL_H */
