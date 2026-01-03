package main

import "fmt"

/*
Topic 62: FMT PACKAGE

CONCEPT: The fmt package provides functions for formatted I/O (input/output).

PRINT FUNCTIONS (Printing to console):
- Print(v...)         - Output values
- Println(v...)       - Output values + newline
- Printf(format, ...) - Output with format string

FORMAT FUNCTIONS (Save to string):
- Sprint(v...)         - Return formatted string
- Sprintln(v...)       - Return formatted string + newline
- Sprintf(format, ...) - Return formatted string

SCAN FUNCTIONS (Reading input):
- Scan(&v...)         - Read from console
- Scanln(&v...)       - Read line from console
- Scanf(format, &v...) - Read formatted input

POINTERS WITH SCAN:
CRITICAL: Scan functions need POINTERS (addresses) to store values!
Use & to get address: Scan(&variable)

ERROR HANDLING:
- Errorf(format, ...) - Create error messages

ANALOGY: Jar of cookies
- Variable = the jar (contains value)
- Pointer = the address/location label on jar
- & = "get the address of this jar"
- * = "dereference, get what's at this address"
*/

// Helper function for error handling example
func checkAge(age int) error {
	if age < 0 {
		return fmt.Errorf("invalid age: %d (age cannot be negative)", age)
	}
	if age < 18 {
		return fmt.Errorf("age %d: must be at least 18 (got %d)", age, age)
	}
	return nil
}

