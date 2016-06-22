#include "laszip_api.h"

#include "laszip.hpp"
#include "lasattributer.hpp"
#include "bytestreamin_file.hpp"
#include "bytestreamin_array.hpp"
#include "lasreadpoint.hpp"

int32_t
laszip_get_error(
    laszip_dll_struct *                     laszip_dll
    , char**                    error
)
{
  if (laszip_dll == 0) return 1;

  try
  {
    *error = laszip_dll->error;
  }
  catch (...)
  {
    sprintf(laszip_dll->error, "internal error in laszip_get_error");
    return 1;
  }

  return 0;
}

int32_t
laszip_get_warning(
    laszip_dll_struct *                     laszip_dll
    , char**                    warning
)
{
  try
  {
    *warning = laszip_dll->warning;
  }
  catch (...)
  {
    sprintf(laszip_dll->error, "internal error in laszip_get_warning");
    return 1;
  }

  return 0;
}

int32_t laszip_create(laszip_dll_struct **pointer)
{
  if (pointer == 0) return 1;

  laszip_dll_struct *laszip_dll  = new laszip_dll_struct;
  memset(laszip_dll, 0, sizeof(laszip_dll_struct));
  laszip_clean(laszip_dll);
  *pointer = laszip_dll;

  return 0;
}

int32_t laszip_clean(laszip_dll_struct *laszip_dll)
{
  if (laszip_dll == 0) return 1;

  try
  {
    if (laszip_dll->reader)
    {
      sprintf(laszip_dll->error, "cannot clean while reader is open.");
      return 1;
    }

    // dealloc everything alloc in the header

    if (laszip_dll->header.user_data_in_header)
    {
      delete [] laszip_dll->header.user_data_in_header;
      laszip_dll->header.user_data_in_header = 0;
    }

    if (laszip_dll->header.vlrs)
    {
      U32 i;
      for (i = 0; i < laszip_dll->header.number_of_variable_length_records; i++)
      {
        if (laszip_dll->header.vlrs[i].data)
        {
          delete [] laszip_dll->header.vlrs[i].data;
        }
      }
      free(laszip_dll->header.vlrs);
      laszip_dll->header.vlrs = 0;
    }

    if (laszip_dll->header.user_data_after_header)
    {
      delete [] laszip_dll->header.user_data_after_header;
      laszip_dll->header.user_data_after_header = 0;
    }

    // dealloc everything alloc in the point

    if (laszip_dll->point.extra_bytes)
    {
      delete [] laszip_dll->point.extra_bytes;
      laszip_dll->point.extra_bytes = 0;
    }

    // dealloc the attributer

    if (laszip_dll->attributer)
    {
      delete laszip_dll->attributer;
      laszip_dll->attributer = 0;
    }

    // zero everything

    memset(laszip_dll, 0, sizeof(laszip_dll_struct));

    // create default header

    sprintf(laszip_dll->header.generating_software, "LASzip DLL %d.%d r%d (%d)", LASZIP_VERSION_MAJOR, LASZIP_VERSION_MINOR, LASZIP_VERSION_REVISION, LASZIP_VERSION_BUILD_DATE);
    laszip_dll->header.version_major = 1;
    laszip_dll->header.version_minor = 2;
    laszip_dll->header.header_size = 227;
    laszip_dll->header.offset_to_point_data = 227;
    laszip_dll->header.point_data_format = 1;
    laszip_dll->header.point_data_record_length = 28;
    laszip_dll->header.x_scale_factor = 0.01;
    laszip_dll->header.y_scale_factor = 0.01;
    laszip_dll->header.z_scale_factor = 0.01;
  }
  catch (...)
  {
    sprintf(laszip_dll->error, "internal error in laszip_clean");
    return 1;
  }

  return 0;
}

int32_t laszip_destroy(laszip_dll_struct *laszip_dll)
{
  if (laszip_dll == 0) return 1;

  int err = 0;

  try
  {
    err = laszip_clean(laszip_dll);
    delete laszip_dll;
  }
  catch (...)
  {
    return 1;
  }

  return err;
}

int32_t
laszip_get_header_pointer(
    laszip_dll_struct *laszip_dll
    , laszip_header_struct**           header_pointer
)
{
  if (laszip_dll == 0) return 1;

  if (header_pointer == 0)
  {
    sprintf(laszip_dll->error, "laszip_header_struct pointer 'header_pointer' is zero");
    return 1;
  }

  *header_pointer = &laszip_dll->header;
  laszip_dll->error[0] = '\0';
  return 0;
}

int32_t
laszip_get_point_pointer(
    laszip_dll_struct *laszip_dll
    , laszip_point_struct**            point_pointer
)
{
  if (laszip_dll == 0) return 1;

  if (point_pointer == 0)
  {
    sprintf(laszip_dll->error, "laszip_point_struct pointer 'point_pointer' is zero");
    return 1;
  }

  *point_pointer = &laszip_dll->point;
  laszip_dll->error[0] = '\0';
  return 0;
}

int32_t
laszip_get_point_count(
    laszip_dll_struct *laszip_dll
    , uint64_t*                      count
)
{
  if (laszip_dll == 0) return 1;

  if (count == 0)
  {
    sprintf(laszip_dll->error, "uint64_t pointer 'count' is zero");
    return 1;
  }

  if (laszip_dll->reader == 0)
  {
    sprintf(laszip_dll->error, "getting count before reader or writer was opened");
    return 1;
  }
  *count = laszip_dll->p_count;

  laszip_dll->error[0] = '\0';
  return 0;
}

int32_t
laszip_check_for_integer_overflow(
    laszip_dll_struct *laszip_dll
)
{
  if (laszip_dll == 0) return 1;

  try
  {
    // get a pointer to the header

    laszip_header_struct* header = &(laszip_dll->header);

    // quantize and dequantize the bounding box with current scale_factor and offset

    I32 quant_min_x = I32_QUANTIZE((header->min_x-header->x_offset)/header->x_scale_factor);
    I32 quant_max_x = I32_QUANTIZE((header->max_x-header->x_offset)/header->x_scale_factor);
    I32 quant_min_y = I32_QUANTIZE((header->min_y-header->y_offset)/header->y_scale_factor);
    I32 quant_max_y = I32_QUANTIZE((header->max_y-header->y_offset)/header->y_scale_factor);
    I32 quant_min_z = I32_QUANTIZE((header->min_z-header->z_offset)/header->z_scale_factor);
    I32 quant_max_z = I32_QUANTIZE((header->max_z-header->z_offset)/header->z_scale_factor);

    F64 dequant_min_x = header->x_scale_factor*quant_min_x+header->x_offset;
    F64 dequant_max_x = header->x_scale_factor*quant_max_x+header->x_offset;
    F64 dequant_min_y = header->y_scale_factor*quant_min_y+header->y_offset;
    F64 dequant_max_y = header->y_scale_factor*quant_max_y+header->y_offset;
    F64 dequant_min_z = header->z_scale_factor*quant_min_z+header->z_offset;
    F64 dequant_max_z = header->z_scale_factor*quant_max_z+header->z_offset;

    // make sure that there is not sign flip (a 32-bit integer overflow) for the bounding box

    if ((header->min_x > 0) != (dequant_min_x > 0))
    {
      sprintf(laszip_dll->error, "quantization sign flip for min_x from %g to %g. set scale factor for x coarser than %g\n", header->min_x, dequant_min_x, header->x_scale_factor);
      return 1;
    }
    if ((header->max_x > 0) != (dequant_max_x > 0))
    {
      sprintf(laszip_dll->error, "quantization sign flip for max_x from %g to %g. set scale factor for x coarser than %g\n", header->max_x, dequant_max_x, header->x_scale_factor);
      return 1;
    }
    if ((header->min_y > 0) != (dequant_min_y > 0))
    {
      sprintf(laszip_dll->error, "quantization sign flip for min_y from %g to %g. set scale factor for y coarser than %g\n", header->min_y, dequant_min_y, header->y_scale_factor);
      return 1;
    }
    if ((header->max_y > 0) != (dequant_max_y > 0))
    {
      sprintf(laszip_dll->error, "quantization sign flip for max_y from %g to %g. set scale factor for y coarser than %g\n", header->max_y, dequant_max_y, header->y_scale_factor);
      return 1;
    }
    if ((header->min_z > 0) != (dequant_min_z > 0))
    {
      sprintf(laszip_dll->error, "quantization sign flip for min_z from %g to %g. set scale factor for z coarser than %g\n", header->min_z, dequant_min_z, header->z_scale_factor);
      return 1;
    }
    if ((header->max_z > 0) != (dequant_max_z > 0))
    {
      sprintf(laszip_dll->error, "quantization sign flip for max_z from %g to %g. set scale factor for z coarser than %g\n", header->max_z, dequant_max_z, header->z_scale_factor);
      return 1;
    }
  }
  catch (...)
  {
    sprintf(laszip_dll->error, "internal error in laszip_auto_offset");
    return 1;
  }

  laszip_dll->error[0] = '\0';
  return 0;
}

