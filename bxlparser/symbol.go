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

func (s Symbol) GetOwner() *BxlParser {
	return s.Owner
}

func (s *Symbol) AddLine(l Line) {
	s.Lines = append(s.Lines, l)
}

func (s *Symbol) Data() *[]string {
	return &s.data
}

func (s *Symbol) AddAttribute(a Attribute) {
	s.Attributes = append(s.Attributes, a)
}

// FindComponents Find components
func (b *BxlParser) FindSymbols() {
	var i int
	for i < len(b.rawlines) {
		if strings.HasPrefix(b.rawlines[i], "Symbol ") {
			var s Symbol
			s.Owner = b
			s.Value.owner = b
			s.Reference.owner = b
			s.Type.owner = b
			s.Name = DoubleQuoteContents(b.rawlines[i])
			j := i
			for j < len(b.rawlines) {
				if strings.HasPrefix(b.rawlines[j], "EndSymbol") {
					s.data = b.rawlines[i+1 : j]
					s.FindPins()
					FindAttributes(&s)

					FindLines(&s)
					i = j
					b.Symbols = append(b.Symbols, s)
					i = j
					break
				}
				j = j + 1
			}
		}
		i = i + 1
	}
}

func (bs *Symbol) Kicad() (*gokicadlib.Symbol, error) {
	var sk gokicadlib.Symbol
	var err error
	sk.Name = bs.Name

	ref, err := AttributeSlice(bs.Attributes).Contains("RefDes")
	if err == nil {
		bs.Reference = *ref
		fmt.Println(bs.Reference.Origin)
	} else {
		fmt.Println(err)
	}

	v, err := AttributeSlice(bs.Attributes).Contains("Value")
	if err == nil {
		bs.Value = *v
	} else {
		fmt.Println(err)
	}
	tp, err := AttributeSlice(bs.Attributes).Contains("Type")
	if err == nil {
		bs.Type = *tp
	} else {
		fmt.Println(err)
	}

	des, err := AttributeSlice(bs.GetOwner().component.Attributes).Contains("DESCRIPTION")
	if err == nil {
		sk.Description = des.Attr[1]
	}
	datasheet, err := AttributeSlice(bs.GetOwner().component.Attributes).Contains("Datasheet URL")
	if err == nil {
		sk.Documentation = datasheet.Attr[1]
	}

	for _, p := range bs.Owner.Patterns {
		sk.FootPrints = append(sk.FootPrints, p.Name)
	}

	bs.Reference.Text.Style = "H50S3"
	bs.Reference.Text.Layer = TOP_SILKSCREEN
	t, err := bs.Reference.Text.ToKicadText(false)
	if err != nil {
		fmt.Println(err)
	}
	sk.Reference = *t
	bs.Value.Text.Style = "H50S3"
	bs.Value.Text.Layer = TOP_ASSEMBLY
	t, err = bs.Value.Text.ToKicadText(false)
	sk.Value = *t
	sk.Reference.Text = string(bs.Name[0])
	sk.Value.Text = bs.Name

	for _, l := range bs.Lines {
		sk.Lines = append(sk.Lines, *l.ToKicadLine(false))
	}
	for _, p := range bs.Pins {
		sk.Pins = append(sk.Pins, *p.Kicad())
	}
	for _, t := range bs.Text {
		tt, e := t.ToKicadText(false)
		if e != nil {
			err = e
			return &sk, fmt.Errorf("%v \r\n %v", t, e)
		}
		sk.Text = append(sk.Text, *tt)
	}

	if err != nil {
		return nil, fmt.Errorf("%v \r\n %v", bs, err)
	}
	return &sk, err
}
