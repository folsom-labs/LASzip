#ifndef LASZIP_DLL_H
#define LASZIP_DLL_H

#ifdef __cplusplus
extern "C"
{
#endif

#include <stdio.h>
#include <string.h>
#include <stdlib.h>
#include <stdint.h>

#include "mydefs.hpp"

typedef int                laszip_BOOL;

class ByteStreamIn;
class LASreadPoint;
class LASattributer;
class LASindex;

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

typedef struct laszip_dll {
  laszip_header_struct header;
  I64 p_count;
  I64 npoints;
  laszip_point_struct point;
  U8** point_items;
  FILE* file;
  ByteStreamIn* streamin;
  LASreadPoint* reader;
  LASattributer* attributer;
  CHAR error[1024];
  CHAR warning[1024];
  F64 lax_r_min_x;
  F64 lax_r_min_y;
  F64 lax_r_max_x;
  F64 lax_r_max_y;
  BOOL preserve_generating_software;
  I32 start_scan_angle;
  I32 start_extended_returns;
  I32 start_classification;
  I32 start_flags_and_channel;
  I32 start_NIR_band;
} laszip_dll_struct;

int32_t laszip_create(laszip_dll_struct **pointer);
int32_t laszip_get_error(laszip_dll_struct *pointer, char** error);
int32_t laszip_get_warning(laszip_dll_struct *pointer, char**warning);
int32_t laszip_clean(laszip_dll_struct *pointer);
int32_t laszip_destroy(laszip_dll_struct *pointer);
int32_t laszip_get_header_pointer(laszip_dll_struct *pointer, laszip_header_struct **header_pointer);
int32_t laszip_get_point_pointer(laszip_dll_struct *pointer, laszip_point_struct **point_pointer);
int32_t laszip_get_point_count(laszip_dll_struct *pointer, uint64_t *count);
int32_t laszip_check_for_integer_overflow(laszip_dll_struct *pointer);
int32_t laszip_auto_offset(laszip_dll_struct *pointer);

int32_t laszip_preserve_generating_software(laszip_dll_struct *pointer, const laszip_BOOL preserve);
int32_t laszip_open_reader(laszip_dll_struct *pointer, const char *file_name, laszip_BOOL *is_compressed);

int32_t laszip_read_point(laszip_dll_struct *pointer);
int32_t laszip_close_reader(laszip_dll_struct *pointer);

#ifdef __cplusplus
}
#endif

#endif /* LASZIP_DLL_H */