int32_t laszip_auto_offset(laszip_dll_struct *laszip_dll)
{
  if (laszip_dll == 0) return 1;

  try
  {
    if (laszip_dll->reader)
    {
      sprintf(laszip_dll->error, "cannot auto offset after reader was opened");
      return 1;
    }

    // get a pointer to the header

    laszip_header_struct* header = &(laszip_dll->header);

    // check scale factor

    F64 x_scale_factor = header->x_scale_factor;
    F64 y_scale_factor = header->y_scale_factor;
    F64 z_scale_factor = header->z_scale_factor;

    if ((x_scale_factor <= 0) || !F64_IS_FINITE(x_scale_factor))
    {
      sprintf(laszip_dll->error, "invalid x scale_factor %g in header", header->x_scale_factor);
      return 1;
    }

    if ((y_scale_factor <= 0) || !F64_IS_FINITE(y_scale_factor))
    {
      sprintf(laszip_dll->error, "invalid y scale_factor %g in header", header->y_scale_factor);
      return 1;
    }

    if ((z_scale_factor <= 0) || !F64_IS_FINITE(z_scale_factor))
    {
      sprintf(laszip_dll->error, "invalid z scale_factor %g in header", header->z_scale_factor);
      return 1;
    }

    F64 center_bb_x = (header->min_x + header->max_x) / 2;
    F64 center_bb_y = (header->min_y + header->max_y) / 2;
    F64 center_bb_z = (header->min_z + header->max_z) / 2;

    if (!F64_IS_FINITE(center_bb_x))
    {
      sprintf(laszip_dll->error, "invalid x coordinate at center of bounding box (min: %g max: %g)", header->min_x, header->max_x);
      return 1;
    }

    if (!F64_IS_FINITE(center_bb_y))
    {
      sprintf(laszip_dll->error, "invalid y coordinate at center of  bounding box (min: %g max: %g)", header->min_y, header->max_y);
      return 1;
    }

    if (!F64_IS_FINITE(center_bb_z))
    {
      sprintf(laszip_dll->error, "invalid z coordinate at center of  bounding box (min: %g max: %g)", header->min_z, header->max_z);
      return 1;
    }

    F64 x_offset = header->x_offset;
    F64 y_offset = header->y_offset;
    F64 z_offset = header->z_offset;

    header->x_offset = (I64_FLOOR(center_bb_x/x_scale_factor/10000000))*10000000*x_scale_factor;
    header->y_offset = (I64_FLOOR(center_bb_y/y_scale_factor/10000000))*10000000*y_scale_factor;
    header->z_offset = (I64_FLOOR(center_bb_z/z_scale_factor/10000000))*10000000*z_scale_factor;

    if (laszip_check_for_integer_overflow(laszip_dll))
    {
      header->x_offset = x_offset;
      header->y_offset = y_offset;
      header->z_offset = z_offset;
      return 1;
    }
  }
  catch (...)
  {
    sprintf(laszip_dll->error, "internal error in laszip_auto_offset");
    return 1;
  }

  laszip_dll->error[0] = '\0';
  return 0;
}

int32_t
laszip_get_coordinates(laszip_dll_struct *laszip_dll, double *coordinates)
{
  if (laszip_dll == 0) return 1;

  try
  {
    if (coordinates == 0)
    {
      sprintf(laszip_dll->error, "double pointer 'coordinates' is zero");
      return 1;
    }

    // get a pointer to the header

    laszip_header_struct* header = &(laszip_dll->header);

    // get a pointer to the point

    laszip_point_struct* point = &(laszip_dll->point);

    // get the coordinates

    coordinates[0] = header->x_scale_factor*point->X+header->x_offset;
    coordinates[1] = header->y_scale_factor*point->Y+header->y_offset;
    coordinates[2] = header->z_scale_factor*point->Z+header->z_offset;
  }
  catch (...)
  {
    sprintf(laszip_dll->error, "internal error in laszip_get_coordinates");
    return 1;
  }

  laszip_dll->error[0] = '\0';
  return 0;
}

int32_t
laszip_add_vlr(
    laszip_dll_struct *                     laszip_dll
    , const char*               user_id
    , uint16_t                       record_id
    , uint16_t                       record_length_after_header
    , const char*               description
    , const uint8_t*                 data
)
{
  if (laszip_dll == 0) return 1;

  try
  {
    if (user_id == 0)
    {
      sprintf(laszip_dll->error, "char pointer 'user_id' is zero");
      return 1;
    }

    if ((record_length_after_header > 0) && (data == 0))
    {
      sprintf(laszip_dll->error, "record_length_after_header of VLR is %u but data pointer is zero", (U32)record_length_after_header);
      return 1;
    }

    if (laszip_dll->reader)
    {
      sprintf(laszip_dll->error, "cannot add vlr after reader was opened");
      return 1;
    }

    U32 i = 0;

    if (laszip_dll->header.vlrs)
    {
      // overwrite existing VLR ?

      for (i = 0; i < laszip_dll->header.number_of_variable_length_records; i++)
      {
        if ((strncmp(laszip_dll->header.vlrs[i].user_id, user_id, 16) == 0) && (laszip_dll->header.vlrs[i].record_id == record_id))
        {
          if (laszip_dll->header.vlrs[i].record_length_after_header)
          {
            laszip_dll->header.offset_to_point_data -= laszip_dll->header.vlrs[i].record_length_after_header;
            laszip_dll->header.vlrs[i].record_length_after_header = 0;
            delete [] laszip_dll->header.vlrs[i].data;
            laszip_dll->header.vlrs[i].data = 0;
          }
          break;
        }
      }

      // create new VLR

      if (i == laszip_dll->header.number_of_variable_length_records)
      {
        laszip_dll->header.number_of_variable_length_records++;
        laszip_dll->header.offset_to_point_data += 54;
        laszip_dll->header.vlrs = (laszip_vlr_struct*)realloc(laszip_dll->header.vlrs, sizeof(laszip_vlr_struct)*laszip_dll->header.number_of_variable_length_records);
        if (laszip_dll->header.vlrs == 0)
        {
          sprintf(laszip_dll->error, "reallocating vlrs[%u] array", laszip_dll->header.number_of_variable_length_records);
          return 1;
        }
      }
    }
    else
    {
      laszip_dll->header.number_of_variable_length_records = 1;
      laszip_dll->header.offset_to_point_data += 54;
      laszip_dll->header.vlrs = (laszip_vlr_struct*)malloc(sizeof(laszip_vlr_struct));
      if (laszip_dll->header.vlrs == 0)
      {
        sprintf(laszip_dll->error, "allocating vlrs[1] array");
        return 1;
      }
    }

    // zero the VLR

    memset(&(laszip_dll->header.vlrs[i]), 0, sizeof(laszip_vlr_struct));

    // copy the VLR

    laszip_dll->header.vlrs[i].reserved = 0xAABB;
    strncpy(laszip_dll->header.vlrs[i].user_id, user_id, 16);
    laszip_dll->header.vlrs[i].record_id = record_id;
    laszip_dll->header.vlrs[i].record_length_after_header = record_length_after_header;
    if (description)
    {
      strncpy(laszip_dll->header.vlrs[i].description, description, 32);
    }
    else
    {
      sprintf(laszip_dll->header.vlrs[i].description, "LASzip DLL %d.%d r%d (%d)", LASZIP_VERSION_MAJOR, LASZIP_VERSION_MINOR, LASZIP_VERSION_REVISION, LASZIP_VERSION_BUILD_DATE);
    }
    if (record_length_after_header)
    {
      laszip_dll->header.offset_to_point_data += record_length_after_header;
      laszip_dll->header.vlrs[i].data = new U8[record_length_after_header];
      memcpy(laszip_dll->header.vlrs[i].data, data, record_length_after_header);
    }
  }
  catch (...)
  {
    sprintf(laszip_dll->error, "internal error in laszip_add_vlr");
    return 1;
  }

  laszip_dll->error[0] = '\0';
  return 0;
}

