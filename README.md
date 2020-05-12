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

If you run `avocado` you should see something like this
```
$ avocado
Avocado is a small surface command line interface to use Avocado SDK

Usage:
  avocado [command]

Available Commands:
  dencrypt    Encrypt and store value, if the key
  encrypt     Encrypt and store value, if the key
  help        Help about any command

Flags:
      --config string   config file (default is ~/.avocado/config.json)
  -h, --help            help for avocado
  -v, --version         version for avocado
```

## Getting Started
Before using Avocado CLI, you need to create your `config.json` file, by default CLI read it in `$HOME/.avocado/` directory
```json
{
   "storage":"",
   "bolt":{
      "path":"/User/me/.avocado/avocado.db",
      "bucket":"avocado"
   },
   "redis":{
      "address":"localhost:6379",
      "password":"",
      "db":0
   },
   "publicPath":"/User/me/.avocado/public_key.pem"
}
```

### Encrypt
A new value can be encrypted and stored using the following command
```sh
$ avocado encrypt [YOUR_KEY]
```

### Encrypt
Stored values can be decrypted using the following command
```sh
$ avocado decrypt [YOUR_KEY] [/PATH/TO/PRIVATE_KEY.pem]
```