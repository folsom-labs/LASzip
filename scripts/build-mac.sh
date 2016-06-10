#!/bin/bash

set -o nounset
set -o errexit
set -o pipefail

clang -g -DUNORDERED=1 -DLASZIPDLL_EXPORTS=1 -Isrc -Iinclude -lstdc++ src/*.cpp example/laszipdllexample.cpp -o lasz_tool
