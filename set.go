package set

// Set set implementation using go
type Set struct {
	elements []int
}

func valueInSlice(a int, list []int) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}

// GetElements get all elements of set
func (s *Set) GetElements() []int {
	return s.elements
}

// Add add element in set
func (s *Set) Add(i ...int) {
	for _, val := range i {
		if !valueInSlice(val, s.elements) {
			s.elements = append(s.elements, val)
		}
	}
}

// Union get union of sets
func (s *Set) Union(s1 *Set) *Set {
	s2 := NewSet()
	for _, val := range s.elements {
		s2.Add(val)
	}
	for _, val := range s1.elements {
		s2.Add(val)
	}
	return s2
}

// Intersect get intersection of set
func (s *Set) Intersect(s1 *Set) *Set {
	s2 := NewSet()
	for _, val := range s.elements {
		for _, val1 := range s1.elements {
			if val == val1 {
				s2.Add(val)
			}
		}
	}
	return s2
}

// Difference get difference of set
func (s *Set) Difference(s1 *Set) *Set {
	s2 := NewSet()
	for _, val := range s.elements {
		add := true
		for _, val1 := range s1.elements {
			if val == val1 {
				add = false
			}
		}
		if add {
			s2.Add(val)
		}
	}
	return s2
}

// NewSet get new set
func NewSet() *Set {
	s := &Set{}
	return s
}
