package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
	"time"
)

var cmd = struct{ name, flags string }{
	"sleep",
	"seconds",
}

func usage() {
	fmt.Fprintln(os.Stderr, "Usage:", cmd.name, cmd.flags)
	os.Exit(2)
}

func main() {
	flag.Usage = usage
	flag.Parse()
	args := flag.Args()
	if len(args) != 1 {
		usage()
	}
	n, err := strconv.Atoi(args[0])
	if err != nil {
		usage()
	}
	for ; n > 0; n-- {
		time.Sleep(1 * time.Second)
	}
}
