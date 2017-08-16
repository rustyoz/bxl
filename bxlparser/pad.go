package bxlparser

import (
	"fmt"

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
	Rotate         int
	owner          *BxlParser
}

type PadSlice []Pad

func (p *Pad) ToKicadPad() (gokicadlib.Pad, error) {
	var kkp gokicadlib.Pad
	kkp.Number = p.Number
	kkp.PinName = p.PinName

	kkp.Origin.X = MiltoMM(p.Origin.X)
	kkp.Origin.Y = MiltoMM(-p.Origin.Y)
	kkp.Origin.R = p.Rotate

	padstack, err := PadStackSlice(p.owner.padstacks).Contains(p.Style)
	if err != nil {
		return kkp, fmt.Errorf("Error: ToKicadPad %v  %v", p, err)
	}

	// find padshape for top layer
	for _, ps := range padstack.PadShapes {
		if ps.Layer == TOP {
			kkp.Layers = gokicadlib.LayerSlice{gokicadlib.F_Cu, gokicadlib.F_Paste, gokicadlib.F_Mask}
			// padsize
			kkp.Size.X = MiltoMM(ps.Width)
			kkp.Size.Y = MiltoMM(ps.Height)

			// pad stype
			if padstack.Surface {
				kkp.Padtype = gokicadlib.Surface
				kkp.Layers = gokicadlib.LayerSlice{gokicadlib.F_Cu, gokicadlib.F_Paste, gokicadlib.F_Mask}
			} else {
				kkp.Drillsize = MiltoMM(float64(padstack.HoleDiam))
				if padstack.Plated {
					kkp.Padtype = gokicadlib.ThroughHole
					kkp.Layers = gokicadlib.LayerSlice{gokicadlib.A_Cu, gokicadlib.A_Mask}
				} else {
					kkp.Padtype = gokicadlib.NotPlatedThroughHole
				}
			}

			// kicad pad shape from PadShape Description
			switch ps.Description {
			case "Rectangle":
				kkp.Shape = gokicadlib.Rectangle
			case "Round":
				kkp.Shape = gokicadlib.Circle
			case "Oblong":
				kkp.Shape = gokicadlib.Oval
			}
		}
	}

	return kkp, nil

}

func (ps PadSlice) ToKicadPads() ([]gokicadlib.Pad, error) {
	var kkps []gokicadlib.Pad
	for _, p := range ps {
		nkkp, err := p.ToKicadPad()
		if err != nil {
			return kkps, fmt.Errorf("PadSlice.ToKicadPads error: %v", err)
		}
		kkps = append(kkps, nkkp)
	}
	return kkps, nil
}
