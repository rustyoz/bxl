package bxlparser

import "strings"

type PadStack struct {
	Name string
	Data []string
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
					return
				}
				j = j + 1

			}

		}
		i = i + 1
	}
}
