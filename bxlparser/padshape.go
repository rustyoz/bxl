package bxlparser

import (
	"log"
	"strconv"
	"strings"
)

type PadShape struct {
	Description string
	Width       float64
	Height      float64
	PadType     int
	Layer       XlrLayer
}

func (ps *PadShape) parsePadShape(l string) {
	fields := strings.FieldsFunc(l, feildfuncer())

	for j, f := range fields {
		var err error
		switch f {
		case "PadShape":
			ps.Description = fields[j+1]
		case "Layer":
			ps.Layer, err = XlrLayerString(fields[j+1])
		case "Width":
			ps.Width, err = strconv.ParseFloat(fields[j+1], 64)
		case "Height":
			ps.Height, err = strconv.ParseFloat(fields[j+1], 64)
		case "PadType":
			var i int64
			i, err = strconv.ParseInt(fields[j+1], 10, 32)
			ps.PadType = int(i)
		}
		if err != nil {
			log.Fatal(err)
		}
	}
}
