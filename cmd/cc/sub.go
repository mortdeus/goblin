package cc

import (
	"errors"
	"fmt"
)

var (
	Tokens *tokens
)

type tokens struct {
	THash  [NALLTYPES]uint32
	BNames [NALIGN]string
	TNames [NALLTYPES]string
	GNames [NGTYPES]string
	QNames [NALLTYPES]string
	CNames [NCTYPES]string
	ONames [OEND + 1]string

	TypeI, TypeU, TypeC, TypeH, TypeV,
	TypeIL, TypeFD, TypeAF, TypeSU,
	TypeSUV, TypeILP, TypeCHL, TypeCHLP, TypeCHLPFD [NALLTYPES]byte

	//This has to match 
	TypeWord    [NALLTYPES]byte
	TypeComplex [NALLTYPES]byte

	//In kencc these are int32 arrays. 
	//As far as I can tell they dont have to be, and uint32 makes sense. 
	Asign, AsAdd, Cast, Add,
	Sub, Mul, And, Rel []uint32
}

func (t *tokens) Init() error {
	if t != nil {
		return errors.New("Tokens are already initialized.")
	}
	t = new(tokens)
	urk := func(name string, max, i int) (err error) {
		if i >= max {
			err = errors.New(fmt.Sprintf("Bad Tinit: %s %v >= %v", name, i, max))
		}
		return
	}

	for _, p := range tHashInit {
		if err := urk("thash", len(t.THash), p.Code); err != nil {
			return err
		}
		t.THash[p.Code] = p.Value
	}
	for _, p := range bNamesInit {
		if err := urk("bnames", len(t.BNames), p.Code); err != nil {
			return err
		}
		t.BNames[p.Code] = p.S
	}
	for _, p := range tNamesInit {
		if err := urk("tnames", len(t.TNames), p.Code); err != nil {
			return err
		}
		t.TNames[p.Code] = p.S
	}
	for _, p := range gNamesInit {
		if err := urk("gnames", len(t.GNames), p.Code); err != nil {
			return err
		}
		t.GNames[p.Code] = p.S
	}
	for _, p := range qNamesInit {
		if err := urk("qnames", len(t.QNames), p.Code); err != nil {
			return err
		}
		t.QNames[p.Code] = p.S
	}
	for _, p := range cNamesInit {
		if err := urk("cnames", len(t.CNames), p.Code); err != nil {
			return err
		}
		t.CNames[p.Code] = p.S
	}
	for _, p := range oNamesInit {
		if err := urk("onames", len(t.ONames), p.Code); err != nil {
			return err
		}
		t.ONames[p.Code] = p.S
	}
	for _, p := range typeiInit {
		if err := urk("typei", len(t.TypeI), p); err != nil {
			return err
		}
		t.TypeI[p] = 1
	}
	for _, p := range typeuInit {
		if err := urk("typeu", len(t.TypeU), p); err != nil {
			return err
		}
		t.TypeU[p] = 1
	}
	for _, p := range typesuvInit {
		if err := urk("typesuv", len(t.TypeSUV), p); err != nil {
			return err
		}
		t.TypeSUV[p] = 1
	}
	for _, p := range typeilpInit {
		if err := urk("typeilp", len(t.TypeILP), p); err != nil {
			return err
		}
		t.TypeILP[p] = 1
	}
	for _, p := range typechlInit {
		if err := urk("typechl", len(t.TypeCHL), p); err != nil {
			return err
		}
		t.TypeCHL[p] = 1
	}
	for _, p := range typechlpInit {
		if err := urk("typechlp", len(t.TypeCHLP), p); err != nil {
			return err
		}
		t.TypeCHLP[p] = 1
	}
	for _, p := range typechlpfdInit {
		if err := urk("typechlpfd", len(t.TypeCHLPFD), p); err != nil {
			return err
		}
		t.TypeCHLPFD[p] = 1
	}
	for _, p := range typecInit {
		if err := urk("typec", len(t.TypeC), p); err != nil {
			return err
		}
		t.TypeC[p] = 1
	}
	for _, p := range typehInit {
		if err := urk("typeh", len(t.TypeH), p); err != nil {
			return err
		}
		t.TypeH[p] = 1
	}
	for _, p := range typeilInit {
		if err := urk("typeil", len(t.TypeIL), p); err != nil {
			return err
		}
		t.TypeIL[p] = 1
	}
	for _, p := range typevInit {
		if err := urk("typev", len(t.TypeV), p); err != nil {
			return err
		}
		t.TypeV[p] = 1
	}
	for _, p := range typefdInit {
		if err := urk("typefd", len(t.TypeFD), p); err != nil {
			return err
		}
		t.TypeFD[p] = 1
	}
	for _, p := range typeafInit {
		if err := urk("typeaf", len(t.TypeAF), p); err != nil {
			return err
		}
		t.TypeAF[p] = 1
	}
	for _, p := range typesuInit {
		if err := urk("typesu", len(t.TypeSU), p); err != nil {
			return err
		}
		t.TypeSU[p] = 1
	}
	for _, p := range tasignInit {
		if err := urk("tasign", len(t.Asign), p.Code); err != nil {
			return err
		}
		t.Asign[p.Code] = p.Value
	}
	for _, p := range tasaddInit {
		if err := urk("tasadd", len(t.AsAdd), p.Code); err != nil {
			return err
		}
		t.AsAdd[p.Code] = p.Value
	}
	for _, p := range tcastInit {
		if err := urk("tcast", len(t.Cast), p.Code); err != nil {
			return err
		}
		t.Cast[p.Code] = p.Value
	}
	for _, p := range taddInit {
		if err := urk("tadd", len(t.Add), p.Code); err != nil {
			return err
		}
		t.Add[p.Code] = p.Value
	}
	for _, p := range tsubInit {
		if err := urk("tsub", len(t.Sub), p.Code); err != nil {
			return err
		}
		t.Sub[p.Code] = p.Value
	}
	for _, p := range tmulInit {
		if err := urk("tmul", len(t.Mul), p.Code); err != nil {
			return err
		}
		t.Mul[p.Code] = p.Value
	}
	for _, p := range tandInit {
		if err := urk("tand", len(t.And), p.Code); err != nil {
			return err
		}
		t.And[p.Code] = p.Value
	}
	for _, p := range trelInit {
		if err := urk("trel", len(t.Rel), p.Code); err != nil {
			return err
		}
		t.Rel[p.Code] = p.Value
	}
	t.TypeWord = t.TypeCHLP
	t.TypeComplex = t.TypeSUV
	return nil
}

