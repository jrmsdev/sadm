#!/bin/sh
set -eu
docker build --rm --network none -t jrmsdev/sadm \
	--build-arg SADM_UID=$(id -u) \
	--build-arg SADM_GID=$(id -g) \
	--build-arg SADM_UMASK=$(umask) \
	./devel
exit 0
