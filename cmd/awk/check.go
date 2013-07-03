package main

import (
	"unicode"
)

func checkIdent(s string) error {
	rs := []rune(s)
	if unicode.IsNumber(rs[0]) {
		return numIdentErr(0)
	}
	for _, r := range rs {
		if !unicode.IsPrint(r) {
			return invalRuneInIdentErr(0)
		}
	}
	return nil
}

func checkStringType(s string) error {
	return nil
}
