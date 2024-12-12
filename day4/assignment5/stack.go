package assignment5

import "fmt"

type Stack struct {
	arr []int
}

func (s *Stack) Push(para int) {
	s.arr = append(s.arr, para)
}
func (s *Stack) Pop() (int, error) {
	if len(s.arr) == 0 {
		return 0, fmt.Errorf("stack is empty")
	}
	res := s.arr[len(s.arr)-1]
	s.arr = s.arr[:len(s.arr)-1] // Remove the last element
	return res, nil
}
func (s *Stack) GetStack() []int {
	return s.arr
}
