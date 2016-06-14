#!/bin/bash

# run unit tests for go code

set -o nounset
set -o errexit
set -o pipefail

go test *.go

