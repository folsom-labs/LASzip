#!/bin/bash

# based on https://gist.github.com/markd2/9ba66145135321fa4830

set -o nounset
set -o errexit
set -o pipefail

D=/usr/local/opt/llvm/bin

xcrun $D/clang -g -Wall -fprofile-instr-generate -fcoverage-mapping -Isrc -lstdc++ src/*.cpp lazinfo.cpp -o lazinfo_cov
rm -rf la
./lazinfo_cov ~/data/lidar/971.laz
xcrun $D/llvm-profdata merge -o lazinfo_cov.profdata default.profraw
xcrun $D/llvm-cov show ./lazinfo_cov -instr-profile=lazinfo_cov.profdata src/* lazinfo.cpp
