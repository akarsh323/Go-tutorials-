package intermediate

import (
	"fmt"
	"regexp"
)

// ============================================================================
// Topic 71: STRING FORMATTING AND INTERPOLATION
// ============================================================================
//
// CORE CONCEPT: Go provides two main approaches to string formatting:
//
//   1. FORMAT SPECIFIERS (fmt package)
//      - Printf, Sprintf, etc.
//      - Uses % verbs to specify type and width
//      - Example: fmt.Printf("%5d", 42) → "   42"
//
//   2. STRING LITERALS (double quotes vs backticks)
//      - Double quotes: Interpret escape sequences (\n, \t, etc.)
//      - Backticks: Treat everything as-is (raw strings)
//      - Critical for regex, SQL, multiline strings
//
// ============================================================================

func main() {
	fmt.Println("=============== 71 String Formatting and Interpolation ===============")

	// ========================================================================
	// SECTION 1: The fmt Package Fundamentals
	// ========================================================================
	fmt.Println("\n--- SECTION 1: fmt Package Overview ---")
	demonstrateFmtPackage()

	// ========================================================================
	// SECTION 2: Format Specifiers (% Verbs)
	// ========================================================================
	fmt.Println("\n--- SECTION 2: Format Specifiers (% Verbs) ---")
	demonstrateFormatSpecifiers()

	// ========================================================================
	// SECTION 3: Number Formatting (Padding and Width)
	// ========================================================================
	fmt.Println("\n--- SECTION 3: Number Formatting (Padding and Width) ---")
	demonstrateNumberFormatting()

	// ========================================================================
	// SECTION 4: String Alignment (Left vs Right)
	// ========================================================================
	fmt.Println("\n--- SECTION 4: String Alignment ---")
	demonstrateStringAlignment()

	// ========================================================================
	// SECTION 5: Float Precision Formatting
	// ========================================================================
	fmt.Println("\n--- SECTION 5: Float Precision ---")
	demonstrateFloatPrecision()

	// ========================================================================
	// SECTION 6: Type Formatting (%v, %T, %#v)
	// ========================================================================
	fmt.Println("\n--- SECTION 6: Type Formatting ---")
	demonstrateTypeFormatting()

	// ========================================================================
	// SECTION 7: String Literals - Double Quotes vs Backticks
	// ========================================================================
	fmt.Println("\n--- SECTION 7: String Literals (Quotes vs Backticks) ---")
	demonstrateStringLiterals()

	// ========================================================================
	// SECTION 8: Use Case - Regex and Raw Strings
	// ========================================================================
	fmt.Println("\n--- SECTION 8: Regex and Raw Strings ---")
	demonstrateRegexStrings()

	// ========================================================================
	// SECTION 9: Use Case - SQL and Multiline Strings
	// ========================================================================
	fmt.Println("\n--- SECTION 9: SQL and Multiline Strings ---")
	demonstrateSQLStrings()

	// ========================================================================
	// SECTION 10: Practical Example - Formatted Receipt
	// ========================================================================
	fmt.Println("\n--- SECTION 10: Real-World Example - Formatted Receipt ---")
	demonstrateFormattedReceipt()

	// ========================================================================
	// SECTION 11: Best Practices Summary
	// ========================================================================
	fmt.Println("\n--- SECTION 11: Best Practices ---")
	fmt.Println(`
DO:
  * Use %v for generic values, %T for type inspection
  * Use %d for integers, %f for floats, %s for strings
  * Use %05d for zero-padded integers
  * Use %-10s for left-aligned, %10s for right-aligned
  * Use backticks for regex patterns and multiline strings
  * Use %.2f for floating-point precision control
  * Use fmt.Sprintf for string creation without printing
  * Use %q for quoted strings (with escape sequences visible)

DON'T:
  * Mix %d with strings (use %s instead)
  * Use double quotes for regex patterns (use backticks)
  * Assume alignment without testing in monospace font
  * Use + operator in loops (use strings.Builder or fmt.Sprintf)
  * Forget newlines in Printf (it doesn't add them automatically)
  * Use %v when you need type-specific formatting

KEY INSIGHT:
  Format specifiers are about DISPLAY, not CONVERSION
  Use them to control how values appear in output
`)
}

