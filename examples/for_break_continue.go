package main

// Loop: For (break/continue) — demonstrates control flow inside loops.

import "fmt"

func main() {
	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			// skip even numbers
			continue
		}
		if i > 7 {
			// stop the loop when i is greater than 7
			break
		}
		fmt.Println("odd", i)
	}

	// Labeled break/continue for nested loops
outer:
	for i := 1; i <= 3; i++ {
		for j := 1; j <= 3; j++ {
			if i*j > 4 {
				fmt.Println("breaking outer at", i, j)
				break outer
			}
			fmt.Println(i, j)
		}
	}
}
