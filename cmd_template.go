//+build ignore
package main

import (
	"flag"
	"fmt"
	"log"
	"os"
)

var (
	cmd = struct{ name, flags string }{
		"name",
		"[ –f foo] [ –b bar ] [ file ... ]",
	}
	//Flags
	fflag = flag.String(name, value, usage)
	bflag = flag.Bool(name, value, usage)
)

func usage() {
	fmt.Fprintf(os.Stderr, "Usage:"+cmd.name+"\t"+cmd.flags)
	flag.PrintDefaults()
	os.Exit(2)
}

func fatal(err error) {
	//silent flag
	//if *sflag {
	//	os.Exit(1)
	//}

	log.Fatalln(cmd.name + ":\t" + err.Error())
}

func main() {
	flag.Usage = usage
	flag.Parse()
	args := flag.Args()
}
