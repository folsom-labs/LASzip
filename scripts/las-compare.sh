#!/bin/bash

set -o nounset
set -o errexit
set -o pipefail

go run binaryreader.go lasreader.go testgo.go geotiff.go -compare-with-las2txt $1
# go run testgo.go $@
