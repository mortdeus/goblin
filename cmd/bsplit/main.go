package main

import (
	"bytes"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"strconv"
)

var (
	cmd = struct{ name, flags string }{
		"bsplit",
		"[ –p prefix ] [ –s size ] [ file ... ]",
	}
	prefix = flag.String("p", "bs",
		"Prefix for the set of output filenames returned by bsplit.")
	size = flag.Int("s", 524288, "Size of each contiguous block of data"+
		"to be split into seperate files in KB.")
)

func usage() {
	fmt.Fprintf(os.Stderr, "Usage:"+cmd.name+"\t"+cmd.flags)
	flag.PrintDefaults()
	os.Exit(2)
}

func fatal(err error) {
	log.Fatalln(cmd.name + ":\t" + err.Error())
}

func main() {
	flag.Usage = usage
	flag.Parse()
	args := flag.Args()
	data := make([]byte, 0)
	if len(args) == 0 {
		fatal(fmt.Errorf("%s", "No files were given as input."+
			" Enter 'bsplit -help' for more information.\n"))
	}
	for _, s := range args {
		f, err := os.Open(s)
		if err != nil {
			fatal(err)
		}
		if fdata, err := ioutil.ReadAll(f); err != nil {
			fatal(err)
		} else {
			data = append(data, fdata...)
		}
		if err := f.Close(); err != nil {
			fatal(err)
		}
	}
}

func bsplit(data []byte) {
	buf := bytes.NewBuffer(data)
	count := 0
	for {
		f, err := os.Create(*prefix + strconv.Itoa(count))
		if err != nil {
			fatal(err)
		}
		if _, err := io.CopyN(f, buf, int64(*size*1024)); err != nil {
			if err == io.EOF {
				return
			}
			fatal(err)
		}
		f.Close()
		count++

	}
}
