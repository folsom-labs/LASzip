#!/bin/bash

set -o nounset
set -o errexit
set -o pipefail

rm -rf ./lasz_tool

clang -g -Wall -DLASZIPDLL_EXPORTS=1 -Isrc -lstdc++ src/*.cpp example/laszipdllexample.cpp example/examples.cpp -o lasz_tool

echo "runing lasz_tool"
./lasz_tool in.laz out.las