int32_t
laszip_remove_vlr(
    laszip_dll_struct *                     laszip_dll
    , const char*               user_id
    , uint16_t                       record_id
)
{
  if (laszip_dll == 0) return 1;

  try
  {
    if (user_id == 0)
    {
      sprintf(laszip_dll->error, "char pointer 'user_id' is zero");
      return 1;
    }

    if (laszip_dll->reader)
    {
      sprintf(laszip_dll->error, "cannot remove vlr after reader was opened");
      return 1;
    }

    U32 i = 0;

    if (laszip_dll->header.vlrs)
    {
      for (i = 0; i < laszip_dll->header.number_of_variable_length_records; i++)
      {
        if ((strncmp(laszip_dll->header.vlrs[i].user_id, user_id, 16) == 0) && (laszip_dll->header.vlrs[i].record_id == record_id))
        {
          if (laszip_dll->header.vlrs[i].record_length_after_header)
          {
            laszip_dll->header.offset_to_point_data -= (54 + laszip_dll->header.vlrs[i].record_length_after_header);
            delete [] laszip_dll->header.vlrs[i].data;
            laszip_dll->header.vlrs[i].data = 0;
          }
          laszip_dll->header.number_of_variable_length_records--;
          //
          // HOBU: should this be for ( i = 0; ...) ?
          // TODO: verify this is correct
          //for (i = i; i < laszip_dll->header.number_of_variable_length_records; i++)
          for (; i < laszip_dll->header.number_of_variable_length_records; i++)
          {
            laszip_dll->header.vlrs[i] = laszip_dll->header.vlrs[i+1];
          }
          if (laszip_dll->header.number_of_variable_length_records)
          {
            laszip_dll->header.vlrs = (laszip_vlr_struct*)realloc(laszip_dll->header.vlrs, sizeof(laszip_vlr_struct)*laszip_dll->header.number_of_variable_length_records);
            if (laszip_dll->header.vlrs == 0)
            {
              sprintf(laszip_dll->error, "reallocating vlrs[%u] array", laszip_dll->header.number_of_variable_length_records);
              return 1;
            }
          }
          else
          {
            free(laszip_dll->header.vlrs);
            laszip_dll->header.vlrs = 0;
          }
          i = U32_MAX;
          break;
        }
      }
      if (i != U32_MAX)
      {
        sprintf(laszip_dll->error, "cannot find VLR with user_id '%s' and record_id %d among the %u VLRs in the header", user_id, (I32)record_id, laszip_dll->header.number_of_variable_length_records);
        return 1;
      }
    }
    else
    {
      sprintf(laszip_dll->error, "cannot remove VLR with user_id '%s' and record_id %d because header has no VLRs", user_id, (I32)record_id);
      return 1;
    }
  }
  catch (...)
  {
    sprintf(laszip_dll->error, "internal error in laszip_add_vlr");
    return 1;
  }

  laszip_dll->error[0] = '\0';
  return 0;
}

int32_t
laszip_preserve_generating_software(
    laszip_dll_struct *                     laszip_dll
    , const laszip_BOOL                preserve
)
{
  if (laszip_dll == 0) return 1;

  try
  {
    if (laszip_dll->reader)
    {
      sprintf(laszip_dll->error, "reader is already open");
      return 1;
    }

    laszip_dll->preserve_generating_software = preserve;
  }
  catch (...)
  {
    sprintf(laszip_dll->error, "internal error in laszip_preserve_generating_software");
    return 1;
  }

  laszip_dll->error[0] = '\0';
  return 0;
}

int32_t
laszip_request_compatibility_mode(
    laszip_dll_struct *                     laszip_dll
    , const laszip_BOOL                request
)
{
  if (laszip_dll == 0) return 1;

  if (laszip_dll->reader)
  {
    sprintf(laszip_dll->error, "reader is already open");
    return 1;
  }

  laszip_dll->request_compatibility_mode = request;

  laszip_dll->error[0] = '\0';
  return 0;
}

