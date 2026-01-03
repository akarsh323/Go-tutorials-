package main

import "fmt"

/*
Topic 57: CLOSURES - A Comprehensive Guide

═══════════════════════════════════════════════════════════════════════════════

DEFINITION:
A closure is a function value that references variables from outside its body.
The function may access and assign to the captured variables, and these variables
persist as long as the closure itself is referenced.

KEY CONCEPTS:
- Closures leverage lexical scoping, capturing variables from their surrounding context
- Functions in Go are first-class citizens, enabling closures to work effectively
- Closures can be assigned to variables, passed as arguments, and returned from functions

HOW IT WORKS:
Closures work by capturing variables from their enclosing scope. Once a closure is created,
it maintains a reference to those variables even after the outer function has finished
execution. This allows closures to access and modify captured variables across multiple calls.

PRACTICAL USE CASES:
1. Maintaining state across multiple calls without exposing state directly
2. Encapsulating functionality and data for cleaner, modular code
3. Callback functions that capture context during async operations
4. Creating function factories that produce specialized functions

CONSIDERATIONS:
- Memory Usage: Closures can keep large objects in memory longer than expected
- Concurrency: Be careful with race conditions when using closures in concurrent programs
- Code Clarity: Avoid overusing closures as they can make code harder to maintain
- Variable Scope: Keep the scope of captured variables as narrow as possible

═══════════════════════════════════════════════════════════════════════════════
*/

// adder returns a closure that maintains state across multiple calls
// Each call increments and returns the current value
func adder() func() int {
	i := 0
	fmt.Println("adder() called - initializing i to:", i)

	return func() int {
		// This anonymous function captures the variable 'i' from adder's scope
		i++ // Increment and update the captured variable
		fmt.Println("Adding 1 to i. New value:", i)
		return i
	}
}

