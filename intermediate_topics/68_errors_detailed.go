package intermediate

import (
	"errors"
	"fmt"
	"math"
)

// ============================================================================
// Topic 68: ERROR HANDLING IN GO
// ============================================================================
//
// CORE CONCEPT: Errors in Go are TREATED AS VALUES, not exceptions.
// Think of errors like integers or strings — you handle them explicitly.
//
// KEY PRINCIPLE: "Don't just panic. Check errors and handle them properly."
//
// In other languages (Java, Python): Use try/catch/finally
// In Go: Return errors as values and check them with "if err != nil"
//
// ============================================================================

func main() {
	fmt.Println("=============== 68 Error Handling in Go ===============")

	// ========================================================================
	// SECTION 1: The Error Interface - Understanding the Foundation
	// ========================================================================
	fmt.Println("\n--- SECTION 1: The Error Interface ---")
	fmt.Println(`
In Go, an error is ANYTHING that implements this interface:

    type error interface {
        Error() string
    }

That's it. Just ONE method that returns a string.

This simple design means:
  • Any type with an Error() method IS an error
  • You can create custom errors easily
  • Errors are just regular values you pass around
`)

	// ========================================================================
	// SECTION 2: Creating Errors - Four Common Ways
	// ========================================================================
	fmt.Println("\n--- SECTION 2: Creating Errors ---")

	// Method 1: errors.New() - Simple string errors
	fmt.Println("\nMethod 1: errors.New() - Simple string errors")
	err1 := errors.New("something bad happened")
	fmt.Printf("err1: %v (type: %T)\n", err1, err1)

	// Method 2: fmt.Errorf() - Formatted errors
	fmt.Println("\nMethod 2: fmt.Errorf() - Formatted error messages")
	err2 := fmt.Errorf("an error: %s", "oops")
	fmt.Printf("err2: %v\n", err2)

	// Method 3: Custom error struct - Rich error information
	fmt.Println("\nMethod 3: Custom error struct (holds more data)")
	err3 := &ValidationError{
		Field:   "email",
		Message: "invalid email format",
		Code:    400,
	}
	fmt.Printf("err3: %v\n", err3)

	// Method 4: Nil - No error occurred
	fmt.Println("\nMethod 4: nil - Represents NO error")
	var err4 error = nil
	fmt.Printf("err4: %v (nil means success)\n", err4)

	// ========================================================================
	// SECTION 3: The Standard Error-Checking Pattern
	// ========================================================================
	fmt.Println("\n" + "="*60)
	fmt.Println("--- SECTION 3: Error Checking Pattern (if err != nil) ---")

	result, err := Sqrt(-16) // This will return an error
	if err != nil {
		fmt.Printf("Error occurred: %v\n", err)
		// In real code, you would handle the error here:
		//   - Log it
		//   - Return to caller
		//   - Provide default value
		//   - Retry operation
	} else {
		fmt.Printf("Result: %v\n", result)
	}

	result2, err2 := Sqrt(25) // This will succeed
	if err2 != nil {
		fmt.Printf("Error: %v\n", err2)
	} else {
		fmt.Printf("Square root of 25 is: %v\n", result2)
	}

	// ========================================================================
	// SECTION 4: Error Wrapping - Adding Context to Errors
	// ========================================================================
	fmt.Println("\n" + "="*60)
	fmt.Println("--- SECTION 4: Error Wrapping (Adding Context) ---")

	fmt.Println(`
When errors bubble up through your code, they lose context.

Example:
  readConfig() calls readFile()
  readFile() returns: "file not found"
  readConfig() returns: "file not found"  ← No context about WHERE

Solution: Wrap errors with fmt.Errorf and %w verb

  readConfig() calls readFile()
  readFile() returns: "file not found"
  readConfig() returns: fmt.Errorf("config init: %w", originalErr)
  Result: "config init: file not found"  ← Full context!
`)

	wrappedErr := readConfig()
	if wrappedErr != nil {
		fmt.Printf("Wrapped error: %v\n", wrappedErr)
	}

	// ========================================================================
	// SECTION 5: Unwrapping Errors - Checking Inner Errors
	// ========================================================================
	fmt.Println("\n" + "="*60)
	fmt.Println("--- SECTION 5: Unwrapping (Finding Original Error) ---")

	fmt.Println(`
When errors are wrapped, you might need to check if the original error
was a specific type.

Go 1.13+ provides: errors.Is() and errors.As()

errors.Is(err, target):
  Check if err (or its chain) contains target
  
errors.As(err, &target):
  If err (or its chain) matches target type, extract it
`)

	// Test errors.Is
	outerErr := fmt.Errorf("operation failed: %w", errors.New("file not found"))
	originalErr := errors.New("file not found")

	if errors.Is(outerErr, originalErr) {
		fmt.Println("✓ Found the original 'file not found' error inside the wrapped error")
	}

	// Test errors.As with custom error type
	customErr := &ValidationError{
		Field:   "password",
		Message: "too short",
		Code:    422,
	}
	wrappedCustom := fmt.Errorf("validation failed: %w", customErr)

	var ve *ValidationError
	if errors.As(wrappedCustom, &ve) {
		fmt.Printf("✓ Extracted custom error: Field=%s, Message=%s, Code=%d\n",
			ve.Field, ve.Message, ve.Code)
	}

	// ========================================================================
	// SECTION 6: The Guard Clause Pattern (Early Return)
	// ========================================================================
	fmt.Println("\n" + "="*60)
	fmt.Println("--- SECTION 6: Guard Clause Pattern (Preferred Style) ---")

	fmt.Println(`
Pattern 1: Immediate check (verbose)
    err := doSomething()
    if err != nil {
        return err
    }

Pattern 2: Guard clause (preferred - shorter scope)
    if err := doSomething(); err != nil {
        return err
    }

WHY guard clause is better:
  • 'err' variable scope is limited to the if block
  • Less variable pollution in outer scope
  • Cleaner and more Go-idiomatic
  • Encourages early returns
`)

	// Demonstrating guard clause
	if err := processData(); err != nil {
		fmt.Printf("Processing failed: %v\n", err)
	} else {
		fmt.Println("Processing succeeded")
	}

	// ========================================================================
	// SECTION 7: Custom Errors - The "Power User" Move
	// ========================================================================
	fmt.Println("\n" + "="*60)
	fmt.Println("--- SECTION 7: Custom Errors (Structs + Error() method) ---")

	fmt.Println(`
Simple string errors (errors.New, fmt.Errorf) are great for simple cases.
For complex scenarios, create a STRUCT and implement Error().

Steps:
  1. Define a struct with fields for your error data
  2. Write an Error() method on the struct
  3. Go automatically recognizes it as an error!

Benefits:
  • Store error codes, timestamps, context fields
  • Type-assert to extract rich information
  • Different error types can behave differently
`)

	// Create and use a custom error
	apiErr := &APIError{
		StatusCode: 500,
		Message:    "Database connection failed",
		Endpoint:   "/api/users",
	}
	fmt.Printf("Custom error: %v\n", apiErr)

	// Type-assert to get more details
	if apiError, ok := apiErr.(*APIError); ok {
		fmt.Printf("Details: Status=%d, Endpoint=%s\n",
			apiError.StatusCode, apiError.Endpoint)
	}

	// ========================================================================
	// SECTION 8: Error Handling Patterns & Best Practices
	// ========================================================================
	fmt.Println("\n" + "="*60)
	fmt.Println("--- SECTION 8: Error Handling Patterns ---")

	demonstrateErrorPatterns()

	// ========================================================================
	// SECTION 9: The Golden Rules
	// ========================================================================
	fmt.Println("\n" + "="*60)
	fmt.Println("--- SECTION 9: The Golden Rules of Go Error Handling ---")
	fmt.Println(`
✓ DO:
  • Check every error with "if err != nil"
  • Be descriptive: "Failed to connect to DB on port 8080"
  • Wrap errors to add context: fmt.Errorf("outer: %w", err)
  • Log errors with timestamps for debugging
  • Use custom error types for different error scenarios
  • Return errors instead of panicking
  • Handle errors at the appropriate level

✗ DON'T:
  • Ignore errors (no: _ = operation())
  • Use panic for recoverable errors
  • Hide original errors (no: fmt.Errorf("error: %v", err))
  • Return generic "Error occurred" messages
  • Panic when you should return an error
  • Swallow error information in logs

PHILOSOPHY:
  "Don't panic. Errors happen. Check them, log them, handle them gracefully."
`)
}

