// Copyright 2015 The GOMACRO Authors. All rights reserved.
// Use of this source code is governed by a GPLv2-style
// license that can be found in the LICENSE file.

// Package heap provides a heap (a priority queue) operations on a slice.
package heap

import (
	heap32 "github.com/gomacro/heap/32/heap"
	heap64 "github.com/gomacro/heap/64/heap"
	heap8 "github.com/gomacro/heap/8/heap"
	"reflect"
	"unsafe"
)

////////////////////////////////////////////////////////////////////////////////
func elemsize(slice interface{}) uintptr {
	return uintptr(reflect.TypeOf(slice).Elem().Size())
}
func elemsize2(slice interface{}) uintptr {
	return uintptr(reflect.TypeOf(slice).Elem().Elem().Size())
}
func mvetype(dst, src *interface{}) {
	*(*uintptr)(unsafe.Pointer(dst)) = *(*uintptr)(unsafe.Pointer(src))
}
func arg8(fun interface{}) (dst func(*uint8, *uint8) int) {
	var ction interface{}
	ction = dst
	mvetype(&fun, &ction)
	return fun.(func(*uint8, *uint8) int)
}
func arg32(fun interface{}) (dst func(*uint32, *uint32) int) {
	var ction interface{}
	ction = dst
	mvetype(&fun, &ction)
	return fun.(func(*uint32, *uint32) int)
}
func arg64(fun interface{}) (dst func(*uint64, *uint64) int) {
	var ction interface{}
	ction = dst
	mvetype(&fun, &ction)
	return fun.(func(*uint64, *uint64) int)
}
func u8(slice interface{}, size uintptr) (src []uint8) {
	var dst interface{}
	dst = src
	mvetype(&slice, &dst)
	src = slice.([]uint8)
	h := (*reflect.SliceHeader)(unsafe.Pointer(&src))
	h.Len *= int(size)
	h.Cap *= int(size)
	return src
}
func u32(slice interface{}, size uintptr) (src []uint32) {
	var dst interface{}
	dst = src
	mvetype(&slice, &dst)
	src = slice.([]uint32)
	h := (*reflect.SliceHeader)(unsafe.Pointer(&src))
	h.Len *= int(size)
	h.Cap *= int(size)
	return src
}
func u64(slice interface{}, size uintptr) (src []uint64) {
	var dst interface{}
	dst = src
	mvetype(&slice, &dst)
	src = slice.([]uint64)
	h := (*reflect.SliceHeader)(unsafe.Pointer(&src))
	h.Len *= int(size)
	h.Cap *= int(size)
	return src
}
func pu8(pointer interface{}, size uintptr) (src []uint8) {
	x := reflect.ValueOf(pointer).Pointer()
	h := (*reflect.SliceHeader)(unsafe.Pointer(&src))
	h.Data = x
	h.Len = int(size)
	h.Cap = int(size)
	return src
}
func pu32(pointer interface{}, size uintptr) (src []uint32) {
	x := reflect.ValueOf(pointer).Pointer()
	h := (*reflect.SliceHeader)(unsafe.Pointer(&src))
	h.Data = x
	h.Len = int(size)
	h.Cap = int(size)
	return src
}
func pu64(pointer interface{}, size uintptr) (src []uint64) {
	x := reflect.ValueOf(pointer).Pointer()
	h := (*reflect.SliceHeader)(unsafe.Pointer(&src))
	h.Data = x
	h.Len = int(size)
	h.Cap = int(size)
	return src
}
func su8(slicepointer interface{}, size uintptr) (u []uint8, v *reflect.SliceHeader) {
	var src *[]uint8
	var dst interface{}
	dst = src
	mvetype(&slicepointer, &dst)
	w := slicepointer.(*[]uint8)
	v = (*reflect.SliceHeader)(unsafe.Pointer(w))
	u = *(w)
	h := (*reflect.SliceHeader)(unsafe.Pointer(&u))
	h.Len *= int(size)
	h.Cap *= int(size)
	return u, v
}
func su32(slicepointer interface{}, size uintptr) (u []uint32, v *reflect.SliceHeader) {
	var src *[]uint32
	var dst interface{}
	dst = src
	mvetype(&slicepointer, &dst)
	w := slicepointer.(*[]uint32)
	v = (*reflect.SliceHeader)(unsafe.Pointer(w))
	u = *(w)
	h := (*reflect.SliceHeader)(unsafe.Pointer(&u))
	h.Len *= int(size)
	h.Cap *= int(size)
	return u, v
}
func su64(slicepointer interface{}, size uintptr) (u []uint64, v *reflect.SliceHeader) {
	var src *[]uint64
	var dst interface{}
	dst = src
	mvetype(&slicepointer, &dst)
	w := slicepointer.(*[]uint64)
	v = (*reflect.SliceHeader)(unsafe.Pointer(w))
	u = *(w)
	h := (*reflect.SliceHeader)(unsafe.Pointer(&u))
	h.Len *= int(size)
	h.Cap *= int(size)
	return u, v
}
func fu8(u []uint8, v *reflect.SliceHeader, size uintptr) {
	if len(u) != 0 {
		v.Data = uintptr(unsafe.Pointer(&u[0]))
	}
	v.Len = len(u) / int(size)
	v.Cap = cap(u) / int(size)
}
func fu32(u []uint32, v *reflect.SliceHeader, size uintptr) {
	if len(u) != 0 {
		v.Data = uintptr(unsafe.Pointer(&u[0]))
	}
	v.Len = len(u) / int(size)
	v.Cap = cap(u) / int(size)
}
func fu64(u []uint64, v *reflect.SliceHeader, size uintptr) {
	if len(u) != 0 {
		v.Data = uintptr(unsafe.Pointer(&u[0]))
	}
	v.Len = len(u) / int(size)
	v.Cap = cap(u) / int(size)
}

