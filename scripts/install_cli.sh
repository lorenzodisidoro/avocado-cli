#!/usr/bin/env bash

set -e
set -o pipefail

RELEASE_TAG=$1

if [ -z "$RELEASE_TAG" ]
then
  echo -e "Missing release tag argument, using default Avocado version 0.0.2"
  RELEASE_TAG=0.0.2
fi

case $OSTYPE in
  darwin*) OS="darwin";;
  linux-gnu*) OS="linux";;
  *) echo "OS $OSTYPE not supported"; exit 1;;
esac

case $(uname -m) in
  armv7l) ARCH="arm";;
  amd64) ARCH="amd64";;
  x86_64) ARCH="amd64";;
  *) echo "OS type $ARCH not supported"; exit 1;;
esac

echo "Install or update Avocado CLI"

BIN="$HOME/bin"
mkdir -p $BIN

RELEASE_URL="https://github.com/lorenzodisidoro/avocado-cli/releases/download/$RELEASE_TAG/avocato-$OS-$ARCH"

echo "Downloading $RELEASE_URL ..."

curl -sL $RELEASE_URL > avocado

case $SHELL in
  *zsh) PROFILE="$HOME/.zshrc";;
  *ksh) PROFILE="$HOME/.kshrc";;
  *bash)
    if [ -f "$HOME/.bash_profile" ]; then
      PROFILE="$HOME/.bash_profile"
    elif [ -f "$HOME/.bash_login" ]; then
      PROFILE="$HOME/.bash_login"
    elif [ -f "$HOME/.profile" ]; then
      PROFILE="$HOME/.profile"
    fi
    ;;
  *csh)
    if [ -f "$HOME/.tcshrc" ]; then
      PROFILE="$HOME/.tcshrc"
    elif [ -f "$HOME/.cshrc" ]; then
      PROFILE="$HOME/.cshrc"
    fi
    ;;
esac

mv avocado $BIN
chmod +x $BIN/avocado

if [ "$PATH" != *"$BIN"* ]; then
  echo "Exporting path..."
  echo "export PATH=\$PATH:\$HOME/bin" >> $PROFILE
fi

echo "Installation completed successfully!"
echo "Run command 'avocado' and start using it 🥑"