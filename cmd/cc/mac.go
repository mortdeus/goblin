package cc

const VARMAC = 0x80

func getnsn(void) int32 {
	c := getnsc()
	if c < '0' || c > '9' {
		return -1
	}
	n := 0
	for _; c >= '0' && c <= '9'; {
		n = n*10 + c - '0'
		c = getc()
	}
	unget(c)
	return n
}

func getsym() *Sym {
	var cp string
	c := getnsc()
	if !isalpha(c) && c != '_' && c < 0x80 {
		unget(c)
		return S
	}
	for cp = symb; ; {
		if cp <= symb+NSYMB-4 {
			//MORT: This is annoying.
			//	*cp++ = c;
		}
		c = getc()
		if isalnum(c) || c == '_' || c >= 0x80 {
			continue
		}
		unget(c)
		break
	}
	*cp = 0
	if cp > symb+NSYMB-4 {
		yyerror("symbol too large: %s", symb)
	}
	return lookup()
}

func getsymdots(int *dots) *Sym {
	s := getsym()
	if s != nil {
		return s
	}

	c := getnsc()
	if c != '.' {
		unget(c)
		return S
	}
	if getc() != '.' || getc() != '.' {
		yyerror("bad dots in macro")
	}
	*dots = 1
	return slookup("__VA_ARGS__")
}

func getcom(void) int {
	var c int

	for {
		c = getnsc()
		if c != '/' {
			break
		}
		c = getc()
		if c == '/' {
			while(c != '\n')
			c = getc()
			break
		}
		if c != '*' {
			break
		}
		c = getc()
		for {
			if c == '*' {
				c = getc()
				if c != '/' {
					continue
				}
				c = getc()
				break
			}
			if c == '\n' {
				yyerror("comment across newline")
				break
			}
			c = getc()
		}
		if c == '\n' {
			break
		}
	}
	return c
}

func dodefine(char *cp) {
	var (
		s *Sym
		p string
		l int32
	)
	ensuresymb(strlen(cp))
	strcpy(symb, cp)
	p = strchr(symb, '=')
	if p {
		//*p++ = 0;
		s = lookup()
		l = strlen(p) + 2 /* +1 null, +1 nargs */
		//s.macro = alloc(l);
		strcpy(s.macro+1, p)
	} else {
		s = lookup()
		s.macro = "\0001" /* \000 is nargs */
	}
	if debug['m'] {
		print("#define (-D) %s %s\n", s.name, s.macro+1)
	}
}

var mactab = []struct {
	macname string
	macf    func()
}{
	"ifdef", 0, /* macif(0) */
	"ifndef", 0, /* macif(1) */
	"else", 0, /* macif(2) */

	"line", maclin,
	"define", macdef,
	"include", macinc,
	"undef", macund,

	"pragma", macprag,
	"endif", macend,
}

func domacro() {
	s := getsym()
	if s == nil {
		s = slookup("endif")
	}
	for i := 0; mactab[i].macname; i++ {
		if strcmp(s.name, mactab[i].macname) == 0 {
			if mactab[i].macf {
				(*mactab[i].macf)()
			} else {
				macif(i)
			}
			return

		}
	}
	yyerror("unknown #: %s", s.name)
	macend()
}

func macund() {
	s := getsym()
	macend()
	if s == nil {
		yyerror("syntax in #undef")
		return
	}
	s.macro = 0
}

const NARG = 25

