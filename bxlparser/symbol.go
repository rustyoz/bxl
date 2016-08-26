package bxlparser

import (
	"fmt"
	"strings"

	"github.com/rustyoz/gokicadlib"
)

type Symbol struct {
	Name  string
	data  []string
	Lines []Line
	Pins  []SymbolPin
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
			s.Name = DoubleQuoteContents(b.rawlines[i])
			j := i
			for j < len(b.rawlines) {
				if strings.HasPrefix(b.rawlines[j], "EndSymbol") {
					s.data = b.rawlines[i+1 : j]
					s.FindSymbolPins()
					FindLines(&s)
					i = j
					b.Symbol = s

					fmt.Println(s.Lines)
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
	sk.Reference.Text = bs.Name
	for _, l := range bs.Lines {
		sk.Lines = append(sk.Lines, *l.ToKicadLine())
	}

	return &sk
}
