package linkedlists

import "fmt"

type node struct{
	data int
	next *node
	previous *node
}

type LinkedList struct{
	head *node
	tail *node
	Size int
}

func (l *LinkedList) AddFirst(data int) {
	newNode := &node{data: data}
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

func (l *LinkedList) AddLast(data int) {
	newNode := &node{data: data}
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

func (l *LinkedList) RemoveFirst() int {
	if l.Size == 0 {
		fmt.Println("LinkedList is empty. Nothing to remove")
		return -1
	}
	oldHead := l.head
	l.head = oldHead.next
	if l.Size == 1 {
		l.tail = l.head
	} else {
		l.head.previous = nil
	}
	l.Size--
	return oldHead.data
}

func (l *LinkedList) RemoveLast() int {
	if l.Size == 0 {
		fmt.Println("LinkedList is empty. Nothing to remove")
		return -1
	}
	oldTail := l.tail
	l.tail = oldTail.previous
	if l.Size == 1 {
		l.head = l.tail
	} else {
		l.tail.next = nil
	}
	l.Size--
	return oldTail.data
}

func (l *LinkedList) Remove(data int) int{
	if l.Size == 0 {
		fmt.Println("LinkedList is empty. Nothing to remove")
		return -1
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
			return removed
		}
		current = current.next
	}
	return -1
}

func (l LinkedList) Traverse() {
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
