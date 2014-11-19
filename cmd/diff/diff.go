package main

import (
	"bufio"
	"fmt"
	"hash/adler32"
	"io"
	"os"
	"sort"
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
	scan(f1, r1)
	scan(f2, r2)

	fmt.Println(f2, "\n")
	f2.Sort()
	fmt.Println(f2, "\n")
	return nil

}

type File struct {
	Line  []line
	Equiv map[int]*equivClass
}

func (f *File) Sort() {
	sort.Sort(f)

	for hash, class := range f.Equiv {
		for i := 0; ; i++ {
			if hash+i > len(f.Line) || f.Line[hash+i].Hash == hash {
				break
			}
			for _, ln := range f.Line[hash : hash+i] {
				class.index = append(class.index, ln.Offset)
			}
			sort.Ints(class.index)
		}
	}
}
func (f *File) Len() int {
	return len(f.Line)
}
func (f *File) Less(i, j int) bool {
	il := f.Line[i]
	jl := f.Line[j]

	if il.Hash < jl.Hash {
		return true
	}
	if il.Hash == jl.Hash {
		if _, ok := f.Equiv[il.Hash]; !ok {
			f.Equiv[il.Hash] = new(equivClass)
		}
		if il.Offset < jl.Offset {
			return true
		}
	}
	return false
}
func (f *File) Swap(i, j int) {
	tmp := f.Line[i]
	f.Line[i] = f.Line[j]
	f.Line[j] = tmp
}

type line struct {
	Hash   int
	Offset int
}

func NewLine(s []byte, pos int) line { return line{int(adler32.Checksum(s)), pos} }

type equivClass struct {
	index []int
}
