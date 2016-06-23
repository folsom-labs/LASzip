
#ifndef BYTE_STREAM_IN_HPP
#define BYTE_STREAM_IN_HPP

#include "mydefs.hpp"
#include <stdio.h>

#if defined(_MSC_VER) && (_MSC_VER < 1300)
extern "C" __int64 _cdecl _ftelli64(FILE*);
extern "C" int _cdecl _fseeki64(FILE*, __int64, int);
#endif

class ByteStreamIn
{
public:
  ByteStreamIn(FILE *f);
  ~ByteStreamIn() {};

  U32 getBits(U32 num_bits);
  U32 getByte();
  void getBytes(U8* bytes, const U32 num_bytes);
  void get16bitsLE(U8* bytes);
  void get32bitsLE(U8* bytes);
  void get64bitsLE(U8* bytes);
  BOOL isSeekable() const;
  I64 tell() const;
  BOOL seek(const I64 position);
  BOOL seekEnd(const I64 distance=0);
private:
  U64 bit_buffer;
  U32 num_buffer;

  FILE *file;
};

inline ByteStreamIn::ByteStreamIn(FILE *file) {
  this->file = file;
  bit_buffer = 0; num_buffer = 0;
};

inline U32 ByteStreamIn::getBits(U32 num_bits)
{
  if (num_buffer < num_bits)
  {
    U32 input_bits;
    get32bitsLE((U8*)&input_bits);
    bit_buffer = bit_buffer | (((U64)input_bits) << num_buffer);
    num_buffer = num_buffer + 32;
  }
  U32 new_bits = (U32)(bit_buffer & ((1 << num_bits) - 1));
  bit_buffer = bit_buffer >> num_bits;
  num_buffer = num_buffer - num_bits;
  return new_bits;
};

inline U32 ByteStreamIn::getByte()
{
  int byte = getc(file);
  if (byte == EOF)
  {
    throw EOF;
  }
  return (U32)byte;
}

inline void ByteStreamIn::getBytes(U8* bytes, const U32 num_bytes)
{
  if (fread(bytes, 1, num_bytes, file) != num_bytes)
  {
    throw EOF;
  }
}

inline BOOL ByteStreamIn::isSeekable() const
{
  return (file != stdin);
}

inline I64 ByteStreamIn::tell() const
{
#if defined _WIN32 && ! defined (__MINGW32__)
  return _ftelli64(file);
#elif defined (__MINGW32__)
  return (I64)ftello64(file);
#else
  return (I64)ftello(file);
#endif
}

inline BOOL ByteStreamIn::seek(const I64 position)
{
  if (tell() != position)
  {
#if defined _WIN32 && ! defined (__MINGW32__)
    return !(_fseeki64(file, position, SEEK_SET));
#elif defined (__MINGW32__)
    return !(fseeko64(file, (off_t)position, SEEK_SET));
#else
    return !(fseeko(file, (off_t)position, SEEK_SET));
#endif
  }
  return TRUE;
}

inline BOOL ByteStreamIn::seekEnd(const I64 distance)
{
#if defined _WIN32 && ! defined (__MINGW32__)
  return !(_fseeki64(file, -distance, SEEK_END));
#elif defined (__MINGW32__)
  return !(fseeko64(file, (off_t)-distance, SEEK_END));
#else
  return !(fseeko(file, (off_t)-distance, SEEK_END));
#endif
}

inline void ByteStreamIn::get16bitsLE(U8* bytes)
{
  getBytes(bytes, 2);
}

inline void ByteStreamIn::get32bitsLE(U8* bytes)
{
  getBytes(bytes, 4);
}

inline void ByteStreamIn::get64bitsLE(U8* bytes)
{
  getBytes(bytes, 8);
}

#endif
