package intermediate

import (
	"errors"
	"fmt"
)

// Topic 68: errors
// Combines pragmatic examples with GoBootcamp patterns


func main() {

	fmt.Println("-- 68 Errors --")

	// Simple error from fmt.Errorf
	var err error = fmt.Errorf("an error: %s", "oops")
	fmt.Println(err)

	// errors.New
	err2 := errors.New("something bad happened")
	fmt.Println(err2)

	// Error wrapping with fmt.Errorf %w (Go 1.13+)
	wrapped := fmt.Errorf("outer error: %w", err2)
	fmt.Println(wrapped)

	// Unwrapping
	if errors.Is(wrapped, err2) {
		fmt.Println("wrapped error contains err2")
	}
}

func bonusErrorPatterns() {
	fmt.Println("\n=== BONUS: Error Checking Patterns ===")
	fmt.Println(`
Pattern 1: Immediate check
	err := doSomething()
	if err != nil {
		// handle
	}

Pattern 2: Guard clause (early return)
	if err := doSomething(); err != nil {
		return err
	}

Pattern 3: Multiple operations
	if err := op1(); err != nil {
		return err
	}
	if err := op2(); err != nil {
		return err
	}

Pattern 4: Type checking
	if ve, ok := err.(ValidationError); ok {
		// handle validation error
	}

Pattern 5: Sentinel error checking
	if errors.Is(err, ErrNotFound) {
		// handle not found
	}

Pattern 6: Error wrapping
	if err := operation(); err != nil {
		return fmt.Errorf("context: %w", err)
	}
	`)
}
