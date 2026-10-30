package main

import (
	"fmt"
)

type Slices[T any] struct {
	header   *T
	length   int
	capacity int
}

func myMake[T any](length, capacity int) *Slices[T] {
	if length > capacity || length < 0 || capacity < 0 {
		panic("invalid length or capacity")
	}

	if capacity == 0 {
		return &Slices[T]{
			header:   nil,
			length:   0,
			capacity: 0,
		}
	}

	// Allocate a slice to hold the raw backing array
	backing := make([]T, capacity)

	return &Slices[T]{
		header:   &backing[0],
		length:   length,
		capacity: capacity,
	}
}

func main() {
	s := myMake[int](3, 5)
	fmt.Printf("Header: %p, Len: %d, Cap: %d\n", s.header, s.length, s.capacity)
}
