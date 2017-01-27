package bxlparser

import (
	"strings"
	"unicode"
)

func feildfuncer() func(c rune) bool {
	quoted := false
	f := func(c rune) bool {
		if quoted && c == '"' {
			quoted = false
			return true
		}
		if quoted && c != '"' {
			return false
		}
		if !quoted && c == '"' {
			quoted = true
			return true
		}

		return strings.Contains(" (),", string(c))
	}
	return f
}

func Feilds(s string) []string {
	var fs []string
	var quote bool
	var a string
	for _, c := range s {
		if quote == false && s == "\"" {
			quote = true
			continue
		}
		if quote == true && s == "\"" {
			fs = append(fs, a)
			a = ""
		}
		if quote == false {
			switch {
			case strings.ContainsAny(string(c), " (),"):
				fs = append(fs, a)
				a = ""
				continue
			case unicode.IsLetter(c), unicode.IsDigit(c):
				a = a + string(c)

			}
		}
	}
	return fs
}
