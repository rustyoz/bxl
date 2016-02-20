package bxlparser

import (
	"strconv"

	"github.com/rustyoz/gokicadlib"
)

// Pad ...
type Pad struct {
	Number         int
	PinName        string
	Origin         Point
	Style          string
	OriginalStyle  string
	OriginalNumber int
}

type PadSlice []Pad

func (p *Pad) ToKicadPad() gokicadlib.Pad {
	var kkp gokicadlib.Pad
	kkp.Number = p.Number
	kkp.PinName = p.PinName

	kkp.Origin.X = MiltoMM(p.Origin.X)
	kkp.Origin.Y = MiltoMM(-p.Origin.Y)
	styleelements := StyleToElements(p.Style)
	for i, e := range styleelements {
		switch e {
		case "x", "X":
			xs, _ := strconv.Atoi(styleelements[i+1])
			kkp.Size.X = MiltoMM(float64(xs))
		case "y", "Y":
			xs, _ := strconv.Atoi(styleelements[i+1])
			kkp.Size.Y = MiltoMM(float64(xs))
		case "d", "D":
			ds, _ := strconv.Atoi(styleelements[i+1])
			kkp.Drillsize = MiltoMM(float64(ds))
		}
	}

	for _, e := range styleelements {
		switch e {
		case "r", "R":
			kkp.Shape = gokicadlib.Rectangle
		case "s", "S":
			kkp.Shape = gokicadlib.Rectangle
		case "e", "E":
			if kkp.Size.X == kkp.Size.Y {
				kkp.Shape = gokicadlib.Circle
			} else {
				kkp.Shape = gokicadlib.Oval
			}
		case "p", "P":
			kkp.Padtype = gokicadlib.ThroughHole
			kkp.Layers = []gokicadlib.Layer{gokicadlib.A_Cu, gokicadlib.A_Mask, gokicadlib.F_SilkS}
		case "t", "T":
			kkp.Padtype = gokicadlib.Surface
			kkp.Layers = []gokicadlib.Layer{gokicadlib.F_Cu, gokicadlib.F_Paste, gokicadlib.F_Mask}
		}
	}

	return kkp

}

func (ps PadSlice) ToKicadPads() []gokicadlib.Pad {
	var kkps []gokicadlib.Pad
	for _, p := range ps {
		kkps = append(kkps, p.ToKicadPad())
	}
	return kkps
}
