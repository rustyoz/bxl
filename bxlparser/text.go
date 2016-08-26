package bxlparser

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/rustyoz/gokicadlib"
)

// Text
type Text struct {
	owner         HasText
	Text          string
	Layer         XlrLayer
	Origin        Point
	Visible       bool
	Justification string
	Style         string
	width         float64
}

// HasText ...
type HasText interface {
	AddText(t Text)
	Data() *[]string
	TextStyles() TextStyleSlice
}

type TextSlice []Text

func FindText(ht HasText) {
	quoted := false
	f := func(c rune) bool {
		if quoted && c == '"' {
			quoted = false
			return true
		}
		if quoted && c != '"' {
			return false
		}
		if !quoted && c == '"' {
			quoted = true
			return true
		}

		return strings.Contains(" (),", string(c))
	}

	for _, l := range *ht.Data() {
		if strings.HasPrefix(strings.TrimSpace(l), "Text") {
			var text Text
			fields := strings.FieldsFunc(l, f)

			for j, f := range fields {
				switch f {
				case "Layer":
					text.Layer, _ = XlrLayerString(fields[j+1])
				case "Origin":
					text.Origin.FromString(fields[j+1], fields[j+2])
				case "Text":
					text.Text = fields[j+1]
				case "IsVisible":
					text.Visible, _ = strconv.ParseBool(fields[j+1])
				case "Justify":
					text.Justification = fields[j+1]
				case "TextStyle":
					text.Style = fields[j+1]
				}
			}

			ht.AddText(text)
		}
	}
}

func (t Text) ToKicadText() (kct *gokicadlib.Text) {
	kct = &gokicadlib.Text{}
	kct.Text = t.Text
	var err error
	kct.Layer, err = t.Layer.ToKicadLayer()
	if err != nil {
		return nil
	}
	kct.Visible = true

	kct.Origin.X = MiltoMM(t.Origin.X)
	kct.Origin.Y = MiltoMM(-t.Origin.Y)
	var height, charwidth float64
	ts, err := t.owner.TextStyles().Contains(t.Style)
	if err != nil {
		fmt.Println(err)
	}
	height = MiltoMM(float64(ts.fontHeight))
	charwidth = MiltoMM(float64(ts.fontCharWidth))
	kct.Font.Size.X = height
	kct.Font.Size.Y = charwidth
	kct.Font.Thickness = float32(MiltoMM(float64(ts.fontWidth)))

	return kct
}

func (ts TextSlice) ToKicadText() []gokicadlib.Text {
	var kcts []gokicadlib.Text
	for _, t := range ts {
		kt := t.ToKicadText()
		if kt != nil {
			kcts = append(kcts, *kt)
		}
	}
	return kcts
}
