package main

// Data Types — concise explanation of core types and zero values.

import "fmt"

func main() {
	var i int     // zero value 0
	var f float64 // zero value 0.0
	var s string  // zero value ""
	var b bool    // zero value false
	fmt.Printf("zero values -> int=%d float=%g string=%q bool=%v\n", i, f, s, b)

	// Short variable declarations with type inference
	i = 42
	f = 3.14
	s = "a string"
	b = true
	fmt.Printf("values -> %d %g %q %v\n", i, f, s, b)

	// Other useful types: slices, maps, structs, channels
	var sl []int         // nil slice
	var m map[string]int // nil map
	fmt.Println("nil slice len/cap:", len(sl), cap(sl), "nil map is nil:", m == nil)
}
