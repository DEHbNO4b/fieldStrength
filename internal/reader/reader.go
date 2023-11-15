package reader

import (
	"encoding/csv"
	"fmt"
	"io"
	"io/fs"
	"os"
	"path/filepath"

	"github.com/DEHbNO4b/fieldStrength/internal/data"
	"github.com/DEHbNO4b/fieldStrength/internal/logger"
	"go.uber.org/zap"
)

var (
	filenameTemplate = `C:\Users\user\Desktop\go\fieldStrength\public\*.csv`
	fileDir          = `C:\Users\user\Desktop\go\fieldStrength\public`
)

func SearchNewEnFiles() ([]string, error) {

	var files []string
	err := filepath.WalkDir(fileDir, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if !d.IsDir() {
			fmt.Println(path)
			matched, err := filepath.Match(filepath.FromSlash(filenameTemplate), path)
			if err != nil {
				return err
			}
			if matched {
				//читаем файлы совпавшие с заданной строкой в требуемой директории
				files = append(files, path)
			}
		}
		return nil
	})

	return files, err
}
func ReadFenomens(f string) ([]*data.Measure, error) {
	logger.Log.Info("", zap.String("reading file:", f))
	file, err := os.Open(f)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	//create new csv reader
	cr := csv.NewReader(file)
	cr.Comma = ';'

	fields := make([]*data.Measure, 0, 100000)

	for {
		record, err := cr.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			logger.Log.Info(err.Error())
			continue
		}
		fenomen, err := MakeFenomen(record)
		if err != nil {
			logger.Log.Error(err.Error())
			continue
		}
		df, err := ToDomainFenomen(fenomen)
		if err != nil {
			logger.Log.Error(err.Error())
			continue
		}
		fields = append(fields, df)
	}
	return fields, nil
}
