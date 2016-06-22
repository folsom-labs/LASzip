#ifndef LAS_INDEX_HPP
#define LAS_INDEX_HPP

#include "mydefs.hpp"

class LASquadtree;
class LASinterval;
class LASreadPoint;
class ByteStreamIn;

class LASindex
{
public:
  LASindex();
  ~LASindex();

  // create spatial index
  void prepare(LASquadtree* spatial, I32 threshold=1000);
  BOOL add(const F64 x, const F64 y, const U32 index);
  void complete(U32 minimum_points=100000, I32 maximum_intervals=-1, const BOOL verbose=TRUE);

  // read from file or write to file
  BOOL read(const char* file_name);
  BOOL read(ByteStreamIn* stream);

  // intersect
  BOOL intersect_rectangle(const F64 r_min_x, const F64 r_min_y, const F64 r_max_x, const F64 r_max_y);
  BOOL intersect_tile(const F32 ll_x, const F32 ll_y, const F32 size);
  BOOL intersect_circle(const F64 center_x, const F64 center_y, const F64 radius);

  // access the intersected intervals
  BOOL get_intervals();
  BOOL has_intervals();

  U32 start;
  U32 end;
  U32 full;
  U32 total;
  U32 cells;

  // seek to next interval
  BOOL seek_next(LASreadPoint* reader, I64 &p_count);

  // for debugging
  void print(BOOL verbose);

  // for visualization
  LASquadtree* get_spatial() const;
  LASinterval* get_interval() const;

private:
  BOOL merge_intervals();

  LASquadtree* spatial;
  LASinterval* interval;
  BOOL have_interval;
};

#endif
