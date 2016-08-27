package bxlparser

import (
	"fmt"
	"strconv"
	"strings"
)

type TextStyle struct {
	feilds        []string
	style         string
	fontWidth     int
	fontHeight    int
	fontCharWidth int
}

type TextStyleSlice []TextStyle

func (b *BxlParser) FindTextStyles() {

	for _, l := range b.rawlines {
		if strings.HasPrefix(strings.TrimSpace(l), "TextStyle") {
			var ts TextStyle
			ts.feilds = strings.FieldsFunc(l, feildfuncer())
			ts.style = strings.ToLower(ts.feilds[1])
			for j, f := range ts.feilds {
				switch f {
				case "FontWidth":
					ts.fontWidth, _ = strconv.Atoi(ts.feilds[j+1])
				case "FontHeight":
					ts.fontHeight, _ = strconv.Atoi(ts.feilds[j+1])
				case "FontCharWidth":
					ts.fontCharWidth, _ = strconv.Atoi(ts.feilds[j+1])
				}
			}
			if ts.fontCharWidth == 0 {
				ts.fontCharWidth = ts.fontHeight
			}
			b.textStyles = append(b.textStyles, ts)

		}
	}
}

func (tss TextStyleSlice) Contains(s string) (ts TextStyle, err error) {
	for _, ss := range []TextStyle(tss) {
		if ss.style == strings.ToLower(s) {
			return ss, nil
		}
	}
	err = fmt.Errorf("TextStyle %s not found", s)
	return TextStyle{}, err
}
