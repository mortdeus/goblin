package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"os"
)

var cmd = struct{ name, flags string }{
	"cat",
	"[ file ... ]",
}

func usage() {
	fmt.Fprintf(os.Stderr, "Usage:"+cmd.name+"\t"+cmd.flags)
	flag.PrintDefaults()
	os.Exit(2)
}
func fatal(err error) {
	log.Fatalln(cmd.name + ":\t" + err.Error())
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
				log.Fatal(err)
			}
			cat(f, s)
		}
	}
}
func cat(f io.Reader, s string) {
	b := make([]byte, 8*1024)
	for {
		n, err := f.Read(b)
		if n == 0 {
			return
		}
		if err != nil {
			fatal(err)
		}
		if n, err = os.Stdout.Write(b[:n]); err != nil {
			fatal(err)
		}
	}
}
