package main

import (
	"fmt"
	"go/build"
	"log"
	"os"
	"os/exec"
	//"runtime"
	"path"
	"strings"
)

func init() {
	log.SetPrefix("goblin:\t  ")
	log.SetFlags(0)
}

func getPkgPath() (string, error) {
	p, err := exec.LookPath(os.Args[0])
	if err != nil {
		return "", err
	}
	cmd := exec.Command("go", "tool", "objdump", "-s=main.main", p)
	out, err := cmd.Output()
	if err != nil {
		return "", err
	}

	s := strings.SplitN(fmt.Sprintf("%s", out), "\n", 2)[0]
	s = s[len("TEXT main.main(SB) "):]

	//makesure this source file's name matches whats hardcoded here.
	// Otherwise things break.
	return s[:len(s)-len("/tool.go")], nil
}

func usage() (s string) {
	install := "install:\tinstalls the goblin tools to $GOTOOLDIR"
	return fmt.Sprintln(install)
}

func main() {
	if len(os.Args) == 1 {
		log.Println(usage())
		return
	}
	for _, arg := range os.Args[1:] {
		switch arg {
		case "install":
			install()
		default:
			log.Println("invalid input arg:", arg)
			log.Println(usage())
			return
		}
	}
}

func install() {
	if p, err := getPkgPath(); err != nil {
		log.Fatal(err)
	} else {
		if err := os.Chdir(p); err != nil {
			log.Fatal(err)
		}
	}

	if err := os.Chdir("cmd"); err != nil {
		log.Fatal(err)
	}
	f, err := os.Open(".")
	if err != nil {
		log.Fatal(err)
	}
	dirs, err := f.Readdirnames(0)
	if err != nil {
		log.Fatal(err)
	}
	for _, d := range dirs {
		if err := os.Chdir(d); err != nil {
			log.Fatal(err)
		}
		if d, err := os.Getwd(); err != nil {
			log.Fatal(err)
		} else {
			log.Println(d)
			nm := path.Base(d)
			out, err := exec.Command("go", "build").CombinedOutput()
			if err != nil {
				log.Printf("go build %s failed\n\n", nm)
				fmt.Fprintf(os.Stderr, "[go-build-log]\n%s", out)
				goto quit
			}

			log.Printf("mv %s to $GOTOOLDIR\n", nm)
			if err := os.Rename(d+"/"+nm, build.ToolDir+"/"+nm); err != nil {
				log.Fatal(err)
			}
		}
	quit:
		if err := os.Chdir(".."); err != nil {
			log.Fatal(err)
		}
		fmt.Println()
	}
}
