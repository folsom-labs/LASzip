#!/bin/bash

set -o nounset
set -o errexit
set -o pipefail

go run binaryreader.go lasreader.go testgo.go geotiff.go geotiff_constants.go $@
