#!/bin/bash

set -o nounset
set -o errexit
set -o pipefail

clang -std=c++11 -g -Wall -Isrc -lstdc++ src/*.cpp lazinfo.cpp -o lazinfo

# if gcc 5 is installed with: brew install gcc
#/usr/local/bin/gcc-5 -g -Wall -Isrc -lstdc++ src/*.cpp lazinfo.cpp -o lazinfo_gcc

# if latest clang is installed with: brew install llvm --with-clang --with-clang-extra-tool --with-compiler-rt --with-lld --with-python --with-utils --HEAD
# xcrun is needed to point to system c++ headers
#xcrun /usr/local/opt/llvm/bin/clang -g -Wall -Isrc -lstdc++ src/*.cpp lazinfo.cpp -o lazinfo_clang
