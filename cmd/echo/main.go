package main

import (
	"flag"
	"fmt"
	"os"
)

var nflag = flag.Bool("n", false, "suppresses newline.")

func usage() {
	fmt.Fprintf(os.Stderr, "Usage: echo [ -n ] [ arg ... ] \n")
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
	if !(*nflag) {
		fmt.Print("\n")
	}
}
