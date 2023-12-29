package main

import (
	"fmt"
	q "github.com/AlessandroFazio/DSAinGolang/queues"
)

func main() {
	queue := q.Queue[string]{}
	fmt.Println("Queue size: ", queue.Size())
	queue.Push("1")
	fmt.Println("Pop: ", queue.Pop())
	queue.Push("2")
	queue.Push("5")
	queue.Push("3")
	fmt.Println("Queue size: ", queue.Size())
	fmt.Println("Pop: ", queue.Pop())
	fmt.Println("Pop: ", queue.Pop())
	fmt.Println("Pop: ", queue.Pop())
	fmt.Println("Queue size: ", queue.Size())
}