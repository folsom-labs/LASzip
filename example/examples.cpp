
#include <time.h>
#include <stdio.h>
#include <stdlib.h>
#include <string.h>

#include "laszip_api.h"
#include "mydefs.hpp"

static void dll_error(void * laszip)
{
  if (laszip)
  {
    char* error;
    if (laszip_get_error(laszip, &error))
    {
      fprintf(stderr,"DLL ERROR: getting error messages\n");
    }
    fprintf(stderr,"DLL ERROR MESSAGE: %s\n", error);
  }
}

static void byebye(bool error=false, bool wait=false, void * laszip=0)
{
  if (error)
  {
    dll_error(laszip);
  }
  if (wait)
  {
    fprintf(stderr,"<press ENTER>\n");
    getc(stdin);
  }
  exit(error);
}

static double taketime()
{
  return (double)(clock())/CLOCKS_PER_SEC;
}

void example1(int argc, char* file_name_in, char* file_name_out)
{
  fprintf(stderr,"running EXAMPLE_ONE (reading *without* and writing *without* compatibility mode)\n");

  double start_time = taketime();
  void * laszip_reader;
  if (laszip_create(&laszip_reader))
  {
    fprintf(stderr,"DLL ERROR: creating laszip reader\n");
    byebye(true, argc==1);
  }

  laszip_BOOL is_compressed = 0;
  if (laszip_open_reader(laszip_reader, file_name_in, &is_compressed))
  {
    fprintf(stderr,"DLL ERROR: opening laszip reader for '%s'\n", file_name_in);
    byebye(true, argc==1, laszip_reader);
  }

  fprintf(stderr,"file '%s' is %scompressed\n", file_name_in, (is_compressed ? "" : "un"));

  // get a pointer to the header of the reader that was just populated

  laszip_header* header;

  if (laszip_get_header_pointer(laszip_reader, &header))
  {
    fprintf(stderr,"DLL ERROR: getting header pointer from laszip reader\n");
    byebye(true, argc==1, laszip_reader);
  }

  // how many points does the file have

  uint64_t npoints = (header->number_of_point_records ? header->number_of_point_records : header->extended_number_of_point_records);

  // report how many points the file has

  fprintf(stderr,"file '%s' contains %lld points\n", file_name_in, npoints);
  //fprintf(stderr,"file '%s' contains %I64d points\n", file_name_in, npoints);

  // get a pointer to the points that will be read

  laszip_point* point;

  if (laszip_get_point_pointer(laszip_reader, &point))
  {
    fprintf(stderr,"DLL ERROR: getting point pointer from laszip reader\n");
    byebye(true, argc==1, laszip_reader);
  }

  void * laszip_writer;
  if (laszip_create(&laszip_writer))
  {
    fprintf(stderr,"DLL ERROR: creating laszip writer\n");
    byebye(true, argc==1);
  }

  // initialize the header for the writer using the header of the reader

  if (laszip_set_header(laszip_writer, header))
  {
    fprintf(stderr,"DLL ERROR: setting header for laszip writer\n");
    byebye(true, argc==1, laszip_writer);
  }

  laszip_BOOL compress = (strstr(file_name_out, ".laz") != 0);

  if (laszip_open_writer(laszip_writer, file_name_out, compress))
  {
    fprintf(stderr,"DLL ERROR: opening laszip writer for '%s'\n", file_name_out);
    byebye(true, argc==1, laszip_writer);
  }

  fprintf(stderr,"writing file '%s' %scompressed\n", file_name_out, (compress ? "" : "un"));

  // read the points

  uint64_t p_count = 0;

  while (p_count < npoints)
  {
    // read a point

    if (laszip_read_point(laszip_reader))
    {
      fprintf(stderr,"DLL ERROR: reading point %lld\n", p_count);
      //fprintf(stderr,"DLL ERROR: reading point %" PRId64 "\n", p_count);
      byebye(true, argc==1, laszip_reader);
    }

    // copy the point

    if (laszip_set_point(laszip_writer, point))
    {
      fprintf(stderr,"DLL ERROR: setting point %lld\n", p_count);
      //fprintf(stderr,"DLL ERROR: setting point %I64d\n", p_count);
      byebye(true, argc==1, laszip_writer);
    }

    // write the point

    if (laszip_write_point(laszip_writer))
    {
      fprintf(stderr,"DLL ERROR: writing point %lld\n", p_count);
      //fprintf(stderr,"DLL ERROR: writing point %I64d\n", p_count);
      byebye(true, argc==1, laszip_writer);
    }

    p_count++;
  }

  fprintf(stderr,"successfully read and written %lld points\n", p_count);
  //fprintf(stderr,"successfully read and written %I64d points\n", p_count);

  if (laszip_close_writer(laszip_writer))
  {
    fprintf(stderr,"DLL ERROR: closing laszip writer\n");
    byebye(true, argc==1, laszip_writer);
  }

  if (laszip_destroy(laszip_writer))
  {
    fprintf(stderr,"DLL ERROR: destroying laszip writer\n");
    byebye(true, argc==1);
  }

  if (laszip_close_reader(laszip_reader))
  {
    fprintf(stderr,"DLL ERROR: closing laszip reader\n");
    byebye(true, argc==1, laszip_reader);
  }

  if (laszip_destroy(laszip_reader))
  {
    fprintf(stderr,"DLL ERROR: destroying laszip reader\n");
    byebye(true, argc==1);
  }

  fprintf(stderr,"total time: %g sec for reading %scompressed and writing %scompressed\n", taketime()-start_time, (is_compressed ? "" : "un"), (compress ? "" : "un"));
}

