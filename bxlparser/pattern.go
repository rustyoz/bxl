package bxlparser

import (
	"strconv"
	"strings"
	"unicode"

	"github.com/rustyoz/gokicadlib"
)

// Pattern is the bxl name for a footprint
type Pattern struct {
	owner       *BxlParser
	component   *Component
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
			p.component = &b.component
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
					p.owner = b
					b.Patterns = append(b.Patterns, p)

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
	t.owner = p
	p.Texts = append(p.Texts, t)
}

func (p *Pattern) Data() *[]string {
	return &p.data
}

func (p *Pattern) TextStyles() TextStyleSlice {
	return p.owner.textStyles
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
					pad.PinName = fields[j+1]
				case "Number":
					pad.Number, _ = strconv.Atoi(fields[j+1])
				case "OriginalNumber":
					pad.OriginalNumber, _ = strconv.Atoi(fields[j+1])
				case "PadStyle":
					pad.Style = fields[j+1]
				case "OriginalPadStyle":
					pad.OriginalStyle = fields[j+1]
				case "OriginalPinNumber":
					pad.OriginalNumber, _ = strconv.Atoi(fields[j+1])
				case "Origin":
					pad.Origin.FromString(fields[j+1], fields[j+2])
				}
			}
			p.Pads = append(p.Pads, pad)
		}
	}
}

func (p *Pattern) ToKicad() gokicadlib.Module {
	var m gokicadlib.Module
	m.Layer = gokicadlib.F_Cu
	m.Name = p.Name
	m.Descr = p.Name
	m.Reference.Text = "REF**"
	m.Reference.Type = "reference"
	m.Reference.Layer = gokicadlib.F_SilkS
	m.Value.Type = "value"
	m.Value.Text = p.component.Name
	m.Value.Layer = gokicadlib.F_Fab
	m.Tags = []string{p.Name}
	m.Pads = PadSlice(p.Pads).ToKicadPads()
	m.Lines = LineSlice(p.Lines).ToKicadLines()
	m.Text = TextSlice(p.Texts).ToKicadText()
	m.Tstamp.Stamp()

	return m
}
