package main

import "fmt"

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

func (s *Set) Add(i ...int) {
	for _, val := range i {
		if !valueInSlice(val, s.elements) {
			s.elements = append(s.elements, val)
		}
	}
}

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

func NewSet() *Set {
	s := &Set{}
	return s
}

func main() {
	s := NewSet()
	s.Add(1, 2, 3)
	s.Add(5, 6)
	s.Add(5)
	s.Add(1)

	s1 := NewSet()
	s1.Add(1, 2, 3, 4)

	fmt.Println(s.elements)
	fmt.Println(s1.elements)
	fmt.Println(s.Union(s1).elements)
	fmt.Println(s.Intersect(s1).elements)
	fmt.Println(s.Difference(s1).elements)
	fmt.Println(s1.Difference(s).elements)
}
