// Package medianheap implements a running median algorithm for integers
// using 2 heaps. The provided operations are for adding elements and
// retrieving the median.
//
// The time complexity for updating the median is O(logN), retrieving it O(1).
//
// The median value is equal to the k/2th element of a sorted array of size
// k if k is odd, or (k/2)-1th element if size is even. No averages are
// computed.
//
// Here is equivalent code (with time complexity O(N logN)):
//
//   sort.Ints(A)
//   if len(A) % 2 == 0 {
//      m = A[(len(A)-1) / 2]
//   } else {
//      m = A[len(A) / 2]
//   }
//
package medianheap

import (
	"container/heap"
)

type intMinHeap []int

func (h intMinHeap) Len() int           { return len(h) }
func (h intMinHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h intMinHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *intMinHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *intMinHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

type intMaxHeap []int

func (h intMaxHeap) Len() int           { return len(h) }
func (h intMaxHeap) Less(i, j int) bool { return h[i] > h[j] }
func (h intMaxHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

// Peek at the heap root element without popping it out.
func (h intMaxHeap) Peek() int { return h[0] }

func (h *intMaxHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *intMaxHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

// An IntMedianHeap maintains references to added elements.
type IntMedianHeap struct {
	max *intMaxHeap
	min *intMinHeap
}

// New returns an initialized median heap.
func New() *IntMedianHeap {
	h := &IntMedianHeap{
		max: &intMaxHeap{},
		min: &intMinHeap{},
	}
	heap.Init(h.max)
	heap.Init(h.min)

	return h
}

// Add inserts an integer into the medianheap.
func (h IntMedianHeap) Add(n interface{}) {
	if h.max.Len() == 0 {
		heap.Push(h.max, n)
		return
	}

	if n.(int) > h.max.Peek() {
		heap.Push(h.min, n)
	} else {
		heap.Push(h.max, n)
	}

	// Balance the heaps so that either the 'max' always
	// contains 1 more element, or both heaps are of equal size.
	if h.max.Len()-h.min.Len() == 2 {
		heap.Push(h.min, heap.Pop(h.max))
	} else if h.min.Len()-h.max.Len() == 1 {
		heap.Push(h.max, heap.Pop(h.min))
	}
}

// Median returns a median value for added elements.
func (h IntMedianHeap) Median() int {
	// There is no default median value.
	if h.max.Len() == 0 {
		panic("median heap is empty")
	}

	// With well-balanced heaps, the running median is
	// always the root element of the 'max' heap.
	return h.max.Peek()
}

// Update adds an element to the heap and returns the running median.
// It is a utility concatenation of Add(), followed by Median().
func (h IntMedianHeap) Update(n interface{}) int {
	h.Add(n)
	return h.Median()
}
