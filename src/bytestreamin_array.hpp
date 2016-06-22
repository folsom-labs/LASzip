#ifndef BYTE_STREAM_IN_ARRAY_H
#define BYTE_STREAM_IN_ARRAY_H

#include "bytestreamin.hpp"

class ByteStreamInArray : public ByteStreamIn
{
public:
  ByteStreamInArray(U8* data, I64 size);
/* read a single byte                                        */
  U32 getByte();
/* read an array of bytes                                    */
  void getBytes(U8* bytes, const U32 num_bytes);
/* is the stream seekable (e.g. stdin is not)                */
  BOOL isSeekable() const;
/* get current position of stream                            */
  I64 tell() const;
/* seek to this position in the stream                       */
  BOOL seek(const I64 position);
/* seek to the end of the stream                             */
  BOOL seekEnd(const I64 distance=0);
/* destructor                                                */
  ~ByteStreamInArray(){};
protected:
  U8* data;
  I64 size;
  I64 curr;
};

class ByteStreamInArrayLE : public ByteStreamInArray
{
public:
  ByteStreamInArrayLE(U8* data, I64 size);
/* read 16 bit low-endian field                              */
  void get16bitsLE(U8* bytes);
/* read 32 bit low-endian field                              */
  void get32bitsLE(U8* bytes);
/* read 64 bit low-endian field                              */
  void get64bitsLE(U8* bytes);
};

inline ByteStreamInArray::ByteStreamInArray(U8* data, I64 size)
{
  this->data = data;
  this->size = size;
  this->curr = 0;
}

inline U32 ByteStreamInArray::getByte()
{
  if (curr == size)
  {
    throw EOF;
  }
  U32 byte = data[curr];
  curr++;
  return byte;
}

inline void ByteStreamInArray::getBytes(U8* bytes, const U32 num_bytes)
{
  if ((curr + num_bytes) > size)
  {
    throw EOF;
  }
  memcpy((void*)bytes, (void*)(data+curr), num_bytes);
  curr += num_bytes;
}

inline BOOL ByteStreamInArray::isSeekable() const
{
  return TRUE;
}

inline I64 ByteStreamInArray::tell() const
{
  return curr;
}

inline BOOL ByteStreamInArray::seek(const I64 position)
{
  if ((0 <= position) && (position <= size))
  {
    curr = position;
    return TRUE;
  }
  return FALSE;
}

inline BOOL ByteStreamInArray::seekEnd(const I64 distance)
{
  if ((0 <= distance) && (distance <= size))
  {
    curr = size - distance;
    return TRUE;
  }
  return FALSE;
}

inline ByteStreamInArrayLE::ByteStreamInArrayLE(U8* data, I64 size) : ByteStreamInArray(data, size)
{
}

inline void ByteStreamInArrayLE::get16bitsLE(U8* bytes)
{
  getBytes(bytes, 2);
}

inline void ByteStreamInArrayLE::get32bitsLE(U8* bytes)
{
  getBytes(bytes, 4);
}

inline void ByteStreamInArrayLE::get64bitsLE(U8* bytes)
{
  getBytes(bytes, 8);
}

#endif
