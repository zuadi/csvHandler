package csvHandler

import (
	"encoding/csv"
	"os"

	"github.com/zuadi/csvHandler/models"
)

type CSVHandler struct {
	delimiter rune
	Data      *models.CSVData
}

func NewCSVHandler(delimiter rune) *CSVHandler {
	return &CSVHandler{delimiter: delimiter}
}

func (cH *CSVHandler) LoadAllData(path string) (err error) {
	f, err := os.Open(path)
	if err != nil {
		return err
	}
	defer f.Close()

	reader := csv.NewReader(f)
	reader.Comma = cH.delimiter
	data, err := reader.ReadAll()
	if err != nil {
		return err
	}
	cH.Data = models.NewCSVData(data)
	return
}

func (cH *CSVHandler) GetRowValuesAsMap(row int) map[string]int {
	return cH.Data.GetRowValuesAsMap(row)
}

func (cH *CSVHandler) FindValue(key string) (column *models.Row) {
	return cH.Data.FindValue(key)
}

func (cH *CSVHandler) FindValueInColumn(i int, keys ...string) (column []*models.Row, notFound *[]string) {
	return cH.Data.FindValueInColumn(i, keys...)
}
