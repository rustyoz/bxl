package bxlparser

import "strings"

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
