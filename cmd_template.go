//+build ignore
package main

import (
	"flag"
	"fmt"
	"log"
	"os"
)

var (
	cmd = struct{ name, flags string }{
		"name",
		"[ –f foo] [ –b bar ] [ file ... ]",
	}
	fflag = flag.String(name, value, usage)
	bflag = flag.Bool(name, value, usage)
)

func usage() {
	fmt.Fprintln(os.Stderr, "Usage:", cmd.name, cmd.flags)
	flag.PrintDefaults()
	os.Exit(2)
}

func fatal(err error) {
	fmt.Fprintf(os.Stderr, "%s: %s\n", cmd.name, err)
	os.Exit(1)
}

func main() {
	flag.Usage = usage
	flag.Parse()
	args := flag.Args()
}
