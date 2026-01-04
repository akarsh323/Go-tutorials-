package intermediate

import (
	"errors"
	"fmt"
	"strings"
)

// ============================================================================
// Topic 69: CUSTOM ERRORS - Building Rich Error Types
// ============================================================================
//
// CORE CONCEPT: Standard errors (errors.New()) are too simple for professional
// applications. Custom error types let you:
//
//   1. Add error codes (404, 500, 422, etc.)
//   2. Wrap original errors (preserve the root cause)
//   3. Add context (userID, field name, operation)
//   4. Support decision-making (CanRetry, IsTimeout, etc.)
//   5. Create structured error responses (JSON, logs)
//
// HOW: Create a struct with an Error() method = custom error type
//
// ============================================================================

func main() {
	fmt.Println("=============== 69 Custom Errors - Rich Error Types ===============")

	// ========================================================================
	// SECTION 0: The Big Picture - Why Custom Errors Matter
	// ========================================================================
	fmt.Println("\n--- SECTION 0: The Big Picture ---")
	fmt.Println(`
PROBLEM IN SIMPLE SYSTEMS
────────────────────────
Standard errors (errors.New("failed")) lose too much information:

  Function A calls Function B calls Function C
  C fails with: "disk failure"
  B passes it to A
  A sees: "disk failure"
  → Lost: What was B doing? What was A trying to do?

This is like a crash report with just:
  "Error occurred"
Instead of:
  "Error 500: Failed to save user profile. Root cause: disk failure"


SOLUTION: Custom Error Types
────────────────────────────
Create structs that hold:
  1. ERROR CODE     → Severity (404, 500, 422)
  2. CONTEXT        → What operation failed ("save file")
  3. ORIGINAL ERROR → Root cause ("disk failure")

This creates a CHAIN of information that tells the complete story.

In the HTTP API world:
  • 404 Not Found      → Client's fault (bad input)
  • 422 Unprocessable  → Client's fault (validation failed)
  • 500 Server Error   → Our fault (system error)

Without custom errors: You can't distinguish them
With custom errors: You return the right HTTP status automatically
`)

	// ========================================================================
	// SECTION 1: The Foundation - What Makes an Error?
	// ========================================================================
	fmt.Println("\n--- SECTION 1: Error Interface Review ---")
	fmt.Println(`
Remember from Topic 68: An error is anything with an Error() method:

    type error interface {
        Error() string
    }

This means: ANY struct with an Error() method IS an error!
Let's create custom error types that hold MORE information.
`)

	// ========================================================================
	// SECTION 2: Simple Custom Error - No Wrapping
	// ========================================================================
	fmt.Println("\n--- SECTION 2: Simple Custom Error (Basic) ---")

	// Create a simple custom error
	demoErr := DemoError{
		Code:    500,
		Message: "database connection failed",
	}
	fmt.Printf("Simple custom error: %v\n", demoErr)

	// Type-assert to get the fields
	fmt.Printf("  • Error output: %v\n", demoErr)
	fmt.Printf("  • Error Code: %d\n", demoErr.Code)
	fmt.Printf("  • Message: %s\n", demoErr.Message)

	// ========================================================================
	// SECTION 3: Validation Error - Domain-Specific Custom Error
	// ========================================================================
	fmt.Println("\n" + strings.Repeat("=", 60))
	fmt.Println("--- SECTION 3: Validation Error (Domain-Specific) ---")

	valErr := ValidationError{
		Field: "email",
		Issue: "invalid format",
		Value: "notanemail",
	}
	fmt.Printf("Validation error: %v\n", valErr)

	// Type-assert and inspect
	fmt.Printf("  • Error output: %v\n", valErr)
	fmt.Printf("  • Field: %s\n", valErr.Field)
	fmt.Printf("  • Issue: %s\n", valErr.Issue)
	fmt.Printf("  • Value received: %s\n", valErr.Value)

	// ========================================================================
	// SECTION 4: The Power Move - Error Wrapping
	// ========================================================================
	fmt.Println("\n" + strings.Repeat("=", 60))
	fmt.Println("--- SECTION 4: Error Wrapping (Preserving Root Cause) ---")

	fmt.Println(`
THE PROBLEM: Lost Context in Error Chains
──────────────────────────────────────────

Call Chain: main() → doSomething() → doSomethingElse()

Scenario 1: WITHOUT Wrapping (BAD ✗)
  doSomethingElse() fails with: "internal disk failure"
  doSomething() returns it directly
  main() sees: "internal disk failure"
  → Lost all context! Where did this fail? In what operation?

Scenario 2: WITH Wrapping (GOOD ✓)
  doSomethingElse() fails with: "internal disk failure"
  doSomething() catches it and wraps with context
  main() sees: "Error 500: failed to save file, caused by: internal disk failure"
  → Full story! WHAT happened, HOW bad it is, WHY it happened!


THE SOLUTION: Error Wrapping Pattern
─────────────────────────────────────

Step 1: Lower-level function fails
  func doSomethingElse() error {
      return errors.New("internal disk failure")
  }

Step 2: Upper-level function intercepts and wraps
  func doSomething() error {
      err := doSomethingElse()
      if err != nil {
          // Wrap with context!
          return &CustomError{
              Code:    500,
              Message: "failed to save file",
              Err:     err,  // ← Preserve original error!
          }
      }
      return nil
  }

Step 3: Caller sees the full chain
  Error 500: failed to save file, caused by: internal disk failure
  ↑Code       ↑Context message                   ↑Root cause


WHY This Matters for Debugging
───────────────────────────────
• Chain of events: Like a stack trace in other languages
• Error codes: Tells you severity (404 vs 500 vs 422)
• Context: Tells you WHERE in the code it failed
• Root cause: Tells you WHY it ultimately failed

Without wrapping: You're flying blind
With wrapping: You have a complete story
`)

	// Demonstrate error wrapping
	err := doSomething()
	if err != nil {
		fmt.Printf("\n✓ Error received in main():\n")
		fmt.Printf("  %v\n\n", err)
		fmt.Println("  Notice how it shows:")
		fmt.Println("  • Error code (500)")
		fmt.Println("  • Context message (failed to save file)")
		fmt.Println("  • Root cause (internal disk failure)")
		fmt.Println("  → This is the FULL chain of what went wrong!")
	}

	// ========================================================================
	// SECTION 5: Unwrapping Wrapped Errors
	// ========================================================================
	fmt.Println("\n" + strings.Repeat("=", 60))
	fmt.Println("--- SECTION 5: Extracting Wrapped Errors (Type Assertion) ---")

	fmt.Println(`
PROBLEM: When you receive a wrapped error, how do you get the details?

The error is the string representation:
  "Error 500: failed to save file, caused by: internal disk failure"

But you need the FIELDS:
  • Code: 500 (to decide what HTTP status to return)
  • Message: "failed to save file" (context)
  • Err: errors.New("internal disk failure") (original error object)


SOLUTION: Type Assertion (Type Check + Extract)
───────────────────────────────────────────────

Go's type assertion lets you "unwrap" an error and access its fields:

  err := doSomething()
  if err != nil {
      if wrapped, ok := err.(*WrappedError); ok {
          fmt.Printf("Code: %d\n", wrapped.Code)
          fmt.Printf("Message: %s\n", wrapped.Message)
          fmt.Printf("Inner error: %v\n", wrapped.Err)
      }
  }

Breakdown:
  err.(*WrappedError) → Try to convert err to *WrappedError type
  ok → True if conversion succeeded, false otherwise
  wrapped → The actual WrappedError struct with all fields accessible
`)

	err = doSomething()
	if err != nil {
		fmt.Println("\n✓ Extracting wrapped error details:")
		// Type-assert to see if it's a WrappedError
		if wrapped, ok := err.(*WrappedError); ok {
			fmt.Printf("  • Is a WrappedError? YES\n")
			fmt.Printf("  • Code: %d\n", wrapped.Code)
			fmt.Printf("  • Message: %s\n", wrapped.Message)
			fmt.Printf("  • Inner error: %v\n", wrapped.Err)
			fmt.Println("\n  Now you can:")
			fmt.Println("    - Return HTTP 500 status (Code field)")
			fmt.Println("    - Log the context (Message field)")
			fmt.Println("    - Inspect root cause (Err field)")
		}
	}

	// ========================================================================
	// SECTION 6: Real-World Example - API Error Handling
	// ========================================================================
	fmt.Println("\n" + strings.Repeat("=", 60))
	fmt.Println("--- SECTION 6: Real-World Example - API Errors ---")

	// Simulate fetching a user
	user := getUserByID("invalid-id")
	if user == nil {
		// This demonstrates how custom errors help in API handlers
		fmt.Println("Could not fetch user - see custom error above")
	}

	// ========================================================================
	// SECTION 7: Multiple Custom Error Types
	// ========================================================================
	fmt.Println("\n" + strings.Repeat("=", 60))
	fmt.Println("--- SECTION 7: Multiple Error Types (Type Switching) ---")

	fmt.Println(`
In real applications, you have MULTIPLE error types:
  • ValidationError: Input validation failed
  • AuthError: Authentication/authorization failed
  • DatabaseError: Database operation failed
  • NetworkError: Network call failed

You handle each type differently:
`)

	// Example: Process multiple operations that could fail
	operations := []struct {
		name string
		fn   func() error
	}{
		{"validate email", func() error {
			return ValidationError{
				Field: "email",
				Issue: "already exists",
				Value: "user@example.com",
			}
		}},
		{"check auth", func() error {
			return AuthError{
				Reason:  "invalid token",
				TokenID: "abc123",
			}
		}},
		{"connect database", func() error {
			return &DatabaseError{
				Operation: "INSERT",
				Table:     "users",
				Inner:     errors.New("connection timeout"),
			}
		}},
	}

	for _, op := range operations {
		err := op.fn()
		handleErrorByType(op.name, err)
	}

	// ========================================================================
	// SECTION 8: Error Methods - Adding Decision-Making Logic
	// ========================================================================
	fmt.Println("\n" + strings.Repeat("=", 60))
	fmt.Println("--- SECTION 8: Error Methods (Is it retryable? Is it temporary?) ---")

	fmt.Println(`
Custom error types can have METHODS to help with decisions:

type DatabaseError struct {
    Operation string
    Table     string
    Inner     error
}

// Can we retry this operation?
func (d *DatabaseError) CanRetry() bool {
    return d.Operation == "SELECT" // Only retry reads, not writes
}

// Is this a temporary/transient error?
func (d *DatabaseError) IsTemporary() bool {
    // Check if the inner error is temporary
    if te, ok := d.Inner.(interface{ Temporary() bool }); ok {
        return te.Temporary()
    }
    return false
}

This lets callers DECIDE how to handle the error:
  if dbErr.CanRetry() {
      retry()
  } else {
      fail()
  }
`)

	// Demonstrate CanRetry() method
	dbErr := DatabaseError{
		Operation: "SELECT",
		Table:     "users",
		Inner:     errors.New("timeout"),
	}
	if dbErr.CanRetry() {
		fmt.Println("✓ Can retry a SELECT operation")
	}

	dbErr2 := DatabaseError{
		Operation: "DELETE",
		Table:     "users",
		Inner:     errors.New("timeout"),
	}
	if !dbErr2.CanRetry() {
		fmt.Println("✗ Cannot retry a DELETE operation (too risky)")
	}

	// ========================================================================
	// SECTION 9: The Golden Rules for Custom Errors
	// ========================================================================
	fmt.Println("\n" + strings.Repeat("=", 60))
	fmt.Println("--- SECTION 9: Golden Rules for Custom Errors ---")
	fmt.Println(`
✓ DO:
  • Create custom error types for different error categories
  • Wrap errors to preserve the root cause
  • Use meaningful field names (Operation, Table, Field, etc.)
  • Implement Error() method to format error message
  • Add helper methods for decision-making (CanRetry, IsTimeout)
  • Export error types (start with capital letter) if used outside package
  • Use pointer receivers in Error() method

✗ DON'T:
  • Ignore the original error when wrapping
  • Hide original errors: fmt.Sprintf("error: %v", err) loses type info
  • Create too many custom error types (keep it simple)
  • Duplicate the error message in the wrapper
  • Forget to initialize custom errors properly
  • Use panic when you should return a custom error

PATTERN (The Standard Way):
  type CustomError struct {
      Code    int
      Message string
      Inner   error  // The wrapped error
  }

  func (c *CustomError) Error() string {
      return fmt.Sprintf("Error %d: %s, %v", c.Code, c.Message, c.Inner)
  }
`)

	// ========================================================================
	// SECTION 10: Comparison - Simple vs Custom Errors
	// ========================================================================
	fmt.Println("\n" + strings.Repeat("=", 60))
	fmt.Println("--- SECTION 10: When to Use Each Approach ---")
	fmt.Println(`
USE errors.New() or fmt.Errorf() when:
  • Error message is simple: "file not found"
  • You don't need to check the error type later
  • No additional context needed beyond the message

USE Custom Error Struct when:
  • You need error codes (404, 500, 422)
  • You need to wrap other errors
  • You need to carry context (userID, field, operation)
  • You need to distinguish error types
  • You need helper methods (CanRetry, IsTemporary)
  • Building APIs or complex systems

RULE OF THUMB:
  Start simple (errors.New)
  When you need more, switch to custom errors
  Don't over-engineer early
`)
}

