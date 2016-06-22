
#ifndef BYTE_STREAM_IN_HPP
#define BYTE_STREAM_IN_HPP

#include "mydefs.hpp"

class ByteStreamIn
{
public:
/* write single bits                                         */
  inline U32 getBits(U32 num_bits)
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
/* read a single byte                                        */
  virtual U32 getByte() = 0;
/* read an array of bytes                                    */
  virtual void getBytes(U8* bytes, const U32 num_bytes) = 0;
/* read 16 bit low-endian field                              */
  virtual void get16bitsLE(U8* bytes) = 0;
/* read 32 bit low-endian field                              */
  virtual void get32bitsLE(U8* bytes) = 0;
/* read 64 bit low-endian field                              */
  virtual void get64bitsLE(U8* bytes) = 0;
/* is the stream seekable (e.g. stdin is not)                */
  virtual BOOL isSeekable() const = 0;
/* get current position of stream                            */
  virtual I64 tell() const = 0;
/* seek to this position in the stream                       */
  virtual BOOL seek(const I64 position) = 0;
/* seek to the end of the file                               */
  virtual BOOL seekEnd(const I64 distance=0) = 0;
/* constructor                                               */
  inline ByteStreamIn() { bit_buffer = 0; num_buffer = 0; };
/* destructor                                                */
  virtual ~ByteStreamIn() {};
private:
  U64 bit_buffer;
  U32 num_buffer;
};

#endif
