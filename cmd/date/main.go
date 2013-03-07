package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
	"time"
)

var (
	cmd = struct{ name, flags string }{
		"date",
		"[ -u ] [ -n ] [ seconds ]",
	}
	nflag = flag.Bool("n", false, "print as number of seconds since epoch")
	uflag = flag.Bool("u", false, "print utc")
)

func usage() {
	fmt.Fprintln(os.Stderr, "Usage", cmd.name, cmd.flags)
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
	var now time.Time
	if len(args) == 0 {
		now = time.Now()
	} else if len(args) != 1 {
		usage()
	} else {
		s, err := strconv.ParseInt(args[0], 0, 64)
		if err != nil {
			fatal(fmt.Errorf("bad number: %s", err))
		}
		now = time.Unix(s, 0)
	}
	if *nflag {
		fmt.Printf("%d\n", now.Unix())
	} else if *uflag {
		fmt.Printf("%s\n", now.UTC().Format("Mon Jan 2 15:04:05 MST 2006"))
	} else {
		fmt.Printf("%s\n", now.Format("Mon Jan 2 15:04:05 MST 2006"))
	}
}
