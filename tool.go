package main

import (
	"fmt"
	"os"
	//"runtime"
)

var goblinroot string = os.Getenv("GOBLIN")
var cmds = map[string]string{
	"install": "installs core goblin system",
	"update":  "update core goblin system",
	"backup":  "create a compressed snapshot of the current goblin filesystem",
	"get":     "fetch go packages from goblin repo",
	"env":     "print information pertaining to the goblin environment",
}

func usage() {
	for k, v := range cmds {
		fmt.Printf("[%s]:\n%s\n\n", k, v)
	}
}

func main() {
	for i := range os.Args {
		if i == 0 {
			continue
		}
		switch os.Args[i] {
		case "install":
			install()
		case "update":
			update()
		case "backup":
			backup()
		case "get":
			get()
		case "env":
			env()
		case "doc":
			doc()
		default:
			usage()
			return
		}

	}
}

func install() {
	unimplemented()
}
func update() {
	unimplemented()
}
func backup() {
	unimplemented()
}
func get() {
	unimplemented()
}
func env() {
	unimplemented()
}
func doc() {
	unimplemented()
}

func unimplemented() {
	fmt.Println(os.Args[1], "is not implemented yet")
	os.Exit(2)
}
