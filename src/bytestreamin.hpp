
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
  ByteStreamIn() { bit_buffer = 0; num_buffer = 0; };
  virtual ~ByteStreamIn() {};

  U32 getBits(U32 num_bits)
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

  virtual U32 getByte() = 0;
  virtual void getBytes(U8* bytes, const U32 num_bytes) = 0;
  virtual void get16bitsLE(U8* bytes) = 0;
  virtual void get32bitsLE(U8* bytes) = 0;
  virtual void get64bitsLE(U8* bytes) = 0;
  virtual BOOL isSeekable() const = 0;
  virtual I64 tell() const = 0;
  virtual BOOL seek(const I64 position) = 0;
  virtual BOOL seekEnd(const I64 distance=0) = 0;
private:
  U64 bit_buffer;
  U32 num_buffer;
};

class ByteStreamInFileLE : public ByteStreamIn
{
public:
  ByteStreamInFileLE(FILE* file);

  U32 getByte() override;
  void getBytes(U8* bytes, const U32 num_bytes) override;
  void get16bitsLE(U8* bytes) override;
  void get32bitsLE(U8* bytes) override;
  void get64bitsLE(U8* bytes) override;
  virtual BOOL isSeekable() const override;
  virtual I64 tell() const override;
  virtual BOOL seek(const I64 position) override;
  virtual BOOL seekEnd(const I64 distance=0) override;

  FILE* file;
};

inline ByteStreamInFileLE::ByteStreamInFileLE(FILE* file)
{
  this->file = file;
}

inline U32 ByteStreamInFileLE::getByte()
{
  int byte = getc(file);
  if (byte == EOF)
  {
    throw EOF;
  }
  return (U32)byte;
}

inline void ByteStreamInFileLE::getBytes(U8* bytes, const U32 num_bytes)
{
  if (fread(bytes, 1, num_bytes, file) != num_bytes)
  {
    throw EOF;
  }
}

inline BOOL ByteStreamInFileLE::isSeekable() const
{
  return (file != stdin);
}

inline I64 ByteStreamInFileLE::tell() const
{
#if defined _WIN32 && ! defined (__MINGW32__)
  return _ftelli64(file);
#elif defined (__MINGW32__)
  return (I64)ftello64(file);
#else
  return (I64)ftello(file);
#endif
}

inline BOOL ByteStreamInFileLE::seek(const I64 position)
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

inline BOOL ByteStreamInFileLE::seekEnd(const I64 distance)
{
#if defined _WIN32 && ! defined (__MINGW32__)
  return !(_fseeki64(file, -distance, SEEK_END));
#elif defined (__MINGW32__)
  return !(fseeko64(file, (off_t)-distance, SEEK_END));
#else
  return !(fseeko(file, (off_t)-distance, SEEK_END));
#endif
}

inline void ByteStreamInFileLE::get16bitsLE(U8* bytes)
{
  getBytes(bytes, 2);
}

inline void ByteStreamInFileLE::get32bitsLE(U8* bytes)
{
  getBytes(bytes, 4);
}

inline void ByteStreamInFileLE::get64bitsLE(U8* bytes)
{
  getBytes(bytes, 8);
}


#endif
