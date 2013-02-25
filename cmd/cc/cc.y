%{
package cc
%}
%union{
	node *Node
	sym *Sym
	type_ *Type
	tycl struct{
		t *Type
		c byte
	}
	tyty struct{
		t1,
		t2,
		t3 *Type
		c byte
	}
	sval struct{
		s string
		l int
	}
	lval int32
	dval double
	vval int64
}

%type	<sym>	ltag
%type	<lval>	gctname gcname cname gname tname
%type	<lval>	gctnlist gcnlist zgnlist
%type	<Type>	tlist sbody complex
%type	<tycl>	types
%type	<node>	zarglist arglist zcexpr
%type	<node>	name block stmnt cexpr expr xuexpr pexpr
%type	<node>	zelist elist adecl slist uexpr string lstring
%type	<node>	xdecor xdecor2 labels label ulstmnt
%type	<node>	adlist edecor tag qual qlist
%type	<node>	abdecor abdecor1 abdecor2 abdecor3
%type	<node>	zexpr lexpr init ilist forexpr

%left	';'
%left	','
%right	'=' LPE LME LMLE LDVE LMDE LRSHE LLSHE LANDE LXORE LORE
%right	'?' ':'
%left	LOROR
%left	LANDAND
%left	'|'
%left	'^'
%left	'&'
%left	LEQ LNE
%left	'<' '>' LLE LGE
%left	LLSH LRSH
%left	'+' '-'
%left	'*' '/' '%'
%right	LMM LPP LMG '.' '[' '('

%token	<sym>	LNAME LTYPE
%token	<dval>	LFCONST LDCONST
%token	<vval>	LCONST LLCONST LUCONST LULCONST LVLCONST LUVLCONST
%token	<sval>	LSTRING LLSTRING

%token	LAUTO LBREAK LCASE LCHAR LCONTINUE LDEFAULT LDO
%token	LDOUBLE LELSE LEXTERN LFLOAT LFOR LGOTO
%token	LIF LINT LLONG LPREFETCH LREGISTER LRETURN LSHORT LSIZEOF LUSED
%token	LSTATIC LSTRUCT LSWITCH LTYPEDEF LTYPESTR LUNION LUNSIGNED
%token	LWHILE LVOID LENUM LSIGNED LCONSTNT LVOLATILE LSET LSIGNOF
%token	LRESTRICT LINLINE
%%

prog:
|	prog xdecl

/*
 * external declarator
 */
xdecl:
	zctlist ';'
	{
		dodecl(xdecl, lastclass, lasttype, new(Node))
	}
|	zctlist xdlist ';'
|	zctlist xdecor
	{
		lastdcl = new(Type)
		firstarg = new(Sym)
		dodecl(xdecl, lastclass, lasttype, $2)
		if lastdcl == nil || lastdcl.Etype != TFUNC {
			diag($2, "not a function")
			lastdcl = Types[TFUNC];
		}
		thisfn = lastdcl;
		markdcl()
		firstdcl = dclstack;
		argmark($2, 0)
	}
	pdecl
	{
		argmark($2, 1)
	}
	block
	{
		n := revertdcl()
		if n{
			$6 = Node.new(OLIST, n, $6)
		}
		if !debug['a'] && !debug['Z']{
			codgen($6, $2)
		}
	}

xdlist:
	xdecor
	{
		dodecl(xdecl, lastclass, lasttype, $1)
	}
|	xdecor
	{
		$1 = dodecl(xdecl, lastclass, lasttype, $1)
	}
	'=' init
	{
		doinit($1.sym, $1.type_, 0, $4)
	}
|	xdlist ',' xdlist

xdecor:
	xdecor2
|	'*' zgnlist xdecor
	{
		$$ = Node.new(OIND, $3, new(Node))
		$$.garb = simpleg($2)
	}

xdecor2:
	tag
|	'(' xdecor ')'
	{
		$$ = $2;
	}
|	xdecor2 '(' zarglist ')'
	{
		$$ = Node.new(OFUNC, $1, $3)
	}
|	xdecor2 '[' zexpr ']'
	{
		$$ = Node.new(OARRAY, $1, $3)
	}

/*
 * automatic declarator
 */