void example2(int argc, char* file_name_in, char* file_name_out)
{
  uint32_t i;
  fprintf(stderr,"running EXAMPLE_TWO (another way of reading *without* and writing *without* compatibility mode)\n");

  double start_time = taketime();

  void * laszip_reader;
  if (laszip_create(&laszip_reader))
  {
    fprintf(stderr,"DLL ERROR: creating laszip reader\n");
    byebye(true, argc==1);
  }

  laszip_BOOL is_compressed = 0;
  if (laszip_open_reader(laszip_reader, file_name_in, &is_compressed))
  {
    fprintf(stderr,"DLL ERROR: opening laszip reader for '%s'\n", file_name_in);
    byebye(true, argc==1, laszip_reader);
  }

  fprintf(stderr,"file '%s' is %scompressed\n", file_name_in, (is_compressed ? "" : "un"));

  // get a pointer to the header of the reader that was just populated

  laszip_header* header_read;

  if (laszip_get_header_pointer(laszip_reader, &header_read))
  {
    fprintf(stderr,"DLL ERROR: getting header pointer from laszip reader\n");
    byebye(true, argc==1, laszip_reader);
  }

  // how many points does the file have

  uint64_t npoints = (header_read->number_of_point_records ? header_read->number_of_point_records : header_read->extended_number_of_point_records);

  // report how many points the file has

  fprintf(stderr,"file '%s' contains %lld points\n", file_name_in, npoints);
  //fprintf(stderr,"file '%s' contains %I64d points\n", file_name_in, npoints);

  void * laszip_writer;
  if (laszip_create(&laszip_writer))
  {
    fprintf(stderr,"DLL ERROR: creating laszip writer\n");
    byebye(true, argc==1);
  }

  // get a pointer to the header of the writer so we can populate it

  laszip_header* header_write;

  if (laszip_get_header_pointer(laszip_writer, &header_write))
  {
    fprintf(stderr,"DLL ERROR: getting header pointer from laszip writer\n");
    byebye(true, argc==1, laszip_writer);
  }

  // copy entries from the reader header to the writer header

  header_write->file_source_ID = header_read->file_source_ID;
  header_write->global_encoding = header_read->global_encoding;
  header_write->project_ID_GUID_data_1 = header_read->project_ID_GUID_data_1;
  header_write->project_ID_GUID_data_2 = header_read->project_ID_GUID_data_2;
  header_write->project_ID_GUID_data_3 = header_read->project_ID_GUID_data_3;
  memcpy(header_write->project_ID_GUID_data_4, header_read->project_ID_GUID_data_4, 8);
  header_write->version_major = header_read->version_major;
  header_write->version_minor = header_read->version_minor;
  memcpy(header_write->system_identifier, header_read->system_identifier, 32);
  memcpy(header_write->generating_software, header_read->generating_software, 32);
  header_write->file_creation_day = header_read->file_creation_day;
  header_write->file_creation_year = header_read->file_creation_year;
  header_write->header_size = header_read->header_size;
  header_write->offset_to_point_data = header_read->header_size; /* note !!! */
  header_write->number_of_variable_length_records = header_read->number_of_variable_length_records;
  header_write->point_data_format = header_read->point_data_format;
  header_write->point_data_record_length = header_read->point_data_record_length;
  header_write->number_of_point_records = header_read->number_of_point_records;
  for (i = 0; i < 5; i++)
  {
    header_write->number_of_points_by_return[i] = header_read->number_of_points_by_return[i];
  }
  header_write->x_scale_factor = header_read->x_scale_factor;
  header_write->y_scale_factor = header_read->y_scale_factor;
  header_write->z_scale_factor = header_read->z_scale_factor;
  header_write->x_offset = header_read->x_offset;
  header_write->y_offset = header_read->y_offset;
  header_write->z_offset = header_read->z_offset;
  header_write->max_x = header_read->max_x;
  header_write->min_x = header_read->min_x;
  header_write->max_y = header_read->max_y;
  header_write->min_y = header_read->min_y;
  header_write->max_z = header_read->max_z;
  header_write->min_z = header_read->min_z;

  // LAS 1.3 and higher only
  header_write->start_of_waveform_data_packet_record = header_read->start_of_waveform_data_packet_record;

  // LAS 1.4 and higher only
  header_write->start_of_first_extended_variable_length_record = header_read->start_of_first_extended_variable_length_record;
  header_write->number_of_extended_variable_length_records = header_read->number_of_extended_variable_length_records;
  header_write->extended_number_of_point_records = header_read->extended_number_of_point_records;
  for (i = 0; i < 15; i++)
  {
    header_write->extended_number_of_points_by_return[i] = header_read->extended_number_of_points_by_return[i];
  }

  // we may modify output because we omit any user defined data that may be ** the header

  if (header_read->user_data_in_header_size)
  {
    header_write->header_size -= header_read->user_data_in_header_size;
    header_write->offset_to_point_data -= header_read->user_data_in_header_size;
    fprintf(stderr,"omitting %d bytes of user_data_in_header\n", header_read->user_data_after_header_size);
  }

  // add all the VLRs

  if (header_read->number_of_variable_length_records)
  {
    fprintf(stderr,"offset_to_point_data before adding %u VLRs is      : %d\n", header_read->number_of_variable_length_records, (int32_t)header_write->offset_to_point_data);
    for (i = 0; i < header_read->number_of_variable_length_records; i++)
    {
      if (laszip_add_vlr(laszip_writer, header_read->vlrs[i].user_id, header_read->vlrs[i].record_id, header_read->vlrs[i].record_length_after_header, header_read->vlrs[i].description, header_read->vlrs[i].data))
      {
        fprintf(stderr,"DLL ERROR: adding VLR %u of %u to the header of the laszip writer\n", i+i, header_read->number_of_variable_length_records);
        byebye(true, argc==1, laszip_writer);
      }
      fprintf(stderr,"offset_to_point_data after adding VLR number %u is : %d\n", i+1, (int32_t)header_write->offset_to_point_data);
    }
  }

  // we may modify output because we omit any user defined data that may be *after* the header

  if (header_read->user_data_after_header_size)
  {
    fprintf(stderr,"omitting %d bytes of user_data_after_header\n", header_read->user_data_after_header_size);
  }

  laszip_BOOL compress = (strstr(file_name_out, ".laz") != 0);

  if (laszip_open_writer(laszip_writer, file_name_out, compress))
  {
    fprintf(stderr,"DLL ERROR: opening laszip writer for '%s'\n", file_name_out);
    byebye(true, argc==1, laszip_writer);
  }

  fprintf(stderr,"writing file '%s' %scompressed\n", file_name_out, (compress ? "" : "un"));

  // get a pointer to the point of the reader will be read

  laszip_point* point_read;

  if (laszip_get_point_pointer(laszip_reader, &point_read))
  {
    fprintf(stderr,"DLL ERROR: getting point pointer from laszip reader\n");
    byebye(true, argc==1, laszip_reader);
  }

  // get a pointer to the point of the writer that we will populate and write

  laszip_point* point_write;

  if (laszip_get_point_pointer(laszip_writer, &point_write))
  {
    fprintf(stderr,"DLL ERROR: getting point pointer from laszip writer\n");
    byebye(true, argc==1, laszip_writer);
  }

  // read the points

  uint64_t p_count = 0;

  while (p_count < npoints)
  {
    // read a point

    if (laszip_read_point(laszip_reader))
    {
      fprintf(stderr,"DLL ERROR: reading point %lld\n", p_count);
      //fprintf(stderr,"DLL ERROR: reading point %I64d\n", p_count);
      byebye(true, argc==1, laszip_reader);
    }

    // copy the point

    point_write->X = point_read->X;
    point_write->Y = point_read->Y;
    point_write->Z = point_read->Z;
    point_write->intensity = point_read->intensity;
    point_write->return_number = point_read->return_number;
    point_write->number_of_returns = point_read->number_of_returns;
    point_write->scan_direction_flag = point_read->scan_direction_flag;
    point_write->edge_of_flight_line = point_read->edge_of_flight_line;
    point_write->classification = point_read->classification;
    point_write->withheld_flag = point_read->withheld_flag;
    point_write->keypoint_flag = point_read->keypoint_flag;
    point_write->synthetic_flag = point_read->synthetic_flag;
    point_write->scan_angle_rank = point_read->scan_angle_rank;
    point_write->user_data = point_read->user_data;
    point_write->point_source_ID = point_read->point_source_ID;

    point_write->gps_time = point_read->gps_time;
    memcpy(point_write->rgb, point_read->rgb, 8);
    memcpy(point_write->wave_packet, point_read->wave_packet, 29);

    // LAS 1.4 only
    point_write->extended_point_type = point_read->extended_point_type;
    point_write->extended_scanner_channel = point_read->extended_scanner_channel;
    point_write->extended_classification_flags = point_read->extended_classification_flags;
    point_write->extended_classification = point_read->extended_classification;
    point_write->extended_return_number = point_read->extended_return_number;
    point_write->extended_number_of_returns = point_read->extended_number_of_returns;
    point_write->extended_scan_angle = point_read->extended_scan_angle;

    if (point_read->num_extra_bytes)
    {
      memcpy(point_write->extra_bytes, point_read->extra_bytes, point_read->num_extra_bytes);
    }

    // write the point

    if (laszip_write_point(laszip_writer))
    {
      fprintf(stderr,"DLL ERROR: writing point %lld\n", p_count);
      //fprintf(stderr,"DLL ERROR: writing point %I64d\n", p_count);
      byebye(true, argc==1, laszip_writer);
    }

    p_count++;
  }

  fprintf(stderr,"successfully read and written %lld points\n", p_count);
  //fprintf(stderr,"successfully read and written %I64d points\n", p_count);

  if (laszip_close_writer(laszip_writer))
  {
    fprintf(stderr,"DLL ERROR: closing laszip writer\n");
    byebye(true, argc==1, laszip_writer);
  }

  if (laszip_destroy(laszip_writer))
  {
    fprintf(stderr,"DLL ERROR: destroying laszip writer\n");
    byebye(true, argc==1);
  }

  if (laszip_close_reader(laszip_reader))
  {
    fprintf(stderr,"DLL ERROR: closing laszip reader\n");
    byebye(true, argc==1, laszip_reader);
  }

  if (laszip_destroy(laszip_reader))
  {
    fprintf(stderr,"DLL ERROR: destroying laszip reader\n");
    byebye(true, argc==1);
  }

  fprintf(stderr,"total time: %g sec for reading %scompressed and writing %scompressed\n", taketime()-start_time, (is_compressed ? "" : "un"), (compress ? "" : "un"));
}

void example4(int argc, char* file_name_in, char* file_name_out)
{
  fprintf(stderr,"running EXAMPLE_FOUR (reading area-of-interest from a file exploiting possibly existing spatial indexing information)\n");

  double start_time = taketime();
  void * laszip_reader;
  if (laszip_create(&laszip_reader))
  {
    fprintf(stderr,"DLL ERROR: creating laszip reader\n");
    byebye(true, argc==1);
  }

  // signal that spatial queries are coming

  laszip_BOOL exploit = 1;
  if (laszip_exploit_spatial_index(laszip_reader, exploit))
  {
    fprintf(stderr,"DLL ERROR: signaling laszip reader that spatial queries are coming for '%s'\n", file_name_in);
    byebye(true, argc==1, laszip_reader);
  }

  laszip_BOOL is_compressed = 0;
  if (laszip_open_reader(laszip_reader, file_name_in, &is_compressed))
  {
    fprintf(stderr,"DLL ERROR: opening laszip reader for '%s'\n", file_name_in);
    byebye(true, argc==1, laszip_reader);
  }

  fprintf(stderr,"file '%s' is %scompressed\n", file_name_in, (is_compressed ? "" : "un"));

  // check whether spatial indexing information is available

  laszip_BOOL is_indexed = 0;
  laszip_BOOL is_appended = 0;
  if (laszip_has_spatial_index(laszip_reader, &is_indexed, &is_appended))
  {
    fprintf(stderr,"DLL ERROR: checking laszip reader whether spatial indexing information is present for '%s'\n", file_name_in);
    byebye(true, argc==1, laszip_reader);
  }

  fprintf(stderr,"file '%s' does %shave spatial indexing information\n", file_name_in, (is_indexed ? "" : "not "));

  laszip_header* header;

  if (laszip_get_header_pointer(laszip_reader, &header))
  {
    fprintf(stderr,"DLL ERROR: getting header pointer from laszip reader\n");
    byebye(true, argc==1, laszip_reader);
  }

  // how many points does the file have

  uint64_t npoints = (header->number_of_point_records ? header->number_of_point_records : header->extended_number_of_point_records);

  // report how many points the file has

  fprintf(stderr,"file '%s' contains %lld points\n", file_name_in, npoints);
  //fprintf(stderr,"file '%s' contains %I64d points\n", file_name_in, npoints);

  // create a rectangular box enclosing a subset of points at the center of the full bounding box

  const double sub = 0.05;

  double mid_x = (header->min_x + header->max_x) / 2;
  double mid_y = (header->min_y + header->max_y) / 2;

  double range_x = header->max_x - header->min_x;
  double range_y = header->max_y - header->min_y;

  double sub_min_x = mid_x - sub * range_x;
  double sub_min_y = mid_y - sub * range_y;

  double sub_max_x = mid_x + sub * range_x;
  double sub_max_y = mid_y + sub * range_y;

  // request the reader to only read this specified rectangular subset of points
  laszip_BOOL is_empty = 0;
  if (laszip_inside_rectangle(laszip_reader, sub_min_x, sub_min_y, sub_max_x, sub_max_y, &is_empty))
  {
    fprintf(stderr,"DLL ERROR: requesting points inside of rectangle [%g,%g] (%g,%g) from laszip reader\n", sub_min_x, sub_min_y, sub_max_x, sub_max_y);
    byebye(true, argc==1, laszip_reader);
  }

  laszip_point* point;

  if (laszip_get_point_pointer(laszip_reader, &point))
  {
    fprintf(stderr,"DLL ERROR: getting point pointer from laszip reader\n");
    byebye(true, argc==1, laszip_reader);
  }

  void * laszip_writer;
  if (laszip_create(&laszip_writer))
  {
    fprintf(stderr,"DLL ERROR: creating laszip writer\n");
    byebye(true, argc==1);
  }

  if (laszip_set_header(laszip_writer, header))
  {
    fprintf(stderr,"DLL ERROR: setting header for laszip writer\n");
    byebye(true, argc==1, laszip_writer);
  }

  laszip_BOOL compress = (strstr(file_name_out, ".laz") != 0);

  if (laszip_open_writer(laszip_writer, file_name_out, compress))
  {
    fprintf(stderr,"DLL ERROR: opening laszip writer for '%s'\n", file_name_out);
    byebye(true, argc==1, laszip_writer);
  }

  fprintf(stderr,"writing file '%s' %scompressed\n", file_name_out, (compress ? "" : "un"));

  laszip_BOOL is_done = 0;
  uint64_t p_count = 0;

  while (p_count < npoints)
  {
    if (laszip_read_inside_point(laszip_reader, &is_done))
    {
      fprintf(stderr,"DLL ERROR: reading point %lld\n", p_count);
      //fprintf(stderr,"DLL ERROR: reading point %I64d\n", p_count);
      byebye(true, argc==1, laszip_reader);
    }

    if (is_done)
    {
      break;
    }

    if (laszip_set_point(laszip_writer, point))
    {
      fprintf(stderr,"DLL ERROR: setting point %lld\n", p_count);
      //fprintf(stderr,"DLL ERROR: setting point %I64d\n", p_count);
      byebye(true, argc==1, laszip_writer);
    }

    if (laszip_write_point(laszip_writer))
    {
      fprintf(stderr,"DLL ERROR: writing point %lld\n", p_count);
      //fprintf(stderr,"DLL ERROR: writing point %I64d\n", p_count);
      byebye(true, argc==1, laszip_writer);
    }

    if (laszip_update_inventory(laszip_writer))
    {
      fprintf(stderr,"DLL ERROR: updating inventory for point %lld\n", p_count);
      //fprintf(stderr,"DLL ERROR: updating inventory for point %I64d\n", p_count);
      byebye(true, argc==1, laszip_writer);
    }

    p_count++;
  }

  fprintf(stderr,"successfully read and written %lld points\n", p_count);
  //fprintf(stderr,"successfully read and written %I64d points\n", p_count);

  if (laszip_close_writer(laszip_writer))
  {
    fprintf(stderr,"DLL ERROR: closing laszip writer\n");
    byebye(true, argc==1, laszip_writer);
  }

  if (laszip_destroy(laszip_writer))
  {
    fprintf(stderr,"DLL ERROR: destroying laszip writer\n");
    byebye(true, argc==1);
  }

  if (laszip_close_reader(laszip_reader))
  {
    fprintf(stderr,"DLL ERROR: closing laszip reader\n");
    byebye(true, argc==1, laszip_reader);
  }

  if (laszip_destroy(laszip_reader))
  {
    fprintf(stderr,"DLL ERROR: destroying laszip reader\n");
    byebye(true, argc==1);
  }

  fprintf(stderr,"total time: %g sec for reading %scompressed and writing %scompressed\n", taketime()-start_time, (is_compressed ? "" : "un"), (compress ? "" : "un"));
}

