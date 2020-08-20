package main

import (
	"flag"
	"fmt"
	"os"
	"time"
)

var (
	coverage bool
	get      bool
	verbose  bool
	exclude  excludePaths
	timeout  time.Duration
	novet    bool
)

const defaultTimeout = 10 * time.Minute

type excludePaths []string

func (e *excludePaths) String() string     { return fmt.Sprint(*e) }
func (e *excludePaths) Set(v string) error { *e = append(*e, v); return nil }

func init() {
	exclude = []string{"vendor"}
	flag.BoolVar(&coverage, "c", false, "enable coverage analysis")
	flag.BoolVar(&get, "g", false, "install necessary tools with go get")
	flag.BoolVar(&verbose, "v", false, "be verbose")
	flag.DurationVar(&timeout, "timeout", defaultTimeout, "timeout for go test")
	flag.Var(&exclude, "e", "exclude subdirectory (can be specified repeatedly)")
	flag.BoolVar(&novet, "novet", false, "disable go vet")
}

func fatal(err error) {
	fmt.Fprintf(os.Stderr, "%s: error: %s\n", os.Args[0], err)
	os.Exit(1)
}

func usage() {
	fmt.Fprintf(os.Stderr, "usage: %s [flags] [path ...]\n", os.Args[0])
	flag.PrintDefaults()
	os.Exit(2)
}

func main() {
	flag.Usage = usage
	flag.Parse()
	if get {
		if err := getTools(verbose); err != nil {
			fatal(err)
		}
	}
	var paths []string
	if flag.NArg() == 0 {
		paths = []string{"."}
	} else {
		paths = flag.Args()
	}
	if err := check(paths, exclude, coverage, verbose); err != nil {
		fatal(err)
	}
}
