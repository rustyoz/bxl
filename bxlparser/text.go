package bxlparser

import (
	"strconv"
	"strings"
)

// Text
type Text struct {
	Text          string
	Layer         Layer
	Origin        Point
	Visible       bool
	Justification string
	Style         string
	width         float64
}

// HasText ...
type HasText interface {
	AddText(t Text)
	Data() *[]string
}

func FindText(ht HasText) {
	f := func(c rune) bool {
		return c != '(' || c != ')'
	}

	for _, l := range *ht.Data() {
		if strings.HasPrefix(strings.TrimSpace(l), "Text") {
			var text Text
			fields := strings.FieldsFunc(l, f)
			for j, f := range fields {
				switch f {
				case "Layer":
					text.Layer = Layer{fields[j+1]}
				case "Origin":
					text.Origin = Point{fields[j+1], fields[j+2]}
				case "Text":
					text.Text = DoubleQuoteContents(fields[j+1])
				case "IsVisible":
					text.Visible, _ = strconv.ParseBool(fields[j+1])
				case "Justify":
					text.Justification = fields[j+1]
				case "TextStyle":
					text.Style = fields[j+1]
				}
			}
			ht.AddText(text)
		}
	}
}
