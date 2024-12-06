package calculator

var global int
var global1 int

func Add(a, b int) int {
	global = a + b
	return global
}

func Subtract(a, b int) int {
	if a < b {
		global = b - a
		return global
	}
	global = a - b
	return global
}

func Multiply(a, b int) int {
	global = a * b
	return global
}

func Divide(a, b int) float64 {
	global = a / b
	return float64(global)
}

func AddToLastValue(para int) int {
	global1 = global + para
	return global1
}
func SubtractFromLastValue(para int) int {
	global1 = global - para
	return global1
}
