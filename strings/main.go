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
	minSpan = flag.Int("m", 6, "Defines the minimum span size for a series "+
		"of runes to be considered a string.")
	verbose = flag.Bool("v", false, "Print the file offsets for the "+
		"beginning of strings.")
)

func usage() {
	fmt.Fprintf(os.Stderr, "Usage: strings [-v] [-m min] [file ...]\n")
	flag.PrintDefaults()
	os.Exit(2)
}

func error(s string) {
	fmt.Fprint(os.Stderr, s, "\n")
	os.Exit(1)
}

func makeString(f *os.File) {
	buf := make([]rune, BUFSIZE)
	b := bufio.NewReader(f)

	var position, start int64
	
	printString := func(buf []rune, start int64) {
		if *verbose {
			fmt.Printf("%8d: %s\n", start, string(buf))
		} else {
			fmt.Printf("%s\n", string(buf))
	}
	
	for {
		c, size, err := b.ReadRune()
		if err != nil {
			break
		}
		position += int64(size)

		if unicode.IsGraphic(c) {
			if start == 0 {
				start = position
			}
			buf = append(buf, c)
			if len(buf) == BUFSIZE {
				buf = append(buf, []rune("...")...)
				printString(buf, start)
				buf = buf[0:0]
				start = 0
			}
		} else {
			if len(buf) >= *minSpan {
				printString(buf, start)
				buf = buf[0:0]
				start = 0
			}
		}
	}

	if len(buf) >= *minSpan {
		printString(buf, start)
	}
}

func main() {
	flag.Usage = usage
	flag.Parse()
	args := flag.Args()

	if len(args) == 0 {
		makeString(os.Stdin)
	} else {
		for _, path := range args {
			f, err := os.Open(path)
			if err != nil {
				fmt.Fprintf(os.Stderr, "open %s: %s", path, err)
				continue
			}
			makeString(f)
			f.Close()
		}
	}

	os.Exit(0)
}