// ============================================================================
// SECTION 1: The fmt Package
// ============================================================================

func demonstrateFmtPackage() {
	fmt.Println(`
Three Main Printf Functions:

1. fmt.Print() - Basic output (no formatting)
   • Prints arguments separated by spaces
   • No newline at end
   • No % verbs

2. fmt.Printf() - Formatted output (to stdout)
   • Uses format specifiers (% verbs)
   • No newline at end (must add \n yourself)
   • Output goes directly to terminal

3. fmt.Sprintf() - Formatted string (returns string)
   • Uses same format specifiers
   • Returns string instead of printing
   • No newline added
   • Store result in variable for later use

Example:`)
	name := "Alice"
	age := 30

	fmt.Print("Using Print: ", name, " ", age, "\n")
	fmt.Printf("Using Printf: %s is %d years old\n", name, age)
	result := fmt.Sprintf("Using Sprintf: %s is %d years old", name, age)
	fmt.Println("Stored result:", result)
}

// ============================================================================
// SECTION 2: Format Specifiers (% Verbs)
// ============================================================================

func demonstrateFormatSpecifiers() {
	fmt.Println(`
%v and %T - Generic Value and Type:
`)
	values := []interface{}{"hello", 42, 3.14, true, []int{1, 2, 3}}
	for _, v := range values {
		fmt.Printf("  Value: %v (Type: %T)\n", v, v)
	}

	fmt.Println(`
Strings and Characters:
  %s    → string value
  %q    → quoted string (with escape sequences)
  %c    → character from Unicode code point
`)
	fmt.Printf("  %%s: %s\n", "Hello")
	fmt.Printf("  %%q: %q\n", "Hello\nWorld")
	fmt.Printf("  %%c: %c\n", 65) // ASCII 65 = 'A'

	fmt.Println(`
Integers (Different Bases):
  %d    → decimal (base 10)
  %x    → hexadecimal lowercase
  %X    → hexadecimal uppercase
  %b    → binary (base 2)
  %o    → octal (base 8)
`)
	num := 42
	fmt.Printf("  %%d: %d\n", num)
	fmt.Printf("  %%x: %x\n", num)
	fmt.Printf("  %%X: %X\n", num)
	fmt.Printf("  %%b: %b\n", num)
	fmt.Printf("  %%o: %o\n", num)

	fmt.Println(`
Boolean:
  %t    → true or false
`)
	fmt.Printf("  %%t: %t\n", true)
	fmt.Printf("  %%t: %t\n", false)
}

// ============================================================================
// SECTION 3: Number Formatting (Padding and Width)
// ============================================================================

func demonstrateNumberFormatting() {
	fmt.Println(`
Basic Width Formatting:
  %5d   → minimum width of 5 (right-aligned, padded with spaces)
  %05d  → minimum width of 5 (padded with ZEROS)
  %-5d  → minimum width of 5 (left-aligned)

Example: ID Numbers
`)
	ids := []int{1, 42, 1234, 12345, 123456}
	fmt.Println("  Standard (%d):")
	for _, id := range ids {
		fmt.Printf("    %d\n", id)
	}

	fmt.Println("  With width 5 (%5d):")
	for _, id := range ids {
		fmt.Printf("    |%5d|\n", id)
	}

	fmt.Println("  With zero-padding (%05d):")
	for _, id := range ids {
		fmt.Printf("    |%05d|\n", id)
	}

	fmt.Println("  Left-aligned (%-5d):")
	for _, id := range ids {
		fmt.Printf("    |%-5d|\n", id)
	}

	fmt.Println(`
Real-World Use Case: Timestamps
`)
	hour, minute, second := 9, 5, 3
	timestamp := fmt.Sprintf("%02d:%02d:%02d", hour, minute, second)
	fmt.Printf("  Timestamp: %s\n", timestamp)
}

