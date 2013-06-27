package main

import (
	"unicode"
)

type numIdentErr struct{}

func (numIdentErr) Error() string {
	return "Variable identifiers may not start with a number."
}

type invalRuneInIdentErr struct{}

func (invalRuneInIdentErr) Error() string {
	return "Variable identifiers must only contain printable runes."
}

func checkIdent(s string) error {
	rs := []rune(s)
	if unicode.IsNumber(rs[0]) {
		return numIdentErr{}
	}
	for _, r := range rs {
		if !unicode.IsPrint(r) {
			return invalRuneInIdentErr{}
		}
	}
	return nil
}
