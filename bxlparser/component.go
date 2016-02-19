package bxlparser

import (
	"strconv"
	"strings"
)

// Component BXL Component
type Component struct {
	owner           *BxlParser
	data            []string
	Name            string
	PatternName     string
	OriginalName    string
	SourceLibrary   string
	RefDesPrefix    string
	NumberOfPins    int
	NumParts        int
	Composition     string
	AltIEEE         bool
	AltDeMorgan     bool
	PatternPins     int
	RevisionLevel   string
	RevisionNote    string
	CompPins        []CompPin
	Compdata        string
	AttachedSymbols []Symbol
	PinMap          map[int]int // Padnum to ComponentPin
}

// FindComponents Find components
func (b *BxlParser) FindComponents() {
	var i int
	for i < len(b.rawlines) {
		if strings.HasPrefix(b.rawlines[i], "Component ") {
			var c Component
			c.Name = DoubleQuoteContents(b.rawlines[i])
			j := i
			for j < len(b.rawlines) {
				if strings.HasPrefix(b.rawlines[j], "EndComponent") {
					c.data = b.rawlines[i+1 : j]
					c.FindCompPins()
					i = j
					b.component = c
					break
				}
				j = j + 1
			}
		}
		i = i + 1
	}
}

// ParseDescription ParseDescription
func (c *Component) ParseDescription() {

	for _, l := range c.data {
		fields := strings.Fields(l)
		switch fields[0] {
		case "PatternName":
			c.PatternName = DoubleQuoteContents(l)
		case "OriginalName":
			c.OriginalName = DoubleQuoteContents(l)
		case "SourceLibrary":
			c.SourceLibrary = DoubleQuoteContents(l)
		case "NumberOfPins":
			i, _ := strconv.ParseInt(fields[1], 10, 0)
			c.NumberOfPins = int(i)
		case "NumParts":
			j, _ := strconv.ParseInt(fields[1], 10, 0)
			c.NumParts = int(j)
		case "Composition":
			c.Composition = fields[1]
		case "AltIEEE":
			c.AltIEEE, _ = strconv.ParseBool(fields[1])
		case "AltDeMorgan":
			c.AltDeMorgan, _ = strconv.ParseBool(fields[1])
		case "PatternPins":
			k, _ := strconv.ParseInt(fields[1], 10, 0)
			c.PatternPins = int(k)
		case "Revision":
			switch fields[1] {
			case "Level":
				c.RevisionLevel = fields[2]
			case "Note":
				c.RevisionNote = fields[2]
			}
		}
	}
}
