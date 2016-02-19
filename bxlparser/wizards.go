package bxlparser

import (
	"strings"
	"unicode"
)

type Wizard struct {
	VarName string
	VarData string
	Number  int
	Origin  Point
}

type HasWizards interface {
	AddWizard(w Wizard)
	Data() *[]string
}

func (b *BxlParser) FindWizards(hw HasWizards) {
	f := func(c rune) bool {
		return !unicode.IsLetter(c) && !unicode.IsNumber(c) && c != '-' && c != '.'
	}

	for _, l := range *hw.Data() {
		if strings.HasPrefix(strings.TrimSpace(l), "Line") {
			var w Wizard
			fields := strings.FieldsFunc(l, f)
			for j, f := range fields {
				switch f {
				case "Origin":
					w.Origin.FromString(fields[j+1], fields[j+2])
				case "VarName":
					w.VarName = fields[j+1]
				case "VarData":
					w.VarData = fields[j+1]
				}
			}
			hw.AddWizard(w)
		}
	}
}
