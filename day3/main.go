package main

import (
	"fmt"
)

func main() {
	/*
		l := []int{4, 1, 7, 3}
		fmt.Println("before reversing")
		fmt.Println(l)
		assignment4.Reverse(l)
		fmt.Println("after reversing")
		fmt.Println(l)

		assignment4.Takeword("waheed")
	*/

	//myMap := map[string][]int{
	//	"A": {1, 2, 3},
	//	"B": {4, 3, 2},
	//	"C": {8, 2, 6},
	//}
	//
	//assignment4.SumValuesByKey(myMap)
	//ans := []int{1, 4, 8, 3}
	//fmt.Println("the slice to map")
	//fmt.Println("slice", ans)
	//
	//assignment4.Slicetomap(ans)

	//s := assignment4.From(1, 2, 3)
	//
	//fmt.Println("Initial set:")
	//s.Add(4)
	//s.Remove(2)
	//s.Add(5)
	//s.Remove(1)
	//
	//fmt.Println("\nCheck if elements are present:")
	//fmt.Println("Set contains 3:", s.Has(3))
	//fmt.Println("Set contains 2:", s.Has(2))
	//fmt.Println("Set contains 5:", s.Has(5))

	newSet := NewSet()

	AddIntoSet(newSet, 1)
	AddIntoSet(newSet, 1)
	PrintSet(newSet)

	AddIntoSet(newSet, 2)
	PrintSet(newSet)

	RemoveElement(newSet, 1)
	PrintSet(newSet)
}

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
