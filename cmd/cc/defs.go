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

var fi = struct{
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

	dynexp  *Dynexp/
	ndynexp int

	en = struct {
		tenum     *Type   /* type of entire enum */
		cenum     *Type   /* type of current enum run */
		lastenum  int64   /* value of current enum */
		floatenum float64 /* value of current enum */
	}{}

	autobn       int
	autoffset    int32
	blockno      int
	dclstack     *Decl
	debug        [256]int
	ehist        *Hist
	firstbit     int32
	firstarg     *Sym
	firstargtype *Type
	firstdcl     *Decl
	fperror      int
	hash         [NHASH]*Sym
	hunk         string
	include      []string
	iofree       *Io
	ionext       *Io
	iostack      *Io
	lastbit      int32
	lastclass    byte
	lastdcl      *Type
	lastfield    int32
	lasttype     *Type
	lineno       int32
	nearln       int32
	nerrors      int
	newflag      int
	nhunk        int32
	ninclude     int
	nodproto     *Node
	nodcast      *Node
	nsymb        int32
	//Biobuf	outbuf;
	//Biobuf	diagbuf;
	outfile         string
	pathname        string
	peekc           int
	stkoff          int32
	strf            *Type
	strl            *Type
	symb            string
	symstring       *Sym
	taggen          int
	tfield          *Type
	tufield         *Type
	thechar         int
	thestring       string
	thisfn          *Type
	thunk           int32
	types           [NALLTYPES]*Type
	fntypes         [NALLTYPES]*Type
	initlist        *Node
	term            [NTERM]Term
	nterm           int
	packflg         int
	fproundflg      int
	textflag        int
	dataflag        int
	flag_largemodel int
	ncontin         int
	canreach        int
	warnreach       int
	zbits           Bits

	onames,
	tnames,
	gnames,
	cnames,
	qnames,
	bnames []string
	tab                    [NTYPE][NTYPE]byte
	comrel, invrel, logrel []byte
	ncast, tadd, tand,
	targ, tasadd, tasign, tcast,
	tdot, tfunct, tindir, tmul,
	tnot, trel, tsub []int32

	typeaf    []byte
	typefd    []byte
	typei     []byte
	typesu    []byte
	typesuv   []byte
	typeu     []byte
	typev     []byte
	typec     []byte
	typeh     []byte
	typeil    []byte
	typeilp   []byte
	typechl   []byte
	typechlv  []byte
	typechlvp []byte
	typechlp  []byte
	typechlpf []byte

	typeword  *byte
	typecmplx *byte

	TypeHash1    uint32 = 0x2edab8c9
	TypeHash2    uint32 = 0x1dc74fb8
	TypeHash3    uint32 = 0x1f241331
	TypeHash     [NALLTYPES]uint32
	TypeHashInit = [...]Init{
		{TXXX, 0x17527bbd, 0},
		{TCHAR, 0x5cedd32b, 0},
		{TUCHAR, 0x552c4454, 0},
		{TSHORT, 0x63040b4b, 0},
		{TUSHORT, 0x32a45878, 0},
		{TINT, 0x4151d5bd, 0},
		{TUINT, 0x5ae707d6, 0},
		{TLONG, 0x5ef20f47, 0},
		{TULONG, 0x36d8eb8f, 0},
		{TVLONG, 0x6e5e9590, 0},
		{TUVLONG, 0x75910105, 0},
		{TFLOAT, 0x25fd7af1, 0},
		{TDOUBLE, 0x7c40a1b2, 0},
		{TIND, 0x1b832357, 0},
		{TFUNC, 0x6babc9cb, 0},
		{TARRAY, 0x7c50986d, 0},
		{TVOID, 0x44112eff, 0},
		{TSTRUCT, 0x7c2da3bf, 0},
		{TUNION, 0x3eb25e98, 0},
		{TENUM, 0x44b54f61, 0},
		{TFILE, 0x19242ac3, 0},
		{TOLD, 0x22b15988, 0},
		{TDOT, 0x0204f6b3, 0},
		{-1, 0, 0},
	}

	TypeNames [NALLTYPES]string

	TypeNamesInit = [...]Init{
		{TXXX, 0, "TXXX"},
		{TCHAR, 0, "CHAR"},
		{TUCHAR, 0, "UCHAR"},
		{TSHORT, 0, "SHORT"},
		{TUSHORT, 0, "USHORT"},
		{TINT, 0, "INT"},
		{TUINT, 0, "UINT"},
		{TLONG, 0, "LONG"},
		{TULONG, 0, "ULONG"},
		{TVLONG, 0, "VLONG"},
		{TUVLONG, 0, "UVLONG"},
		{TFLOAT, 0, "FLOAT"},
		{TDOUBLE, 0, "DOUBLE"},
		{TIND, 0, "IND"},
		{TFUNC, 0, "FUNC"},
		{TARRAY, 0, "ARRAY"},
		{TVOID, 0, "VOID"},
		{TSTRUCT, 0, "STRUCT"},
		{TUNION, 0, "UNION"},
		{TENUM, 0, "ENUM"},
		{TFILE, 0, "FILE"},
		{TOLD, 0, "OLD"},
		{TDOT, 0, "DOT"},
		{-1, 0, 0},
	}

// 	schar	ewidth[];
)
