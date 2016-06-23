#ifndef LAS_READ_ITEM_RAW_HPP
#define LAS_READ_ITEM_RAW_HPP

#include "lasreaditem.hpp"

#include <assert.h>

class LASreadItemRaw_POINT10_LE : public LASreadItemRaw
{
public:
  LASreadItemRaw_POINT10_LE(){};
  inline void read(U8* item)
  {
    instream->getBytes(item, 20);
  }
};

class LASreadItemRaw_GPSTIME11_LE : public LASreadItemRaw
{
public:
  LASreadItemRaw_GPSTIME11_LE(){};
  inline void read(U8* item)
  {
    instream->getBytes(item, 8);
  };
};

class LASreadItemRaw_RGB12_LE : public LASreadItemRaw
{
public:
  LASreadItemRaw_RGB12_LE(){};
  inline void read(U8* item)
  {
    instream->getBytes(item, 6);
  };
};

class LASreadItemRaw_WAVEPACKET13_LE : public LASreadItemRaw
{
public:
  LASreadItemRaw_WAVEPACKET13_LE(){}
  inline void read(U8* item)
  {
    instream->getBytes(item, 29);
  };
};

class LASreadItemRaw_BYTE : public LASreadItemRaw
{
public:
  LASreadItemRaw_BYTE(U32 number)
  {
    this->number = number;
  }
  inline void read(U8* item)
  {
    instream->getBytes(item, number);
  };
private:
  U32 number;
};

class LAStempReadPoint10
{
public:
  I32 x;
  I32 y;
  I32 z;
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

class LAStempReadPoint14
{
public:
  I32 x;
  I32 y;
  I32 z;
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

class LASreadItemRaw_POINT14_LE : public LASreadItemRaw
{
public:
  LASreadItemRaw_POINT14_LE(){};
  inline void read(U8* item)
  {
    instream->getBytes(buffer, 30);
    LAStempReadPoint10 *p10 = (LAStempReadPoint10*)item;
    LAStempReadPoint14 *p14 = (LAStempReadPoint14*)buffer;
    p10->x = p14->x;
    p10->y = p14->y;
    p10->z = p14->z;
    p10->intensity = p14->intensity;
    if (p14->number_of_returns > 7)
    {
      if (p14->return_number > 6)
      {
        if (p14->return_number >= p14->number_of_returns)
        {
          p10->return_number = 7;
        }
        else
        {
          p10->return_number = 6;
        }
      }
      else
      {
        p10->return_number = p14->return_number;
      }
      p10->number_of_returns = 7;
    }
    else
    {
      p10->return_number = p14->return_number;
      p10->number_of_returns = p14->number_of_returns;
    }
    p10->scan_direction_flag = p14->scan_direction_flag;
    p10->edge_of_flight_line = p14->edge_of_flight_line;
    p10->classification = (p14->classification_flags << 5);
    if (p14->classification < 32) p10->classification |= p14->classification;
    p10->scan_angle_rank = I8_CLAMP(I16_QUANTIZE(0.006f*p14->scan_angle));
    p10->user_data = p14->user_data;
    p10->point_source_ID = p14->point_source_ID;
    p10->extended_scanner_channel = p14->scanner_channel;
    p10->extended_classification_flags = p14->classification_flags;
    p10->extended_classification = p14->classification;
    p10->extended_return_number = p14->return_number;
    p10->extended_number_of_returns = p14->number_of_returns;
    p10->extended_scan_angle = p14->scan_angle;
    p10->gps_time = *((F64*)&buffer[22]);
  }
private:
  U8 buffer[30];
};

class LASreadItemRaw_RGBNIR14_LE : public LASreadItemRaw
{
public:
  LASreadItemRaw_RGBNIR14_LE(){};
  inline void read(U8* item)
  {
    instream->getBytes(item, 8);
  };
};

#endif
