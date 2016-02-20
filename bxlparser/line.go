package bxlparser

import (
	"strconv"
	"strings"

	"github.com/rustyoz/gokicadlib"
)

type Line struct {
	Origin Point
	End    Point
	Layer  XlrLayer
	Width  float64
}

type LineSlice []Line

type HasLines interface {
	AddLine(l Line)
	Data() *[]string
}

func FindLines(hl HasLines) {

	for _, l := range *hl.Data() {
		if strings.HasPrefix(strings.TrimSpace(l), "Line") {
			var line Line
			fields := strings.FieldsFunc(l, SplitFields)
			for j, f := range fields {
				switch f {
				case "Layer":
					line.Layer, _ = XlrLayerString(fields[j+1])
				case "Origin":
					line.Origin.FromString(fields[j+1], fields[j+2])
				case "EndPoint":
					line.End.FromString(fields[j+1], fields[j+2])
				case "Width":
					line.Width, _ = strconv.ParseFloat(fields[j+1], 64)
				}
			}
			hl.AddLine(line)
		}
	}
}

func (l Line) ToKicadLine() gokicadlib.Line {
	var kcl gokicadlib.Line

	kcl.Origin.X = MiltoMM(l.Origin.X)
	kcl.Origin.Y = MiltoMM(-l.Origin.Y)

	kcl.End.X = MiltoMM(l.End.X)
	kcl.End.Y = MiltoMM(-l.End.Y)
	kcl.Layer = l.Layer.ToKicadLayer()
	kcl.Width = MiltoMM(l.Width)
	return kcl
}

func (ls LineSlice) ToKicadLines() []gokicadlib.Line {
	var kcls []gokicadlib.Line
	for _, l := range ls {
		kcls = append(kcls, l.ToKicadLine())
	}
	return kcls
}
