package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	f1, err := os.Open("test1")
	if err != nil {
		panic(err)
	}
	f2, err := os.Open("test2")
	if err != nil {
		panic(err)
	}
	Diff(f1, f2)
}

func Diff(r1, r2 io.Reader) io.Reader {
	f1 := &File{Equiv: make(map[int]*equivClass, 0)}
	f2 := &File{Equiv: make(map[int]*equivClass, 0)}

	scan := func(f *File, r io.Reader) {
		s := bufio.NewScanner(r)
		var x int
		for s.Scan() {
			f.Line = append(f.Line, NewLine(s.Bytes(), x))
			x++
		}
		if err := s.Err(); err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
	}
	scan(f1, r2)
	scan(f2, r1)
	f2.Sort()

	return nil

}
