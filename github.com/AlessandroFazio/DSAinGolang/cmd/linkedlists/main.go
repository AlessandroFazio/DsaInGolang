package main

import (
	// "fmt"
	"fmt"

	ll "github.com/AlessandroFazio/DSAinGolang/linkedlists"
)

func main() {
	l := ll.LinkedList{}
	l.AddLast(10)
	l.AddLast(20)
	l.AddLast(30)
	l.AddLast(40)
	l.AddLast(50)

	// for i := 0; i < 5; i++ {
	//	fmt.Printf("Size of the LinkedList is: %v\n", l.Size)
	//	fmt.Printf("Removed: %v\n", l.RemoveFirst())
	// }

	l.Traverse()
	fmt.Println("\nRemove with Value: ", l.Remove(10))
	fmt.Println("newSize: ", l.Size)
	l.Traverse()
	fmt.Println("Removed 10 again: ", l.Remove(10))
}