adecl:
	ctlist ';'
	{
		$$ = dodecl(adecl, lastclass, lasttype, new(Node))
	}
|	ctlist adlist ';'
	{
		$$ = $2;
	}

adlist:
	xdecor
	{
		dodecl(adecl, lastclass, lasttype, $1)
		$$ = new(Node)
	}
|	xdecor
	{
		$1 = dodecl(adecl, lastclass, lasttype, $1)
	}
	'=' init
	{
		w := $1.sym.type_.width
		$$ = doinit($1.sym, $1.type, 0, $4)
		$$ = contig($1.sym, $$, w)
	}
|	adlist ',' adlist
	{
		$$ = $1;
		if $3 != nil{ 
			$$ = $3;
		
		if $1 != nil{
			$$ = Node.new(OLIST, $1, $3)
			}
		}
	}

/*
 * parameter declarator
 */
pdecl:
|	pdecl ctlist pdlist ';'

pdlist:
	xdecor
	{
		dodecl(pdecl, lastclass, lasttype, $1)
	}
|	pdlist ',' pdlist

/*
 * structure element declarator
 */
edecl:
	tlist
	{
		lasttype = $1;
	}
	zedlist ';'
|	edecl tlist
	{
		lasttype = $2;
	}
	zedlist ';'

zedlist:					/* extension */
	{
		lastfield = 0;
		edecl(CXXX, lasttype, S)
	}
|	edlist

edlist:
	edecor
	{
		dodecl(edecl, CXXX, lasttype, $1)
	}
|	edlist ',' edlist

edecor:
	xdecor
	{
		lastbit = 0;
		firstbit = 1;
	}
|	tag ':' lexpr
	{
		$$ = Node.new(OBIT, $1, $3)
	}
|	':' lexpr
	{
		$$ = Node.new(OBIT, new(Node), $2)
	}

/*
 * abstract declarator
 */
abdecor:
	{
		$$ = new(Node)
	}
|	abdecor1

abdecor1:
	'*' zgnlist
	{
		$$ = Node.new(OIND, new(Node), new(Node))
		$$.garb = simpleg($2)
	}
|	'*' zgnlist abdecor1
	{
		$$ = Node.new(OIND, $3, new(Node))
		$$.garb = simpleg($2)
	}
|	abdecor2

abdecor2:
	abdecor3
|	abdecor2 '(' zarglist ')'
	{
		$$ = Node.new(OFUNC, $1, $3)
	}
|	abdecor2 '[' zexpr ']'
	{
		$$ = Node.new(OARRAY, $1, $3)
	}

abdecor3:
	'(' ')'
	{
		$$ = Node.new(OFUNC, new(Node), new(Node))
	}
|	'[' zexpr ']'
	{
		$$ = Node.new(OARRAY, new(Node), $2)
	}
|	'(' abdecor1 ')'
	{
		$$ = $2;
	}

init:
	expr
|	'{' ilist '}'
	{
		$$ = Node.new(OINIT, invert($2), new(Node))
	}

qual:
	'[' lexpr ']'
	{
		$$ = Node.new(OARRAY, $2, new(Node))
	}
|	'.' ltag
	{
		$$ = Node.new(OELEM, new(Node), new(Node))
		$$.sym = $2;
	}
|	qual '='

qlist:
	init ','
|	qlist init ','
	{
		$$ = Node.new(OLIST, $1, $2)
	}
|	qual
|	qlist qual
	{
		$$ = Node.new(OLIST, $1, $2)
	}

ilist:
	qlist
|	init
|	qlist init
	{
		$$ = Node.new(OLIST, $1, $2)
	}

zarglist:
	{
		$$ = new(Node);
	}
|	arglist
	{
		$$ = invert($1)
	}


arglist:
	name
|	tlist abdecor
	{
		$$ = Node.new(OPROTO, $2, new(Node))
		$$.type = $1;
	}
|	tlist xdecor
	{
		$$ = Node.new(OPROTO, $2, new(Node))
		$$.type = $1;
	}
|	'.' '.' '.'
	{
		$$ = Node.new(ODOTDOT, new(Node), new(Node))
	}
|	arglist ',' arglist
	{
		$$ = Node.new(OLIST, $1, $3)
	}

