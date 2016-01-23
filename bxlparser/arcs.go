package bxlparser

import (
	"strconv"
	"strings"
	"unicode"
)

type HasArcs interface {
	AddArc(l Arc)
	Data() *[]string
}

func FindArcs(harcs HasArcs) {
	f := func(c rune) bool {
		return !unicode.IsLetter(c) && !unicode.IsNumber(c) && c != '-' && c != '.'
	}

	for _, l := range *harcs.Data() {
		if strings.HasPrefix(strings.TrimSpace(l), "Arc") {
			var arc Arc
			fields := strings.FieldsFunc(l, f)
			for j, f := range fields {
				switch f {
				case "Layer":
					arc.Layer = Layer{fields[j+1]}
				case "Origin":
					arc.Origin = Point{fields[j+1], fields[j+2]}
				case "StartAngle":
					arc.StartAngle, _ = strconv.ParseFloat(fields[j+2], 64)
				case "Radius":
					arc.Radius, _ = strconv.ParseFloat(fields[j+2], 64)
				case "Width":
					arc.Width = fields[j+1]
				}
			}
			harcs.AddArc(arc)
		}
	}
}
