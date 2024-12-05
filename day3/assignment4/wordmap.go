package assignment4

import "fmt"

func Takeword(para string) {

	count := make(map[string]int)
	for i := 0; i < len(para); i++ {
		count[string(para[i])]++
	}

	for k, v := range count {
		fmt.Printf("%d : %d\n", k, v)
	}
}