// ============================================================================
// HELPER FUNCTION 1: Sqrt - Demonstrates basic error returning
// ============================================================================
//
// Pattern: Return (value, error)
// - If successful: return (result, nil)
// - If failed: return (zeroValue, error)

func Sqrt(x float64) (float64, error) {
	if x < 0 {
		// Create an error when input is invalid
		return 0, errors.New("math error: square root of negative number")
	}
	// Success case: return result and nil (no error)
	return math.Sqrt(x), nil
}

// ============================================================================
// HELPER FUNCTION 2: readConfig - Demonstrates error wrapping
// ============================================================================
//
// Shows how to add context when forwarding errors up the call stack.

func readConfig() error {
	// Simulate an error from a lower-level function
	originalErr := errors.New("config file not found")

	// Wrap it with context using %w (wrap verb)
	// This preserves the original error in the chain
	return fmt.Errorf("failed to initialize config: %w", originalErr)
}

// ============================================================================
// HELPER FUNCTION 3: processData - Demonstrates guard clause
// ============================================================================

func processData() error {
	// Simulate successful processing
	return nil // No error
}

// ============================================================================
// HELPER FUNCTION 4: Custom Error Type 1 - ValidationError
// ============================================================================
//
// A custom error struct that holds validation-specific information.
// By implementing the Error() method, Go treats it as an error.

