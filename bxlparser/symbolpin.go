package bxlparser

import (
	"strconv"
	"strings"

	"github.com/rustyoz/gokicadlib"
)

// SymbolPin BXL Component pin
type SymbolPin struct {
	Owner     *Symbol
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
func (s *Symbol) FindPins() {
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

	fields := strings.FieldsFunc(lines[0], feildfuncer())
	for i, s := range fields {
		switch s {
		case "PinNum":
			p.Number, _ = strconv.Atoi(fields[i+1])

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
	p.Owner = s
	p.parseSymbolPinDes(lines[1])
	p.parseSymbolPinName(lines[2])

	s.Pins = append(s.Pins, p)
}

func (p *SymbolPin) parseSymbolPinDes(l string) {
	text := &p.PinDes
	text.owner = p.Owner.Owner

	fields := strings.FieldsFunc(l, feildfuncer())

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
		case "TextStyleRef":
			text.Style = fields[j+1]
		}
	}
	ts, err := TextStyleSlice(*text.owner.TextStyles()).Contains(text.Style)
	if err == nil {
		text.width = float64(ts.fontCharWidth) * 0.59
	}
}

func (p *SymbolPin) parseSymbolPinName(l string) {
	text := &p.PinName
	text.owner = p.Owner.Owner

	fields := strings.FieldsFunc(l, feildfuncer())

	for j, f := range fields {
		switch f {
		case "Layer":
			text.Layer, _ = XlrLayerString(fields[j+1])
		case "Origin":
			text.Origin.FromString(fields[j+1], fields[j+2])
		case "PinName":
			text.Text = fields[j+1]
			p.Name = text.Text
		case "IsVisible":
			text.Visible, _ = strconv.ParseBool(fields[j+1])
		case "Justify":
			text.Justification = fields[j+1]
		case "TextStyleRef":
			text.Style = fields[j+1]
		}
	}
	ts, err := TextStyleSlice(*text.owner.TextStyles()).Contains(text.Style)
	if err == nil {
		text.width = float64(ts.fontCharWidth) * 0.59
	}
}

func (bs SymbolPin) Kicad() *gokicadlib.Pin {
	p := &gokicadlib.Pin{}
	p.PinName = bs.Name
	p.Number = bs.Number
	p.Type = gokicadlib.Input
	p.Shape = gokicadlib.Normal
	p.Origin = gokicadlib.Point{bs.Origin.X, -bs.Origin.Y}
	p.Length = float64(bs.Length)
	p.Sizename = bs.PinName.width
	p.Sizenum = bs.PinDes.width

	switch {
	case bs.Rotation < 90:
		p.Orientation = gokicadlib.Left
		p.Origin.X += p.Length
	case bs.Rotation < 180:
		p.Orientation = gokicadlib.Up
		p.Origin.Y += p.Length
	case bs.Rotation < 270:
		p.Orientation = gokicadlib.Right
		p.Origin.X -= p.Length
	case bs.Rotation <= 360:
		p.Orientation = gokicadlib.Down
		p.Origin.Y -= p.Length
	}

	return p
}
