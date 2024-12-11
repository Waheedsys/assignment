package assignment4

func Reverse(list []int) (l []int, err error) {
	for i, j := 0, len(list)-1; i < j; {
		list[i], list[j] = list[j], list[i]
		i++
		j--
	}
	return list, nil
}
