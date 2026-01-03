package main

import (
	"fmt"
)

/*
Topic 58: RECURSION - A Comprehensive Guide

═══════════════════════════════════════════════════════════════════════════════

DEFINITION:
Recursion is a fundamental programming concept where a function calls itself
directly or indirectly to solve a problem.

KEY CONCEPT:
Recursion breaks down a problem into smaller subproblems of the same type
until they become simple enough to solve directly.

ESSENTIAL COMPONENTS:

1. BASE CASE:
   - A condition where the function STOPS calling itself
   - Returns a specific value without further recursion
   - Without a base case: INFINITE RECURSION → STACK OVERFLOW

2. RECURSIVE CASE:
   - The function calls itself with a smaller or simpler input
   - Makes progress towards the base case
   - Each call should bring us closer to the base case

HOW RECURSION WORKS:
When a function calls itself, a new function frame is added to the call stack.
The execution is paused, the new function starts executing. When that function
finishes, the result is returned and the previous function continues.

CALL STACK EXAMPLE (factorial):
  factorial(5)
    → 5 * factorial(4)
      → 4 * factorial(3)
        → 3 * factorial(2)
          → 2 * factorial(1)
            → BASE CASE: return 1
          ← return 2 * 1 = 2
        ← return 3 * 2 = 6
      ← return 4 * 6 = 24
    ← return 5 * 24 = 120

PRACTICAL USE CASES:
1. Mathematical Algorithms: Factorial, Fibonacci, permutations
2. Tree/Graph Traversal: DFS, preorder, inorder, postorder traversals
3. Divide and Conquer: Merge sort, quicksort, binary search
4. Backtracking: Maze solving, N-queens problem
5. Dynamic Programming: Problems with overlapping subproblems

ADVANTAGES:
- Simplifies implementation of complex problems
- Code mirrors the problem statement
- Elegant solutions for naturally recursive structures
- Easier to understand for recursive problems

DISADVANTAGES:
- Less efficient than iterative solutions (function call overhead)
- Risk of stack overflow with deep recursion
- Higher memory usage due to call stack
- Slower execution compared to loops

CONSIDERATIONS:
- Must have a reachable base case to avoid infinite recursion
- Be careful with large inputs (memory and performance)
- Use iterative approaches for very deep recursion
- Consider memoization to cache results
- Test with multiple input sizes and edge cases

═══════════════════════════════════════════════════════════════════════════════
*/

// ─────────────────────────────────────────────────────────────────────────────
// EXAMPLE 1: FACTORIAL
// ─────────────────────────────────────────────────────────────────────────────

// factorial calculates n! (n factorial)
// Base case: factorial(0) = 1
// Recursive case: factorial(n) = n * factorial(n-1)
func factorial(n int) int {
	fmt.Printf("Calling factorial(%d)\n", n)

	// BASE CASE: When n reaches 0, stop recursion
	if n == 0 {
		fmt.Println("  → BASE CASE reached: returning 1")
		return 1
	}

	// RECURSIVE CASE: Call factorial with n-1
	result := n * factorial(n-1)
	fmt.Printf("  ← factorial(%d) returning %d\n", n, result)
	return result
}

// ─────────────────────────────────────────────────────────────────────────────
// EXAMPLE 2: SUM OF DIGITS
// ─────────────────────────────────────────────────────────────────────────────

// sumOfDigits calculates the sum of all digits in a number
// Example: sumOfDigits(12345) = 1+2+3+4+5 = 15
// Base case: if n < 10, return n
// Recursive case: (n % 10) + sumOfDigits(n / 10)
func sumOfDigits(n int) int {
	if n < 0 {
		n = -n // Handle negative numbers
	}

	// BASE CASE: Single digit numbers
	if n < 10 {
		fmt.Printf("  → BASE CASE: digit %d\n", n)
		return n
	}

	// RECURSIVE CASE: last digit + sum of remaining digits
	lastDigit := n % 10
	remainingSum := sumOfDigits(n / 10)
	total := lastDigit + remainingSum

	fmt.Printf("  → sumOfDigits(%d): %d + sumOfDigits(%d) = %d\n",
		n, lastDigit, n/10, total)

	return total
}

// ─────────────────────────────────────────────────────────────────────────────
// EXAMPLE 3: FIBONACCI SEQUENCE
// ─────────────────────────────────────────────────────────────────────────────

