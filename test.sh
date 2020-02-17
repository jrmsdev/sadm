#!/bin/sh
set -eu
ARGS=${@:-'./...'}
go test -mod=vendor ${ARGS}
exit 0
