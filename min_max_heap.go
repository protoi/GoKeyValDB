package main

import "fmt"

// HeapType enum for type of heap, MIN or MAX
type HeapType int

const (
	MAXHEAP = iota
	MINHEAP
)

type Element struct {
	data   string
	weight int
}

type Heap struct {
	contents []*Element
	size     int
	ordering HeapType
}

func (heap *Heap) printHeap() {
	for _, elem := range heap.contents {
		fmt.Print(*elem)
	}
	fmt.Println()
}

func (heap *Heap) comparator(a, b int, operator string) bool {
	switch heap.ordering {
	case MAXHEAP:
		switch operator {
		case ">":
			return a > b
		case "<":
			return a < b
		}
	case MINHEAP:
		switch operator {
		case ">":
			return a < b
		case "<":
			return a > b
		}
	}
	return false
}

func (heap *Heap) heapPush(data string, weight int) {
	// Create a new element
	newElement := &Element{data, weight}

	// Add the new element to the end of the heap
	heap.contents = append(heap.contents, newElement)

	// Increase the heap size
	heap.size++

	// propagate the new element up the heap as necessary
	heap.propagateUp(heap.size - 1)
}

func (heap *Heap) heapPop() (Element, bool) {
	// Check if the heap is empty
	if heap.size == 0 {
		return Element{
			data:   "",
			weight: 0,
		}, false
	}

	// Get the element with the maximum weight (the root)
	maxElement := heap.contents[0]

	// Move the last element in the heap to the root
	heap.contents[0] = heap.contents[heap.size-1]

	// Remove the last element from the heap
	heap.contents = heap.contents[:heap.size-1]

	// Decrease the heap size
	heap.size--

	// propagate the new root element down the heap as necessary
	heap.propagateDown(0)

	return Element{
		data:   maxElement.data,
		weight: maxElement.weight,
	}, true
}

func (heap *Heap) heapify() {
	// Starting from the middle of the heap, propagate each element down the heap
	for i := heap.size/2 - 1; i >= 0; i-- {
		heap.propagateDown(i)
	}
}

func (heap *Heap) propagateUp(childIndex int) {
	// Get the parent index
	parentIndex := (childIndex - 1) / 2

	// If the parent element is less than the current element, swap them
	if parentIndex >= 0 && heap.comparator(heap.contents[parentIndex].weight, heap.contents[childIndex].weight, "<") {
		heap.contents[parentIndex], heap.contents[childIndex] = heap.contents[childIndex], heap.contents[parentIndex]

		// Recursively propagate the swapped element up the heap
		heap.propagateUp(parentIndex)
	}
}

func (heap *Heap) propagateDown(parentIndex int) {
	// Get the left and right child indices
	leftChildIndex := 2*parentIndex + 1
	rightChildIndex := 2*parentIndex + 2

	// Assume the current element is the largest
	largestIndex := parentIndex

	// If the left child exists and has a greater weight than the current element, set it as the largest
	if leftChildIndex < heap.size && heap.comparator(heap.contents[leftChildIndex].weight, heap.contents[largestIndex].weight, ">") {
		largestIndex = leftChildIndex
	}

	// If the right child exists and has a greater weight than the current element, set it as the largest
	if rightChildIndex < heap.size && heap.comparator(heap.contents[rightChildIndex].weight, heap.contents[largestIndex].weight, ">") {
		largestIndex = rightChildIndex
	}

	// perform a swap if the largest element has changed
	if largestIndex != parentIndex {
		heap.contents[largestIndex], heap.contents[parentIndex] = heap.contents[parentIndex], heap.contents[largestIndex]

		// Recursively propagate the swapped element down the heap
		heap.propagateDown(largestIndex)
	}
}

func testing() {
	// Create a new max heap
	heap := &Heap{}
	heap.ordering = MINHEAP
	// Add some elements to the heap
	heap.heapPush("A", 30)
	heap.heapPush("B", 2)
	heap.heapPush("C", 40)
	heap.heapPush("D", 14)
	heap.heapPush("E", 3)
	heap.heapPush("F", 2)
	heap.heapPush("G", 40)
	heap.heapPush("H", 100)
	heap.heapPush("I", 32)
	heap.heapPush("J", 22)
	heap.heapPush("K", 44)
	heap.heapPush("L", 14)

	// Heapify the heap
	heap.heapify()

	heap.printHeap()

	// Pop elements from the heap until it is empty
	for heap.size > 0 {
		fmt.Println(heap.heapPop())
	}
	heap.heapPop()
	heap.heapPop()
	heap.heapPop()
	fmt.Println(heap.heapPop())
	heap.heapPush("aaa", 100)
	fmt.Println(heap.heapPop())
}
