package binarysearchtrees

import (
	"fmt"
	"strings"
	q "github.com/AlessandroFazio/DSAinGolang/queues"
)

type node struct {
	left *node
	right *node
	data int
}

type BinarySearchTree struct {
	Root *node
	NumNodes int
}

// Insert Iteratively
func (t *BinarySearchTree) InsertIteratively(e int) {
	if t.IsEmpty() {
		t.Root = &node{data: e}
		t.NumNodes++
		return
	}
	current := t.Root
	for current != nil {
		if e == current.data {
			return
		} else if e > current.data {
			if current.left == nil {
				current.left = &node{data: e}
				t.NumNodes++
				return
			} else {
				current = current.left
			}
		} else if e < current.data {
			if current.right == nil {
				current.right = &node{data: e}
				t.NumNodes++
				return
			} else {
				current = current.right
			}
		}
	}
}

func (t *BinarySearchTree) InsertRecursively(e int) {
	if t.IsEmpty() {
		t.Root = &node{data: e}
		t.NumNodes++
		return
	} 
	t.Root.doInsertRecursively(e)
	t.NumNodes++
}

// Insert Recursively
func (n *node) doInsertRecursively(e int) {
	if e == n.data {
		return
	} else if e > n.data {
		if n.left == nil {
			n.left = &node{data: e}
		} else {
			n.left.doInsertRecursively(e)
		}
	} else if e < n.data {
		if n.right == nil {
			n.right = &node{data: e}
		} else {
			n.right.doInsertRecursively(e)
		}
	}
}

func (t *BinarySearchTree) IsEmpty() bool {
	return t.Root == nil 
}

//Search
func (t *BinarySearchTree) SearchIteratively(e int) bool {
	if t.IsEmpty() {
		fmt.Println("Tree is empty. Nothing to search")
		return false
	}

	current := t.Root
	for current != nil {
		if e < current.data {
			if current.right == nil {
				return false
			} else {
				current = current.right
			}
		} else if e > current.data {
			if current.left == nil {
				return false
			} else {
				current = current.left
			}
		} else {
			return true
		}
	}
	return false
}

func (t *BinarySearchTree) SearchRecursively(e int) bool {
	if t.IsEmpty() {
		fmt.Println("Tree is empty. Nothing to search")
		return false
	}
	return t.Root.doSearchRecursively(e)
}

func (n *node) doSearchRecursively(e int) bool {
	if n == nil {
		return false
	}
	if e < n.data {
		return n.right.doSearchRecursively(e)
	} else if e > n.data {
		return n.left.doSearchRecursively(e)
	}
	return true
}

// Traverse
func (t *BinarySearchTree) BfsTraverse() {
	if t.IsEmpty() {
		fmt.Println("Tree is empty. Nothing to traverse")
		return
	}
	start := t.Root
	queue := q.Queue[*node]{}
	levelMap := make(map[*node]int)
	levelMap[start] = 1

	currentLevel := 1
	var result strings.Builder
	
	queue.Push(start)
	for !queue.IsEmpty() {
		current := queue.Pop()
		if levelMap[current] > currentLevel {
			result.WriteString("\n")
			currentLevel = levelMap[current]
		}

		if current != start {
			result.WriteString(strings.Repeat("  ", levelMap[current]-1))
			result.WriteString("|__ ")
		}
		fmt.Fprintf(&result, "%d", current.data)

		if current.left != nil {
			queue.Push(current.left)
			levelMap[current.left] = levelMap[current] + 1
		}
		if current.right != nil {
			queue.Push(current.right)
			levelMap[current.right] = levelMap[current] + 1
		}
	}

	fmt.Println(result.String())
}