int32_t
laszip_open_reader(
    laszip_dll_struct *                     laszip_dll
    , const char*               file_name
    , laszip_BOOL*                     is_compressed
)
{
  if (laszip_dll == 0) return 1;

  try
  {
    if (file_name == 0)
    {
      sprintf(laszip_dll->error, "char pointer 'file_name' is zero");
      return 1;
    }

    if (is_compressed == 0)
    {
      sprintf(laszip_dll->error, "laszip_BOOL pointer 'is_compressed' is zero");
      return 1;
    }

    if (laszip_dll->reader)
    {
      sprintf(laszip_dll->error, "reader is already open");
      return 1;
    }

    // open the file

    laszip_dll->file = fopen(file_name, "rb");

    if (laszip_dll->file == 0)
    {
      sprintf(laszip_dll->error, "cannot open file '%s'", file_name);
      return 1;
    }

    if (setvbuf(laszip_dll->file, NULL, _IOFBF, 262144) != 0)
    {
      sprintf(laszip_dll->warning, "setvbuf() failed with buffer size 262144\n");
    }

    laszip_dll->streamin = new ByteStreamInFileLE(laszip_dll->file);

    if (laszip_dll->streamin == 0)
    {
      sprintf(laszip_dll->error, "could not alloc ByteStreamInFile");
      return 1;
    }

    // read the header variable after variable

    U32 i;

    CHAR file_signature[5];
    try { laszip_dll->streamin->getBytes((U8*)file_signature, 4); } catch(...)
    {
      sprintf(laszip_dll->error, "reading header.file_signature");
      return 1;
    }
    if (strncmp(file_signature, "LASF", 4) != 0)
    {
      sprintf(laszip_dll->error, "wrong file_signature. not a LAS/LAZ file.");
      return 1;
    }
    try { laszip_dll->streamin->get16bitsLE((U8*)&(laszip_dll->header.file_source_ID)); } catch(...)
    {
      sprintf(laszip_dll->error, "reading header.file_source_ID");
      return 1;
    }
    try { laszip_dll->streamin->get16bitsLE((U8*)&(laszip_dll->header.global_encoding)); } catch(...)
    {
      sprintf(laszip_dll->error, "reading header.global_encoding");
      return 1;
    }
    try { laszip_dll->streamin->get32bitsLE((U8*)&(laszip_dll->header.project_ID_GUID_data_1)); } catch(...)
    {
      sprintf(laszip_dll->error, "reading header.project_ID_GUID_data_1");
      return 1;
    }
    try { laszip_dll->streamin->get16bitsLE((U8*)&(laszip_dll->header.project_ID_GUID_data_2)); } catch(...)
    {
      sprintf(laszip_dll->error, "reading header.project_ID_GUID_data_2");
      return 1;
    }
    try { laszip_dll->streamin->get16bitsLE((U8*)&(laszip_dll->header.project_ID_GUID_data_3)); } catch(...)
    {
      sprintf(laszip_dll->error, "reading header.project_ID_GUID_data_3");
      return 1;
    }
    try { laszip_dll->streamin->getBytes((U8*)laszip_dll->header.project_ID_GUID_data_4, 8); } catch(...)
    {
      sprintf(laszip_dll->error, "reading header.project_ID_GUID_data_4");
      return 1;
    }
    try { laszip_dll->streamin->getBytes((U8*)&(laszip_dll->header.version_major), 1); } catch(...)
    {
      sprintf(laszip_dll->error, "reading header.version_major");
      return 1;
    }
    try { laszip_dll->streamin->getBytes((U8*)&(laszip_dll->header.version_minor), 1); } catch(...)
    {
      sprintf(laszip_dll->error, "reading header.version_minor");
      return 1;
    }
    try { laszip_dll->streamin->getBytes((U8*)laszip_dll->header.system_identifier, 32); } catch(...)
    {
      sprintf(laszip_dll->error, "reading header.system_identifier");
      return 1;
    }
    try { laszip_dll->streamin->getBytes((U8*)laszip_dll->header.generating_software, 32); } catch(...)
    {
      sprintf(laszip_dll->error, "reading header.generating_software");
      return 1;
    }
    try { laszip_dll->streamin->get16bitsLE((U8*)&(laszip_dll->header.file_creation_day)); } catch(...)
    {
      sprintf(laszip_dll->error, "reading header.file_creation_day");
      return 1;
    }
    try { laszip_dll->streamin->get16bitsLE((U8*)&(laszip_dll->header.file_creation_year)); } catch(...)
    {
      sprintf(laszip_dll->error, "reading header.file_creation_year");
      return 1;
    }
    try { laszip_dll->streamin->get16bitsLE((U8*)&(laszip_dll->header.header_size)); } catch(...)
    {
      sprintf(laszip_dll->error, "reading header.header_size");
      return 1;
    }
    try { laszip_dll->streamin->get32bitsLE((U8*)&(laszip_dll->header.offset_to_point_data)); } catch(...)
    {
      sprintf(laszip_dll->error, "reading header.offset_to_point_data");
      return 1;
    }
    try { laszip_dll->streamin->get32bitsLE((U8*)&(laszip_dll->header.number_of_variable_length_records)); } catch(...)
    {
      sprintf(laszip_dll->error, "reading header.number_of_variable_length_records");
      return 1;
    }
    try { laszip_dll->streamin->getBytes((U8*)&(laszip_dll->header.point_data_format), 1); } catch(...)
    {
      sprintf(laszip_dll->error, "reading header.point_data_format");
      return 1;
    }
    try { laszip_dll->streamin->get16bitsLE((U8*)&(laszip_dll->header.point_data_record_length)); } catch(...)
    {
      sprintf(laszip_dll->error, "reading header.point_data_record_length");
      return 1;
    }
    try { laszip_dll->streamin->get32bitsLE((U8*)&(laszip_dll->header.number_of_point_records)); } catch(...)
    {
      sprintf(laszip_dll->error, "reading header.number_of_point_records");
      return 1;
    }
    for (i = 0; i < 5; i++)
    {
      try { laszip_dll->streamin->get32bitsLE((U8*)&(laszip_dll->header.number_of_points_by_return[i])); } catch(...)
      {
        sprintf(laszip_dll->error, "reading header.number_of_points_by_return %d", i);
        return 1;
      }
    }
    try { laszip_dll->streamin->get64bitsLE((U8*)&(laszip_dll->header.x_scale_factor)); } catch(...)
    {
      sprintf(laszip_dll->error, "reading header.x_scale_factor");
      return 1;
    }
    try { laszip_dll->streamin->get64bitsLE((U8*)&(laszip_dll->header.y_scale_factor)); } catch(...)
    {
      sprintf(laszip_dll->error, "reading header.y_scale_factor");
      return 1;
    }
    try { laszip_dll->streamin->get64bitsLE((U8*)&(laszip_dll->header.z_scale_factor)); } catch(...)
    {
      sprintf(laszip_dll->error, "reading header.z_scale_factor");
      return 1;
    }
    try { laszip_dll->streamin->get64bitsLE((U8*)&(laszip_dll->header.x_offset)); } catch(...)
    {
      sprintf(laszip_dll->error, "reading header.x_offset");
      return 1;
    }
    try { laszip_dll->streamin->get64bitsLE((U8*)&(laszip_dll->header.y_offset)); } catch(...)
    {
      sprintf(laszip_dll->error, "reading header.y_offset");
      return 1;
    }
    try { laszip_dll->streamin->get64bitsLE((U8*)&(laszip_dll->header.z_offset)); } catch(...)
    {
      sprintf(laszip_dll->error, "reading header.z_offset");
      return 1;
    }
    try { laszip_dll->streamin->get64bitsLE((U8*)&(laszip_dll->header.max_x)); } catch(...)
    {
      sprintf(laszip_dll->error, "reading header.max_x");
      return 1;
    }
    try { laszip_dll->streamin->get64bitsLE((U8*)&(laszip_dll->header.min_x)); } catch(...)
    {
      sprintf(laszip_dll->error, "reading header.min_x");
      return 1;
    }
    try { laszip_dll->streamin->get64bitsLE((U8*)&(laszip_dll->header.max_y)); } catch(...)
    {
      sprintf(laszip_dll->error, "reading header.max_y");
      return 1;
    }
    try { laszip_dll->streamin->get64bitsLE((U8*)&(laszip_dll->header.min_y)); } catch(...)
    {
      sprintf(laszip_dll->error, "reading header.min_y");
      return 1;
    }
    try { laszip_dll->streamin->get64bitsLE((U8*)&(laszip_dll->header.max_z)); } catch(...)
    {
      sprintf(laszip_dll->error, "reading header.max_z");
      return 1;
    }
    try { laszip_dll->streamin->get64bitsLE((U8*)&(laszip_dll->header.min_z)); } catch(...)
    {
      sprintf(laszip_dll->error, "reading header.min_z");
      return 1;
    }

    // special handling for LAS 1.3
    if ((laszip_dll->header.version_major == 1) && (laszip_dll->header.version_minor >= 3))
    {
      if (laszip_dll->header.header_size < 235)
      {
        sprintf(laszip_dll->error, "for LAS 1.%d header_size should at least be 235 but it is only %d", laszip_dll->header.version_minor, laszip_dll->header.header_size);
        return 1;
      }
      else
      {
        try { laszip_dll->streamin->get64bitsLE((U8*)&(laszip_dll->header.start_of_waveform_data_packet_record)); } catch(...)
        {
          sprintf(laszip_dll->error, "reading header.start_of_waveform_data_packet_record");
          return 1;
        }
        laszip_dll->header.user_data_in_header_size = laszip_dll->header.header_size - 235;
      }
    }
    else
    {
      laszip_dll->header.user_data_in_header_size = laszip_dll->header.header_size - 227;
    }

    // special handling for LAS 1.4
    if ((laszip_dll->header.version_major == 1) && (laszip_dll->header.version_minor >= 4))
    {
      if (laszip_dll->header.header_size < 375)
      {
        sprintf(laszip_dll->error, "for LAS 1.%d header_size should at least be 375 but it is only %d", laszip_dll->header.version_minor, laszip_dll->header.header_size);
        return 1;
      }
      else
      {
        try { laszip_dll->streamin->get64bitsLE((U8*)&(laszip_dll->header.start_of_first_extended_variable_length_record)); } catch(...)
        {
          sprintf(laszip_dll->error, "reading header.start_of_first_extended_variable_length_record");
          return 1;
        }
        try { laszip_dll->streamin->get32bitsLE((U8*)&(laszip_dll->header.number_of_extended_variable_length_records)); } catch(...)
        {
          sprintf(laszip_dll->error, "reading header.number_of_extended_variable_length_records");
          return 1;
        }
        try { laszip_dll->streamin->get64bitsLE((U8*)&(laszip_dll->header.extended_number_of_point_records)); } catch(...)
        {
          sprintf(laszip_dll->error, "reading header.extended_number_of_point_records");
          return 1;
        }
        for (i = 0; i < 15; i++)
        {
          try { laszip_dll->streamin->get64bitsLE((U8*)&(laszip_dll->header.extended_number_of_points_by_return[i])); } catch(...)
          {
            sprintf(laszip_dll->error, "reading header.extended_number_of_points_by_return[%d]", i);
            return 1;
          }
        }
        laszip_dll->header.user_data_in_header_size = laszip_dll->header.header_size - 375;
      }
    }

    // load any number of user-defined bytes that might have been added to the header
    if (laszip_dll->header.user_data_in_header_size)
    {
      if (laszip_dll->header.user_data_in_header)
      {
        delete [] laszip_dll->header.user_data_in_header;
      }
      laszip_dll->header.user_data_in_header = new U8[laszip_dll->header.user_data_in_header_size];

      try { laszip_dll->streamin->getBytes((U8*)laszip_dll->header.user_data_in_header, laszip_dll->header.user_data_in_header_size); } catch(...)
      {
        sprintf(laszip_dll->error, "reading %u bytes of data into header.user_data_in_header", laszip_dll->header.user_data_in_header_size);
        return 1;
      }
    }

    // read variable length records into the header

    U32 vlrs_size = 0;
    LASzip* laszip = 0;

    if (laszip_dll->header.number_of_variable_length_records)
    {
      U32 i;

      laszip_dll->header.vlrs = (laszip_vlr*)malloc(sizeof(laszip_vlr)*laszip_dll->header.number_of_variable_length_records);

      if (laszip_dll->header.vlrs == 0)
      {
        sprintf(laszip_dll->error, "allocating %u VLRs", laszip_dll->header.number_of_variable_length_records);
        return 1;
      }

      for (i = 0; i < laszip_dll->header.number_of_variable_length_records; i++)
      {
        // make sure there are enough bytes left to read a variable length record before the point block starts

        if (((int)laszip_dll->header.offset_to_point_data - vlrs_size - laszip_dll->header.header_size) < 54)
        {
          sprintf(laszip_dll->warning, "only %d bytes until point block after reading %d of %d vlrs. skipping remaining vlrs ...", (int)laszip_dll->header.offset_to_point_data - vlrs_size - laszip_dll->header.header_size, i, laszip_dll->header.number_of_variable_length_records);
          laszip_dll->header.number_of_variable_length_records = i;
          break;
        }

        // read variable length records variable after variable (to avoid alignment issues)

        try { laszip_dll->streamin->get16bitsLE((U8*)&(laszip_dll->header.vlrs[i].reserved)); } catch(...)
        {
          sprintf(laszip_dll->error, "reading header.vlrs[%u].reserved", i);
          return 1;
        }

        try { laszip_dll->streamin->getBytes((U8*)laszip_dll->header.vlrs[i].user_id, 16); } catch(...)
        {
          sprintf(laszip_dll->error, "reading header.vlrs[%u].user_id", i);
          return 1;
        }
        try { laszip_dll->streamin->get16bitsLE((U8*)&(laszip_dll->header.vlrs[i].record_id)); } catch(...)
        {
          sprintf(laszip_dll->error, "reading header.vlrs[%u].record_id", i);
          return 1;
        }
        try { laszip_dll->streamin->get16bitsLE((U8*)&(laszip_dll->header.vlrs[i].record_length_after_header)); } catch(...)
        {
          sprintf(laszip_dll->error, "reading header.vlrs[%u].record_length_after_header", i);
          return 1;
        }
        try { laszip_dll->streamin->getBytes((U8*)laszip_dll->header.vlrs[i].description, 32); } catch(...)
        {
          sprintf(laszip_dll->error, "reading header.vlrs[%u].description", i);
          return 1;
        }

        // keep track on the number of bytes we have read so far

        vlrs_size += 54;

        // check variable length record contents

        if (laszip_dll->header.vlrs[i].reserved != 0xAABB)
        {
          sprintf(laszip_dll->warning,"wrong header.vlrs[%d].reserved: %d != 0xAABB", i, laszip_dll->header.vlrs[i].reserved);
        }

        // make sure there are enough bytes left to read the data of the variable length record before the point block starts

        if (((int)laszip_dll->header.offset_to_point_data - vlrs_size - laszip_dll->header.header_size) < laszip_dll->header.vlrs[i].record_length_after_header)
        {
          sprintf(laszip_dll->warning, "only %d bytes until point block when trying to read %d bytes into header.vlrs[%d].data", (int)laszip_dll->header.offset_to_point_data - vlrs_size - laszip_dll->header.header_size, laszip_dll->header.vlrs[i].record_length_after_header, i);
          laszip_dll->header.vlrs[i].record_length_after_header = (int)laszip_dll->header.offset_to_point_data - vlrs_size - laszip_dll->header.header_size;
        }

        // load data following the header of the variable length record

        if (laszip_dll->header.vlrs[i].record_length_after_header)
        {
          if ((strcmp(laszip_dll->header.vlrs[i].user_id, "laszip encoded") == 0) && (laszip_dll->header.vlrs[i].record_id == 22204))
          {
            if (laszip)
            {
              delete laszip;
            }

            laszip = new LASzip();

            if (laszip == 0)
            {
              sprintf(laszip_dll->error, "could not alloc LASzip");
              return 1;
            }

            // read the LASzip VLR payload

            //     U16  compressor                2 bytes
            //     U32  coder                     2 bytes
            //     U8   version_major             1 byte
            //     U8   version_minor             1 byte
            //     U16  version_revision          2 bytes
            //     U32  options                   4 bytes
            //     I32  chunk_size                4 bytes
            //     I64  number_of_special_evlrs   8 bytes
            //     I64  offset_to_special_evlrs   8 bytes
            //     U16  num_items                 2 bytes
            //        U16 type                2 bytes * num_items
            //        U16 size                2 bytes * num_items
            //        U16 version             2 bytes * num_items
            // which totals 34+6*num_items

            try { laszip_dll->streamin->get16bitsLE((U8*)&(laszip->compressor)); } catch(...)
            {
              sprintf(laszip_dll->error, "reading compressor %d", (I32)laszip->compressor);
              return 1;
            }
            try { laszip_dll->streamin->get16bitsLE((U8*)&(laszip->coder)); } catch(...)
            {
              sprintf(laszip_dll->error, "reading coder %d", (I32)laszip->coder);
              return 1;
            }
            try { laszip_dll->streamin->getBytes((U8*)&(laszip->version_major), 1); } catch(...)
            {
              sprintf(laszip_dll->error, "reading version_major %d", (I32)laszip->version_major);
              return 1;
            }
            try { laszip_dll->streamin->getBytes((U8*)&(laszip->version_minor), 1); } catch(...)
            {
              sprintf(laszip_dll->error, "reading version_minor %d", (I32)laszip->version_minor);
              return 1;
            }
            try { laszip_dll->streamin->get16bitsLE((U8*)&(laszip->version_revision)); } catch(...)
            {
              sprintf(laszip_dll->error, "reading version_revision %d", (I32)laszip->version_revision);
              return 1;
            }
            try { laszip_dll->streamin->get32bitsLE((U8*)&(laszip->options)); } catch(...)
            {
              sprintf(laszip_dll->error, "reading options %u", laszip->options);
              return 1;
            }
            try { laszip_dll->streamin->get32bitsLE((U8*)&(laszip->chunk_size)); } catch(...)
            {
              sprintf(laszip_dll->error, "reading chunk_size %u", laszip->chunk_size);
              return 1;
            }
            try { laszip_dll->streamin->get64bitsLE((U8*)&(laszip->number_of_special_evlrs)); } catch(...)
            {
              sprintf(laszip_dll->error, "reading number_of_special_evlrs %d", (I32)laszip->number_of_special_evlrs);
              return 1;
            }
            try { laszip_dll->streamin->get64bitsLE((U8*)&(laszip->offset_to_special_evlrs)); } catch(...)
            {
              sprintf(laszip_dll->error, "reading offset_to_special_evlrs %d", (I32)laszip->offset_to_special_evlrs);
              return 1;
            }
            try { laszip_dll->streamin->get16bitsLE((U8*)&(laszip->num_items)); } catch(...)
            {
              sprintf(laszip_dll->error, "reading num_items %d", (I32)laszip->num_items);
              return 1;
            }
            laszip->items = new LASitem[laszip->num_items];
            U32 j;
            for (j = 0; j < laszip->num_items; j++)
            {
              U16 type;
              try { laszip_dll->streamin->get16bitsLE((U8*)&type); } catch(...)
              {
                sprintf(laszip_dll->error, "reading type of item %u", j);
                return 1;
              }
              laszip->items[j].type = (LASitem::Type)type;
              try { laszip_dll->streamin->get16bitsLE((U8*)&(laszip->items[j].size)); } catch(...)
              {
                sprintf(laszip_dll->error, "reading size of item %u", j);
                return 1;
              }
              try { laszip_dll->streamin->get16bitsLE((U8*)&(laszip->items[j].version)); } catch(...)
              {
                sprintf(laszip_dll->error, "reading version of item %u", j);
                return 1;
              }
            }
          }
          else
          {
            laszip_dll->header.vlrs[i].data = new U8[laszip_dll->header.vlrs[i].record_length_after_header];

            try { laszip_dll->streamin->getBytes(laszip_dll->header.vlrs[i].data, laszip_dll->header.vlrs[i].record_length_after_header); } catch(...)
            {
              sprintf(laszip_dll->error, "reading %d bytes of data into header.vlrs[%u].data", (I32)laszip_dll->header.vlrs[i].record_length_after_header, i);
              return 1;
            }
          }
        }
        else
        {
          laszip_dll->header.vlrs[i].data = 0;
        }

        // keep track on the number of bytes we have read so far

        vlrs_size += laszip_dll->header.vlrs[i].record_length_after_header;

        // special handling for LASzip VLR

        if ((strcmp(laszip_dll->header.vlrs[i].user_id, "laszip encoded") == 0) && (laszip_dll->header.vlrs[i].record_id == 22204))
        {
          // we take our the VLR for LASzip away
          laszip_dll->header.offset_to_point_data -= (54+laszip_dll->header.vlrs[i].record_length_after_header);
          vlrs_size -= (54+laszip_dll->header.vlrs[i].record_length_after_header);
          i--;
          laszip_dll->header.number_of_variable_length_records--;
          // free or resize the VLR array
          if (laszip_dll->header.number_of_variable_length_records == 0)
          {
            free(laszip_dll->header.vlrs);
            laszip_dll->header.vlrs = 0;
          }
          else
          {
            laszip_dll->header.vlrs = (laszip_vlr*)realloc(laszip_dll->header.vlrs, sizeof(laszip_vlr)*laszip_dll->header.number_of_variable_length_records);
          }
        }
      }
    }

    // load any number of user-defined bytes that might have been added after the header

    laszip_dll->header.user_data_after_header_size = (I32)laszip_dll->header.offset_to_point_data - vlrs_size - laszip_dll->header.header_size;
    if (laszip_dll->header.user_data_after_header_size)
    {
      if (laszip_dll->header.user_data_after_header)
      {
        delete [] laszip_dll->header.user_data_after_header;
      }
      laszip_dll->header.user_data_after_header = new U8[laszip_dll->header.user_data_after_header_size];

      try { laszip_dll->streamin->getBytes((U8*)laszip_dll->header.user_data_after_header, laszip_dll->header.user_data_after_header_size); } catch(...)
      {
        sprintf(laszip_dll->error, "reading %u bytes of data into header.user_data_after_header", laszip_dll->header.user_data_after_header_size);
        return 1;
      }
    }

    // remove extra bits in point data type

    if ((laszip_dll->header.point_data_format & 128) || (laszip_dll->header.point_data_format & 64))
    {
      if (!laszip)
      {
        sprintf(laszip_dll->error, "this file was compressed with an experimental version of LASzip. contact 'martin.isenburg@rapidlasso.com' for assistance");
        return 1;
      }
      laszip_dll->header.point_data_format &= 127;
    }

    // check if file is compressed

    if (laszip)
    {
      // yes. check the compressor state
      *is_compressed = 1;
      if (!laszip->check())
      {
        sprintf(laszip_dll->error, "%s upgrade to the latest release of LASzip or contact 'martin.isenburg@rapidlasso.com' for assistance", laszip->get_error());
        return 1;
      }
    }
    else
    {
      // no. setup an un-compressed read
      *is_compressed = 0;
      laszip = new LASzip;
      if (laszip == 0)
      {
        sprintf(laszip_dll->error, "could not alloc LASzip");
        return 1;
      }
      if (!laszip->setup(laszip_dll->header.point_data_format, laszip_dll->header.point_data_record_length, LASZIP_COMPRESSOR_NONE))
      {
        sprintf(laszip_dll->error, "invalid combination of point_data_format %d and point_data_record_length %d", (I32)laszip_dll->header.point_data_format, (I32)laszip_dll->header.point_data_record_length);
        return 1;
      }
    }

    // create point's item pointers

    laszip_dll->point_items = new U8*[laszip->num_items];

    if (laszip_dll->point_items == 0)
    {
      sprintf(laszip_dll->error, "could not alloc point_items");
      return 1;
    }

    for (i = 0; i < laszip->num_items; i++)
    {
      switch (laszip->items[i].type)
      {
      case LASitem::POINT14:
      case LASitem::POINT10:
        laszip_dll->point_items[i] = (U8*)&(laszip_dll->point.X);
        break;
      case LASitem::GPSTIME11:
        laszip_dll->point_items[i] = (U8*)&(laszip_dll->point.gps_time);
        break;
      case LASitem::RGBNIR14:
      case LASitem::RGB12:
        laszip_dll->point_items[i] = (U8*)laszip_dll->point.rgb;
        break;
      case LASitem::WAVEPACKET13:
        laszip_dll->point_items[i] = (U8*)&(laszip_dll->point.wave_packet);
        break;
      case LASitem::BYTE:
        laszip_dll->point.num_extra_bytes = laszip->items[i].size;
        if (laszip_dll->point.extra_bytes) delete [] laszip_dll->point.extra_bytes;
        laszip_dll->point.extra_bytes = new U8[laszip_dll->point.num_extra_bytes];
        laszip_dll->point_items[i] = laszip_dll->point.extra_bytes;
        break;
      default:
        sprintf(laszip_dll->error, "unknown LASitem type %d", (I32)laszip->items[i].type);
        return 1;
      }
    }

    // did the user request to recode the compatibility mode points?

    laszip_dll->compatibility_mode = FALSE;

    if (laszip_dll->request_compatibility_mode && (laszip_dll->header.version_minor < 4))
    {
      // does this file contain compatibility mode recoded LAS 1.4 content

      struct laszip_vlr* compatibility_VLR = 0;

      if (laszip_dll->header.point_data_format == 1 || laszip_dll->header.point_data_format == 3 || laszip_dll->header.point_data_format == 5)
      {
        // if we find the compatibility VLR

        for (i = 0; i < laszip_dll->header.number_of_variable_length_records; i++)
        {
          if ((strncmp(laszip_dll->header.vlrs[i].user_id, "lascompatible\0\0", 16) == 0) && (laszip_dll->header.vlrs[i].record_id == 22204))
          {
            if (laszip_dll->header.vlrs[i].record_length_after_header == 2+2+4+148)
            {
              compatibility_VLR = &(laszip_dll->header.vlrs[i]);
              break;
            }
          }
        }

        if (compatibility_VLR)
        {
          // and we also find the extra bytes VLR with the right attributes

          LASattributer attributer;
          for (i = 0; i < laszip_dll->header.number_of_variable_length_records; i++)
          {
            if ((strncmp(laszip_dll->header.vlrs[i].user_id, "LASF_Spec\0\0\0\0\0\0", 16) == 0) && (laszip_dll->header.vlrs[i].record_id == 4))
            {
              attributer.init_attributes(laszip_dll->header.vlrs[i].record_length_after_header/192, (LASattribute*)laszip_dll->header.vlrs[i].data);
              laszip_dll->start_scan_angle = attributer.get_attribute_start("LAS 1.4 scan angle");
              laszip_dll->start_extended_returns = attributer.get_attribute_start("LAS 1.4 extended returns");
              laszip_dll->start_classification = attributer.get_attribute_start("LAS 1.4 classification");
              laszip_dll->start_flags_and_channel = attributer.get_attribute_start("LAS 1.4 flags and channel");
              laszip_dll->start_NIR_band = attributer.get_attribute_start("LAS 1.4 NIR band");
              break;
            }
          }

          // can we do it ... ?

          if ((laszip_dll->start_scan_angle != -1) && (laszip_dll->start_extended_returns != -1) && (laszip_dll->start_classification != -1) && (laszip_dll->start_flags_and_channel != -1))
          {
            // yes ... so let's fix the header (using the content from the compatibility VLR)

            ByteStreamInArray* in;
            in = new ByteStreamInArrayLE(compatibility_VLR->data, compatibility_VLR->record_length_after_header);
            // read control info
            U16 laszip_version;
            in->get16bitsLE((U8*)&laszip_version);
            U16 compatible_version;
            in->get16bitsLE((U8*)&compatible_version);
            U32 unused;
            in->get32bitsLE((U8*)&unused);
            // read the 148 bytes of the extended LAS 1.4 header
            U64 start_of_waveform_data_packet_record;
            in->get64bitsLE((U8*)&start_of_waveform_data_packet_record);
            if (start_of_waveform_data_packet_record != 0)
            {
#ifdef _WIN32
              fprintf(stderr,"WARNING: start_of_waveform_data_packet_record is %I64d. reading 0 instead.\n", start_of_waveform_data_packet_record);
#else
              fprintf(stderr,"WARNING: start_of_waveform_data_packet_record is %lld. reading 0 instead.\n", start_of_waveform_data_packet_record);
#endif
            }
            laszip_dll->header.start_of_waveform_data_packet_record = 0;
            U64 start_of_first_extended_variable_length_record;
            in->get64bitsLE((U8*)&start_of_first_extended_variable_length_record);
            if (start_of_first_extended_variable_length_record != 0)
            {
#ifdef _WIN32
              fprintf(stderr,"WARNING: EVLRs not supported. start_of_first_extended_variable_length_record is %I64d. reading 0 instead.\n", start_of_first_extended_variable_length_record);
#else
              fprintf(stderr,"WARNING: EVLRs not supported. start_of_first_extended_variable_length_record is %lld. reading 0 instead.\n", start_of_first_extended_variable_length_record);
#endif
            }
            laszip_dll->header.start_of_first_extended_variable_length_record = 0;
            U32 number_of_extended_variable_length_records ;
            in->get32bitsLE((U8*)&number_of_extended_variable_length_records);
            if (number_of_extended_variable_length_records != 0)
            {
              fprintf(stderr,"WARNING: EVLRs not supported. number_of_extended_variable_length_records is %u. reading 0 instead.\n", number_of_extended_variable_length_records);
            }
            laszip_dll->header.number_of_extended_variable_length_records = 0;
            U64 extended_number_of_point_records = 0;
            in->get64bitsLE((U8*)&extended_number_of_point_records);
            if (((U64)(laszip_dll->header.number_of_point_records)) != extended_number_of_point_records)
            {
#ifdef _WIN32
              fprintf(stderr,"WARNING: number_of_point_records is %u. but extended_number_of_point_records is %I64d.\n", laszip_dll->header.number_of_point_records, number_of_extended_variable_length_records);
#else
              fprintf(stderr,"WARNING: number_of_point_records is %u. but extended_number_of_point_records is %u.\n", laszip_dll->header.number_of_point_records, number_of_extended_variable_length_records);
#endif
            }
            laszip_dll->header.extended_number_of_point_records = extended_number_of_point_records;
            U64 extended_number_of_points_by_return;
            for (U32 i = 0; i < 15; i++)
            {
              in->get64bitsLE((U8*)&extended_number_of_points_by_return);
              if ((i < 5) && ((U64)(laszip_dll->header.number_of_points_by_return[i])) != extended_number_of_points_by_return)
              {
#ifdef _WIN32
                fprintf(stderr,"WARNING: number_of_points_by_return[%d] is %u. but extended_number_of_points_by_return[%llu] is %I64d.\n", laszip_dll->header.number_of_points_by_return[i], i, extended_number_of_points_by_return, i);
#else
                fprintf(stderr,"WARNING: number_of_points_by_return[%d] is %u. but extended_number_of_points_by_return[%llu] is %u.\n", laszip_dll->header.number_of_points_by_return[i], i, extended_number_of_points_by_return, i);
#endif
              }
              laszip_dll->header.extended_number_of_points_by_return[i] = extended_number_of_points_by_return;
            }
            delete in;

            // remove the compatibility VLR

            if (laszip_remove_vlr(laszip_dll, "lascompatible\0\0", 22204))
            {
              sprintf(laszip_dll->error, "removing the compatibility VLR");
              return 1;
            }

            // remove the LAS 1.4 attributes from the "extra bytes" description

            if (laszip_dll->start_NIR_band != -1) attributer.remove_attribute("LAS 1.4 NIR band");
            attributer.remove_attribute("LAS 1.4 flags and channel");
            attributer.remove_attribute("LAS 1.4 classification");
            attributer.remove_attribute("LAS 1.4 extended returns");
            attributer.remove_attribute("LAS 1.4 scan angle");

            // either rewrite or remove the "extra bytes" VLR

            if (attributer.number_attributes)
            {
              if (laszip_add_vlr(laszip_dll, "LASF_Spec\0\0\0\0\0\0", 4, attributer.number_attributes*sizeof(LASattribute), 0, (uint8_t*)attributer.attributes))
              {
                sprintf(laszip_dll->error, "rewriting the extra bytes VLR without 'LAS 1.4 compatibility mode' attributes");
                return 1;
              }
            }
            else
            {
              if (laszip_remove_vlr(laszip_dll, "LASF_Spec\0\0\0\0\0\0", 4))
              {
                sprintf(laszip_dll->error, "removing the LAS 1.4 attribute VLR");
                return 1;
              }
            }

            // upgrade to LAS 1.4
            if (laszip_dll->header.version_minor < 3)
            {
              // LAS 1.2 header is 148 bytes less than LAS 1.4+ header
              laszip_dll->header.header_size += 148;
              laszip_dll->header.offset_to_point_data += 148;
            }
            else
            {
              // LAS 1.3 header is 140 bytes less than LAS 1.4+ header
              laszip_dll->header.header_size += 148;
              laszip_dll->header.offset_to_point_data += 140;
            }
            laszip_dll->header.version_minor = 4;

            // maybe turn on the bit indicating the presence of the OGC WKT
            for (i = 0; i < laszip_dll->header.number_of_variable_length_records; i++)
            {
              if ((strncmp(laszip_dll->header.vlrs[i].user_id, "LASF_Projection", 16) == 0) && (laszip_dll->header.vlrs[i].record_id == 2112))
              {
                laszip_dll->header.global_encoding |= (1<<4);
                break;
              }
            }

            // update point type and size

            laszip_dll->point.extended_point_type = 1;

            if (laszip_dll->header.point_data_format == 1)
            {
              laszip_dll->header.point_data_format = 6;
              laszip_dll->header.point_data_record_length += (2 - 5); // record is 2 bytes larger but minus 5 extra bytes
            }
            else if (laszip_dll->header.point_data_format == 3)
            {
              if (laszip_dll->start_NIR_band == -1)
              {
                laszip_dll->header.point_data_format = 7;
                laszip_dll->header.point_data_record_length += (2 - 5); // record is 2 bytes larger but minus 5 extra bytes
              }
              else
              {
                laszip_dll->header.point_data_format = 8;
                laszip_dll->header.point_data_record_length += (4 - 7); // record is 4 bytes larger but minus 7 extra bytes
              }
            }
            else
            {
              if (laszip_dll->start_NIR_band == -1)
              {
                laszip_dll->header.point_data_format = 9;
                laszip_dll->header.point_data_record_length += (2 - 5);
              }
              else
              {
                laszip_dll->header.point_data_format = 10;
                laszip_dll->header.point_data_record_length += (4 - 7);
              }
            }

            // we are operating in compatibility mode

            laszip_dll->compatibility_mode = TRUE;
          }
        }
      }
    }

    // create the point reader

    laszip_dll->reader = new LASreadPoint();
    if (laszip_dll->reader == 0)
    {
      sprintf(laszip_dll->error, "could not alloc LASreadPoint");
      return 1;
    }

    if (!laszip_dll->reader->setup(laszip->num_items, laszip->items, laszip))
    {
      sprintf(laszip_dll->error, "setup of LASreadPoint failed");
      return 1;
    }

    if (!laszip_dll->reader->init(laszip_dll->streamin))
    {
      sprintf(laszip_dll->error, "init of LASreadPoint failed");
      return 1;
    }

    delete laszip;

    // set the point number and point count

    laszip_dll->npoints = (laszip_dll->header.number_of_point_records ? laszip_dll->header.number_of_point_records : laszip_dll->header.extended_number_of_point_records);
    laszip_dll->p_count = 0;
  }
  catch (...)
  {
    sprintf(laszip_dll->error, "internal error in laszip_open_reader");
    return 1;
  }

  laszip_dll->error[0] = '\0';
  return 0;
}

