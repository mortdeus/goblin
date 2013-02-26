package main

import (
	"flag"
	"fmt"
	"io"
	"os"
)

func cat(f io.Reader, s string) {
	b := make([]byte, 8*1024)
	for {
		n, err := f.Read(b)
		if n == 0 {
			return
		}
		if err != nil {
			errExit(err)
		}
		if n, err = os.Stdout.Write(b[:n]); err != nil {
			errExit(err)
		}
	}
}

func usage() {
	fmt.Fprintf(os.Stderr, "Usage: cat [file ...]\n")
	flag.PrintDefaults()
	os.Exit(2)
}

func errExit(err error) {
	fmt.Fprintln(os.Stderr, "cat:", err)
	os.Exit(1)
}

func main() {
	flag.Usage = usage
	flag.Parse()
	args := flag.Args()
	if len(args) == 0 {
		cat(os.Stdin, "<stdin>")
	} else {
		for _, s := range args {
			f, err := os.Open(s)
			if err != nil {
				errExit(err)
			}
			cat(f, s)
		}
	}
}
