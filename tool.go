package main

import (
	"fmt"
	"go/build"
	"log"
	"os"
	"os/exec"
	//"runtime"
	"path/filepath"
	"strings"
)

func init() {
	log.SetPrefix("[goblin]  ")
	log.SetFlags(0)
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

var pkgpath string
var cmdpath string

func getPkgPath() error {
	p, err := exec.LookPath(os.Args[0])
	if err != nil {
		return err
	}
	cmd := exec.Command("go", "tool", "objdump", "-s=main.main", p)
	out, err := cmd.Output()
	if err != nil {
		return err
	}
	s := strings.SplitN(fmt.Sprintf("%s", out), "\n", 2)[0]
	pkgpath = filepath.Dir(s[len("TEXT main.main(SB) "):])
	cmdpath = pkgpath + "/cmd"
	return nil
}

func install() {
	if err := getPkgPath(); err != nil {
		log.Fatal(err)
	}
	if err := walkCmds(cmdpath); err != nil {
		log.Fatal(err)
	}

}
func getDirs(path string) ([]string, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	fids, err := f.Readdir(0)
	if err != nil {
		return nil, err
	}
	var dirs []string
	for _, finfo := range fids {
		if finfo.IsDir() {
			dirs = append(dirs, finfo.Name())
		}
	}
	return dirs, nil
}

func walkCmds(dpath string) error {
	fmt.Println("----------------------------------------------------------------------------------")
	log.Println("visiting", dpath)
	dirs, err := getDirs(dpath)
	if err != nil {
		log.Fatal(err)
	}
	for _, d := range dirs {
		if err := walkCmds(dpath + "/" + d); err != nil {
			return err
		}
	}
	if gosrc, err := filepath.Glob(dpath + "/*.go"); err != nil {
		log.Fatal(err)
	} else if len(gosrc) == 0 {
		return nil
	}
	var nm string
	p := strings.Split(dpath, cmdpath)
	if len(p) > 1 {
		p = strings.Split(p[1], string(os.PathSeparator))
		if len(p) > 1 {
			nm = p[len(p)-2] + p[len(p)-1]
		} else {
			nm = p[0]
		}
	}
	log.Printf("chdir(%s)", dpath)
	if err := os.Chdir(dpath); err != nil {
		log.Fatal(err)
	}
	defer os.Chdir("..")
	cmd := exec.Command("go", "build", "-v")
	cmd.Stdout, cmd.Stderr = os.Stdout, os.Stderr
	fmt.Fprintf(os.Stderr, "\n[go-build-log]\n")
	if cmd.Run(); err != nil {
		log.Printf("go build %s failed\n\n", nm)
		return err
	}
	fmt.Fprint(os.Stderr, "\n")

	log.Println("installing", nm, "to $GOTOOLDIR")
	fmt.Println("__________________________________________________________________________________")
	os.Rename(dpath+"/"+filepath.Base(dpath), build.ToolDir+"/"+nm)
	return nil
}