void example5(int argc, char* file_name_in, char* file_name_out)
{
	fprintf(stderr,"running EXAMPLE_FIVE (reading from one file and writing to another file while simultaneously generating a spatial index)\n");

  double start_time = taketime();
	void * laszip_reader;
	if (laszip_create(&laszip_reader))
	{
	  fprintf(stderr,"DLL ERROR: creating laszip reader\n");
	  byebye(true, argc==1);
	}

	laszip_BOOL is_compressed = 0;
	if (laszip_open_reader(laszip_reader, file_name_in, &is_compressed))
	{
	  fprintf(stderr,"DLL ERROR: opening laszip reader for '%s'\n", file_name_in);
	  byebye(true, argc==1, laszip_reader);
	}

	fprintf(stderr,"file '%s' is %scompressed\n", file_name_in, (is_compressed ? "" : "un"));

	// get a pointer to the header of the reader that was just populated

	laszip_header* header;

	if (laszip_get_header_pointer(laszip_reader, &header))
	{
	  fprintf(stderr,"DLL ERROR: getting header pointer from laszip reader\n");
	  byebye(true, argc==1, laszip_reader);
	}

	// how many points does the file have

	uint64_t npoints = (header->number_of_point_records ? header->number_of_point_records : header->extended_number_of_point_records);

	// report how many points the file has

	fprintf(stderr,"file '%s' contains %lld points\n", file_name_in, npoints);
	//fprintf(stderr,"file '%s' contains %I64d points\n", file_name_in, npoints);

	// get a pointer to the points that will be read

	laszip_point* point;

	if (laszip_get_point_pointer(laszip_reader, &point))
	{
	  fprintf(stderr,"DLL ERROR: getting point pointer from laszip reader\n");
	  byebye(true, argc==1, laszip_reader);
	}

	void * laszip_writer;
	if (laszip_create(&laszip_writer))
	{
	  fprintf(stderr,"DLL ERROR: creating laszip writer\n");
	  byebye(true, argc==1);
	}

	// initialize the header for the writer using the header of the reader

	if (laszip_set_header(laszip_writer, header))
	{
	  fprintf(stderr,"DLL ERROR: setting header for laszip writer\n");
	  byebye(true, argc==1, laszip_writer);
	}

	// enable the creation of spatial indices

	laszip_BOOL create = 1;
	laszip_BOOL append = 0; /* not supported yet */

	if (laszip_create_spatial_index(laszip_writer, create, append))
	{
	  fprintf(stderr,"DLL ERROR: signaling laszip writer to create spatial indexing information\n");
	  byebye(true, argc==1, laszip_writer);
	}

	laszip_BOOL compress = (strstr(file_name_out, ".laz") != 0);

	if (laszip_open_writer(laszip_writer, file_name_out, compress))
	{
	  fprintf(stderr,"DLL ERROR: opening laszip writer for '%s'\n", file_name_out);
	  byebye(true, argc==1, laszip_writer);
	}

	fprintf(stderr,"writing file '%s' spatially indexed and %scompressed\n", file_name_out, (compress ? "" : "un"));

	// read the points

	uint64_t p_count = 0;

	while (p_count < npoints)
	{
	  // read a point

	  if (laszip_read_point(laszip_reader))
	  {
	    fprintf(stderr,"DLL ERROR: reading point %lld\n", p_count);
	    //fprintf(stderr,"DLL ERROR: reading point %I64d\n", p_count);
	    byebye(true, argc==1, laszip_reader);
	  }

	  // copy the point

	  if (laszip_set_point(laszip_writer, point))
	  {
	    fprintf(stderr,"DLL ERROR: setting point %lld\n", p_count);
	    //fprintf(stderr,"DLL ERROR: setting point %I64d\n", p_count);
	    byebye(true, argc==1, laszip_writer);
	  }

	  // write the point

	  if (laszip_write_indexed_point(laszip_writer))
	  {
	    fprintf(stderr,"DLL ERROR: writing indexed point %lld\n", p_count);
	    //fprintf(stderr,"DLL ERROR: writing indexed point %I64d\n", p_count);
	    byebye(true, argc==1, laszip_writer);
	  }

	  p_count++;
	}

	fprintf(stderr,"successfully read and written %lld indexed points\n", p_count);
	//fprintf(stderr,"successfully read and written %I64d indexed points\n", p_count);

	if (laszip_close_writer(laszip_writer))
	{
	  fprintf(stderr,"DLL ERROR: closing laszip writer\n");
	  byebye(true, argc==1, laszip_writer);
	}

	if (laszip_destroy(laszip_writer))
	{
	  fprintf(stderr,"DLL ERROR: destroying laszip writer\n");
	  byebye(true, argc==1);
	}

	if (laszip_close_reader(laszip_reader))
	{
	  fprintf(stderr,"DLL ERROR: closing laszip reader\n");
	  byebye(true, argc==1, laszip_reader);
	}

	if (laszip_destroy(laszip_reader))
	{
	  fprintf(stderr,"DLL ERROR: destroying laszip reader\n");
	  byebye(true, argc==1);
	}

	fprintf(stderr,"total time: %g sec for reading %scompressed and writing indexed & %scompressed\n", taketime()-start_time, (is_compressed ? "" : "un"), (compress ? "" : "un"));
}

