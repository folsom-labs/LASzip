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

  void * laszip_reader;
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
