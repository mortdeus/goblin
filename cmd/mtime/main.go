package main

import (
	"os"
	"fmt"
)

func main() {
	for _, s := range os.Args[1:] {
		fi, _ := os.Stat(s)
		mt := fi.ModTime()
		fmt.Printf("%d %s", mt, s)
	}
}
