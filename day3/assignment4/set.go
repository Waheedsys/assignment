package assignment4

import "fmt"

//
//type Set map[int]struct{}
//
//func (s Set) Add(e int) Set {
//	s[e] = struct{}{}
//	s.printSet()
//	return s
//}
//
//func (s Set) Remove(e int) Set {
//	delete(s, e)
//	s.printSet()
//	return s
//}
//
//func (s Set) Has(e int) bool {
//	_, ok := s[e]
//	return ok
//}
//
//func From(elms ...int) Set {
//	s := make(Set, len(elms))
//	for _, e := range elms {
//		s.Add(e)
//	}
//	return s
//}
//
//func (s Set) String() string {
//	if len(s) == 0 {
//		return "{}"
//	}
//	str := "{"
//	for e := range s {
//		str += fmt.Sprintf(" %d", e)
//	}
//	str += " }"
//	return str
//}
//
//func (s Set) printSet() {
//	fmt.Println("Current Set:", s.String())
//}
//

func NewSet() map[int]bool {
	mp := make(map[int]bool)
	return mp
}

func AddIntoSet(mp map[int]bool, element int) {
	fmt.Println("adding to set: ", element)
	mp[element] = true
}

func RemoveElement(mp map[int]bool, element int) {
	fmt.Println("deleting element from set: ", element)
	delete(mp, element)
}

func PrintSet(mp map[int]bool) {
	fmt.Println("current element in the set: ")
	for k, _ := range mp {
		fmt.Print(k, " ")
	}

	fmt.Println()
}
