#!/bin/bash
set -ex

PARENT_PATH=$(dirname $(cd $(dirname $0); pwd -P))

pushd $PARENT_PATH

mkdir -p build
GOOS=$1 GOARCH=$2 go build -o build/avocato-$1-$2

popd