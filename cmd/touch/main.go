package main

import (
	"flag"
	"fmt"
	"os"
	"time"
)

var (
	cmd = struct{ name, flags string }{
		"touch",
		"[ -c ] [ -t time] files...",
	}
	dontCreate = flag.Bool("c", false, "don't create non-existend file")
	newt       = flag.Int64("t", time.Now().Unix(), "new modification time")
)

func fatal(err error) {
	fmt.Fprintf(os.Stderr, "%s: %s\n", cmd.name, err)
	os.Exit(1)
}

func main() {
	flag.Usage = usage
	flag.Parse()

	var t time.Time
	if newt != nil {
		t = time.Unix(*newt, 0)
	} else {
		t = time.Now()
	}

	for _, a := range flag.Args() {
		err := os.Chtimes(a, t, t) // ??? We set the atime too

		if os.IsNotExist(err) {
			if *dontCreate {
				fatal(err)
			}
			f, err := os.Create(a)
			if err != nil {
				f.Close()
			}

			if err == nil && newt != nil {
				err = os.Chtimes(a, t, t)
			}

			if err != nil {
				fatal(err)
			}
		}
	}
}

func usage() {
	fmt.Fprintln(os.Stderr, "Usage:", cmd.name, cmd.flags)
	flag.PrintDefaults()
	os.Exit(2)
}