int32_t
laszip_seek_point(
    laszip_dll_struct *                     laszip_dll
    , uint64_t                       index
)
{
  if (laszip_dll == 0) return 1;

  try
  {
    // seek to the point
    if (!laszip_dll->reader->seek((U32)laszip_dll->p_count, (U32)index))
    {
#ifdef _WIN32
      sprintf(laszip_dll->error, "seeking from index %I64d to index %I64d for file with %I64d points", laszip_dll->p_count, index, laszip_dll->npoints);
#else
      sprintf(laszip_dll->error, "seeking from index %lld to index %lld for file with %lld points", laszip_dll->p_count, index, laszip_dll->npoints);
#endif
      return 1;
    }
    laszip_dll->p_count = index;
  }
  catch (...)
  {
    sprintf(laszip_dll->error, "internal error in laszip_seek_point");
    return 1;
  }

  laszip_dll->error[0] = '\0';
  return 0;
}

int32_t
laszip_read_point(
    laszip_dll_struct *                     laszip_dll
)
{
  if (laszip_dll == 0) return 1;

  try
  {
    // read the point
    if (!laszip_dll->reader->read(laszip_dll->point_items))
    {
#ifdef _WIN32
      sprintf(laszip_dll->error, "reading point %I64d of %I64d total points", laszip_dll->p_count, laszip_dll->npoints);
#else
      sprintf(laszip_dll->error, "reading point %lld of %lld total points", laszip_dll->p_count, laszip_dll->npoints);
#endif
      return 1;
    }

    // special recoding of points (in compatibility mode only)

    if (laszip_dll->compatibility_mode)
    {
      I16 scan_angle_remainder;
      U8 extended_returns;
      U8 classification;
      U8 flags_and_channel;
      I32 return_number_increment;
      I32 number_of_returns_increment;
      I32 overlap_bit;
      I32 scanner_channel;

      // instill extended attributes
      struct laszip_point* point = &laszip_dll->point;

      // get extended attributes from extra bytes
      scan_angle_remainder = *((I16*)(point->extra_bytes + laszip_dll->start_scan_angle));
      extended_returns = point->extra_bytes[laszip_dll->start_extended_returns];
      classification = point->extra_bytes[laszip_dll->start_classification];
      flags_and_channel = point->extra_bytes[laszip_dll->start_flags_and_channel];
      if (laszip_dll->start_NIR_band != -1)
      {
        point->rgb[3] = *((U16*)(point->extra_bytes + laszip_dll->start_NIR_band));
      }

      // decompose into individual attributes
      return_number_increment = (extended_returns >> 4) & 0x0F;
      number_of_returns_increment = extended_returns & 0x0F;
      scanner_channel = (flags_and_channel >> 1) & 0x03;
      overlap_bit = flags_and_channel & 0x01;

      // instill into point
      point->extended_scan_angle = scan_angle_remainder + I16_QUANTIZE(((F32)point->scan_angle_rank) / 0.006f);
      point->extended_return_number = return_number_increment + point->return_number;
      point->extended_number_of_returns = number_of_returns_increment + point->number_of_returns;
      point->extended_classification = classification + point->classification;
      point->extended_scanner_channel = scanner_channel;
      point->extended_classification_flags = (overlap_bit << 3) | (point->classification >> 5);
    }

    laszip_dll->p_count++;
  }
  catch (...)
  {
    sprintf(laszip_dll->error, "internal error in laszip_read_point");
    return 1;
  }

  laszip_dll->error[0] = '\0';
  return 0;
}

