package data

import (
	"strconv"
	"strings"

	"github.com/montanaflynn/stats"
)

var GenMean float64

type Fenomen struct {
	Fenomen  string
	Strength []float64
	Stats
}
type Stats struct {
	Mean   float64
	Median float64
	Max    float64
	Min    float64
	SD     float64
	N      int
	T      float64
}

func NewFenomen(f string) Fenomen {
	// size := utf8.RuneCountInString(f)
	// for i := size - 1; i < 30; i++ {
	// 	f = f + " "
	// }
	return Fenomen{Fenomen: f}
}

func (f *Fenomen) Research() error {
	f.N = len(f.Strength)
	m, err := stats.Median(f.Strength)
	if err != nil {
		return err
	}
	f.Median = m
	m, err = stats.Mean(f.Strength)
	if err != nil {
		return err
	}
	f.Mean = m
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
func (f Fenomen) Slice() []string {
	n := strconv.Itoa(f.N)
	mean := strconv.FormatFloat(f.Mean, 'f', -1, 64)
	median := strconv.FormatFloat(f.Median, 'f', -1, 64)
	max := strconv.FormatFloat(f.Max, 'f', -1, 64)
	min := strconv.FormatFloat(f.Min, 'f', -1, 64)
	sd := strconv.FormatFloat(f.SD, 'f', -1, 64)
	t := strconv.FormatFloat(f.T, 'f', -1, 64)

	ans := []string{
		f.Fenomen,
		n,
		mean,
		median,
		max,
		min,
		sd,
		t,
	}
	for i, el := range ans {
		ans[i] = strings.Replace(el, ".", ",", 1)
	}
	return ans
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
	return f[i].SD < f[j].SD
}
func (f Fenomens) Swap(i, j int) {
	f[i], f[j] = f[j], f[i]
}
