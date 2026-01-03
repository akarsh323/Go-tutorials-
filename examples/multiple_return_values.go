package main

// Multiple Return Values — covers multiple returns, error handling idiom,
// named returns, and blank identifier usage.

import "fmt"

// Simple multiple returns
func divmod(a, b int) (int, int) {
	return a / b, a % b
}

// Multiple return with error (idiomatic Go)
func safeDivide(a, b int) (int, error) {
	if b == 0 {
		return 0, fmt.Errorf("division by zero")
	}
	return a / b, nil
}

// Named return values
func swap(x, y int) (swapped_x, swapped_y int) {
	swapped_x = y
	swapped_y = x
	return
}

func main() {
	q, r := divmod(7, 3)
	fmt.Println("divmod(7,3):", q, r)

	// Error handling
	result, err := safeDivide(10, 2)
	if err != nil {
		fmt.Println("error:", err)
	} else {
		fmt.Println("safeDivide(10,2):", result)
	}

	// Named returns
	a, b := swap(3, 7)
	fmt.Println("swap(3,7):", a, b)
}
