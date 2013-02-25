package cc

import (
	"errors"
)

var Compiler *compiler

type compiler struct {
	Autobn,
	Autoffset,
	Blockno,
	Lineno,
	Peekc,
	Nhunk,
	Nerrors int

	Iostack,
	Iofree []Io //This may be a *Io, not sure yet.
	IoNext *Io

	Types [NALLTYPES]*Type
}

func (c *compiler) Init() error {
	if c != nil {
		return errors.New("The compiler is already initialized.")
	}
	peekc = IGN

	types[TXXX] = new(Type)
	types[TCHAR] = typ(TCHAR, new(Type))
	types[TUCHAR] = typ(TUCHAR, new(Type))
	types[TSHORT] = typ(TSHORT, new(Type))
	types[TUSHORT] = typ(TUSHORT, new(Type))
	types[TINT] = typ(TINT, new(Type))
	types[TUINT] = typ(TUINT, new(Type))
	types[TLONG] = typ(TLONG, new(Type))
	types[TULONG] = typ(TULONG, new(Type))
	types[TVLONG] = typ(TVLONG, new(Type))
	types[TUVLONG] = typ(TUVLONG, new(Type))
	types[TFLOAT] = typ(TFLOAT, new(Type))
	types[TDOUBLE] = typ(TDOUBLE, new(Type))
	types[TVOID] = typ(TVOID, new(Type))
	types[TENUM] = typ(TENUM, new(Type))
	types[TFUNC] = typ(TFUNC, types[TINT])
	types[TIND] = typ(TIND, types[TVOID])

	return nil
}

func EnsureSymb(n int) {
	if symb == nil {
		symb = alloc(NSYMB + 1)
		nsymb = NSYMB
	}

	if n > nsymb {
		symb = allocn(symb, nsymb, n+1-nsymb)
		nsymb = n
	}
}
