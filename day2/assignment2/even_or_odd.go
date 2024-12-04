package assignment2

import "fmt"

func EvenOrodd(para1 int) {
	i := para1

	if i%2 == 0 {
		fmt.Println("even")
	} else {
		fmt.Println("odd")
	}
}
