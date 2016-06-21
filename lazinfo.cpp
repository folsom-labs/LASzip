#include <time.h>
#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <inttypes.h>

#include "laszip_api.h"

void usage()
{
  fprintf(stderr,"usage:\n");
  fprintf(stderr,"unlaz in.laz\n");
  exit(1);
}

static void dll_error(void * laszip)
{
  if (laszip) {
    return;
  }
  char* error;
  if (laszip_get_error(laszip, &error)) {
      fprintf(stderr,"DLL ERROR: getting error messages\n");
  }
  fprintf(stderr,"DLL ERROR MESSAGE: %s\n", error);
}

static void fatal(void * laszip=nullptr)
{
  dll_error(laszip);
  exit(1);
}

static void unlaz(char *file_name_in) {
  printf("file: %s\n", file_name_in);

  void * laszip_reader = nullptr;
  if (laszip_create(&laszip_reader))
  {
    fprintf(stderr,"DLL ERROR: creating laszip reader\n");
    fatal();
  }

  laszip_BOOL is_compressed = 0;
  if (laszip_open_reader(laszip_reader, file_name_in, &is_compressed))
  {
    fprintf(stderr,"DLL ERROR: opening laszip reader for '%s'\n", file_name_in);
    fatal(laszip_reader);
  }

  fprintf(stderr,"file '%s' is %scompressed\n", file_name_in, (is_compressed ? "" : "un"));

  laszip_header* header;
  if (laszip_get_header_pointer(laszip_reader, &header))
  {
    fprintf(stderr,"DLL ERROR: getting header pointer from laszip reader\n");
    fatal(laszip_reader);
  }

  uint64_t npoints = (header->number_of_point_records ? header->number_of_point_records : header->extended_number_of_point_records);

  fprintf(stderr,"file '%s' contains %lld points\n", file_name_in, npoints);

  // this is where point data is after reading
  laszip_point* point;

  int32_t maxX = 0, maxY = 0, maxZ = 0;
  int32_t minX = 0, minY = 0, minZ = 0;

  if (laszip_get_point_pointer(laszip_reader, &point))
  {
    fprintf(stderr,"DLL ERROR: getting point pointer from laszip reader\n");
    fatal(laszip_reader);
  }

  for (uint64_t n = 0; n < npoints; n++) {
    if (laszip_read_point(laszip_reader))
    {
      fprintf(stderr,"DLL ERROR: reading point %lld\n", n);
      fatal(laszip_reader);
    }
    if (n == 0) {
      minX = maxX = point->X;
      minY = maxY = point->Y;
      minZ = maxZ = point->Z;
    } else {
      if (point->X > maxX) {
        maxX = point->X;
      }
      if (point->Y > maxY) {
        maxY = point->Y;
      }
      if (point->Z > maxZ) {
        maxZ = point->Z;
      }

      if (point->X < minX) {
        minX = point->X;
      }
      if (point->Y < minY) {
        minY = point->Y;
      }
      if (point->Z > minZ) {
        minZ = point->Z;
      }
    }
  }

  if (laszip_close_reader(laszip_reader))
  {
    fprintf(stderr,"DLL ERROR: closing laszip reader\n");
    fatal(laszip_reader);
  }

  if (laszip_destroy(laszip_reader))
  {
    fprintf(stderr,"DLL ERROR: destroying laszip reader\n");
    fatal();
  }

  fprintf(stdout, "minX: %d maxX: %d\n", (int)minX, (int)maxX);
  fprintf(stdout, "minY: %d maxY: %d\n", (int)minY, (int)maxY);
  fprintf(stdout, "minZ: %d maxZ: %d\n", (int)minZ, (int)maxZ);
}

int main(int argc, char *argv[])
{
  char* file_name = nullptr;

  if (argc !=2) {
    usage();
  }

  file_name = argv[1];
  unlaz(file_name);
  return 0;
}
