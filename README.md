[![Go Report Card](https://goreportcard.com/badge/github.com/lorenzodisidoro/avocado-cli)](https://goreportcard.com/report/github.com/lorenzodisidoro/avocado-cli)
[![Build Status](https://travis-ci.com/lorenzodisidoro/avocado-cli.svg?branch=master)](https://travis-ci.com/lorenzodisidoro/avocado-cli)

# Avocado CLI
Avocado is a small surface command line interface to use Avocado SDK.

## Supported OS
- OSX
- Linux

## Installation
Generate your RSA key pair and run the following command
```sh
curl https://raw.githubusercontent.com/lorenzodisidoro/avocado-cli/master/scripts/install_cli.sh | bash
```

or specify the version [tag](https://github.com/lorenzodisidoro/avocado-cli/tags) of the CLI
```sh
curl https://raw.githubusercontent.com/lorenzodisidoro/avocado-cli/master/scripts/install_cli.sh | bash -s VERSION_TAG
```

You can do something like this

[![asciicast](https://asciinema.org/a/RiSS8iQtoFAa76wtqWNFgli80.svg)](https://asciinema.org/a/RiSS8iQtoFAa76wtqWNFgli80)

## Getting Started

### Configuration
Before using Avocado CLI, you need to initialize the `config.json` file, you cna do this running the following command
```sh
$ avocado init [/PATH/TO/PRIVATE_KEY.pem]
```

The CLI creates and reads this file in `$HOME/.avocado/` directory
```json
{
   "storage":"",
   "bolt":{
      "path":"$HOME/.avocado/avocado.db",
      "bucket":"avocado"
   },
   "redis":{
      "address":"localhost:6379",
      "password":"",
      "db":0
   },
   "publicPath":"$HOME/.avocado/public_key.pem"
}
```

### Encrypt
A new value can be encrypted and stored using the following command
```sh
$ avocado encrypt [STORAGE_KEY]
```

### Encrypt
Stored values can be decrypted using the following command
```sh
$ avocado decrypt [STORAGE_KEY] [/PATH/TO/PRIVATE_KEY.pem]
```