// ============================================================================
// CUSTOM ERROR TYPE 1: DemoError - Simple (No Wrapping)
// ============================================================================
//
// A basic custom error struct that holds:
//   - Code: HTTP-style error code
//   - Message: Description of what went wrong

type DemoError struct {
	Code    int
	Message string
}

// Implement the error interface
func (d DemoError) Error() string {
	return fmt.Sprintf("DemoError (%d): %s", d.Code, d.Message)
}

// ============================================================================
// CUSTOM ERROR TYPE 2: ValidationError - Domain-Specific
// ============================================================================
//
// Used when input validation fails.
// Holds information about which field failed and why.

type ValidationError struct {
	Field string // Which field failed?
	Issue string // What's wrong with it?
	Value string // What value was provided?
}

// Implement the error interface
func (v ValidationError) Error() string {
	return fmt.Sprintf("Validation Error: field '%s' %s (received: %q)",
		v.Field, v.Issue, v.Value)
}

// ============================================================================
// CUSTOM ERROR TYPE 3: WrappedError - With Original Error (KEY PATTERN)
// ============================================================================
//
// This is the MOST IMPORTANT pattern for professional Go code.
// It solves the "lost context" problem in error chains.
//
// Struct Fields:
//   - Code: int           → HTTP-style status code (404, 500, 422, etc.)
//   - Message: string     → Context message describing the operation
//   - Err: error          → The ORIGINAL error (the root cause)
//
// Why Three Fields?
//
//   Code: Tells the CALLER how severe the error is
//     • 404 Not Found → User input was bad
//     • 500 Server Error → Our system failed
//     • 422 Unprocessable → Validation error
//
//   Message: Tells the DEVELOPER what operation failed
//     • "failed to save file"
//     • "database connection refused"
//     • "email service timeout"
//
//   Err: Tells the ROOT CAUSE
//     • "internal disk failure"
//     • "connection refused"
//     • "read timeout"
//
// The Error() Method Output:
//   When you print this error, all three pieces combine:
//   "Error 500: failed to save file, caused by: internal disk failure"
//
// This is the COMPLETE picture: WHAT + HOW BAD + WHY
//
// ============================================================================

