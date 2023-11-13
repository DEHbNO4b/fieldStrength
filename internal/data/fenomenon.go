package data

import "fmt"

type Fenomen struct {
	fenomen  string
	strength float64
}

func NewFenomen(f string, s float64) (*Fenomen, error) {
	if f == "" {
		return nil, fmt.Errorf("empty data in field fenomen: %s %f", f, s)
	}
	return &Fenomen{fenomen: f, strength: s}, nil
}
func (f *Fenomen) Fenomen() string {
	return f.fenomen
}
func (f *Fenomen) Strength() float64 {
	return f.strength
}
func (f *Fenomen) SetFenomen(fen string) {
	f.fenomen = fen
}
func (f *Fenomen) SetStrength(s float64) {
	f.strength = s
}
