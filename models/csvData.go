package models

type CSVData struct {
	Rows []Row
}

func NewCSVData(data [][]string) *CSVData {
	csvData := CSVData{}
	for i := range data {
		row := Row{}
		row.AddColum(data[i]...)
		csvData.Rows = append(csvData.Rows, row)
	}
	return &csvData
}

func (data *CSVData) GetRowByIndex(i int) (row Row) {
	if i > len(data.Rows)+1 {
		return Row{}
	}
	return (data.Rows)[i]
}

func (data *CSVData) GetRowValuesAsMap(row int) map[string]int {
	if row > len(data.Rows)+1 {
		return nil
	}

	return data.Rows[row].GetColumnValueAsMap()
}

func (data *CSVData) FindValue(key string) *Row {
	length := len(data.Rows) - 1
	for i := range data.Rows {
		if i+1 > length-i {
			return nil
		} else if value := data.Rows[i].FindValue(key); value != "" {
			return &data.Rows[i]
		} else if value := data.Rows[length-i].FindValue(key); value != "" {
			return &data.Rows[i]
		}
	}
	return nil
}

func (data *CSVData) FindValueInColumn(i int, keys ...string) (rows []*Row, notFound *[]string) {
	search := NewSearch()
	search.AddKeys(keys...)

	length := len(data.Rows) - 1
	for j := range data.Rows {
		if j+1 > length-j {
			break
		} else if data.Rows[j].FindValueInColumn(i, search) != "" {
			rows = append(rows, &data.Rows[j])
		} else if data.Rows[length-j].FindValueInColumn(i, search) != "" {
			rows = append(rows, &data.Rows[length-j])
		}
		if search.AllFound() {
			return
		}
	}
	return rows, search.GetNotFound()
}
