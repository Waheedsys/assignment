package assignment2

import "fmt"

func Calculator(op string, num1, num2 float64) {
	switch op {
	case "+":

		fmt.Println(num1 + num2)
	case "-":
		fmt.Println(num1 - num2)
	case "*":
		fmt.Println(num1 * num2)
	case "/":
		fmt.Println(num1 / num2)
	default:
		fmt.Println("value not recognized")
	}
}
