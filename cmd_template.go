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
	fflag = flag.String(name, value, usage)
	bflag = flag.Bool(name, value, usage)
)

func usage() {
<<<<<<< HEAD
	fmt.Fprintf(os.Stderr, "Usage: %s\t%s\n", cmd.name, cmd.flags)
=======
	fmt.Fprintln(os.Stderr, "Usage:", cmd.name, cmd.flags)
>>>>>>> a749fdef772193a37dec178541b533bd2f77c7b5
	flag.PrintDefaults()
	os.Exit(2)
}

func fatal(err error) {
<<<<<<< HEAD
	//silent flag
	//if *sflag {
	//	os.Exit(1)
	//}
	fmt.Fprintf(os.Stderr, "%s:\t%s\n", cmd.name, err.Error())
=======
	fmt.Fprintf(os.Stderr, "%s: %s\n", cmd.name, err)
	os.Exit(1)
>>>>>>> a749fdef772193a37dec178541b533bd2f77c7b5
}

func main() {
	flag.Usage = usage
	flag.Parse()
	args := flag.Args()
}
