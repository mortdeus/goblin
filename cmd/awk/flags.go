package main

import (
	"fmt"
	"github.com/mortdeus/flag9"
	"os"
	"strconv"
)

var (
	safeMode            bool
	maxField, maxRecord int
)

func processFlags() {
	args := flag9.NewArgs(os.Args[1:])
	for args.Next() {
		switch args.Argc() {
		case 'h':
			usage()
		case 'F':
			if argf, ok := args.Argf(); ok {
				vTable["FS"].sval = argf
			} else {
				fatal(fmt.Errorf("No flag value."))
			}
		case 'm':
			if args.Next() {
				c := args.Argc()
				if argf, ok := args.Argf(); ok {
					v, _ := strconv.Atoi(argf)

					if c == 'f' {
						maxField = v
					} else if c == 'r' {
						maxRecord = v
					}
					continue
				}
				fatal(fmt.Errorf("No flag value."))
			}
		case 's':
			if argf, ok := args.Argf(); ok {
				if "s"+argf[0:3] == "safe" {
					safeMode = true
				}
			}
		case 'v':
		case 'f':
		}
	}
}
