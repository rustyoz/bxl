package bxlparser

import (
	"strconv"
	"strings"
	"unicode"
)

type TextStyle struct {
	feilds        []string
	style         string
	fontWidth     int
	fontHeight    int
	fontCharWidth int
}

func (b *BxlParser) FindTextStyles() {
	f := func(c rune) bool {
		return !unicode.IsLetter(c) && !unicode.IsNumber(c)
	}

	for _, l := range b.rawlines {
		if strings.HasPrefix(strings.TrimSpace(l), "TextStyle") {
			var ts TextStyle
			ts.feilds = strings.FieldsFunc(l, f)
			ts.style = ts.feilds[1]
			for j, f := range ts.feilds {
				switch f {
				case "FontWidth":
					ts.fontWidth, _ = strconv.Atoi(ts.feilds[j+1])
				case "FontHeight":
					ts.fontHeight, _ = strconv.Atoi(ts.feilds[j+1])
				case "FontCharWidth":
					ts.fontCharWidth, _ = strconv.Atoi(ts.feilds[j+1])
				}
			}
			b.textStyles = append(b.textStyles, ts)

		}
	}
}
