package bxlparser

import (
	"strings"
	"unicode"
)

type Line struct {
	Origin Point
	End    Point
	Layer  Layer
	Width  string
}

type LineSlice []Line

type HasLines interface {
	AddLine(l Line)
	Data() *[]string
}

func FindLines(hl HasLines) {
	f := func(c rune) bool {
		return !unicode.IsLetter(c) && !unicode.IsNumber(c) && c != '-' && c != '.'
	}

	for _, l := range *hl.Data() {
		if strings.HasPrefix(strings.TrimSpace(l), "Line") {
			var line Line
			fields := strings.FieldsFunc(l, f)
			for j, f := range fields {
				switch f {
				case "Layer":
					line.Layer = Layer{fields[j+1]}
				case "Origin":
					line.Origin = Point{fields[j+1], fields[j+2]}
				case "EndPoint":
					line.End = Point{fields[j+1], fields[j+2]}
				case "Width":
					line.Width = fields[j+1]
				}
			}
			hl.AddLine(line)
		}
	}
}
