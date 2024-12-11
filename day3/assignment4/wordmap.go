package assignment4

func Takeword(para string) (s map[string]int, err error) {

	count := make(map[string]int)
	for i := 0; i < len(para); i++ {
		count[string(para[i])]++
	}

	return count, nil
}