type ValidationError struct {
	Field   string // Which field failed validation?
	Message string // What's wrong with it?
	Code    int    // HTTP-style error code
}

// Error() implements the error interface
// This is the method Go calls when you do fmt.Println(validationError)
func (v *ValidationError) Error() string {
	return fmt.Sprintf("ValidationError(%d) on field '%s': %s",
		v.Code, v.Field, v.Message)
}

// ============================================================================
// HELPER FUNCTION 5: Custom Error Type 2 - APIError
// ============================================================================
//
// Another custom error struct for API-related errors.

type APIError struct {
	StatusCode int
	Message    string
	Endpoint   string
}

func (a *APIError) Error() string {
	return fmt.Sprintf("API Error %d at %s: %s",
		a.StatusCode, a.Endpoint, a.Message)
}

// ============================================================================
// HELPER FUNCTION 6: demonstrateErrorPatterns
// ============================================================================

func demonstrateErrorPatterns() {
	fmt.Println(`
Pattern 1: Basic check (when you need the error for logic)
    err := operation()
    if err != nil {
        // Handle or return
        return fmt.Errorf("context: %w", err)
    }

Pattern 2: Guard clause (most Go-idiomatic)
    if err := operation(); err != nil {
        return fmt.Errorf("context: %w", err)
    }
    // Continue with success case

Pattern 3: Sentinel error (check for specific error)
    if errors.Is(err, ErrNotFound) {
        // Handle "not found" specially
    }

Pattern 4: Type assertion (check for specific error TYPE)
    if ve, ok := err.(*ValidationError); ok {
        fmt.Printf("Validation failed on field: %s\n", ve.Field)
    }

Pattern 5: Error wrapping (add context while forwarding)
    if err := deepOperation(); err != nil {
        return fmt.Errorf("operation context: %w", err)
    }

Pattern 6: Multiple steps with guard clauses
    if err := step1(); err != nil {
        return fmt.Errorf("step 1: %w", err)
    }
    if err := step2(); err != nil {
        return fmt.Errorf("step 2: %w", err)
    }
    if err := step3(); err != nil {
        return fmt.Errorf("step 3: %w", err)
    }
    return nil

IMPORTANT NOTE:
  Use guard clauses (Pattern 2) for most of your code.
  It's the Go way.
`)
}

// ============================================================================
// COMPREHENSIVE COMPARISON TABLE
// ============================================================================
//
// When to use each error creation method:
//
// Method               | Use Case              | Example
// ================================================================================
// errors.New()         | Simple, static text   | errors.New("user not found")
// fmt.Errorf()         | Formatted message     | fmt.Errorf("user %s not found", id)
// fmt.Errorf + %w      | Wrap existing error   | fmt.Errorf("init: %w", err)
// Custom struct        | Rich error data       | &ValidationError{Field: "email", ...}
// Nil                  | Success (no error)    | if err != nil { /* error */ }
//
// ============================================================================