// ============================================================================
// SECTION 4: String Alignment
// ============================================================================

func demonstrateStringAlignment() {
	fmt.Println(`
String Width Formatting:
  %10s  → minimum width 10, right-aligned (padding on left)
  %-10s → minimum width 10, left-aligned (padding on right)

Key: Use pipes | | to visualize whitespace
`)
	names := []string{"Go", "Python", "Rust", "JavaScript"}

	fmt.Println("Right-aligned (%10s):")
	for _, name := range names {
		fmt.Printf("  |%10s|\n", name)
	}

	fmt.Println("Left-aligned (%-10s):")
	for _, name := range names {
		fmt.Printf("  |%-10s|\n", name)
	}

	fmt.Println(`
Real-World Use Case: Table Headers
`)
	fmt.Println("  Column Format:")
	fmt.Printf("  |%-15s|%10s|%10s|\n", "Product", "Price", "Quantity")
	fmt.Printf("  |%-15s|%10s|%10s|\n", "---", "---", "---")
	fmt.Printf("  |%-15s|%10.2f|%10d|\n", "Apple", 1.50, 5)
	fmt.Printf("  |%-15s|%10.2f|%10d|\n", "Banana", 0.75, 12)
	fmt.Printf("  |%-15s|%10.2f|%10d|\n", "Orange", 2.00, 8)
}

// ============================================================================
// SECTION 5: Float Precision
// ============================================================================

func demonstrateFloatPrecision() {
	fmt.Println(`
Float Formatting:
  %f     → decimal notation (default 6 decimal places)
  %.2f   → exactly 2 decimal places
  %.0f   → no decimal places (rounded)
  %e     → scientific notation
  %g     → compact notation (switches between %f and %e)
`)
	value := 3.14159265

	fmt.Printf("  %%f:    %f\n", value)
	fmt.Printf("  %%.2f:  %.2f\n", value)
	fmt.Printf("  %%.0f:  %.0f\n", value)
	fmt.Printf("  %%e:    %e\n", value)
	fmt.Printf("  %%g:    %g\n", value)

	fmt.Println(`
Currency Formatting:
`)
	prices := []float64{9.5, 19.99, 100.00, 1234.5678}
	for _, price := range prices {
		fmt.Printf("  $%.2f\n", price)
	}

	fmt.Println(`
Width + Precision:
  %10.2f → width 10, 2 decimal places
`)
	for _, price := range prices {
		fmt.Printf("  |%10.2f|\n", price)
	}
}

// ============================================================================
// SECTION 6: Type Formatting
// ============================================================================

func demonstrateTypeFormatting() {
	fmt.Println(`
%v - Default Value (Smart Formatting)
  → Shows value in a reasonable way
  → Works with ANY type
`)
	fmt.Printf("  String: %v\n", "hello")
	fmt.Printf("  Integer: %v\n", 42)
	fmt.Printf("  Float: %v\n", 3.14)
	fmt.Printf("  Slice: %v\n", []int{1, 2, 3})
	fmt.Printf("  Map: %v\n", map[string]int{"a": 1, "b": 2})

	fmt.Println(`
%T - Type Name
  → Shows the TYPE, not the value
  → Useful for debugging
`)
	fmt.Printf("  Type of \"hello\": %T\n", "hello")
	fmt.Printf("  Type of 42: %T\n", 42)
	fmt.Printf("  Type of 3.14: %T\n", 3.14)
	fmt.Printf("  Type of []int: %T\n", []int{1, 2, 3})

	fmt.Println(`
%#v - Go Syntax
  → Shows value as valid Go code
  → Useful for printing out data structures
`)
	fmt.Printf("  %#v\n", "hello")
	fmt.Printf("  %#v\n", []int{1, 2, 3})
	fmt.Printf("  %#v\n", map[string]int{"a": 1})
}

