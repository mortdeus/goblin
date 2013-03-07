package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"strconv"
)

var (
	cmd = struct{ name, flags string }{
		"cmp",
		"[ –l ] [ -s ] [ -L ] [ file ... ]",
	}

	lflag = flag.Bool("l", false, "do not print and character")
	sflag = flag.Bool("s", false, "silent, do not print errors")
	Lflag = flag.Bool("L", false, "print line number of difference")
)

func usage() {
	fmt.Fprintln(os.Stderr, "Usage:", cmd.name, cmd.flags)
	flag.PrintDefaults()
	os.Exit(2)
}

func fatal(err error) {
	if !*sflag {
		fmt.Fprintf(os.Stderr, "%s: %s\n", cmd.name, err)
	}
	os.Exit(1)
}

func main() {
	flag.Usage = usage
	flag.Parse()
	args := flag.Args()
	if len(args) < 2 || len(args) > 4 {
		usage()
	}
	f1, err := os.Open(args[0])
	if err != nil {
		fatal(err)
	}
	f2, err := os.Open(args[1])
	if err != nil {
		fatal(err)
	}

	xseek := func(f *os.File, fn string, os string) {
		o, err := strconv.ParseInt(os, 0, 64)
		if err != nil {
			fatal(fmt.Errorf("bad offset %s: %s", os, err))
		}
		_, err = f.Seek(o, 0)
		if err != nil {
			fatal(err)
		}
	}

	if len(args) >= 3 {
		xseek(f1, args[0], args[2])
	}
	if len(args) >= 4 {
		xseek(f2, args[1], args[2])
	}

	r1 := bufio.NewReader(f1)
	r2 := bufio.NewReader(f2)
	nc := 1
	l := 1
	for {
		b1, err1 := r1.ReadByte()
		b2, err2 := r2.ReadByte()
		if err1 == io.EOF && err2 == io.EOF {
			break
		}
		if err1 != nil {
			fatal(err1)
		}
		if err2 != nil {
			fatal(err2)
		}
		if b1 != b2 {
			if *sflag {
				os.Exit(1)
			}
			if !*lflag {
				fmt.Printf("%s %s differ: char %d", args[0], args[1], nc)
				if !*Lflag {
					fmt.Printf(" line %d", l)
				}
				fmt.Print("\n")
				os.Exit(1)
			}
			fmt.Printf("%6d 0x%.2x 0x%.2x\n", nc, b1, b2)
		}
		nc++
		if b1 == '\n' {
			l++
		}
	}
}
