#!/bin/bash

set -o nounset
set -o errexit
set -o pipefail

clang -g -Wall -DUNORDERED=1 -DLASZIPDLL_EXPORTS=1 -Isrc -lstdc++ src/*.cpp example/laszipdllexample.cpp -o lasz_tool

# if gcc 5 is installed with: brew install gcc
#/usr/local/bin/gcc-5 -g -Wall -DUNORDERED=1 -DLASZIPDLL_EXPORTS=1 -Isrc -lstdc++ src/*.cpp example/laszipdllexample.cpp -o lasz_tool_gcc


# if latest clang is installed with: brew install llvm --HEAD
# xcrun is needed to point to system c++ headers
#xcrun /usr/local/opt/llvm/bin/clang -g -Wall -DUNORDERED=1 -DLASZIPDLL_EXPORTS=1 -Isrc -lstdc++ src/*.cpp example/laszipdllexample.cpp -o lasz_tool_clang
