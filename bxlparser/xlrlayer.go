package bxlparser

import (
	"fmt"
	"log"

	"github.com/rustyoz/gokicadlib"
)

type XlrLayer int

var xlrLayerFilter = map[XlrLayer]bool{
	TOP_ASSEMBLY:        true,
	TOP_SILKSCREEN:      true,
	TOP_SOLDER_PASTE:    true,
	TOP_SOLDER_MASK:     true,
	TOP:                 true,
	INNER:               true,
	BOTTOM:              true,
	BOTTOM_SOLDER_MASK:  true,
	BOTTOM_SOLDER_PASTE: true,
	BOTTOM_SILKSCREEN:   true,
	BOTTOM_ASSEMBLY:     true,
	TOP_PLACE_BOUND:     true,
	BOTTOM_PLACE_BOUND:  true,
	INTERNAL1:           true,
	INTERNAL2:           true,
	INTERNAL3:           true,
	INTERNAL4:           true,
	INTERNAL5:           true,
	INTERNAL6:           true,
	INTERNAL7:           true,
	INTERNAL8:           true,
	INTERNAL9:           true,
	INTERNAL10:          true,
	INTERNAL11:          true,
	INTERNAL12:          true,
	INTERNAL13:          true,
	INTERNAL14:          true,
	INTERNAL15:          true,
	INTERNAL16:          true,
	USER1:               true,
	USER2:               true,
	USER3:               true,
	USER4:               true,
	USER5:               true,
	USER6:               true,
	USER7:               true,
	USER8:               true,
	USER9:               true,
	USER10:              true,
	L3D_DXF:             true,
	PIN1MARKER:          true,
	PINTEST:             true,
	TOP_BGA_PLACE_BOARD: true,
	ATTRIBUTE4:          true,
	ATTRIBUTE3:          true,
	ATTRIBUTE2:          true,
	ATTRIBUTE1:          true,
	PIN_NUMBER:          true,
	CONSTRAINT_AREA:     true,
	CONTACT_AREA:        true,
	INPUTDIMENSIONS:     true,
	ROUTE_KEEPOUT:       true,
	VIA_KEEPOUT:         true,
	DRILL_FIGURE:        true,
	TOP_COMP_BOUND:      true,
	BOTTOM_COMP_BOUND:   true,
	TOP_NOPROBE:         true,
	BOTTOM_NOPROBE:      true,
	PRO_E:               false,
	PIN_DETAIL:          false,
	DIMENSION:           false,
	PAD_DIMENSIONS:      false,
	BOARD:               true,
}

const (
	TOP_ASSEMBLY XlrLayer = iota
	TOP_SILKSCREEN
	TOP_SOLDER_PASTE
	TOP_SOLDER_MASK
	TOP
	INNER
	BOTTOM
	BOTTOM_SOLDER_MASK
	BOTTOM_SOLDER_PASTE
	BOTTOM_SILKSCREEN
	BOTTOM_ASSEMBLY
	TOP_PLACE_BOUND
	BOTTOM_PLACE_BOUND
	INTERNAL1
	INTERNAL2
	INTERNAL3
	INTERNAL4
	INTERNAL5
	INTERNAL6
	INTERNAL7
	INTERNAL8
	INTERNAL9
	INTERNAL10
	INTERNAL11
	INTERNAL12
	INTERNAL13
	INTERNAL14
	INTERNAL15
	INTERNAL16
	USER1
	USER2
	USER3
	USER4
	USER5
	USER6
	USER7
	USER8
	USER9
	USER10
	L3D_DXF
	PIN1MARKER
	PINTEST
	TOP_BGA_PLACE_BOARD
	ATTRIBUTE4
	ATTRIBUTE3
	ATTRIBUTE2
	ATTRIBUTE1
	PIN_NUMBER
	CONSTRAINT_AREA
	CONTACT_AREA
	INPUTDIMENSIONS
	ROUTE_KEEPOUT
	VIA_KEEPOUT
	DRILL_FIGURE
	TOP_COMP_BOUND
	BOTTOM_COMP_BOUND
	TOP_NOPROBE
	BOTTOM_NOPROBE
	PRO_E
	PIN_DETAIL
	DIMENSION
	PAD_DIMENSIONS
	BOARD
)

var layermap = map[string]gokicadlib.Layer{
	"TOP_ASSEMBLY":        "F.Fab",
	"TOP_SILKSCREEN":      "F.SilkS",
	"TOP_SOLDER_PASTE":    "F.Paste",
	"TOP_SOLDER_MASK":     "F.Mask",
	"TOP":                 "F.Cu",
	"INNER":               "In1.Cu",
	"BOTTOM":              "B.Cu",
	"BOTTOM_SOLDER_MASK":  "B.Mask",
	"BOTTOM_SOLDER_PASTE": "B.Paste",
	"BOTTOM_SILKSCREEN":   "B.SilkS",
	"BOTTOM_ASSEMBLY":     "B.Fab",
	"TOP_PLACE_BOUND":     "F.CrtYd",
	"BOTTOM_PLACE_BOUND":  "B.CrtYd",
	"INTERNAL1":           "In1.Cu",
	"INTERNAL2":           "In2.Cu",
	"INTERNAL3":           "In3.Cu",
	"INTERNAL4":           "In4.Cu",
	"INTERNAL5":           "In5.Cu",
	"INTERNAL6":           "In6.Cu",
	"INTERNAL7":           "In7.Cu",
	"INTERNAL8":           "In8.Cu",
	"INTERNAL9":           "In9.Cu",
	"INTERNAL10":          "In10.Cu",
	"INTERNAL11":          "In11.Cu",
	"INTERNAL12":          "In12.Cu",
	"INTERNAL13":          "In13.Cu",
	"INTERNAL14":          "In14.Cu",
	"INTERNAL15":          "In15.Cu",
	"INTERNAL16":          "In16.Cu",
	"USER1":               "Eco1.User",
	"USER2":               "Eco2.User",
	"DIMENSION":           "Cmts.User",
	"PAD_DIMENSIONS":      "Cmts.User",
	"PRO_E":               "Cmts.User",
	"BOARD":               "Edge.Cuts",
	"INPUTDIMENSIONS":     "Cmts.User",
	"PIN_DETAIL":          "Cmts.User",
}

func (l XlrLayer) ToKicadLayer() (gokicadlib.Layer, error) {

	kcl, err := layermap[l.String()]
	if err != true {
		log.Fatal(l.String(), err)
	}
	if xlrlayerFilter(l) {
		return kcl, nil
	}
	e := fmt.Errorf("excluded layer %v", l)
	return gokicadlib.Exclude, e
}

func xlrlayerFilter(l XlrLayer) bool {
	return xlrLayerFilter[l]
}