type WrappedError struct {
	Code    int    // HTTP-style code
	Message string // Context message
	Err     error  // The ORIGINAL error (the root cause)
}

// Implement the error interface
func (w *WrappedError) Error() string {
	// Format: "Error [Code]: [Message], caused by: [Original Error]"
	return fmt.Sprintf("Error %d: %s, caused by: %v", w.Code, w.Message, w.Err)
}

// ============================================================================
// HELPER FUNCTION 1: doSomethingElse - The "Inner" Function
// ============================================================================
//
// This simulates a low-level function that encounters an error.
// In real code, this might be a database call or file operation.

func doSomethingElse() error {
	// Simulate an error from a lower-level operation
	return errors.New("internal disk failure")
}

// ============================================================================
// HELPER FUNCTION 2: doSomething - The "Outer" Function (Wrapper)
// ============================================================================
//
// This function demonstrates the ERROR WRAPPING PATTERN.
// When doSomethingElse() fails, we don't just pass the error up.
// Instead, we "intercept" it and wrap it with additional context.
//
// The Wrapping Logic:
// ──────────────────
//
// Step 1: Call the lower-level function
//   err := doSomethingElse()
//
// Step 2: Check if it failed
//   if err != nil {
//
// Step 3: Create a new error that wraps the original
//   return &WrappedError{
//       Code:    500,           ← Our decision: This is a server error
//       Message: "failed to save file",  ← Our context
//       Err:     err,           ← Original error (preserved!)
//   }
//
// Result: The original error is PRESERVED inside the new error
//         And we ADD our own context layer
//
// When main() receives this error, it gets the COMPLETE story:
//   "Error 500: failed to save file, caused by: internal disk failure"
//                 ↑ Our context         ↑ Original cause
//            ↑ Severity
//
// Without wrapping: main() would just see "internal disk failure"
//                   (Lost the fact that it was a save operation!)
//
// With wrapping: main() knows:
//   • IT'S SEVERE (500)
//   • WHAT FAILED (save file)
//   • WHY IT FAILED (disk failure)

