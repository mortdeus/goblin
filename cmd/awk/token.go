package main

type token int

const (
	NONE = iota
	NAME
	NUMBER
	STRING
	ERE

	FUNC_NAME

	BEGIN
	END

	BREAK
	CLOSE
	CONTINUE
	DELETE
	DO
	ELSE
	EXIT
	FOR
	FUNCTION
	IF
	IN
	NEXT
	NEXTFILE
	PRINT
	PRINTF
	RETURN
	WHILE

	BUILTIN
	GETLINE

	AS    // =
	ADDAS // +=
	SUBAS // -=
	MULAS // *=
	DIVAS // /=
	MODAS // %=
	POWAS // ^=

	OR      // ||
	AND     // &&
	NOMATCH // !~
	MATCH   // ~
	EQ      // ==
	NEQ     // !=
	LEQ     // <=
	GEQ     // >=
	INC     // ++
	DEC     // --
	APPEND  // >>

	LBRACE // {
	RBRACE // }
	LPAREN // (
	RPAREN // )
	LBRACK // [
	RBRACK // ]
	LCHEV  // <
	RCHEV  // >

	COMMA     // ,
	SEMICOLON // ;
	NEWL      // \n
	ADD       // +
	SUB       // -
	MUL       // *
	DIV       // /
	MOD       // %
	POW       // ^
	NOT       // !
	BAR       // |
	COND      // ?
	COLON     // :
	FIELDREF  // $
)

const (
	FLENGTH = iota
	FSQRT
	FEXP
	FLOG
	FINT
	FSYSTEM
	FRAND
	FSRAND
	FSIN
	FCOS
	FATAN
	FTOLOWER
	FTOUPPER
	FFLUSH
	FUTF
	FMATCH
	FSPLIT
	FGSUB
	FSUBSTR
	FSPRINTF
	FINDEX
)

func keyword(l *lexeme) bool {
	switch l.val {
	case "BEGIN":
		l.typ = BEGIN
	case "END":
		l.typ = END
	case "atan2":
		l.typ, l.sub = BUILTIN, FATAN
	case "break":
		l.typ = BREAK
	case "close":
		l.typ = CLOSE
	case "continue":
		l.typ = CONTINUE
	case "cos":
		l.typ, l.sub = BUILTIN, FCOS
	case "delete":
		l.typ = DELETE
	case "do":
		l.typ = DO
	case "else":
		l.typ = ELSE
	case "exit":
		l.typ = EXIT
	case "exp":
		l.typ, l.sub = BUILTIN, FEXP
	case "fflush":
		l.typ, l.sub = BUILTIN, FFLUSH
	case "for":
		l.typ = FOR
	case "func", "function":
		l.typ = FUNCTION
	case "getline":
		l.typ = GETLINE
	case "gsub":
		l.typ, l.sub = BUILTIN, FGSUB
	case "if":
		l.typ = IF
	case "in":
		l.typ = IN
	case "index":
		l.typ, l.sub = BUILTIN, FINDEX
	case "int":
		l.typ, l.sub = BUILTIN, FINT
	case "length":
		l.typ, l.sub = BUILTIN, FLENGTH
	case "log":
		l.typ, l.sub = BUILTIN, FLOG
	case "match":
		l.typ, l.sub = BUILTIN, FMATCH
	case "next":
		l.typ = NEXT
	case "nextfile":
		l.typ = NEXTFILE
	case "print":
		l.typ = PRINT
	case "printf":
		l.typ = PRINTF
	case "rand":
		l.typ, l.sub = BUILTIN, FRAND
	case "return":
		l.typ = RETURN
	case "sin":
		l.typ, l.sub = BUILTIN, FSIN
	case "split":
		l.typ, l.sub = BUILTIN, FSPLIT
	case "sprintf":
		l.typ, l.sub = BUILTIN, FSPRINTF
	case "sqrt":
		l.typ, l.sub = BUILTIN, FSQRT
	case "srand":
		l.typ, l.sub = BUILTIN, FSRAND
	case "sub":
		l.typ = SUB
	case "substr":
		l.typ, l.sub = BUILTIN, FSUBSTR
	case "system":
		l.typ, l.sub = BUILTIN, FSYSTEM
	case "tolower":
		l.typ, l.sub = BUILTIN, FTOLOWER
	case "toupper":
		l.typ, l.sub = BUILTIN, FTOUPPER
	case "utf":
		l.typ, l.sub = BUILTIN, FUTF
	case "while":
		l.typ = WHILE
	default:
		return false
	}
	return true
}