block:
	'{' slist '}'
	{
		$$ = invert($2)
	//	if $2 != nil
	//		$$ = Node.new(OLIST, $2, $$)
		if $$ == nil{
			$$ = Node.new(OLIST, new(Node), new(Node))
		}
	}

slist:
	{
		$$ = new(Node)
	}
|	slist adecl
	{
		$$ = Node.new(OLIST, $1, $2)
	}
|	slist stmnt
	{
		$$ = Node.new(OLIST, $1, $2)
	}

labels:
	label
|	labels label
	{
		$$ = Node.new(OLIST, $1, $2)
	}

label:
	LCASE expr ':'
	{
		$$ = Node.new(OCASE, $2, new(Node))
	}
|	LDEFAULT ':'
	{
		$$ = Node.new(OCASE, new(Node), new(Node))
	}
|	LNAME ':'
	{
		$$ = Node.new(OLABEL, dcllabel($1, 1), new(Node))
	}

stmnt:
	error ';'
	{
		$$ = new(Node)
	}
|	ulstmnt
|	labels ulstmnt
	{
		$$ = Node.new(OLIST, $1, $2)
	}

forexpr:
	zcexpr
|	ctlist adlist
	{
		$$ = $2;
	}

ulstmnt:
	zcexpr ';'
|	{
		markdcl()
	}
	block
	{
		$$ = revertdcl()
		if $$
			$$ = Node.new(OLIST, $$, $2)
		else
			$$ = $2;
	}
|	LIF '(' cexpr ')' stmnt
	{
		$$ = Node.new(OIF, $3, new(OLIST, $5, new(Node)))
		if $5 == nil{
			warn($3, "empty if body")
		}
	}
|	LIF '(' cexpr ')' stmnt LELSE stmnt
	{
		$$ = Node.new(OIF, $3, new(OLIST, $5, $7))
		if $5 == nil{
			warn($3, "empty if body")
		}
		if $7 == nil{
			warn($3, "empty else body")
		}
	}
|	{ markdcl() } LFOR '(' forexpr ';' zcexpr ';' zcexpr ')' stmnt
	{
		$$ = revertdcl()
		if $${
			if $4 {
				$4 = Node.new(OLIST, $$, $4)
			}else{
				$4 = $$
			}
		}
		$$ = Node.new(OFOR, Node.new(OLIST, $6, Node.new(OLIST, $4, $8)), $10)
	}
|	LWHILE '(' cexpr ')' stmnt
	{
		$$ = Node.new(OWHILE, $3, $5)
	}
|	LDO stmnt LWHILE '(' cexpr ')' ';'
	{
		$$ = Node.new(ODWHILE, $5, $2)
	}
|	LRETURN zcexpr ';'
	{
		$$ = Node.new(ORETURN, $2, new(Node))
		$$.type = thisfn.link;
	}
|	LSWITCH '(' cexpr ')' stmnt
	{
		$$ = Node.new(OCONST, new(Node), new(Node))
		$$.vconst = 0;
		$$.type = types[TINT];
		$3 = Node.new(OSUB, $$, $3)

		$$ = Node.new(OCONST, new(Node), new(Node))
		$$.vconst = 0;
		$$.type = types[TINT];
		$3 = Node.new(OSUB, $$, $3)

		$$ = Node.new(OSWITCH, $3, $5)
	}
|	LBREAK ';'
	{
		$$ = Node.new(OBREAK, new(Node), new(Node))
	}
|	LCONTINUE ';'
	{
		$$ = Node.new(OCONTINUE, new(Node), new(Node))
	}
|	LGOTO ltag ';'
	{
		$$ = Node.new(OGOTO, dcllabel($2, 0), new(Node))
	}
|	LUSED '(' zelist ')' ';'
	{
		$$ = Node.new(OUSED, $3, new(Node))
	}
|	LPREFETCH '(' zelist ')' ';'
	{
		$$ = Node.new(OPREFETCH, $3, new(Node))
	}
|	LSET '(' zelist ')' ';'
	{
		$$ = Node.new(OSET, $3, new(Node))
	}

zcexpr:
	{
		$$ = new(Node)
	}
