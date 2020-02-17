#!/bin/sh
set -eu
go mod edit -fmt
gofmt -w -l -s .
exit 0
