package cc

import (
	"unsafe"
)

const (
	BUFSIZ     = 8192
	NSYMB      = 500
	NHASH      = 1024
	STRINGSZ   = 200
	HISTSZ     = 20
	YYMAXDEPTH = 500
	NTERM      = 10
	MAXALIGN   = 7
	BITS       = 5
	NVAR       = BITS * 32

	SIGNONE   = 0
	SIGDONE   = 1
	SIGINTERN = 2

	SIGNINTERN = 1729 * 325 * 1729

	Plan9 = 1 << iota
	Unix
	Windows

	EOF = -1
	IGN = -2
	ESC = 1 << 20
)
const (
	Numdec = 1 << iota
	Numlong
	Numuns
	Numvlong
	Numflt
)

func SIGN(n uint64) uint64 {
	return 1 << (n - 1)
}
func MASK(n uint64) uint64 {
	return SIGN(n) | (SIGN(n) - 1)
}

type Bits struct {
	B [BITS]uint32
}
type Node struct {
	Left    *Node
	Right   *Node
	Label   unsafe.Pointer
	Pc      int32
	Reg     int
	Xoffset int32
	Fconst  float64 /* fp constant */
	Vconst  int64   /* non fp const */
	Cstring []byte  /* character string */
	Rstring []rune  /* rune string */

	Sym     Sym
	Type    *Type
	Lineno  int32
	Op      byte
	Oldop   byte
	Xcast   byte
	Class   byte
	Etype   byte
	Complex byte
	Addable byte
	Scale   byte
	Garb    byte
}

type Sym struct {
	Link      *Sym
	Type      *Type
	Suetag    *Type
	Tenum     *Type
	Macro     []byte
	Varlineno int32
	Offset    int32
	Vconst    int64
	Fconst    float64
	Label     *Node
	Lexical   uint16
	Name      string
	Block     uint16
	Sueblock  uint16
	Class     byte
	Sym       byte
	Aused     byte
	Sig       byte
	Dataflag  byte
}

type Decl struct {
	Link      *Decl
	Sym       *Sym
	Type      *Type
	Varlineno int32
	Offset    int32
	Val       int16
	Block     uint16
	Class     byte
	Aused     byte
}

type Type struct {
	Sym    *Sym
	Tag    *Sym
	Funct  *Funct
	Link   *Type
	Down   *Type
	Width  int32
	Offset int32
	Lineno int32
	Shift  byte
	Nbits  byte
	Etype  byte
	Garb   byte
	Align  byte
}

//#define	NODECL	((void(*)(int, Type*, Sym*))0)

/* general purpose initialization */
type Init struct {
	Code  int
	Value uint32
	S     string
}

var fi = struct {
	p string
	c int
}{}

//io
type Io struct {
	Link *Io
	P    []byte
	B    [BUFSIZ]byte
	C    int16
	F    int16
}

//history?
type Hist struct {
	Link   *Hist
	Name   string
	Line   int32
	Offset int32
}

//terminal
type Term struct {
	Mult int64
	Node *Node
}

const (
	Axxx = iota
	Ael1
	Ael2
	Asu2
	Aarg0
	Aarg1
	Aarg2
	Aaut3
	NALIGN
)

const (
	DMARK = iota
	DAUTO
	DSUE
	DLABEL
)

