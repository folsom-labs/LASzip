/*
===============================================================================

  FILE:  laswriteitemraw.hpp

  CONTENTS:

    Implementation of LASwriteItemRaw for *all* items that compose a point.

  PROGRAMMERS:

    martin.isenburg@rapidlasso.com  -  http://rapidlasso.com

  COPYRIGHT:

    (c) 2007-2013, martin isenburg, rapidlasso - fast tools to catch reality

    This is free software; you can redistribute and/or modify it under the
    terms of the GNU Lesser General Licence as published by the Free Software
    Foundation. See the COPYING file for more information.

    This software is distributed WITHOUT ANY WARRANTY and without even the
    implied warranty of MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.

  CHANGE HISTORY:

    10 January 2011 -- licensing change for LGPL release and liblas integration
    7 January 2011 -- introduced swap buffers to reduce number of fwrite calls
    12 December 2010 -- refactored after watching two movies with silke

===============================================================================
*/
#ifndef LAS_WRITE_ITEM_RAW_HPP
#define LAS_WRITE_ITEM_RAW_HPP

#include "laswriteitem.hpp"

#include <assert.h>

class LASwriteItemRaw_POINT10_LE : public LASwriteItemRaw
{
public:
  LASwriteItemRaw_POINT10_LE(){};
  inline BOOL write(const U8* item)
  {
    return outstream->putBytes(item, 20);
  };
};

class LASwriteItemRaw_GPSTIME11_LE : public LASwriteItemRaw
{
public:
  LASwriteItemRaw_GPSTIME11_LE() {};
  inline BOOL write(const U8* item)
  {
    return outstream->putBytes(item, 8);
  };
};

class LASwriteItemRaw_RGB12_LE : public LASwriteItemRaw
{
public:
  LASwriteItemRaw_RGB12_LE(){}
  inline BOOL write(const U8* item)
  {
    return outstream->putBytes(item, 6);
  };
};

class LASwriteItemRaw_WAVEPACKET13_LE : public LASwriteItemRaw
{
public:
  LASwriteItemRaw_WAVEPACKET13_LE(){}
  inline BOOL write(const U8* item)
  {
    return outstream->putBytes(item, 29);
  };
};

class LASwriteItemRaw_BYTE : public LASwriteItemRaw
{
public:
  LASwriteItemRaw_BYTE(U32 number)
  {
    this->number = number;
  }
  inline BOOL write(const U8* item)
  {
    return outstream->putBytes(item, number);
  };
private:
  U32 number;
};

class LAStempWritePoint10
{
public:
  I32 X;
  I32 Y;
  I32 Z;
  U16 intensity;
  U8 return_number : 3;
  U8 number_of_returns : 3;
  U8 scan_direction_flag : 1;
  U8 edge_of_flight_line : 1;
  U8 classification;
  I8 scan_angle_rank;
  U8 user_data;
  U16 point_source_ID;

  // LAS 1.4 only
  I16 extended_scan_angle;
  U8 extended_point_type : 2;
  U8 extended_scanner_channel : 2;
  U8 extended_classification_flags : 4;
  U8 extended_classification;
  U8 extended_return_number : 4;
  U8 extended_number_of_returns : 4;

  // for 8 byte alignment of the GPS time
  U8 dummy[3];

  // LASlib only
  U32 deleted_flag;

  F64 gps_time;
};

class LAStempWritePoint14
{
public:
  I32 X;
  I32 Y;
  I32 Z;
  U16 intensity;
  U8 return_number : 4;
  U8 number_of_returns : 4;
  U8 classification_flags : 4;
  U8 scanner_channel : 2;
  U8 scan_direction_flag : 1;
  U8 edge_of_flight_line : 1;
  U8 classification;
  U8 user_data;
  I16 scan_angle;
  U16 point_source_ID;
};

class LASwriteItemRaw_POINT14_LE : public LASwriteItemRaw
{
public:
  LASwriteItemRaw_POINT14_LE(){};
  inline BOOL write(const U8* item)
  {
    ((LAStempWritePoint14*)buffer)->X = ((LAStempWritePoint10*)item)->X;
    ((LAStempWritePoint14*)buffer)->Y = ((LAStempWritePoint10*)item)->Y;
    ((LAStempWritePoint14*)buffer)->Z = ((LAStempWritePoint10*)item)->Z;
    ((LAStempWritePoint14*)buffer)->intensity = ((LAStempWritePoint10*)item)->intensity;
    ((LAStempWritePoint14*)buffer)->scan_direction_flag = ((LAStempWritePoint10*)item)->scan_direction_flag;
    ((LAStempWritePoint14*)buffer)->edge_of_flight_line = ((LAStempWritePoint10*)item)->edge_of_flight_line;
    ((LAStempWritePoint14*)buffer)->classification = (((LAStempWritePoint10*)item)->classification & 31);
    ((LAStempWritePoint14*)buffer)->user_data = ((LAStempWritePoint10*)item)->user_data;
    ((LAStempWritePoint14*)buffer)->point_source_ID = ((LAStempWritePoint10*)item)->point_source_ID;

    if (((LAStempWritePoint10*)item)->extended_point_type)
    {
      ((LAStempWritePoint14*)buffer)->classification_flags = (((LAStempWritePoint10*)item)->extended_classification_flags & 8) | (((LAStempWritePoint10*)item)->classification >> 5);
      if (((LAStempWritePoint10*)item)->extended_classification > 31) ((LAStempWritePoint14*)buffer)->classification = ((LAStempWritePoint10*)item)->extended_classification;
      ((LAStempWritePoint14*)buffer)->scanner_channel = ((LAStempWritePoint10*)item)->extended_scanner_channel;
      ((LAStempWritePoint14*)buffer)->return_number = ((LAStempWritePoint10*)item)->extended_return_number;
      ((LAStempWritePoint14*)buffer)->number_of_returns = ((LAStempWritePoint10*)item)->extended_number_of_returns;
      ((LAStempWritePoint14*)buffer)->scan_angle = ((LAStempWritePoint10*)item)->extended_scan_angle;
    }
    else
    {
      ((LAStempWritePoint14*)buffer)->classification_flags = (((LAStempWritePoint10*)item)->classification >> 5);
      ((LAStempWritePoint14*)buffer)->scanner_channel = 0;
      ((LAStempWritePoint14*)buffer)->return_number = ((LAStempWritePoint10*)item)->return_number;
      ((LAStempWritePoint14*)buffer)->number_of_returns = ((LAStempWritePoint10*)item)->number_of_returns;
      ((LAStempWritePoint14*)buffer)->scan_angle = I16_QUANTIZE(((LAStempWritePoint10*)item)->scan_angle_rank/0.006f);
    }

    *((F64*)&buffer[22]) = ((LAStempWritePoint10*)item)->gps_time;
    return outstream->putBytes(buffer, 30);
  }
private:
  U8 buffer[30];
};

class LASwriteItemRaw_RGBNIR14_LE : public LASwriteItemRaw
{
public:
  LASwriteItemRaw_RGBNIR14_LE(){}
  inline BOOL write(const U8* item)
  {
    return outstream->putBytes(item, 8);
  };
};

#endif
