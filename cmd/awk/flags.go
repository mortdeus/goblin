package main

import (
	// "fmt"
	"github.com/mortdeus/flag9"
	"os"
	"strconv"
)

func init() {
	var (
		s    = os.Args
		args = flag9.NewArgs(s[1:])

		fl []rune
		f  string
	)
	cmd.flags = map[string]*flag{
		"F":    &flag{[2]string{"[ -F fs ]", "input field seperator regexp."}, " "},
		"mf":   &flag{[2]string{"[ -mf n ]", "maximum number of fields."}, -1},
		"mr":   &flag{[2]string{"[ -mr n ]", "maximum size of the input records."}, -1},
		"safe": &flag{[2]string{"[ -safe ]", "No shell commands, open files, or the ENVIRON variable."}, false},
		"v":    &flag{[2]string{"[ -v var=value ]", "Preallocated global variables."}, varTab},
		"f":    &flag{[2]string{"[ -f progfile | prog ]", "Awk script."}, "{}"},
		"":     &flag{[2]string{"[ file ... ]", "Accepts files, '-'(stdin), or 'var=value'"}, varTab},
	}
	for args.Next() {
		fl = append(fl, args.Argc())
		switch f = string(fl); f {
		case "h":
			usage(true)
		case "F", "f":
			if val, ok := args.Argf(); ok {
				if val[0] == '=' {
					val = val[1:]
				}
				cmd.flags[f].value = val
			}
			fl = []rune{}

		case "mf", "mr":
			if val, ok := args.Argf(); ok {
				if val[0] == '=' {
					val = val[1:]
				}
				if i, err := strconv.Atoi(val); err != nil {
					panic(err)
				} else {
					cmd.flags[f].value = i
				}
			}
			fl = []rune{}

		case "v":
			if val, ok := args.Argf(); ok {
				//TODO: Implement -v
				_ = val
			}
			fl = []rune{}
		case "safe":
			cmd.flags[f].value = true
			fl = []rune{}
		}
	}

}
