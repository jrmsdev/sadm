#!/bin/sh
set -eu
ARGS=${@:-'./...'}
rm -vrf ./_test
go clean -i -testcache ${ARGS}
exit 0
