package main

import (
	"fmt"
	"github.com/waheedsys/newproject/day5/assignment6"
)

func main() {
	//p := &assignment6.Calculator{3, 3}
	//fmt.Println(p.Add())
	//fmt.Println(p.Subtraction())
	//fmt.Println(p.Multiplication())
	//fmt.Println(p.Division())

	// compare function
	fmt.Println(assignment6.SortSlice([]int{6, 2, 9, 5}, assignment6.Compare))

}
