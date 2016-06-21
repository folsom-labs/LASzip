/*
===============================================================================

  FILE:  lasindex.cpp

  CONTENTS:

    see corresponding header file

  PROGRAMMERS:

    martin.isenburg@rapidlasso.com  -  http://rapidlasso.com

  COPYRIGHT:

    (c) 2011-2015, martin isenburg, rapidlasso - fast tools to catch reality

    This is free software; you can redistribute and/or modify it under the
    terms of the GNU Lesser General Licence as published by the Free Software
    Foundation. See the LICENSE.txt file for more information.

    This software is distributed WITHOUT ANY WARRANTY and without even the
    implied warranty of MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.

  CHANGE HISTORY:

    see corresponding header file

===============================================================================
*/
#include "lasindex.hpp"

#include <stdio.h>
#include <stdlib.h>
#include <string.h>

#include "lasquadtree.hpp"
#include "lasinterval.hpp"
#include "lasreadpoint.hpp"
#include "bytestreamin_file.hpp"

// Check if on OS X and using cland (unordered map isn't part of tr1 namespace)
#if defined(__APPLE__) && defined(__clang__)
  #include <unordered_map>
  using namespace std;
#else
  #include <tr1/unordered_map>
  using namespace std;
  using namespace tr1;
#endif
typedef unordered_map<I32,U32> my_cell_hash;

LASindex::LASindex()
{
  spatial = 0;
  interval = 0;
  have_interval = FALSE;
  start = 0;
  end = 0;
  full = 0;
  total = 0;
  cells = 0;
}

LASindex::~LASindex()
{
  if (spatial) delete spatial;
  if (interval) delete interval;
}

void LASindex::prepare(LASquadtree* spatial, I32 threshold)
{
  if (this->spatial) delete this->spatial;
  this->spatial = spatial;
  if (this->interval) delete this->interval;
  this->interval = new LASinterval(threshold);
}

BOOL LASindex::add(const F64 x, const F64 y, const U32 p_index)
{
  I32 cell = spatial->get_cell_index(x, y);
  return interval->add(p_index, cell);
}

void LASindex::complete(U32 minimum_points, I32 maximum_intervals, const BOOL verbose)
{
  if (verbose)
  {
    fprintf(stderr,"before complete %d %d\n", minimum_points, maximum_intervals);
    print(FALSE);
  }
  if (minimum_points)
  {
    I32 hash1 = 0;
    my_cell_hash cell_hash[2];
    // insert all cells into hash1
    interval->get_cells();
    while (interval->has_cells())
    {
      cell_hash[hash1][interval->index] = interval->full;
    }
    while (cell_hash[hash1].size())
    {
      I32 hash2 = (hash1+1)%2;
      cell_hash[hash2].clear();
      // coarsen if a coarser cell will still have fewer than minimum_points (and points in all subcells)
      BOOL coarsened = FALSE;
      U32 i, full;
      I32 coarser_index;
      U32 num_indices;
      U32 num_filled;
      I32* indices;
      my_cell_hash::iterator hash_element_inner;
      my_cell_hash::iterator hash_element_outer = cell_hash[hash1].begin();
      while (hash_element_outer != cell_hash[hash1].end())
      {
        if ((*hash_element_outer).second)
        {
          if (spatial->coarsen((*hash_element_outer).first, &coarser_index, &num_indices, &indices))
          {
            full = 0;
            num_filled = 0;
            for (i = 0; i < num_indices; i++)
            {
              if ((*hash_element_outer).first == indices[i])
              {
                hash_element_inner = hash_element_outer;
              }
              else
              {
                hash_element_inner = cell_hash[hash1].find(indices[i]);
              }
              if (hash_element_inner != cell_hash[hash1].end())
              {
                full += (*hash_element_inner).second;
                (*hash_element_inner).second = 0;
                num_filled++;
              }
            }
            if ((full < minimum_points) && (num_filled == num_indices))
            {
              interval->merge_cells(num_indices, indices, coarser_index);
              coarsened = TRUE;
              cell_hash[hash2][coarser_index] = full;
            }
          }
        }
        hash_element_outer++;
      }
      if (!coarsened) break;
      hash1 = (hash1+1)%2;
    }
    // tell spatial about the existing cells
    interval->get_cells();
    while (interval->has_cells())
    {
      spatial->manage_cell(interval->index);
    }
    if (verbose)
    {
      fprintf(stderr,"after minimum_points %d\n", minimum_points);
      print(FALSE);
    }
  }
  if (maximum_intervals < 0)
  {
    maximum_intervals = -maximum_intervals*interval->get_number_cells();
  }
  if (maximum_intervals)
  {
    interval->merge_intervals(maximum_intervals, verbose);
    if (verbose)
    {
      fprintf(stderr,"after maximum_intervals %d\n", maximum_intervals);
      print(FALSE);
    }
  }
}

