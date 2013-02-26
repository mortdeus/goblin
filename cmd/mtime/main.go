package main

import (
	"fmt"
	"os"
)

func main() {
	for _, s := range os.Args[1:] {
		fi, err := os.Stat(s)
		if err != nil {
			fmt.Fprintln(os.Stderr, "mtime:", err.Error())
		}
		mt := fi.ModTime()
		fmt.Printf("%d %s", mt, s)
	}
}