func doSomething() error {
	// Call the inner function
	err := doSomethingElse()

	// If the inner function failed, we intercept and wrap it
	if err != nil {
		// Instead of just returning the raw error,
		// we wrap it in our custom error to add context
		return &WrappedError{
			Code:    500,
			Message: "failed to save file",
			Err:     err, // Preserve the original error!
		}
	}

	// Success case: no error
	return nil
}

// ============================================================================
// CUSTOM ERROR TYPE 4: ValidationError (Domain-Specific)
// ============================================================================

type AuthError struct {
	Reason  string // Why authentication failed
	TokenID string // Which token had the issue
}

func (a AuthError) Error() string {
	return fmt.Sprintf("Auth Error: %s (token: %s)", a.Reason, a.TokenID)
}

// ============================================================================
// CUSTOM ERROR TYPE 5: DatabaseError (With Helper Methods)
// ============================================================================
//
// This error type demonstrates helper methods for decision-making.

type DatabaseError struct {
	Operation string // SELECT, INSERT, UPDATE, DELETE
	Table     string // Which table
	Inner     error  // The original database error
}

func (d *DatabaseError) Error() string {
	return fmt.Sprintf("Database Error: %s on table '%s', caused by: %v",
		d.Operation, d.Table, d.Inner)
}

// Helper method: Can we retry this operation?
func (d *DatabaseError) CanRetry() bool {
	// Only retry SELECT operations (safe to retry reads)
	// Don't retry writes (INSERT, UPDATE, DELETE) to avoid duplicates
	return d.Operation == "SELECT"
}

