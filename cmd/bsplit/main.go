package main

import (
	"bytes"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strconv"
)

var (
	prefix = flag.String("p", "bs",
		"Prefix for the set of output filenames returned by bsplit.")
	size = flag.Int("s", 524288, "Size of each contiguous block of data"+
		"to be split into seperate files in KB.")
)

func usage() {
	fmt.Fprintf(os.Stderr,
		"Usage:	bsplit [ –p prefix ] [ –s size ] [ file ... ]\n")
	flag.PrintDefaults()
	os.Exit(2)
}

func error(s string) {
	fmt.Fprint(os.Stderr, s, "\n")
	os.Exit(1)
}

func main() {
	flag.Usage = usage
	flag.Parse()
	args := flag.Args()
	data := make([]byte, 0)
	if len(args) == 0 {
		error("No files were given as input. Enter 'bsplit -help' for more information.\n")
	}
	for _, s := range args {
		f, err := os.Open(s)
		if err != nil {
			error(err.Error())
		}
		if fdata, err := ioutil.ReadAll(f); err != nil {
			error(err.Error())
		} else {
			data = append(data, fdata...)
		}
		if err := f.Close(); err != nil {
			error(err.Error())
		}
	}
}

func bsplit(data []byte) {
	buf := bytes.NewBuffer(data)
	count := 0
	for {
		f, err := os.Create(*prefix + strconv.Itoa(count))
		if err != nil {
			error(err.Error())
		}
		if _, err := io.CopyN(f, buf, int64(*size*1024)); err != nil {
			if err == io.EOF {
				return
			}
			error(err.Error())
		}
		f.Close()
		count++

	}
}