void example6(int argc, char* file_name_in, char* file_name_out)
{
	  int i;
    fprintf(stderr,"running EXAMPLE_SIX (writing five points of type 6 to LAS 1.4 without compatibility)\n");

    double start_time = taketime();
    void * laszip_writer;
    if (laszip_create(&laszip_writer))
    {
      fprintf(stderr,"DLL ERROR: creating laszip writer\n");
      byebye(true, argc==1);
    }

    // get a pointer to the header of the writer so we can populate it

    laszip_header* header;

    if (laszip_get_header_pointer(laszip_writer, &header))
    {
      fprintf(stderr,"DLL ERROR: getting header pointer from laszip writer\n");
      byebye(true, argc==1, laszip_writer);
    }

    // populate the header

    header->file_source_ID = 4711;
    header->global_encoding = (1<<0) | (1<<4);     // see LAS specification for details
    header->version_major = 1;
    header->version_minor = 4;
    strncpy(header->system_identifier, "LASzip DLL example 6", 32);
    header->file_creation_day = 30;
    header->file_creation_year = 2015;
    header->header_size = 375;
    header->offset_to_point_data = 375;
    header->point_data_format = 6;
    header->point_data_record_length = 30;
    header->number_of_point_records = 0;           // legacy 32-bit counters should be zero for new point types > 5
    for (i = 0; i < 5; i++)
    {
      header->number_of_points_by_return[i] = 0;
    }
    header->extended_number_of_point_records = 5;
    header->extended_number_of_points_by_return[0] = 1;
    header->extended_number_of_points_by_return[1] = 2;
    header->extended_number_of_points_by_return[7] = 1;
    header->extended_number_of_points_by_return[8] = 1;
    header->max_x = 630499.95;
    header->min_x = 630498.56;
    header->max_y = 4834749.66;
    header->min_y = 4834748.73;
    header->max_z = 63.68;
    header->min_z = 61.33;

    // optional: use the bounding box and the scale factor to create a "good" offset

    if (laszip_auto_offset(laszip_writer))
    {
      fprintf(stderr,"DLL ERROR: during automatic offset creation\n");
      byebye(true, argc==1, laszip_writer);
    }

    fprintf(stderr,"offset_to_point_data before adding  empty OGC WKT VLR is : %d\n", (int32_t)header->offset_to_point_data);

    // add intentionally empty OGC WKT

    if (laszip_add_vlr(laszip_writer, "LASF_Projection", 2112, 0, "intentionally empty OGC WKT", 0))
    {
      fprintf(stderr,"DLL ERROR: adding intentionally empty OGC WKT VLR to the header\n");
      byebye(true, argc==1, laszip_writer);
    }

    fprintf(stderr,"offset_to_point_data before adding funny VLR is          : %d\n", (int32_t)header->offset_to_point_data);

    // add some funny VLR

    if (laszip_add_vlr(laszip_writer, "funny", 12345, 0, "just a funny VLR", 0))
    {
      fprintf(stderr,"DLL ERROR: adding funny VLR to the header\n");
      byebye(true, argc==1, laszip_writer);
    }

    fprintf(stderr,"offset_to_point_data after adding VLRs                   : %d\n", (int32_t)header->offset_to_point_data);

    laszip_BOOL compress = (strstr(file_name_out, ".laz") != 0);

    // this should fail if compress is true

    if (laszip_open_writer(laszip_writer, file_name_out, compress))
    {
      fprintf(stderr,"DLL ERROR: opening laszip writer for '%s'\n", file_name_out);
      byebye(true, argc==1, laszip_writer);
    }

    fprintf(stderr,"writing file '%s' %scompressed\n", file_name_out, (compress ? "" : "un"));

    // get a pointer to the point of the writer that we will populate and write

    laszip_point* point;

    if (laszip_get_point_pointer(laszip_writer, &point))
    {
      fprintf(stderr,"DLL ERROR: getting point pointer from laszip writer\n");
      byebye(true, argc==1, laszip_writer);
    }

    // set points to "extended type"

    point->extended_point_type = 1;

    // write five points

    uint64_t p_count = 0;
    double coordinates[3];

    // populate the first point

    coordinates[0] = 630499.95;
    coordinates[1] = 4834749.17;
    coordinates[2] = 62.15;

    if (laszip_set_coordinates(laszip_writer, coordinates))
    {
      fprintf(stderr,"DLL ERROR: setting coordinates for point %lld\n", p_count);
      //fprintf(stderr,"DLL ERROR: setting coordinates for point %I64d\n", p_count);
      byebye(true, argc==1, laszip_writer);
    }

    point->intensity = 60;
    point->extended_return_number = 2;
    point->extended_number_of_returns = 2;
    point->classification = 2;                // it must be set because it "fits" in 5 bits
    point->extended_classification = 2;
    point->extended_scan_angle = 3500;
    point->extended_scanner_channel = 1;
    point->extended_classification_flags = 8; // overflag flag is set
    point->gps_time = 53413162.560400;

    // write the first point

    if (laszip_write_point(laszip_writer))
    {
      fprintf(stderr,"DLL ERROR: writing point %lld\n", p_count);
      //fprintf(stderr,"DLL ERROR: writing point %I64d\n", p_count);
      byebye(true, argc==1, laszip_writer);
    }
    p_count++;

    // populate the second point

    coordinates[0] = 630499.83;
    coordinates[1] = 4834748.88;
    coordinates[2] = 62.68;

    if (laszip_set_coordinates(laszip_writer, coordinates))
    {
      fprintf(stderr,"DLL ERROR: setting coordinates for point %lld\n", p_count);
      //fprintf(stderr,"DLL ERROR: setting coordinates for point %I64d\n", p_count);
      byebye(true, argc==1, laszip_writer);
    }

    point->intensity = 90;
    point->extended_return_number = 8;
    point->extended_number_of_returns = 9;
    point->classification = 0;                // it must be set to zero as the real value is stored in the extended field
    point->extended_classification = 41;
    point->extended_scan_angle = 3567;
    point->extended_scanner_channel = 1;
    point->extended_classification_flags = 0; // no flag is not set
    point->gps_time = 53413162.563600;

    // write the second point

    if (laszip_write_point(laszip_writer))
    {
      fprintf(stderr,"DLL ERROR: writing point %lld\n", p_count);
      //fprintf(stderr,"DLL ERROR: writing point %I64d\n", p_count);
      byebye(true, argc==1, laszip_writer);
    }
    p_count++;

    // populate the third point

    coordinates[0] = 630499.54;
    coordinates[1] = 4834749.66;
    coordinates[2] = 62.66;

    if (laszip_set_coordinates(laszip_writer, coordinates))
    {
      fprintf(stderr,"DLL ERROR: setting coordinates for point %lld\n", p_count);
      //fprintf(stderr,"DLL ERROR: setting coordinates for point %I64d\n", p_count);
      byebye(true, argc==1, laszip_writer);
    }

    point->intensity = 70;
    point->extended_return_number = 9;
    point->extended_number_of_returns = 9;
    point->classification = 0;                // it must be set to zero as the real value is stored in the extended field
    point->extended_classification = 42;
    point->extended_scan_angle = 3633;
    point->extended_scanner_channel = 1;
    point->extended_classification_flags = 0; // no flag is not set
    point->gps_time = 53413162.566800;

    // write the third point

    if (laszip_write_point(laszip_writer))
    {
      fprintf(stderr,"DLL ERROR: writing point %lld\n", p_count);
      //fprintf(stderr,"DLL ERROR: writing point %I64d\n", p_count);
      byebye(true, argc==1, laszip_writer);
    }
    p_count++;

    // populate the fourth point

    coordinates[0] = 630498.56;
    coordinates[1] = 4834749.41;
    coordinates[2] = 63.68;

    if (laszip_set_coordinates(laszip_writer, coordinates))
    {
      fprintf(stderr,"DLL ERROR: setting coordinates for point %lld\n", p_count);
      //fprintf(stderr,"DLL ERROR: setting coordinates for point %I64d\n", p_count);
      byebye(true, argc==1, laszip_writer);
    }

    point->intensity = 20;
    point->extended_return_number = 1;
    point->extended_number_of_returns = 2;
    point->classification = 5;                // it must be set because it "fits" in 5 bits
    point->extended_classification = 5;
    point->extended_scan_angle = 3700;
    point->extended_scanner_channel = 1;
    point->extended_classification_flags = 0; // no flag is not set
    point->gps_time = 53413162.580200;

    // write the fourth point

    if (laszip_write_point(laszip_writer))
    {
      fprintf(stderr,"DLL ERROR: writing point %lld\n", p_count);
      //fprintf(stderr,"DLL ERROR: writing point %I64d\n", p_count);
      byebye(true, argc==1, laszip_writer);
    }
    p_count++;

    // populate the fifth point

    coordinates[0] = 630498.80;
    coordinates[1] = 4834748.73;
    coordinates[2] = 62.16;

    if (laszip_set_coordinates(laszip_writer, coordinates))
    {
      fprintf(stderr,"DLL ERROR: setting coordinates for point %lld\n", p_count);
      //fprintf(stderr,"DLL ERROR: setting coordinates for point %I64d\n", p_count);
      byebye(true, argc==1, laszip_writer);
    }

    point->intensity = 110;
    point->extended_return_number = 2;
    point->extended_number_of_returns = 2;
    point->classification = 2;                // it must be set because it "fits" in 5 bits
    point->extended_classification = 2;
    point->extended_scan_angle = 3767;
    point->extended_scanner_channel = 1;
    point->extended_classification_flags = 0; // no flag is not set
    point->gps_time = 53413162.580200;

    // write the fifth point

    if (laszip_write_point(laszip_writer))
    {
      fprintf(stderr,"DLL ERROR: writing point %lld\n", p_count);
      //fprintf(stderr,"DLL ERROR: writing point %I64d\n", p_count);
      byebye(true, argc==1, laszip_writer);
    }
    p_count++;

    // get the number of points written so far

    if (laszip_get_point_count(laszip_writer, &p_count))
    {
      fprintf(stderr,"DLL ERROR: getting point count\n");
      byebye(true, argc==1, laszip_writer);
    }

    fprintf(stderr,"successfully written %lld points\n", p_count);
    //fprintf(stderr,"successfully written %I64d points\n", p_count);

    if (laszip_close_writer(laszip_writer))
    {
      fprintf(stderr,"DLL ERROR: closing laszip writer\n");
      byebye(true, argc==1, laszip_writer);
    }

    if (laszip_destroy(laszip_writer))
    {
      fprintf(stderr,"DLL ERROR: destroying laszip writer\n");
      byebye(true, argc==1);
    }

    fprintf(stderr,"total time: %g sec for writing %scompressed\n", taketime()-start_time, (compress ? "" : "un"));
  } // end of EXAMPLE_SIX

