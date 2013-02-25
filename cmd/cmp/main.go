package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"strconv"
)

var lflag = flag.Bool("l", false, "do not print and character")
var sflag = flag.Bool("s", false, "silent, do not print errors")
var Lflag = flag.Bool("L", false, "print line number of difference")

func usage() {
	fmt.Fprintf(os.Stderr, "usage: cmp [-lsL] file1 file2 [offset1 [offset2]]\n")
	os.Exit(2)
}

func error(s string) {
	if !*sflag {
		fmt.Fprint(os.Stderr, s, "\n")
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
		error(err.Error())
	}
	f2, err := os.Open(args[1])
	if err != nil {
		error(err.Error())
	}

	xseek := func(f *os.File, fn string, os string) {
		o, err := strconv.ParseInt(os, 0, 64)
		if err != nil {
			error(fmt.Sprintf("bad offset %s: %s", os, err))
		}
		_, err = f.Seek(o, 0)
		if err != nil {
			error(err.Error())
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
			error(err1.Error())
		}
		if err2 != nil {
			error(err2.Error())
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
	os.Exit(0)
}
