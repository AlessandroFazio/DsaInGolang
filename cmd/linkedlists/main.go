package main

import (
	// "fmt"
	"fmt"

	ll "github.com/AlessandroFazio/DSAinGolang/linkedlists"
)

func main() {
	l := ll.LinkedList[string]{}
	// with integer
	// l := ll.LinkedList[int]{}
	// l.AddLast(10)
	// l.AddLast(20)
	// l.AddLast(30)
	// l.AddLast(40)
	// l.AddLast(50)

	l.AddLast("mio")
	l.AddLast("tuo")
	l.AddLast("suo")
	l.AddLast("nostro")
	l.AddLast("vostro")
	l.AddLast("loro")

	// for i := 0; i < 5; i++ {
	//	fmt.Printf("Size of the LinkedList is: %v\n", l.Size)
	//	fmt.Printf("Removed: %v\n", l.RemoveFirst())
	// }

	l.Traverse()
	if value,err := l.Remove("mio"); err != nil {
		fmt.Println("Error removing element:", err)
	} else {
		fmt.Println("\nRemove with Value: ", value)
	}
	fmt.Println("newSize: ", l.Size)
	l.Traverse()
	if value,err := l.Remove("mio"); err != nil {
		fmt.Println("Error removing element:", err)
	} else {
		fmt.Println("\nRemove with Value: ", value)
	}
}