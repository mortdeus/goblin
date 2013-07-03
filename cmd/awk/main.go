package main

import (
	"fmt"
	"os"
)

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
	fmt.Fprintf(os.Stderr, "%s:\t%s\n", "awk", err.Error())
	usage()
}

func main() {
	processFlags()
}
