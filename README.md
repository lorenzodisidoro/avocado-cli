[![Go Report Card](https://goreportcard.com/badge/github.com/lorenzodisidoro/avocado-cli)](https://goreportcard.com/report/github.com/lorenzodisidoro/avocado-cli)
[![Build Status](https://travis-ci.com/lorenzodisidoro/avocado-cli.svg?branch=master)](https://travis-ci.com/lorenzodisidoro/avocado-cli)

# Avocado CLI
Avocado is a small surface command line interface to use [Avocado SDK](https://github.com/lorenzodisidoro/avocado-sdk), it can be used to manage your passwords with RSA key.

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
Before using Avocado CLI, you need to generate an RSA private key, for example running:
```sh
openssl genrsa -out private_key.pem
```


After that inizialize the CLI configuration file `config.json`, running the following command
```sh
avocado init [/PATH/TO/private_key.pem]
```

The CLI creates and reads this file in `$HOME/.avocado/` directory.

## How to use

### Encrypt
A new value can be encrypted and stored using the following command
```sh
avocado encrypt [STORED_KEY]
```

### Get
Print the stored keys
```sh
avocado get
```

### Decrypt
Stored values can be decrypted using the following command
```sh
avocado decrypt [STORED_KEY] [/PATH/TO/PRIVATE_KEY.pem]
```
decrypted value is written in to the clipboard.

## For Developers
Clone this repository into your `$GOPATH` using
```sh
git clone https://github.com/lorenzodisidoro/avocado-cli
```

moved to the project folder install dependencie with
```sh
go mod tidy
```

The application can be used running 
```sh
go run main.go [command]
```

or create a build and use it
```sh
./scripts/build.sh <GOOS> <GOARCH>
```
You can build this applications for different Operating Systems and architectures (To find this list of possible platforms, run the following: `go tool dist list`)