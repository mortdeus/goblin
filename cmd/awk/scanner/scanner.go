// http://mortdeus.mit-license.org

// Derived from go's go/token standard pkg.
// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package scanner implements a scanner for AWK script.
// It takes a []byte as source which can then be tokenized
// through repeated calls to the Scan method.
//
package scanner

import (
	//"bytes"
	"fmt"
	"github.com/mortdeus/goblin/cmd/awk/token"
	"path/filepath"
	//"strconv"
	"unicode"
	"unicode/utf8"
)

// An ErrorHandler may be provided to Scanner.Init. If a syntax error is
// encountered and a handler was installed, the handler is called with a
// position and an error message. The position points to the beginning of
// the offending token.
//
type ErrorHandler func(pos token.Position, msg string)

// A Scanner holds the scanner's internal state while processing
// a given text.  It can be allocated as part of another data
// structure but must be initialized via Init before use.
//
type Scanner struct {
	// immutable state
	file *token.File  // source file handle
	dir  string       // directory portion of file.Name()
	src  []byte       // source
	err  ErrorHandler // error reporting; or nil

	// scanning state
	ch         rune // current character
	offset     int  // character offset
	rdOffset   int  // reading offset (position after current character)
	lineOffset int  // current line offset

	// public state - ok to modify
	ErrorCount int // number of errors encountered
}

const bom = 0xFEFF // byte order mark, only permitted as very first character

// Read the next Unicode char into s.ch.
// s.ch < 0 means end-of-file.
//
func (s *Scanner) next() {
	if s.rdOffset < len(s.src) {
		s.offset = s.rdOffset
		if s.ch == '\n' {
			s.lineOffset = s.offset
			s.file.AddLine(s.offset)
		}
		r, w := rune(s.src[s.rdOffset]), 1
		switch {
		case r == 0:
			s.error(s.offset, "illegal character NUL")
		case r >= 0x80:
			// not ASCII
			r, w = utf8.DecodeRune(s.src[s.rdOffset:])
			if r == utf8.RuneError && w == 1 {
				s.error(s.offset, "illegal UTF-8 encoding")
			} else if r == bom && s.offset > 0 {
				s.error(s.offset, "illegal byte order mark")
			}
		}
		s.rdOffset += w
		s.ch = r
	} else {
		s.offset = len(s.src)
		if s.ch == '\n' {
			s.lineOffset = s.offset
			s.file.AddLine(s.offset)
		}
		s.ch = -1 // eof
	}
}

// Init prepares the scanner s to tokenize the text src by setting the
// scanner at the beginning of src. The scanner uses the file set file
// for position information and it adds line information for each line.
// It is ok to re-use the same file when re-scanning the same file as
// line information which is already present is ignored. Init causes a
// panic if the file size does not match the src size.
//
// Calls to Scan will invoke the error handler err if they encounter a
// syntax error and err is not nil. Also, for each error encountered,
// the Scanner field ErrorCount is incremented by one. The mode parameter
// determines how comments are handled.
//
// Note that Init may call err if there is an error in the first character
// of the file.
//
func (s *Scanner) Init(file *token.File, src []byte, err ErrorHandler) {
	// Explicitly initialize all fields since a scanner may be reused.
	if file.Size() != len(src) {
		panic(fmt.Sprintf("file size (%d) does not match src len (%d)", file.Size(), len(src)))
	}
	s.file = file
	s.dir, _ = filepath.Split(file.Name())
	s.src = src
	s.err = err

	s.ch = ' '
	s.offset = 0
	s.rdOffset = 0
	s.lineOffset = 0
	s.ErrorCount = 0

	s.next()
	if s.ch == bom {
		s.next() // ignore BOM at file beginning
	}
}

func (s *Scanner) error(offs int, msg string) {
	if s.err != nil {
		s.err(s.file.Position(s.file.Pos(offs)), msg)
	}
	s.ErrorCount++
}
func (s *Scanner) scanComment() string {
	// initial '/' already consumed; s.ch == '/' || s.ch == '*'
	offs := s.offset - 1 // position of initial '/'
	hasCR := false

	for s.ch != '\n' && s.ch >= 0 {
		if s.ch == '\r' {
			hasCR = true
		}
		s.next()
	}
	lit := s.src[offs:s.offset]
	if hasCR {
		lit = stripCR(lit)
	}

	return string(lit)
}

func isLetter(ch rune) bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_' || ch >= 0x80 && unicode.IsLetter(ch)
}

func isDigit(ch rune) bool {
	return '0' <= ch && ch <= '9' || ch >= 0x80 && unicode.IsDigit(ch)
}

