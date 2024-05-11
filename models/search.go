package models

type Search struct {
	NotFound int
	Keys     []Key
}

func NewSearch() *Search {
	return &Search{}
}

func (s *Search) AddKeys(value ...string) {
	for i := range value {
		s.Keys = append(s.Keys, Key{
			Value: value[i],
		})
		s.NotFound += 1
	}
}

func (s *Search) AllFound() bool {
	return s.NotFound == 0
}

func (s *Search) GetNotFound() (list *[]string) {
	if s.AllFound() {
		return
	}
	list = &[]string{}

	for i := range s.Keys {
		if !s.Keys[i].Found {
			*list = append(*list, s.Keys[i].Value)
		}
	}
	return
}
