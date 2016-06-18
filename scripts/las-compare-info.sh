#!/bin/bash

set -o nounset
set -o errexit
set -o pipefail

go run testgo.go *reader.go geotiff*.go -compare-with-lasinfo $1