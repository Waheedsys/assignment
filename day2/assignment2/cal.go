package assignment2

import "fmt"

func Calculator(op string, num1, num2 float64) (res float64, err error) {
	switch op {
	case "+":

		return num1 + num2, nil
	case "-":

		return num1 - num2, nil
	case "*":

		return num1 * num2, nil
	case "/":
		if num2 == 0 {
			return 0, fmt.Errorf("cannot divide by zero")
		}
		return num1 / num2, nil
	default:
		return 0, fmt.Errorf("invalid operator %s", op)
	}
}
