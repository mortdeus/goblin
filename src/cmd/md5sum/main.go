package main

import (
	"crypto/md5"
	"flag"
	"fmt"
	"io"
	"os"
)

var cmd = struct{ name, flags string }{
	"md5sum",
	"[file ...]",
}

func usage() {
	fmt.Fprintln(os.Stderr, "Usage:", cmd.name, cmd.flags)
	os.Exit(2)
}

func fatal(err error) {
	fmt.Fprintf(os.Stderr, "%s: %s\n", cmd.name, err)
	os.Exit(1)
}

func sum(f *os.File, path string) {
	h := md5.New()
	if _, err := io.Copy(h, f); err != nil {
		fatal(err)
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
				fmt.Fprintf(os.Stderr, "%s: %s\n", cmd.name, err)
				continue
			}
			sum(f, path)
			f.Close()
		}
	}
}
