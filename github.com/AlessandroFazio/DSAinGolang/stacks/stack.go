package stacks

import "fmt"

type Stack struct {
	array []int
}

func (s *Stack) Push(e int) {
	s.array = append(s.array, e)
}

func (s *Stack) Pop() int {
	l := len(s.array)
	if l == 0 {
		fmt.Println("Stack is empty. Nothing to pop")
		return -1
	}
	e := s.array[l-1]
	s.array = s.array[:l-1]
	return e
}

func (s *Stack) Size() int {
	return len(s.array)
}