void example7(int argc, char* file_name_in, char* file_name_out)
{
  	int i;
    fprintf(stderr,"running EXAMPLE_SEVEN (writing five points of type 6 to LAS 1.4 *with* compatibility to compressed LAZ *and* also uncompressed LAS)\n");

    double start_time = taketime();
    void * laszip_writer;
    if (laszip_create(&laszip_writer))
    {
      fprintf(stderr,"DLL ERROR: creating laszip writer\n");
      byebye(true, argc==1);
    }

    // get a pointer to the header of the writer so we can populate it

    laszip_header* header;

    if (laszip_get_header_pointer(laszip_writer, &header))
    {
      fprintf(stderr,"DLL ERROR: getting header pointer from laszip writer\n");
      byebye(true, argc==1, laszip_writer);
    }

    // populate the header

    header->file_source_ID = 4711;
    header->global_encoding = (1<<0) | (1<<4);     // see LAS specification for details
    header->version_major = 1;
    header->version_minor = 4;
    strncpy(header->system_identifier, "LASzip DLL example 7", 32);
    header->file_creation_day = 30;
    header->file_creation_year = 2015;
    header->header_size = 375;
    header->offset_to_point_data = 375;
    header->point_data_format = 6;
    header->point_data_record_length = 30;
    header->number_of_point_records = 0;           // legacy 32-bit counters should be zero for new point types > 5
    for (i = 0; i < 5; i++)
    {
      header->number_of_points_by_return[i] = 0;
    }
    header->extended_number_of_point_records = 5;
    header->extended_number_of_points_by_return[0] = 1;
    header->extended_number_of_points_by_return[1] = 2;
    header->extended_number_of_points_by_return[7] = 1;
    header->extended_number_of_points_by_return[8] = 1;
    header->max_x = 630499.95;
    header->min_x = 630498.56;
    header->max_y = 4834749.66;
    header->min_y = 4834748.73;
    header->max_z = 63.68;
    header->min_z = 61.33;

    // optional: use the bounding box and the scale factor to create a "good" offset

    if (laszip_auto_offset(laszip_writer))
    {
      fprintf(stderr,"DLL ERROR: during automatic offset creation\n");
      byebye(true, argc==1, laszip_writer);
    }

    fprintf(stderr,"offset_to_point_data before adding  empty OGC WKT VLR is : %d\n", (int32_t)header->offset_to_point_data);

    // add intentionally empty OGC WKT

    if (laszip_add_vlr(laszip_writer, "LASF_Projection", 2112, 0, "intentionally empty OGC WKT", 0))
    {
      fprintf(stderr,"DLL ERROR: adding intentionally empty OGC WKT VLR to the header\n");
      byebye(true, argc==1, laszip_writer);
    }

    fprintf(stderr,"offset_to_point_data before adding funny VLR is          : %d\n", (int32_t)header->offset_to_point_data);

    // add some funny VLR

    if (laszip_add_vlr(laszip_writer, "funny", 12345, 0, "just a funny VLR", 0))
    {
      fprintf(stderr,"DLL ERROR: adding funny VLR to the header\n");
      byebye(true, argc==1, laszip_writer);
    }

    fprintf(stderr,"offset_to_point_data after adding VLRs                   : %d\n", (int32_t)header->offset_to_point_data);

    // enable the compatibility mode

    laszip_BOOL request = 1;
    if (laszip_request_compatibility_mode(laszip_writer, request))
    {
      fprintf(stderr,"DLL ERROR: enabling laszip LAS 1.4 compatibility mode\n");
      byebye(true, argc==1, laszip_writer);
    }

    laszip_BOOL compress = (strstr(file_name_out, ".laz") != 0);

    if (laszip_open_writer(laszip_writer, file_name_out, compress))
    {
      fprintf(stderr,"DLL ERROR: opening laszip writer for '%s'\n", file_name_out);
      byebye(true, argc==1, laszip_writer);
    }

    fprintf(stderr,"writing file '%s' %scompressed\n", file_name_out, (compress ? "" : "un"));

    // get a pointer to the point of the writer that we will populate and write

    laszip_point* point;

    if (laszip_get_point_pointer(laszip_writer, &point))
    {
      fprintf(stderr,"DLL ERROR: getting point pointer from laszip writer\n");
      byebye(true, argc==1, laszip_writer);
    }

    // set points to "extended type"

    point->extended_point_type = 1;

    // write five points

    uint64_t p_count = 0;
    double coordinates[3];

    // populate the first point

    coordinates[0] = 630499.95;
    coordinates[1] = 4834749.17;
    coordinates[2] = 62.15;

    if (laszip_set_coordinates(laszip_writer, coordinates))
    {
      fprintf(stderr,"DLL ERROR: setting coordinates for point %lld\n", p_count);
      //fprintf(stderr,"DLL ERROR: setting coordinates for point %I64d\n", p_count);
      byebye(true, argc==1, laszip_writer);
    }

    point->intensity = 60;
    point->extended_return_number = 2;
    point->extended_number_of_returns = 2;
    point->classification = 2;                // it must be set because it "fits" in 5 bits
    point->extended_classification = 2;
    point->extended_scan_angle = 3500;
    point->extended_scanner_channel = 1;
    point->extended_classification_flags = 8; // overflag flag is set
    point->gps_time = 53413162.560400;

    // write the first point

    if (laszip_write_point(laszip_writer))
    {
      fprintf(stderr,"DLL ERROR: writing point %lld\n", p_count);
      //fprintf(stderr,"DLL ERROR: writing point %I64d\n", p_count);
      byebye(true, argc==1, laszip_writer);
    }
    p_count++;

    // populate the second point

    coordinates[0] = 630499.83;
    coordinates[1] = 4834748.88;
    coordinates[2] = 62.68;

    if (laszip_set_coordinates(laszip_writer, coordinates))
    {
      fprintf(stderr,"DLL ERROR: setting coordinates for point %lld\n", p_count);
      //fprintf(stderr,"DLL ERROR: setting coordinates for point %I64d\n", p_count);
      byebye(true, argc==1, laszip_writer);
    }

    point->intensity = 90;
    point->extended_return_number = 8;
    point->extended_number_of_returns = 9;
    point->classification = 0;                // it must be set to zero as the real value is stored in the extended field
    point->extended_classification = 41;
    point->extended_scan_angle = 3567;
    point->extended_scanner_channel = 1;
    point->extended_classification_flags = 0; // no flag is not set
    point->gps_time = 53413162.563600;

    // write the second point

    if (laszip_write_point(laszip_writer))
    {
      fprintf(stderr,"DLL ERROR: writing point %lld\n", p_count);
      //fprintf(stderr,"DLL ERROR: writing point %I64d\n", p_count);
      byebye(true, argc==1, laszip_writer);
    }
    p_count++;

    // populate the third point

    coordinates[0] = 630499.54;
    coordinates[1] = 4834749.66;
    coordinates[2] = 62.66;

    if (laszip_set_coordinates(laszip_writer, coordinates))
    {
      fprintf(stderr,"DLL ERROR: setting coordinates for point %lld\n", p_count);
      //fprintf(stderr,"DLL ERROR: setting coordinates for point %I64d\n", p_count);
      byebye(true, argc==1, laszip_writer);
    }

    point->intensity = 70;
    point->extended_return_number = 9;
    point->extended_number_of_returns = 9;
    point->classification = 0;                // it must be set to zero as the real value is stored in the extended field
    point->extended_classification = 42;
    point->extended_scan_angle = 3633;
    point->extended_scanner_channel = 1;
    point->extended_classification_flags = 0; // no flag is not set
    point->gps_time = 53413162.566800;

    // write the third point

    if (laszip_write_point(laszip_writer))
    {
      fprintf(stderr,"DLL ERROR: writing point %lld\n", p_count);
      //fprintf(stderr,"DLL ERROR: writing point %I64d\n", p_count);
      byebye(true, argc==1, laszip_writer);
    }
    p_count++;

    // populate the fourth point

    coordinates[0] = 630498.56;
    coordinates[1] = 4834749.41;
    coordinates[2] = 63.68;

    if (laszip_set_coordinates(laszip_writer, coordinates))
    {
      fprintf(stderr,"DLL ERROR: setting coordinates for point %lld\n", p_count);
      //fprintf(stderr,"DLL ERROR: setting coordinates for point %I64d\n", p_count);
      byebye(true, argc==1, laszip_writer);
    }

    point->intensity = 20;
    point->extended_return_number = 1;
    point->extended_number_of_returns = 2;
    point->classification = 5;                // it must be set because it "fits" in 5 bits
    point->extended_classification = 5;
    point->extended_scan_angle = 3700;
    point->extended_scanner_channel = 1;
    point->extended_classification_flags = 0; // no flag is not set
    point->gps_time = 53413162.580200;

    // write the fourth point

    if (laszip_write_point(laszip_writer))
    {
      fprintf(stderr,"DLL ERROR: writing point %lld\n", p_count);
      //fprintf(stderr,"DLL ERROR: writing point %I64d\n", p_count);
      byebye(true, argc==1, laszip_writer);
    }
    p_count++;

    // populate the fifth point

    coordinates[0] = 630498.80;
    coordinates[1] = 4834748.73;
    coordinates[2] = 62.16;

    if (laszip_set_coordinates(laszip_writer, coordinates))
    {
      fprintf(stderr,"DLL ERROR: setting coordinates for point %lld\n", p_count);
      //fprintf(stderr,"DLL ERROR: setting coordinates for point %I64d\n", p_count);
      byebye(true, argc==1, laszip_writer);
    }

    point->intensity = 110;
    point->extended_return_number = 2;
    point->extended_number_of_returns = 2;
    point->classification = 2;                // it must be set because it "fits" in 5 bits
    point->extended_classification = 2;
    point->extended_scan_angle = 3767;
    point->extended_scanner_channel = 1;
    point->extended_classification_flags = 0; // no flag is not set
    point->gps_time = 53413162.580200;

    // write the fifth point

    if (laszip_write_point(laszip_writer))
    {
      fprintf(stderr,"DLL ERROR: writing point %lld\n", p_count);
      //fprintf(stderr,"DLL ERROR: writing point %I64d\n", p_count);
      byebye(true, argc==1, laszip_writer);
    }
    p_count++;

    // get the number of points written so far

    if (laszip_get_point_count(laszip_writer, &p_count))
    {
      fprintf(stderr,"DLL ERROR: getting point count\n");
      byebye(true, argc==1, laszip_writer);
    }

    fprintf(stderr,"successfully written %lld points\n", p_count);
    //fprintf(stderr,"successfully written %I64d points\n", p_count);

    if (laszip_close_writer(laszip_writer))
    {
      fprintf(stderr,"DLL ERROR: closing laszip writer\n");
      byebye(true, argc==1, laszip_writer);
    }

    if (laszip_destroy(laszip_writer))
    {
      fprintf(stderr,"DLL ERROR: destroying laszip writer\n");
      byebye(true, argc==1);
    }

    fprintf(stderr,"total time: %g sec for writing %scompressed\n", taketime()-start_time, (compress ? "" : "un"));
} // end of EXAMPLE_SEVEN

