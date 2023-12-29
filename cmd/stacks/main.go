package main

import (
	"fmt"
	s "github.com/AlessandroFazio/DSAinGolang/stacks"
)

func main() {
	stack := s.Stack[string]{}
	fmt.Println("Stack size: ", stack.Size())
	stack.Push("1")
	fmt.Println("Pop: ", stack.Pop())
	stack.Push("2")
	stack.Push("5")
	stack.Push("3")
	fmt.Println("Stack size: ", stack.Size())
	fmt.Println("Pop: ", stack.Pop())
	fmt.Println("Pop: ", stack.Pop())
	fmt.Println("Pop: ", stack.Pop())
	fmt.Println("Stack size: ", stack.Size())
}