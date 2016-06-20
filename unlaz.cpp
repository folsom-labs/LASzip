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

/*
static void dll_error(void * laszip)
{
  if (laszip) {
    return;
  }
  char* error;
  if (laszip_get_error(laszip, &error) {
      fprintf(stderr,"DLL ERROR: getting error messages\n");
  }
  fprintf(stderr,"DLL ERROR MESSAGE: %s\n", error);
}

static void fatal(void * laszip=nullptr)
{
  dll_error(laszip);
  exit(1);
}
*/

static void unlaz(char *file_name) {
  printf("file: %s\n", file_name);
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
