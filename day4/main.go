package main

import (
	"fmt"
	"github.com/waheedsys/assignment5/assignment5/calculator"
)

func main() {
	var op string
	for {
		fmt.Println("enter the cammand")
		fmt.Scanln(&op)
		switch op {
		case "+":
			fmt.Println(calculator.Add(2, 3))
		case "-":
			fmt.Println(calculator.Subtract(2, 3))
		case "*":
			fmt.Println(calculator.Multiply(2, 3))
		case "/":
			fmt.Println(calculator.Divide(2, 3))
		case "++":
			fmt.Println(calculator.AddToLastValue(3))
		case "--":

			fmt.Println(calculator.SubtractFromLastValue(1))
		}
	}

}

//func Add(a, b int) int {
//
//}
