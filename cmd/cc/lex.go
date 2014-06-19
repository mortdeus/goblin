package cc

import (
	"errors"
	"os"
)

var (
	Compiler *compiler
	symb     string

	itab = [...]struct {
		name    string
		lexical uint16
		type_   uint16
	}{
		{"auto", LAUTO, 0},
		{"break", LBREAK, 0},
		{"case", LCASE, 0},
		{"char", LCHAR, TCHAR},
		{"const", LCONSTNT, 0},
		{"continue", LCONTINUE, 0},
		{"default", LDEFAULT, 0},
		{"do", LDO, 0},
		{"double", LDOUBLE, TDOUBLE},
		{"else", LELSE, 0},
		{"enum", LENUM, 0},
		{"extern", LEXTERN, 0},
		{"float", LFLOAT, TFLOAT},
		{"for", LFOR, 0},
		{"goto", LGOTO, 0},
		{"if", LIF, 0},
		{"inline", LINLINE, 0},
		{"int", LINT, TINT},
		{"long", LLONG, TLONG},
		{"PREFETCH", LPREFETCH, 0},
		{"register", LREGISTER, 0},
		{"restrict", LRESTRICT, 0},
		{"return", LRETURN, 0},
		{"SET", LSET, 0},
		{"short", LSHORT, TSHORT},
		{"signed", LSIGNED, 0},
		{"signof", LSIGNOF, 0},
		{"sizeof", LSIZEOF, 0},
		{"static", LSTATIC, 0},
		{"struct", LSTRUCT, 0},
		{"switch", LSWITCH, 0},
		{"typedef", LTYPEDEF, 0},
		{"typestr", LTYPESTR, 0},
		{"union", LUNION, 0},
		{"unsigned", LUNSIGNED, 0},
		{"USED", LUSED, 0},
		{"void", LVOID, TVOID},
		{"volatile", LVOLATILE, 0},
		{"while", LWHILE, 0},
	}
	Types [NALLTYPES]*Type
)

type compiler struct {
	Autobn,
	Autoffset,
	Blockno,
	Lineno,
	Peekc,
	Nhunk,
	Nerrors int

	Io struct {
		Stack,
		Free []Io //This may be a *Io, not sure yet.
		Next *Io
	}

	OutBuf,
	DiagBuf,
	OutFile os.File
}

func (c *compiler) Init() error {
	if c != nil {
		return errors.New("The compiler is already initialized.")
	}

	c.Lineno = 1
	c.Peekc = IGN

	Types[TXXX] = new(Type)
	Types[TCHAR] = typ(TCHAR, new(Type))
	Types[TUCHAR] = typ(TUCHAR, new(Type))
	Types[TSHORT] = typ(TSHORT, new(Type))
	Types[TUSHORT] = typ(TUSHORT, new(Type))
	Types[TINT] = typ(TINT, new(Type))
	Types[TUINT] = typ(TUINT, new(Type))
	Types[TLONG] = typ(TLONG, new(Type))
	Types[TULONG] = typ(TULONG, new(Type))
	Types[TVLONG] = typ(TVLONG, new(Type))
	Types[TUVLONG] = typ(TUVLONG, new(Type))
	Types[TFLOAT] = typ(TFLOAT, new(Type))
	Types[TDOUBLE] = typ(TDOUBLE, new(Type))
	Types[TVOID] = typ(TVOID, new(Type))
	Types[TENUM] = typ(TENUM, new(Type))
	Types[TFUNC] = typ(TFUNC, Types[TINT])
	Types[TIND] = typ(TIND, Types[TVOID])

	for _, h := range hash {
		h = new(Sym)
	}

	for _, inst := range itab {
		s := slookup(inst.name)
		s.lexical = inst.lexical
		if inst.type_ != 0 {
			s.Type = c.Types[inst.type_]
		}
	}
	t := typ(TARRAY, c.Types[TCHAR])

	return nil
}
func slookup(key string) *Sym {
	symb = key
	return lookup()
}

func lookup() *Sym {

	if symb[0] == 0xc2 && symb[1] == 0xb7 {
		// turn leading · into ""·
		symb = append([]byte(`""`), symb)
	}
	h := uint(0)
	for i, s := range symb {
		if s == '·' {
			symb = append([]byte('.'), symb[i+1:])
		} else if string(symb[i:i+2]) == '∕' {
			symb = append([]byte('/'), symb[i+1:])
		}
	}
	sym, ok := hash[symb]
	if !ok {
		sym := new(Sym)
		sym.Name = symb
	}
}

//I dont think these functions are necessary in go.

// func slookup(s string){
// 	EnsureSymb(len(s))
// }

// func EnsureSymb(n int) {
// 	if symb == []byte("") {
// 		symb = make([]byte, NSYMB+1)
// 		nsymb = NSYMB
// 	}

// 	if n > nsymb {
// 		ext := make([]byte, n - nsymb)
// 		symb = append(symb, ext)
// 		nsymb = n
// 	}
// }
