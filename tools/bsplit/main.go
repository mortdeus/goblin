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
			error(fmt.Sprintf("open %s: %s", s, err))
		}
		if fdata, err := ioutil.ReadAll(f); err != nil {
			error(fmt.Sprintf("read %s: %s", s, err))
		} else {
			data = append(data, fdata...)
		}
		if err := f.Close(); err != nil {
			error(fmt.Sprintf("close %s: %s", s, err))
		}
	}
}

func bsplit(data []byte) {
	buf := bytes.NewBuffer(data)
	count := 0
	for {
		f, err := os.Create(*prefix + strconv.Itoa(count))
		if err != nil {
			error(fmt.Sprintf("create %s: %s", f.Name(), err))
		}
		if _, err := io.CopyN(f, buf, int64(*size*1024)); err != nil {
			if err == io.EOF {
				return
			}
			error(fmt.Sprintf("write %s: %s", f.Name(), err))
		}
		f.Close()
		count++

	}
}
