package reader

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/DEHbNO4b/fieldStrength/internal/data"
)

type Fenomen struct {
	Fenomen  string  `json:"fenomen"`
	Strength float64 `json:"strength"`
}

func MakeFenomen(record []string) (Fenomen, error) {
	if len(record) != 2 {
		return Fenomen{}, fmt.Errorf("invalid record slice: %v", record)
	}
	fen := record[0]
	strength := strings.ReplaceAll(record[1], ",", ".")
	s, err := strconv.ParseFloat(strength, 64)
	if err != nil {
		return Fenomen{}, fmt.Errorf("invalid strength field in record: %v", record)
	}
	return Fenomen{Fenomen: fen, Strength: s}, nil
}

func ToDomainFenomen(f Fenomen) (*data.Measure, error) {
	return data.NewStrength(f.Fenomen, f.Strength)
}
