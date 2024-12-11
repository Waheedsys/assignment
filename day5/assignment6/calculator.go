package assignment6

// calculator structure have two feilds
type Calculator struct {
	X, Y int
}

// Add is a method return addition of two numbers
func (c Calculator) Add() int {
	return c.X + c.Y
}

// Subtraction is a method return subtraction of two numbers
func (c Calculator) Subtraction() int {
	return c.X - c.Y
}

// Multiplication is a method return multiplication of two numbers
func (c Calculator) Multiplication() int {
	return c.X * c.Y
}

// Division is a method return division of two numbers
func (c Calculator) Division() int {
	return c.X / c.Y
}
