package main

import (
	"flag"
	"fmt"
	"log"
	"os"
)

var (
	cmd = struct{name, flags string}{
		"rm",
		"[ -fr ] file...",
	}
	ignore = flag.Bool("f", false,
		"Don't report files that can't be removed.")
	recurse = flag.Bool("r", false,
		"Recursively delete the entire contents of a directory and the directory itself.")
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
		usage()
	}
	for _, s := range args {
		info, err := os.Lstat(s)
		if err != nil {
			fatal(err)
		}
		if info.IsDir() && !*recurse {
			log.Println("Recurse flag must be set to remove directories. Skipping...")
			continue
		} else if info.IsDir() && *recurse {
			if err := os.RemoveAll(info.Name()); err != nil {
				fatal(err)
			}
		} else {
			if err := os.Remove(s); err != nil {
				fatal(err)
			}
		}

	}
}
