package assignment2

import "fmt"

func Sum(n int) {
	var sum int
	for j := 1; j <= n; j++ {
		sum = sum + j
	}
	fmt.Println(sum)
}