void LASindex::print(BOOL verbose)
{
  U32 total_cells = 0;
  U32 total_full = 0;
  U32 total_total = 0;
  U32 total_intervals = 0;
  U32 total_check;
  U32 intervals;
  interval->get_cells();
  while (interval->has_cells())
  {
    total_check = 0;
    intervals = 0;
    while (interval->has_intervals())
    {
      total_check += interval->end-interval->start+1;
      intervals++;
    }
    if (total_check != interval->total)
    {
      fprintf(stderr,"ERROR: total_check %d != interval->total %d\n", total_check, interval->total);
    }
    if (verbose) fprintf(stderr,"cell %d intervals %d full %d total %d (%.2f)\n", interval->index, intervals, interval->full, interval->total, 100.0f*interval->full/interval->total);
    total_cells++;
    total_full += interval->full;
    total_total += interval->total;
    total_intervals += intervals;
  }
  if (verbose) fprintf(stderr,"total cells/intervals %d/%d full %d (%.2f)\n", total_cells, total_intervals, total_full, 100.0f*total_full/total_total);
}

LASquadtree* LASindex::get_spatial() const
{
  return spatial;
}

LASinterval* LASindex::get_interval() const
{
  return interval;
}

BOOL LASindex::intersect_rectangle(const F64 r_min_x, const F64 r_min_y, const F64 r_max_x, const F64 r_max_y)
{
  have_interval = FALSE;
  cells = spatial->intersect_rectangle(r_min_x, r_min_y, r_max_x, r_max_y);
//  fprintf(stderr,"%d cells of %g/%g %g/%g intersect rect %g/%g %g/%g\n", num_cells, spatial->get_min_x(), spatial->get_min_y(), spatial->get_max_x(), spatial->get_max_y(), r_min_x, r_min_y, r_max_x, r_max_y);
  if (cells)
    return merge_intervals();
  return FALSE;
}

BOOL LASindex::intersect_tile(const F32 ll_x, const F32 ll_y, const F32 size)
{
  have_interval = FALSE;
  cells = spatial->intersect_tile(ll_x, ll_y, size);
//  fprintf(stderr,"%d cells of %g/%g %g/%g intersect tile %g/%g/%g\n", num_cells, spatial->get_min_x(), spatial->get_min_y(), spatial->get_max_x(), spatial->get_max_y(), ll_x, ll_y, size);
  if (cells)
    return merge_intervals();
  return FALSE;
}

BOOL LASindex::intersect_circle(const F64 center_x, const F64 center_y, const F64 radius)
{
  have_interval = FALSE;
  cells = spatial->intersect_circle(center_x, center_y, radius);
//  fprintf(stderr,"%d cells of %g/%g %g/%g intersect circle %g/%g/%g\n", num_cells, spatial->get_min_x(), spatial->get_min_y(), spatial->get_max_x(), spatial->get_max_y(), center_x, center_y, radius);
  if (cells)
    return merge_intervals();
  return FALSE;
}

