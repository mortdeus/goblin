package main

import (
	"flag"
	"fmt"
	"os"
)

var (
	cmd = struct{ name, flags string }{
		"echo",
		"[ -n ] [ arg ...]",
	}
	nflag = flag.Bool("n", false, "suppress newline")
)

func usage() {
	fmt.Fprintln(os.Stderr, "Usage:", cmd.name, cmd.flags)
	flag.PrintDefaults()
	os.Exit(2)
}

func main() {
	flag.Usage = usage
	flag.Parse()
	args := flag.Args()
	for _, arg := range args {
		fmt.Printf("%v", arg)
	}
	if !*nflag {
		fmt.Print("\n")
	}
}
