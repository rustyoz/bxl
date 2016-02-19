package bxlparser

import (
	"log"
	"regexp"
	"strings"
	"unicode"
)

type Polygon struct {
	Layer  XlrLayer
	Origin Point
	End    Point
	Points []Point
	Width  string
}

type HasPolygon interface {
	AddPolygon(p Polygon)
	Data() *[]string
}

func FindPolygon(hp HasPolygon) {
	f := func(c rune) bool {
		return !unicode.IsLetter(c) && !unicode.IsNumber(c) && c != '-' && c != '.' && c != '_'
	}

	for _, l := range *hp.Data() {
		if strings.HasPrefix(strings.TrimSpace(l), "Poly") {
			var poly Polygon
			fields := strings.FieldsFunc(l, SplitFields)
			for j, f := range fields {
				switch f {
				case "Layer":
					poly.Layer, _ = XlrLayerString(fields[j+1])
				case "Origin":
					poly.Origin.FromString(fields[j+1], fields[j+2])
				case "EndPoint":
					poly.End.FromString(fields[j+1], fields[j+2])
				case "Width":
					poly.Width = fields[j+1]

				}
			}
			tuples := FindTuples(l)
			for _, tup := range tuples {
				cords := strings.FieldsFunc(tup, f)
				var p Point
				p.FromString(cords[0], cords[1])
				poly.Points = append(poly.Points, p)
			}
			hp.AddPolygon(poly)
		}
	}
}

func FindTuples(input string) []string {
	reg, err := regexp.Compile(`\([^a-zA-Z\)]*\)`)
	if err != nil {
		log.Fatal(err)
	}
	return reg.FindAllString(input, -1)

}
