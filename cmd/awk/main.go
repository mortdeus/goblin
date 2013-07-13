package main

import (
	"fmt"
	"github.com/mortdeus/flag9"
	"os"
	"strconv"
)

func main() {
	c := newContext()
	flags(c)
	lex(c)
	fmt.Println(c.script.lexemes)
}

func usage() {
	fmt.Println("Usage: awk [ -F fs ][ -mf n ][ -mr n ][ -safe ][ -v var=value ][ -f progfile | prog ][ file ... ]")
	fmt.Println("\n" +
		"F:  \tinput field seperator regexp.\n" +
		"mf: \tmaximum number of fields.\n" +
		"mr: \tmaximum size of the input records.\n" +
		"safe: \tNo shell commands, open files, or the ENVIRON variable.\n" +
		"v:  \tPreallocated global variables.\n" +
		"f:  \tAwk script.\n" +
		"file: \tAccepts files, '-'(stdin), or 'var=value'.")
	os.Exit(2)
}

func fatal(err error) {
	panic(err)
	//fmt.Fprintf(os.Stderr, "%s:\t%s\n", "awk", err.Error())
	os.Exit(2)
}

func flags(ctx *context) {
	args := flag9.NewArgs(os.Args[1:])
	for args.Next() {
		switch args.Argc() {
		case 'h':
			usage()
		case 'F':
			if argf, ok := args.Argf(); ok {
				ctx.vars.tab["FS"].sval = argf
			} else {
				fatal(flagValErr(0))
			}
		case 'm':
			if args.Next() {
				c := args.Argc()
				if argf, ok := args.Argf(); ok {
					v, _ := strconv.Atoi(argf)
					if c == 'f' {
						ctx.maxField = v
					} else if c == 'r' {
						ctx.maxRecord = v
					}
					continue
				}
				fatal(flagValErr(0))
			}
		case 's':
			if argf, ok := args.Argf(); ok {
				if "s"+argf[0:3] == "safe" {
					ctx.safeMode = true
				}
			}
		case 'v':
			if argf, ok := args.Argf(); ok {
				for i, ch := range argf {
					if ch == '=' {
						id, val := argf[:i], argf[i:]
						if len(val) == 1 {
							val = ""
						} else {
							val = val[1:]
						}
						ctx.vars.add(id, val)
					}
				}
				continue
			}
			fatal(flagValErr(0))
		case 'f':
			if argf, ok := args.Argf(); ok {
				if f, err := os.Open(argf); err != nil {
					fatal(err)
				} else {
					_, err := ctx.script.buf.ReadFrom(f)
					if err != nil {
						fatal(err)
					}
				}
				continue
			}
			fatal(flagValErr(0))
		}

	}
	in := args.Argv()

	if ctx.script.buf.Len() <= 0 {
		ctx.script.buf.WriteString(in[0])
	}
	ctx.input = in[1:]
}
