package assignment1

func Squareperimeter(sides int) int {
	//return 4 * sides formula for square perimeter
	return 4 * sides
}
func Reactangle(l, b int) int {
	return 2 * (l + b)
}

func Cube(l, b, h int) int {
	return l * b * h
}

func Sphere(r float64) float64 {
	return float64(4/3) * Pi * r * r * r
}
