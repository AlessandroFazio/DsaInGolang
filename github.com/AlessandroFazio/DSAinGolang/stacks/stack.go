package stacks

import "fmt"

type Stack[T any] struct {
	array []T
}

func (s *Stack[T]) Push(e T) {
	s.array = append(s.array, e)
}

func (s *Stack[T]) Pop() T {
	l := len(s.array)
	if l == 0 {
		fmt.Println("Stack is empty. Nothing to pop")
		return *new(T)
	}
	e := s.array[l-1]
	s.array = s.array[:l-1]
	return e
}

func (s *Stack[T]) Size() int {
	return len(s.array)
}