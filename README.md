Serveo
=======================================================================

[![GoDoc](https://godoc.org/github.com/Axelen123/serveo?status.svg)](https://godoc.org/github.com/Axelen123/serveo)

- [What is Serveo?](#what-is-serveo)
- [Installation](#installation)
- [Usage](#usage)
- [Docs](#docs)
- [Contributing](#contributing)
- [License](#license)

## What is Serveo?
This is an open source (unofficial) CLI for serveo.net,
a service that allows you to expose local servers without port forwarding.

## Installation

Right now the only way to install serveo is from source.

You need to install [go](https://golang.org) first, I recommend version 1.12 or higher.
You also need openssh installed:
### Debian/Ubuntu Linux

Open a terminal and run:
```sh
sudo apt-get install openssh-client
go get -v -u github.com/Axelen123/serveo/cmd/serveo
```

### MacOS

Open a terminal and run:
```sh
go get -v -u github.com/Axelen123/serveo/cmd/serveo
```

### Windows 10

Open the start menu and search for "*features*".
Click on "*Apps and features*" -> "*Manage optional features*". Scroll down until you find "*OpenSSH Client*" (Not the server!) and click install.

Open a command prompt and run:
```sh
go get -v -u github.com/Axelen123/serveo/cmd/serveo
```

Now the binary should be in your GOPATH/bin folder.
Make sure to add GOPATH/bin to your PATH.


## Usage

### HTTP
Exposing an http server is easy:
```sh
serveo http

#On another port:

# Replace 8080 with the port your server uses
serveo http -p 8080 

# You can also specify a custom [sub]domain:
# (if you don't specify a domain, your username will be used)

# These do the same thing
serveo http -d hello
serveo http -d hello.serveo.net

# For info on using a custom domain, see the wiki.
serveo http -d mydomain.com
```

### SSH
You can expose SSH as well!
```sh
serveo ssh

# ssh also works with domains

serveo ssh -d hello
```

### Config

The cli also supports config files, to initialize one you need to run:
```sh
serveo init
```

It should give you a file that looks like this:
```json
{
  "http": 80,
  "ssh": false,
  "domain": "",
  "tcp": []
}
```

The configuration options are similar to the CLI options:

- http: contains the port that an http server uses. You can disable it by setting to -1.
- ssh: expose ssh server or not
- domain: domain/alias for http and ssh
- tcp: contains paths to expose, example (remove comments if you want to use these):
```js
{
    "http": -1,
    "ssh": false,
    "domain": "",
    "tcp": [
        {
            "local": {
                // Expose http server on
                // raspberry pi
                "host": "raspberrypi",
                "port": 80
            },
            "remote": {
                "host": "pisrv",
                "port": 80
            }
        },
        {
            "local": {
                // expose local redis
                "host": "",
                "port": 6379
            },
            "remote": {
                "host": "redis",
                "port": 6379
            }
        }
    ]
}
```

#### Using the config

You can use the config like this:
```sh
# If you run without arguments it will
# use "serveo.config.json" as config file
serveo

# You can also specify config file:
serveo -c myconf.json
```

## Docs

You can find reference documentation at [godoc](https://godoc.org/github.com/Axelen123).

## Contributing

Note: by contributing code to the Redis project in any form, including sending a pull request via Github, a code fragment or patch via private email or public discussion groups, you agree to release your code under the terms of the MIT license that you can find in the [LICENSE](LICENSE) file included in the source distribution.

See CONTRIBUTING.md for more info.

## License

[MIT](LICENSE)

This project is not affiliated with serveo.net or its creator(s)

