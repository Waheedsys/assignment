package assignment4

func Slicetomap(list []int) (s map[int]int, err error) {
	result := make(map[int]int)
	for k := range list {
		result[k] = list[k]
	}

	//for k, v := range result {
	//	fmt.Printf("%v : %v\n", k, v)
	//}

	return result, nil
}
