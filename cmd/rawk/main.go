package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
)

var (
	flagF   = flag.String("F", "", "input field Separator regexp")
	flagV   = flag.String("v", "", "predefined vars (id=v)")
	flagf   = flag.String("f", "", "rawk script")
	scripts []string
	inFiles []*os.File
)

func main() {
	flag.Parse()
	if *flagF != "" {
		fmt.Println(*flagF)
	}
	if *flagV != "" {
		vdecl := strings.Split(*flagV, " ")
		fmt.Println(vdecl)
	}
	if *flagf != "" {
		fmt.Println(*flagf)
	}
	args := flag.Args()
	if len(args) == 0 {
		fmt.Println("missing files")
		os.Exit(2)
	}
	for _, arg := range args {
		if arg[0] == '{' {
			scripts = append(scripts, arg)
		} else {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			} else {
				inFiles = append(inFiles, f)
			}
			defer f.Close()
		}
	}
	fmt.Println(scripts)
	fmt.Println(inFiles)
}
