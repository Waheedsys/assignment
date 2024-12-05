package assignment4

import "fmt"

type Set map[int]struct{}

func (s Set) Add(e int) Set {
	s[e] = struct{}{}
	s.printSet()
	return s
}

func (s Set) Remove(e int) Set {
	delete(s, e)
	s.printSet()
	return s
}

func (s Set) Has(e int) bool {
	_, ok := s[e]
	return ok
}

func From(elms ...int) Set {
	s := make(Set, len(elms))
	for _, e := range elms {
		s.Add(e)
	}
	return s
}

func (s Set) String() string {
	if len(s) == 0 {
		return "{}"
	}
	str := "{"
	for e := range s {
		str += fmt.Sprintf(" %d", e)
	}
	str += " }"
	return str
}

func (s Set) printSet() {
	fmt.Println("Current Set:", s.String())
}
