package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
	"unicode"
)

var (
	infile  = flag.String("if", "", "File to open for input")
	outfile = flag.String("of", "", "File to open for output")
	ibss     = flag.String("ibs", "", "Set input block size to n bytes (default 512)")
	obss     = flag.String("obs", "", "Set output block size to n bytes (default 512)")
	bss      = flag.String("bs", "", "Set both input and output block size, superseding ibs and obs.")
	skip    = flag.Int("skip", 0, "Skip n records before copying.")
	iseek   = flag.Int("iseek", 0, "Seek n records forward on input file before copying.")
	oseek   = flag.Int("oseek", 0, "Seek n records from beginning of output file before copying.")
	count   = flag.Int("count", 0, "Copy only n input records.")
	trunc   = flag.Int("trunc", 1, "By default, dd truncates the output file when it opens it; -trunc=0 opens without truncation.")
	quiet   = flag.Bool("quiet", false, "Don't print number of blocks read and written when finished.")
	//files = flag.Int("files", 1, "Catenate n input files.")
	conv     = flag.String("conv", "", "Set conversion mode.")
	conbufsz = flag.Int("cbs", 0, "Set conversion buffer size.")
	ibs = 512
	obs = 512
	bs = 0
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

func fatal(err error) {
	fmt.Fprint(os.Stderr, "dd:", err)
	os.Exit(1)
}

func errHandler(err error) {
	if err != nil {
		if convmod&noerror != noerror {
			fatal(err)
		} else {
			fmt.Fprint(os.Stderr,"dd:",err)
		}
	}
}

func convert_byte_string(s string) int {
	buf := make([]rune, 0, 128)
	var val int = 0
	var ret int = 0
	var prod bool = false
	
	s = strings.Join(strings.Fields(s), "")

	fmt.Println("A:",s)

	l := len(s)

	fmt.Println("l =",l)

	if l == 0 {
		fatal(nil)
	}

	for i:=0; i < l; i++ {
		if unicode.IsDigit(rune(s[i])) {
			buf = append(buf, rune(s[i]))
		} else {
			val, _ = strconv.Atoi(string(buf))
			buf = buf[0:0]
			
			switch(s[i]) {
			case 'k':
				val *= 1024
			case 'b':
				val *= 512
			case 'x':
				prod = true
				ret = val
			}
		}
	}
	fmt.Println("buf =",string(buf))
	val, _ = strconv.Atoi(string(buf))
	fmt.Println("val =",val)

	if prod {
		ret *= val
	} else {
		ret = val;
	}
	
	return ret
}

func main() {
	flag.Usage = usage
	flag.Parse()

	var in *os.File = os.Stdin
	var out *os.File = os.Stdout
	var err error

	fmt.Println("bss =",*bss)
	convstr := convert_byte_string(*bss)
	fmt.Println("bs =",convstr)

	if bs > 0 {
		ibs = bs
		obs = bs
	}

	if len(*infile) > 0 {
		in, err = os.Open(*infile)
		if err != nil {
			fatal(err)
		}

		if *iseek > 0 {
			_, err = in.Seek((int64)((*iseek)*(ibs)), 0)
			if err != nil {
				fatal(err)
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
			fatal(err)
		}

		if *oseek > 0 {
			_, err = out.Seek((int64)((*oseek)*(obs)), 0)
			if err != nil {
				fatal(err)
			}
		}
	}

	if len(*conv) > 0 {
		clist := strings.Split(*conv, ",")
		for _, opt := range clist {
			switch(opt) {
			case "lcase":
				convmod |= lcase
			case "ucase":
				convmod |= ucase
			case "swab":
				convmod |= swab
			case "noerror":
				convmod |= noerror
			case "sync":
				convmod |= sync
			}
		}
	}

	Dd(in, out)

	in.Close()
	out.Close()

	os.Exit(0)
}


func Dd(inf *os.File, outf *os.File) {
	ibuf := make([]byte, 0, ibs)
	obuf := make([]byte, 0, obs)
	in := bufio.NewReader(inf)
	out := bufio.NewWriter(outf)

	var incnt int = 0
	var outcnt int = 0

	var skipcnt int = 0

	var incblkin int = 0
	var incblkout int = 0

	passThrough := func() {
		_, err := out.Write(ibuf)
		errHandler(err)
		
		err = out.Flush()
		errHandler(err)
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

		if ibs == obs {
			passThrough()
			outcnt++
			return
		}

		for _, b := range ibuf {
			obuf = append(obuf, b)

			if len(obuf) >= obs {
				_, err := out.Write(obuf)
				errHandler(err)
				
				err = out.Flush()
				errHandler(err)
				
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

			if ibs == obs {
				passThrough()
				incblkout++
			} else {
				writeOut()
			}
			incblkin++
		}

		if len(obuf) > 0 {
			l := len(obuf)
			_, err := out.Write(obuf)

			errHandler(err)

			if convmod&sync == sync && (l<ibs || l<obs) {
				outcnt++
				for i := l; i<ibs || i<obs; i++ {
					err = out.WriteByte(0)
					errHandler(err)
				}
			}

			err = out.Flush()
			errHandler(err)
			
			incblkout++
		}
	}

	for {
		c, err := in.ReadByte()
		if err != nil {
			if err != io.EOF && convmod&noerror != noerror {
				fatal(err)
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
		if len(ibuf) >= ibs {
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
