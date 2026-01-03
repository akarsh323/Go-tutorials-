package intermediate

import (
	"fmt"
)

// Topic 69: custom_errors
// Combines pragmatic examples with GoBootcamp patterns


func main() {

	fmt.Println("-- 69 Custom Errors --")

	// Use custom error
	var e error = DemoError69{Msg: "bad thing", Code: 500}
	fmt.Println(e)

	// Another custom error
	var ve error = ValidationError{Field: "email", Issue: "invalid format"}
	fmt.Println(ve)

	// Type assert to get fields
	if err, ok := e.(DemoError69); ok {
		fmt.Println("error code:", err.Code)
	}
}

func (e DemoError69) Error() string {
	return fmt.Sprintf("DemoError69 (%d): %s", e.Code, e.Msg)
}

func customErrorExample10() {
	
	fmt.Println(`
Use custom errors when you need to:
  ✓ Check error types (type assertion)
  ✓ Carry business context (userID, operation, field)
  ✓ Provide methods for decision making (CanRetry, ShouldAlert)
  ✓ Distinguish between error categories (validation vs database)
  ✓ Wrap errors with additional context
  ✓ Support structured error responses (JSON, logs)

Simple rule:
  - errors.New() → Simple string-only errors
  - Custom error type → Need to check type or carry context
	`)
}
