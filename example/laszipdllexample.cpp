/*
===============================================================================

  FILE:  laszipdllexample.cpp

  CONTENTS:

    This source code implements several different  easy-to-follow examples on
    how to use the LASzip DLL. The first and the second examples implement a
    small compression and decompression utilitity. The third example shows
    how to use the DLL to export points to a proper geo-referenced LAZ file.

  PROGRAMMERS:

    martin.isenburg@rapidlasso.com  -  http://rapidlasso.com

  COPYRIGHT:

    (c) 2007-2015, martin isenburg, rapidlasso - fast tools to catch reality

    This is free software; you can redistribute and/or modify it under the
    terms of the GNU Lesser General Licence as published by the Free Software
    Foundation. See the LICENSE.txt file for more information.

    This software is distributed WITHOUT ANY WARRANTY and without even the
    implied warranty of MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.

  CHANGE HISTORY:

    23 September 2015 -- 11th example writes without a-priori bounding box or counters
    22 September 2015 -- 10th upconverts to LAS 1.4 with pre-existing "extra bytes"
     5 September 2015 -- eighth and nineth example show pre-existing "extra bytes"
    19 July 2015 -- sixth and seventh example show LAS 1.4 compatibility mode
     2 April 2015 -- fourth and fifth example with integrated spatially indexing
    11 August 2013 -- added third example for exporting geo-referenced points
    29 July 2013 -- created for the LASzip DLL after returning to Sommerhausen

===============================================================================
*/

#include <time.h>
#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <inttypes.h>

#include "laszip_api.h"

void usage(bool wait=false)
{
  fprintf(stderr,"usage:\n");
  fprintf(stderr,"laszipdllexample\n");
  fprintf(stderr,"laszipdllexample in.las out.laz\n");
  fprintf(stderr,"laszipdllexample in.laz out.las\n");
  fprintf(stderr,"laszipdllexample in.las out.las\n");
  fprintf(stderr,"laszipdllexample in.laz out.laz\n");
  fprintf(stderr,"laszipdllexample -h\n");
  if (wait)
  {
    fprintf(stderr,"<press ENTER>\n");
    getc(stdin);
  }
  exit(1);
}

extern void example1(int argc, char* file_name_in, char* file_name_out);
extern void example2(int argc, char* file_name_in, char* file_name_out);
extern void example3(int argc, char* file_name_in, char* file_name_out);
extern void example4(int argc, char* file_name_in, char* file_name_out);
extern void example5(int argc, char* file_name_in, char* file_name_out);
extern void example6(int argc, char* file_name_in, char* file_name_out);
extern void example7(int argc, char* file_name_in, char* file_name_out);
extern void example8(int argc, char* file_name_in, char* file_name_out);
extern void example9(int argc, char* file_name_in, char* file_name_out);
extern void example10(int argc, char* file_name_in, char* file_name_out);
extern void example11(int argc, char* file_name_in, char* file_name_out);

int main(int argc, char *argv[])
{
  char* file_name_in = 0;
  char* file_name_out = 0;

  if (argc == 1)
  {
    char file_name[256];
    fprintf(stderr,"%s is better run in the command line\n", argv[0]);
    // not used for example3
    fprintf(stderr,"enter input file: "); fgets(file_name, 256, stdin);
    file_name[strlen(file_name)-1] = '\0';
    file_name_in = strdup(file_name);
    fprintf(stderr,"enter output file: "); fgets(file_name, 256, stdin);
    file_name[strlen(file_name)-1] = '\0';
    file_name_out = strdup(file_name);
  }
  else if (argc == 3)
  {
    file_name_in = strdup(argv[1]);
    file_name_out = strdup(argv[2]);
  }
  else
  {
    if ((argc != 2) || (strcmp(argv[1], "-h") != 0))
    {
      fprintf(stderr, "ERROR: cannot understand arguments\n");
    }
    usage();
  }

  //example1(argc, file_name_in, file_name_out);
  //example2(argc, file_name_in, file_name_out);
  //example3(argc, file_name_in, file_name_out);
  //example4(argc, file_name_in, file_name_out);
  //example5(argc, file_name_in, file_name_out);
  example6(argc, file_name_in, file_name_out);
  //example7(argc, file_name_in, file_name_out);
  //example8(argc, file_name_in, file_name_out);
  //example9(argc, file_name_in, file_name_out);
  //example10(argc, file_name_in, file_name_out);
  //example11(argc, file_name_in, file_name_out);

  return 0;
}
