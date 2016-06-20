/*
===============================================================================

  FILE:  bytestreamin_istream.hpp

  CONTENTS:

    Class for istream-based input streams with endian handling.

  PROGRAMMERS:

    martin.isenburg@rapidlasso.com  -  http://rapidlasso.com

  COPYRIGHT:

    (c) 2007-2012, martin isenburg, rapidlasso - fast tools to catch reality

    This is free software; you can redistribute and/or modify it under the
    terms of the GNU Lesser General Licence as published by the Free Software
    Foundation. See the COPYING file for more information.

    This software is distributed WITHOUT ANY WARRANTY and without even the
    implied warranty of MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.

  CHANGE HISTORY:

     1 October 2011 -- added 64 bit file support in MSVC 6.0 at McCafe at Hbf Linz
    10 January 2011 -- licensing change for LGPL release and liblas integration
    12 December 2010 -- created from ByteStreamOutFile after Howard got pushy (-;

===============================================================================
*/
#ifndef BYTE_STREAM_IN_ISTREAM_H
#define BYTE_STREAM_IN_ISTREAM_H

#include "bytestreamin.hpp"

#ifdef LZ_WIN32_VC6
#include <fstream.h>
#else
#include <istream>
#include <fstream>
using namespace std;
#endif

class ByteStreamInIstream : public ByteStreamIn
{
public:
  ByteStreamInIstream(istream& stream);
/* read a single byte                                        */
  U32 getByte();
/* read an array of bytes                                    */
  void getBytes(U8* bytes, const U32 num_bytes);
/* is the stream seekable (e.g. standard in is not)          */
  BOOL isSeekable() const;
/* get current position of stream                            */
  I64 tell() const;
/* seek to this position in the stream                       */
  BOOL seek(const I64 position);
/* seek to the end of the file                               */
  BOOL seekEnd(const I64 distance=0);
/* destructor                                                */
  ~ByteStreamInIstream(){};
protected:
  istream& stream;
};

class ByteStreamInIstreamLE : public ByteStreamInIstream
{
public:
  ByteStreamInIstreamLE(istream& stream);
/* read 16 bit low-endian field                              */
  void get16bitsLE(U8* bytes);
/* read 32 bit low-endian field                              */
  void get32bitsLE(U8* bytes);
/* read 64 bit low-endian field                              */
  void get64bitsLE(U8* bytes);
};

inline ByteStreamInIstream::ByteStreamInIstream(istream& stream_param) :
  stream(stream_param)
{
}

inline U32 ByteStreamInIstream::getByte()
{
  int byte = stream.get();
  if (stream.eof())
  {
    throw EOF;
  }
  return (U32)byte;
}

inline void ByteStreamInIstream::getBytes(U8* bytes, const U32 num_bytes)
{
  stream.read((char*)bytes, num_bytes);
  if (!stream.good())
  {
    throw EOF;
  }
}

inline BOOL ByteStreamInIstream::isSeekable() const
{
  return !!(static_cast<ifstream&>(stream));
}

inline I64 ByteStreamInIstream::tell() const
{
  return (I64)stream.tellg();
}

inline BOOL ByteStreamInIstream::seek(const I64 position)
{
  if (tell() != position)
  {
    stream.seekg(static_cast<streamoff>(position));
    return stream.good();
  }
  return TRUE;
}

inline BOOL ByteStreamInIstream::seekEnd(const I64 distance)
{
  stream.seekg(static_cast<streamoff>(-distance), ios::end);
  return stream.good();
}

inline ByteStreamInIstreamLE::ByteStreamInIstreamLE(istream& stream) : ByteStreamInIstream(stream)
{
}

inline void ByteStreamInIstreamLE::get16bitsLE(U8* bytes)
{
  getBytes(bytes, 2);
}

inline void ByteStreamInIstreamLE::get32bitsLE(U8* bytes)
{
  getBytes(bytes, 4);
}

inline void ByteStreamInIstreamLE::get64bitsLE(U8* bytes)
{
  getBytes(bytes, 8);
}

#endif
