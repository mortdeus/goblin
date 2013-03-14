package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"unicode"
)

const BUFSIZE = 70

var (
	cmd = struct{ name, flags string }{
		"strings",
		"[-m min] [file ...]",
	}
	minSpan = flag.Int("m", 6, "minimum length of recognized strings")
)

func usage() {
	fmt.Fprintln(os.Stderr, "Usage:", cmd.name, cmd.flags)
	flag.PrintDefaults()
	os.Exit(2)
}

func fatal(err error) {
	fmt.Fprintf(os.Stderr, "%s: %s\n", cmd.name, err)
	os.Exit(1)
}

func makeString(f *os.File) {
	buf := make([]rune, 0)
	b := bufio.NewReader(f)

	var position, start int

	for {
		c, offset, err := b.ReadRune()
		position += offset
		if err != nil {
			if err != io.EOF {
				fatal(err)
			}
			break
		}
		if unicode.IsPrint(c) {
			if start == 0 {
				start = position
			}
			if c == unicode.ReplacementChar {
				buf = buf[0:0]
				start = 0
				continue
			}
			buf = append(buf, c)
			if len(buf) >= BUFSIZE {
				fmt.Printf("%d:%s ...\n", start, string(buf))
				buf = buf[0:0]
				start = 0
			}
		} else {
			if len(buf) >= *minSpan {
				fmt.Printf("%d:%s\n", start, string(buf))
			}
			buf = buf[0:0]
			start = 0
		}
	}
	if len(buf) >= *minSpan {
		fmt.Printf("%d:%s\n", start, string(buf))
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
				fatal(err)
			}
			makeString(f)
			f.Close()
		}
	}
}
