package main

import (
	"fmt"
)

/*
Topic 62: THE FMT PACKAGE - A Comprehensive Guide

═══════════════════════════════════════════════════════════════════════════════

WHAT IS THE FMT PACKAGE?

The `fmt` (Format) package is the most frequently used package in Go.
It handles three major tasks:

1. PRINTING - Write text directly to the console (Standard Output)
2. FORMATTING - Create formatted strings and save them to variables
3. SCANNING - Read input from the user (keyboard/Standard Input)

═══════════════════════════════════════════════════════════════════════════════

PART 1: PRINTING TO CONSOLE (OUTPUT)

These functions write text directly to "Standard Output" (your terminal).

═════════════════════════════════════════════════════════════════════════════

FUNCTION: fmt.Print()
- Prints raw data as-is
- NO automatic newlines
- NO automatic spaces between arguments (unless explicitly provided)
- Best for: Raw output without formatting

FUNCTION: fmt.Println()
- ADDS spaces between arguments automatically
- ADDS a newline at the end
- Easy to use
- Best for: Debugging, simple output, logging

FUNCTION: fmt.Printf()
- Uses FORMAT VERBS (%s, %d, %f, etc.) for precise formatting
- NO automatic newline (must use \n)
- Full control over output
- Best for: Structured, formatted output

═════════════════════════════════════════════════════════════════════════════

PART 2: FORMATTING STRINGS (THE "S" FAMILY)

Sometimes you don't want to PRINT the text directly.
Instead, you want to SAVE it to a variable for later use
(e.g., sending to database, API, file, etc.)

The functions are identical to printing, but start with "S" (String)
and RETURN the result instead of printing.

FUNCTION: fmt.Sprint()
- Like Print(), but returns a string
- Best for: Saving raw formatted text to variable

FUNCTION: fmt.Sprintln()
- Like Println(), but returns a string
- Adds spaces and newline to the returned string
- Best for: Saving multi-argument output to variable

FUNCTION: fmt.Sprintf()
- Like Printf(), but returns a string
- Uses format verbs for precise formatting
- NO automatic newline (add \n if needed)
- Best for: Saving formatted strings with control

═════════════════════════════════════════════════════════════════════════════

PART 3: SCANNING (USER INPUT)

This allows your program to PAUSE and WAIT for user input from keyboard.

CRUCIAL CONCEPT: POINTERS (&)

When using Scan functions, you MUST pass the ADDRESS of the variable using &

WHY?
- Scan functions need to know the exact memory location of the variable
- They write the user's input directly into that memory location
- Without &, Scan only gets a COPY of the variable
- Modifying the copy doesn't affect your original variable
- Result: Your original variable stays empty!

ANALOGY:
- Without &: You give Scan a photocopy of your empty jar
- Scan fills the photocopy
- You get back an empty jar
- With &: You give Scan the actual address of your jar
- Scan fills the ACTUAL jar
- You get back a filled jar

═════════════════════════════════════════════════════════════════════════════

FUNCTION: fmt.Scan()
- Reads space-separated values
- IGNORES newlines (treats Enter like a space)
- Waits until ALL variables are filled
- Flexible with input format
- Best for: Simple input across multiple lines

Example:
var name string
var age int
fmt.Scan(&name, &age)
// User can type: "John" [Enter] "25" [Enter] ✓ Works!
// OR: "John 25" [Enter] ✓ Also works!

═════════════════════════════════════════════════════════════════════════════

FUNCTION: fmt.Scanln()
- STOPS at newline (End of line)
- Expects all values on ONE line
- Strict about line boundaries
- Best for: Single-line input

Example:
var name string
fmt.Scanln(&name)
// User types: "John" [Enter] ✓ Works!
// User types only [Enter] ✗ Gets empty string and waits for next

═════════════════════════════════════════════════════════════════════════════

FUNCTION: fmt.Scanf()
- The STRICTEST input function
- Expects input to match format string EXACTLY
- Uses same format verbs as Printf (%d, %s, %f, etc.)
- Best for: Highly structured input (dates, measurements, etc.)

Example:
var day, month int
fmt.Scanf("%d/%d", &day, &month)
// User MUST type: "12/05" [Enter] ✓ Works!
// User types: "12 / 05" ✗ Fails (extra spaces not allowed)
// User types: "12-05" ✗ Fails (wrong delimiter)

═════════════════════════════════════════════════════════════════════════════

PART 4: ERROR FORMATTING (fmt.Errorf)

In Go, errors are values. You often need to create error messages dynamically.
`fmt.Errorf` works exactly like Sprintf, but returns an `error` type.

FUNCTION: fmt.Errorf()
- Creates error objects with dynamic messages
- Uses same format verbs as Printf (%d, %s, etc.)
- Returns type: error
- Best for: Creating meaningful, contextual error messages

═════════════════════════════════════════════════════════════════════════════
*/

