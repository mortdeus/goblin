package main

import (
	"os"
	"flag"
	"fmt"
	"time"
)

var dontCreate = flag.Bool("c", false, "If the file doesn't exist, don't create it.")
var newt = flag.Int64("t", time.Now().Unix(), "Set the modified time.")

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
				fmt.Fprintf(os.Stderr, "touch: %s\n", err)
				os.Exit(2)
			}
			f, err := os.Create(a)
			if err != nil {
				f.Close()
			}

			if err == nil  && newt != nil {
				err = os.Chtimes(a, t, t)
			}

			if err != nil {
				fmt.Fprintf(os.Stderr, "touch: %s\n", err)
				os.Exit(2)
			}
		}
	}
}

func usage() {
	fmt.Print("usage: touch [-c] [-t time] files...")
	flag.PrintDefaults()
	os.Exit(2)
}
