package main

import (
	"fmt"
	"strconv"

	"github.com/DEHbNO4b/fieldStrength/internal/logger"
	"github.com/DEHbNO4b/fieldStrength/internal/reader"
	"github.com/DEHbNO4b/fieldStrength/internal/science"
	"github.com/DEHbNO4b/fieldStrength/internal/writer"
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

	//search *.csv files in public dir
	files, err := reader.SearchNewEnFiles()
	if err != nil {
		logger.Log.Error(err.Error())
		return nil
	}
	if len(files) == 0 {
		logger.Log.Info("no files in public dir")
		return nil
	}

	//get all measurements from first file
	all, err := reader.ReadFenomens(files[0])
	if err != nil {
		logger.Log.Error(err.Error())
		return nil
	}
	fmt.Printf("all measurements len: %d\n", len(all))

	f := science.MakeResearch(all)

	f, err = science.TTest(f)
	if err != nil {
		logger.Log.Error(err.Error())
	}
	for _, el := range f {
		fmt.Printf("%s:	len:%d;			stats = %+v \n", strconv.Quote(el.Fenomen), len(el.Strength), el.Stats)
	}
	//writer.WriteOutCsv(f)
	err = writer.OutStats(f)
	if err != nil {
		logger.Log.Error(err.Error())
	}

	return nil
}
