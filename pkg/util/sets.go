package utils

type StringSet struct {
	Storage map[string]bool
}

func (s *StringSet) Size() int {
	return len(s.Storage)
}

func (s *StringSet) Items() []string {
	i := 0
	ret := make([]string, len(s.Storage))

	for key := range s.Storage {
		ret[i] = key
		i++
	}

	return ret
}

func (s *StringSet) Insert(val ...string) {
	for _, v := range val {
		s.Storage[v] = true
	}
}

func (s *StringSet) Remove(val ...string) {
	for _, v := range val {
		delete(s.Storage, v)
	}
}

func (s *StringSet) Contains(val string) bool {
	_, ok := s.Storage[val]
	return ok
}

func CreateStringSet(items []string) *StringSet {
	self := &StringSet{Storage: make(map[string]bool)}
	self.Insert(items...)
	return self
}
