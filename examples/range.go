package main

// Range — iteration over arrays, slices, maps, strings, and channels.
// Returns index (or key) and value.

import "fmt"

func main() {
	// Range over slice
	s := []string{"a", "b", "c"}
	fmt.Println("slice:")
	for i, v := range s {
		fmt.Printf("  [%d]=%s\n", i, v)
	}

	// Range with blank identifier (ignore index)
	fmt.Println("\nvalues only:")
	for _, v := range s {
		fmt.Println("  -", v)
	}

	// Range with index only
	fmt.Println("\nindexes only:")
	for i := range s {
		fmt.Println("  index:", i)
	}

	// Range over map
	m := map[string]int{"one": 1, "two": 2}
	fmt.Println("\nmap:")
	for k, v := range m {
		fmt.Printf("  %s -> %d\n", k, v)
	}

	// Range over string (iterates by rune)
	fmt.Println("\nstring (runes):")
	for i, r := range "hello" {
		fmt.Printf("  [%d]=%q\n", i, r)
	}
}
