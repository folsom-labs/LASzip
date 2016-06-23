#!/bin/bash

set -o nounset
set -o errexit
set -o pipefail

rm -rf lazinfo
rm -rf lazinfo.dSYM

clang -std=c++11 -g -Wall -Isrc -lstdc++ src/*.cpp lazinfo.cpp -o lazinfo
./lazinfo ~/data/lidar/971.laz