void example8(int argc, char* file_name_in, char* file_name_out)
{
    fprintf(stderr,"running EXAMPLE_EIGHT (always *with* compatibility mode when reading but when writing *only* for compressed output)\n");

    double start_time = taketime();
    void * laszip_reader;
    if (laszip_create(&laszip_reader))
    {
      fprintf(stderr,"DLL ERROR: creating laszip reader\n");
      byebye(true, argc==1);
    }

    // enable the compatibility mode for the reader

    laszip_BOOL request_reader = 1;
    if (laszip_request_compatibility_mode(laszip_reader, request_reader))
    {
      fprintf(stderr,"DLL ERROR: enabling laszip LAS 1.4 compatibility mode for the reader\n");
      byebye(true, argc==1, laszip_reader);
    }

    laszip_BOOL is_compressed = 0;
    if (laszip_open_reader(laszip_reader, file_name_in, &is_compressed))
    {
      fprintf(stderr,"DLL ERROR: opening laszip reader for '%s'\n", file_name_in);
      byebye(true, argc==1, laszip_reader);
    }

    fprintf(stderr,"file '%s' is %scompressed\n", file_name_in, (is_compressed ? "" : "un"));

    // get a pointer to the header of the reader that was just populated

    laszip_header* header;

    if (laszip_get_header_pointer(laszip_reader, &header))
    {
      fprintf(stderr,"DLL ERROR: getting header pointer from laszip reader\n");
      byebye(true, argc==1, laszip_reader);
    }

    // how many points does the file have

    uint64_t npoints = (header->number_of_point_records ? header->number_of_point_records : header->extended_number_of_point_records);

    // report how many points the file has

    fprintf(stderr,"file '%s' contains %lld points\n", file_name_in, npoints);
    //fprintf(stderr,"file '%s' contains %I64d points\n", file_name_in, npoints);

    // get a pointer to the points that will be read

    laszip_point* point;

    if (laszip_get_point_pointer(laszip_reader, &point))
    {
      fprintf(stderr,"DLL ERROR: getting point pointer from laszip reader\n");
      byebye(true, argc==1, laszip_reader);
    }

    void * laszip_writer;
    if (laszip_create(&laszip_writer))
    {
      fprintf(stderr,"DLL ERROR: creating laszip writer\n");
      byebye(true, argc==1);
    }

    // check if the output is compressed

    laszip_BOOL compress = (strstr(file_name_out, ".laz") != 0);

    // *only* enable the compatibility mode for the writer for compressed output

    if (compress)
    {
      laszip_BOOL request_writer = 1;
      if (laszip_request_compatibility_mode(laszip_writer, request_writer))
      {
        fprintf(stderr,"DLL ERROR: enabling laszip LAS 1.4 compatibility mode for the writer\n");
        byebye(true, argc==1, laszip_writer);
      }
    }

    // initialize the header for the writer using the header of the reader

    if (laszip_set_header(laszip_writer, header))
    {
      fprintf(stderr,"DLL ERROR: setting header for laszip writer\n");
      byebye(true, argc==1, laszip_writer);
    }

    if (laszip_open_writer(laszip_writer, file_name_out, compress))
    {
      fprintf(stderr,"DLL ERROR: opening laszip writer for '%s'\n", file_name_out);
      byebye(true, argc==1, laszip_writer);
    }

    fprintf(stderr,"writing file '%s' %scompressed\n", file_name_out, (compress ? "" : "un"));

    // read the points

    uint64_t p_count = 0;

    while (p_count < npoints)
    {
      // read a point

      if (laszip_read_point(laszip_reader))
      {
        fprintf(stderr,"DLL ERROR: reading point %lld\n", p_count);
        //fprintf(stderr,"DLL ERROR: reading point %I64\n", p_count);
        byebye(true, argc==1, laszip_reader);
      }

      // copy the point

      if (laszip_set_point(laszip_writer, point))
      {
        fprintf(stderr,"DLL ERROR: setting point %lld\n", p_count);
        //fprintf(stderr,"DLL ERROR: setting point %I64\n", p_count);
        byebye(true, argc==1, laszip_writer);
      }

      if (laszip_write_point(laszip_writer))
      {
        fprintf(stderr,"DLL ERROR: writing point %lld\n", p_count);
        //fprintf(stderr,"DLL ERROR: writing point %I64\n", p_count);
        byebye(true, argc==1, laszip_writer);
      }

      p_count++;
    }

    fprintf(stderr,"successfully read and written %lld points\n", p_count);
    //fprintf(stderr,"successfully read and written %I64d points\n", p_count);

    if (laszip_close_writer(laszip_writer))
    {
      fprintf(stderr,"DLL ERROR: closing laszip writer\n");
      byebye(true, argc==1, laszip_writer);
    }

    if (laszip_destroy(laszip_writer))
    {
      fprintf(stderr,"DLL ERROR: destroying laszip writer\n");
      byebye(true, argc==1);
    }

    if (laszip_close_reader(laszip_reader))
    {
      fprintf(stderr,"DLL ERROR: closing laszip reader\n");
      byebye(true, argc==1, laszip_reader);
    }

    if (laszip_destroy(laszip_reader))
    {
      fprintf(stderr,"DLL ERROR: destroying laszip reader\n");
      byebye(true, argc==1);
    }

    fprintf(stderr,"total time: %g sec for reading %scompressed and writing %scompressed\n", taketime()-start_time, (is_compressed ? "" : "un"), (compress ? "" : "un"));
} // end of EXAMPLE_EIGHT

