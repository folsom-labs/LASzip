#!/bin/bash

# based on https://gist.github.com/markd2/9ba66145135321fa4830
# and https://alastairs-place.net/blog/2016/05/20/code-coverage-from-the-command-line-with-clang/
# http://llvm.org/docs/CommandGuide/llvm-cov.html#llvm-cov-report

set -o nounset
set -o errexit
set -o pipefail

D=/usr/local/opt/llvm/bin

xcrun $D/clang -std=c++11 -g -Wall -fprofile-instr-generate -fcoverage-mapping -Isrc -lstdc++ src/*.cpp lazinfo.cpp -o lazinfo_cov

rm -rf lazinfo_cov.profdata default.profraw

./lazinfo_cov ~/data/lidar/971.laz

xcrun $D/llvm-profdata merge -o lazinfo_cov.profdata default.profraw

xcrun $D/llvm-cov report -instr-profile=lazinfo_cov.profdata ./lazinfo_cov

xcrun $D/llvm-cov show -instr-profile=lazinfo_cov.profdata ./lazinfo_cov
