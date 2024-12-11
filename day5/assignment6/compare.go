package assignment6

func SortSlice(ar []int, Compare func(int, int) bool) []int {
	for i := 0; i < len(ar)-1; i++ {

		for j := i + 1; j < len(ar); j++ {
			if Compare(ar[i], ar[j]) {
				ar[j], ar[i] = ar[i], ar[j]
			}
		}

	}
	return ar

}

func Compare(a, b int) bool {
	return a > b
}
