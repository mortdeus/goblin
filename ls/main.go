package main

import (
	"flag"
	"fmt"
	"io"
	"os"
	"os/user"
	"strconv"
	"syscall"
)

const MAXDIRREAD int = 50

var uwidth int
var gwidth int
var swidth int

var (
	usedir   = flag.Bool("d", false, "List a directory instead of its contents.")
	long     = flag.Bool("l", false, "Long format list.")
	usrname  = flag.Bool("m", false, "List the user who last modified the file.")
	nosort   = flag.Bool("n", false, "Don't sort the list.")
	fnlpath  = flag.Bool("p", false, "Only print the last path element.")
	reverse  = flag.Bool("r", false, "Reverse the sorting order.")
	kbytes   = flag.Bool("s", false, "Use KBytes for size.")
	timesort = flag.Bool("t", false, "Sort by latest-modified first.")
	useatime = flag.Bool("u", false, "If -t sort by access time; if -u print"+
		"last access time.")
	tlslash = flag.Bool("F", false, "Add / after all directories and * after"+
		"all executables.")
)

func usage() {
	fmt.Fprintf(os.Stderr, "Usage: ls [-dlmnprstuF] [file ...]\n")
	flag.PrintDefaults()
	os.Exit(2)
}

func error(s string) {
	fmt.Fprint(os.Stderr, s, "\n")
	os.Exit(1)
}

type dent struct {
	mode string
	u    string
	g    uint32
	s    int64
	n    string
}

/* Probably not portable */
func (self *dent) getInfo(fi os.FileInfo) {
	self.n = fi.Name()
	self.s = fi.Size()
	self.mode = fi.Mode().String()

	l := len(strconv.Itoa(int(self.s)))
	if l > swidth {
		swidth = l
	}

	s := fi.Sys().(*syscall.Stat_t)
	u, _ := user.LookupId(strconv.Itoa(int(s.Uid)))

	if len(u.Username) > uwidth {
		uwidth = len(u.Username)
	}

	gl := strconv.Itoa(int(s.Gid))
	if len(gl) > gwidth {
		gwidth = len(gl)
	}

	self.u = u.Username
	self.g = s.Gid
}

func (self *dent) String() string {
	return fmt.Sprintf("%s %*s %*d %*d %s",
		self.mode,
		uwidth, self.u,
		gwidth, self.g,
		swidth, self.s,
		self.n)
}

func ls(path string) {
	dents := make([]*dent, 0, MAXDIRREAD)

	addDent := func(d *dent) {
		l := len(dents)
		if l+1 > cap(dents) {
			n_dents := make([]*dent, l+MAXDIRREAD)
			copy(n_dents, dents)
			dents = n_dents
		}
		dents = dents[0 : l+1]
		dents[l] = d
	}

	f, err := os.Open(path)
	if err != nil {
		error(fmt.Sprint("%s", err))
	}

	for {
		fi, err := f.Readdir(MAXDIRREAD)
		if err != nil {
			if err == io.EOF {
				break
			} else {
				error(fmt.Sprint("%s", err))
			}
		}

		for _, file := range fi {
			if *long {
				d := new(dent)
				d.getInfo(file)
				addDent(d)
			} else {
				fmt.Printf("%s\n", file.Name())
			}
		}
	}

	if *long {
		for _, d := range dents {
			fmt.Println(d.String())
		}
	}
}

func main() {
	flag.Usage = usage
	flag.Parse()
	args := flag.Args()

	if len(args) == 0 {
		ls(".")
	} else {
		for _, path := range args {
			ls(path)
		}
	}

	os.Exit(0)
}
