package bxlparser

// MiltoMM Convert Enlish Mils to milimetres
func MiltoMM(mil float64) float64 {
	return mil * 0.0254
}

// MMtoMil Convert Milimeters	to English Mils
func MMtoMil(mm float64) float64 {
	return mm / 0.0254
}
