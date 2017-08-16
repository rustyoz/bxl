package bxlparser

import (
	"fmt"
	"strings"
)

type HasAttributes interface {
	AddAttribute(a Attribute)
	Data() *[]string
	GetOwner() *BxlParser
}

type AttributeSlice []Attribute

type Attribute struct {
	Text
	Attr [2]string
}

func FindAttributes(ha HasAttributes) {
	for _, l := range *ha.Data() {
		if strings.HasPrefix(l, "Attribute") {
			var a Attribute

			a.owner = ha.GetOwner()
			a.parseAttribute(l)
			ha.AddAttribute(a)
		}
	}
}

func (a *Attribute) parseAttribute(l string) {
	fields := strings.FieldsFunc(l, feildfuncer())
	a.Text.parseText(l)
	fmt.Println(l)
	for j, f := range fields {
		switch f {
		case "Attr":
			a.Attr[0] = fields[j+1]
			a.Attr[1] = fields[j+2]
		}
	}
	fmt.Println(a.Origin)
}

func (ats AttributeSlice) Contains(name string) (*Attribute, error) {
	for _, a := range ats {
		if a.Attr[0] == name {
			return &a, nil
		}
	}
	return nil, fmt.Errorf("Attribute not found: %v", name)
}