func macdef() {

	var args [NARG]string 
	var np, base string
	var i, len, ischr int

	s := getsym()
	if s == nil {
		goto bad
	}
	if s.macro {
		yyerror("macro redefined: %s", s.name)
	}
	c := getc()
	n := -1
	dots := 0
	if c == '(' {
		n++
		c = getnsc()
		if c != ')' {
			unget(c)
			for {
				a := getsymdots(&dots)
				if a == nil {
					goto bad
				}
				if n >= NARG {
					yyerror("too many arguments in #define: %s", s.name)
					goto bad
				}
				//MORT: Some where I screwed this a[n++] like logic up by putting a[n+1]. That wont work.

				args[n] = a.name
				n++
				c = getnsc()
				if c == ')' {
					break
				}
				if c != ',' || dots {
					goto bad
				}
			}
		}
		c = getc()
	}
	if isspace(c) {
		if c != '\n' {
			c = getnsc()
		}
	}
	base = hunk
	len = 1
	ischr = 0
	for {
		if isalpha(c) || c == '_' {
			np = symb
			//*np++ = c;
			c = getc()
			for _; isalnum(c) || c == '_'; {
				//*np++ = c;
				c = getc()
			}
			*np = 0
			for i = 0; i < n; i++ {
				if strcmp(symb, args[i]) == 0 {
					break
				}
				if i >= n {
					i = strlen(symb)
					base = allocn(base, len, i)
					memcpy(base+len, symb, i)
					len += i
					continue
				}
				base = allocn(base, len, 2)

				base[len] = '#'
				len++
				base[len] = 'a' + i
				len++
				continue
			}
			if ischr {
				if c == '\\' {
					base = allocn(base, len, 1)
					base[len] = c
					len++
					c = getc()
				} else if c == ischr {
					ischr = 0
				}
			} else {
				if c == '"' || c == '\'' {
					base = allocn(base, len, 1)
					base[len] = c
					len++
					ischr = c
					c = getc()
					continue
				}
				if c == '/' {
					c = getc()
					if c == '/' {
						c = getc()
						for {
							if c == '\n' {
								break
							}
							c = getc()
						}
						continue
					}
					if c == '*' {
						c = getc()
						for {
							if c == '*' {
								c = getc()
								if c != '/' {
									continue
								}
								c = getc()
								break
							}
							if c == '\n' {
								yyerror("comment and newline in define: %s", s.name)
								break
							}
							c = getc()
						}
						continue
					}
					base = allocn(base, len, 1)
					base[len] = '/'
					len++
					continue
				}
			}
			if c == '\\' {
				c = getc()
				if c == '\n' {
					c = getc()
					continue
				} else if c == '\r' {
					c = getc()
					if c == '\n' {
						c = getc()
						continue
					}
				}
				base = allocn(base, len, 1)
				base[len] = '\\'
				len++
				continue
			}
			if c == '\n' {
				break
			}
			if c == '#' {
				if n > 0 {
					base = allocn(base, len, 1)
					base[len] = c
					len++
				}
			}
			base = allocn(base, len, 1)
			base[len] = c
			len++
			fi.c--
			if fi.c < 0 {
				c = filbuf()
			} else {
				c = *fi.p & 0xff
				//fi.p++	
			}
			if c == '\n' {
				lineno++
			}
			if c == -1 {
				yyerror("eof in a macro: %s", s.name)
				break
			}
		}

		base = allocn(base, len, 1)
		base[len] = 0
		len++
		for _; len & 3; {
			base = allocn(base, len, 1)
			base[len] = 0
			len++
		}

		*base = n + 1
		if dots {
			*base |= VARMAC
		}
		s.macro = base
		if debug['m'] {
			print("#define %s %s\n", s.name, s.macro+1)
		}
		return
	}

bad:
	if s == nil {
		yyerror("syntax in #define")
	} else {
		yyerror("syntax in #define: %s", s.name)
	}
	macend()
}

func macexpand(s *Sym, b string){
	var( buf [2000]byte
	 n, l, c, nargs int
	arg [NARG]string, cp, ob, ecp *rune, dots int8;
	)
	ob = b;
	if(*s.macro == 0) {
		strcpy(b, s.macro+1);
		if(debug['m'])
			print("#expand %s %s\n", s.name, ob);
		return;
	}

	nargs = (char)(*s.macro & ~VARMAC) - 1;
	dots = *s.macro & VARMAC;

	c = getnsc();
	if(c != '(')
		goto bad;
	n = 0;
	c = getc();
	if(c != ')') {
		unget(c);
		l = 0;
		cp = buf;
		ecp = cp + sizeof(buf)-4;
		arg[n++] = cp;
		for(;;) {
			if(cp >= ecp)
				goto toobig;
			c = getc();
			if(c == '"')
				for(;;) {
					if(cp >= ecp)
						goto toobig;
					*cp++ = c;
					c = getc();
					if(c == '\\') {
						*cp++ = c;
						c = getc();
						continue;
					}
					if(c == '\n')
						goto bad;
					if(c == '"')
						break;
				}
			if(c == '\'')
				for(;;) {
					if(cp >= ecp)
						goto toobig;
					*cp++ = c;
					c = getc();
					if(c == '\\') {
						*cp++ = c;
						c = getc();
						continue;
					}
					if(c == '\n')
						goto bad;
					if(c == '\'')
						break;
				}
			if(c == '/') {
				c = getc();
				switch(c) {
				case '*':
					for(;;) {
						c = getc();
						if(c == '*') {
							c = getc();
							if(c == '/')
								break;
						}
					}
					*cp++ = ' ';
					continue;
				case '/':
					while((c = getc()) != '\n')
						;
					break;
				default:
					unget(c);
					c = '/';
				}
			}
			if(l == 0) {
				if(c == ',') {
					if(n == nargs && dots) {
						*cp++ = ',';
						continue;
					}
					*cp++ = 0;
					arg[n++] = cp;
					if(n > nargs)
						break;
					continue;
				}
				if(c == ')')
					break;
			}
			if(c == '\n')
				c = ' ';
			*cp++ = c;
			if(c == '(')
				l++;
			if(c == ')')
				l--;
		}
		*cp = 0;
	}
	if(n != nargs) {
		yyerror("argument mismatch expanding: %s", s.name);
		*b = 0;
		return;
	}
	cp = s.macro+1;
	for(;;) {
		c = *cp++;
		if(c == '\n')
			c = ' ';
		if(c != '#') {
			*b++ = c;
			if(c == 0)
				break;
			continue;
		}
		c = *cp++;
		if(c == 0)
			goto bad;
		if(c == '#') {
			*b++ = c;
			continue;
		}
		c -= 'a';
		if(c < 0 || c >= n)
			continue;
		strcpy(b, arg[c]);
		b += strlen(arg[c]);
	}
	*b = 0;
	if(debug['m'])
		print("#expand %s %s\n", s.name, ob);
	return;

bad:
	yyerror("syntax in macro expansion: %s", s.name);
	*b = 0;
	return;

toobig:
	yyerror("too much text in macro expansion: %s", s.name);
	*b = 0;
}

