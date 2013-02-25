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
	Name      []byte
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

	TAUTO = NTYPE
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
	BCHAR     int64 = 1 << TCHAR
	BUCHAR    int64 = 1 << TUCHAR
	BSHORT    int64 = 1 << TSHORT
	BUSHORT   int64 = 1 << TUSHORT
	BINT      int64 = 1 << TINT
	BUINT     int64 = 1 << TUINT
	BLONG     int64 = 1 << TLONG
	BULONG    int64 = 1 << TULONG
	BVLONG    int64 = 1 << TVLONG
	BUVLONG   int64 = 1 << TUVLONG
	BFLOAT    int64 = 1 << TFLOAT
	BDOUBLE   int64 = 1 << TDOUBLE
	BIND      int64 = 1 << TIND
	BFUNC     int64 = 1 << TFUNC
	BARRAY    int64 = 1 << TARRAY
	BVOID     int64 = 1 << TVOID
	BSTRUCT   int64 = 1 << TSTRUCT
	BUNION    int64 = 1 << TUNION
	BENUM     int64 = 1 << TENUM
	BFILE     int64 = 1 << TFILE
	BDOT      int64 = 1 << TDOT
	BCONSTNT  int64 = 1 << TCONSTNT
	BVOLATILE int64 = 1 << TVOLATILE
	BUNSIGNED int64 = 1 << TUNSIGNED
	BSIGNED   int64 = 1 << TSIGNED
	BAUTO     int64 = 1 << TAUTO
	BEXTERN   int64 = 1 << TEXTERN
	BSTATIC   int64 = 1 << TSTATIC
	BTYPEDEF  int64 = 1 << TTYPEDEF
	BTYPESTR  int64 = 1 << TTYPESTR
	BREGISTER int64 = 1 << TREGISTER

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

	en = struct {
		tenum     *Type   /* type of entire enum */
		cenum     *Type   /* type of current enum run */
		lastenum  int64   /* value of current enum */
		floatenum float64 /* value of current enum */
	}{}

	autobn    int
	autoffset int32
	blockno   int

	//Array?
	dclstack *Decl
	debug    [256]int

	ehist *Hist

	firstbit     int32
	firstarg     *Sym
	firstargtype *Type
	firstdcl     *Decl
	fperror      int

	hash [NHASH]*Sym
	hunk string

	include []string

	iofree,
	ionext,
	iostack *Io

	lastclass byte

	lastdcl,
	lasttype *Type

	lastbit,
	nsymb,
	lastfield,
	lineno,
	nearln,
	nhunk int32

	nerrors,
	newflag,
	ninclude int

	nodproto,
	nodcast *Node

	//Whats the go equiv of biobuf? 
	//bufio.Reader, Bytes.Buffer, slice???

	//Biobuf	outbuf;
	//Biobuf	diagbuf;

	outfile,
	pathname,
	symb string

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

	thunk    int32
	types    [NALLTYPES]*Type
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

	EWidth []int8
)
