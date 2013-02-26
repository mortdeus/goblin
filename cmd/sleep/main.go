package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
	"time"
)

func usage() {
	fmt.Fprintf(os.Stderr, "usage: sleep seconds\n")
	flag.PrintDefaults()
	os.Exit(2)
}

func main() {
	flag.Usage = usage
	flag.Parse()
	args := flag.Args()
	if len(args) != 1 {
		usage()
	}
	n, err := strconv.Atoi(args[0])
	if err != nil {
		usage()
	}
	for ; n > 0; n-- {
		time.Sleep(1 * time.Second)
	}
	os.Exit(0)
}
