package main

import "log"

type MaxHeap struct {
	array []int
}

func (h *MaxHeap) InsertElement(k int) {
	h.array = append(h.array, k)
	h.HeapifyUp()
	log.Println(h.array)
}

func (h *MaxHeap) HeapifyUp() {
	index := len(h.array) - 1

	for h.array[Parent(index)] < h.array[index] {
		h.swap(Parent(index), index)
		index = Parent(index)
	}
}

func LeftChild(k int) int {
	return k*2 + 1
}

func RightChild(k int) int {
	return k*2 + 2
}

func Parent(k int) int {
	return (k - 1) / 2
}

func (h *MaxHeap) swap(i1, i2 int) {
	h.array[i1], h.array[i2] = h.array[i2], h.array[i1]
}

func (h *MaxHeap) ExtractElement() {
	// log.Printf("Extracted Element : %d\n", h.array[0])
	h.array[0] = h.array[len(h.array)-1]
	h.array = h.array[0 : len(h.array)-1]

	h.HeapifyDown(0)

	log.Println(h.array)
}

func (h *MaxHeap) HeapifyDown(index int) {
	l, r := LeftChild(index), RightChild(index)

	for l < len(h.array)-1 {
		childToCompare := r
		if h.array[l] > h.array[r] {
			childToCompare = l
		}

		if h.array[childToCompare] > h.array[index] {
			h.swap(index, childToCompare)
			index = childToCompare
		} else {
			break
		}
		l, r = LeftChild(index), RightChild(index)
	}
}

func main() {
	m := MaxHeap{}
	// buildHeap := []int{50, 16, 48, 14, 8, 34, 20, 9, 1, 5, 7, 63}
	buildHeap := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12}

	for _, v := range buildHeap {
		m.InsertElement(v)
	}

	m.ExtractElement()
	m.ExtractElement()
	m.ExtractElement()

}
