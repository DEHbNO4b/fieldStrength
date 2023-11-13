package main

import (
	"fmt"
	"strconv"

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

	//get all fenomens from file
	allFens, err := reader.ReadFenomens(files[0])
	if err != nil {
		logger.Log.Error(err.Error())
		return nil
	}
	fmt.Printf("all fenomens len: %d\n", len(allFens))

	fenomens := make(map[string][]float64)
	for _, el := range allFens {
		fenomens[el.Fenomen()] = append(fenomens[el.Fenomen()], el.Strength())
	}

	for key, val := range fenomens {
		fmt.Printf("%s  len:  %d \n", strconv.Quote(key), len(val))
	}
	// //get specified fenomens
	// withoutFens := make([]float64, 0, 300000)
	// rosa := make([]float64, 0, 1000)
	// inei := make([]float64, 0, 1000)
	// moros := make([]float64, 0, 1000)
	// gololed := make([]float64, 0, 1000)

	// for _, el := range allFens {
	// 	switch el.Fenomen() {
	// 	case "Нет явлений":
	// 		withoutFens = append(withoutFens, el.Strength())
	// 	case "Мр -морось":
	// 		moros = append(moros, el.Strength())
	// 	case "Р  -роса":
	// 		rosa = append(rosa, el.Strength())
	// 	case "И  -иней":
	// 		inei = append(inei, el.Strength())
	// 	case "Гл -гололед":
	// 		gololed = append(gololed, el.Strength())
	// 	}
	// }
	// fmt.Printf("without fenomens len: %d\n", len(withoutFens))
	// sort.Float64s(withoutFens)
	// l := len(withoutFens)
	// new := withoutFens[l/6 : 5*l/6]
	// fmt.Printf("new slice without fenomens len : %d\n", len(new))
	// fmt.Printf("moros len : %d\n", len(moros))
	// fmt.Printf("rosa len : %d\n", len(rosa))
	// fmt.Printf("inei len : %d\n", len(inei))
	// fmt.Printf("gololed len : %d\n", len(gololed))

	return nil
}
