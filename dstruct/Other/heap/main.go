package main

import "log"

type MaxHeap struct {
	array []int
}

// Insert an element into the heap
func (h *MaxHeap) Insert(num int) {
	h.array = append(h.array, num)
	h.maxHeapifyUp()
}

func (h *MaxHeap) maxHeapifyUp() {
	index := len(h.array) - 1
	for h.array[parent(index)] < h.array[index] {
		h.swap(parent(index), index)
		index = parent(index)
	}
}

func (h *MaxHeap) swap(i1, i2 int) {
	h.array[i1], h.array[i2] = h.array[i2], h.array[i1]
}

func parent(i int) int {
	return (i - 1) / 2
}

func leftChild(i int) int {
	return 2*i + 1
}

func rightChild(i int) int {
	return 2*i + 2
}

// Extract an element from the heap
func (h *MaxHeap) Extract() int {
	extractedNum := h.array[0]
	l := len(h.array) - 1
	h.array[0] = h.array[l]
	h.array = h.array[:l]
	h.maxHeapifyDown()
	return extractedNum
}

func (h *MaxHeap) maxHeapifyDown() {
	index := 0
	lastIndex := len(h.array) - 1
	l, r := leftChild(index), rightChild(index)
	childToCompare := 0

	for l <= lastIndex {
		if l == lastIndex {
			childToCompare = l
		} else if h.array[l] < h.array[r] {
			childToCompare = r
		} else {
			childToCompare = l
		}
		if h.array[index] < h.array[childToCompare] {
			h.swap(index, childToCompare)
			index = childToCompare
			l, r = leftChild(index), rightChild(index)
		}
	}
}

func main() {
	m := MaxHeap{}
	buildHeap := []int{50, 16, 48, 14, 8, 34, 20, 9, 1, 5, 7, 63}

	for _, i := range buildHeap {
		m.Insert(i)
		log.Printf("%v", m.array)
	}

	for i := 0; i < 10; i++ {
		m.Extract()
		log.Printf("%v", m.array)
	}
}
