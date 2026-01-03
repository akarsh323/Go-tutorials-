package main

import "fmt"

/*
Topic 61: FORMATTING VERBS - A Comprehensive Guide

═══════════════════════════════════════════════════════════════════════════════

WHAT ARE FORMATTING VERBS?

Formatting verbs (also called format verbs or placeholders) are special sequences
in Printf statements that tell Go how to format and display values.

Syntax: %<flag><width>.<precision><verb>

EXAMPLE: %6.2f
  6 = width (minimum total width)
  2 = precision (decimal places)
  f = verb (format type)

═══════════════════════════════════════════════════════════════════════════════

GENERAL FORMATTING VERBS:

%v = Default value format (most appropriate for the type)
%# = Go syntax format (how the value would appear in Go code)
%T = Type of the value (returns the type name)
%% = Prints a literal % sign

═══════════════════════════════════════════════════════════════════════════════

INTEGER FORMATTING VERBS:

%b = Binary (base 2) representation
%d = Decimal (base 10) - most common for integers
%+d = Decimal with explicit + or - sign
%o = Octal (base 8) representation
%#o = Octal with 0 prefix
%x = Hexadecimal (base 16) lowercase (a-f)
%X = Hexadecimal (base 16) uppercase (A-F)
%#x = Hexadecimal with 0x prefix

WIDTH & PADDING (for integers):
%4d = Right-justify with width 4, pad with spaces
%04d = Right-justify with width 4, pad with zeros
%-4d = Left-justify with width 4

═══════════════════════════════════════════════════════════════════════════════

STRING FORMATTING VERBS:

%s = Plain string
%q = Quoted string (adds double quotes)
%8s = Right-justify with width 8
%-8s = Left-justify with width 8
%x = Hexadecimal representation of bytes
%X = Hexadecimal representation (uppercase)
%X with space = Hexadecimal with space separation

═══════════════════════════════════════════════════════════════════════════════

BOOLEAN FORMATTING VERBS:

%t = Boolean value (true or false)
%v = Default boolean format

═══════════════════════════════════════════════════════════════════════════════

FLOAT FORMATTING VERBS:

%e = Scientific notation (lowercase e): 9.18e+02
%E = Scientific notation (uppercase E): 9.18E+02
%f = Decimal notation: 918.00
%.2f = Decimal notation with 2 decimal places: 918.00
%6.2f = Width 6, 2 decimal places, right-justified
%g = Uses %e or %f format as needed (most compact)
%G = Uses %E or %f format as needed

PRECISION FOR FLOATS:
%.2f = 2 digits after decimal point
%.0f = No decimal places (rounds to nearest integer)
%g = Eliminates trailing zeros and uses exponent if necessary

═══════════════════════════════════════════════════════════════════════════════

USEFUL TIP: Number Readability in Source Code
In Go, you can use underscores in numeric literals for readability:
- num := 1_000_000 (one million)
- price := 19.99 (no underscore before/after decimal)
- binary := 0b1010_1010

═══════════════════════════════════════════════════════════════════════════════
*/