func (s *Scanner) scanIdentifier() string {
	offs := s.offset
	for isLetter(s.ch) || isDigit(s.ch) {
		s.next()
	}
	return string(s.src[offs:s.offset])
}

func digitVal(ch rune) int {
	if '0' <= ch && ch <= '9' {
		return int(ch - '0')
	}

	return 10 // larger than any legal digit val
}

func (s *Scanner) scanMantissa(base int) {
	for digitVal(s.ch) < base {
		s.next()
	}
}

func (s *Scanner) scanNumber(seenDecimalPoint bool) (token.Token, string) {
	// digitVal(s.ch) < 10
	offs := s.offset
	tok := token.NUM

	if seenDecimalPoint {
		offs--
		s.scanMantissa(10)
		goto exponent
	}

	if s.ch == '0' {
		// int or float
		offs = s.offset
		s.next()

		s.scanMantissa(10)
		if s.ch == '.' || s.ch == 'e' || s.ch == 'E' {
			goto fraction
		}

		goto exit
	}

	// decimal int or float
	s.scanMantissa(10)

fraction:
	if s.ch == '.' {
		s.next()
		s.scanMantissa(10)
	}

exponent:
	if s.ch == 'e' || s.ch == 'E' {
		s.next()
		if s.ch == '-' || s.ch == '+' {
			s.next()
		}
		s.scanMantissa(10)
	}
exit:
	return token.Token(tok), string(s.src[offs:s.offset])
}

func (s *Scanner) scanEscape(quote rune) {
	offs := s.offset

	var i, base, max uint32
	switch s.ch {
	case 'a', 'b', 'f', 'n', 'r', 't', 'v', '\\', quote:
		s.next()
		return
	case '0', '1', '2', '3', '4', '5', '6', '7':
		i, base, max = 3, 8, 255
	case 'x':
		s.next()
		i, base, max = 2, 16, 255
	case 'u':
		s.next()
		i, base, max = 4, 16, unicode.MaxRune
	case 'U':
		s.next()
		i, base, max = 8, 16, unicode.MaxRune
	default:
		s.next() // always make progress
		s.error(offs, "unknown escape sequence")
		return
	}

	var x uint32
	for ; i > 0 && s.ch != quote && s.ch >= 0; i-- {
		d := uint32(digitVal(s.ch))
		if d >= base {
			s.error(s.offset, "illegal character in escape sequence")
			break
		}
		x = x*base + d
		s.next()
	}
	// in case of an error, consume remaining chars
	for ; i > 0 && s.ch != quote && s.ch >= 0; i-- {
		s.next()
	}
	if x > max || 0xD800 <= x && x < 0xE000 {
		s.error(offs, "escape sequence is invalid Unicode code point")
	}
}

func (s *Scanner) scanChar() string {
	// '\'' opening already consumed
	offs := s.offset - 1

	n := 0
	for s.ch != '\'' {
		ch := s.ch
		n++
		s.next()
		if ch == '\n' || ch < 0 {
			s.error(offs, "character literal not terminated")
			n = 1
			break
		}
		if ch == '\\' {
			s.scanEscape('\'')
		}
	}

	s.next()

	if n != 1 {
		s.error(offs, "illegal character literal")
	}

	return string(s.src[offs:s.offset])
}

func (s *Scanner) scanString() string {
	// '"' opening already consumed
	offs := s.offset - 1

	for s.ch != '"' {
		ch := s.ch
		s.next()
		if ch == '\n' || ch < 0 {
			s.error(offs, "string not terminated")
			break
		}
		if ch == '\\' {
			s.scanEscape('"')
		}
	}

	s.next()

	return string(s.src[offs:s.offset])
}

func stripCR(b []byte) []byte {
	c := make([]byte, len(b))
	i := 0
	for _, ch := range b {
		if ch != '\r' {
			c[i] = ch
			i++
		}
	}
	return c[:i]
}

func (s *Scanner) skipWhitespace() {
	for s.ch == ' ' || s.ch == '\t' || s.ch == '\r' {
		s.next()
	}
}

// Helper functions for scanning multi-byte tokens such as >> += >>= .
// Different routines recognize different length tok_i based on matches
// of ch_i. If a token ends in '=', the result is tok1 or tok3
// respectively. Otherwise, the result is tok0 if there was no other
// matching character, or tok2 if the matching character was ch2.

func (s *Scanner) switch2(tok0, tok1 token.Token) token.Token {
	if s.ch == '=' {
		s.next()
		return tok1
	}
	return tok0
}

func (s *Scanner) switch3(tok0, tok1 token.Token, ch2 rune, tok2 token.Token) token.Token {
	if s.ch == '=' {
		s.next()
		return tok1
	}
	if s.ch == ch2 {
		s.next()
		return tok2
	}
	return tok0
}

