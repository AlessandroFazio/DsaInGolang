package heaps

import "fmt"

// MaxHeap struct has a slice that holds the array
type MaxHeap struct{
	array []int
}

// Insert adds an element to the heap
func (h *MaxHeap) Insert(key int) {
	h.array = append(h.array, key)
	h.maxHeapifyUp(len(h.array)-1)
}

func (h *MaxHeap) Extract() int{
	l := len(h.array)
	if l == 0 {
		fmt.Println("Heap is empty. Nothing to extract")
		return -1
	}

	extracted := h.array[0]

	if l == 1 {
		h.array = []int{}
		return extracted
	}

	h.array[0] = h.array[l-1]
	h.array = h.array[:l-1]
	h.maxHeapifyDown(0)

	return extracted
}

// maxHeapify will heapify up the MaxHeap 
func (h *MaxHeap) maxHeapifyUp(index int) {
	for h.array[parent(index)] < h.array[index] {
		h.swap(parent(index), index)
		index = parent(index)
	}
}

func (h *MaxHeap) maxHeapifyDown(index int) {
	lastIndex := len(h.array) - 1
	l,r := left(index), right(index)
	childToCompare := 0

	for l <= lastIndex {
		if l == lastIndex {
			childToCompare = l
		} else if h.array[l] > h.array[r] {
			childToCompare = l
		} else {
			childToCompare = r
		}

		if h.array[index] < h.array[childToCompare] {
			h.swap(index, childToCompare)
			index = childToCompare
			l = left(childToCompare)
			r = right(childToCompare)
		} else {
			return
		}
	}
} 

func parent(i int) int{
	return (i-1)/2
}

func left(i int) int{
	return i*2 + 1 
}

func right(i int) int{
	return i*2 + 2
}

func (h *MaxHeap) swap(i1, i2 int) {
	h.array[i1], h.array[i2] = h.array[i2], h.array[i1]
}
