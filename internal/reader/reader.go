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
	filenameTemplate = `C:\Users\guzoe\OneDrive\Рабочий стол\go\fieldStrength\public\*.csv`
	fileDir          = `C:\Users\guzoe\OneDrive\Рабочий стол\go\fieldStrength\public`
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

// func ReadEnstrokes(files []string) []data.Stroke {
// 	strokes := make([]data.Stroke, 0, 1000)
// 	for _, f := range files {
// 		file, err := os.Open(f)
// 		if err != nil {
// 			fmt.Println(err)
// 			continue
// 		}
// 		defer file.Close()
// 		scanner := bufio.NewScanner(file)
// 		for scanner.Scan() {
// 			line := scanner.Text()
// 			stroke, err := parseEnStroke(line)
// 			if err != nil {
// 				continue
// 			}
// 			strokes = append(strokes, stroke)
// 		}
// 	}

// 	return strokes
// }

// func parseEnStroke(record string) (data.Stroke, error) {

// 	rec := strings.Split(record, ";")
// 	if len(rec) != 8 {
// 		return data.Stroke{}, fmt.Errorf("invalid record: %v", rec)
// 	}
// 	//определение типа разряда
// 	var cloud bool
// 	if len(rec[0]) == 12 {
// 		cloud = false
// 	} else if len(rec[0]) == 13 {
// 		cloud = true
// 	} else {
// 		return data.Stroke{}, fmt.Errorf("invalid record: %v", rec[0])
// 	}
// 	//парсинг времени
// 	layout := "2006-01-02 15:04:05"
// 	time, _ := time.Parse(layout, strings.Trim(rec[1], "\""))

// 	nano, err := strconv.Atoi(rec[2])
// 	if err != nil {
// 		logger.Log.Error("unable to parse nanoseconds from file", zap.Error(err))
// 	}
// 	lat, err := strconv.ParseFloat(rec[3], 64)
// 	if err != nil {
// 		logger.Log.Error("unable to parse latitude from file", zap.Error(err))
// 	}
// 	long, err := strconv.ParseFloat(rec[4], 64)
// 	if err != nil {
// 		logger.Log.Error("unable to parse longitude from file", zap.Error(err))
// 	}
// 	c, _ := strconv.Atoi(rec[5])

// 	signal := c / 1000
// 	height, err := strconv.Atoi(rec[6])
// 	if err != nil {
// 		logger.Log.Error("unable to parse height from file", zap.Error(err))
// 	}
// 	sensors, err := strconv.Atoi(rec[7])
// 	if err != nil {
// 		logger.Log.Error("unable to parse sensors from file", zap.Error(err))
// 	}
// 	return data.Stroke{Cloud: cloud, Time: time, Nano: nano,
// 		Lat: lat, Long: long, Signal: signal, Height: height, Sensors: sensors}, nil

// }