void example10(int argc, char* file_name_in, char* file_name_out)
{
	  int i;
    fprintf(stderr,"running EXAMPLE_TEN (read LAS 1.0-1.3 file, upconvert old to new point types, write LAS 1.4 compatibility mode *only* for compressed output)\n");
    double start_time = taketime();

    void * laszip_reader;
    if (laszip_create(&laszip_reader))
    {
      fprintf(stderr,"DLL ERROR: creating laszip reader\n");
      byebye(true, argc==1);
    }

    // enable the compatibility mode for the reader

    laszip_BOOL request_reader = 1;
    if (laszip_request_compatibility_mode(laszip_reader, request_reader))
    {
      fprintf(stderr,"DLL ERROR: enabling laszip LAS 1.4 compatibility mode for the reader\n");
      byebye(true, argc==1, laszip_reader);
    }

    laszip_BOOL is_compressed = 0;
    if (laszip_open_reader(laszip_reader, file_name_in, &is_compressed))
    {
      fprintf(stderr,"DLL ERROR: opening laszip reader for '%s'\n", file_name_in);
      byebye(true, argc==1, laszip_reader);
    }

    fprintf(stderr,"file '%s' is %scompressed\n", file_name_in, (is_compressed ? "" : "un"));

    // get a pointer to the header of the reader that was just populated

    laszip_header* header_read;

    if (laszip_get_header_pointer(laszip_reader, &header_read))
    {
      fprintf(stderr,"DLL ERROR: getting header pointer from laszip reader\n");
      byebye(true, argc==1, laszip_reader);
    }

    // make sure it is LAS 1.0, LAS 1.1, LAS 1.2, or LAS 1.3

    if (header_read->version_minor > 3)
    {
      fprintf(stderr,"USER ERROR: input should be LAS 1.0 to LAS 1.3\n");
      byebye(true, argc==1);
    }

    // how many points does the LAS 1.x (x < 4) file have

    uint64_t npoints = header_read->number_of_point_records;

    // report how many points the file has

    fprintf(stderr,"file '%s' contains %lld points\n", file_name_in, npoints);
    //fprintf(stderr,"file '%s' contains %I64d points\n", file_name_in, npoints);

    // get a pointer to the points that will be read

    laszip_point* point_read;

    if (laszip_get_point_pointer(laszip_reader, &point_read))
    {
      fprintf(stderr,"DLL ERROR: getting point pointer from laszip reader\n");
      byebye(true, argc==1, laszip_reader);
    }

    void * laszip_writer;
    if (laszip_create(&laszip_writer))
    {
      fprintf(stderr,"DLL ERROR: creating laszip writer\n");
      byebye(true, argc==1);
    }

    // check if the output is compressed

    laszip_BOOL compress = (strstr(file_name_out, ".laz") != 0);

    // *only* enable the compatibility mode for the writer for compressed output

    if (compress)
    {
      laszip_BOOL request_writer = 1;
      if (laszip_request_compatibility_mode(laszip_writer, request_writer))
      {
        fprintf(stderr,"DLL ERROR: enabling laszip LAS 1.4 compatibility mode for the writer\n");
        byebye(true, argc==1, laszip_writer);
      }
    }

    // get a pointer to the header of the writer so we can populate it

    laszip_header* header_write;

    if (laszip_get_header_pointer(laszip_writer, &header_write))
    {
      fprintf(stderr,"DLL ERROR: getting header pointer from laszip writer\n");
      byebye(true, argc==1, laszip_writer);
    }

    // copy entries from the reader header to the writer header

    header_write->file_source_ID = header_read->file_source_ID;
    header_write->global_encoding = header_read->global_encoding;
    header_write->project_ID_GUID_data_1 = header_read->project_ID_GUID_data_1;
    header_write->project_ID_GUID_data_2 = header_read->project_ID_GUID_data_2;
    header_write->project_ID_GUID_data_3 = header_read->project_ID_GUID_data_3;
    memcpy(header_write->project_ID_GUID_data_4, header_read->project_ID_GUID_data_4, 8);
    header_write->version_major = header_read->version_major;
    header_write->version_minor = header_read->version_minor;
    memcpy(header_write->system_identifier, header_read->system_identifier, 32);
    memcpy(header_write->generating_software, header_read->generating_software, 32);
    header_write->file_creation_day = header_read->file_creation_day;
    header_write->file_creation_year = header_read->file_creation_year;
    header_write->header_size = header_read->header_size;
    header_write->offset_to_point_data = header_read->header_size; /* note !!! */
    header_write->number_of_variable_length_records = header_read->number_of_variable_length_records;
    header_write->point_data_format = header_read->point_data_format;
    header_write->point_data_record_length = header_read->point_data_record_length;
    header_write->number_of_point_records = header_read->number_of_point_records;
    for (i = 0; i < 5; i++)
    {
      header_write->number_of_points_by_return[i] = header_read->number_of_points_by_return[i];
    }
    header_write->x_scale_factor = header_read->x_scale_factor;
    header_write->y_scale_factor = header_read->y_scale_factor;
    header_write->z_scale_factor = header_read->z_scale_factor;
    header_write->x_offset = header_read->x_offset;
    header_write->y_offset = header_read->y_offset;
    header_write->z_offset = header_read->z_offset;
    header_write->max_x = header_read->max_x;
    header_write->min_x = header_read->min_x;
    header_write->max_y = header_read->max_y;
    header_write->min_y = header_read->min_y;
    header_write->max_z = header_read->max_z;
    header_write->min_z = header_read->min_z;

    // LAS 1.3 and higher only
    header_write->start_of_waveform_data_packet_record = header_read->start_of_waveform_data_packet_record;

    // we may modify output because we omit any user defined data that may be *before* the header

    header_write->user_data_in_header_size = 0;
    if (header_read->user_data_in_header_size)
    {
      header_write->header_size -= header_read->user_data_in_header_size;
      header_write->offset_to_point_data -= header_read->user_data_in_header_size;
      fprintf(stderr,"omitting %d bytes of user_data_in_header\n", header_read->user_data_after_header_size);
    }

    // add all the VLRs

    if (header_read->number_of_variable_length_records)
    {
      fprintf(stderr,"offset_to_point_data before adding %u VLRs is      : %d\n", header_read->number_of_variable_length_records, (int32_t)header_write->offset_to_point_data);
      for (U32 i = 0; i < header_read->number_of_variable_length_records; i++)
      {
        if (laszip_add_vlr(laszip_writer, header_read->vlrs[i].user_id, header_read->vlrs[i].record_id, header_read->vlrs[i].record_length_after_header, header_read->vlrs[i].description, header_read->vlrs[i].data))
        {
          fprintf(stderr,"DLL ERROR: adding VLR %u of %u to the header of the laszip writer\n", i+i, header_read->number_of_variable_length_records);
          byebye(true, argc==1, laszip_writer);
        }
        fprintf(stderr,"offset_to_point_data after adding VLR number %u is : %d\n", i+1, (int32_t)header_write->offset_to_point_data);
      }
    }

    // we may modify output because we omit any user defined data that may be *after* the header

    header_write->user_data_after_header_size = 0;
    if (header_read->user_data_after_header_size)
    {
      fprintf(stderr,"omitting %d bytes of user_data_after_header\n", header_read->user_data_after_header_size);
    }

    // upgrade the header and the points to LAS 1.4

    header_write->version_minor = 4;
    if (header_read->version_minor == 3)
    {
      header_write->header_size += 140;
      header_write->offset_to_point_data += 140;
    }
    else
    {
      header_write->header_size += 148;
      header_write->offset_to_point_data += 148;
    }

    if (header_read->point_data_format == 0)
    {
      header_write->point_data_format = 6;
      header_write->point_data_record_length += 10;
    }
    else if (header_read->point_data_format == 1)
    {
      header_write->point_data_format = 6;
      header_write->point_data_record_length += 2;
    }
    else if (header_read->point_data_format == 2)
    {
      header_write->point_data_format = 7;
      header_write->point_data_record_length += 10;
    }
    else if (header_read->point_data_format == 3)
    {
      header_write->point_data_format = 7;
      header_write->point_data_record_length += 2;
    }
    else
    {
      fprintf(stderr,"USER ERROR: input point type should be 0 to 3 and not %d\n", header_read->point_data_format);
      byebye(true, argc==1);
    }

    // we do not add EVLRs
    header_write->start_of_first_extended_variable_length_record = 0;
    header_write->number_of_extended_variable_length_records = 0;

    // zero the legacy counters
    header_write->number_of_point_records = 0;
    for (i = 0; i < 5; i++)
    {
      header_write->number_of_points_by_return[i] = 0;
    }

    // populate the extended counters
    header_write->extended_number_of_point_records = header_read->number_of_point_records;
    for (i = 0; i < 5; i++)
    {
      header_write->extended_number_of_points_by_return[i] = header_read->number_of_points_by_return[i];
    }
    for (i = 5; i < 15; i++)
    {
      header_write->extended_number_of_points_by_return[i] = 0;
    }

    if (laszip_open_writer(laszip_writer, file_name_out, compress))
    {
      fprintf(stderr,"DLL ERROR: opening laszip writer for '%s'\n", file_name_out);
      byebye(true, argc==1, laszip_writer);
    }

    fprintf(stderr,"writing file '%s' %scompressed\n", file_name_out, (compress ? "" : "un"));

    // get a pointer to the point of the writer that we will populate and write

    laszip_point* point_write;

    if (laszip_get_point_pointer(laszip_writer, &point_write))
    {
      fprintf(stderr,"DLL ERROR: getting point pointer from laszip writer\n");
      byebye(true, argc==1, laszip_writer);
    }

    // read the points

    uint64_t p_count = 0;

    while (p_count < npoints)
    {
      // read a point

      if (laszip_read_point(laszip_reader))
      {
        fprintf(stderr,"DLL ERROR: reading point %lld\n", p_count);
        //fprintf(stderr,"DLL ERROR: reading point %I64\n", p_count);
        byebye(true, argc==1, laszip_reader);
      }

      // copy the point

      point_write->X = point_read->X;
      point_write->Y = point_read->Y;
      point_write->Z = point_read->Z;
      point_write->intensity = point_read->intensity;
      point_write->scan_direction_flag = point_read->scan_direction_flag;
      point_write->edge_of_flight_line = point_read->edge_of_flight_line;
      point_write->user_data = point_read->user_data;
      point_write->point_source_ID = point_read->point_source_ID;

      point_write->gps_time = point_read->gps_time;
      memcpy(point_write->rgb, point_read->rgb, 8);

      point_write->extended_point_type = 1;
      point_write->extended_scanner_channel = 0;
      point_write->extended_classification_flags = (point_read->withheld_flag << 2) | (point_read->keypoint_flag << 1) | (point_read->synthetic_flag << 0);;
      point_write->extended_classification = point_read->classification;
      point_write->extended_return_number = point_read->return_number;
      point_write->extended_number_of_returns = point_read->number_of_returns;
      point_write->extended_scan_angle = (int16_t)( (point_read->scan_angle_rank > 0) ? ((1.0 / 0.006 * point_read->scan_angle_rank) + 0.5) : ((1.0 / 0.006 * point_read->scan_angle_rank) - 0.5) );

      if (point_read->num_extra_bytes)
      {
        memcpy(point_write->extra_bytes, point_read->extra_bytes, point_read->num_extra_bytes);
      }

      // write the point

      if (laszip_write_point(laszip_writer))
      {
        fprintf(stderr,"DLL ERROR: writing point %lld\n", p_count);
        //fprintf(stderr,"DLL ERROR: writing point %I64\n", p_count);
        byebye(true, argc==1, laszip_writer);
      }

      p_count++;
    }

    fprintf(stderr,"successfully read and written %lld points\n", p_count);
    //fprintf(stderr,"successfully read and written %I64d points\n", p_count);

    if (laszip_close_writer(laszip_writer))
    {
      fprintf(stderr,"DLL ERROR: closing laszip writer\n");
      byebye(true, argc==1, laszip_writer);
    }

    if (laszip_destroy(laszip_writer))
    {
      fprintf(stderr,"DLL ERROR: destroying laszip writer\n");
      byebye(true, argc==1);
    }

    if (laszip_close_reader(laszip_reader))
    {
      fprintf(stderr,"DLL ERROR: closing laszip reader\n");
      byebye(true, argc==1, laszip_reader);
    }

    if (laszip_destroy(laszip_reader))
    {
      fprintf(stderr,"DLL ERROR: destroying laszip reader\n");
      byebye(true, argc==1);
    }

    fprintf(stderr,"total time: %g sec for reading %scompressed and writing %scompressed\n", taketime()-start_time, (is_compressed ? "" : "un"), (compress ? "" : "un"));

  } // end of EXAMPLE_TEN