|	cexpr

zexpr:
	{
		$$ = new(Node)
	}
|	lexpr

lexpr:
	expr
	{
		$$ = Node.new(OCAST, $1, new(Node))
		$$.type = types[TLONG];
	}

cexpr:
	expr
|	cexpr ',' cexpr
	{
		$$ = Node.new(OCOMMA, $1, $3)
	}

expr:
	xuexpr
|	expr '*' expr
	{
		$$ = Node.new(OMUL, $1, $3)
	}
|	expr '/' expr
	{
		$$ = Node.new(ODIV, $1, $3)
	}
|	expr '%' expr
	{
		$$ = Node.new(OMOD, $1, $3)
	}
|	expr '+' expr
	{
		$$ = Node.new(OADD, $1, $3)
	}
|	expr '-' expr
	{
		$$ = Node.new(OSUB, $1, $3)
	}
|	expr LRSH expr
	{
		$$ = Node.new(OASHR, $1, $3)
	}
|	expr LLSH expr
	{
		$$ = Node.new(OASHL, $1, $3)
	}
|	expr '<' expr
	{
		$$ = Node.new(OLT, $1, $3)
	}
|	expr '>' expr
	{
		$$ = Node.new(OGT, $1, $3)
	}
|	expr LLE expr
	{
		$$ = Node.new(OLE, $1, $3)
	}
|	expr LGE expr
	{
		$$ = Node.new(OGE, $1, $3)
	}
|	expr LEQ expr
	{
		$$ = Node.new(OEQ, $1, $3)
	}
|	expr LNE expr
	{
		$$ = Node.new(ONE, $1, $3)
	}
|	expr '&' expr
	{
		$$ = Node.new(OAND, $1, $3)
	}
|	expr '^' expr
	{
		$$ = Node.new(OXOR, $1, $3)
	}
|	expr '|' expr
	{
		$$ = Node.new(OOR, $1, $3)
	}
|	expr LANDAND expr
	{
		$$ = Node.new(OANDAND, $1, $3)
	}
|	expr LOROR expr
	{
		$$ = Node.new(OOROR, $1, $3)
	}
|	expr '?' cexpr ':' expr
	{
		$$ = Node.new(OCOND, $1, Node.new(OLIST, $3, $5))
	}
|	expr '=' expr
	{
		$$ = Node.new(OAS, $1, $3)
	}
|	expr LPE expr
	{
		$$ = Node.new(OASADD, $1, $3)
	}
|	expr LME expr
	{
		$$ = Node.new(OASSUB, $1, $3)
	}
|	expr LMLE expr
	{
		$$ = Node.new(OASMUL, $1, $3)
	}
|	expr LDVE expr
	{
		$$ = Node.new(OASDIV, $1, $3)
	}
|	expr LMDE expr
	{
		$$ = Node.new(OASMOD, $1, $3)
	}
|	expr LLSHE expr
	{
		$$ = Node.new(OASASHL, $1, $3)
	}
|	expr LRSHE expr
	{
		$$ = Node.new(OASASHR, $1, $3)
	}
|	expr LANDE expr
	{
		$$ = Node.new(OASAND, $1, $3)
	}
|	expr LXORE expr
	{
		$$ = Node.new(OASXOR, $1, $3)
	}
|	expr LORE expr
	{
		$$ = Node.new(OASOR, $1, $3)
	}

xuexpr:
	uexpr
|	'(' tlist abdecor ')' xuexpr
	{
		$$ = Node.new(OCAST, $5, new(Node))
		dodecl(NODECL, CXXX, $2, $3)
		$$.type = lastdcl;
		$$.xcast = 1;
	}
|	'(' tlist abdecor ')' '{' ilist '}'	/* extension */
	{
		$$ = Node.new(OSTRUCT, $6, new(Node))
		dodecl(NODECL, CXXX, $2, $3)
		$$.type = lastdcl;
	}

uexpr:
	pexpr
|	'*' xuexpr
	{
		$$ = Node.new(OIND, $2, new(Node))
	}
|	'&' xuexpr
	{
		$$ = Node.new(OADDR, $2, new(Node))
	}
|	'+' xuexpr
	{
		$$ = Node.new(OPOS, $2, new(Node))
	}