// fibonacci returns the nth Fibonacci number
// Base cases: fib(0)=0, fib(1)=1
// Recursive case: fib(n) = fib(n-1) + fib(n-2)
// WARNING: This naive implementation is VERY slow for large n (exponential time)
func fibonacci(n int) int {
	// BASE CASES
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}

	// RECURSIVE CASE: Sum of previous two Fibonacci numbers
	return fibonacci(n-1) + fibonacci(n-2)
}

// ─────────────────────────────────────────────────────────────────────────────
// EXAMPLE 4: FIBONACCI WITH MEMOIZATION (Optimized)
// ─────────────────────────────────────────────────────────────────────────────

// fibonacciMemo returns the nth Fibonacci number using memoization
// Memoization: Caching results of expensive function calls
// This reduces exponential time to linear time!
func fibonacciMemo(n int, memo map[int]int) int {
	// Check if result is already cached
	if val, exists := memo[n]; exists {
		fmt.Printf("  → Using cached value for fib(%d) = %d\n", n, val)
		return val
	}

	// BASE CASES
	if n == 0 {
		memo[n] = 0
		return 0
	}
	if n == 1 {
		memo[n] = 1
		return 1
	}

	// RECURSIVE CASE with memoization
	result := fibonacciMemo(n-1, memo) + fibonacciMemo(n-2, memo)
	memo[n] = result
	return result
}

// ─────────────────────────────────────────────────────────────────────────────
// EXAMPLE 5: POWER FUNCTION
// ─────────────────────────────────────────────────────────────────────────────

// power calculates base^exponent recursively
// Base case: base^0 = 1
// Recursive case: base^n = base * base^(n-1)
func power(base int, exponent int) int {
	// BASE CASE: Any number to the power of 0 is 1
	if exponent == 0 {
		return 1
	}

	// RECURSIVE CASE
	return base * power(base, exponent-1)
}

// ─────────────────────────────────────────────────────────────────────────────
// EXAMPLE 6: STRING REVERSAL
// ─────────────────────────────────────────────────────────────────────────────

// reverseString recursively reverses a string
// Base case: empty or single character string
// Recursive case: last character + reverse of remaining string
func reverseString(s string) string {
	// BASE CASE: Empty or single character
	if len(s) <= 1 {
		return s
	}

	// RECURSIVE CASE: Last character + reverse of rest
	lastChar := string(s[len(s)-1])
	restReversed := reverseString(s[:len(s)-1])
	return lastChar + restReversed
}

// ─────────────────────────────────────────────────────────────────────────────
// MAIN FUNCTION WITH EXAMPLES
// ─────────────────────────────────────────────────────────────────────────────

