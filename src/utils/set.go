package utils

type empty struct{}

type t interface{}

type Set map[t]empty

func NewSet() Set {
	return make(map[t]empty)
}
func (s Set) Has(item t) bool {
	_, exists := s[item]
	return exists
}

func (s Set) Insert(item t) {
	s[item] = empty{}
}

func (s Set) Delete(item t) {
	delete(s, item)
}
