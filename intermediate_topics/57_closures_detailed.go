package main

import "fmt"

/*
Topic 57: CLOSURES

CONCEPT: A closure is a function that captures and remembers variables from
its surrounding scope, even after the outer function finishes executing.

KEY IDEAS:
- Closures capture variables by reference (changes persist)
- Each closure instance maintains its own captured state
- Great for encapsulation and state management without globals

EXAMPLE: adder() returns a closure that remembers the variable 'i'.
Each time we call the closure, it increments 'i' and returns the new value.

USE CASES:
✓ Maintaining state across calls (counters, accumulators)
✓ Data encapsulation (hiding internal variables)
✓ Creating specialized functions (function factories)
✓ Callbacks that need to remember context
*/

// Example 1: Counter closure
func adder() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

// Example 2: Countdown closure
func makeCountdown() func(int) int {
	value := 100
	return func(x int) int {
		value -= x
		return value
	}
}

// Example 3: Bank account encapsulation
func createAccount(initialBalance float64) func(string, float64) float64 {
	balance := initialBalance
	return func(op string, amount float64) float64 {
		switch op {
		case "deposit":
			balance += amount
		case "withdraw":
			if amount <= balance {
				balance -= amount
			} else {
				fmt.Println("❌ Insufficient funds!")
			}
		}
		return balance
	}
}

// Example 4: Function factory
func makeMultiplier(factor int) func(int) int {
	return func(x int) int {
		return x * factor
	}
}

func main() {
	fmt.Println("═══════════════════════════════════════════════════════════")
	fmt.Println("EXAMPLE 1: Counter Closure")
	fmt.Println("═══════════════════════════════════════════════════════════\n")

	counter := adder()
	fmt.Println("Calling counter() 5 times:")
	for i := 1; i <= 5; i++ {
		fmt.Printf("  Call %d: %d\n", i, counter())
	}

	fmt.Println("\nEach closure has independent state:")
	counter2 := adder()
	fmt.Printf("  New counter: %d (starts fresh)\n\n", counter2())

	// ─────────────────────────────────────────────────────────────────────────

	fmt.Println("═══════════════════════════════════════════════════════════")
	fmt.Println("EXAMPLE 2: Countdown Closure")
	fmt.Println("═══════════════════════════════════════════════════════════\n")

	countdown := makeCountdown()
	fmt.Println("Subtracting from 100:")
	fmt.Printf("  subtract(10):  %d\n", countdown(10))
	fmt.Printf("  subtract(20):  %d\n", countdown(20))
	fmt.Printf("  subtract(15):  %d\n\n", countdown(15))

	// ─────────────────────────────────────────────────────────────────────────

	fmt.Println("═══════════════════════════════════════════════════════════")
	fmt.Println("EXAMPLE 3: Bank Account Encapsulation")
	fmt.Println("═══════════════════════════════════════════════════════════\n")

	myAccount := createAccount(1000.0)
	fmt.Println("Initial balance: $1000.00\n")

	fmt.Printf("deposit($500):   balance = $%.2f\n", myAccount("deposit", 500))
	fmt.Printf("withdraw($200):  balance = $%.2f\n", myAccount("withdraw", 200))
	fmt.Printf("withdraw($100):  balance = $%.2f\n", myAccount("withdraw", 100))
	fmt.Printf("current():       balance = $%.2f\n\n", myAccount("", 0))

	// ─────────────────────────────────────────────────────────────────────────

	fmt.Println("═══════════════════════════════════════════════════════════")
	fmt.Println("EXAMPLE 4: Function Factory")
	fmt.Println("═══════════════════════════════════════════════════════════\n")

	double := makeMultiplier(2)
	triple := makeMultiplier(3)
	quadruple := makeMultiplier(4)

	fmt.Println("Created specialized functions using factory:\n")
	fmt.Printf("double(5)     = %d\n", double(5))
	fmt.Printf("triple(5)     = %d\n", triple(5))
	fmt.Printf("quadruple(5)  = %d\n\n", quadruple(5))

	// ─────────────────────────────────────────────────────────────────────────

	fmt.Println("═══════════════════════════════════════════════════════════")
	fmt.Println("KEY CONCEPTS & BEST PRACTICES")
	fmt.Println("═══════════════════════════════════════════════════════════\n")

	fmt.Println(`
VARIABLE CAPTURE:
- Closures capture variables by reference
- Changes persist across function calls
- Each closure instance has independent captured variables

STATE ENCAPSULATION:
- Hide internal state from the outside world
- Provide controlled access through closures
- No need for global variables

FUNCTION FACTORIES:
- Create specialized functions with captured parameters
- Reduces code duplication
- Each closure remembers its captured values

MEMORY CONSIDERATIONS:
- Closures keep captured variables alive in memory
- Be careful with large objects that won't be garbage collected
- Consider lifecycle implications

CONCURRENCY:
- Avoid race conditions when multiple goroutines access same closure
- Use synchronization primitives if needed

BEST PRACTICES:
✓ Use closures for encapsulation and state management
✓ Keep captured variable scope narrow
✓ Document what variables are captured
✓ Use only when they genuinely improve code clarity
✓ Be mindful of memory implications
✓ Ensure thread-safe access in concurrent scenarios
	`)
}
