package main

import "fmt"

/*
Topic 58: RECURSION

CONCEPT: A function calling itself to solve a problem by breaking it into
smaller subproblems of the same type.

ESSENTIAL COMPONENTS:
1. BASE CASE: The stopping condition (return without recursion)
2. RECURSIVE CASE: Call itself with simpler input

CALL STACK EXAMPLE (factorial(5)):
  factorial(5) → 5 * factorial(4)
                 → 4 * factorial(3)
                 → 3 * factorial(2)
                 → 2 * factorial(1)
                 → BASE: return 1
                 ← 2*1=2, 3*2=6, 4*6=24, 5*24=120

USE CASES:
✓ Tree/graph traversal, divide and conquer algorithms
✓ Mathematical problems (factorial, Fibonacci)
✓ Backtracking problems (maze solving)

PROS: Elegant for naturally recursive problems
CONS: Stack overflow risk, less efficient than iteration
*/

// Example 1: Factorial
func factorial(n int) int {
	if n == 0 {
		return 1 // BASE CASE
	}
	return n * factorial(n-1) // RECURSIVE CASE
}

// Example 2: Sum of digits
func sumOfDigits(n int) int {
	if n < 0 {
		n = -n
	}
	if n < 10 {
		return n // BASE CASE
	}
	return (n % 10) + sumOfDigits(n/10) // RECURSIVE CASE
}

// Example 3: Fibonacci (naive - exponential time)
func fibonacci(n int) int {
	if n <= 1 {
		return n // BASE CASES
	}
	return fibonacci(n-1) + fibonacci(n-2) // RECURSIVE CASE
}

// Example 4: Fibonacci with memoization (optimized)
func fibMemo(n int, memo map[int]int) int {
	if val, exists := memo[n]; exists {
		return val // Use cached result
	}
	if n <= 1 {
		return n // BASE CASE
	}
	result := fibMemo(n-1, memo) + fibMemo(n-2, memo)
	memo[n] = result
	return result
}

// Example 5: Power function
func power(base, exp int) int {
	if exp == 0 {
		return 1 // BASE CASE
	}
	return base * power(base, exp-1) // RECURSIVE CASE
}

// Example 6: String reversal
func reverseString(s string) string {
	if len(s) <= 1 {
		return s // BASE CASE
	}
	return string(s[len(s)-1]) + reverseString(s[:len(s)-1]) // RECURSIVE CASE
}

func main() {
	fmt.Println("═══════════════════════════════════════════════════════════")
	fmt.Println("EXAMPLE 1: Factorial")
	fmt.Println("═══════════════════════════════════════════════════════════\n")

	fmt.Println("Calculating factorial(5):")
	result := factorial(5)
	fmt.Printf("5! = %d\n\n", result)

	for i := 0; i <= 10; i++ {
		fmt.Printf("factorial(%d) = %d\n", i, factorial(i))
	}
	fmt.Println()

	// ─────────────────────────────────────────────────────────────────────────

	fmt.Println("═══════════════════════════════════════════════════════════")
	fmt.Println("EXAMPLE 2: Sum of Digits")
	fmt.Println("═══════════════════════════════════════════════════════════\n")

	nums := []int{9, 12, 123, 12345, 9876}
	for _, num := range nums {
		sum := sumOfDigits(num)
		fmt.Printf("sumOfDigits(%d) = %d\n", num, sum)
	}
	fmt.Println()

	// ─────────────────────────────────────────────────────────────────────────

	fmt.Println("═══════════════════════════════════════════════════════════")
	fmt.Println("EXAMPLE 3: Fibonacci (Naive - Slow)")
	fmt.Println("═══════════════════════════════════════════════════════════\n")

	fmt.Println("Naive Fibonacci (exponential time):")
	for i := 0; i <= 10; i++ {
		fmt.Printf("fibonacci(%d) = %d\n", i, fibonacci(i))
	}
	fmt.Printf("\nfibonacci(20) = %d (still fast)\n", fibonacci(20))
	fmt.Println("⚠️  fibonacci(30+) would be very slow!\n")

	// ─────────────────────────────────────────────────────────────────────────

	fmt.Println("═══════════════════════════════════════════════════════════")
	fmt.Println("EXAMPLE 4: Fibonacci with Memoization (Fast)")
	fmt.Println("═══════════════════════════════════════════════════════════\n")

	memo := make(map[int]int)
	fmt.Println("Memoized Fibonacci (linear time):")
	for i := 0; i <= 20; i++ {
		fmt.Printf("fibonacci(%d) = %d\n", i, fibMemo(i, memo))
	}
	fmt.Printf("\nfibonacci(30) = %d (very fast!)\n", fibMemo(30, make(map[int]int)))
	fmt.Println("✓ Memoization caches results\n")

	// ─────────────────────────────────────────────────────────────────────────

	fmt.Println("═══════════════════════════════════════════════════════════")
	fmt.Println("EXAMPLE 5: Power Function")
	fmt.Println("═══════════════════════════════════════════════════════════\n")

	fmt.Println("Powers of 2:")
	for i := 0; i <= 10; i++ {
		fmt.Printf("2^%d = %d\n", i, power(2, i))
	}
	fmt.Printf("\n3^4 = %d\n", power(3, 4))
	fmt.Printf("5^3 = %d\n\n", power(5, 3))

	// ─────────────────────────────────────────────────────────────────────────

	fmt.Println("═══════════════════════════════════════════════════════════")
	fmt.Println("EXAMPLE 6: String Reversal")
	fmt.Println("═══════════════════════════════════════════════════════════\n")

	words := []string{"hello", "Go", "recursion", "world"}
	for _, word := range words {
		reversed := reverseString(word)
		fmt.Printf("reverseString(\"%s\") = \"%s\"\n", word, reversed)
	}
	fmt.Println()

	// ─────────────────────────────────────────────────────────────────────────

	fmt.Println("═══════════════════════════════════════════════════════════")
	fmt.Println("KEY CONCEPTS & BEST PRACTICES")
	fmt.Println("═══════════════════════════════════════════════════════════\n")

	fmt.Println(`
ESSENTIAL COMPONENTS:
1. BASE CASE: Without it → infinite recursion → stack overflow
2. RECURSIVE CASE: Must make progress towards base case

HOW IT WORKS:
- Each function call adds to call stack
- When base case reached, unwind and combine results
- Deep recursion can cause stack overflow

PERFORMANCE:
- Naive recursion can be exponential time (slow!)
- Memoization converts exponential → linear time
- Iteration is usually more efficient than recursion

WHEN TO USE RECURSION:
✓ Tree/graph traversal (DFS, inorder, postorder)
✓ Divide and conquer (merge sort, quicksort)
✓ Natural recursive structures
✓ Backtracking problems
✗ Very deep recursion (use iteration instead)
✗ Without careful base case handling

BEST PRACTICES:
✓ Always define a reachable base case
✓ Test with edge cases (empty, single element, large input)
✓ Use memoization to optimize slow recursive solutions
✓ Consider stack depth limits for large inputs
✓ Measure performance - recursion isn't always optimal!
✓ Document the base case and recursive case clearly

OPTIMIZATION TECHNIQUES:
- Memoization: Cache results to avoid redundant calculations
- Tail recursion: Some languages optimize, Go doesn't
- Convert to iteration: More memory efficient
	`)
}
