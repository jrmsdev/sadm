#!/bin/sh
set -eu
ARGS=${@:-'./...'}
rm -vrf ./_test ./docs/Gemfile.lock ./docs/_site
go clean -mod=vendor -i -testcache ${ARGS}
exit 0