func (Node) New(t int, l, r *Node) *Node {
	n := new(Node)
	n.Op = t
	n.Left = l
	n.Right = r
	if l && t != OGOTO {
		n.Lineno = l.Lineno
	} else if r != nil {
		n.Lineno = r.Lineno
	} else {
		n.Lineno = Compiler.Lineno
	}
	newflag = 1
	return n
}
func (Node) New1(o int, l, r *Node) *Node {
	n := Node.New(o, l, r)
	n.Lineno = nearln
	return n
}

var (
	tHashInit = [...]Init{
		{TXXX, 0x17527bbd, ""},
		{TCHAR, 0x5cedd32b, ""},
		{TUCHAR, 0x552c4454, ""},
		{TSHORT, 0x63040b4b, ""},
		{TUSHORT, 0x32a45878, ""},
		{TINT, 0x4151d5bd, ""},
		{TUINT, 0x5ae707d6, ""},
		{TLONG, 0x5ef20f47, ""},
		{TULONG, 0x36d8eb8f, ""},
		{TVLONG, 0x6e5e9590, ""},
		{TUVLONG, 0x75910105, ""},
		{TFLOAT, 0x25fd7af1, ""},
		{TDOUBLE, 0x7c40a1b2, ""},
		{TIND, 0x1b832357, ""},
		{TFUNC, 0x6babc9cb, ""},
		{TARRAY, 0x7c50986d, ""},
		{TVOID, 0x44112eff, ""},
		{TSTRUCT, 0x7c2da3bf, ""},
		{TUNION, 0x3eb25e98, ""},
		{TENUM, 0x44b54f61, ""},
		{TFILE, 0x19242ac3, ""},
		{TOLD, 0x22b15988, ""},
		{TDOT, 0x0204f6b3, ""},
	}
	bNamesInit = [...]Init{
		{Axxx, 0, "Axxx"},
		{Ael1, 0, "el1"},
		{Ael2, 0, "el2"},
		{Asu2, 0, "su2"},
		{Aarg0, 0, "arg0"},
		{Aarg1, 0, "arg1"},
		{Aarg2, 0, "arg2"},
		{Aaut3, 0, "aut3"},
	}
	tNamesInit = [...]Init{
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
	}
	gNamesInit = [...]Init{
		{GXXX, 0, "GXXX"},
		{GCONSTNT, 0, "CONST"},
		{GVOLATILE, 0, "VOLATILE"},
		{GVOLATILE | GCONSTNT, 0, "CONST-VOLATILE"},
	}
	qNamesInit = [...]Init{
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
		{TAUTO, 0, "AUTO"},
		{TEXTERN, 0, "EXTERN"},
		{TSTATIC, 0, "STATIC"},
		{TTYPEDEF, 0, "TYPEDEF"},
		{TTYPESTR, 0, "TYPESTR"},
		{TREGISTER, 0, "REGISTER"},
		{TCONSTNT, 0, "CONSTNT"},
		{TVOLATILE, 0, "VOLATILE"},
		{TUNSIGNED, 0, "UNSIGNED"},
		{TSIGNED, 0, "SIGNED"},
		{TDOT, 0, "DOT"},
		{TFILE, 0, "FILE"},
		{TOLD, 0, "OLD"},
	}
	cNamesInit = [...]Init{
		{CXXX, 0, "CXXX"},
		{CAUTO, 0, "AUTO"},
		{CEXTERN, 0, "EXTERN"},
		{CGLOBL, 0, "GLOBL"},
		{CSTATIC, 0, "STATIC"},
		{CLOCAL, 0, "LOCAL"},
		{CTYPEDEF, 0, "TYPEDEF"},
		{CTYPESTR, 0, "TYPESTR"},
		{CPARAM, 0, "PARAM"},
		{CSELEM, 0, "SELEM"},
		{CLABEL, 0, "LABEL"},
		{CEXREG, 0, "EXREG"},
	}
	oNamesInit = [...]Init{
		{OXXX, 0, "OXXX"},
		{OADD, 0, "ADD"},
		{OADDR, 0, "ADDR"},
		{OAND, 0, "AND"},
		{OANDAND, 0, "ANDAND"},
		{OARRAY, 0, "ARRAY"},
		{OAS, 0, "AS"},
		{OASI, 0, "ASI"},
		{OASADD, 0, "ASADD"},
		{OASAND, 0, "ASAND"},
		{OASASHL, 0, "ASASHL"},
		{OASASHR, 0, "ASASHR"},
		{OASDIV, 0, "ASDIV"},
		{OASHL, 0, "ASHL"},
		{OASHR, 0, "ASHR"},
		{OASLDIV, 0, "ASLDIV"},
		{OASLMOD, 0, "ASLMOD"},
		{OASLMUL, 0, "ASLMUL"},
		{OASLSHR, 0, "ASLSHR"},
		{OASMOD, 0, "ASMOD"},
		{OASMUL, 0, "ASMUL"},
		{OASOR, 0, "ASOR"},
		{OASSUB, 0, "ASSUB"},
		{OASXOR, 0, "ASXOR"},
		{OBIT, 0, "BIT"},
		{OBREAK, 0, "BREAK"},
		{OCASE, 0, "CASE"},
		{OCAST, 0, "CAST"},
		{OCOMMA, 0, "COMMA"},
		{OCOND, 0, "COND"},
		{OCONST, 0, "CONST"},
		{OCONTINUE, 0, "CONTINUE"},
		{ODIV, 0, "DIV"},
		{ODOT, 0, "DOT"},
		{ODOTDOT, 0, "DOTDOT"},
		{ODWHILE, 0, "DWHILE"},
		{OENUM, 0, "ENUM"},
		{OEQ, 0, "EQ"},
		{OEXREG, 0, "EXREG"},
		{OFOR, 0, "FOR"},
		{OFUNC, 0, "FUNC"},
		{OGE, 0, "GE"},
		{OGOTO, 0, "GOTO"},
		{OGT, 0, "GT"},
		{OHI, 0, "HI"},
		{OHS, 0, "HS"},
		{OIF, 0, "IF"},
		{OIND, 0, "IND"},
		{OINDREG, 0, "INDREG"},
		{OINIT, 0, "INIT"},
		{OLABEL, 0, "LABEL"},
		{OLDIV, 0, "LDIV"},
		{OLE, 0, "LE"},
		{OLIST, 0, "LIST"},
		{OLMOD, 0, "LMOD"},
		{OLMUL, 0, "LMUL"},
		{OLO, 0, "LO"},
		{OLS, 0, "LS"},
		{OLSHR, 0, "LSHR"},
		{OLT, 0, "LT"},
		{OMOD, 0, "MOD"},
		{OMUL, 0, "MUL"},
		{ONAME, 0, "NAME"},
		{ONE, 0, "NE"},
		{ONOT, 0, "NOT"},
		{OOR, 0, "OR"},
		{OOROR, 0, "OROR"},
		{OPOSTDEC, 0, "POSTDEC"},
		{OPOSTINC, 0, "POSTINC"},
		{OPREDEC, 0, "PREDEC"},
		{OPREINC, 0, "PREINC"},
		{OPREFETCH, 0, "PREFETCH"},
		{OPROTO, 0, "PROTO"},
		{OREGISTER, 0, "REGISTER"},
		{ORETURN, 0, "RETURN"},
		{OSET, 0, "SET"},
		{OSIGN, 0, "SIGN"},
		{OSIZE, 0, "SIZE"},
		{OSTRING, 0, "STRING"},
		{OLSTRING, 0, "LSTRING"},
		{OSTRUCT, 0, "STRUCT"},
		{OSUB, 0, "SUB"},
		{OSWITCH, 0, "SWITCH"},
		{OUNION, 0, "UNION"},
		{OUSED, 0, "USED"},
		{OWHILE, 0, "WHILE"},
		{OXOR, 0, "XOR"},
		{OPOS, 0, "POS"},
		{ONEG, 0, "NEG"},
		{OCOM, 0, "COM"},
		{OELEM, 0, "ELEM"},
		{OTST, 0, "TST"},
		{OINDEX, 0, "INDEX"},
		{OFAS, 0, "FAS"},
		{OREGPAIR, 0, "REGPAIR"},
		{OROTL, 0, "ROTL"},
		{OEND, 0, "END"},
	}

	comrel = [...]byte{ONE, OEQ, OGT, OHI, OGE, OHS, OLT, OLO, OLE, OLS}
	invrel = [...]byte{ONE, OEQ, OGT, OHI, OGE, OHS, OLT, OLO, OLE, OLS}
	logrel = [...]byte{OEQ, ONE, OLS, OLS, OLO, OLO, OHS, OHS, OHI, OHI}

	typeiInit      = [...]int{TCHAR, TUCHAR, TSHORT, TUSHORT, TINT, TUINT, TLONG, TULONG, TVLONG, TUVLONG}
	typeuInit      = [...]int{TUCHAR, TUSHORT, TUINT, TULONG, TUVLONG, TIND}
	typesuvInit    = [...]int{TVLONG, TUVLONG, TSTRUCT, TUNION}
	typeilpInit    = [...]int{TINT, TUINT, TLONG, TULONG, TIND}
	typechlInit    = [...]int{TCHAR, TUCHAR, TSHORT, TUSHORT, TINT, TUINT, TLONG, TULONG}
	typechlpInit   = [...]int{TCHAR, TUCHAR, TSHORT, TUSHORT, TINT, TUINT, TLONG, TULONG, TIND}
	typechlpfdInit = [...]int{TCHAR, TUCHAR, TSHORT, TUSHORT, TINT, TUINT, TLONG, TULONG, TFLOAT, TDOUBLE, TIND}
	typecInit      = [...]int{TCHAR, TUCHAR}
	typehInit      = [...]int{TSHORT, TUSHORT}
	typeilInit     = [...]int{TINT, TUINT, TLONG, TULONG}
	typevInit      = [...]int{TVLONG, TUVLONG}
	typefdInit     = [...]int{TFLOAT, TDOUBLE}
	typeafInit     = [...]int{TFUNC, TARRAY}
	typesuInit     = [...]int{TSTRUCT, TUNION}

	tasignInit = [...]Init{
		{TCHAR, BNUMBER, ""},
		{TUCHAR, BNUMBER, ""},
		{TSHORT, BNUMBER, ""},
		{TUSHORT, BNUMBER, ""},
		{TINT, BNUMBER, ""},
		{TUINT, BNUMBER, ""},
		{TLONG, BNUMBER, ""},
		{TULONG, BNUMBER, ""},
		{TVLONG, BNUMBER, ""},
		{TUVLONG, BNUMBER, ""},
		{TFLOAT, BNUMBER, ""},
		{TDOUBLE, BNUMBER, ""},
		{TIND, BIND, ""},
		{TSTRUCT, BSTRUCT, ""},
		{TUNION, BUNION, ""},
	}
	tasaddInit = [...]Init{
		{TCHAR, BNUMBER, ""},
		{TUCHAR, BNUMBER, ""},
		{TSHORT, BNUMBER, ""},
		{TUSHORT, BNUMBER, ""},
		{TINT, BNUMBER, ""},
		{TUINT, BNUMBER, ""},
		{TLONG, BNUMBER, ""},
		{TULONG, BNUMBER, ""},
		{TVLONG, BNUMBER, ""},
		{TUVLONG, BNUMBER, ""},
		{TFLOAT, BNUMBER, ""},
		{TDOUBLE, BNUMBER, ""},
		{TIND, BINTEGER, ""},
	}
	tcastInit = [...]Init{
		{TCHAR, BNUMBER | BIND | BVOID, ""},
		{TUCHAR, BNUMBER | BIND | BVOID, ""},
		{TSHORT, BNUMBER | BIND | BVOID, ""},
		{TUSHORT, BNUMBER | BIND | BVOID, ""},
		{TINT, BNUMBER | BIND | BVOID, ""},
		{TUINT, BNUMBER | BIND | BVOID, ""},
		{TLONG, BNUMBER | BIND | BVOID, ""},
		{TULONG, BNUMBER | BIND | BVOID, ""},
		{TVLONG, BNUMBER | BIND | BVOID, ""},
		{TUVLONG, BNUMBER | BIND | BVOID, ""},
		{TFLOAT, BNUMBER | BVOID, ""},
		{TDOUBLE, BNUMBER | BVOID, ""},
		{TIND, BINTEGER | BIND | BVOID, ""},
		{TVOID, BVOID, ""},
		{TSTRUCT, BSTRUCT | BVOID, ""},
		{TUNION, BUNION | BVOID, ""},
	}
	taddInit = [...]Init{
		{TCHAR, BNUMBER | BIND, ""},
		{TUCHAR, BNUMBER | BIND, ""},
		{TSHORT, BNUMBER | BIND, ""},
		{TUSHORT, BNUMBER | BIND, ""},
		{TINT, BNUMBER | BIND, ""},
		{TUINT, BNUMBER | BIND, ""},
		{TLONG, BNUMBER | BIND, ""},
		{TULONG, BNUMBER | BIND, ""},
		{TVLONG, BNUMBER | BIND, ""},
		{TUVLONG, BNUMBER | BIND, ""},
		{TFLOAT, BNUMBER, ""},
		{TDOUBLE, BNUMBER, ""},
		{TIND, BINTEGER, ""},
	}
	tsubInit = [...]Init{
		{TCHAR, BNUMBER, ""},
		{TUCHAR, BNUMBER, ""},
		{TSHORT, BNUMBER, ""},
		{TUSHORT, BNUMBER, ""},
		{TINT, BNUMBER, ""},
		{TUINT, BNUMBER, ""},
		{TLONG, BNUMBER, ""},
		{TULONG, BNUMBER, ""},
		{TVLONG, BNUMBER, ""},
		{TUVLONG, BNUMBER, ""},
		{TFLOAT, BNUMBER, ""},
		{TDOUBLE, BNUMBER, ""},
		{TIND, BINTEGER | BIND, ""},
	}
	tmulInit = [...]Init{
		{TCHAR, BNUMBER, ""},
		{TUCHAR, BNUMBER, ""},
		{TSHORT, BNUMBER, ""},
		{TUSHORT, BNUMBER, ""},
		{TINT, BNUMBER, ""},
		{TUINT, BNUMBER, ""},
		{TLONG, BNUMBER, ""},
		{TULONG, BNUMBER, ""},
		{TVLONG, BNUMBER, ""},
		{TUVLONG, BNUMBER, ""},
		{TFLOAT, BNUMBER, ""},
		{TDOUBLE, BNUMBER, ""},
	}
	tandInit = [...]Init{
		{TCHAR, BINTEGER, ""},
		{TUCHAR, BINTEGER, ""},
		{TSHORT, BINTEGER, ""},
		{TUSHORT, BINTEGER, ""},
		{TINT, BNUMBER, ""},
		{TUINT, BNUMBER, ""},
		{TLONG, BINTEGER, ""},
		{TULONG, BINTEGER, ""},
		{TVLONG, BINTEGER, ""},
		{TUVLONG, BINTEGER, ""},
	}
	trelInit = [...]Init{
		{TCHAR, BNUMBER, ""},
		{TUCHAR, BNUMBER, ""},
		{TSHORT, BNUMBER, ""},
		{TUSHORT, BNUMBER, ""},
		{TINT, BNUMBER, ""},
		{TUINT, BNUMBER, ""},
		{TLONG, BNUMBER, ""},
		{TULONG, BNUMBER, ""},
		{TVLONG, BNUMBER, ""},
		{TUVLONG, BNUMBER, ""},
		{TFLOAT, BNUMBER, ""},
		{TDOUBLE, BNUMBER, ""},
		{TIND, BIND, ""},
	}
)

func typ(et byte, d *Type) *Type {

	t := new(Type)
	t.Etype = et
	t.Link = d
	t.Down = new(Type)
	t.Sym = new(Sym)

	if et < NTYPE {
		t.Width = int32(Ewidth[et])
	} else {
		// for TDOT or TOLD in prototype
		t.Width = -1
	}

	return t
}
