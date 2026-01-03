package intermediate

import "fmt"

// Topic 58: recursion
// Combines pragmatic examples with GoBootcamp patterns


func main() {

	fmt.Println("-- 58 Recursion --")
	fmt.Println("factorial(0) =", factorial58(0))
	fmt.Println("factorial(1) =", factorial58(1))
	fmt.Println("factorial(5) =", factorial58(5))

	// Note: Go does not guarantee tail-call optimization. For very deep
	// recursive algorithms prefer iterative approaches or explicit stacks.
}

func factorial58(n int) int {
	if n <= 1 {
		return 1
	}
	return n * factorial58(n-1)
}

func factorialTail(n, accumulator int) int {
	if n <= 1 {
		return accumulator
	}
	return factorialTail(n-1, n*accumulator)
}

func recursionBonusExample() {
	fmt.Println("\n=== BONUS: Tail Recursion ===")
	fmt.Printf("5! (tail recursive) = %d\n", factorialTail(5, 1))
	fmt.Println("\nTail recursion passes accumulator to avoid stack buildup")
	fmt.Println("(Note: Go doesn't optimize TCO, but the pattern is useful)")
}
