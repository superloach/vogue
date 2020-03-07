package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/superloach/vogue"
)

var debug = flag.Bool("debug", false, "print extra info to stderr")

func init() {
	flag.Parse()
}

func debugln(args ...interface{}) {
	if !*debug {
		return
	}

	fmt.Fprintln(os.Stderr, args...)
}

func main() {
	vg, err := vogue.MkVogue()
	if err != nil {
		panic(err)
	}

	err = vg.Init()
	if err != nil {
		panic(err)
	}

	vg.Fresh()
	vg.Run()
}
