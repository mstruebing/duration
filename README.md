# Duration

[![CI](https://github.com/mstruebing/duration/actions/workflows/main.yml/badge.svg)](https://github.com/mstruebing/duration/actions/workflows/main.yml)
[![Go Report Card](https://goreportcard.com/badge/github.com/mstruebing/duration)](https://goreportcard.com/report/github.com/mstruebing/duration)

1. [Why](#why) 
2. [Usage](#usage) 
2.1 [Important](#important) 
3. [Installation](#installation)  
3.1 [From Source](#from-source)  
3.2 [With Go](#with-go)  
3.3 [From Release Page](#from-release-page)  
3.3 [via AUR](#via-aur)  

## Why?

If you have processes which takes a bit longer, especially if you know round about how much time it takes to complete, it is cool to know how long the process is already running.

![example](https://raw.githubusercontent.com/mstruebing/duration/master/example/example.gif "example")

## Usage

```sh 
USAGE: duration [h|help|v|version|<command>]
where
        [v|version] - prints the version
        [h|help] - prints the help
        <command> - a SINGLE command or script to execute

would work:
        duration sleep 5
        duration script.sh
would NOT work:
        duration sleep 5 && sleep 4
```


### Important

It's __important__ to understand that currently only a __single__  command with or without arguments can be executed. If you are chaining commands with `||` or `&&` or for example, you need to put your command into a script and execute that.

## Installation

### From Source

In order to build from source you need `go` and `make` installed.
Run: `make build`, this will place a binary `duration` in a `$PWD/bin`, directory.

```sh 
cd ~/ && \
git clone git@github.com:mstruebing/duration.git && \
cd duration && \
make build && \
duration
```

Then put the binary somewhere in your `$PATH`


### With Go

`go get -u github.com/mstruebing/duration`

### From Release Page

You can also grab a release from the [release page](https://github.com/mstruebing/duration/releases).

### Via AUR

If you use Arch-Linux you can simply install the `duration-git`-package.
[AUR-link](https://aur.archlinux.org/packages/duration-git/)

i.e:

```sh 
trizen -S duration-git 
```

### License

Duration is open source software licensed under the MIT License.

See the [LICENSE](./LICENSE) file for more.
