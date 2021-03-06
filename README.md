## gocheck

[![Build Status](https://img.shields.io/travis/frankbraun/gocheck.svg?style=flat-square)](https://travis-ci.org/frankbraun/gocheck) [![Go Report Card](https://goreportcard.com/badge/github.com/frankbraun/gocheck?style=flat-square)](https://goreportcard.com/report/github.com/frankbraun/gocheck)

gocheck checks Go source code by running common source code checkers and
executing unit tests. It executes the following checkers:

```
goimports -l -w
gofmt -l -w -s
golint
go vet
```

gocheck also executes `go test` in all subdirectories which contain test files.


### Installation

```
go get github.com/frankbraun/gocheck
```


### Usage

```
usage: gocheck [flags] [path ...]
  -c	enable coverage analysis
  -e value
    	exclude subdirectory (can be specified repeatedly) (default [vendor])
  -g	install necessary tools with go get
  -novet
    	disable go vet
  -timeout duration
    	timeout for go test (default 10m0s)
  -v	be verbose
```


### Integration into Travis

A typical `.travis.yml` file for gocheck integration into Travis looks like this:

```
language: go
go: 1.13
before_install:
  - go get github.com/frankbraun/gocheck
script:
  - gocheck -g -c
```
