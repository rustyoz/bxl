package bxlparser

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/rustyoz/gokicadlib"
)

// Text
type Text struct {
	owner         HasTextStyles
	Text          string
	Layer         XlrLayer
	Origin        Point
	Visible       bool
	Justification string
	Style         string
	width         float64
	data          string
}

// HasText ...
type HasText interface {
	AddText(t Text)
	Data() *[]string
	TextStyles() *[]TextStyle
}
type HasTextStyles interface {
	TextStyles() *[]TextStyle
}

type TextSlice []Text

func FindText(ht HasText) {

	for _, l := range *ht.Data() {
		l = strings.TrimSpace(l)
		if strings.HasPrefix(l, "Text") {
			var text Text
			text.parseText(l)
			ht.AddText(text)
		}
	}
}

func (t *Text) parseText(l string) {
	fields := strings.FieldsFunc(l, feildfuncer())

	for j, f := range fields {
		switch f {
		case "Layer":
			t.Layer, _ = XlrLayerString(fields[j+1])
		case "Origin":
			t.Origin.FromString(fields[j+1], fields[j+2])
		case "Text":
			t.Text = fields[j+1]
		case "IsVisible":
			t.Visible, _ = strconv.ParseBool(fields[j+1])
		case "Justify":
			t.Justification = fields[j+1]
		case "TextStyle", "TextStyleRef":
			t.Style = strings.ToLower(fields[j+1])
		}
	}
}

func (t Text) ToKicadText(mm bool) (*gokicadlib.Text, error) {
	kct := &gokicadlib.Text{}
	kct.Text = t.Text
	var err error

	kct.Visible = true

	if t.owner == nil {
		return nil, fmt.Errorf("no owner of text")
	}
	tss := t.owner.TextStyles()
	ts, err := TextStyleSlice(*tss).Contains(t.Style)
	if err != nil {
		return nil, fmt.Errorf("error converting to kicad text  %v", err)
	}

	if mm {
		kct.Origin.X = MiltoMM(t.Origin.X)
		kct.Origin.Y = MiltoMM(-t.Origin.Y)
		kct.Font.Size.X = MiltoMM(float64(ts.fontHeight))
		kct.Font.Size.Y = MiltoMM(float64(ts.fontCharWidth))
		kct.Font.Thickness = float32(MiltoMM(float64(ts.fontWidth)))
		return kct, nil
	}

	kct.Origin.X = t.Origin.X
	kct.Origin.Y = -t.Origin.Y
	kct.Font.Size.X = float64(ts.fontHeight)
	kct.Font.Size.Y = float64(ts.fontCharWidth)
	kct.Font.Thickness = float32(float64(ts.fontWidth))

	kct.Layer, err = t.Layer.ToKicadLayer()
	if err != nil {
		if !strings.HasPrefix(err.Error(), "ex") {
			return kct, err
		}
	}
	return kct, nil

}

func (ts TextSlice) ToKicadText(mm bool) ([]gokicadlib.Text, error) {
	var kcts []gokicadlib.Text
	for _, t := range ts {
		kt, err := t.ToKicadText(mm)
		if kt != nil && err != nil {
			kcts = append(kcts, *kt)
		}
		if err != nil {
			if !strings.HasPrefix(err.Error(), "ex") {
				return kcts, err
			}
			fmt.Println(err)
		}

	}
	return kcts, nil
}
