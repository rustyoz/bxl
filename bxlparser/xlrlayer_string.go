// Code generated by "stringer -type XlrLayer"; DO NOT EDIT

package bxlparser

import "fmt"

const _XlrLayer_name = "TOP_ASSEMBLYTOP_SILKSCREENTOP_SOLDER_PASTETOP_SOLDER_MASKTOPINNERBOTTOMBOTTOM_SOLDER_MASKBOTTOM_SOLDER_PASTEBOTTOM_SILKSCREENBOTTOM_ASSEMBLYTOP_PLACE_BOUNDBOTTOM_PLACE_BOUNDINTERNAL1INTERNAL2INTERNAL3INTERNAL4INTERNAL5INTERNAL6INTERNAL7INTERNAL8INTERNAL9INTERNAL10INTERNAL11INTERNAL12INTERNAL13INTERNAL14INTERNAL15INTERNAL16USER1USER2USER3USER4USER5USER6USER7USER8USER9USER10L3D_DXFPIN1MARKERPINTESTTOP_BGA_PLACE_BOARDATTRIBUTE4ATTRIBUTE3ATTRIBUTE2ATTRIBUTE1PIN_NUMBERCONSTRAINT_AREACONTACT_AREAINPUTDIMENSIONSROUTE_KEEPOUTVIA_KEEPOUTDRILL_FIGURETOP_COMP_BOUNDBOTTOM_COMP_BOUNDTOP_NOPROBEBOTTOM_NOPROBEPRO_EPIN_DETAILDIMENSIONPAD_DIMENSIONSBOARD"

var _XlrLayer_index = [...]uint16{0, 12, 26, 42, 57, 60, 65, 71, 89, 108, 125, 140, 155, 173, 182, 191, 200, 209, 218, 227, 236, 245, 254, 264, 274, 284, 294, 304, 314, 324, 329, 334, 339, 344, 349, 354, 359, 364, 369, 375, 382, 392, 399, 418, 428, 438, 448, 458, 468, 483, 495, 510, 523, 534, 546, 560, 577, 588, 602, 607, 617, 626, 640, 645}

func (i XlrLayer) String() string {
	if i < 0 || i >= XlrLayer(len(_XlrLayer_index)-1) {
		return fmt.Sprintf("XlrLayer(%d)", i)
	}
	return _XlrLayer_name[_XlrLayer_index[i]:_XlrLayer_index[i+1]]
}

var _XlrLayerNameToValue_map = map[string]XlrLayer{
	_XlrLayer_name[0:12]:    0,
	_XlrLayer_name[12:26]:   1,
	_XlrLayer_name[26:42]:   2,
	_XlrLayer_name[42:57]:   3,
	_XlrLayer_name[57:60]:   4,
	_XlrLayer_name[60:65]:   5,
	_XlrLayer_name[65:71]:   6,
	_XlrLayer_name[71:89]:   7,
	_XlrLayer_name[89:108]:  8,
	_XlrLayer_name[108:125]: 9,
	_XlrLayer_name[125:140]: 10,
	_XlrLayer_name[140:155]: 11,
	_XlrLayer_name[155:173]: 12,
	_XlrLayer_name[173:182]: 13,
	_XlrLayer_name[182:191]: 14,
	_XlrLayer_name[191:200]: 15,
	_XlrLayer_name[200:209]: 16,
	_XlrLayer_name[209:218]: 17,
	_XlrLayer_name[218:227]: 18,
	_XlrLayer_name[227:236]: 19,
	_XlrLayer_name[236:245]: 20,
	_XlrLayer_name[245:254]: 21,
	_XlrLayer_name[254:264]: 22,
	_XlrLayer_name[264:274]: 23,
	_XlrLayer_name[274:284]: 24,
	_XlrLayer_name[284:294]: 25,
	_XlrLayer_name[294:304]: 26,
	_XlrLayer_name[304:314]: 27,
	_XlrLayer_name[314:324]: 28,
	_XlrLayer_name[324:329]: 29,
	_XlrLayer_name[329:334]: 30,
	_XlrLayer_name[334:339]: 31,
	_XlrLayer_name[339:344]: 32,
	_XlrLayer_name[344:349]: 33,
	_XlrLayer_name[349:354]: 34,
	_XlrLayer_name[354:359]: 35,
	_XlrLayer_name[359:364]: 36,
	_XlrLayer_name[364:369]: 37,
	_XlrLayer_name[369:375]: 38,
	_XlrLayer_name[375:382]: 39,
	_XlrLayer_name[382:392]: 40,
	_XlrLayer_name[392:399]: 41,
	_XlrLayer_name[399:418]: 42,
	_XlrLayer_name[418:428]: 43,
	_XlrLayer_name[428:438]: 44,
	_XlrLayer_name[438:448]: 45,
	_XlrLayer_name[448:458]: 46,
	_XlrLayer_name[458:468]: 47,
	_XlrLayer_name[468:483]: 48,
	_XlrLayer_name[483:495]: 49,
	_XlrLayer_name[495:510]: 50,
	_XlrLayer_name[510:523]: 51,
	_XlrLayer_name[523:534]: 52,
	_XlrLayer_name[534:546]: 53,
	_XlrLayer_name[546:560]: 54,
	_XlrLayer_name[560:577]: 55,
	_XlrLayer_name[577:588]: 56,
	_XlrLayer_name[588:602]: 57,
	_XlrLayer_name[602:607]: 58,
	_XlrLayer_name[607:617]: 59,
	_XlrLayer_name[617:626]: 60,
	_XlrLayer_name[626:640]: 61,
	_XlrLayer_name[640:645]: 62,
}

func XlrLayerString(s string) (XlrLayer, error) {
	if val, ok := _XlrLayerNameToValue_map[s]; ok {
		return val, nil
	}
	return 0, fmt.Errorf("%s does not belong to XlrLayer values", s)
}
