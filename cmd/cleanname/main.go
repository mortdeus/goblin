package main

import (
	"flag"
	"fmt"
	"os"
	"path"
	"strings"
)

var dir = flag.String("d", "", "directory to use as working directory")

func usage() {
	fmt.Fprintf(os.Stderr, "usage: cleanname [-d pwd] name ...\n")
	os.Exit(2)
}

func main() {
	flag.Usage = usage
	flag.Parse()
	args := flag.Args()
	if len(args) == 0 {
		usage()
	}

	for _, p := range args {
		if *dir == "" && strings.HasPrefix(p, "/") {
			fmt.Print(path.Clean(p), "\n")
		} else {
			xdir := *dir
			if xdir == "" {
				xdir = "."
			}
			fmt.Print(path.Clean(fmt.Sprint(xdir, "/", p)), "\n")
		}
	}
	os.Exit(0)
}
