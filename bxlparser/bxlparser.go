package bxlparser

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

type BxlParser struct {
	input      string
	lines      []string
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
	x float64
	y float64
}
type Pattern struct {
	Name        string
	OriginPoint Point
	PickPoint   Point
	GluePoint   Point
	Data        []string
}

type PadStack struct {
	Name string
	Data []string
}
type Symbol struct {
	Name string
	Data []string
}

type Pin struct {
	Number int
	Name   string
	Data   []string
}
type Pad struct {
	Number          int
	Name            string
	Origin          Point
	Style           string
	OriginalStyle   string
	OringinalNumber int
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
	Width      int
}

type Line struct {
	Origin Point
	End    Point
	Layer  Layer
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
	b.lines = strings.SplitAfter(b.input, "\n")
	for i, l := range b.lines {
		b.lines[i] = strings.TrimSpace(l)
	}
	//fmt.Println("Lines:", len(b.lines))
	b.FindTextStyles()
	b.FindPadStacks()
	b.FindPatterns()
}

func (b *BxlParser) FindTextStyles() {
	f := func(c rune) bool {
		return !unicode.IsLetter(c) && !unicode.IsNumber(c)
	}

	for _, l := range b.lines {
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

func (b *BxlParser) FindPatterns() {
	var i int
	for i < len(b.lines) {
		if strings.HasPrefix(b.lines[i], "Pattern ") {
			var p Pattern
			p.Name = DoubleQuoteContents(b.lines[i])
			for j, end := range b.lines[i:] {
				if strings.HasPrefix(end, "EndPattern") {
					p.Data = b.lines[i+1 : j+i]
					i = j
				}
			}
			fmt.Print(p.Name)
			fmt.Print(p.Data)
			b.patterns = append(b.patterns, p)
		}
		i = i + 1
	}
}

func (b *BxlParser) FindPadStacks() {
	var i int
	for i < len(b.lines) {
		if strings.HasPrefix(b.lines[i], "PadStack ") {
			var p PadStack
			p.Name = DoubleQuoteContents(b.lines[i])
			for j, end := range b.lines[i:] {
				if strings.HasPrefix(end, "EndPadStack") {
					p.Data = b.lines[i : j+i]
					i = j
				}
			}
			fmt.Print(p.Name)
			fmt.Print(p.Data)
			b.padstacks = append(b.padstacks, p)
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
