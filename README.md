## gocheck [![Build Status](http://img.shields.io/travis/frankbraun/gocheck.svg?style=flat-square)](https://travis-ci.org/frankbraun/gocheck)

gocheck checks Go source code by running common source code checker and
executing unit tests. It executes the following checkers:

```
goimports -l -w
gofmt -l -w -s
golint
go tool vet
```

gocheck also executes `go test` in all subdirectories which contain test files.


### Installation

```
go get github.com/frankbraun/gocheck
```


### Usage

```
usage: gocheck [flags] [path ...]
  -c    enable coverage analysis
  -e value
        exclude subdirectory (can be specified repeatedly) (default [vendor])
  -g    install necessary tools with go get
  -v    be verbose
```


### Integration into Travis

A typical `.travis.yml` file for gocheck looks like this:

```
language: go
go: 1.5
env: GO15VENDOREXPERIMENT=1
before_install:
  - go get github.com/frankbraun/gocheck
script:
  - gocheck -g -c -e vendor
```
