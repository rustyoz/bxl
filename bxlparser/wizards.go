package bxlparser

import (
	"fmt"
	"strings"
)

type Wizard struct {
	VarName string
	VarData string
	Number  int
	Origin  Point
}
type WizardSlice []Wizard

type HasWizards interface {
	AddWizard(w Wizard)
	Data() *[]string
}

func FindWizards(hw HasWizards) {

	for _, l := range *hw.Data() {
		if strings.HasPrefix(strings.TrimSpace(l), "Wizard") {
			var w Wizard
			fields := strings.FieldsFunc(l, feildfuncer())
			for j, _ := range fields {
				if j == 0 {
					continue
				}
				switch fields[j-1] {
				case "Origin":
					w.Origin.FromString(fields[j], fields[j+1])
				case "VarName":
					w.VarName = fields[j]
				case "VarData":
					w.VarData = fields[j]
				}
			}
			hw.AddWizard(w)
		}
	}
}

func (ws WizardSlice) Contains(name string) (*Wizard, error) {
	for _, w := range ws {
		if w.VarName == name {
			return &w, nil
		}
	}
	return nil, fmt.Errorf("Wizard not found: %v", name)
}
