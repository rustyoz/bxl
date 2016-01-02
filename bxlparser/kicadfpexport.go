package bxlparser

import (
	"fmt"
	"math"
	"strings"
)

func NewSExp(name string, newline bool, contents ...string) string {
	var output string
	if newline {
		output = "(" + name + "\n"
		for _, s := range contents {
			output = output + "    " + strings.Replace(s, "\n", "\n    ", -1) + "\n"
		}
		output = output + ")"
	} else {
		output = "(" + name + " "
		for _, s := range contents {
			output = output + s + ")"
		}
	}

	return output
}

func (line Line) ToSExp() string {
	start := NewSExp("start", false, line.Origin.ToString())
	end := NewSExp("end", false, line.End.ToString())
	layer := NewSExp("layer", false, line.Layer.Name)
	width := NewSExp("width", false, line.Width)
	return NewSExp("fp_line", true, start, end, layer, width)
}

func (p Point) ToString() string {
	return p.x + " " + p.y
}

func (arc Arc) ToSExp() string {
	start := NewSExp("start", false, arc.Origin.ToString())
	arc.CalculateEndPoint()
	end := NewSExp("end", false, arc.End.ToString())
	angle := NewSExp("angle", false, fmt.Sprintf("%f", arc.StartAngle))
	layer := NewSExp("layer", false, arc.Layer.Name)
	width := NewSExp("width", false, arc.Width)
	return NewSExp("fp_line", true, start, end, angle, layer, width)
}

func (arc *Arc) CalculateEndPoint() {
	origin := arc.Origin.ToPointFloat()
	vector := PointFloat{arc.Radius * math.Cos(arc.StartAngle*180/math.Pi), arc.Radius * math.Sin(arc.StartAngle*180/math.Pi)}
	end := PointFloat{origin.x + vector.x, origin.y + vector.y}
	arc.End = Point{fmt.Sprintf("%f.4", end.x), fmt.Sprintf("%f.4", end.y)}
}

func (t *Text) ToSExp() string {
	layer := NewSExp("layer", false, t.Layer.Name)
	origin := NewSExp("start", false, t.Origin.ToString())
	return NewSExp("fp_text", false, layer, origin)
}
