package intermediate

import (
	"fmt"
	"strings"
)

// Topic 71: string_formatting
// Combines pragmatic examples with GoBootcamp patterns


func main() {

	fmt.Println("-- 71 String Formatting --")

	// Sprintf: Create formatted string without printing
	name := "Gopher"
	greeting := fmt.Sprintf("Hello, %s!", name)
	fmt.Println("Sprintf result:", greeting)

	// Printf: Formatted output directly to stdout
	age := 5
	fmt.Printf("Name: %s, Age: %d\n", name, age)

	// Multiple subtopic: Building strings efficiently
	// Concatenation with + operator (simple, but creates new strings each time)
	str1 := "Hello"
	str2 := "World"
	result := str1 + " " + str2
	fmt.Println("Concatenation with +:", result)

	// Sprintf for complex formatting
	data := map[string]interface{}{"a": 1, "b": "text", "c": 3.14}
	formatted := fmt.Sprintf("Data: %v", data)
	fmt.Println("Sprintf with map:", formatted)

	// strings.Builder: Efficient string building (better for loops)
	var builder strings.Builder
	builder.WriteString("Building")
	builder.WriteString(" ")
	builder.WriteString("a")
	builder.WriteString(" ")
	builder.WriteString("string")
	fmt.Println("strings.Builder:", builder.String())

	// Combining multiple values
	components := []string{"Alpha", "Beta", "Gamma"}
	joined := strings.Join(components, " -> ")
	fmt.Println("Joined with delimiter:", joined)

	// Width and precision formatting
	fmt.Printf("Width 10, left-aligned: |%-10s|\n", "Go")
	fmt.Printf("Width 10, right-aligned: |%10s|\n", "Go")
	fmt.Printf("Float precision: %.2f\n", 3.14159)
}

func stringFormattingExample1() {
	
	fmt.Println("📚 Convert numbers to their string representations")

	// Itoa (integer to ASCII) - quick for ints
	age := 30
	ageStr := fmt.Sprint(age) // or use strconv.Itoa
	fmt.Printf("Age as string: %q (type: %T)\n", ageStr, ageStr)

	// Sprintf for formatted conversion
	price := 19.99
	priceStr := fmt.Sprintf("$%.2f", price)
	fmt.Printf("Price: %s\n", priceStr)

	// Bool to string
	isActive := true
	statusStr := fmt.Sprint(isActive)
	fmt.Printf("Status: %s\n", statusStr)
}

func stringFormattingExample10() {
	
	fmt.Println(`
Format Verbs Quick Reference:

Type Conversion:
  %v    → value (default)           "hello"
  %T    → type name                 string
  %#v   → Go syntax                 "hello"
  %t    → boolean                   true/false

String/Char:
  %s    → string value              hello
  %q    → quoted string             "hello"
  %c    → character from code point A

Integer:
  %d    → decimal                   42
  %x,%X → hex (lower/upper)         2a/2A
  %b    → binary                    101010
  %o    → octal                     52

Float:
  %f    → decimal (default 6 places) 3.140000
  %.2f  → decimal (2 places)        3.14
  %e    → scientific notation       3.140000e+00

Formatting Tips:
  ✓ Use %q for strings with special chars
  ✓ Use %#v for debugging complex types
  ✓ Use %-10s for left-aligned, %10s for right-aligned
  ✓ Use %03d for zero-padded numbers
  ✓ Combine with strings.Builder for efficiency
	`)
}
