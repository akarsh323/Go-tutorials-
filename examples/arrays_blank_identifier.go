package main

// Arrays and Blank Identifier — teaching version.
// Shows array literals, indexing, iteration and the blank identifier when
// you want to ignore values returned by range.

import "fmt"

func main() {
	// Fixed-size array (length is part of the type)
	arr := [3]int{1, 2, 3}

	// Access by index
	fmt.Println("arr[0] =", arr[0])

	// Range returns index and value. Use the blank identifier '_' to ignore
	// a return value if it's not needed.
	fmt.Println("indexes:")
	for i := range arr { // shorthand when only index is needed
		fmt.Println(i)
	}

	fmt.Println("values:")
	for _, v := range arr { // ignore index, keep value
		fmt.Println(v)
	}

	// Arrays are values: assignment copies the entire array.
	b := arr
	b[0] = 99
	fmt.Println("original:", arr, "copied/modified:", b)
}
