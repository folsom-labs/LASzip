#!/bin/bash

set -o nounset
set -o errexit
set -o pipefail

go run binaryreader.go lasreader.go testgo.go geotiff.go lidar.las
# go run testgo.go $@
