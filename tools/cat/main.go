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
			error(fmt.Sprintf("%s: read: %s", s, err))
		}
		if n, err = os.Stdout.Write(b[:n]); err != nil {
			error(fmt.Sprintf("write: %s", err))
		}
	}
}

func usage() {
	fmt.Fprintf(os.Stderr, "Usage: cat [file ...]\n")
	flag.PrintDefaults()
	os.Exit(2)
}

func error(s string) {
	fmt.Fprint(os.Stderr, s, "\n")
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
				error(fmt.Sprintf("open %s: %s", s, err))
			}
			cat(f, s)
		}
	}
}
