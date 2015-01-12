// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package heap

import (
//	"reflect"
//	"unsafe"
)

// FIXME: bounded heap, deheap

// does not check that the array is indeed heap-ordered
func Push( /*ts0 *[1]uintptr, */ compar func(*int32, *int32) int, heap *[]int32, elem *int32) {
	l := len(*heap)
	*heap = append(*heap, *elem)
	up( /*ts0, */ compar, *heap, l)
}

// the return value shall not be ignored
// deletes item from the heap at position N
// pop is done by inspecting heap[0] and calling Remove(..,0)
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

// another loads the second smallest value to heap[1]
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

func Fix( /*ts0 *[1]uintptr, */ compar func(*int32, *int32) int, heap []int32, i int) {
	down( /*ts0, */ compar, heap, i, len(heap)) //FIXME: shouldn't be len(heap)-1?
	up( /*ts0, */ compar, heap, i)
}

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