// ============================================================================
// SECTION 7: String Literals - Quotes vs Backticks
// ============================================================================

func demonstrateStringLiterals() {
	fmt.Println(`
TWO TYPES OF STRING LITERALS:

1. DOUBLE QUOTES ("") - Interpreted String Literal
   • Go INTERPRETS escape sequences
   • \n becomes newline, \t becomes tab
   • Backslashes must be escaped (\\)
   • Single-line only (can't have literal newlines)

2. BACKTICKS ("") - Raw String Literal
   • Go IGNORES all escape sequences
   • \n stays as \n (two characters)
   • No escaping needed
   • Can span multiple lines
`)

	fmt.Println("Example 1: Newline character")
	msg1 := "Line 1\nLine 2"
	fmt.Printf("Double quotes: %s\n", msg1)

	msg2 := `Line 1\nLine 2`
	fmt.Printf("Backticks: %s\n", msg2)

	fmt.Println("\nExample 2: Backslash in Windows path")
	path1 := "C:\\Users\\Name\\Documents"
	fmt.Printf("Double quotes: %s\n", path1)

	path2 := `C:\Users\Name\Documents`
	fmt.Printf("Backticks: %s\n", path2)

	fmt.Println("\nExample 3: Multiline string (backticks only)")
	description := `This is a description
that spans multiple lines
without needing concatenation
or \n characters`
	fmt.Printf("Backticks:\n%s\n", description)

	fmt.Println("\nQuote Behavior with %q:")
	fmt.Printf("Double quoted with %%q: %q\n", "Hello\nWorld")
	fmt.Printf("Backtick quoted with %%q: %q\n", `Hello\nWorld`)
}

// ============================================================================
// SECTION 8: Regex and Raw Strings
// ============================================================================

func demonstrateRegexStrings() {
	fmt.Println(`
WHY BACKTICKS FOR REGEX?

Regex patterns use backslashes:
  \d  → digit
  \w  → word character
  \s  → whitespace

Double Quotes (Wrong):
  regexp.MustCompile("\\\\d+")  ← Must escape TWICE!

Backticks (Correct):
  regexp.MustCompile(` + "`" + `\d+` + "`" + `)  ← Natural!
`)

	fmt.Println("Example 1: Find all digits")
	text := "I have 2 apples and 15 oranges"
	re1 := regexp.MustCompile(`\d+`) // Backticks!
	matches1 := re1.FindAllString(text, -1)
	fmt.Printf("Text: %q\n", text)
	fmt.Printf("Numbers found: %v\n", matches1)

	fmt.Println("\nExample 2: Find words starting with capital letter")
	text2 := "Go is Python and Rust"
	re2 := regexp.MustCompile(`[A-Z]\w*`) // Backticks for natural regex!
	matches2 := re2.FindAllString(text2, -1)
	fmt.Printf("Text: %q\n", text2)
	fmt.Printf("Capitalized words: %v\n", matches2)

	fmt.Println("\nExample 3: Email pattern")
	emailPattern := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	re3 := regexp.MustCompile(emailPattern)
	emails := []string{"user@example.com", "invalid.email", "test@domain.co.uk"}
	for _, email := range emails {
		fmt.Printf("  %q is valid: %v\n", email, re3.MatchString(email))
	}
}

// ============================================================================
// SECTION 9: SQL and Multiline Strings
// ============================================================================

