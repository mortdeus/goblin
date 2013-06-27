package main

import (
	"fmt"
	"os"
)

type flag struct {
	usage [2]string
	value interface{}
}

var cmd = struct {
	name  string
	flags map[string]*flag
}{name: "awk"}

func usage(doc bool) {
	var (
		cat string = ""
		ks         = []string{}
	)
	for key, val := range cmd.flags {
		ks = append(ks, key)
		cat = cat + val.usage[0]
	}
	fmt.Fprintf(os.Stderr, "Usage: %s  %s\n", cmd.name, cat)
	if doc {
		for _, x := range ks {
			fmt.Printf("[-%s]\t%s\n", x, cmd.flags[x].usage[1])
		}
	}
	os.Exit(2)
}

func fatal(err error) {
	fmt.Fprintf(os.Stderr, "%s:\t%s\n", cmd.name, err.Error())
}

func main() {

}