// Helper method: Is this a timeout error?
func (d *DatabaseError) IsTimeout() bool {
	return d.Inner.Error() == "timeout"
}

// ============================================================================
// HELPER FUNCTION 3: getUserByID - Demonstrates Real-World Error Handling
// ============================================================================
//
// This simulates an API endpoint that fetches a user.
// Shows how custom errors help in real applications.

func getUserByID(id string) *struct{ ID, Name string } {
	// Validate input
	if id == "" || id == "invalid-id" {
		err := ValidationError{
			Field: "id",
			Issue: "invalid format",
			Value: id,
		}
		fmt.Printf("getUserByID error: %v\n", err)
		return nil
	}

	// In real code, this would query a database
	return &struct{ ID, Name string }{ID: id, Name: "John Doe"}
}

// ============================================================================
// HELPER FUNCTION 4: handleErrorByType - Type Switching
// ============================================================================
//
// Shows how to handle different error types differently.

func handleErrorByType(operation string, err error) {
	fmt.Printf("\nOperation: %s\n", operation)

	// Switch on error type
	switch e := err.(type) {
	case ValidationError:
		fmt.Printf("  → Validation error on field '%s': %s\n", e.Field, e.Issue)
		fmt.Println("  → Action: Return 400 Bad Request to client")

	case AuthError:
		fmt.Printf("  → Auth error: %s\n", e.Reason)
		fmt.Println("  → Action: Return 401 Unauthorized to client")

	case *DatabaseError:
		fmt.Printf("  → Database error: %s on %s\n", e.Operation, e.Table)
		if e.CanRetry() {
			fmt.Println("  → Action: Retry the operation")
		} else {
			fmt.Println("  → Action: Fail fast (unsafe to retry)")
		}

	default:
		fmt.Printf("  → Unknown error: %v\n", err)
		fmt.Println("  → Action: Return 500 Internal Server Error")
	}
}

// ============================================================================
// COMPREHENSIVE PATTERN EXAMPLES
// ============================================================================
//
// Pattern 1: Simple Custom Error (No Wrapping)
//   err := SomeFunc()
//   if err != nil {
//       return err  // Just return it
//   }
//
// Pattern 2: Wrapped Custom Error (Add Context)
//   err := SomeFunc()
//   if err != nil {
//       return &CustomError{
//           Code:    500,
//           Message: "operation context",
//           Err:     err,  // Preserve original!
//       }
//   }
//
// Pattern 3: Type Assertion (Extract Information)
//   if ve, ok := err.(ValidationError); ok {
//       fmt.Printf("Field: %s, Issue: %s\n", ve.Field, ve.Issue)
//   }
//
// Pattern 4: Helper Methods (Decision Making)
//   if dbErr, ok := err.(*DatabaseError); ok {
//       if dbErr.CanRetry() {
//           retry()
//       }
//   }
//
// ============================================================================