func main() {
	fmt.Println("════════════════════════════════════════════════════════════")
	fmt.Println("EXAMPLE 1: FACTORIAL")
	fmt.Println("════════════════════════════════════════════════════════════\n")

	fmt.Println("Calculating factorial(5):")
	result := factorial(5)
	fmt.Printf("\n✓ Result: 5! = %d\n\n", result)

	fmt.Println("─────────────────────────────────────────────────────────────")
	fmt.Println("Testing different inputs:")
	fmt.Printf("factorial(0) = %d\n", factorial(0))
	fmt.Printf("factorial(1) = %d\n", factorial(1))
	fmt.Printf("factorial(10) = %d\n\n", factorial(10))

	// ─────────────────────────────────────────────────────────────────────────

	fmt.Println("════════════════════════════════════════════════════════════")
	fmt.Println("EXAMPLE 2: SUM OF DIGITS")
	fmt.Println("════════════════════════════════════════════════════════════\n")

	testNumbers := []int{9, 12, 123, 12345}
	for _, num := range testNumbers {
		fmt.Printf("Sum of digits in %d:\n", num)
		sum := sumOfDigits(num)
		fmt.Printf("✓ Result: %d\n\n", sum)
	}

	// ─────────────────────────────────────────────────────────────────────────

	fmt.Println("════════════════════════════════════════════════════════════")
	fmt.Println("EXAMPLE 3: FIBONACCI SEQUENCE")
	fmt.Println("════════════════════════════════════════════════════════════\n")

	fmt.Println("Naive Fibonacci (slow for large n):")
	for i := 0; i <= 10; i++ {
		fib := fibonacci(i)
		fmt.Printf("fibonacci(%d) = %d\n", i, fib)
	}

	fmt.Println("\n⚠️  WARNING: Naive recursion is VERY slow for large values!")
	fmt.Println("   For fibonacci(30), it would make hundreds of thousands of calls")
	fmt.Printf("   fibonacci(20) = %d (relatively fast)\n", fibonacci(20))

	// ─────────────────────────────────────────────────────────────────────────

	fmt.Println("\n════════════════════════════════════════════════════════════")
	fmt.Println("EXAMPLE 4: FIBONACCI WITH MEMOIZATION (Optimized)")
	fmt.Println("════════════════════════════════════════════════════════════\n")

	fmt.Println("Fibonacci with Memoization (fast even for large n):")
	memo := make(map[int]int)
	for i := 0; i <= 20; i++ {
		fib := fibonacciMemo(i, memo)
		fmt.Printf("fibonacci(%d) = %d\n", i, fib)
	}

	fmt.Println("\n✓ Much faster! Memoization caches results to avoid redundant calculations")
	fmt.Printf("   fibonacci(30) with memoization = %d (very fast!)\n", fibonacciMemo(30, make(map[int]int)))

	// ─────────────────────────────────────────────────────────────────────────

	fmt.Println("\n════════════════════════════════════════════════════════════")
	fmt.Println("EXAMPLE 5: POWER FUNCTION")
	fmt.Println("════════════════════════════════════════════════════════════\n")

	fmt.Println("Calculating powers recursively:")
	fmt.Printf("2^0 = %d\n", power(2, 0))
	fmt.Printf("2^3 = %d\n", power(2, 3))
	fmt.Printf("2^5 = %d\n", power(2, 5))
	fmt.Printf("2^10 = %d\n", power(2, 10))
	fmt.Printf("3^4 = %d\n", power(3, 4))

	// ─────────────────────────────────────────────────────────────────────────

	fmt.Println("\n════════════════════════════════════════════════════════════")
	fmt.Println("EXAMPLE 6: STRING REVERSAL")
	fmt.Println("════════════════════════════════════════════════════════════\n")

	strings := []string{"hello", "Go", "recursion", "programming"}
	for _, str := range strings {
		reversed := reverseString(str)
		fmt.Printf("reverseString(\"%s\") = \"%s\"\n", str, reversed)
	}

	// ─────────────────────────────────────────────────────────────────────────

	fmt.Println("\n════════════════════════════════════════════════════════════")
	fmt.Println("SUMMARY & BEST PRACTICES")
	fmt.Println("════════════════════════════════════════════════════════════\n")

	fmt.Println(`
KEY TAKEAWAYS:

1. BASE CASE IS CRITICAL:
   Every recursive function MUST have a reachable base case
   Without it: infinite recursion → stack overflow

2. RECURSIVE CASE:
   Must make progress towards the base case
   Each call should work with a simpler/smaller version of the problem

3. CALL STACK:
   Each function call adds a frame to the call stack
   Deep recursion can cause stack overflow
   For very large inputs, consider iterative approaches

4. EFFICIENCY:
   Naive recursion can be slow (exponential time complexity)
   Memoization can dramatically improve performance (exponential → linear)
   Iterative solutions are often more efficient

5. WHEN TO USE RECURSION:
   ✓ Tree and graph traversals
   ✓ Divide and conquer algorithms
   ✓ Problems with natural recursive structure
   ✓ Backtracking problems
   ✗ Very deep recursion (use iteration instead)
   ✗ Performance-critical code without memoization

BEST PRACTICES:

✓ Always define a clear base case
✓ Ensure the base case is reachable
✓ Test with edge cases (empty, single element, large input)
✓ Use memoization to optimize recursive solutions
✓ Consider stack depth for large inputs
✓ Document which variables are used in recursion
✓ Use iterative approaches for very deep recursion
✓ Measure performance - recursion isn't always optimal!

PRACTICE:
Try implementing these recursive problems:
- Sum of natural numbers: sum(n) = n + sum(n-1)
- Greatest Common Divisor (GCD)
- Palindrome checker
- Binary search
- Merge sort
- Tree traversal (if you know trees)

Remember: Recursion is powerful but requires careful consideration of base cases,
performance, and stack depth. Master it, but use it wisely!
	`)
}