func main() {
	fmt.Println("════════════════════════════════════════════════════════════")
	fmt.Println("PART 1: PRINTING TO CONSOLE")
	fmt.Println("════════════════════════════════════════════════════════════\n")

	name, age := "John", 25

	// 1. Print() - Raw, no newlines, no automatic spaces
	fmt.Println("Example 1: fmt.Print()")
	fmt.Print("Hello")
	fmt.Print("World")
	fmt.Print("\n") // Must add newline manually
	fmt.Println("Output: HelloWorld (No space, no newline added)\n")

	// 2. Println() - Adds spaces and newlines automatically
	fmt.Println("Example 2: fmt.Println()")
	fmt.Println("Hello", "World", "from", "Go")
	fmt.Println("Output: Hello World from Go (Spaces added automatically)\n")

	// 3. Printf() - Formatted control using format verbs
	fmt.Println("Example 3: fmt.Printf()")
	fmt.Printf("User %s is %d years old.\n", name, age)
	fmt.Printf("User %s will be %d next year.\n", name, age+1)
	fmt.Println("Output: Formatted with %s and %d\n")

	// ─────────────────────────────────────────────────────────────────────────

	fmt.Println("════════════════════════════════════════════════════════════")
	fmt.Println("PART 2: FORMATTING STRINGS (THE \"S\" FAMILY)")
	fmt.Println("════════════════════════════════════════════════════════════\n")

	// 1. Sprint() - Like Print() but returns string
	fmt.Println("Example 1: fmt.Sprint()")
	result1 := fmt.Sprint("Go", "is", "awesome")
	fmt.Printf("Result: %q\n", result1)
	fmt.Println("(Spaces NOT added, saved to variable)\n")

	// 2. Sprintln() - Like Println() but returns string
	fmt.Println("Example 2: fmt.Sprintln()")
	result2 := fmt.Sprintln("Go", "is", "awesome")
	fmt.Printf("Result: %q\n", result2)
	fmt.Println("(Spaces ADDED, saved to variable)\n")

	// 3. Sprintf() - Like Printf() but returns string
	fmt.Println("Example 3: fmt.Sprintf()")
	country := "Germany"
	capital := "Berlin"
	result3 := fmt.Sprintf("The capital of %s is %s.\n", country, capital)
	fmt.Printf("Result: %q\n", result3)
	fmt.Println("(Formatted and saved to variable)")

	// Practical use: Save to variable and reuse
	greeting := fmt.Sprintf("Hello, %s! You are %d years old.", name, age)
	fmt.Printf("Stored greeting: %s\n\n", greeting)

	// ─────────────────────────────────────────────────────────────────────────

	fmt.Println("════════════════════════════════════════════════════════════")
	fmt.Println("PART 3: SCANNING USER INPUT (COMMENTED FOR DEMO)")
	fmt.Println("════════════════════════════════════════════════════════════\n")

	fmt.Println(`
SCANNING FUNCTIONS (Not executed in demo, but explained):

fmt.Scan(&variable):
  - Reads space-separated values
  - Ignores newlines (treats Enter like space)
  - Keeps waiting until all variables filled
  - Flexible with input format
  
  Example:
  var name string
  var age int
  fmt.Scan(&name, &age)
  // User input: "John" [Enter] "25" [Enter] ✓

fmt.Scanln(&variable):
  - Stops at newline (End of line)
  - Expects all values on one line
  - More strict than Scan
  
  Example:
  var name string
  fmt.Scanln(&name)
  // User input: "John" [Enter] ✓

fmt.Scanf("%d/%d", &day, &month):
  - Strictest input function
  - Matches exact format
  - Uses format verbs like Printf
  
  Example:
  var day, month int
  fmt.Scanf("%d/%d", &day, &month)
  // User input: "12/05" [Enter] ✓
  // User input: "12 / 05" [Enter] ✗ (fails: extra spaces)

IMPORTANT: Always use & to pass variable address!
Without &, the function gets a copy, not the original variable.
`)

	// ─────────────────────────────────────────────────────────────────────────

	fmt.Println("════════════════════════════════════════════════════════════")
	fmt.Println("PART 4: ERROR FORMATTING (fmt.Errorf)")
	fmt.Println("════════════════════════════════════════════════════════════\n")

	// Example 1: Simple error
	err1 := fmt.Errorf("invalid age: %d", 10)
	fmt.Printf("Error 1: %v\n\n", err1)

	// Example 2: Error with context
	userID := 42
	err2 := fmt.Errorf("user %d not found in database", userID)
	fmt.Printf("Error 2: %v\n\n", err2)

	// Example 3: Error from function
	err3 := checkAge(15)
	if err3 != nil {
		fmt.Printf("Error 3: %v\n\n", err3)
	}

	// Example 4: Error with multiple values
	email := "invalid.email"
	err4 := fmt.Errorf("invalid email format: '%s' (expected format: user@domain.com)", email)
	fmt.Printf("Error 4: %v\n\n", err4)

	// ─────────────────────────────────────────────────────────────────────────

	fmt.Println("════════════════════════════════════════════════════════════")
	fmt.Println("SUMMARY TABLE: FMT PACKAGE FUNCTIONS")
	fmt.Println("════════════════════════════════════════════════════════════\n")

	fmt.Println("┌──────────────────┬────────────────────────┬──────────────────┐")
	fmt.Println("│ Function         │ Action                 │ Output           │")
	fmt.Println("├──────────────────┼────────────────────────┼──────────────────┤")
	fmt.Println("│ Print()          │ Print raw              │ Console          │")
	fmt.Println("│ Println()        │ Print with spaces & \\n │ Console          │")
	fmt.Println("│ Printf()         │ Formatted print        │ Console          │")
	fmt.Println("├──────────────────┼────────────────────────┼──────────────────┤")
	fmt.Println("│ Sprint()         │ Save raw               │ Variable         │")
	fmt.Println("│ Sprintln()       │ Save with spaces & \\n  │ Variable         │")
	fmt.Println("│ Sprintf()        │ Formatted save         │ Variable         │")
	fmt.Println("├──────────────────┼────────────────────────┼──────────────────┤")
	fmt.Println("│ Scan()           │ Read flexible input    │ Variable         │")
	fmt.Println("│ Scanln()         │ Read single line       │ Variable         │")
	fmt.Println("│ Scanf()          │ Read formatted input   │ Variable         │")
	fmt.Println("├──────────────────┼────────────────────────┼──────────────────┤")
	fmt.Println("│ Errorf()         │ Create error message   │ Error object     │")
	fmt.Println("└──────────────────┴────────────────────────┴──────────────────┘\n")

	// ─────────────────────────────────────────────────────────────────────────

	fmt.Println("════════════════════════════════════════════════════════════")
	fmt.Println("BEST PRACTICES & TIPS")
	fmt.Println("════════════════════════════════════════════════════════════\n")

	fmt.Println(`
PRINTING OUTPUT:
✓ Use Println() for debugging (easiest)
✓ Use Printf() for formatted output (precise control)
✓ Use Print() rarely (difficult to format correctly)
✓ Always include \n in Printf() if you want newlines

FORMATTING STRINGS:
✓ Use Sprint family to save formatted text to variables
✓ Useful for building messages, logs, database entries
✓ Use Sprintf() for precise formatting with format verbs
✓ Use Sprintln() for simple multi-argument formatting

SCANNING INPUT:
✓ ALWAYS use & with Scan functions (pass variable address!)
✓ Use Scan() for flexible multi-line input
✓ Use Scanln() for single-line input
✓ Use Scanf() only for strictly formatted input
✓ Remember: Scan functions block until input is received

ERROR HANDLING:
✓ Use Errorf() to create meaningful error messages
✓ Include context (variable names, values, expected format)
✓ Use same format verbs as Printf (%d, %s, %v, etc.)
✓ Errors should be informative for debugging

COMMON MISTAKES:
✗ Forgetting & in Scan(&variable) - Variable stays empty!
✗ Forgetting \n in Printf() - No newline added
✗ Using Print() for formatted output - Use Printf() instead
✗ Not checking Scan return value - May silently fail
✗ Scanf() too strict for user input - Users make typos!

PRINT vs PRINTLN vs PRINTF:
- Print("a", "b")   → ab (no spaces or newline)
- Println("a", "b") → a b\n (adds spaces and newline)
- Printf("%s %s\n", "a", "b") → a b\n (you control format)

REAL-WORLD EXAMPLES:

1. LOGGING:
   fmt.Printf("[%s] User %s logged in\n", time.Now().Format("15:04:05"), userName)

2. FORMATTING DATA FOR STORAGE:
   csvLine := fmt.Sprintf("%d,%s,%f\n", id, name, price)
   // Save to file later

3. ERROR HANDLING:
   if age < 18 {
       return fmt.Errorf("age %d is below minimum (18)", age)
   }

4. USER INTERACTION:
   fmt.Print("Enter your name: ")
   var name string
   fmt.Scanln(&name)
   fmt.Printf("Hello, %s!\n", name)
`)
}

// checkAge validates age and returns error if invalid
func checkAge(age int) error {
	if age < 18 {
		return fmt.Errorf("age %d is too low, minimum is 18", age)
	}
	return nil
}
