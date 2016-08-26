package bxlparser

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

// SymbolPin BXL Component pin
type SymbolPin struct {
	Number    int
	Name      string
	Origin    Point
	PinName   Text
	PinDes    Text
	Rotation  int
	Width     int
	Length    int
	IsVisible bool
}

// FindSymbolPins FindSymbolPins
func (s *Symbol) FindSymbolPins() {
	var a, b int
	for i, l := range s.data {
		if strings.HasPrefix(l, "Pin ") {
			a = i
			b = i + 3
			s.parseSymbolPin(s.data[a:b])
		}

	}

}

func (s *Symbol) parseSymbolPin(lines []string) {
	var p SymbolPin

	f := func(c rune) bool {
		return !unicode.IsLetter(c) && !unicode.IsNumber(c) || c == '(' || c == ')' || c == '"'
	}

	fields := strings.FieldsFunc(lines[0], f)
	for i, s := range fields {
		switch s {
		case "PinNum":
			p.Number, _ = strconv.Atoi(fields[i+1])
			p.Name = fields[i+2]
		case "Origin":
			p.Origin.FromString(fields[i+1], fields[i+2])
		case "PinLength":
			p.Length, _ = strconv.Atoi(fields[i+1])
		case "Rotate":
			p.Rotation, _ = strconv.Atoi(fields[i+1])
		case "Width":
			p.Width, _ = strconv.Atoi(fields[i+1])
		case "IsVisible":
			p.IsVisible, _ = strconv.ParseBool(fields[i+1])
		}
	}
	p.parseSymbolPinDes(lines[1])
	p.parseSymbolPinName(lines[2])

}

func (p *SymbolPin) parseSymbolPinDes(l string) {
	text := &p.PinDes
	quoted := false
	f := func(c rune) bool {
		if quoted && c == '"' {
			quoted = false
			return true
		}
		if quoted && c != '"' {
			return false
		}
		if !quoted && c == '"' {
			quoted = true
			return true
		}

		return strings.Contains(" (),", string(c))
	}

	fields := strings.FieldsFunc(l, f)

	for j, f := range fields {
		switch f {
		case "Layer":
			text.Layer, _ = XlrLayerString(fields[j+1])
		case "Origin":
			text.Origin.FromString(fields[j+1], fields[j+2])
		case "PinDes":
			text.Text = fields[j+1]
		case "IsVisible":
			text.Visible, _ = strconv.ParseBool(fields[j+1])
		case "Justify":
			text.Justification = fields[j+1]
		case "TextStyle":
			text.Style = fields[j+1]
		}
	}

	fmt.Println(p)
}

func (p *SymbolPin) parseSymbolPinName(l string) {
	text := &p.PinName
	quoted := false
	f := func(c rune) bool {
		if quoted && c == '"' {
			quoted = false
			return true
		}
		if quoted && c != '"' {
			return false
		}
		if !quoted && c == '"' {
			quoted = true
			return true
		}

		return strings.Contains(" (),", string(c))
	}

	fields := strings.FieldsFunc(l, f)

	for j, f := range fields {
		switch f {
		case "Layer":
			text.Layer, _ = XlrLayerString(fields[j+1])
		case "Origin":
			text.Origin.FromString(fields[j+1], fields[j+2])
		case "PinName":
			text.Text = fields[j+1]
		case "IsVisible":
			text.Visible, _ = strconv.ParseBool(fields[j+1])
		case "Justify":
			text.Justification = fields[j+1]
		case "TextStyle":
			text.Style = fields[j+1]
		}
	}
	fmt.Println(p)
}