func demonstrateSQLStrings() {
	fmt.Println(`
WHY BACKTICKS FOR SQL?

SQL queries are often long and multiline:
  • Preserve formatting
  • No need for + concatenation
  • No need for \n escape sequences
  • More readable
`)

	fmt.Println("Example 1: Basic SQL query")
	query1 := `
		SELECT id, name, email
		FROM users
		WHERE age > 30
		ORDER BY name
	`
	fmt.Printf("Query:\n%s\n", query1)

	fmt.Println("Example 2: Complex SQL with JOIN")
	query2 := `
		SELECT 
		    u.id,
		    u.name,
		    COUNT(o.id) as order_count
		FROM users u
		LEFT JOIN orders o ON u.id = o.user_id
		WHERE u.status = 'active'
		GROUP BY u.id
		HAVING COUNT(o.id) > 0
		ORDER BY order_count DESC
	`
	fmt.Printf("Query:\n%s\n", query2)

	fmt.Println("Example 3: HTML template")
	htmlTemplate := `
		<!DOCTYPE html>
		<html>
		<head>
		    <title>My Page</title>
		</head>
		<body>
		    <h1>Welcome</h1>
		    <p>This is a paragraph.</p>
		</body>
		</html>
	`
	fmt.Printf("HTML:\n%s\n", htmlTemplate)

	fmt.Println("Example 4: JSON document")
	jsonData := `
		{
		    "name": "Alice",
		    "age": 30,
		    "email": "alice@example.com",
		    "active": true
		}
	`
	fmt.Printf("JSON:\n%s\n", jsonData)
}

// ============================================================================
// SECTION 10: Real-World Example - Formatted Receipt
// ============================================================================

func demonstrateFormattedReceipt() {
	fmt.Println(`
Create a nicely formatted receipt using alignment and width:
`)
	fmt.Println()

	// Receipt header
	fmt.Println("═══════════════════════════════════════")
	fmt.Println("           STORE RECEIPT")
	fmt.Println("═══════════════════════════════════════")

	// Items header
	fmt.Printf("%-20s %10s %10s\n", "ITEM", "PRICE", "TOTAL")
	fmt.Println("───────────────────────────────────────")

	// Items
	items := []struct {
		name     string
		price    float64
		quantity int
	}{
		{"Apple", 1.50, 5},
		{"Banana", 0.75, 12},
		{"Orange", 2.00, 8},
		{"Milk", 3.50, 2},
		{"Bread", 2.25, 1},
	}

	totalAmount := 0.0
	for _, item := range items {
		total := item.price * float64(item.quantity)
		totalAmount += total
		fmt.Printf("%-20s %10.2f %10.2f\n", item.name, item.price, total)
	}

	// Footer
	fmt.Println("───────────────────────────────────────")
	fmt.Printf("%-20s %20.2f\n", "TOTAL:", totalAmount)
	fmt.Println("═══════════════════════════════════════")
	fmt.Printf("%-20s %20s\n", "Thank you!", "Come again!")
	fmt.Println()
}

// ============================================================================
// COMPREHENSIVE FORMAT SPECIFIER REFERENCE
// ============================================================================
//
// VALUE DISPLAY:
//   %v    → value (default)
//   %T    → type name
//   %#v   → Go syntax representation
//
// STRINGS & CHARACTERS:
//   %s    → string value
//   %q    → quoted string
//   %c    → character (from code point)
//
// INTEGERS (Different Bases):
//   %d    → decimal (base 10)
//   %x    → hexadecimal (lowercase)
//   %X    → hexadecimal (uppercase)
//   %o    → octal (base 8)
//   %b    → binary (base 2)
//
// FLOATS:
//   %f    → decimal notation
//   %.2f  → 2 decimal places
//   %e    → scientific notation
//   %g    → compact notation
//
// BOOLEANS:
//   %t    → true or false
//
// WIDTH & PRECISION:
//   %5d     → width 5, right-aligned
//   %-5d    → width 5, left-aligned
//   %05d    → width 5, zero-padded
//   %10s    → width 10 string, right-aligned
//   %-10s   → width 10 string, left-aligned
//   %.2f    → 2 decimal places
//   %10.2f  → width 10, 2 decimal places
//
// STRING LITERALS:
//   ""      → interpreted (understands \n, \t, \\, etc.)
//   ``      → raw (ignores escape sequences)
//
// ============================================================================
