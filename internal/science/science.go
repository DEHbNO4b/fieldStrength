package science

import (
	"math"
	"sort"

	"github.com/DEHbNO4b/fieldStrength/internal/data"
)

func MakeResearch(all []*data.Measure) data.Fenomens {

	//создается мапа измерений для отдельных явлений
	fenomens := make(map[string][]float64)
	for _, el := range all {
		fenomens[el.Fenomen()] = append(fenomens[el.Fenomen()], el.Strength())
	}

	//мапа переводится в слайс измерений ( Fenomen)
	fSlices := data.NewFenomens()
	for key, val := range fenomens {
		f := data.NewFenomen(key)
		f.Strength = val
		fSlices = append(fSlices, f)
	}

	for i, el := range fSlices {
		if el.Fenomen != "Гд -град" && el.Fenomen != "СМ -снег мокрый" && el.Fenomen != "МО -метель общая" {
			sort.Float64s(el.Strength) //сортировка измерений
			l := len(el.Strength)
			new := el.Strength[l/5 : 4*l/5] //отсечение выбросов
			fSlices[i].Strength = new
		}

		fSlices[i].Research()

	}
	sort.Sort(fSlices)

	return fSlices
}
func TTest(f data.Fenomens) (data.Fenomens, error) {
	var (
		genMean float64
		genN    int
		genSD   float64
	)
	for _, el := range f {
		if el.Fenomen != "Нет явлений" {
			continue
		}
		genMean = el.Mean
		genN = el.N
		genSD = el.SD
	}
	for i, el := range f {
		f[i].T = (el.Mean - genMean) / math.Sqrt(el.SD*el.SD/float64(el.N)+(genSD*genSD/float64(genN)))
	}
	return f, nil
}
