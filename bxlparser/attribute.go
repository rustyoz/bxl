package bxlparser

import "strings"

type Attribute struct {
	Text
	Attr [2]string
}

func (s *Symbol) FindAttributes() {
	for i, l := range s.data {
		if strings.HasPrefix(l, "Attribute") {
			var a Attribute

			a.owner = s.Owner
			a.parseAttribute(s.data[i])
			switch a.Attr[0] {
			case "RefDes":
				s.Reference = a
			case "Value":
				s.Value = a
			case "Type":
				s.Type = a
			default:
				s.Attributes = append(s.Attributes, a)

			}
		}
	}
}

func (a *Attribute) parseAttribute(l string) {
	fields := strings.FieldsFunc(l, feildfuncer())
	a.Text.parseText(l)
	for j, f := range fields {
		switch f {
		case "Attr":
			a.Attr[0] = fields[j+1]
			a.Attr[1] = fields[j+2]

		}
	}

}
