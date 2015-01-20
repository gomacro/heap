// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package heap provides a heap (a priority queue) operations on an int32 slice.
package heap

// Pop is not provided. Reason: Consistency with the generic code.
// Please use popped = heap[0] ; Remove(compar, heap, 0);
func Pop() {
	panic("The Pop is not provided. Use Remove(compar, heap, 0).")
}

// Push pushes the element x onto the heap.
// The compar is a compare function.
// The heap is a heapified slice.
// The elem element is a pointer to an element of the same type.
// The complexity is O(log(n)) where n = h.Len().
func Push( /*ts0 *[1]uintptr, */ compar func(*int32, *int32) int, heap *[]int32, elem *int32) {
	l := len(*heap)
	*heap = append(*heap, *elem)
	up( /*ts0, */ compar, *heap, l)
}

// Remove removes the element at index i from the heap.
// The compar is a compare function.
// The heap is a heapified slice.
// The complexity is O(log(n)) where n = h.Len().
func Remove( /*ts0 *[1]uintptr, */ compar func(*int32, *int32) int, heap *[]int32, i int) {

	n := len(*heap) - 1
	if n != i {
		{ // swap
			x := (*heap)[i]
			(*heap)[i] = (*heap)[n]
			(*heap)[n] = x
		}
		down( /*ts0, */ compar, (*heap), i, n)
		if i != 0 {
			up( /*ts0, */ compar, (*heap), i)
		}
	}
	(*heap) = (*heap)[:n]
}

// Another loads the second-top value to heap[1]
// The compar is a compare function.
// The heap is a heapified slice.
func Another( /*ts0 *[1]uintptr, */ compar func(*int32, *int32) int, heap []int32) {
	// first we check that [1] < [2]

	if len(heap) <= 2 || compar(&heap[1], &heap[2]) <= 0 {
		// ok
		return
	}

	{ // swap
		x := heap[1]
		heap[1] = heap[2]
		heap[2] = x
	}

	down( /*ts0, */ compar, heap, 2, len(heap)) //FIXME: shouldn't be len(heap)-1?
}

// Fix re-establishes the heap ordering after the element at index i has
// changed its value. Changing the value of the element at index i and then
// calling Fix is equivalent to, but less expensive than, calling Remove(h, i)
// followed by a Push of the new value.
// The compar is a compare function.
// The heap is a slice.
// The complexity is O(log(n)) where n = h.Len().
func Fix( /*ts0 *[1]uintptr, */ compar func(*int32, *int32) int, heap []int32, i int) {
	down( /*ts0, */ compar, heap, i, len(heap)) //FIXME: shouldn't be len(heap)-1?
	up( /*ts0, */ compar, heap, i)
}

// A heap must be initialized before any of the heap operations can be used.
// Heapify is idempotent with respect to the heap invariants and may be called
// whenever the heap invariants may have been invalidated.
// The compar is a compare function.
// Then heap is a source slice. Dst is a result slice. In place is supported.
// Its complexity is O(n) where n = h.Len().
func Heapify( /*ts0 *[1]uintptr, */ compar func(*int32, *int32) int, dst []int32, heap []int32) {
	n := len(heap)
	if &dst[0] == &heap[0] {
		for i := n/2 - 1; i >= 0; i-- {
			down( /*ts0, */ compar, heap, i, n)
		}
	} else {
		// FIXME: out of place heapify not implemented
		panic("FIXME: out of place heapify not implemented")
	}
}

func up( /*ts0 *[1]uintptr, */ compar func(*int32, *int32) int, heap []int32, j int) {
	for {
		i := (j - 1) / 2 // parent
		if i == j || compar(&heap[j], &heap[i]) >= 0 {
			break
		}
		{ // swap
			x := heap[i]
			heap[i] = heap[j]
			heap[j] = x
		}
		j = i
	}
}

func down( /*ts0 *[1]uintptr, */ compar func(*int32, *int32) int, heap []int32, i, n int) {
	for {
		j1 := 2*i + 1
		if j1 >= n || j1 < 0 { // j1 < 0 after int32 overflow
			break
		}
		j := j1 // left child
		if j2 := j1 + 1; j2 < n && compar(&heap[j1], &heap[j2]) >= 0 {
			j = j2 // = 2*i + 2  // right child
		}
		if compar(&heap[j], &heap[i]) >= 0 {
			break
		}
		{ // swap
			x := heap[i]
			heap[i] = heap[j]
			heap[j] = x
		}
		i = j
	}
}
