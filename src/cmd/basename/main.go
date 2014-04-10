package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
)

var (
	cmd = struct{ name, flags string }{
		"basename",
		"[ -d ] string [ suffix ]",
	}
	dflag = flag.Bool("d", false, "print directories, not file")
)

func usage() {
	fmt.Fprintln(os.Stderr, "Usage:", cmd.name, cmd.flags)
	flag.PrintDefaults()
	os.Exit(2)
}

func main() {
	flag.Usage = usage
	flag.Parse()
	args := flag.Args()
	if len(args) != 1 && len(args) != 2 {
		usage()
	}

	slash := strings.LastIndex(args[0], "/")
	pr := ""
	if slash >= 0 {
		pr = args[0][slash:]
	}
	if *dflag {
		if pr != "" {
			fmt.Print(args[0][:slash], "\n")
		} else {
			fmt.Print(".\n")
		}
		return
	}
	if pr != "" {
		pr = pr[1:]
	} else {
		pr = args[0]
	}
	if len(args) == 2 && strings.HasSuffix(pr, args[1]) {
		pr = pr[:len(pr)-len(args[1])]
	}
	fmt.Print(pr, "\n")
}
