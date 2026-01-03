package main

// Conditions: If / else / else-if — covers if, else-if chains, nested conditions,
// short statements, and scope rules.

import "fmt"

func main() {
	x := 5
	// Basic if / else
	if x%2 == 0 {
		fmt.Println("even")
	} else {
		fmt.Println("odd")
	}

	// If with a short statement
	if v := x * 2; v > 8 {
		fmt.Println("x*2 > 8", v)
	} else {
		fmt.Println("x*2 <= 8", v)
	}

	// else-if chain
	grade := 85
	if grade >= 90 {
		fmt.Println("A")
	} else if grade >= 80 {
		fmt.Println("B")
	} else if grade >= 70 {
		fmt.Println("C")
	} else {
		fmt.Println("F")
	}

	// Nested conditions
	if x > 0 {
		if x < 10 {
			fmt.Println("x is between 0 and 10")
		}
	}
}
