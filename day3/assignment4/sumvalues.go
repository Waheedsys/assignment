package assignment4

import "fmt"

func SumValuesByKey(m map[string][]int) {
	result := make(map[string]int)
	for key, values := range m {
		sum := 0
		for _, value := range values {
			sum += value
		}
		result[key] = sum
	}

	for k, v := range result {
		fmt.Printf("%v : %v\n", k, v)
	}
}
