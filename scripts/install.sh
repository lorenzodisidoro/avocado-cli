  
#!/bin/bash
set -ex

PARENT_PATH=$(dirname $(cd $(dirname $0); pwd -P))
CLI_PATH=$HOME/.avocado

pushd $PARENT_PATH

go mod tidy

echo "Installing Avocado CLI into $CLI_PATH"

mkdir -p $HOME/.avocado
cp -R resources/* $HOME/.avocado

popd