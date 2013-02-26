package main

import (
	"crypto/md5"
	"flag"
	"fmt"
	"io"
	"os"
)

func usage() {
	fmt.Fprintf(os.Stderr, "usage: md5sum [file ...]\n")
	os.Exit(2)
}

func errExit(err error) {
	fmt.Fprintln(os.Stderr, "ms5sum:", err)
	os.Exit(1)
}

func sum(f *os.File, path string) {
	h := md5.New()
	if _, err := io.Copy(h, f); err != nil {
		errExit(err)
	}
	if path == "" {
		fmt.Printf("%x\n", h.Sum(nil))
	} else {
		fmt.Printf("%x\t%s\n", h.Sum(nil), path)
	}
}

func main() {
	flag.Usage = usage
	flag.Parse()
	args := flag.Args()
	if len(args) == 0 {
		sum(os.Stdin, "")
	} else {
		for _, path := range args {
			f, err := os.Open(path)
			if err != nil {
				fmt.Fprintf(os.Stderr, err.Error())
				continue
			}
			sum(f, path)
			f.Close()
		}
	}
	os.Exit(0)
}
