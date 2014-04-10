package main

import (
	"flag"
	"fmt"
	"io"
	"os"
	"os/user"
	"sort"
	"strconv"
	"syscall"
	"time"
)

/*
 * We only read in a certain number of dirents at a time, which should make
 * things a little more robust on unreliable file systems.
 */
const MAXDIRREAD int = 50

var (
	cmd = struct{ name, flags string }{
		"ls",
		"[-dlmnprstuF] [file ...]",
	}

	uwidth int
	gwidth int
	swidth int
	dwidth int

	usedir   = flag.Bool("d", false, "list directory instead of its contents")
	long     = flag.Bool("l", false, "use long format")
	usrname  = flag.Bool("m", false, "list user who last modified the file")
	nosort   = flag.Bool("n", false, "don't sort list")
	nopath   = flag.Bool("p", false, "only print the last path element")
	reverse  = flag.Bool("r", false, "reverse sorting order")
	kbytes   = flag.Bool("s", false, "give size in KBytes for each file")
	timesort = flag.Bool("t", false, "sort by latest-modified first")
	useatime = flag.Bool("u", false, "use atime for sorting / printing")
	tlslash = flag.Bool("f", false, "add / after directories and * after executables")
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

func main() {
	flag.Usage = usage
	flag.Parse()
	args := flag.Args()

	if len(args) == 0 {
		ls("")
	} else {
		for _, path := range args {
			ls(path)
		}
	}
}

type dent struct {
	mode string
	p    string
	u    string
	dev  uint32
	Atim syscall.Timespec
	Ctim syscall.Timespec
	t    time.Time
	qver int32
	qpth uint64
	g    uint32
	s    int64
	n    string
	d    bool
	e    bool
}

func (self *dent) getInfo(fi os.FileInfo) {
	self.n = fi.Name()
	self.s = fi.Size()

	l := len(strconv.Itoa(int(self.s)))
	if l > swidth {
		swidth = l
	}

	self.mode = fi.Mode().String()
	self.d = fi.Mode().IsDir()
	self.t = fi.ModTime()

	if fi.Mode().Perm()&0111 != 0 && !self.d {
		self.e = true
	}

	/* The following is probably not portable */
	s := fi.Sys().(*syscall.Stat_t)

	self.Ctim = s.Ctim
	self.Atim = s.Atim
	self.qver = s.Mtim.Sec + s.Ctim.Sec
	self.qpth = s.Ino
	self.dev = s.Nlink

	devs := strconv.Itoa(int(s.Dev))
	if len(devs) > dwidth {
		dwidth = len(devs)
	}

	if *useatime {
		self.t = time.Unix(int64(self.Atim.Sec), int64(self.Atim.Nsec))
	}

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
	/* Put the time parsing here so we can sort easier elsewhere */
	yr := self.t.Year()
	mon := self.t.Format("Jan")
	day := self.t.Day()
	hr, min, _ := self.t.Clock()
	var m string

	if time.Now().Year() == yr {
		m = fmt.Sprintf("%s %2d %02d:%02d", mon, day, hr, min)
	} else {
		m = fmt.Sprintf("%s %2d %5d", mon, day, yr)
	}

	return fmt.Sprintf("%s %*d %*s %*d %*d %s",
		self.mode,
		dwidth, self.dev,
		uwidth, self.u,
		gwidth, self.g,
		swidth, self.s,
		m)
}

/* Sorting Helpers */
type dents []*dent

func (d dents) Len() int      { return len(d) }
func (d dents) Swap(i, j int) { d[i], d[j] = d[j], d[i] }

type ByName struct{ dents }

func (s ByName) Less(i, j int) bool { return s.dents[i].n < s.dents[j].n }

type ByMTime struct{ dents }

func (s ByMTime) Less(i, j int) bool {
	return s.dents[i].t.Unix() > s.dents[j].t.Unix()
}

type Reverse struct{ sort.Interface }

func (r Reverse) Less(i, j int) bool { return r.Interface.Less(j, i) }

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

	printFName := func(d *dent) {
		fmt.Printf("%s%s", d.p, d.n)
		if *tlslash {
			if d.d {
				fmt.Print("/")
			} else if d.e {
				fmt.Print("*")
			}
		}
		fmt.Print("\n")
	}

	processDent := func(file os.FileInfo, p string) {
		d := new(dent)

		if *nopath {
			p = ""
		}
		d.p = p
		if p != "" {
			d.p += "/"
		}

		d.getInfo(file)
		addDent(d)
	}

	var pth string
	if path == "" {
		pth = "."
	} else {
		pth = path
	}

	f, err := os.Open(pth)
	if err != nil {
		fatal(err)
	}

	s, err := f.Stat()
	if err != nil {
		fatal(err)
	}

	if !s.Mode().IsDir() || *usedir {
		processDent(s, "")
	} else {
		for {
			fi, err := f.Readdir(MAXDIRREAD)
			if err != nil {
				if err == io.EOF {
					break
				} else {
					fatal(err)
				}
			}

			for _, file := range fi {
				processDent(file, path)
			}
		}

		if !*nosort {
			if !*timesort {
				if !*reverse {
					sort.Sort(ByName{dents})
				} else {
					sort.Sort(Reverse{ByName{dents}})
				}
			} else {
				if !*reverse {
					sort.Sort(ByMTime{dents})
				} else {
					sort.Sort(Reverse{ByMTime{dents}})
				}
			}
		}
	}

	for _, d := range dents {
		if *kbytes {
			nswidth := swidth - 3
			if nswidth < 1 {
				nswidth = 1
			}
			fmt.Printf("%*d ", nswidth, d.s/1024)
		}

		/* Provided for compatibility only */
		if *usrname {
			fmt.Print("[] ")
		}

		if *long {
			fmt.Printf("%s ", d.String())
		}

		printFName(d)
	}
}
