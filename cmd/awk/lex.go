package main

import (
	"bytes"
	"fmt"
	"io"
	"unicode"
	"unicode/utf8"
)

type lexeme struct {
	val string
	typ token
	sub token
}
type prog struct {
	buf       *bytes.Buffer
	lineno    int
	lexemes   []lexeme
	arithExpr bool
	nlbrack,
	nlparen,
	nlbrace int
	running bool
}

func lex(c *context) {
	s := c.script
	s.lexemes = make([]lexeme, 0, 256)
	s.running = true
	for s.running {
		switch r := string(s.get()); {

		case r == "_", unicode.IsLetter(rune(r[0])):
			var b = []byte(r)

			for c := string(s.peek()); unicode.IsDigit(rune(c[0])) ||
				c == "_" || unicode.IsLetter(rune(c[0])); c = string(s.peek()) {
				b = append(b, s.get()...)
			}

			l := lexeme{val: string(b)}
			if !keyword(&l) {
				if string(s.peek()) == "(" {
					l.typ = FUNC_NAME
					s.nlparen++
				} else {
					l.typ = NAME
				}
			}
			s.lexemes = append(s.lexemes, l)
		case unicode.IsNumber(rune(r[0])), r == ".":
			seenDot := false
			seenE := false
			seenNegOp := false
			if r == "." {
				seenDot = true
			}
			var b = []byte(r)
			for {
				c := string(s.peek())

				if unicode.IsNumber(rune(c[0])) {
					b = append(b, c...)
				} else if c == "." && !seenDot {
					seenDot = true
					b = append(b, byte('.'))
				} else if c == "e" && seenDot && !seenE {
					seenE = true
					b = append(b, byte('e'))
				} else if c == "-" && seenDot && seenE && !seenNegOp {
					seenNegOp = true
					b = append(b, byte('-'))
				} else {
					break
				}
			}
			s.lexemes = append(s.lexemes, lexeme{val: string(b)})

		case r == "#":
			_, _ = s.buf.ReadString('\n')
			s.lineno++

		case r == "\n":
			s.lexemes = append(s.lexemes, lexeme{"\\n", NEWL, NONE})
			s.lineno++
		case r == " ", r == "\t", r == "\r":

		case r == "/":
			if !s.arithExpr {
				var b []byte
				for {
					switch string(s.peek()) {
					case "\n":
						fatal(fmt.Errorf("Error: Newline in regexp.\n"))
					case "/":
						s.get()
						s.lexemes = append(s.lexemes, lexeme{string(b), ERE, NONE})
						break
					default:
						b = append(b, s.get()...)
					}
				}
			} else {
				if string(s.peek()) == "=" {
					s.get()
					s.lexemes = append(s.lexemes, lexeme{"/=", DIVAS, NONE})

				} else {
					s.lexemes = append(s.lexemes, lexeme{"/", DIV, NONE})
				}
			}
		case r == ";":
			s.lexemes = append(s.lexemes, lexeme{";", SEMICOLON, NONE})
		case r == "+":
			switch string(s.peek()) {
			case "+":
				s.get()
				s.lexemes = append(s.lexemes, lexeme{"++", INC, NONE})

			case "=":
				s.get()
				s.lexemes = append(s.lexemes, lexeme{"+=", ADDAS, NONE})

			default:
				s.lexemes = append(s.lexemes, lexeme{"+", ADD, NONE})
			}
		case r == "-":
			switch string(s.peek()) {
			case "-":
				s.get()
				s.lexemes = append(s.lexemes, lexeme{"-", DEC, NONE})

			case "=":
				s.get()
				s.lexemes = append(s.lexemes, lexeme{"-=", SUBAS, NONE})

			default:
				s.lexemes = append(s.lexemes, lexeme{"-", SUB, NONE})
			}
		case r == "*":
			switch string(s.peek()) {
			case "=":
				s.get()
				s.lexemes = append(s.lexemes, lexeme{"*=", MULAS, NONE})

			case "*":
				s.get()
				if string(s.peek()) == "=" {
					s.lexemes = append(s.lexemes, lexeme{"**=", POWAS, NONE})
					s.get()
				} else {
					s.lexemes = append(s.lexemes, lexeme{"**", POW, NONE})
				}

			default:
				s.lexemes = append(s.lexemes, lexeme{"*", MUL, NONE})
			}
		case r == ">":
			switch string(s.peek()) {
			case ">":
				s.get()
				s.lexemes = append(s.lexemes, lexeme{">>", APPEND, NONE})
			case "=":
				s.get()
				s.lexemes = append(s.lexemes, lexeme{">=", GEQ, NONE})
			default:
				s.lexemes = append(s.lexemes, lexeme{">", GT, NONE})
			}
		case r == "<":
			if string(s.peek()) == "=" {
				s.get()
				s.lexemes = append(s.lexemes, lexeme{"<<", LEQ, NONE})
			} else {
				s.lexemes = append(s.lexemes, lexeme{"<", LT, NONE})
			}
		case r == "~":
			s.lexemes = append(s.lexemes, lexeme{"~", MATCH, NONE})
		case r == "=":
			if string(s.peek()) == "=" {
				s.get()
				s.lexemes = append(s.lexemes, lexeme{"==", EQ, NONE})
			} else {
				s.lexemes = append(s.lexemes, lexeme{"=", AS, NONE})
			}
		case r == `\`:
			s.lexemes = append(s.lexemes, lexeme{";", SEMICOLON, NONE})

		case r == "&":
			if string(s.peek()) == "&" {
				s.lexemes = append(s.lexemes, lexeme{"&&", AND, NONE})
			}

		case r == "|":
			if string(s.peek()) == "|" {
				s.lexemes = append(s.lexemes, lexeme{"||", OR, NONE})
				s.get()
			}
		case r == "!":
			switch string(s.peek()) {
			case "=":
				s.get()
				s.lexemes = append(s.lexemes, lexeme{"!=", NEQ, NONE})
			case "~":
				s.get()
				s.lexemes = append(s.lexemes, lexeme{"!~", NOMATCH, NONE})
			default:
				s.lexemes = append(s.lexemes, lexeme{"!", NOT, NONE})
			}
		case r == "?":
			s.lexemes = append(s.lexemes, lexeme{"?", COND, NONE})
		case r == ":":
			s.lexemes = append(s.lexemes, lexeme{":", COLON, NONE})
		/*case r == "$":
		s.lexemes = append(s.lexemes, lexeme{"$", FIELDREF, NONE})
		*/
		case r == `%`:
			if string(s.peek()) == "=" {
				s.get()
				s.lexemes = append(s.lexemes, lexeme{`%=`, MODAS, NONE})
			} else {
				s.lexemes = append(s.lexemes, lexeme{`%`, MOD, NONE})
			}
		case r == "^":
			if string(s.peek()) == "=" {
				s.lexemes = append(s.lexemes, lexeme{"^=", POWAS, NONE})
			}
			s.lexemes = append(s.lexemes, lexeme{"^", POW, NONE})

		/*case r == `"`:
		s.lexemes = append(s.lexemes, lexeme{`"`, st, NONE})
		*/

		case r == "{":
			s.nlbrack++
			s.lexemes = append(s.lexemes, lexeme{"{", LBRACK, NONE})
		case r == "[":
			s.nlbrace++
			s.lexemes = append(s.lexemes, lexeme{"[", LBRACE, NONE})
		case r == "(":
			s.nlparen++
			s.lexemes = append(s.lexemes, lexeme{"(", LPAREN, NONE})

		case r == "}":
			s.nlbrack--
			if s.nlbrack > 0 && s.nlbrack%2 != 1 {
				fatal(fmt.Errorf("Error: missing { \n"))
			}
			s.lexemes = append(s.lexemes, lexeme{"}", RBRACK, NONE})
		case r == "]":
			s.nlbrace--
			if s.nlbrack > 0 && s.nlbrack%2 != 1 {
				fatal(fmt.Errorf("Error: missing [ \n"))
			}
			s.lexemes = append(s.lexemes, lexeme{"]", RBRACE, NONE})
		case r == ")":
			s.nlparen--
			if s.nlbrack > 0 && s.nlbrack%2 != 1 {
				fatal(fmt.Errorf("Error: missing ( \n"))
			}
			s.lexemes = append(s.lexemes, lexeme{")", RPAREN, NONE})
		default:
			s.lexemes = append(s.lexemes, lexeme{r, NONE, NONE})

		}
	}
}

func (p *prog) peek() (b []byte) {
	r, i, err := p.buf.ReadRune()
	if err != nil {
		if err == io.EOF {
			p.running = false
		}
	}
	b = make([]byte, i)
	utf8.EncodeRune(b, r)

	if err := p.buf.UnreadRune(); err != nil {
		fatal(err)
	}
	return

}
func (p *prog) get() (b []byte) {

	r, i, err := p.buf.ReadRune()
	if err != nil {
		if err == io.EOF {
			p.running = false

		} else {
			fatal(err)
		}
	}
	if i == 0 {
		return []byte{0}
	}
	b = make([]byte, i)
	utf8.EncodeRune(b, r)

	return
}
