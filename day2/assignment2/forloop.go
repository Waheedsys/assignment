package assignment2

import "fmt"

func ForLoop(para2 int) {
	//a := [3]int{2, 3, 5}
	i := para2

	for j := 1; j < 11; j += 1 {
		fmt.Printf("%d  x %d = %d\n", i, j, i*j)
	}
	//for i, v := range a {
	//	print(i, v)
	//}

}
