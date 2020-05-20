[![Go Report Card](https://goreportcard.com/badge/github.com/lorenzodisidoro/avocado-cli)](https://goreportcard.com/report/github.com/lorenzodisidoro/avocado-cli)
[![Build Status](https://travis-ci.com/lorenzodisidoro/avocado-cli.svg?branch=master)](https://travis-ci.com/lorenzodisidoro/avocado-cli)

# Avocado CLI
Avocado is a small surface command line interface to use Avocado SDK, it can be used to encrypt with RSA key the values ​​which will be stored.

## Supported OS
- OSX
- Linux

## Installation
Generate your RSA key pair (e.g. using [ssh-keygen](https://www.ssh.com/ssh/keygen/) tool) and run the following command to install Avocado CLI
```sh
curl https://raw.githubusercontent.com/lorenzodisidoro/avocado-cli/master/scripts/install_cli.sh | bash
```
or specify the version [tag](https://github.com/lorenzodisidoro/avocado-cli/tags) of the CLI
```sh
curl https://raw.githubusercontent.com/lorenzodisidoro/avocado-cli/master/scripts/install_cli.sh | bash -s VERSION_TAG
```
After that try to run the following command
```sh
$ avocado
```

## Getting Started

### Configuration
Before using Avocado CLI, you need to initialize the `config.json` file, you can do this running the following command
```sh
$ avocado init [/PATH/TO/PRIVATE_KEY.pem]
```

The CLI creates and reads this file in `$HOME/.avocado/` directory.

### Encrypt
A new value can be encrypted and s@tored using the following command
```sh
$ avocado encrypt [STORAGE_KEY]
```

### Decrypt
Stored values can be decrypted using the following command
```sh
$ avocado decrypt [STORAGE_KEY] [/PATH/TO/PRIVATE_KEY.pem]
```
decrypted value is written in to the clipboard.