BOOL LASindex::get_intervals()
{
  have_interval = FALSE;
  return interval->get_merged_cell();
}

BOOL LASindex::has_intervals()
{
  if (interval->has_intervals())
  {
    start = interval->start;
    end = interval->end;
    full = interval->full;
    have_interval = TRUE;
    return TRUE;
  }
  have_interval = FALSE;
  return FALSE;
}

BOOL LASindex::read(const char* file_name)
{
  if (file_name == 0) return FALSE;
  char* name = strdup(file_name);
  if (strstr(file_name, ".las") || strstr(file_name, ".laz"))
  {
    name[strlen(name)-1] = 'x';
  }
  else if (strstr(file_name, ".LAS") || strstr(file_name, ".LAZ"))
  {
    name[strlen(name)-1] = 'X';
  }
  else
  {
    name[strlen(name)-3] = 'l';
    name[strlen(name)-2] = 'a';
    name[strlen(name)-1] = 'x';
  }
  FILE* file = fopen(name, "rb");
  free(name);
  if (file == 0)
  {
    return FALSE;
  }
  ByteStreamIn* stream;
  stream = new ByteStreamInFileLE(file);
  if (!read(stream))
  {
    fprintf(stderr,"ERROR (LASindex): cannot read '%s'\n", name);
    delete stream;
    fclose(file);
    return FALSE;
  }
  delete stream;
  fclose(file);
  return TRUE;
}

BOOL LASindex::read(ByteStreamIn* stream)
{
  if (spatial)
  {
    delete spatial;
    spatial = 0;
  }
  if (interval)
  {
    delete interval;
    interval = 0;
  }
  char signature[4];
  try { stream->getBytes((U8*)signature, 4); } catch (...)
  {
    fprintf(stderr,"ERROR (LASindex): reading signature\n");
    return FALSE;
  }
  if (strncmp(signature, "LASX", 4) != 0)
  {
    fprintf(stderr,"ERROR (LASindex): wrong signature %4s instead of 'LASX'\n", signature);
    return FALSE;
  }
  U32 version;
  try { stream->get32bitsLE((U8*)&version); } catch (...)
  {
    fprintf(stderr,"ERROR (LASindex): reading version\n");
    return FALSE;
  }
  // read spatial quadtree
  spatial = new LASquadtree();
  if (!spatial->read(stream))
  {
    fprintf(stderr,"ERROR (LASindex): cannot read LASspatial (LASquadtree)\n");
    return FALSE;
  }
  // read interval
  interval = new LASinterval();
  if (!interval->read(stream))
  {
    fprintf(stderr,"ERROR (LASindex): reading LASinterval\n");
    return FALSE;
  }
  // tell spatial about the existing cells
  interval->get_cells();
  while (interval->has_cells())
  {
    spatial->manage_cell(interval->index);
  }
  return TRUE;
}

// seek to next interval point

BOOL LASindex::seek_next(LASreadPoint* reader, I64 &p_count)
{
  if (!have_interval)
  {
    if (!has_intervals()) return FALSE;
    reader->seek((U32)p_count, start);
    p_count = start;
  }
  if (p_count == end)
  {
    have_interval = FALSE;
  }
  return TRUE;
}

// merge the intervals of non-empty cells
BOOL LASindex::merge_intervals()
{
  if (spatial->get_intersected_cells())
  {
    U32 used_cells = 0;
    while (spatial->has_more_cells())
    {
      if (interval->get_cell(spatial->current_cell))
      {
        interval->add_current_cell_to_merge_cell_set();
        used_cells++;
      }
    }
//    fprintf(stderr,"LASindex: used %d cells of total %d\n", used_cells, interval->get_number_cells());
    if (used_cells)
    {
      BOOL r = interval->merge();
      full = interval->full;
      total = interval->total;
      interval->clear_merge_cell_set();
      return r;
    }
  }
  return FALSE;
}
