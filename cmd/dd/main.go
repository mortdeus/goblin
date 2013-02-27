package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"strings"
)

var (
	infile  = flag.String("if", "", "File to open for input")
	outfile = flag.String("of", "", "File to open for output")
	ibs     = flag.Int("ibs", 512, "Set input block size to n bytes (default 512)")
	obs     = flag.Int("obs", 512, "Set output block size to n bytes (default 512)")
	bs      = flag.Int("bs", 0, "Set both input and output block size, superseding ibs and obs.")
	skip    = flag.Int("skip", 0, "Skip n records before copying.")
	iseek   = flag.Int("iseek", 0, "Seek n records forward on input file before copying.")
	oseek   = flag.Int("oseek", 0, "Seek n records from beginning of output file before copying.")
	count   = flag.Int("count", 0, "Copy only n input records.")
	trunc   = flag.Int("trunc", 1, "By default, dd truncates the output file when it opens it; -trunc=0 opens without truncation.")
	quiet   = flag.Bool("quiet", false, "Don't print number of blocks read and written when finished.")
	//files = flag.Int("files", 1, "Catenate n input files.")
	conv     = flag.String("conv", "", "Set conversion mode.")
	conbufsz = flag.Int("cbs", 0, "Set conversion buffer size.")
)

const (
	lcase   = 1 << iota
	ucase   = 1 << iota
	swab    = 1 << iota
	noerror = 1 << iota
	sync    = 1 << iota
)

var convmod int = 0

func usage() {
	fmt.Fprintf(os.Stderr, "Usage: dd [ option value ] ...\n")
	flag.PrintDefaults()
	os.Exit(2)
}

func errExit(err error) {
	fmt.Fprint(os.Stderr, "dd:", err)
	os.Exit(1)
}

func main() {
	flag.Usage = usage
	flag.Parse()

	var in *os.File = os.Stdin
	var out *os.File = os.Stdout
	var err error

	if *bs > 0 {
		*ibs = *bs
		*obs = *bs
	}

	if len(*infile) > 0 {
		in, err = os.Open(*infile)
		if err != nil {
			errExit(err)
		}

		if *iseek > 0 {
			_, err = in.Seek((int64)((*iseek)*(*ibs)), 0)
			if err != nil {
				errExit(err)
			}
		}
	}

	if len(*outfile) > 0 {
		if *trunc == 1 {
			out, err = os.Create(*outfile)
		} else {
			out, err = os.OpenFile(*outfile, os.O_WRONLY, 0666)
		}
		if err != nil {
			errExit(err)
		}

		if *oseek > 0 {
			_, err = out.Seek((int64)((*oseek)*(*obs)), 0)
			if err != nil {
				errExit(err)
			}
		}
	}

	if len(*conv) > 0 {
		clist := strings.Split(*conv, ",")
		for _, opt := range clist {
			if opt == "lcase" {
				convmod |= lcase
			} else if opt == "ucase" {
				convmod |= ucase
			} else if opt == "swab" {
				convmod |= swab
			} else if opt == "noerror" {
				convmod |= noerror
			} else if opt == "sync" {
				convmod |= sync
			}
		}
	}

	dd(in, out)

	in.Close()
	out.Close()

	os.Exit(0)
}

func dd(inf *os.File, outf *os.File) {
	ibuf := make([]byte, 0)
	obuf := make([]byte, 0)
	in := bufio.NewReader(inf)
	out := bufio.NewWriter(outf)

	var incnt int = 0
	var outcnt int = 0

	var skipcnt int = 0

	var incblkin int = 0
	var incblkout int = 0

	passThrough := func() {
		out.Write(ibuf)
		out.Flush()
	}

	uc := func(b *byte) {
		if *b >= 'a' && *b <= 'z' {
			*b -= 32
		}
	}

	lc := func(b *byte) {
		if *b >= 'A' && *b <= 'Z' {
			*b += 32
		}
	}

	swb := func() {
		l := len(ibuf)
		var a byte
		for i := 0; i < l; i++ {
			a = ibuf[i]
			i++
			ibuf[i-1] = ibuf[i]
			ibuf[i] = a
		}
	}

	writeOut := func() {
		if convmod&swab == swab {
			swb()
		}

		if *ibs == *obs {
			passThrough()
			outcnt++
			return
		}

		for _, b := range ibuf {
			obuf = append(obuf, b)

			if len(obuf) >= *obs {
				out.Write(obuf)
				out.Flush()
				obuf = obuf[0:0]
				outcnt++
			}
		}
	}

	finishUp := func() {
		if len(ibuf) > 0 {
			if convmod&swab == swab {
				swb()
			}

			if *ibs == *obs {
				passThrough()
				incblkout++
			} else {
				writeOut()
			}
			incblkin++
		}

		if len(obuf) > 0 {
			out.Write(obuf)
			out.Flush()
			incblkout++
		}
	}

	for {
		c, err := in.ReadByte()
		if err != nil {
			if err != io.EOF && convmod&noerror != noerror {
				errExit(err)
				continue
			}
			finishUp()
			break
		}

		if convmod&lcase == lcase {
			lc(&c)
		}

		if convmod&ucase == ucase {
			uc(&c)
		}

		ibuf = append(ibuf, c)
		if len(ibuf) >= *ibs {
			if skipcnt < *skip {
				skipcnt++
			} else {
				writeOut()
				incnt++
			}
			ibuf = ibuf[0:0]
		}

		if *count > 0 && incnt >= *count {
			finishUp()
			break
		}
	}

	if !*quiet {
		fmt.Fprintf(os.Stderr, "%d+%d records in\n", incnt, incblkin)
		fmt.Fprintf(os.Stderr, "%d+%d records out\n", outcnt, incblkout)
	}
}