func main() {
	fmt.Println("═══════════════════════════════════════════════════════════")
	fmt.Println("PART 1: PRINT FUNCTIONS (Output to Console)")
	fmt.Println("═══════════════════════════════════════════════════════════\n")

	fmt.Println("Print() - outputs values separated by spaces:")
	fmt.Print("Hello", "World", 42)
	fmt.Print("\n\n")

	fmt.Println("Println() - outputs values + newline:")
	fmt.Println("Hello", "World", 42)

	fmt.Println("Printf() - formatted output:")
	name := "Alice"
	age := 25
	score := 95.5
	fmt.Printf("Name: %s, Age: %d, Score: %.1f%%\n\n", name, age, score)

	// ─────────────────────────────────────────────────────────────────────────

	fmt.Println("═══════════════════════════════════════════════════════════")
	fmt.Println("PART 2: SPRINT FUNCTIONS (Save to String)")
	fmt.Println("═══════════════════════════════════════════════════════════\n")

	fmt.Println("Sprint() - returns formatted string:")
	str1 := fmt.Sprint("Hello", "World", 42)
	fmt.Printf("str1 = \"%s\" (type: %T)\n", str1, str1)

	fmt.Println("\nSprintln() - returns formatted string + newline:")
	str2 := fmt.Sprintln("Hello", "World", 42)
	fmt.Printf("str2 = \"%s\" (type: %T)\n", str2, str2)

	fmt.Println("Sprintf() - formatted string:")
	name2 := "Bob"
	score2 := 88.5
	str3 := fmt.Sprintf("Name: %s, Score: %.1f%%", name2, score2)
	fmt.Printf("str3 = \"%s\"\n\n", str3)

	// ─────────────────────────────────────────────────────────────────────────

	fmt.Println("═══════════════════════════════════════════════════════════")
	fmt.Println("PART 3: POINTER CONCEPT (Why Scan Needs &)")
	fmt.Println("═══════════════════════════════════════════════════════════\n")

	fmt.Println(`
ANALOGY: Box with contents and mailing address

Variable (var x int):
- x = the box containing value (initially 0)
- The box itself

Pointer (&x):
- &x = the mailing ADDRESS of the box
- Where to find the box in memory
- Type: *int (pointer to int)

Dereferencing (*ptr):
- *ptr = look inside the box at this address
- Get the value at the address

MEMORY VISUALIZATION:
Variable: x = 42
  └─ Box labeled "x" containing: 42

Address of x: &x = 0x1234 (memory location)
  └─ Mailing address of the box

Pointer: ptr := &x
  └─ A note saying "the box is at 0x1234"

Dereference: *ptr = 42
  └─ Go to address 0x1234, get the value

WHY SCAN NEEDS POINTERS:
Scan needs to MODIFY your variable by filling it with input.
To modify, it needs the ADDRESS (&) not the value!

This is how Scan knows WHERE to store the value you input!
`)

	// ─────────────────────────────────────────────────────────────────────────

	fmt.Println("═══════════════════════════════════════════════════════════")
	fmt.Println("PART 4: SCAN FUNCTIONS (Reading Input)")
	fmt.Println("═══════════════════════════════════════════════════════════\n")

	fmt.Println("Scan() - reads space-separated values:")
	fmt.Println("(Example: would read input from console - commented out for demo)")
	fmt.Println("// var name string")
	fmt.Println("// var age int")
	fmt.Println("// fmt.Scan(&name, &age)")

	fmt.Println("\nScanln() - reads until newline:")
	fmt.Println("// var input string")
	fmt.Println("// fmt.Scanln(&input)")

	fmt.Println("\nScanf() - reads formatted input:")
	fmt.Println("// var name string")
	fmt.Println("// var age int")
	fmt.Println("// fmt.Scanf(\"%s %d\", &name, &age)")
	fmt.Println()

	// ─────────────────────────────────────────────────────────────────────────

	fmt.Println("═══════════════════════════════════════════════════════════")
	fmt.Println("PART 5: ERROR HANDLING (Errorf)")
	fmt.Println("═══════════════════════════════════════════════════════════\n")

	ages := []int{-5, 15, 25, 40}

	fmt.Println("Validating ages with Errorf():\n")
	for _, a := range ages {
		err := checkAge(a)
		if err != nil {
			fmt.Printf("Age %d: ❌ %v\n", a, err)
		} else {
			fmt.Printf("Age %d: ✓ Valid\n", a)
		}
	}
	fmt.Println()

	// ─────────────────────────────────────────────────────────────────────────

	fmt.Println("═══════════════════════════════════════════════════════════")
	fmt.Println("FMT PACKAGE FUNCTIONS SUMMARY TABLE")
	fmt.Println("═══════════════════════════════════════════════════════════\n")

	fmt.Println("┌──────────┬────────────────────────┬──────────────────────┐")
	fmt.Println("│ Function │ Purpose                │ Returns/Use          │")
	fmt.Println("├──────────┼────────────────────────┼──────────────────────┤")
	fmt.Println("│ Print    │ Output to console      │ Outputs directly     │")
	fmt.Println("│ Println  │ Output + newline       │ Outputs directly     │")
	fmt.Println("│ Printf   │ Formatted output       │ Outputs directly     │")
	fmt.Println("│ Sprint   │ Format to string       │ Returns string       │")
	fmt.Println("│ Sprintln │ Format + newline       │ Returns string       │")
	fmt.Println("│ Sprintf  │ Formatted to string    │ Returns string       │")
	fmt.Println("│ Scan     │ Read space-separated   │ Stores in variables  │")
	fmt.Println("│ Scanln   │ Read line              │ Stores in variables  │")
	fmt.Println("│ Scanf    │ Read formatted         │ Stores in variables  │")
	fmt.Println("│ Errorf   │ Create error message   │ Returns error        │")
	fmt.Println("└──────────┴────────────────────────┴──────────────────────┘\n")

	// ─────────────────────────────────────────────────────────────────────────

	fmt.Println("═══════════════════════════════════════════════════════════")
	fmt.Println("PRACTICAL EXAMPLES")
	fmt.Println("═══════════════════════════════════════════════════════════\n")

	fmt.Println("Example 1: Build log message")
	logMsg := fmt.Sprintf("[INFO] User %s logged in at %s", "alice", "2024-01-15")
	fmt.Println(logMsg)

	fmt.Println("\nExample 2: Format table for display")
	headers := fmt.Sprintf("%-10s | %8s", "Name", "Score")
	fmt.Println(headers)
	fmt.Println("-----------|----------")
	fmt.Printf("%-10s | %8.1f\n", "Alice", 95.5)
	fmt.Printf("%-10s | %8.1f\n", "Bob", 88.0)

	fmt.Println("\nExample 3: Error message creation")
	operation := "delete"
	resource := "file.txt"
	permError := fmt.Errorf("permission denied: cannot %s %s (admin only)", operation, resource)
	fmt.Printf("Error: %v\n\n", permError)

	// ─────────────────────────────────────────────────────────────────────────

	fmt.Println("═══════════════════════════════════════════════════════════")
	fmt.Println("KEY CONCEPTS & BEST PRACTICES")
	fmt.Println("═══════════════════════════════════════════════════════════\n")

	fmt.Println(`
PRINT FUNCTIONS:
Print()   - outputs, NO newline, spaces between args
Println() - outputs, WITH newline, spaces between args
Printf()  - formatted output with format string

SPRINT FUNCTIONS (Save to string):
Sprint()   - like Print but returns string
Sprintln() - like Println but returns string
Sprintf()  - like Printf but returns string
Use when you need the string for later use (logging, storage, etc.)

SCAN FUNCTIONS (Input):
Scan()   - read space-separated values from console
Scanln() - read until newline
Scanf()  - read with format string

CRITICAL POINTER RULES:
✓ Scan functions REQUIRE pointers (&variable)
✓ & means "address of" (where the variable lives)
✓ Scan needs addresses to STORE values in your variables
✓ Without &, Scan won't work!

ERROR CREATION:
✓ Use Errorf() to create formatted error messages
✓ Returns nil if no error, error if something failed
✓ Format just like Printf but returns error type

BEST PRACTICES:
✓ Use Println() for simple output with newline
✓ Use Printf() for formatted output (currency, tables)
✓ Use Sprintf() when you need string for later use
✓ ALWAYS use & with Scan (common beginner mistake!)
✓ Use Errorf() for meaningful error messages
✓ Format errors consistently across your application
✓ Use fmt.Print with multiple values for clean output

COMMON PITFALLS:
✗ Forgetting \n with Printf
✗ Forgetting & with Scan
✗ Using Print instead of Println (no automatic newline)
✗ Wrong format verb for type (%d for string)
✗ Not using Errorf for error creation

MEMORY TIP:
"S" in Sprint/Sprintf = "Save to string"
"f" functions need format strings (Printf, Sprintf, Scanf)
"ln" functions add newlines (Println, Sprintln, Scanln)
	`)
}
