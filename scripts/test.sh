#!/bin/bash

set -o nounset
set -o errexit
set -o pipefail

go run testgo.go lidar.las
# go run testgo.go $@
