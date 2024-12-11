package main

import (
	"fmt"
	"github.com/Waheedsys/assignment7/assignment7"
	"math"
)

func main() {
	var a assignment7.Abser
	f := assignment7.MyFloat(-math.Sqrt2)
	v := assignment7.Vertex{3, 4}

	a = f  // a MyFloat implements Abser
	a = &v // a *Vertex implements Abser

	fmt.Println(a.Abs())
}
