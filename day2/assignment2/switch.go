package assignment2

import "fmt"

func SwitchStatment(para int) {
	i := para
	switch i % 2 {
	case 0:
		fmt.Println("even")
	default:
		fmt.Println("odd")
	}
}