|	'-' xuexpr
	{
		$$ = Node.new(ONEG, $2, new(Node))
	}
|	'!' xuexpr
	{
		$$ = Node.new(ONOT, $2, new(Node))
	}
|	'~' xuexpr
	{
		$$ = Node.new(OCOM, $2, new(Node))
	}
|	LPP xuexpr
	{
		$$ = Node.new(OPREINC, $2, new(Node))
	}
|	LMM xuexpr
	{
		$$ = Node.new(OPREDEC, $2, new(Node))
	}
|	LSIZEOF uexpr
	{
		$$ = Node.new(OSIZE, $2, new(Node))
	}
|	LSIGNOF uexpr
	{
		$$ = Node.new(OSIGN, $2, new(Node))
	}

pexpr:
	'(' cexpr ')'
	{
		$$ = $2;
	}
|	LSIZEOF '(' tlist abdecor ')'
	{
		$$ = Node.new(OSIZE, new(Node), new(Node))
		dodecl(NODECL, CXXX, $3, $4)
		$$.type = lastdcl;
	}
|	LSIGNOF '(' tlist abdecor ')'
	{
		$$ = Node.new(OSIGN, new(Node), new(Node))
		dodecl(NODECL, CXXX, $3, $4)
		$$.type = lastdcl;
	}
|	pexpr '(' zelist ')'
	{
		$$ = Node.new(OFUNC, $1, new(Node))
		if $1.op == ONAME{
			if $1.type == nil{
				dodecl(xdecl, CXXX, types[TINT], $$)
			}
		}
		$$.right = invert($3)
	}
|	pexpr '[' cexpr ']'
	{
		$$ = Node.new(OIND, Node.new(OADD, $1, $3), new(Node))
	}
|	pexpr LMG ltag
	{
		$$ = Node.new(ODOT, Node.new(OIND, $1, new(Node)), new(Node))
		$$.sym = $3;
	}
|	pexpr '.' ltag
	{
		$$ = Node.new(ODOT, $1, new(Node))
		$$.sym = $3;
	}
|	pexpr LPP
	{
		$$ = Node.new(OPOSTINC, $1, new(Node))
	}
|	pexpr LMM
	{
		$$ = Node.new(OPOSTDEC, $1, new(Node))
	}
|	name
|	LCONST
	{
		$$ = Node.new(OCONST, new(Node), new(Node))
		$$.type = types[TINT]
		$$.vconst = $1
		$$.cstring = symb
	}
|	LLCONST
	{
		$$ = Node.new(OCONST, new(Node), new(Node))
		$$.type = types[TLONG]
		$$.vconst = $1
		$$.cstring = symb
	}
|	LUCONST
	{
		$$ = Node.new(OCONST, new(Node), new(Node))
		$$.type = types[TUINT]
		$$.vconst = $1
		$$.cstring = symb
	}
|	LULCONST
	{
		$$ = Node.new(OCONST, new(Node), new(Node))
		$$.type = types[TULONG]
		$$.vconst = $1
		$$.cstring = symb
	}
|	LDCONST
	{
		$$ = Node.new(OCONST, new(Node), new(Node))
		$$.type = types[TDOUBLE]
		$$.fconst = $1
		$$.cstring = symb
	}
|	LFCONST
	{
		$$ = Node.new(OCONST, new(Node), new(Node))
		$$.type = types[TFLOAT]
		$$.fconst = $1
		$$.cstring = symb
	}
|	LVLCONST
	{
		$$ = Node.new(OCONST, new(Node), new(Node))
		$$.type = types[TVLONG]
		$$.vconst = $1
		$$.cstring = symb
	}
|	LUVLCONST
	{
		$$ = Node.new(OCONST, new(Node), new(Node))
		$$.type = types[TUVLONG]
		$$.vconst = $1
		$$.cstring = symb
	}
|	string
|	lstring

string:
	LSTRING
	{
		$$ = Node.new(OSTRING, new(Node), new(Node))
		$$.type = typ(TARRAY, types[TCHAR])
		$$.type.width = $1.l + 1
		$$.cstring = $1.s
		$$.sym = symstring
		$$.etype = TARRAY
		$$.class = CSTATIC
	}
