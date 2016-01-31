package bxlparser

import (
	"fmt"
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
	o := p.Origin.ToPointFloat()
	kkp.Origin.X = fmt.Sprint(MiltoMM(o.X))
	kkp.Origin.Y = fmt.Sprint(MiltoMM(o.Y))
	styleelements := StyleToElements(p.Style)
	for i, e := range styleelements {
		switch e {
		case "x", "X":
			xs, _ := strconv.Atoi(styleelements[i+1])
			kkp.Size.X = fmt.Sprint(MiltoMM(float64(xs)))
		case "y", "Y":
			xs, _ := strconv.Atoi(styleelements[i+1])
			kkp.Size.Y = fmt.Sprint(MiltoMM(float64(xs)))
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
		case "t", "T":
			kkp.Padtype = gokicadlib.Surface
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

func StyleToElements(s string) []string {
	var r []string
	var currentnumber string

	for i, c := range []byte(s) {

		if isLetter(c) { // if it is a letter directly append to output
			r = append(r, string(c))
		} else if isNumber(c) { //
			if !isNumber(s[i+1]) {
				r = append(r, currentnumber+string(c))
				currentnumber = ""
			} else {
				currentnumber = currentnumber + string(c)
			}
		}
	}

	return r
}

func isNumber(b byte) bool {
	return (b > 47 && b < 58)
}

func isLetter(b byte) bool {
	return (b > 64 && b < 91) || (b > 96 && b < 123)
}
