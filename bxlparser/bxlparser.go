package bxlparser

import (
	"strings"

	"github.com/rustyoz/gokicadlib"
)

type BxlParser struct {
	input      string
	rawlines   []string
	textStyles []TextStyle
	patterns   []Pattern
	padstacks  []PadStack
	symbol     Symbol
	component  Component
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
	b.FindComponents()
}

func DoubleQuoteContents(s string) string {
	first := strings.Index(s, "\"")
	last := strings.LastIndex(s, "\"")

	if first != -1 && last != -1 {
		return s[first+1 : last]

	}
	return ""
}

func (p *Pattern) ToKicad() gokicadlib.Module {
	var m gokicadlib.Module

	return m
}