////////////////////////////////////////////////////////////////////////////////

// Pop is not provided. Reason: We do not return generic pointer.
// Please use popped = heap[0] ; Remove(compar, heap, 0);
func Pop() {
	panic("The Pop is not provided. Use Remove(compar, heap, 0).")
}

// Push pushes the element x onto the heap.
// The compar is a compare function.
// The heap is a heapified slice.
// The elem element is a pointer to an element of the same type.
// The complexity is O(log(n)) where n = h.Len().
func Push(compar interface{}, heap interface{}, elem interface{}) {

	// OK
	size := elemsize2(heap) //8,4,1
	//	fmt.Println("ELEM SIZE:", size)

	if (size & 7) == 0 { // use 8 (64bit)
		var m = [1]uintptr{size / 8}
		uheap, fheap := su64(heap, m[0])

		heap64.Push(&m, arg64(compar), &uheap, pu64(elem, m[0]))
		fu64(uheap, fheap, m[0])
		return
	}
	if (size & 3) == 0 { // use 4 (32bit)
		var m = [1]uintptr{size / 4}
		uheap, fheap := su32(heap, m[0])

		heap32.Push(&m, arg32(compar), &uheap, pu32(elem, m[0]))
		fu32(uheap, fheap, m[0])
		return
	}

	// use 1 (8bit)
	var m = [1]uintptr{size}
	uheap, fheap := su8(heap, m[0])

	heap8.Push(&m, arg8(compar), &uheap, pu8(elem, m[0]))
	fu8(uheap, fheap, m[0])

}

// A heap must be initialized before any of the heap operations can be used.
// Heapify is idempotent with respect to the heap invariants and may be called
// whenever the heap invariants may have been invalidated.
// The compar is a compare function.
// Then heap is a source slice. Dst is a result slice. In place is supported.
// Its complexity is O(n) where n = h.Len().
func Heapify(compar interface{}, dst interface{}, heap interface{}) {

	// OK
	size := elemsize(heap) //8,4,1
	//	fmt.Println("ELEM SIZE:", size)

	if (size & 7) == 0 { // use 8 (64bit)
		var m = [1]uintptr{size / 8}
		heap64.Heapify(&m, arg64(compar), u64(dst, m[0]), u64(heap, m[0]))
		return
	}
	if (size & 3) == 0 { // use 4 (32bit)
		var m = [1]uintptr{size / 4}
		heap32.Heapify(&m, arg32(compar), u32(dst, m[0]), u32(heap, m[0]))
		return
	}

	// use 1 (8bit)
	var m = [1]uintptr{size}
	heap8.Heapify(&m, arg8(compar), u8(dst, m[0]), u8(heap, m[0]))
	return

}

// Remove removes the element at index i from the heap.
// The compar is a compare function.
// The heap is a heapified slice.
// The complexity is O(log(n)) where n = h.Len().
func Remove(compar interface{}, heap interface{}, i int) {
	// OK
	size := elemsize2(heap) //8,4,1
	//	fmt.Println("ELEM SIZE:", size)

	if (size & 7) == 0 { // use 8 (64bit)
		var m = [1]uintptr{size / 8}
		uheap, fheap := su64(heap, m[0])

		heap64.Remove(&m, arg64(compar), &uheap, i)
		fu64(uheap, fheap, m[0])
		return
	}
	if (size & 3) == 0 { // use 4 (32bit)
		var m = [1]uintptr{size / 4}
		uheap, fheap := su32(heap, m[0])

		heap32.Remove(&m, arg32(compar), &uheap, i)
		fu32(uheap, fheap, m[0])
		return
	}

	// use 1 (8bit)
	var m = [1]uintptr{size}
	uheap, fheap := su8(heap, m[0])

	heap8.Remove(&m, arg8(compar), &uheap, i)
	fu8(uheap, fheap, m[0])
	return

}

// Fix re-establishes the heap ordering after the element at index i has
// changed its value. Changing the value of the element at index i and then
// calling Fix is equivalent to, but less expensive than, calling Remove(h, i)
// followed by a Push of the new value.
// The compar is a compare function.
// The heap is a slice.
// The complexity is O(log(n)) where n = h.Len().
func Fix(compar interface{}, heap interface{}, i int) {
	size := elemsize(heap) //8,4,1
	//	fmt.Println("ELEM SIZE:", size)

	if (size & 7) == 0 { // use 8 (64bit)
		var m = [1]uintptr{size / 8}
		heap64.Fix(&m, arg64(compar), u64(heap, m[0]), i)
		return
	}
	if (size & 3) == 0 { // use 4 (32bit)
		var m = [1]uintptr{size / 4}
		heap32.Fix(&m, arg32(compar), u32(heap, m[0]), i)
		return
	}

	// use 1 (8bit)
	var m = [1]uintptr{size}
	heap8.Fix(&m, arg8(compar), u8(heap, m[0]), i)
	return

}

// Another loads the second-top value to heap[1]
// The compar is a compare function.
// The heap is a heapified slice.
func Another(compar interface{}, heap interface{}) {
	size := elemsize(heap) //8,4,1
	//	fmt.Println("ELEM SIZE:", size)

	if (size & 7) == 0 { // use 8 (64bit)
		var m = [1]uintptr{size / 8}
		heap64.Another(&m, arg64(compar), u64(heap, m[0]))
		return
	}

	if (size & 3) == 0 { // use 4 (32bit)
		var m = [1]uintptr{size / 4}
		heap32.Another(&m, arg32(compar), u32(heap, m[0]))
		return
	}

	// use 1 (8bit)
	var m = [1]uintptr{size}
	heap8.Another(&m, arg8(compar), u8(heap, m[0]))
	return

}
