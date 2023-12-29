package main

import (
	"fmt"

	bst "github.com/AlessandroFazio/DSAinGolang/binarysearchtrees"
)

func main() {
	t := bst.BinarySearchTree{}
	// t.InsertIteratively(10)
	// t.InsertIteratively(20)
	// t.InsertIteratively(5)
	// t.InsertIteratively(50)
	// t.InsertIteratively(100)
	// t.InsertIteratively(3)
	// t.InsertIteratively(1)

	t.InsertRecursively(10)
	t.InsertRecursively(20)
	t.InsertRecursively(5)
	t.InsertRecursively(50)
	t.InsertRecursively(100)
	t.InsertRecursively(3)
	t.InsertRecursively(1)
	// t.BfsTraverse()
	fmt.Println("Tree size: ", t.NumNodes)
	testNum := 10
	fmt.Printf("%v is in present in Tree: %v\n", testNum, t.SearchRecursively(testNum))
}