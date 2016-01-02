package bxlparser

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

type HasLines interface {
	AddLine(l Line)
	Data() *[]string
}
type HasArcs interface {
	AddArc(l Arc)
	Data() *[]string
}

type HasText interface {
	AddText(t Text)
	Data() *[]string
}

type BxlParser struct {
	input      string
	rawlines   []string
	textStyles []TextStyle
	patterns   []Pattern
	padstacks  []PadStack
	symbols    []Symbol
}

type TextStyle struct {
	feilds        []string
	style         string
	fontWidth     int
	fontHeight    int
	fontCharWidth int
}

type Point struct {
	x string
	y string
}

func (p Point) ToPointFloat() PointFloat {
	x, _ := strconv.ParseFloat(p.x, 64)
	y, _ := strconv.ParseFloat(p.y, 64)
	return PointFloat{x, y}
}

type PointFloat struct {
	x float64
	y float64
}

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
	HasLines
	HasArcs
	HasText
}

type PadStack struct {
	Name string
	Data []string
}
type Symbol struct {
	Name  string
	data  []string
	Lines []Line
}

type Pin struct {
	Number int
	Name   string
	_data  []string
}
type Pad struct {
	Number         int
	PinName        string
	Origin         Point
	Style          string
	OriginalStyle  string
	OriginalNumber int
}

type Poly struct {
	Layer  Layer
	Origin Point
	Points []Point
	Width  int
}

type Arc struct {
	Layer      Layer
	Origin     Point
	Radius     float64
	StartAngle float64
	SweepAngle float64
	Width      string
	End        Point
}

type Line struct {
	Origin Point
	End    Point
	Layer  Layer
	Width  string
}
type Text struct {
	Text          string
	Layer         Layer
	Origin        Point
	Visible       bool
	Justification string
	Style         string
}

type Layer struct {
	Name string
}

func NewBxlParser() *BxlParser {
	var b BxlParser
	return &b
}

func (b *BxlParser) Parse(in string) {
	b.input = in
	b.rawlines = strings.SplitAfter(b.input, "\n")
	for i, l := range b.rawlines {
		b.rawlines[i] = strings.TrimSpace(l)
	}
	b.FindTextStyles()
	b.FindPadStacks()
	b.FindPatterns()
	fmt.Println(len(b.patterns))
	fmt.Println(b.patterns[0].Lines[0].ToSExp())
	fmt.Println(b.patterns[0].Texts[0].ToSExp())
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
				case "Name":
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

func FindArcs(harcs HasArcs) {
	f := func(c rune) bool {
		return !unicode.IsLetter(c) && !unicode.IsNumber(c) && c != '-' && c != '.'
	}

	for _, l := range *harcs.Data() {
		if strings.HasPrefix(strings.TrimSpace(l), "Arc") {
			var arc Arc
			fields := strings.FieldsFunc(l, f)
			fmt.Println(fields)
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

func (p *Pattern) AddLine(l Line) {
	p.Lines = append(p.Lines, l)
}

func (p *Pattern) AddArc(a Arc) {
	p.Arcs = append(p.Arcs, a)
}
func (p *Pattern) AddText(t Text) {
	p.Texts = append(p.Texts, t)
}

func (s *Symbol) AddLine(l Line) {
	s.Lines = append(s.Lines, l)
}

func (s *Symbol) Data() *[]string {
	return &s.data
}

func (s *Pattern) Data() *[]string {
	return &s.data
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
					b.patterns = append(b.patterns, p)
					i = j
					break
				}
				j = j + 1
			}
		}
		i = i + 1
	}
}

func (b *BxlParser) FindPadStacks() {
	var i int
	for i < len(b.rawlines) {
		if strings.HasPrefix(b.rawlines[i], "PadStack ") {
			var p PadStack
			p.Name = DoubleQuoteContents(b.rawlines[i])
			j := i
			for j < len(b.rawlines) {
				if strings.HasPrefix(b.rawlines[j], "EndPadStack") {
					p.Data = b.rawlines[i : j+i]
					b.padstacks = append(b.padstacks, p)
					i = j
					break
				}
				j = j + 1

			}

		}
		i = i + 1
	}
}

func DoubleQuoteContents(s string) string {
	first := strings.Index(s, "\"")
	last := strings.LastIndex(s, "\"")

	if first != -1 && last != -1 {
		return s[first+1 : last]

	}
	return ""
}