func main() {
	fmt.Println("════════════════════════════════════════════════════════════")
	fmt.Println("PART 1: GENERAL FORMATTING VERBS")
	fmt.Println("════════════════════════════════════════════════════════════\n")

	floatVal := 15.5
	strVal := "Hello World"

	fmt.Println("Float value (15.5):")
	fmt.Printf("%%v (default):           %v\n", floatVal)
	fmt.Printf("%%#v (Go syntax):        %#v\n", floatVal)
	fmt.Printf("%%T (type):              %T\n", floatVal)
	fmt.Printf("%%%% (literal %%):       %%\n")

	fmt.Println("\nString value (\"Hello World\"):")
	fmt.Printf("%%v (default):           %v\n", strVal)
	fmt.Printf("%%#v (Go syntax):        %#v\n", strVal)
	fmt.Printf("%%T (type):              %T\n\n", strVal)

	// ─────────────────────────────────────────────────────────────────────────

	fmt.Println("════════════════════════════════════════════════════════════")
	fmt.Println("PART 2: INTEGER FORMATTING VERBS")
	fmt.Println("════════════════════════════════════════════════════════════\n")

	intVal := 18

	fmt.Println("Integer value (18):")
	fmt.Printf("%%b (binary):            %b\n", intVal)
	fmt.Printf("%%d (decimal):           %d\n", intVal)
	fmt.Printf("%%+d (with sign):        %+d\n", intVal)
	fmt.Printf("%%o (octal):             %o\n", intVal)
	fmt.Printf("%%#o (octal with 0):     %#o\n", intVal)
	fmt.Printf("%%x (hex lowercase):     %x\n", intVal)
	fmt.Printf("%%X (hex uppercase):     %X\n", intVal)
	fmt.Printf("%%#x (hex with 0x):      %#x\n\n", intVal)

	// Demonstrate uppercase vs lowercase in hex
	intVal2 := 255 // Will show as 'ff' or 'FF'
	fmt.Println("Integer value (255) - demonstrates uppercase/lowercase:")
	fmt.Printf("%%x (lowercase):         %x\n", intVal2)
	fmt.Printf("%%X (uppercase):         %X\n\n", intVal2)

	// ─────────────────────────────────────────────────────────────────────────

	fmt.Println("════════════════════════════════════════════════════════════")
	fmt.Println("PART 3: INTEGER WIDTH & PADDING")
	fmt.Println("════════════════════════════════════════════════════════════\n")

	num := 42
	negNum := -42

	fmt.Printf("Value: %d\n\n", num)

	fmt.Println("Width with space padding (right-justified):")
	fmt.Printf("%%4d:  '%4d'\n", num)
	fmt.Printf("%%6d:  '%6d'\n", num)
	fmt.Printf("%%8d:  '%8d'\n\n", num)

	fmt.Println("Width with zero padding (right-justified):")
	fmt.Printf("%%04d: '%04d'\n", num)
	fmt.Printf("%%06d: '%06d'\n", num)
	fmt.Printf("%%08d: '%08d'\n\n", num)

	fmt.Println("Negative number with padding:")
	fmt.Printf("%%4d:  '%4d' (right-justified with space)\n", negNum)
	fmt.Printf("%%04d: '%04d' (padded with zeros)\n\n", negNum)

	fmt.Println("Left-justified with padding:")
	fmt.Printf("%%-4d: '%-4d'\n", num)
	fmt.Printf("%%-6d: '%-6d'\n\n", num)

	// ─────────────────────────────────────────────────────────────────────────

	fmt.Println("════════════════════════════════════════════════════════════")
	fmt.Println("PART 4: STRING FORMATTING VERBS")
	fmt.Println("════════════════════════════════════════════════════════════\n")

	text := "World"

	fmt.Println("String value (\"World\"):")
	fmt.Printf("%%s (plain):             %s\n", text)
	fmt.Printf("%%q (quoted):            %q\n", text)
	fmt.Printf("%%x (hexadecimal):       %x\n", text)
	fmt.Printf("%%X (hex uppercase):     %X\n\n", text)

	// ─────────────────────────────────────────────────────────────────────────

	fmt.Println("════════════════════════════════════════════════════════════")
	fmt.Println("PART 5: STRING WIDTH & JUSTIFICATION")
	fmt.Println("════════════════════════════════════════════════════════════\n")

	word := "Go"

	fmt.Println("String value (\"Go\"):")
	fmt.Println("Right-justified (with leading spaces):")
	fmt.Printf("%%8s:  '%8s'\n", word)
	fmt.Printf("%%10s: '%10s'\n\n", word)

	fmt.Println("Left-justified (with trailing spaces):")
	fmt.Printf("%%-8s:  '%- 8s'\n", word)
	fmt.Printf("%%-10s: '%-10s'\n\n", word)

	// ─────────────────────────────────────────────────────────────────────────

	fmt.Println("════════════════════════════════════════════════════════════")
	fmt.Println("PART 6: BOOLEAN FORMATTING VERBS")
	fmt.Println("════════════════════════════════════════════════════════════\n")

	isTrue := true
	isFalse := false

	fmt.Printf("%%t (true):              %t\n", isTrue)
	fmt.Printf("%%t (false):             %t\n", isFalse)
	fmt.Printf("%%v (default true):      %v\n", isTrue)
	fmt.Printf("%%v (default false):     %v\n\n", isFalse)

	// ─────────────────────────────────────────────────────────────────────────

	fmt.Println("════════════════════════════════════════════════════════════")
	fmt.Println("PART 7: FLOAT FORMATTING VERBS")
	fmt.Println("════════════════════════════════════════════════════════════\n")

	floatSmall := 9.18
	floatLarge := 918_000_000.0 // Using underscore for readability

	fmt.Println("Float value (9.18):")
	fmt.Printf("%%e (scientific):        %e\n", floatSmall)
	fmt.Printf("%%f (decimal):           %f\n", floatSmall)
	fmt.Printf("%%.2f (2 decimal places):%% .2f = %.2f\n", floatSmall)
	fmt.Printf("%%g (compact):           %g\n\n", floatSmall)

	fmt.Println("Float value (918,000,000):")
	fmt.Printf("%%e (scientific):        %e\n", floatLarge)
	fmt.Printf("%%f (decimal):           %f\n", floatLarge)
	fmt.Printf("%%.2f (2 decimal places): %.2f\n", floatLarge)
	fmt.Printf("%%g (compact):           %g\n\n", floatLarge)

	// ─────────────────────────────────────────────────────────────────────────

	fmt.Println("════════════════════════════════════════════════════════════")
	fmt.Println("PART 8: FLOAT WIDTH & PRECISION")
	fmt.Println("════════════════════════════════════════════════════════════\n")

	price := 19.5

	fmt.Println("Float value (19.5):")
	fmt.Printf("%%f:      %f\n", price)
	fmt.Printf("%%.0f:    %.0f (no decimals, rounded)\n", price)
	fmt.Printf("%%.1f:    %.1f (1 decimal)\n", price)
	fmt.Printf("%%.2f:    %.2f (2 decimals)\n", price)
	fmt.Printf("%%.3f:    %.3f (3 decimals)\n\n", price)

	fmt.Println("Width and Precision combined:")
	fmt.Printf("%%6.2f:   '%6.2f' (width 6, 2 decimals, right-justified)\n", price)
	fmt.Printf("%%8.2f:   '%8.2f' (width 8, 2 decimals)\n", price)
	fmt.Printf("%%-8.2f:  '%-8.2f' (width 8, 2 decimals, left-justified)\n\n", price)

	// More complex example
	smallPrice := 9.1
	fmt.Println("Small float (9.1) with width 6 and 2 decimals:")
	fmt.Printf("%%6.2f: '%6.2f' (has leading spaces)\n\n", smallPrice)

	// ─────────────────────────────────────────────────────════════────────────

	fmt.Println("════════════════════════════════════════════════════════════")
	fmt.Println("PART 9: SCIENTIFIC VS COMPACT NOTATION")
	fmt.Println("════════════════════════════════════════════════════════════\n")

	values := []float64{0.00000123, 123456789.0, 0.000456, 100.0}

	for _, v := range values {
		fmt.Printf("Value: %g\n", v)
		fmt.Printf("  %%e (scientific):     %e\n", v)
		fmt.Printf("  %%f (decimal):        %f\n", v)
		fmt.Printf("  %%g (compact):        %g\n", v)
		fmt.Println()
	}

	// ─────────────────────────────────────────────────────────────────────────

	fmt.Println("════════════════════════════════════════════════════════════")
	fmt.Println("PART 10: PRACTICAL EXAMPLES")
	fmt.Println("════════════════════════════════════════════════════════════\n")

	// Example 1: Table formatting
	fmt.Println("Example 1: Formatted Table")
	fmt.Printf("%-10s | %8s | %8s\n", "Product", "Price", "Tax")
	fmt.Println("-----------|----------|----------")
	fmt.Printf("%-10s | %8.2f | %8.2f\n", "Laptop", 999.99, 79.99)
	fmt.Printf("%-10s | %8.2f | %8.2f\n", "Mouse", 29.50, 2.35)
	fmt.Printf("%-10s | %8.2f | %8.2f\n\n", "Keyboard", 149.00, 11.92)

	// Example 2: Data display
	fmt.Println("Example 2: Data Display with Multiple Types")
	age := 25
	name := "Alice"
	score := 95.5
	active := true

	fmt.Printf("Name:      %s\n", name)
	fmt.Printf("Age:       %d years\n", age)
	fmt.Printf("Score:     %.1f%%\n", score)
	fmt.Printf("Active:    %t\n\n", active)

	// Example 3: Memory addresses and hex values
	fmt.Println("Example 3: Hexadecimal and Binary Display")
	value := 255
	fmt.Printf("Decimal:    %d\n", value)
	fmt.Printf("Hex:        0x%X\n", value)
	fmt.Printf("Binary:     0b%b\n", value)
	fmt.Printf("Octal:      0%o\n\n", value)

	// ─────────────────────────────────────────────────────────────────────────

	fmt.Println("════════════════════════════════════════════════════════════")
	fmt.Println("SUMMARY & BEST PRACTICES")
	fmt.Println("════════════════════════════════════════════════════════════\n")

	fmt.Println("DETAILED FORMATTING VERBS TABLE:")
	fmt.Println("┌────────────┬─────────────────────────┬──────────────┬─────────────┐")
	fmt.Println("│ Verb       │ Description             │ Example      │ Output      │")
	fmt.Println("├────────────┼─────────────────────────┼──────────────┼─────────────┤")
	fmt.Println("│ %v         │ Default Value           │ 15           │ 15          │")
	fmt.Println("│ %T         │ Type                    │ 15           │ int         │")
	fmt.Println("│ %d         │ Decimal (Base 10)       │ 15           │ 15          │")
	fmt.Println("│ %b         │ Binary (Base 2)         │ 15           │ 1111        │")
	fmt.Println("│ %x         │ Hex (Base 16)           │ 15           │ f           │")
	fmt.Println("│ %f         │ Float (Default)         │ 15.5         │ 15.500000   │")
	fmt.Println("│ %.2f       │ Float (Precision)       │ 15.5         │ 15.50       │")
	fmt.Println("│ %q         │ Quoted String           │ \"Hi\"        │ \"Hi\"      │")
	fmt.Println("│ %t         │ Boolean                 │ true         │ true        │")
	fmt.Println("└────────────┴─────────────────────────┴──────────────┴─────────────┘\n")

	fmt.Println(`
QUICK REFERENCE TABLE:

TYPE       | COMMON VERBS      | EXAMPLES
-----------|-------------------|------------------------------------------
Integer    | %d, %x, %o, %b    | 42, 2a, 52, 101010
Float      | %f, %e, %g        | 3.14, 3.14e+00, 3.14
String     | %s, %q            | hello, "hello"
Boolean    | %t                | true, false
General    | %v, %T            | <type-specific>, int, string, etc.

WIDTH & PRECISION SYNTAX:

%5d       = Width 5, right-justified (spaces)
%05d      = Width 5, zero-padded, right-justified
%-5d      = Width 5, left-justified
%5.2f     = Width 5, 2 decimal places
%.2f      = 2 decimal places, no minimum width
%8s       = Width 8, right-justified string
%-8s      = Width 8, left-justified string

BEST PRACTICES:

✓ Use %d for integers (most common and clear)
✓ Use %.2f for currency and money values
✓ Use %s for strings (simplest format)
✓ Use %q for strings containing special characters
✓ Use %T to debug and check variable types
✓ Use %v when unsure (Go picks appropriate format)
✓ Use %x or %X for hexadecimal display
✓ Use %b for binary display (helpful for bit operations)
✓ Use width specifiers for aligned table output
✓ Use underscores in numeric literals for readability (1_000_000)

COMMON MISTAKES TO AVOID:

✗ Forgetting \n in Printf (use Println if you want newline)
✗ Using wrong verb for type (e.g., %d for string)
✗ Not specifying precision for floats (get many decimals)
✗ Misunderstanding width vs precision (6.2f = 6 total, 2 decimal)
✗ Using %f for very large/small numbers (use %e or %g instead)

FLOAT NOTATION GUIDE:

- %f:  Always shows decimal point (918.000000)
- %e:  Scientific notation (9.18e+02)
- %g:  Compact format (uses %e for large/small, %f for normal)
     Also removes trailing zeros and decimal point if not needed

PRACTICAL USE CASES:

1. CURRENCY:     fmt.Printf("Price: $%.2f\n", price)
2. ALIGNMENT:    fmt.Printf("%-10s %10.2f\n", name, value)
3. DEBUGGING:    fmt.Printf("Type: %T, Value: %v\n", var, var)
4. HEX OUTPUT:   fmt.Printf("Color: 0x%06X\n", colorCode)
5. PADDING ZEROS: fmt.Printf("ID: %05d\n", id)
	`)
}
