package main

import (
	"flag"
	"fmt"
	"io"
	"os"
	"os/signal"
)

var (
	cmd = struct{ name, flags string }{
		"tee",
		"[-ai] [file ...]",
	}
	iflag = flag.Bool("i", false, "ignore interrupts")
	aflag = flag.Bool("a", false, "append instead of overwriting")
	uflag = flag.Bool("u", false, "unix relic, ignored")
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
	if *iflag {
		c := make(chan os.Signal)
		signal.Notify(c, os.Interrupt)
	}
	args := flag.Args()
	files := make([]*os.File, len(args)+1)
	files[len(args)] = os.Stdout
	for i, path := range args {
		var err error
		var f *os.File
		if *aflag {
			f, err = os.OpenFile(path, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0666)
		} else {
			f, err = os.Create(path)
		}
		if err != nil {
			fmt.Fprintf(os.Stderr, "%s, ignoring\n", err)
			continue
		}
		files[i] = f
	}

	buf := make([]byte, 8*1024)
	for {
		n, err := os.Stdin.Read(buf)
		if err == io.EOF {
			break
		}
		if err != nil {
			fatal(err)
		}
		for i, f := range files {
			if f == nil {
				continue
			}
			_, err := f.Write(buf[:n])
			if err != nil {
				fmt.Fprintf(os.Stderr, "%s - dropping\n", err)
				files[i] = nil
			}
		}
	}
}
