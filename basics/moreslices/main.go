package main

import (
	"fmt"
	"slices"
)

func main() {
	// Slices are a more powerful and flexible way to work with sequences of data in Go.
	// They are built on top of arrays and provide dynamic sizing, making them more versatile than arrays.

	// Creating a slice using a slice literal
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Printf("Initial slice: %v length: %d, cap: %d\n", numbers, len(numbers), cap(numbers))

	s1 := numbers[2:5] // Slice from index 2 to 4 (5 is exclusive)
	fmt.Printf("\nSlice s1: %v length: %d, cap: %d\n", s1, len(s1), cap(s1))

	s2 := numbers[5:] // Slice from index 5 to the end
	fmt.Printf("\nSlice s2: %v length: %d, cap: %d\n", s2, len(s2), cap(s2))

	s3 := numbers[:4]
	fmt.Printf("\nSlice s3: %v length: %d, cap: %d\n", s3, len(s3), cap(s3))

	s4 := numbers[:] // Slice from the beginning to the end
	fmt.Printf("\nSlice s4: %v length: %d, cap: %d\n", s4, len(s4), cap(s4))

	contains := slices.Contains(s4, 8)
	fmt.Printf("\nSlice s4 contains 8: %v\n", contains)
}
