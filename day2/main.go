package main

import (
	"fmt"
	day2assignment2 "github.com/waheedsys/newproject/day2/assignment2"
)

func main() {
	fmt.Println("type a number")
	var i, j, z, x, c int
	var y, a float64
	fmt.Scan(&i)
	day2assignment2.EvenOrodd(i)
	fmt.Println("type a number")
	fmt.Scan(&x)
	day2assignment2.ForLoop(x)
	fmt.Println("type a number")
	fmt.Scan(&z)
	fmt.Println(day2assignment2.IsPrime(z))
	fmt.Println("type a number")
	fmt.Scan(&j)
	day2assignment2.SwitchStatment(j)
	fmt.Println("type two numbers")
	fmt.Scan(&y, &a)
	day2assignment2.Calculator("+", y, a)
	fmt.Println("type a number")
	fmt.Scan(&c)
	day2assignment2.Sum(c)

}
