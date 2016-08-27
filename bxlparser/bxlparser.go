package bxlparser

import "strings"

type BxlParser struct {
	input      string
	rawlines   []string
	textStyles []TextStyle
	Patterns   []Pattern
	padstacks  []PadStack
	Symbol     Symbol
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
	b.FindComponents()

	b.FindPadStacks()
	b.FindPatterns()
	b.FindSymbol()

}

func DoubleQuoteContents(s string) string {
	first := strings.Index(s, "\"")
	last := strings.LastIndex(s, "\"")

	if first != -1 && last != -1 {
		return s[first+1 : last]

	}
	return ""
}

func StyleToElements(s string) []string {
	var r []string
	var currentnumber string

	for i, c := range []byte(s) {

		if isLetter(c) { // if it is a letter directly append to output
			r = append(r, string(c))
		} else if isNumber(c) { //
			if !isNumber(s[i+1]) {
				r = append(r, currentnumber+string(c))
				currentnumber = ""
			} else {
				currentnumber = currentnumber + string(c)
			}
		}
	}

	return r
}

func isNumber(b byte) bool {
	return (b > 47 && b < 58)
}

func isLetter(b byte) bool {
	return (b > 64 && b < 91) || (b > 96 && b < 123)
}

func (b *BxlParser) TextStyles() *[]TextStyle {
	return &b.textStyles
}
