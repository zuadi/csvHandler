package models

type Column []string

func (c *Column) FindValueInColumns(key string) (value string, found bool) {
	if c == nil {
		return
	}
	length := len(*c) - 1
	for i := range *c {
		if i+1 > length-i {
			return
		} else if (*c)[i] == key {
			return (*c)[i], true
		} else if (*c)[length-i] == key {
			return (*c)[i], true
		}
	}
	return
}
