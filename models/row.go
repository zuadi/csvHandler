package models

type Row struct {
	Columns []string
}

func (r *Row) AddColum(data ...string) {
	r.Columns = append(r.Columns, data...)
}

func (r *Row) GetColumnValueAsMap() (list map[string]int) {
	if r == nil {
		return
	}
	list = make(map[string]int)
	length := len(r.Columns) - 1
	for i := range r.Columns {
		if i+1 > length-i {
			return
		}
		list[(r.Columns)[i]] = i
		list[(r.Columns)[length-i]] = length - i

	}
	return
}

func (r *Row) FindValue(key string) (value string) {
	length := len(r.Columns) - 1
	for i := range r.Columns {
		if i+1 > length-i {
			return
		} else if r.Columns[i] == key {
			return r.Columns[i]
		} else if r.Columns[length-i] == key {
			return r.Columns[length-1]
		}
	}
	return
}

func (r *Row) FindValueInColumn(i int, search *Search) (value string) {
	for j := range search.Keys {
		if r.Columns[i] == search.Keys[j].Value {
			search.Keys[j].Found = true
			search.NotFound -= 1
			return r.Columns[i]
		}
	}
	return
}
