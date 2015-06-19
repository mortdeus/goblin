package main

import (
	"flag"
	"fmt"
	"io"
	"os"
)

var cmd = struct{ name, flags string }{
	"cat",
	"[ file ... ]",
}

func usage() {
	fmt.Fprintln(os.Stderr, "Usage:", cmd.name, cmd.flags)
	flag.PrintDefaults()
	os.Exit(2)
}

func fatal(err error) {
	fmt.Fprintf(os.Stderr, "%s: %s\n", cmd.name, err)
	os.Exit(1)
}

func main() {
	flag.Usage = usage
	flag.Parse()
	args := flag.Args()
	if len(args) == 0 {
		cat(os.Stdin)
		return
	}
	for _, s := range args {
		f, err := os.Open(s)
		if err != nil {
			fatal(err)
		}
		cat(f)
	}

}

func cat(f io.Reader) {
	if _, err := io.Copy(os.Stdout, f); err != nil {
		fatal(err)
	}
}
