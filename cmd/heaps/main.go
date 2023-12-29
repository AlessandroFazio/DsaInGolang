package main

import (
	"fmt"
	h "github.com/AlessandroFazio/DSAinGolang/heaps"
)

func main() {
	m := &h.MaxHeap{}
	fmt.Println(m)
	buildHeap := []int{10,20,30,40,50,60,70,80,90,100}
	for _,v := range buildHeap {
		m.Insert(v)
		fmt.Println(m)
	}

	for i := 0; i < len(buildHeap); i++ {
		v := m.Extract()
		fmt.Println("Extracted: ", v)
		// fmt.Println(m)
	}
}