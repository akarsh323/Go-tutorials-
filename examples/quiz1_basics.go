package main

// Quiz 1: Basics — test your understanding of basic concepts.
// Try predicting the output before running.

import "fmt"

func main() {
	fmt.Println("Quiz 1 Basics:")

	// Question 1: Arithmetic and type coercion
	fmt.Println("1. What is 1+2?", 1+2)
	fmt.Println("2. What is 10/3?", 10/3)
	fmt.Println("3. What is 10%3?", 10%3)

	// Question 2: Variable scope
	x := 5
	if x > 0 {
		x = 10
	}
	fmt.Println("4. What is x after if block?", x)

	// Question 3: Short variable declaration
	// y := y + 1 // compile error: y not declared
	fmt.Println("5. Can you use := with an already declared variable?", "No")
}
