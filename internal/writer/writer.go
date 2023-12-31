package writer

import (
	"encoding/csv"
	"os"
	"path/filepath"

	"github.com/DEHbNO4b/fieldStrength/internal/data"
	"github.com/DEHbNO4b/fieldStrength/internal/logger"
)

func WriteOutCsv(fSlices data.Fenomens) {
	measurements := make([]*data.Measure, 0, 100000)
	for _, el := range fSlices {
		for _, val := range el.Strength {
			m, err := data.NewMeasure(el.Fenomen, val)
			if err != nil {
				continue
			}
			measurements = append(measurements, m)
		}
	}
	newFile, err := os.Create(filepath.FromSlash("./public/out.csv"))
	if err != nil {
		logger.Log.Error(err.Error())
	}
	w := csv.NewWriter(newFile)
	w.Comma = ';'
	for _, el := range measurements {
		err := w.Write(el.Slice())
		if err != nil {
			logger.Log.Error(err.Error())
			continue
		}
	}
	w.Flush()
	if err := w.Error(); err != nil {
		logger.Log.Error(err.Error())
	}
}
func OutStats(f data.Fenomens) error {
	file, err := os.Create(filepath.FromSlash("./public/stats.csv"))
	if err != nil {
		logger.Log.Error(err.Error())
		return err
	}
	w := csv.NewWriter(file)
	w.Comma = ';'
	err = w.Write([]string{"name", "N", "mean", "median", "max", "min", "sd", "t"})
	if err != nil {
		logger.Log.Error(err.Error())
		return err
	}
	for _, el := range f {
		err := w.Write(el.Slice())
		if err != nil {
			logger.Log.Error(err.Error())
			return err
		}
	}
	w.Flush()
	if err := w.Error(); err != nil {
		logger.Log.Error(err.Error())
		return err
	}
	return nil
}