|	string LSTRING
	{
		char *s
		int n

		n = $1.type.width - 1
		s = alloc(n+$2.l+MAXALIGN)

		memcpy(s, $1.cstring, n)
		memcpy(s+n, $2.s, $2.l)
		s[n+$2.l] = 0

		$$ = $1
		$$.type.width += $2.l
		$$.cstring = s
	}

lstring:
	LLSTRING
	{
		$$ = Node.new(OLSTRING, new(Node), new(Node))
		$$.type = typ(TARRAY, types[TUSHORT])
		$$.type.width = $1.l + sizeof(ushort)
		$$.rstring = (ushort*)$1.s
		$$.sym = symstring
		$$.etype = TARRAY
		$$.class = CSTATIC
	}
|	lstring LLSTRING
	{
		char *s
		int n

		n = $1.type.width - sizeof(ushort)
		s = alloc(n+$2.l+MAXALIGN)

		memcpy(s, $1.rstring, n)
		memcpy(s+n, $2.s, $2.l)
		*(ushort*)(s+n+$2.l) = 0

		$$ = $1
		$$.type.width += $2.l
		$$.rstring = (ushort*)s
	}

zelist:
	{
		$$ = new(Node)
	}
|	elist

elist:
	expr
|	elist ',' elist
	{
		$$ = Node.new(OLIST, $1, $3)
	}

sbody:
	'{'
	{
		$<tyty>$.t1 = strf
		$<tyty>$.t2 = strl
		$<tyty>$.t3 = lasttype
		$<tyty>$.c = lastclass
		strf = new(Type)
		strl = new(Type)
		lastbit = 0
		firstbit = 1
		lastclass = CXXX
		lasttype = new(Type)
	}
	edecl '}'
	{
		$$ = strf
		strf = $<tyty>2.t1
		strl = $<tyty>2.t2
		lasttype = $<tyty>2.t3
		lastclass = $<tyty>2.c
	}

zctlist:
	{
		lastclass = CXXX
		lasttype = types[TINT]
	}
|	ctlist

types:
	complex
	{
		$$.t = $1
		$$.c = CXXX
	}
|	tname
	{
		$$.t = simplet($1)
		$$.c = CXXX
	}
|	gcnlist
	{
		$$.t = simplet($1)
		$$.c = simplec($1)
		$$.t = garbt($$.t, $1)
	}
|	complex gctnlist
	{
		$$.t = $1
		$$.c = simplec($2)
		$$.t = garbt($$.t, $2)
		if $2 & ^BCLASS & ^BGARB {
			diag(new(Node), "duplicate types given: %T and %Q", $1, $2)
		}
	}
|	tname gctnlist
	{
		$$.t = simplet(typebitor($1, $2))
		$$.c = simplec($2)
		$$.t = garbt($$.t, $2)
	}
|	gcnlist complex zgnlist
	{
		$$.t = $2
		$$.c = simplec($1)
		$$.t = garbt($$.t, $1|$3)
	}
|	gcnlist tname
	{
		$$.t = simplet($2)
		$$.c = simplec($1)
		$$.t = garbt($$.t, $1)
	}
|	gcnlist tname gctnlist
	{
		$$.t = simplet(typebitor($2, $3))
		$$.c = simplec($1|$3)
		$$.t = garbt($$.t, $1|$3)
	}

tlist:
	types
	{
		$$ = $1.t
		if $1.c != CXXX{
			diag(new(Node), "illegal combination of class 4: %s", cnames[$1.c])
		}
	}

ctlist:
	types
	{
		lasttype = $1.t
		lastclass = $1.c
	}

complex:
	LSTRUCT ltag
	{
		dotag($2, TSTRUCT, 0)
		$$ = $2.suetag
	}
|	LSTRUCT ltag
	{
		dotag($2, TSTRUCT, autobn)
	}
	sbody
	{
		$$ = $2.suetag
		if $$.link != nil{
			diag(new(Node), "redeclare tag: %s", $2.name)
		}
		$$.link = $4
		sualign($$)
	}
|	LSTRUCT sbody
	{
		taggen++
		sprint(symb, "_%d_", taggen)
		$$ = dotag(lookup(), TSTRUCT, autobn)
		$$.link = $2
		sualign($$)
	}
