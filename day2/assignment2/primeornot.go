package assignment2

func IsPrime(n int) bool {
	if n < 2 {
		return false
	}
	i := 2
	for i*i <= n {
		if n%i == 0 {
			return false
		}
		i++
	}
	return true
}