int32_t
laszip_read_inside_point(
    laszip_dll_struct *                     laszip_dll
    , laszip_BOOL*                     is_done
)
{
  if (laszip_dll == 0) return 1;

  try
  {
    double xy;

    *is_done = 1;

    while (laszip_dll->reader->read(laszip_dll->point_items))
    {
      laszip_dll->p_count++;
      xy = laszip_dll->header.x_scale_factor*laszip_dll->point.X+laszip_dll->header.x_offset;
      if (xy < laszip_dll->lax_r_min_x || xy >= laszip_dll->lax_r_max_x) continue;
      xy = laszip_dll->header.y_scale_factor*laszip_dll->point.Y+laszip_dll->header.y_offset;
      if (xy < laszip_dll->lax_r_min_y || xy >= laszip_dll->lax_r_max_y) continue;
      *is_done = 0;
      break;
    }

    if (*is_done)
    {
      if (laszip_dll->p_count < laszip_dll->npoints)
      {
#ifdef _WIN32
        sprintf(laszip_dll->error, "reading point %I64d of %I64d total points", laszip_dll->p_count, laszip_dll->npoints);
#else
        sprintf(laszip_dll->error, "reading point %lld of %lld total points", laszip_dll->p_count, laszip_dll->npoints);
#endif
        return 1;
      }
    }
  }
  catch (...)
  {
    sprintf(laszip_dll->error, "internal error in laszip_read_inside_point");
    return 1;
  }

  laszip_dll->error[0] = '\0';
  return 0;
}

int32_t
laszip_close_reader(
    laszip_dll_struct *                     laszip_dll
)
{
  if (laszip_dll == 0) return 1;

  try
  {
    if (laszip_dll->reader == 0)
    {
      sprintf(laszip_dll->error, "closing reader before it was opened");
      return 1;
    }

    if (!laszip_dll->reader->done())
    {
      sprintf(laszip_dll->error, "done of LASreadPoint failed");
      return 1;
    }

    delete laszip_dll->reader;
    laszip_dll->reader = 0;

    delete [] laszip_dll->point_items;
    laszip_dll->point_items = 0;

    delete laszip_dll->streamin;
    laszip_dll->streamin = 0;

    fclose(laszip_dll->file);
    laszip_dll->file = 0;
  }
  catch (...)
  {
    sprintf(laszip_dll->error, "internal error in laszip_close_reader");
    return 1;
  }

  laszip_dll->error[0] = '\0';
  return 0;
}