void example11(int argc, char* file_name_in, char* file_name_out)
{
	  int i;
    fprintf(stderr,"running EXAMPLE_ELEVEN (writing points to LAS 1.4 without a-priori knowlegde of bounding box or point count (compatibility only for LAZ))\n");
    double start_time = taketime();

    void * laszip_writer;
    if (laszip_create(&laszip_writer))
    {
      fprintf(stderr,"DLL ERROR: creating laszip writer\n");
      byebye(true, argc==1);
    }

    // get a pointer to the header of the writer so we can populate it

    laszip_header* header;

    if (laszip_get_header_pointer(laszip_writer, &header))
    {
      fprintf(stderr,"DLL ERROR: getting header pointer from laszip writer\n");
      byebye(true, argc==1, laszip_writer);
    }

    // populate the header

    header->file_source_ID = 4711;
    header->global_encoding = (1<<0) | (1<<4);     // see LAS specification for details
    header->version_major = 1;
    header->version_minor = 4;
    strncpy(header->system_identifier, "LASzip DLL example 7", 32);
    header->file_creation_day = 30;
    header->file_creation_year = 2015;
    header->header_size = 375;
    header->offset_to_point_data = 375;
    header->point_data_format = 6;
    header->point_data_record_length = 30;
    header->number_of_point_records = 0;           // legacy 32-bit counters should be zero for new point types > 5
    for (i = 0; i < 5; i++)
    {
      header->number_of_points_by_return[i] = 0;   // legacy 32-bit counters should be zero for new point types > 5
    }
    header->extended_number_of_point_records = 0;  // a-priori unknown number of points
    for (i = 0; i < 15; i++)
    {
      header->extended_number_of_points_by_return[i] = 0;
    }
    header->max_x = 0.0;                           // a-priori unknown bounding box
    header->min_x = 0.0;
    header->max_y = 0.0;
    header->min_y = 0.0;
    header->max_z = 0.0;
    header->min_z = 0.0;

    fprintf(stderr,"offset_to_point_data before adding  empty OGC WKT VLR is : %d\n", (int32_t)header->offset_to_point_data);

    // add intentionally empty OGC WKT

    if (laszip_add_vlr(laszip_writer, "LASF_Projection", 2112, 0, "intentionally empty OGC WKT", 0))
    {
      fprintf(stderr,"DLL ERROR: adding intentionally empty OGC WKT VLR to the header\n");
      byebye(true, argc==1, laszip_writer);
    }

    fprintf(stderr,"offset_to_point_data before adding funny VLR is          : %d\n", (int32_t)header->offset_to_point_data);

    // add some funny VLR

    if (laszip_add_vlr(laszip_writer, "funny", 12345, 0, "just a funny VLR", 0))
    {
      fprintf(stderr,"DLL ERROR: adding funny VLR to the header\n");
      byebye(true, argc==1, laszip_writer);
    }

    fprintf(stderr,"offset_to_point_data after adding VLRs                   : %d\n", (int32_t)header->offset_to_point_data);

    // compressed output or not?

    laszip_BOOL compress = (strstr(file_name_out, ".laz") != 0);

    // if compressed output was requested enable the compatibility mode

    if (compress)
    {
      laszip_BOOL request = 1;
      if (laszip_request_compatibility_mode(laszip_writer, request))
      {
        fprintf(stderr,"DLL ERROR: enabling laszip LAS 1.4 compatibility mode\n");
        byebye(true, argc==1, laszip_writer);
      }
    }

    if (laszip_open_writer(laszip_writer, file_name_out, compress))
    {
      fprintf(stderr,"DLL ERROR: opening laszip writer for '%s'\n", file_name_out);
      byebye(true, argc==1, laszip_writer);
    }

    fprintf(stderr,"writing file '%s' %scompressed\n", file_name_out, (compress ? "" : "un"));

    // get a pointer to the point of the writer that we will populate and write

    laszip_point* point;

    if (laszip_get_point_pointer(laszip_writer, &point))
    {
      fprintf(stderr,"DLL ERROR: getting point pointer from laszip writer\n");
      byebye(true, argc==1, laszip_writer);
    }

    // set points to "extended type"

    point->extended_point_type = 1;

    // write five points

    uint64_t p_count = 0;
    double coordinates[3];

    // populate the first point

    coordinates[0] = 630499.95;
    coordinates[1] = 4834749.17;
    coordinates[2] = 62.15;

    if (laszip_set_coordinates(laszip_writer, coordinates))
    {
      fprintf(stderr,"DLL ERROR: setting coordinates for point %lld\n", p_count);
      //fprintf(stderr,"DLL ERROR: setting coordinates for point %I64d\n", p_count);
      byebye(true, argc==1, laszip_writer);
    }

    point->intensity = 60;
    point->extended_return_number = 2;
    point->extended_number_of_returns = 2;
    point->classification = 2;                // it must be set because it "fits" in 5 bits
    point->extended_classification = 2;
    point->extended_scan_angle = 3500;
    point->extended_scanner_channel = 1;
    point->extended_classification_flags = 8; // overflag flag is set
    point->gps_time = 53413162.560400;

    // write the first point

    if (laszip_write_point(laszip_writer))
    {
      fprintf(stderr,"DLL ERROR: writing point %lld\n", p_count);
      //fprintf(stderr,"DLL ERROR: writing point %I64d\n", p_count);
      byebye(true, argc==1, laszip_writer);
    }

    // update inventory with first point

    if (laszip_update_inventory(laszip_writer))
    {
      fprintf(stderr,"DLL ERROR: updating inventory for point %lld\n", p_count);
      //fprintf(stderr,"DLL ERROR: updating inventory for point %I64d\n", p_count);
      byebye(true, argc==1, laszip_writer);
    }

    p_count++;

    // populate the second point

    coordinates[0] = 630499.83;
    coordinates[1] = 4834748.88;
    coordinates[2] = 62.68;

    if (laszip_set_coordinates(laszip_writer, coordinates))
    {
      fprintf(stderr,"DLL ERROR: setting coordinates for point %lld\n", p_count);
      //fprintf(stderr,"DLL ERROR: setting coordinates for point %I64d\n", p_count);
      byebye(true, argc==1, laszip_writer);
    }

    point->intensity = 90;
    point->extended_return_number = 8;
    point->extended_number_of_returns = 9;
    point->classification = 0;                // it must be set to zero as the real value is stored in the extended field
    point->extended_classification = 41;
    point->extended_scan_angle = 3567;
    point->extended_scanner_channel = 1;
    point->extended_classification_flags = 0; // no flag is not set
    point->gps_time = 53413162.563600;

    // write the second point

    if (laszip_write_point(laszip_writer))
    {
      fprintf(stderr,"DLL ERROR: writing point %lld\n", p_count);
      //fprintf(stderr,"DLL ERROR: writing point %I64d\n", p_count);
      byebye(true, argc==1, laszip_writer);
    }

    // update inventory with second point

    if (laszip_update_inventory(laszip_writer))
    {
      fprintf(stderr,"DLL ERROR: updating inventory for point %lld\n", p_count);
      //fprintf(stderr,"DLL ERROR: updating inventory for point %I64d\n", p_count);
      byebye(true, argc==1, laszip_writer);
    }

    p_count++;

    // populate the third point

    coordinates[0] = 630499.54;
    coordinates[1] = 4834749.66;
    coordinates[2] = 62.66;

    if (laszip_set_coordinates(laszip_writer, coordinates))
    {
      fprintf(stderr,"DLL ERROR: setting coordinates for point %lld\n", p_count);
      //fprintf(stderr,"DLL ERROR: setting coordinates for point %I64d\n", p_count);
      byebye(true, argc==1, laszip_writer);
    }

    point->intensity = 70;
    point->extended_return_number = 9;
    point->extended_number_of_returns = 9;
    point->classification = 0;                // it must be set to zero as the real value is stored in the extended field
    point->extended_classification = 42;
    point->extended_scan_angle = 3633;
    point->extended_scanner_channel = 1;
    point->extended_classification_flags = 0; // no flag is not set
    point->gps_time = 53413162.566800;

    // write the third point

    if (laszip_write_point(laszip_writer))
    {
      fprintf(stderr,"DLL ERROR: writing point %lld\n", p_count);
      //fprintf(stderr,"DLL ERROR: writing point %I64d\n", p_count);
      byebye(true, argc==1, laszip_writer);
    }

    // update inventory with third point

    if (laszip_update_inventory(laszip_writer))
    {
      fprintf(stderr,"DLL ERROR: updating inventory for point %lld\n", p_count);
      //fprintf(stderr,"DLL ERROR: updating inventory for point %I64d\n", p_count);
      byebye(true, argc==1, laszip_writer);
    }

    p_count++;

    // populate the fourth point

    coordinates[0] = 630498.56;
    coordinates[1] = 4834749.41;
    coordinates[2] = 63.68;

    if (laszip_set_coordinates(laszip_writer, coordinates))
    {
      fprintf(stderr,"DLL ERROR: setting coordinates for point %lld\n", p_count);
      //fprintf(stderr,"DLL ERROR: setting coordinates for point %I64d\n", p_count);
      byebye(true, argc==1, laszip_writer);
    }

    point->intensity = 20;
    point->extended_return_number = 1;
    point->extended_number_of_returns = 2;
    point->classification = 5;                // it must be set because it "fits" in 5 bits
    point->extended_classification = 5;
    point->extended_scan_angle = 3700;
    point->extended_scanner_channel = 1;
    point->extended_classification_flags = 0; // no flag is not set
    point->gps_time = 53413162.580200;

    // write the fourth point

    if (laszip_write_point(laszip_writer))
    {
      fprintf(stderr,"DLL ERROR: writing point %lld\n", p_count);
      //fprintf(stderr,"DLL ERROR: writing point %I64d\n", p_count);
      byebye(true, argc==1, laszip_writer);
    }

    // update inventory with fourth point

    if (laszip_update_inventory(laszip_writer))
    {
      fprintf(stderr,"DLL ERROR: updating inventory for point %lld\n", p_count);
      //fprintf(stderr,"DLL ERROR: updating inventory for point %I64d\n", p_count);
      byebye(true, argc==1, laszip_writer);
    }

    p_count++;

    // populate the fifth point

    coordinates[0] = 630498.80;
    coordinates[1] = 4834748.73;
    coordinates[2] = 62.16;

    if (laszip_set_coordinates(laszip_writer, coordinates))
    {
      fprintf(stderr,"DLL ERROR: setting coordinates for point %lld\n", p_count);
      //fprintf(stderr,"DLL ERROR: setting coordinates for point %I64d\n", p_count);
      byebye(true, argc==1, laszip_writer);
    }

    point->intensity = 110;
    point->extended_return_number = 2;
    point->extended_number_of_returns = 2;
    point->classification = 2;                // it must be set because it "fits" in 5 bits
    point->extended_classification = 2;
    point->extended_scan_angle = 3767;
    point->extended_scanner_channel = 1;
    point->extended_classification_flags = 0; // no flag is not set
    point->gps_time = 53413162.580200;

    // write the fifth point

    if (laszip_write_point(laszip_writer))
    {
      fprintf(stderr,"DLL ERROR: writing point %lld\n", p_count);
      //fprintf(stderr,"DLL ERROR: writing point %I64d\n", p_count);
      byebye(true, argc==1, laszip_writer);
    }

    // update inventory with fifth point

    if (laszip_update_inventory(laszip_writer))
    {
      fprintf(stderr,"DLL ERROR: updating inventory for point %lld\n", p_count);
      //fprintf(stderr,"DLL ERROR: updating inventory for point %I64d\n", p_count);
      byebye(true, argc==1, laszip_writer);
    }

    p_count++;

    // get the number of points written so far

    if (laszip_get_point_count(laszip_writer, &p_count))
    {
      fprintf(stderr,"DLL ERROR: getting point count\n");
      byebye(true, argc==1, laszip_writer);
    }

    fprintf(stderr,"successfully written %lld points\n", p_count);
    //fprintf(stderr,"successfully written %" PRId64 " points\n", p_count);

    if (laszip_close_writer(laszip_writer))
    {
      fprintf(stderr,"DLL ERROR: closing laszip writer\n");
      byebye(true, argc==1, laszip_writer);
    }

    if (laszip_destroy(laszip_writer))
    {
      fprintf(stderr,"DLL ERROR: destroying laszip writer\n");
      byebye(true, argc==1);
    }

    fprintf(stderr,"total time: %g sec for writing %scompressed\n", taketime()-start_time, (compress ? "" : "un"));
} // end of EXAMPLE_ELEVEN
