package main

import (
	"fmt"
	"sort"
	"strconv"

	"github.com/DEHbNO4b/fieldStrength/internal/data"
	"github.com/DEHbNO4b/fieldStrength/internal/logger"
	"github.com/DEHbNO4b/fieldStrength/internal/reader"
)

func main() {
	err := run()
	if err != nil {
		panic(err)
	}
}

func run() error {
	//logger
	if err := logger.Initialize("info"); err != nil {
		return err
	}

	//search *.csv files
	files, err := reader.SearchNewEnFiles()
	if err != nil {
		logger.Log.Error(err.Error())
		return nil
	}
	if len(files) == 0 {
		logger.Log.Info("no files in public dir")
	}

	//get all measurements from file
	all, err := reader.ReadFenomens(files[0])
	if err != nil {
		logger.Log.Error(err.Error())
		return nil
	}
	fmt.Printf("all measurements len: %d\n", len(all))

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
	for _, el := range fSlices {
		fmt.Printf("slice: %s -len: 	%d;	stats = %+v \n", strconv.Quote(el.Fenomen), len(el.Strength), el.Stats)
	}

	// new := withoutFens[l/6 : 5*l/6]

	return nil
}
