package main

// Loop: For (using as while) — use 'for' as a while loop in Go.

import "fmt"

func main() {
	n := 3
	for n > 0 { // while-like loop
		fmt.Println("n=", n)
		n--
	}

	// Infinite loop example with break
	i := 0
	for {
		if i >= 3 {
			break
		}
		fmt.Println("loop", i)
		i++
	}
}
