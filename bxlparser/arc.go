package bxlparser

import (
	"math"
	"strconv"
	"strings"

	"github.com/rustyoz/gokicadlib"
)

type Arc struct {
	Layer      XlrLayer
	Origin     Point
	Radius     float64
	StartAngle float64
	SweepAngle float64
	Width      float64
	End        Point
}

type ArcSlice []Arc

func (arc *Arc) CalculateEndPoint() {
	vector := Point{arc.Radius * math.Cos(arc.StartAngle*180/math.Pi), arc.Radius * math.Sin(arc.StartAngle*180/math.Pi)}
	arc.End = Point{arc.Origin.X + vector.X, arc.Origin.Y + vector.Y}
}

type HasArcs interface {
	AddArc(l Arc)
	Data() *[]string
}

func FindArcs(harcs HasArcs) {

	for _, l := range *harcs.Data() {
		if strings.HasPrefix(strings.TrimSpace(l), "Arc") {
			var arc Arc
			fields := strings.FieldsFunc(l, SplitFields)
			for j, f := range fields {
				switch f {
				case "Layer":
					arc.Layer, _ = XlrLayerString(fields[j+1])
				case "Origin":
					arc.Origin.FromString(fields[j+1], fields[j+2])
				case "StartAngle":
					arc.StartAngle, _ = strconv.ParseFloat(fields[j+1], 64)
				case "SweepAngle":
					arc.SweepAngle, _ = strconv.ParseFloat(fields[j+1], 64)
				case "Radius":
					arc.Radius, _ = strconv.ParseFloat(fields[j+1], 64)
				case "Width":
					arc.Width, _ = strconv.ParseFloat(fields[j+1], 64)
				}
			}
			harcs.AddArc(arc)
		}
	}
}

func (a *Arc) ToKicadArc() (ka *gokicadlib.Arc) {
	ka = &gokicadlib.Arc{}
	ka.Angle = a.SweepAngle
	ka.Start = gokicadlib.Point{MiltoMM(a.Origin.X), MiltoMM(-a.Origin.Y)}

	var err error
	ka.Layer, err = a.Layer.ToKicadLayer()
	if err != nil {
		return nil
	}
	v := Vector(MiltoMM(a.Radius), a.StartAngle+a.SweepAngle)

	ka.End = gokicadlib.Point{MiltoMM(a.Origin.X) + v.X, MiltoMM(-a.Origin.Y) + v.Y}
	ka.Width = MiltoMM(a.Width)

	return ka
}

func Vector(radius float64, angle float64) Point {
	vector := Point{radius * math.Cos(angle*math.Pi/180.0), radius * math.Sin(angle*math.Pi/180.0)}
	return vector
}

func (as ArcSlice) ToKicadArcs() []gokicadlib.Arc {
	var kcas []gokicadlib.Arc
	for _, a := range as {
		arc := a.ToKicadArc()
		if arc != nil {
			kcas = append(kcas, *arc)
		}
	}
	return kcas
}
