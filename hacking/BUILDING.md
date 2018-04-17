# Building octl

## Go environment

```bash
$ sudo yum update
$ sudo yum install golang
```

Setting `GOPATH` and `PATH`:
```bash
$ export GOPATH=$HOME/go # or some other path for local Go binaries, packages and sources
$ export PATH=$GOPATH/bin:$PATH
```

Check if all is well:
```bash
$ go version
$ go env
```

## Clone and build

Fetching the sources. You'll get a "no buildable Go source files" error, that's because Octl has its own Makefile instead of using plain `go build`.
```bash
$ go get -d github.com/teo/octl
```

Running make. This will take a while as all dependencies are gathered, built and installed.
```bash
$ cd go/src/github.com/teo/octl
$ make
```

You should find `octld` and `octl-executor` in `bin`.

Continue with [setting up the development cluster](DCOS.md).