//Operators
const (
	OXXX = iota
	OADD
	OADDR
	OAND
	OANDAND
	OARRAY
	OAS
	OASI
	OASADD
	OASAND
	OASASHL
	OASASHR
	OASDIV
	OASHL
	OASHR
	OASLDIV
	OASLMOD
	OASLMUL
	OASLSHR
	OASMOD
	OASMUL
	OASOR
	OASSUB
	OASXOR
	OBIT
	OBREAK
	OCASE
	OCAST
	OCOMMA
	OCOND
	OCONST
	OCONTINUE
	ODIV
	ODOT
	ODOTDOT
	ODWHILE
	OENUM
	OEQ
	OEXREG
	OFOR
	OFUNC
	OGE
	OGOTO
	OGT
	OHI
	OHS
	OIF
	OIND
	OINDREG
	OINIT
	OLABEL
	OLDIV
	OLE
	OLIST
	OLMOD
	OLMUL
	OLO
	OLS
	OLSHR
	OLT
	OMOD
	OMUL
	ONAME
	ONE
	ONOT
	OOR
	OOROR
	OPOSTDEC
	OPOSTINC
	OPREDEC
	OPREINC
	OPREFETCH
	OPROTO
	OREGISTER
	ORETURN
	OSET
	OSIGN
	OSIZE
	OSTRING
	OLSTRING
	OSTRUCT
	OSUB
	OSWITCH
	OUNION
	OUSED
	OWHILE
	OXOR
	ONEG
	OCOM
	OPOS
	OELEM

	OTST /* used in some compilers */
	OINDEX
	OFAS
	OREGPAIR
	OROTL

	OEND
)
const (
	TXXX = iota
	TCHAR
	TUCHAR
	TSHORT
	TUSHORT
	TINT
	TUINT
	TLONG
	TULONG
	TVLONG
	TUVLONG
	TFLOAT
	TDOUBLE
	TIND
	TFUNC
	TARRAY
	TVOID
	TSTRUCT
	TUNION
	TENUM
	NTYPE
)
const (
	TAUTO = NTYPE + iota
	TEXTERN
	TSTATIC
	TTYPEDEF
	TTYPESTR
	TREGISTER
	TCONSTNT
	TVOLATILE
	TUNSIGNED
	TSIGNED
	TDOT
	TFILE
	TOLD
	NALLTYPES
)
const (
	CXXX = iota
	CAUTO
	CEXTERN
	CGLOBL
	CSTATIC
	CLOCAL
	CTYPEDEF
	CTYPESTR
	CPARAM
	CSELEM
	CLABEL
	CEXREG
	NCTYPES
)
const (
	GXXX     = 0
	GCONSTNT = 1 << iota
	GVOLATILE
	NGTYPES

	GINCOMPLETE = 1 << 2
)
const (
	BCHAR     = 1 << TCHAR
	BUCHAR    = 1 << TUCHAR
	BSHORT    = 1 << TSHORT
	BUSHORT   = 1 << TUSHORT
	BINT      = 1 << TINT
	BUINT     = 1 << TUINT
	BLONG     = 1 << TLONG
	BULONG    = 1 << TULONG
	BVLONG    = 1 << TVLONG
	BUVLONG   = 1 << TUVLONG
	BFLOAT    = 1 << TFLOAT
	BDOUBLE   = 1 << TDOUBLE
	BIND      = 1 << TIND
	BFUNC     = 1 << TFUNC
	BARRAY    = 1 << TARRAY
	BVOID     = 1 << TVOID
	BSTRUCT   = 1 << TSTRUCT
	BUNION    = 1 << TUNION
	BENUM     = 1 << TENUM
	BFILE     = 1 << TFILE
	BDOT      = 1 << TDOT
	BCONSTNT  = 1 << TCONSTNT
	BVOLATILE = 1 << TVOLATILE
	BUNSIGNED = 1 << TUNSIGNED
	BSIGNED   = 1 << TSIGNED
	BAUTO     = 1 << TAUTO
	BEXTERN   = 1 << TEXTERN
	BSTATIC   = 1 << TSTATIC
	BTYPEDEF  = 1 << TTYPEDEF
	BTYPESTR  = 1 << TTYPESTR
	BREGISTER = 1 << TREGISTER

	BINTEGER = BCHAR | BUCHAR | BSHORT | BUSHORT | BINT | BUINT | BLONG | BULONG | BVLONG | BUVLONG
	BNUMBER  = BINTEGER | BFLOAT | BDOUBLE

	/* these can be overloaded with complex types */

	BCLASS = BAUTO | BEXTERN | BSTATIC | BTYPEDEF | BTYPESTR | BREGISTER
	BGARB  = BCONSTNT | BVOLATILE
)

type Funct struct {
	Sym    [OEND]*Sym
	Castto [NTYPE]*Sym
	Castfr [NTYPE]*Sym
}

type Dynimp struct {
	Local  []byte
	Remote []byte
	Path   []byte
}

type Dynexp struct {
	Local  []byte
	Remote []byte
}

var (
	dynimp  *Dynimp
	ndynimp int

	dynexp  *Dynexp
	ndynexp int

	enum = struct {
		tenum     *Type   /* type of entire enum */
		cenum     *Type   /* type of current enum run */
		lastenum  int64   /* value of current enum */
		floatenum float64 /* value of current enum */
	}{}

	//Array?
	dclstack *Decl
	debug    [256]int

	ehist *Hist

	firstbit     int32
	firstarg     *Sym
	firstargtype *Type
	firstdcl     *Decl
	fperror      int

	hash map[string]*Sym
	hunk string

	include []string

	lastclass byte

	lastdcl,
	lasttype *Type

	lastbit,
	nsymb,
	lastfield,
	lineno,
	nearln,
	nhunk,
	nerrors,
	newflag,
	ninclude int

	nodproto,
	nodcast *Node

	outfile,
	pathname string

	taggen,
	thechar,
	peekc int

	stkoff int32

	symstring *Sym

	strf,
	strl,
	tfield,
	thisfn,
	tufield *Type

	thestring string

	thunk int32

	fntypes  [NALLTYPES]*Type
	initlist *Node
	term     [NTERM]Term

	nterm,
	packflg,
	fproundflg,
	textflag,
	dataflag,
	flag_largemodel,
	ncontin,
	canreach,
	warnreach int

	zbits Bits

	tab [NTYPE][NTYPE]byte

	ncast, tadd, tand, targ, tasadd, tasign, tcast,
	tdot, tfunct, tindir, tmul,
	tnot, trel, tsub []int32

	typeaf, typefd, typei, typesu,
	typesuv, typeu, typev, typec,
	typeh, typeil, typeilp, typechl,
	typechlv, typechlvp, typechlp, typechlpf []byte

	typeword,
	typecmplx *byte

	typeHash1,
	typeHash2,
	typeHash3 uint32 = 0x2edab8c9, 0x1dc74fb8, 0x1f241331

	//This needs to be created in the backend(e.g. 5c, 6c, 8c, etc.) init func before the compiler is initialized because
	//it is platform dependent and is implemented via preprocessor macros in the C version.  
	Ewidth []int8
)
