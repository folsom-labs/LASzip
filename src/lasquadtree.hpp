#ifndef LAS_QUADTREE_HPP
#define LAS_QUADTREE_HPP

#include "mydefs.hpp"

class ByteStreamIn;
class ByteStreamOut;

#define LAS_SPATIAL_QUAD_TREE 0

class LASquadtree
{
public:
  LASquadtree();
  ~LASquadtree();

  // read from file or write to file
  BOOL read(ByteStreamIn* stream);

  // create or finalize the cell (in the spatial hierarchy)
  BOOL manage_cell(const U32 cell_index, const BOOL finalize=FALSE);

  // map points to cells
  BOOL inside(const F64 x, const F64 y) const;
  U32 get_cell_index(const F64 x, const F64 y) const;

  // map cells to coarser cells
  BOOL coarsen(const I32 cell_index, I32* coarser_cell_index, U32* num_cell_indices, I32** cell_indices) const;

  // describe cells
  void get_cell_bounding_box(const I32 cell_index, F32* min, F32* max) const;
  void get_cell_bounding_box(const F64 x, const F64 y, F32* min, F32* max) const;

  // decribe spatial extend
  F64 get_min_x() const { return min_x; };
  F64 get_min_y() const { return min_y; };
  F64 get_max_x() const { return max_x; };
  F64 get_max_y() const { return max_y; };

  // query spatial intersections
  U32 intersect_rectangle(const F64 r_min_x, const F64 r_min_y, const F64 r_max_x, const F64 r_max_y);
  U32 intersect_tile(const F32 ll_x, const F32 ll_y, const F32 size);
  U32 intersect_circle(const F64 center_x, const F64 center_y, const F64 radius);

  // iterate over cells
  BOOL get_all_cells();
  BOOL get_intersected_cells();
  BOOL has_more_cells();

  // for LASquadtree
  BOOL setup(F64 bb_min_x, F64 bb_max_x, F64 bb_min_y, F64 bb_max_y, F32 cell_size = 1000.0f);
  BOOL setup(F64 bb_min_x, F64 bb_max_x, F64 bb_min_y, F64 bb_max_y, F32 cell_size, F32 offset_x, F32 offset_y);
  BOOL tiling_setup(F32 min_x, F32 max_x, F32 min_y, F32 max_y, U32 levels);
  BOOL subtiling_setup(F32 min_x, F32 max_x, F32 min_y, F32 max_y, U32 sub_level, U32 sub_level_index, U32 levels);

  // additional index queries
  U32 get_level_index(const F64 x, const F64 y, U32 level) const;
  U32 get_level_index(const F64 x, const F64 y) const;
  U32 get_level_index(const F64 x, const F64 y, U32 level, F32* min, F32* max) const;
  U32 get_level_index(const F64 x, const F64 y, F32* min, F32* max) const;
  U32 get_cell_index(const F64 x, const F64 y, U32 level) const;

  // additional bounding box queries
  void get_cell_bounding_box(const F64 x, const F64 y, U32 level, F32* min, F32* max) const;
  void get_cell_bounding_box(U32 level_index, U32 level, F32* min, F32* max) const;
  void get_cell_bounding_box(U32 level_index, U32 level, F64* min, F64* max) const;
  void get_cell_bounding_box(U32 level_index, F32* min, F32* max) const;
  void get_cell_bounding_box(U32 level_index, F64* min, F64* max) const;

  // index conversions
  U32 get_level(U32 cell_index) const;

  U32 get_level_index(U32 cell_index, U32 level) const;
  U32 get_level_index(U32 cell_index) const;

  U32 get_cell_index(U32 level_index, U32 level) const;
  U32 get_cell_index(U32 level_index) const;

  // convenience functions
  U32 get_max_level_index(U32 level) const;
  U32 get_max_level_index() const;

  U32 get_max_cell_index(U32 level) const;
  U32 get_max_cell_index() const;

  U32* raster_occupancy(BOOL(*does_cell_exist)(I32), U32 level) const;
  U32* raster_occupancy(BOOL(*does_cell_exist)(I32)) const;

  U32 levels;
  F32 cell_size;
  F32 min_x;
  F32 max_x;
  F32 min_y;
  F32 max_y;
  U32 cells_x;
  U32 cells_y;

  // spatial queries
  U32 intersect_rectangle(const F64 r_min_x, const F64 r_min_y, const F64 r_max_x, const F64 r_max_y, U32 level);
  U32 intersect_tile(const F32 ll_x, const F32 ll_y, const F32 size, U32 level);
  U32 intersect_circle(const F64 center_x, const F64 center_y, const F64 radius, U32 level);

  I32 current_cell;

private:
  U32 sub_level;
  U32 sub_level_index;
  U32 level_offset[24];
  U32 coarser_indices[4];
  U32 adaptive_alloc;
  U32* adaptive;

  void intersect_rectangle_with_cells(const F64 r_min_x, const F64 r_min_y, const F64 r_max_x, const F64 r_max_y, const F32 cell_min_x, const F32 cell_max_x, const F32 cell_min_y, const F32 cell_max_y, U32 level, U32 level_index);
  void intersect_rectangle_with_cells_adaptive(const F64 r_min_x, const F64 r_min_y, const F64 r_max_x, const F64 r_max_y, const F32 cell_min_x, const F32 cell_max_x, const F32 cell_min_y, const F32 cell_max_y, U32 level, U32 level_index);
  void intersect_tile_with_cells(const F32 ll_x, const F32 ll_y, const F32 ur_x, const F32 ur_y, const F32 cell_min_x, const F32 cell_max_x, const F32 cell_min_y, const F32 cell_max_y, U32 level, U32 level_index);
  void intersect_tile_with_cells_adaptive(const F32 ll_x, const F32 ll_y, const F32 ur_x, const F32 ur_y, const F32 cell_min_x, const F32 cell_max_x, const F32 cell_min_y, const F32 cell_max_y, U32 level, U32 level_index);
  void intersect_circle_with_cells(const F64 center_x, const F64 center_y, const F64 radius, const F64 r_min_x, const F64 r_min_y, const F64 r_max_x, const F64 r_max_y, const F32 cell_min_x, const F32 cell_max_x, const F32 cell_min_y, const F32 cell_max_y, U32 level, U32 level_index);
  void intersect_circle_with_cells_adaptive(const F64 center_x, const F64 center_y, const F64 radius, const F64 r_min_x, const F64 r_min_y, const F64 r_max_x, const F64 r_max_y, const F32 cell_min_x, const F32 cell_max_x, const F32 cell_min_y, const F32 cell_max_y, U32 level, U32 level_index);
  BOOL intersect_circle_with_rectangle(const F64 center_x, const F64 center_y, const F64 radius, const F32 r_min_x, const F32 r_max_x, const F32 r_min_y, const F32 r_max_y);
  void raster_occupancy(BOOL(*does_cell_exist)(I32), U32* data, U32 min_x, U32 min_y, U32 level_index, U32 level, U32 stop_level) const;
  void* current_cells;
  U32 next_cell_index;
};

#endif
