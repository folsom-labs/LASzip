#!/bin/bash

# To install latest llvm on mac:
# brew install llvm --with-clang --with-clang-extra-tool --with-compiler-rt --with-libcxx --with-lld --with-python --with-utils --HEAD
# to install gcc 6 on mac:
# brew install gcc

set -o nounset
set -o errexit
set -o pipefail

rm -rf lazinfo
rm -rf lazinfo.dSYM

clang -std=c++11 -g -Wall -Isrc -lstdc++ src/*.cpp lazinfo.cpp -o lazinfo

#/usr/local/bin/gcc-6 -g -Wall -Isrc -lstdc++ src/*.cpp lazinfo.cpp -o lazinfo_gcc

# xcrun is needed to point to system c++ headers
#xcrun /usr/local/opt/llvm/bin/clang -std=c++11 -g -Wall -Isrc -lstdc++ src/*.cpp lazinfo.cpp -o lazinfo_clang