func (s *Scanner) switch4(tok0, tok1 token.Token, ch2 rune, tok2, tok3 token.Token) token.Token {
	if s.ch == '=' {
		s.next()
		return tok1
	}
	if s.ch == ch2 {
		s.next()
		if s.ch == '=' {
			s.next()
			return tok3
		}
		return tok2
	}
	return tok0
}

// Scan scans the next token and returns the token position, the token,
// and its literal string if applicable. The source end is indicated by
// token.EOF.
//
// If the returned token is a literal (token.IDENT, token.INT, token.FLOAT,
// token.IMAG, token.CHAR, token.STRING) or token.COMMENT, the literal string
// has the corresponding value.
//
// If the returned token is a keyword, the literal string is the keyword.
//
// If the returned token is token.SEMICOLON, the corresponding
// literal string is ";" if the semicolon was present in the source,
// and "\n" if the semicolon was inserted because of a newline or
// at EOF.
//
// If the returned token is token.ILLEGAL, the literal string is the
// offending character.
//
// In all other cases, Scan returns an empty literal string.
//
// For more tolerant parsing, Scan will return a valid token if
// possible even if a syntax error was encountered. Thus, even
// if the resulting token sequence contains no illegal tokens,
// a client may not assume that no error occurred. Instead it
// must check the scanner's ErrorCount or the number of calls
// of the error handler, if there was one installed.
//
// Scan adds line information to the file added to the file
// set with Init. Token positions are relative to that file
// and thus relative to the file set.
//
func (s *Scanner) Scan() (pos token.Pos, tok token.Token, lit string) {
scanAgain:
	s.skipWhitespace()
	// current token start
	pos = s.file.Pos(s.offset)
	// determine token value

	switch ch := s.ch; {
	case isLetter(ch):
		lit = s.scanIdentifier()
		if len(lit) > 1 {
			// keywords are longer than one letter - avoid lookup otherwise
			tok = token.Lookup(lit)
		} else {
			tok = token.IDENT
		}
	case '0' <= ch && ch <= '9':
		tok, lit = s.scanNumber(false)

	default:
		s.next() // always make progress

		switch ch {
		case -1:
			tok = token.EOF
		case '\n':
			tok = token.NEWLINE
		case '"':
			tok = token.STRING
			lit = s.scanString()
		case '\'':
			tok = token.STRING
			lit = s.scanChar()
		case '.':
			if '0' <= s.ch && s.ch <= '9' {
				tok, lit = s.scanNumber(true)
			}
		case ',':
			tok = token.COMMA
		case ';':
			tok = token.SEMICOLON
			lit = ";"
		case '(':
			tok = token.LPAREN
		case ')':
			tok = token.RPAREN
		case '[':
			tok = token.LBRACKET
		case ']':
			tok = token.RBRACKET
		case '{':
			tok = token.LBRACE
		case '}':
			tok = token.RBRACE
		case '+':
			tok = s.switch3(token.ADD, token.ADD_ASSIGN, '+', token.INCR)
		case '-':
			tok = s.switch3(token.SUB, token.SUB_ASSIGN, '-', token.DECR)
		case '*':
			tok = s.switch2(token.MUL, token.MUL_ASSIGN)
		case '/':
			//TODO: Implement REGEXP logic.
			tok = s.switch2(token.DIV, token.DIV_ASSIGN)
		case '#':
			lit = s.scanComment()
			goto scanAgain
		case '%':
			tok = s.switch2(token.MOD, token.MOD_ASSIGN)
		case '^':
			tok = s.switch2(token.POW, token.POW_ASSIGN)
		case '<':
			tok = s.switch2(token.LT, token.LTE)
		case '>':
			tok = s.switch3(token.GT, token.GTE, '>', token.APPEND)
		case '=':
			tok = s.switch2(token.ASSIGN, token.EQ)
		case '!':
			tok = s.switch3(token.NOT, token.NEQ, '~', token.NO_MATCH)
		case '&':
			if s.ch == '&' {
				tok = token.AND
			}

		case '|':
			if s.ch == '|' {
				tok = token.OR
			} else {
				tok = token.PIPE
			}
		case '?':
			tok = token.COND
		case ':':
			tok = token.COLON
		case '~':
			tok = token.MATCH
		case '$':
			tok = token.REF

		default:
			// next reports unexpected BOMs - don't repeat
			if ch != bom {
				s.error(s.file.Offset(pos), fmt.Sprintf("illegal character %#U", ch))
			}
			tok = token.BAD
			lit = string(ch)
		}
	}

	return
}
