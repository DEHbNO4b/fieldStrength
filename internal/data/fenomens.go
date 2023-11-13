package data

import (
	"unicode/utf8"

	"github.com/montanaflynn/stats"
)

type Fenomen struct {
	Fenomen  string
	Strength []float64
	Stats
}
type Stats struct {
	Median float64
	Max    float64
	Min    float64
	SD     float64
}

func NewFenomen(f string) Fenomen {
	size := utf8.RuneCountInString(f)
	for i := size - 1; i < 30; i++ {
		f = f + " "
	}
	return Fenomen{Fenomen: string(f)}
}
func (f *Fenomen) Research() error {
	m, err := stats.Median(f.Strength)
	if err != nil {
		return err
	}
	f.Median = m
	m, err = stats.Max(f.Strength)
	if err != nil {
		return err
	}
	f.Max = m
	m, err = stats.Min(f.Strength)
	if err != nil {
		return err
	}
	f.Min = m
	sd, err := stats.StandardDeviation(f.Strength)
	if err != nil {
		return err
	}
	f.SD = sd
	return nil
}

type Fenomens []Fenomen

func NewFenomens() Fenomens {
	f := make([]Fenomen, 0, 10)
	return f
}

func (f Fenomens) Len() int {
	return len(f)
}
func (f Fenomens) Less(i, j int) bool {
	// return len(f[i].Strength) > len(f[j].Strength)
	return f[i].SD > f[j].SD
}
func (f Fenomens) Swap(i, j int) {
	f[i], f[j] = f[j], f[i]
}

// type SdSort []Fenomen
// func (s SdSort) Len() int {
// 	return len(s)
// }
// func (s SdSort) Less(i, j int) bool {
// 	return len(f[i].Strength) > len(f[j].Strength)
// 	// return len(f[i].Strength) > len(f[j].Strength)
// }
// func (s SdSort) Swap(i, j int) {
// 	f[i], f[j] = f[j], f[i]
// }
