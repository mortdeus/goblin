package main

import (
	"flag"
	"fmt"
	"log"
	"os"
)

var (
	ignore = flag.Bool("f", false,
		"Don't report files that can't be removed.")
	recurse = flag.Bool("r", false,
		"Recursively delete the entire contents of a directory and the directory itself.")
)

func usage() {
	fmt.Fprintf(os.Stderr, "Usage: rm [ –fr ] file ...\n")
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
	if len(args) == 0 {
		error("No files were given as input. Enter 'rm -help' for more information.\n")
	}
	for _, s := range args {
		info, err := os.Lstat(s)
		if err != nil {
			error(fmt.Sprintf("Lstat %s: %s", s, err))
		}
		if info.IsDir() && !*recurse {
			log.Println("Recurse flag must be set to remove directories. Skipping...")
			continue
		} else if info.IsDir() && *recurse {
			if err := os.RemoveAll(info.Name()); err != nil {
				error(fmt.Sprintf("RemoveAll %s: %s", s, err))
			}
		} else {
			if err := os.Remove(s); err != nil {
				error(fmt.Sprintf("Remove %s: %s", s, err))
			}
		}

	}
}
