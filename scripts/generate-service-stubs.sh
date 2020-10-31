#!/usr/bin/env bash

rm -rf rpc/vivsemoapi/*.go
# Get the parent directory of where this script is.
SOURCE="${BASH_SOURCE[0]}"
while [ -h "$SOURCE" ] ; do SOURCE="$(readlink "$SOURCE")"; done
DIR="$( cd -P "$( dirname "$SOURCE" )/.." && pwd )"

# Change into that directory
cd "$DIR" | exit

protoc --proto_path=$GOPATH/src:. --twirp_out=. --go_out=. rpc/vivsemoapi/service.proto
