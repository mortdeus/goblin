package main

import (
	"flag"
	"fmt"
	"os"
)

var (
	cmd = struct{ name, flags string }{
		"mkdir",
		"[-p] [-m mode] dir...",
	}
	mkParents = flag.Bool("p", false, "Create any necessary parent directories.")
	mode      = flag.Int("m", 0777, "Permissions to use when creating the directory.")
)

func usage() {
	fmt.Fprintln(os.Stderr, "Usage:", cmd.name, cmd.flags)
	flag.PrintDefaults()
	os.Exit(2)
}

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
			fmt.Fprintf(os.Stderr, "%s: %s\n", cmd.name, err)
		}
	}
}
