package data

import "fmt"

type Measure struct {
	fenomen  string
	strength float64
}

func NewStrength(f string, s float64) (*Measure, error) {
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
