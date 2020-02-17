#!/bin/sh
set -eu
go build -mod=vendor -o ./_build/sadm -i ./cmd/sadm
exec ./_build/sadm $@
