package main

import (
	"fmt"
)

type MinHeap struct {
	array []int
	size  int
}

func NewMinHeap() *MinHeap {
	return &MinHeap{
		array: make([]int, 0),
		size:  0,
	}
}

func (h *MinHeap) parentIndex(index int) int {
	return (index - 1) / 2
}

func (h *MinHeap) leftChildIndex(index int) int {
	return 2*index + 1
}

func (h *MinHeap) rightChildIndex(index int) int {
	return 2*index + 2
}

func (h *MinHeap) swap(i, j int) {
	h.array[i], h.array[j] = h.array[j], h.array[i]
}

func (h *MinHeap) insert(value int) {
	h.array = append(h.array, value)
	h.size++
	h.heapifyUp(h.size - 1)
}

func (h *MinHeap) extractMin() (int, error) {
	if h.size == 0 {
		return 0, fmt.Errorf("Heap is empty")
	}

	min := h.array[0]
	h.array[0] = h.array[h.size-1]
	h.array = h.array[:h.size-1]
	h.size--
	h.heapifyDown(0)

	return min, nil
}

func (h *MinHeap) heapifyUp(index int) {
	for index > 0 && h.array[index] < h.array[h.parentIndex(index)] {
		parentIndex := h.parentIndex(index)
		h.swap(index, parentIndex)
		index = parentIndex
	}
}

func (h *MinHeap) heapifyDown(index int) {
	smallest := index
	leftChildIndex := h.leftChildIndex(index)
	rightChildIndex := h.rightChildIndex(index)

	if leftChildIndex < h.size && h.array[leftChildIndex] < h.array[smallest] {
		smallest = leftChildIndex
	}

	if rightChildIndex < h.size && h.array[rightChildIndex] < h.array[smallest] {
		smallest = rightChildIndex
	}

	if smallest != index {
		h.swap(index, smallest)
		h.heapifyDown(smallest)
	}
}

func main() {
	minHeap := NewMinHeap()

	minHeap.insert(5)
	minHeap.insert(2)
	minHeap.insert(8)
	minHeap.insert(1)
	minHeap.insert(10)

	min, err := minHeap.extractMin()
	if err == nil {
		fmt.Println("Extracted minimum value:", min)
	}

	min, err = minHeap.extractMin()
	if err == nil {
		fmt.Println("Extracted minimum value:", min)
	}
}