// func macinc(void){
// 	int c0, c, i, f;
// 	char str[STRINGSZ], *hp;

// 	c0 = getnsc();
// 	if(c0 != '"') {
// 		c = c0;
// 		if(c0 != '<')
// 			goto bad;
// 		c0 = '>';
// 	}
// 	for(hp = str;;) {
// 		c = getc();
// 		if(c == c0)
// 			break;
// 		if(c == '\n')
// 			goto bad;
// 		*hp++ = c;
// 	}
// 	*hp = 0;

// 	c = getcom();
// 	if(c != '\n')
// 		goto bad;

// 	f = -1;
// 	for(i=0; i<ninclude; i++) {
// 		if(i == 0 && c0 == '>')
// 			continue;
// 		ensuresymb(strlen(include[i])+strlen(str)+2);
// 		strcpy(symb, include[i]);
// 		strcat(symb, "/");
// 		if(strcmp(symb, "./") == 0)
// 			symb[0] = 0;
// 		strcat(symb, str);
// 		f = open(symb, OREAD);
// 		if(f >= 0)
// 			break;
// 	}
// 	if(f < 0)
// 		strcpy(symb, str);
// 	c = strlen(symb) + 1;
// 	hp = alloc(c);
// 	memcpy(hp, symb, c);
// 	newio();
// 	pushio();
// 	newfile(hp, f);
// 	return;

// bad:
// 	unget(c);
// 	yyerror("syntax in #include");
// 	macend();
// }

// func maclin(void){
// 	char *cp;
// 	int c;
// 	int32 n;

// 	n = getnsn();
// 	c = getc();
// 	if(n < 0)
// 		goto bad;

// 	for(;;) {
// 		if(c == ' ' || c == '\t') {
// 			c = getc();
// 			continue;
// 		}
// 		if(c == '"')
// 			break;
// 		if(c == '\n') {
// 			strcpy(symb, "<noname>");
// 			goto nn;
// 		}
// 		goto bad;
// 	}
// 	cp = symb;
// 	for(;;) {
// 		c = getc();
// 		if(c == '"')
// 			break;
// 		*cp++ = c;
// 	}
// 	*cp = 0;
// 	c = getcom();
// 	if(c != '\n')
// 		goto bad;

// nn:
// 	c = strlen(symb) + 1;
// 	cp = alloc(c);
// 	memcpy(cp, symb, c);
// 	linehist(cp, n);
// 	return;

// bad:
// 	unget(c);
// 	yyerror("syntax in #line");
// 	macend();
// }

// func macif(int f){
// 	int c, l, bol;
// 	Sym *s;

// 	if(f == 2)
// 		goto skip;
// 	s = getsym();
// 	if(s == S)
// 		goto bad;
// 	if(getcom() != '\n')
// 		goto bad;
// 	if((s.macro != 0) ^ f)
// 		return;

