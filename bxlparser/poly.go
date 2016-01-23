package bxlparser

import (
	"log"
	"regexp"
	"strings"
	"unicode"
)

type Polygon struct {
	Layer  Layer
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
		return !unicode.IsLetter(c) && !unicode.IsNumber(c) && c != '-' && c != '.'
	}

	for _, l := range *hp.Data() {
		if strings.HasPrefix(strings.TrimSpace(l), "Poly") {
			var poly Polygon
			fields := strings.FieldsFunc(l, f)
			for j, f := range fields {
				switch f {
				case "Layer":
					poly.Layer = Layer{fields[j+1]}
				case "Origin":
					poly.Origin = Point{fields[j+1], fields[j+2]}
				case "EndPoint":
					poly.End = Point{fields[j+1], fields[j+2]}
				case "Width":
					poly.Width = fields[j+1]

				}
			}
			tuples := FindTuples(l)
			for _, tup := range tuples {
				cords := strings.FieldsFunc(tup, f)
				poly.Points = append(poly.Points, Point{cords[0], cords[1]})
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
