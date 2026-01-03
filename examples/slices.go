package main

// Slices — dynamic arrays: creation, indexing, slicing, append, copy,
// and capacity management.

import "fmt"

func main() {
	// Create slice from literal
	s := []int{1, 2, 3}
	fmt.Println("initial slice:", s, "len=", len(s), "cap=", cap(s))

	// Append
	s = append(s, 4)
	fmt.Println("after append:", s, "len=", len(s), "cap=", cap(s))

	// Slice a slice
	s2 := s[1:3]
	fmt.Println("slice[1:3]:", s2)

	// Make a slice with initial length and capacity
	s3 := make([]int, 3, 5)
	fmt.Println("make([]int, 3, 5):", s3, "len=", len(s3), "cap=", cap(s3))

	// Copy one slice to another
	s4 := make([]int, len(s))
	copy(s4, s)
	fmt.Println("copy(s4, s):", s4)

	// Modify slice element (reflects in original)
	s5 := s[:2]
	s5[0] = 99
	fmt.Println("after s5[0]=99, s:", s)
}
