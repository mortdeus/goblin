package main

import (
	"compress/gzip"
	"fmt"
	"github.com/guelfey/flag9"
	"io"
	"log"
	"os"
	"strings"
)

var (
	cflag bool
	level int = -1
	verbose bool
)

func usage() {
	fmt.Fprintf(os.Stderr, "Usage: gzip [-vcD] [-1-9] [file ...]\n")
	os.Exit(2)
}

func doGzip(in io.Reader, out io.Writer) error {
	// The error is only non-nil if the compression level is wrong, which
	// cannot happen.
	w, _ := gzip.NewWriterLevel(out, level)
	defer w.Close()
	_, err := io.Copy(w, in)
	return err
}

func gzipFile(name string) {
	var newname string

	in, err := os.Open(name)
	if err != nil {
		log.Fatal(err)
	}
	defer in.Close()
	if !cflag {
		if strings.HasSuffix(name, ".tar") {
			newname = name[:len(name)-2] + "gz"
		} else {
			newname = name + ".gz"
		}
		if verbose {
			fmt.Fprintln(os.Stderr, "compressing", name, "to", newname)
		}
		out, err := os.Create(newname)
		if err != nil {
			log.Fatal(err)
		}
		defer out.Close()
		err = doGzip(in, out)
		if err != nil {
			// Don't leave a possibly corrupted file, but try and remove it.
			os.Remove(newname)
			log.Fatal(err)
		}
	} else {
		err = doGzip(in, os.Stdout)
		if err != nil {
			log.Fatal(err)
		}
	}
}

func main() {
	for flag9.Next() {
		c := flag9.Argc()
		switch {
		case c >= '1' && c <= '9':
			level = int(c) - '0'
		case c == 'c':
			cflag = true
		case c == 'v':
			verbose = true
		case c == 'D':
			// debugging information - stub for now
		default:
			usage()
		}
	}
	args := flag9.Argv()
	if len(args) == 0 {
		err := doGzip(os.Stdin, os.Stdout)
		if err != nil {
			log.Fatal(err)
		}
	} else {
		for _, v := range args {
			gzipFile(v)
		}
	}
}
