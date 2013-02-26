package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
)

var pad = flag.Bool("w", false, "Equalize the widths of all numbers by padding with leading zeros as necessary.")
var format = flag.String("f", "%g", "Use the print(3)-style format for printing each number.")

func main() {

	flag.Usage = usage
	flag.Parse()
	args := flag.Args()

	firsts := "1"
	incrs := "1"
	var lasts string
	switch len(args) {
	case 1:
		lasts = args[0]
	case 2:
		incrs = args[0]
		lasts = args[1]
	case 3:
		firsts = args[0]
		incrs = args[1]
		lasts = args[2]
	default:
		usage()
	}

	first, err := strconv.ParseFloat(firsts, 64)
	if err != nil {
		errExit(err)
	}
	incr, err := strconv.ParseFloat(incrs, 64)
	if err != nil {
		errExit(err)
	}
	last, err := strconv.ParseFloat(lasts, 64)
	if err != nil {
		errExit(err)
	}

	if incr == 0 {
		fmt.Fprintf(os.Stderr, "seq: a zero increment will take you nowhere!\n")
	}

	f := *format + "\n"

	/* TODO
	if pad {
		maxi := 0
		for c := first; c <= last; c += incr {
			maxl = len(fmt.Sprintf("%g", c))
		}
	} */
	if incr > 0 {
		for c := first; c <= last; c += incr {
			fmt.Printf(f, c)
		}
	} else {
		for c := first; c >= last; c += incr {
			fmt.Printf(f, c)
		}
	}

}

func errExit(err error) {
	fmt.Fprintln(os.Stderr, "seq:", err)
	os.Exit(2)
}

func usage() {
	fmt.Print("usage: seq [-fformat] [-w] [first [incr]] last")
	flag.PrintDefaults()
	os.Exit(2)
}