func main() {
	fmt.Println("═══════════════════════════════════════════════════════════")
	fmt.Println("EXAMPLE 1: Basic Closure with adder()")
	fmt.Println("═══════════════════════════════════════════════════════════\n")

	// sequence is assigned the returned function from adder()
	sequence := adder()

	fmt.Println("Calling sequence(1):", sequence())
	fmt.Println("Calling sequence(2):", sequence())
	fmt.Println("Calling sequence(3):", sequence())
	fmt.Println("Calling sequence(4):", sequence())
	fmt.Println("Calling sequence(5):", sequence())

	fmt.Println("\n--- Creating a second closure (independent state) ---\n")

	// sequence2 is a NEW closure with its own independent 'i' variable
	sequence2 := adder()
	fmt.Println("Calling sequence2(1):", sequence2())

	fmt.Println("\nNotice how sequence2 starts fresh with its own i = 0")
	fmt.Println("Each closure maintains its own captured variable state!\n")

	// ─────────────────────────────────────────────────────────────────────────────

	fmt.Println("═══════════════════════════════════════════════════════════")
	fmt.Println("EXAMPLE 2: Closure with Parameters (Anonymous Function)")
	fmt.Println("═══════════════════════════════════════════════════════════\n")

	// This is an anonymous function that is immediately executed (IIFE pattern)
	// It returns another function that captures 'countdown'
	subtractor := func() func(int) int {
		countdown := 99
		fmt.Println("Anonymous function called - initializing countdown to:", countdown)

		return func(x int) int {
			// This inner function captures 'countdown' from the outer scope
			countdown -= x
			fmt.Printf("Subtracting %d from countdown. New value: %d\n", x, countdown)
			return countdown
		}
	}() // The () at the end executes the function immediately

	fmt.Println("\nCalling subtractor with different values:")
	fmt.Println("Result:", subtractor(1))
	fmt.Println("Result:", subtractor(2))
	fmt.Println("Result:", subtractor(3))
	fmt.Println("Result:", subtractor(4))
	fmt.Println("Result:", subtractor(5))

	fmt.Println("\nThe countdown state is retained across all calls!")
	fmt.Println("99 - 1 = 98, 98 - 2 = 96, 96 - 3 = 93, 93 - 4 = 89, 89 - 5 = 84\n")

	// ─────────────────────────────────────────────────────────────────────────────

	fmt.Println("═══════════════════════════════════════════════════════════")
	fmt.Println("EXAMPLE 3: Practical Use Case - Creating a Bank Account")
	fmt.Println("═══════════════════════════════════════════════════════════\n")

	// This closure creates a simple bank account with encapsulated balance
	account := func() func(string, float64) float64 {
		balance := 1000.0
		fmt.Printf("Account created with initial balance: $%.2f\n\n", balance)

		return func(operation string, amount float64) float64 {
			switch operation {
			case "deposit":
				balance += amount
				fmt.Printf("Deposited $%.2f. New balance: $%.2f\n", amount, balance)
			case "withdraw":
				if amount <= balance {
					balance -= amount
					fmt.Printf("Withdrew $%.2f. New balance: $%.2f\n", amount, balance)
				} else {
					fmt.Println("Insufficient funds!")
				}
			default:
				fmt.Printf("Current balance: $%.2f\n", balance)
			}
			return balance
		}
	}() // Immediately executed

	fmt.Println("Performing transactions:")
	account("deposit", 500)
	account("withdraw", 200)
	account("deposit", 300)
	account("withdraw", 1000)
	account("", 0) // Check balance
	account("withdraw", 100)

	fmt.Println("\nThe balance is encapsulated - we can't access it directly,")
	fmt.Println("only through the returned closure function!\n")

	// ─────────────────────────────────────────────────────────────────────────────

	fmt.Println("═══════════════════════════════════════════════════════════")
	fmt.Println("EXAMPLE 4: Function Factory - Creating Multipliers")
	fmt.Println("═══════════════════════════════════════════════════════════\n")

	// makeMultiplier is a function factory that creates multiplier functions
	makeMultiplier := func(factor int) func(int) int {
		return func(x int) int {
			return x * factor
		}
	}

	// Create different multiplier functions with different factors
	double := makeMultiplier(2)
	triple := makeMultiplier(3)
	quadruple := makeMultiplier(4)

	fmt.Println("Created specialized functions using a factory:")
	fmt.Printf("double(5) = %d\n", double(5))
	fmt.Printf("triple(5) = %d\n", triple(5))
	fmt.Printf("quadruple(5) = %d\n", quadruple(5))

	fmt.Println("\nEach closure captures a different 'factor' value!")
	fmt.Println("This demonstrates how closures can be used to create specialized functions.\n")

	// ─────────────────────────────────────────────────────────────────────────────

	fmt.Println("═══════════════════════════════════════════════════════════")
	fmt.Println("SUMMARY & BEST PRACTICES")
	fmt.Println("═══════════════════════════════════════════════════════════\n")

	fmt.Println(`
KEY TAKEAWAYS:

1. VARIABLE CAPTURE:
   Closures capture variables from their enclosing scope by reference
   Changes to captured variables persist across closure invocations

2. STATE ENCAPSULATION:
   Closures allow you to hide internal state while providing a public interface
   This creates clean, modular code without global variables

3. INDEPENDENT INSTANCES:
   Each call to the outer function creates a NEW closure with its own captured variables
   Different closures don't interfere with each other's state

4. MEMORY CONSIDERATIONS:
   Closures keep captured variables alive in memory as long as the closure exists
   Be careful with large objects or slices as they won't be garbage collected

5. CONCURRENCY:
   Use caution with closures in concurrent programs to avoid race conditions
   Consider synchronization primitives if multiple goroutines access the same closure

BEST PRACTICES:

✓ Use closures to encapsulate state and create clean abstractions
✓ Keep the scope of captured variables as narrow as possible
✓ Document which variables are captured by your closures
✓ Avoid overusing closures - use them where they genuinely improve code clarity
✓ Be mindful of memory usage when closures capture large objects
✓ In concurrent scenarios, ensure thread-safe access to captured variables

Closures are a powerful feature of Go that enable functional programming patterns
and help create expressive, maintainable code!
	`)
}
