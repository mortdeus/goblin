package main

import (
	"flag"
	"fmt"
	"os"
)

var mkParents = flag.Bool("p", false, "Create any necessary parent directories.")
var mode = flag.Int("m", 0777, "Permissions to use when creating the directory.")

func main() {
	flag.Usage = usage
	flag.Parse()

	m := os.FileMode(*mode)

	for _, a := range flag.Args() {
		var err error
		if *mkParents {
			err = os.MkdirAll(a, m)
		} else {
			err = os.Mkdir(a, m)
		}
		if err != nil {
			fmt.Fprintln(os.Stderr, "mkdir:", err)
		}
	}
}

func usage() {
	fmt.Print("usage: mkdir [-p] [-m mode] dir...")
	flag.PrintDefaults()
	os.Exit(2)
}
