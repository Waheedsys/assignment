package assignment5

func push(para int, arr []int) {
	arr = append(arr, para)
}

func pop(arr []int) int {
	res := arr[len(arr)-1]
	return res
}
func Stack(para int, arr []int) (s []int, err error) {
	push(para, arr)
	pop(arr)
	return arr, nil

}
