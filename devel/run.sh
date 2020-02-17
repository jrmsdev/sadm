#!/bin/sh
set -eu
docker run -it --rm --network none --name jrmsdev-sadm --hostname sadm -u sadm \
	-v ${PWD}:/go/src/sadm jrmsdev/sadm
exit 0
