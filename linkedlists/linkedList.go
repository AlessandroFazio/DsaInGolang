package linkedlists

import "fmt"

type node[T comparable] struct{
	data T
	next *node[T]
	previous *node[T]
}

type LinkedList[T comparable] struct{
	head *node[T]
	tail *node[T]
	Size int
}

func (l *LinkedList[T]) AddFirst (data T) {
	newNode := &node[T]{data: data}
	oldHead := l.head
	l.head = newNode
	newNode.next = oldHead
	if l.Size == 0 {
		l.tail = newNode
	} else {
		oldHead.previous = newNode
	}
	l.Size++
}

func (l *LinkedList[T]) AddLast(data T) {
	newNode := &node[T]{data: data}
	oldTail := l.tail
	l.tail = newNode
	newNode.previous = oldTail
	if l.Size == 0 {
		l.head = newNode
	} else {
		oldTail.next = newNode
	}
	l.Size++
}

func (l *LinkedList[T]) RemoveFirst() (T, error) {
	if l.Size == 0 {
		return *new(T), fmt.Errorf("LinkedList is empty. Nothing to remove")
	}
	oldHead := l.head
	l.head = oldHead.next
	if l.Size == 1 {
		l.tail = l.head
	} else {
		l.head.previous = nil
	}
	l.Size--
	return oldHead.data, nil
}

func (l *LinkedList[T]) RemoveLast() (T, error) {
	if l.Size == 0 {
		return *new(T), fmt.Errorf("LinkedList is empty. Nothing to remove")
	}
	oldTail := l.tail
	l.tail = oldTail.previous
	if l.Size == 1 {
		l.head = l.tail
	} else {
		l.tail.next = nil
	}
	l.Size--
	return oldTail.data, nil
}

func (l *LinkedList[T]) Remove(data T) (T, error) {
	if l.Size == 0 {
		return *new(T), fmt.Errorf("LinkedList is empty. Nothing to remove")
	}
	current := l.head
	for current != nil {
		if current.data == data {
			removed := current.data
			if current.previous != nil {
				current.previous.next = current.next
			} else {
				l.head = current.next
			}
			if current.next != nil { 
				current.next.previous = current.previous
			} else {
				l.tail = current.previous
			}
			l.Size--
			return removed, nil
		}
		current = current.next
	}
	return *new(T), nil
}

func (l LinkedList[T]) Traverse() {
	if l.Size == 0 {
		fmt.Println("LinkedList is empty. Nothing to print")
		return
	}
	current := l.head
	for current != nil {
		fmt.Printf("<-> %v ", current.data)
		current = current.next
	}
	fmt.Printf("\n")
}
