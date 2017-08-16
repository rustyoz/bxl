package bxlparser

import (
	"fmt"
	"strconv"
	"strings"
)

type PadStack struct {
	Name      string
	data      []string
	PadShapes []PadShape
	HoleDiam  int
	Surface   bool
	Plated    bool
}

type PadStackSlice []PadStack

func (b *BxlParser) FindPadStacks() {
	var i int
	for i < len(b.rawlines) {
		if strings.HasPrefix(b.rawlines[i], "PadStack ") {
			var p PadStack
			p.Name = DoubleQuoteContents(b.rawlines[i])
			j := i
			for j < len(b.rawlines) {
				if strings.HasPrefix(b.rawlines[j], "EndPadStack") {
					p.data = b.rawlines[i:j]
					p.parsePadStack()
					for _, l := range p.data {
						var ps PadShape
						ps.parsePadShape(l)
						p.PadShapes = append(p.PadShapes, ps)
					}
					b.padstacks = append(b.padstacks, p)
					break
				}
				j = j + 1

			}

		}
		i = i + 1
	}
}

func (pss PadStackSlice) Contains(name string) (*PadStack, error) {
	for _, ps := range pss {
		if ps.Name == name {
			return &ps, nil
		}
	}
	return nil, fmt.Errorf("Padstack not found: %s", name)
}

func (ps *PadStack) parsePadStack() {
	fields := strings.FieldsFunc(ps.data[0], feildfuncer())

	for j, f := range fields {
		switch f {
		case "Plated":
			ps.Plated, _ = strconv.ParseBool(fields[j+1])

		case "Surface":
			ps.Surface, _ = strconv.ParseBool(fields[j+1])

		case "HoleDiam":
			i, _ := strconv.ParseInt(fields[j+1], 10, 32)
			ps.HoleDiam = int(i)
		}
	}
}
