#!/bin/bash

set -o nounset
set -o errexit
set -o pipefail

go run *.go lidar.las
# go run testgo.go $@