|	LUNION ltag
	{
		dotag($2, TUNION, 0)
		$$ = $2.suetag
	}
|	LUNION ltag
	{
		dotag($2, TUNION, autobn)
	}
	sbody
	{
		$$ = $2.suetag
		if $$.link != nil{
			diag(new(Node), "redeclare tag: %s", $2.name)
		}
		$$.link = $4
		sualign($$)
	}
|	LUNION sbody
	{
		taggen++
		sprint(symb, "_%d_", taggen)
		$$ = dotag(lookup(), TUNION, autobn)
		$$.link = $2
		sualign($$)
	}
|	LENUM ltag
	{
		dotag($2, TENUM, 0)
		$$ = $2.suetag
		if $$.link == nil{
			$$.link = types[TINT]
		}
		$$ = $$.link
	}
|	LENUM ltag
	{
		dotag($2, TENUM, autobn)
	}
	'{'
	{
		en.tenum = new(Type)
		en.cenum = new(Type)
	}
	enum '}'
	{
		$$ = $2.suetag
		if $$.link != nil{
			diag(new(Node), "redeclare tag: %s", $2.name)
		}
		if en.tenum == nil{ 
			diag(new(Node), "enum type ambiguous: %s", $2.name)
			en.tenum = types[TINT]
		}
		$$.link = en.tenum
		$$ = en.tenum
	}
|	LENUM '{'
	{
		en.tenum = T
		en.cenum = T
	}
	enum '}'
	{
		$$ = en.tenum
	}
|	LTYPE
	{
		$$ = tcopy($1.type)
	}

gctnlist:
	gctname
|	gctnlist gctname
	{
		$$ = typebitor($1, $2)
	}

zgnlist:
	{
		$$ = 0
	}
|	zgnlist gname
	{
		$$ = typebitor($1, $2)
	}

gctname:
	tname
|	gname
|	cname

gcnlist:
	gcname
|	gcnlist gcname
	{
		$$ = typebitor($1, $2)
	}

gcname:
	gname
|	cname

enum:
	LNAME
	{
		doenum($1, new(Node))
	}
|	LNAME '=' expr
	{
		doenum($1, $3)
	}
|	enum ','
|	enum ',' enum

tname:	/* type words */
	LCHAR { $$ = BCHAR }
|	LSHORT { $$ = BSHORT }
|	LINT { $$ = BINT }
|	LLONG { $$ = BLONG }
|	LSIGNED { $$ = BSIGNED }
|	LUNSIGNED { $$ = BUNSIGNED }
|	LFLOAT { $$ = BFLOAT }
|	LDOUBLE { $$ = BDOUBLE }
|	LVOID { $$ = BVOID }

cname:	/* class words */
	LAUTO { $$ = BAUTO }
|	LSTATIC { $$ = BSTATIC }
|	LEXTERN { $$ = BEXTERN }
|	LTYPEDEF { $$ = BTYPEDEF }
|	LTYPESTR { $$ = BTYPESTR }
|	LREGISTER { $$ = BREGISTER }
|	LINLINE { $$ = 0 }

gname:	/* garbage words */
	LCONSTNT { $$ = BCONSTNT }
|	LVOLATILE { $$ = BVOLATILE }
|	LRESTRICT { $$ = 0 }

name:
	LNAME
	{
		$$ = Node.new(ONAME, new(Node), new(Node))
		if $1.class == CLOCAL{
			$1 = mkstatic($1)
		}
		$$.sym = $1
		$$.type = $1.type
		$$.etype = TVOID
		if $$.type != nil{
			$$.etype = $$.type.etype
		}
		$$.xoffset = $1.offset
		$$.class = $1.class
		$1.aused = 1
	}
tag:
	ltag
	{
		$$ = Node.new(ONAME, new(Node), new(Node))
		$$.sym = $1
		$$.type = $1.type
		$$.etype = TVOID
		if $$.type != nil{
			$$.etype = $$.type.etype
		}
		$$.xoffset = $1.offset;
		$$.class = $1.class
	}
ltag:
	LNAME
|	LTYPE
%%
