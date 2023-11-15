package data

import (
	"fmt"
	"strconv"
	"strings"
)

type Measure struct {
	fenomen  string
	strength float64
}

func NewMeasure(f string, s float64) (*Measure, error) {
	if f == "" {
		return nil, fmt.Errorf("empty data in field fenomen: %s %f", f, s)
	}
	return &Measure{fenomen: f, strength: s}, nil
}

func (f *Measure) Fenomen() string {
	return f.fenomen
}
func (f *Measure) Strength() float64 {
	return f.strength
}
func (f *Measure) SetFenomen(fen string) {
	f.fenomen = fen
}
func (f *Measure) SetStrength(s float64) {
	f.strength = s
}

func (f *Measure) Slice() []string {
	s := strconv.FormatFloat(f.strength, 'f', -1, 64)
	s = strings.Replace(s, ".", ",", 1)
	return []string{f.fenomen, s}
}
