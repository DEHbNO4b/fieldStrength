package science

import (
	"sort"

	"github.com/DEHbNO4b/fieldStrength/internal/data"
)

func MakeResearch(all []*data.Measure) data.Fenomens {
	fenomens := make(map[string][]float64)
	for _, el := range all {
		fenomens[el.Fenomen()] = append(fenomens[el.Fenomen()], el.Strength())
	}

	fSlices := data.NewFenomens()
	for key, val := range fenomens {
		f := data.NewFenomen(key)
		f.Strength = val
		fSlices = append(fSlices, f)
	}

	for i, el := range fSlices {
		sort.Float64s(el.Strength)
		l := len(el.Strength)
		new := el.Strength[l/5 : 4*l/5]
		fSlices[i].Strength = new
		fSlices[i].Research()

	}
	sort.Sort(fSlices)

	return fSlices
}
