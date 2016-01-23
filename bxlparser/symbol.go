package bxlparser

type Symbol struct {
	Name  string
	data  []string
	Lines []Line
}

func (s *Symbol) AddLine(l Line) {
	s.Lines = append(s.Lines, l)
}

func (s *Symbol) Data() *[]string {
	return &s.data
}
