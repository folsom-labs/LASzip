#ifndef LAS_READ_POINT_HPP
#define LAS_READ_POINT_HPP

#include "mydefs.hpp"
#include "laszip.hpp"
#include "bytestreamin.hpp"

class LASreadItem;
class ArithmeticDecoder;

class LASreadPoint
{
public:
  LASreadPoint();
  ~LASreadPoint();

  // should only be called *once*
  BOOL setup(const U32 num_items, const LASitem* items, const LASzip* laszip=0);

  BOOL init(ByteStreamIn* instream);
  BOOL read(U8* const * point);
  BOOL check_end();
  BOOL done();

  inline const CHAR* error() const { return last_error; };
  inline const CHAR* warning() const { return last_warning; };

private:
  ByteStreamIn* instream;
  U32 num_readers;
  LASreadItem** readers;
  LASreadItem** readers_raw;
  LASreadItem** readers_compressed;
  ArithmeticDecoder* dec;
  // used for chunking
  U32 chunk_size;
  U32 chunk_count;
  U32 current_chunk;
  U32 number_chunks;
  U32 tabled_chunks;
  I64* chunk_starts;
  U32* chunk_totals;
  BOOL init_dec();
  BOOL read_chunk_table();
  U32 search_chunk_table(const U32 index, const U32 lower, const U32 upper);
  // used for seeking
  I64 point_start;
  U32 point_size;
  // used for error and warning reporting
  CHAR* last_error;
  CHAR* last_warning;
};

#endif
