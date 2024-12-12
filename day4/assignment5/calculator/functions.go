package calculator

import (
	"errors"
)

// Add performs addition
func Add(a, b int) int {
	return a + b
}

// Subtract performs subtraction
func Subtract(a, b int) int {
	return a - b
}

// Multiply performs multiplication
func Multiply(a, b int) int {
	return a * b
}

// Divide performs division with error handling for division by zero
func Divide(a, b int) (float64, error) {
	if b == 0 {
		return 0.0, errors.New("cannot divide by zero")
	}
	return float64(a) / float64(b), nil
}

// AddToLastValue adds a number to the last result
// Note: `lastResult` should be passed to this function
func AddToLastValue(lastResult, para int) int {
	return lastResult + para
}

// SubtractFromLastValue subtracts a number from the last result
// Note: `lastResult` should be passed to this function
func SubtractFromLastValue(lastResult, para int) int {
	return lastResult - para
}