// skip:
// 	bol = 1;
// 	l = 0;
// 	for(;;) {
// 		c = getc();
// 		if(c != '#') {
// 			if(!isspace(c))
// 				bol = 0;
// 			if(c == '\n')
// 				bol = 1;
// 			continue;
// 		}
// 		if(!bol)
// 			continue;
// 		s = getsym();
// 		if(s == S)
// 			continue;
// 		if(strcmp(s.name, "endif") == 0) {
// 			if(l) {
// 				l--;
// 				continue;
// 			}
// 			macend();
// 			return;
// 		}
// 		if(strcmp(s.name, "ifdef") == 0 || strcmp(s.name, "ifndef") == 0) {
// 			l++;
// 			continue;
// 		}
// 		if(l == 0 && f != 2 && strcmp(s.name, "else") == 0) {
// 			macend();
// 			return;
// 		}
// 	}

// bad:
// 	yyerror("syntax in #if(n)def");
// 	macend();
// }

// func macprag(void){
// 	Sym *s;
// 	int c0, c;
// 	char *hp;
// 	Hist *h;

// 	s = getsym();

// 	if(s && strcmp(s.name, "lib") == 0)
// 		goto praglib;
// 	if(s && strcmp(s.name, "pack") == 0) {
// 		pragpack();
// 		return;
// 	}
// 	if(s && strcmp(s.name, "fpround") == 0) {
// 		pragfpround();
// 		return;
// 	}
// 	if(s && strcmp(s.name, "textflag") == 0) {
// 		pragtextflag();
// 		return;
// 	}
// 	if(s && strcmp(s.name, "dataflag") == 0) {
// 		pragdataflag();
// 		return;
// 	}
// 	if(s && strcmp(s.name, "varargck") == 0) {
// 		pragvararg();
// 		return;
// 	}
// 	if(s && strcmp(s.name, "incomplete") == 0) {
// 		pragincomplete();
// 		return;
// 	}
// 	if(s && strcmp(s.name, "dynimport") == 0) {
// 		pragdynimport();
// 		return;
// 	}
// 	if(s && strcmp(s.name, "dynexport") == 0) {
// 		pragdynexport();
// 		return;
// 	}
// 	if(s && strcmp(s.name, "dynlinker") == 0) {
// 		pragdynlinker();
// 		return;
// 	}
// 	while(getnsc() != '\n')
// 		;
// 	return;

// praglib:
// 	c0 = getnsc();
// 	if(c0 != '"') {
// 		c = c0;
// 		if(c0 != '<')
// 			goto bad;
// 		c0 = '>';
// 	}
// 	for(hp = symb;;) {
// 		c = getc();
// 		if(c == c0)
// 			break;
// 		if(c == '\n')
// 			goto bad;
// 		*hp++ = c;
// 	}
// 	*hp = 0;
// 	c = getcom();
// 	if(c != '\n')
// 		goto bad;

// 	/*
// 	 * put pragma-line in as a funny history
// 	 */
// 	c = strlen(symb) + 1;
// 	hp = alloc(c);
// 	memcpy(hp, symb, c);

// 	h = alloc(sizeof(Hist));
// 	h.name = hp;
// 	h.line = lineno;
// 	h.offset = -1;
// 	h.link = H;
// 	if(ehist == H) {
// 		hist = h;
// 		ehist = h;
// 		return;
// 	}
// 	ehist.link = h;
// 	ehist = h;
// 	return;

// bad:
// 	unget(c);
// 	yyerror("syntax in #pragma lib");
// 	macend();
// }

// func macend(void){
// 	int c;

// 	for(;;) {
// 		c = getnsc();
// 		if(c < 0 || c == '\n')
// 			return;
// 	}
// }

// func linehist(char *f, int offset){
// 	Hist *h;

// 	/*
// 	 * overwrite the last #line directive if
// 	 * no alloc has happened since the last one
// 	 */
// 	if(newflag == 0 && ehist != H && offset != 0 && ehist.offset != 0)
// 		if(f && ehist.name && strcmp(f, ehist.name) == 0) {
// 			ehist.line = lineno;
// 			ehist.offset = offset;
// 			return;
// 		}

// 	if(debug['f'])
// 		if(f) {
// 			if(offset)
// 				print("%4d: %s (#line %d)\n", lineno, f, offset);
// 			else
// 				print("%4d: %s\n", lineno, f);
// 		} else
// 			print("%4d: <pop>\n", lineno);
// 	newflag = 0;

// 	h = alloc(sizeof(Hist));
// 	h.name = f;
// 	h.line = lineno;
// 	h.offset = offset;
// 	h.link = H;
// 	if(ehist == H) {
// 		hist = h;
// 		ehist = h;
// 		return;
// 	}
// 	ehist.link = h;
// 	ehist = h;
// }
