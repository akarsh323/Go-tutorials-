package main

// Variadic Functions — functions that accept a variable number of arguments.

import "fmt"

// Variadic function accepts any number of ints
func sum(nums ...int) int {
	total := 0
	for _, n := range nums {
		total += n
	}
	return total
}

// Variadic with multiple types
func printAll(args ...interface{}) {
	for i, arg := range args {
		fmt.Printf("  [%d] %v (%T)\n", i, arg, arg)
	}
}

// Passing slice to variadic
func main() {
	// Call variadic with multiple args
	fmt.Println("sum(1,2,3,4) =", sum(1, 2, 3, 4))
	fmt.Println("sum() =", sum())

	// Pass a slice to variadic with ...
	slice := []int{5, 6, 7}
	fmt.Println("sum(5,6,7 from slice) =", sum(slice...))

	// Variadic with interface{}
	fmt.Println("\nusing interface{} variadic:")
	printAll(42, "hello", 3.14, true)
}
