package assignment1

import "fmt"

func Squareperimeter(sides int) {
	//return 4 * sides formula for square perimeter
	fmt.Printf("the perimeter of a square is:  %v\n", 4*sides)
}
func Reactangle(l, b int) {
	fmt.Printf("the perimeter of a square is:  %v\n", 2*l+b)
}

func Sphere(r float64) {
	fmt.Printf("the perimeter of a square is:  %v\n", float64(4/3)*Pi*r*r*r)
}
