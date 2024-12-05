package assignment4

import "fmt"

func Slicetomap(list []int) {
	result := make(map[int]int)
	for k := range list {
		result[k] = list[k]
	}

	for k, v := range result {
		fmt.Printf("%v : %v\n", k, v)
	}
}
