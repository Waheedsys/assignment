package assignment2

func Sum(n int) (s int, err error) {
	var sum int
	for j := 1; j <= n; j++ {
		sum = sum + j
	}
	return sum, nil
}
