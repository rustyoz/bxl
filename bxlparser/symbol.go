package bxlparser

import (
	"fmt"
	"strings"

	"github.com/rustyoz/gokicadlib"
)

type Symbol struct {
	Owner      *BxlParser
	Reference  Attribute
	Value      Attribute
	Type       Attribute
	Name       string
	data       []string
	Lines      []Line
	Pins       []SymbolPin
	Text       []Text
	Attributes []Attribute
}

func (s *Symbol) AddLine(l Line) {
	s.Lines = append(s.Lines, l)
}

func (s *Symbol) Data() *[]string {
	return &s.data
}

// FindComponents Find components
func (b *BxlParser) FindSymbol() {
	var i int
	for i < len(b.rawlines) {
		if strings.HasPrefix(b.rawlines[i], "Symbol ") {
			var s Symbol
			s.Owner = b
			s.Name = DoubleQuoteContents(b.rawlines[i])
			j := i
			for j < len(b.rawlines) {
				if strings.HasPrefix(b.rawlines[j], "EndSymbol") {
					s.data = b.rawlines[i+1 : j]
					s.FindPins()
					s.FindAttributes()
					FindLines(&s)
					i = j
					b.Symbol = s
					break
				}
				j = j + 1
			}
		}
		i = i + 1
	}
}

func (bs *Symbol) Kicad() *gokicadlib.Symbol {
	var sk gokicadlib.Symbol

	sk.Name = bs.Name
	sk.Reference = *bs.Reference.Text.ToKicadText(false)
	fmt.Println("sk.Referece ", sk.Reference)
	sk.Value = *bs.Value.Text.ToKicadText(false)
	fmt.Println("sk.Value ", sk.Value)
	sk.Reference.Text = string(bs.Name[0])
	sk.Value.Text = bs.Name

	for _, l := range bs.Lines {
		sk.Lines = append(sk.Lines, *l.ToKicadLine(false))
	}
	for _, p := range bs.Pins {
		sk.Pins = append(sk.Pins, *p.Kicad())
	}
	for _, t := range bs.Text {
		sk.Text = append(sk.Text, *t.ToKicadText(false))
	}

	return &sk
}
