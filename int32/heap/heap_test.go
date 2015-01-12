// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package heap

import (
	"math/rand"
	"testing"
)

var NULL = [1]uintptr{uintptr(1)}

func Int32(a, b *int32) int {
	r := int(*a>>16) - int(*b>>16)
	if r != 0 {
		return r
	}
	return int(*a) - int(*b)
}

func Bag(a, b *byte) int {
	return 0x1EE7
}

type myHeap []int32

func (h *myHeap) Less(i, j int) bool {
	return (*h)[i] < (*h)[j]
}

func (h *myHeap) Swap(i, j int) {
	(*h)[i], (*h)[j] = (*h)[j], (*h)[i]
}

func (h *myHeap) Len() int {
	return len(*h)
}

func (h *myHeap) Pop() (v interface{}) {
	*h, v = (*h)[:h.Len()-1], (*h)[h.Len()-1]
	return
}

func (h *myHeap) Push(v interface{}) {
	*h = append(*h, v.(int32))
}

func (h myHeap) verify(t *testing.T, i int) {
	n := len(h)
	j1 := 2*i + 1
	j2 := 2*i + 2
	if j1 < n {
		if h.Less(j1, i) {
			t.Errorf("heap invariant invalidated [%d] = %d > [%d] = %d", i, h[i], j1, h[j1])
			return
		}
		h.verify(t, j1)
	}
	if j2 < n {
		if h.Less(j2, i) {
			t.Errorf("heap invariant invalidated [%d] = %d > [%d] = %d", i, h[i], j1, h[j2])
			return
		}
		h.verify(t, j2)
	}
}

func TestInit0(t *testing.T) {
	h := []int32{}
	for i := 20; i > 0; i-- {
		n := int32(0)
		h = Push( /*&NULL, */ Int32, h, &n) /*TYPECAST*/ // all elements are the same
	}
	Heapify( /*&NULL, */ Int32, h, h)
	myHeap(h).verify(t, 0)

	for i := 1; len(h) > 0; i++ {
		x := h[0]
		h = Remove( /*&NULL, */ Int32, h, 0) /*TYPECAST*/
		myHeap(h).verify(t, 0)
		if x != 0 {
			t.Errorf("%d.th pop got %d; want %d", i, x, 0)
		}
	}
}

func TestInit1(t *testing.T) {
	h := []int32{}
	for i := int32(20); i > 0; i-- {
		h = Push( /*&NULL, */ Int32, h, &i) /*TYPECAST*/ // all elements are different
	}
	Heapify( /*&NULL, */ Int32, h, h)
	myHeap(h).verify(t, 0)

	for i := int32(1); len(h) > 0; i++ {
		x := h[0]
		h = Remove( /*&NULL, */ Int32, h, 0) /*TYPECAST*/
		myHeap(h).verify(t, 0)
		if x != i {
			t.Errorf("%d.th pop got %d; want %d", i, x, i)
		}
	}
}

func Test(t *testing.T) {
	h := []int32{}
	myHeap(h).verify(t, 0)

	for i := int32(20); i > 10; i-- {
		h = Push( /*&NULL, */ Int32, h, &i) /*TYPECAST*/ // all elements are different
	}
	Heapify( /*&NULL, */ Int32, h, h)
	myHeap(h).verify(t, 0)

	for i := int32(10); i > 0; i-- {
		h = Push( /*&NULL, */ Int32, h, &i) /*TYPECAST*/ // all elements are different
		myHeap(h).verify(t, 0)
	}

	for i := int32(1); len(h) > 0; i++ {
		x := h[0]
		h = Remove( /*&NULL, */ Int32, h, 0) /*TYPECAST*/
		if i < 20 {
			j := 20 + i
			h = Push( /*&NULL, */ Int32, h, &j) /*TYPECAST*/ // all elements are different
		}
		myHeap(h).verify(t, 0)
		if x != i {
			t.Errorf("%d.th pop got %d; want %d", i, x, i)
		}
	}
}
func TestRemove0(t *testing.T) {
	h := []int32{}

	for i := int32(0); i < 10; i++ {
		h = Push( /*&NULL, */ Int32, h, &i) /*TYPECAST*/
	}

	myHeap(h).verify(t, 0)

	for len(h) > 0 {
		i := len(h) - 1

		x := h[i]
		h = Remove( /*&NULL, */ Int32, h, i) /*TYPECAST*/
		if x != int32(i) {
			t.Errorf("Remove(%d) got %d; want %d", i, x, i)
		}
		myHeap(h).verify(t, 0)
	}

}

func TestRemove1(t *testing.T) {
	h := []int32{}

	for i := int32(0); i < 10; i++ {
		h = Push( /*&NULL, */ Int32, h, &i) /*TYPECAST*/
	}

	myHeap(h).verify(t, 0)

	for i := int32(0); len(h) > 0; i++ {
		x := h[0]
		h = Remove( /*&NULL, */ Int32, h, 0) /*TYPECAST*/
		if x != i {
			t.Errorf("Remove(0) got %d; want %d", x, i)
		}
		myHeap(h).verify(t, 0)
	}
}
func TestRemove2(t *testing.T) {
	N := 10

	h := []int32{}
	for i := int32(0); i < int32(N); i++ {
		h = Push( /*&NULL, */ Int32, h, &i)
	}
	myHeap(h).verify(t, 0)

	m := make(map[int32]bool)
	for len(h) > 0 {
		i := int32((len(h) - 1) / 2)
		x := h[i]
		h = Remove( /*&NULL, */ Int32, h, int(i))
		m[x] = true
		myHeap(h).verify(t, 0)
	}

	if len(m) != N {
		t.Errorf("len(m) = %d; want %d", len(m), N)
	}
	for i := int32(0); i < int32(len(m)); i++ {
		if !m[i] {
			t.Errorf("m[%d] doesn't exist", i)
		}
	}
}

func BenchmarkDup(b *testing.B) {
	const n = 10000
	h := make([]int32, n)
	for i := 0; i < b.N; i++ {
		for j := 0; j < n; j++ {
			var zero int32 = 0
			h = Push( /*&NULL, */ Int32, h, &zero) // all elements are the same
		}
		for len(h) > 0 {
			Remove(Int32, h, 0)
		}
	}
}

func TestFix(t *testing.T) {
	h := []int32{}
	myHeap(h).verify(t, 0)

	for i := int32(200); i > 0; i -= 10 {
		h = Push( /*&NULL, */ Int32, h, &i) /*TYPECAST*/
	}
	myHeap(h).verify(t, 0)

	if h[0] != 10 {
		t.Fatalf("Expected head to be 10, was %d", h[0])
	}

	h[0] = 210
	Fix( /*&NULL, */ Int32, h, 0)
	myHeap(h).verify(t, 0)

	for i := int32(100); i > 0; i-- {
		elem := rand.Intn(len(h))
		if i&1 == 0 {
			h[elem] *= 2
		} else {
			h[elem] /= 2
		}
		Fix( /*&NULL, */ Int32, h, elem)
		myHeap(h).verify(t, 0)
	}
}

func TestAnother0(t *testing.T) {
	q := []int32{0, 10, 100, 11, 12, 101, 102}
	m := []int32{0, 100, 10, 101, 102, 11, 12}
	h := make([]int32, len(m))
	copy(h, m)

	myHeap(q).verify(t, 0)
	myHeap(m).verify(t, 0)
	myHeap(h).verify(t, 0)

	Another( /*&NULL, */ Int32, h)

	if h[1] != 10 {
		t.Errorf("Has %v", h)
	}

	myHeap(h).verify(t, 0)

}
