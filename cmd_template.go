//+build ignore
package main

import (
	"flag"
	"fmt"
	"os"
)

var (
	cmd = struct{ name, flags string }{
		"name",
		"[ –f foo] [ –b bar ] [ file ... ]",
	}
	//Flags
	fflag = flag.String(name, value, usage)
	bflag = flag.Bool(name, value, usage)
)

func usage() {
	fmt.Fprintf(os.Stderr, "Usage: %s\t%s\n", cmd.name, cmd.flags)
	flag.PrintDefaults()
	os.Exit(2)
}

func fatal(err error) {
	//silent flag
	//if *sflag {
	//	os.Exit(1)
	//}
	fmt.Fprintf(os.Stderr, "%s:\t%s\n", cmd.name, err.Error())
}

func main() {
	flag.Usage = usage
	flag.Parse()
	args := flag.Args()
}
