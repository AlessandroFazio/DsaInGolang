package queues

import "fmt"

type Queue[T any] struct {
	array []T
}

func (q *Queue[T]) Push(e T) {
	q.array = append(q.array, e)
}

func (q *Queue[T]) Pop() T {
	if len(q.array) == 0 {
		fmt.Println("Queue is empty. Nothing to pop")
		return *new(T)
	}
	e := q.array[0]
	q.array = q.array[1:]
	return e
}

func (q *Queue[T]) Size() int {
	return len(q.array)
}