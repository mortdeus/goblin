// http://mortdeus.mit-license.org

// Derived from go's go/token standard pkg.
// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package token

import "strconv"

type Token int

const (
	BAD = iota
	EOF
	COMMENT // #

	literalBegin
	IDENT
	NUM
	STRING
	ERE
	FUNC_IDENT
	literalEnd

	keywordBegin
	BEGIN
	END
	BREAK
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
	PRINT
	PRINTF
	RETURN
	WHILE

	BUILTIN

	GETLINE
	keywordEnd

	operatorBegin
	ADD_ASSIGN // +=
	SUB_ASSIGN // -=
	MUL_ASSIGN // *=
	DIV_ASSIGN // /=
	MOD_ASSIGN // *=
	POW_ASSIGN // ^=

	OR       // ||
	AND      // &&
	MATCH    // ~
	NO_MATCH // !~
	EQ       // ==
	LT       // <
	LTE      // <=
	GT       // >
	GTE      // >=
	NEQ      // !=
	INCR     // ++
	DECR     // --
	APPEND   // >>

	LBRACE    // {
	RBRACE    // }
	LPAREN    // (
	RPAREN    // )
	LBRACKET  // [
	RBRACKET  // ]
	COMMA     // ,
	SEMICOLON // ;
	NEWLINE   // \n
	ADD       // +
	SUB       // -
	MUL       // *
	DIV       // /
	MOD       // %
	POW       // ^
	NOT       // !
	PIPE      // |
	COND      // ?
	COLON     // :
	REF       // $
	ASSIGN    // =
	operatorEnd
)

var tokens = [...]string{
	BAD:     "BAD",
	EOF:     "EOF",
	COMMENT: "COMMENT",

	IDENT:      "IDENT",
	NUM:        "NUM",
	STRING:     "STRING",
	ERE:        "ERE",
	FUNC_IDENT: "FUNC_IDENT",

	BEGIN:    "BEGIN",
	END:      "END",
	BREAK:    "break",
	CONTINUE: "continue",
	DELETE:   "delete",
	DO:       "do",
	ELSE:     "else",
	EXIT:     "exit",
	FOR:      "for",
	FUNCTION: "function",
	IF:       "if",
	IN:       "in",
	NEXT:     "next",
	PRINT:    "print",
	PRINTF:   "printf",
	RETURN:   "return",
	WHILE:    "while",

	BUILTIN: "BUILTIN",

	GETLINE: "getline",

	ADD_ASSIGN: "+=",
	SUB_ASSIGN: "-=",
	MUL_ASSIGN: "*=",
	DIV_ASSIGN: "/=",
	MOD_ASSIGN: "*=",
	POW_ASSIGN: "^=",

	OR:       "||",
	AND:      "&&",
	MATCH:    "~",
	NO_MATCH: "!~",
	EQ:       "==",
	LT:       "<",
	LTE:      "<=",
	GT:       ">",
	GTE:      ">=",
	NEQ:      "!=",
	INCR:     "++",
	DECR:     "--",
	APPEND:   ">>",

	LBRACE:    "{",
	RBRACE:    "}",
	LPAREN:    "(",
	RPAREN:    ")",
	LBRACKET:  "[",
	RBRACKET:  "]",
	COMMA:     ",",
	SEMICOLON: ";",
	NEWLINE:   "\n",
	ADD:       "+",
	SUB:       "-",
	MUL:       "*",
	DIV:       "/",
	MOD:       "%",
	POW:       "^",
	NOT:       "!",
	PIPE:      "|",
	COND:      "?",
	COLON:     ":",
	REF:       "$",
	ASSIGN:    "=",
}

func (tok Token) String() string {
	s := ""
	if 0 <= tok && tok < Token(len(tokens)) {
		s = tokens[tok]
	}
	if s == "" {
		s = "Token(" + strconv.Itoa(int(tok)) + ")"
	}
	return s
}

// A set of constants for precedence-based expression parsing.
// Non-operators have lowest precedence, followed by operators
// starting with precedence 1 up to unary operators. The highest
// precedence serves as "catch-all" precedence for selector,
// indexing, and other operator and delimiter tokens.
//
const (
	LowestPrec       = 0 // non-operators
	StringConcatPrec = 8
	UnaryPrec        = 9
	HighestPrec      = 15
)

// Precedence returns the operator precedence of the binary
// operator op. If op is not a binary operator, the result
// is LowestPrecedence.
//
func (op Token) Precedence() int {
	switch op {
	case LPAREN, RPAREN:
		return 14
	case REF:
		return 13
	case INCR, DECR:
		return 12
	case POW:
		return 11
	case ADD, SUB, NOT, MUL, DIV, MOD:
		return 10
	case LT, LTE, GT, GTE, EQ, NEQ, APPEND, PIPE:
		return 7
	case MATCH, NO_MATCH:
		return 6
	case IN:
		return 5
	case AND:
		return 4
	case OR:
		return 3
	case COND:
		return 2
	case ASSIGN, ADD_ASSIGN, SUB_ASSIGN,
		MUL_ASSIGN, DIV_ASSIGN, MOD_ASSIGN, POW_ASSIGN:
		return 1
	}
	return LowestPrec
}

var keywords map[string]Token

func init() {
	keywords = make(map[string]Token)
	for i := keywordBegin + 1; i < keywordEnd; i++ {
		keywords[tokens[i]] = Token(i)
	}
}

// Lookup maps an identifier to its keyword token or IDENT (if not a keyword).
//
func Lookup(ident string) Token {
	if tok, is_keyword := keywords[ident]; is_keyword {
		return tok
	}
	return IDENT
}

// Predicates

// IsLiteral returns true for tokens corresponding to identifiers
// and basic type literals; it returns false otherwise.
//
func (tok Token) IsLiteral() bool { return literalBegin < tok && tok < literalEnd }

// IsOperator returns true for tokens corresponding to operators and
// delimiters; it returns false otherwise.
//
func (tok Token) IsOperator() bool { return operatorBegin < tok && tok < operatorEnd }

// IsKeyword returns true for tokens corresponding to keywords;
// it returns false otherwise.
func (tok Token) IsKeyword() bool { return keywordBegin < tok && tok < keywordEnd }
