package bxlparser

import (
	"strconv"
	"strings"
	"unicode"
)

// Pattern is the bxl name for a footprint
type Pattern struct {
	Name        string
	OriginPoint Point
	PickPoint   Point
	GluePoint   Point
	data        []string
	Pads        []Pad
	Lines       []Line
	Arcs        []Arc
	Texts       []Text
	Polygons    []Polygon
	HasLines
	HasArcs
	HasText
	HasPolygon
}

func (b *BxlParser) FindPatterns() {
	var i int
	for i < len(b.rawlines) {
		if strings.HasPrefix(b.rawlines[i], "Pattern ") {
			var p Pattern
			p.Name = DoubleQuoteContents(b.rawlines[i])
			j := i
			for j < len(b.rawlines) {
				if strings.HasPrefix(b.rawlines[j], "EndPattern") {
					p.data = b.rawlines[i+1 : j+i]
					p.FindPads()
					FindLines(&p)
					FindArcs(&p)
					FindText(&p)
					FindPolygon(&p)
					b.patterns = append(b.patterns, p)
					return
					break
				}
				j = j + 1
			}
		}
		i = i + 1
	}
}

func (p *Pattern) AddLine(l Line) {
	p.Lines = append(p.Lines, l)
}

func (p *Pattern) AddArc(a Arc) {
	p.Arcs = append(p.Arcs, a)
}
func (p *Pattern) AddText(t Text) {
	p.Texts = append(p.Texts, t)
}

func (p *Pattern) Data() *[]string {
	return &p.data
}

func (p *Pattern) AddPolygon(poly Polygon) {
	p.Polygons = append(p.Polygons, poly)
}

func (p *Pattern) FindPads() {
	f := func(c rune) bool {
		return !unicode.IsLetter(c) && !unicode.IsNumber(c) && c != '-' && c != '.'
	}

	for _, l := range p.data {
		if strings.HasPrefix(strings.TrimSpace(l), "Pad") {
			var pad Pad
			fields := strings.FieldsFunc(l, f)

			for j, f := range fields {
				switch f {
				case "PinName":
					pad.PinName = DoubleQuoteContents(fields[j+1])
				case "Number":
					pad.Number, _ = strconv.Atoi(fields[j+1])
				case "OriginalNumber":
					pad.OriginalNumber, _ = strconv.Atoi(fields[j+1])
				case "PadStyle":
					pad.Style = DoubleQuoteContents(fields[j+1])
				case "OriginalPadStyle":
					pad.OriginalStyle = DoubleQuoteContents(fields[j+1])
				case "OriginalPinNumber":
					pad.OriginalNumber, _ = strconv.Atoi(fields[j+1])
				case "Origin":
					pad.Origin = Point{fields[j+1], fields[j+2]}
				}
			}
			p.Pads = append(p.Pads, pad)
		}
	}
}
