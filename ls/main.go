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

func ls(path string) {
	var uwidth int
	var gwidth int
	var swidth int

	/* Probably not portable */
	linuxGetUserGroupInfo := func(fi os.FileInfo) (string, uint32) {
		s := fi.Sys().(*syscall.Stat_t)
		u, _ := user.LookupId(strconv.Itoa(int(s.Uid)))

		if len(u.Username) > uwidth {
			uwidth = len(u.Username)
		}

		gl := strconv.Itoa(int(s.Gid))
		if len(gl) > gwidth {
			gwidth = len(gl)
		}

		return u.Username, s.Gid
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
				usr, grp := linuxGetUserGroupInfo(file)
				sze := file.Size()
				if len(strconv.Itoa(int(sze))) > swidth {
					swidth = len(strconv.Itoa(int(sze)))
				}
				fmt.Printf("%s %*s %*d %*d ",
					file.Mode().String(),
					uwidth, usr,
					gwidth, grp,
					swidth, sze)
			}
			fmt.Printf("%s\n", file.Name())
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
