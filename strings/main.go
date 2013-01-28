package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"unicode"
)

const BUFSIZE int = 70

var (
	minspan = flag.Int("m", 6, "Defines the minimum span size for a series "+
		"of runes to be considered a string.")
	offsets = flag.Bool("o", false, "Print the file offsets for the "+
		"beginning of strings.")
)

func usage() {
	fmt.Fprintf(os.Stderr, "Usage: strings [-o] [-m min] [file ...]\n")
	flag.PrintDefaults()
	os.Exit(2)
}

func error(s string) {
	fmt.Fprint(os.Stderr, s, "\n")
	os.Exit(1)
}

func prntstr(s string, start int64) {
	if *offsets {
		fmt.Printf("%8d: %s\n", start, s)
	} else {
		fmt.Printf("%s\n", s)
	}
}

func stringit(f *os.File) {
	buf := make([]rune, BUFSIZE)
	var start, posn int64

	b := bufio.NewReader(f)

	posn = 0
	start = 0
	for {
		c, size, err := b.ReadRune()
		if err != nil {
			break
		}
		posn += int64(size)

		if unicode.IsGraphic(c) {
			if start == 0 {
				start = posn
			}
			buf = append(buf, c)
			if len(buf) == BUFSIZE {
				prntstr(string(buf)+" ...", start)
				buf = buf[0:0]
				start = 0
			}
		} else {
			if len(buf) >= *minspan {
				prntstr(string(buf), start)
				buf = buf[0:0]
				start = 0
			}
		}
	}

	if len(buf) >= *minspan {
		prntstr(string(buf), start)
	}
}

func main() {
	flag.Usage = usage
	flag.Parse()
	args := flag.Args()

	if len(args) == 0 {
		stringit(os.Stdin)
	} else {
		for _, path := range args {
			f, err := os.Open(path)
			if err != nil {
				fmt.Fprintf(os.Stderr, "open %s: %s", path, err)
				continue
			}
			stringit(f)
			f.Close()
		}
	}

	os.Exit(0)
}
