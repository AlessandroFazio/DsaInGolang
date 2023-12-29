package main

import (
	"fmt"
	h "github.com/AlessandroFazio/DSAinGolang/hashtables"
)

func main() {
	table := h.HashTable{}
	table.Insert("firstKey")
	table.Insert("secondKey")
	table.Insert("thirdKey")
	table.Insert("fourthKey")
	fmt.Println("key (thirdKey) is present: ", table.Search("thirdKey"))
	table.Delete("thirdKey")
	fmt.Println("key (thirdKey) is present: ", table.Search("thirdKey"))
}