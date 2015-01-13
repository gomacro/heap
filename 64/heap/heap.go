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
func Push(ts0 *[1]uintptr, compar func(*uint64, *uint64) int, heap *[]uint64, elem []uint64) {
	incr := int((*ts0)[0])
	_ = incr

	l := (len(*heap) / incr)

	*heap = append(*heap, elem...)
	up(ts0, compar, *heap, l)
}

// the return value shall not be ignored
// deletes item from the heap at position N
// pop is done by inspecting heap[0] and calling Remove(..,0)
func Remove(ts0 *[1]uintptr, compar func(*uint64, *uint64) int, heap *[]uint64, i int) {
	incr := int((*ts0)[0])
	_ = incr

	n := (len(*heap) / incr) - 1
	if n != i {
		for q := 0; q < incr; q++ { // swap
			x := (*heap)[i*incr+q]
			(*heap)[i*incr+q] = (*heap)[n*incr+q]
			(*heap)[n*incr+q] = x
		}
		down(ts0, compar, (*heap), i, n)
		if i != 0 {
			up(ts0, compar, (*heap), i)
		}
	}
	(*heap) = (*heap)[:n*incr]
}

// another loads the second smallest value to heap[1]
func Another(ts0 *[1]uintptr, compar func(*uint64, *uint64) int, heap []uint64) {
	incr := int((*ts0)[0])
	_ = incr

	// first we check that [1] < [2]

	if (len(heap)/incr) <= 2 || compar(&heap[1], &heap[2]) <= 0 {
		// ok
		return
	}
	for q := 0; q < incr; q++ { // swap
		x := heap[1*incr+q]
		heap[1*incr+q] = heap[2*incr+q]
		heap[2*incr+q] = x
	}

	down(ts0, compar, heap, 2, (len(heap) / incr)) //FIXME: shouldn't be len(heap)-1?
}

func Fix(ts0 *[1]uintptr, compar func(*uint64, *uint64) int, heap []uint64, i int) {
	incr := int((*ts0)[0])
	_ = incr

	down(ts0, compar, heap, i, (len(heap) / incr)) //FIXME: shouldn't be len(heap)-1?
	up(ts0, compar, heap, i)
}

func Heapify(ts0 *[1]uintptr, compar func(*uint64, *uint64) int, dst []uint64, heap []uint64) {
	incr := int((*ts0)[0])
	_ = incr

	n := (len(heap) / incr)
	if &dst[0] == &heap[0] {
		for i := n/2 - 1; i >= 0; i-- {
			down(ts0, compar, heap, i, n)
		}
	} else {
		// FIXME: out of place heapify not implemented
		panic("FIXME: out of place heapify not implemented")
	}
}

func up(ts0 *[1]uintptr, compar func(*uint64, *uint64) int, heap []uint64, j int) {
	incr := int((*ts0)[0])
	_ = incr

	for {
		i := (j - 1) / 2 // parent
		if i == j || compar(&heap[j*incr], &heap[i*incr]) >= 0 {
			break
		}
		for q := 0; q < incr; q++ { // swap
			x := heap[i*incr+q]
			heap[i*incr+q] = heap[j*incr+q]
			heap[j*incr+q] = x
		}
		j = i
	}
}

func down(ts0 *[1]uintptr, compar func(*uint64, *uint64) int, heap []uint64, i, n int) {
	incr := int((*ts0)[0])
	_ = incr

	for {
		j1 := 2*i + 1
		if j1 >= n || j1 < 0 { // j1 < 0 after uint64 overflow
			break
		}
		j := j1 // left child
		if j2 := j1 + 1; j2 < n && compar(&heap[j1*incr], &heap[j2*incr]) >= 0 {
			j = j2 // = 2*i + 2  // right child
		}
		if compar(&heap[j*incr], &heap[i*incr]) >= 0 {
			break
		}
		for q := 0; q < incr; q++ { // swap
			x := heap[i*incr+q]
			heap[i*incr+q] = heap[j*incr+q]
			heap[j*incr+q] = x
		}
		i = j
	}
}
