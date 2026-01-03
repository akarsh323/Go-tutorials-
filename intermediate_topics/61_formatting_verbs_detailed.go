package main

import "fmt"

/*
Topic 61: FORMATTING VERBS

CONCEPT: Special sequences in Printf that tell Go how to format values.

SYNTAX: %<flags><width>.<precision><verb>

EXAMPLES:
%v = default format
%d = decimal (integer)
%x = hexadecimal
%f = float
%.2f = float with 2 decimal places
%5d = integer with width 5
%s = string

GENERAL VERBS:
%v = default format, %T = type, %% = literal %

INTEGERS:
%d = decimal, %b = binary, %o = octal, %x = hex (lower), %X = hex (upper)

FLOATS:
%f = decimal, %e = scientific, %g = compact

WIDTH & PRECISION:
%5d = width 5 (right-justified)
%05d = width 5 (zero-padded)
%-5d = width 5 (left-justified)
%.2f = 2 decimal places
%6.2f = width 6, 2 decimal places
*/

func main() {
	fmt.Println("═══════════════════════════════════════════════════════════")
	fmt.Println("PART 1: GENERAL FORMATTING VERBS")
	fmt.Println("═══════════════════════════════════════════════════════════\n")

	fmt.Println("Value formatting verbs:")
	floatVal := 15.5
	fmt.Printf("%%v (default):   %v\n", floatVal)
	fmt.Printf("%%#v (Go syntax): %#v\n", floatVal)
	fmt.Printf("%%T (type):      %T\n", floatVal)
	fmt.Printf("%%%% (literal):  %%\n\n")

	// ─────────────────────────────────────────────────────────────────────────

	fmt.Println("═══════════════════════════════════════════════════════════")
	fmt.Println("PART 2: INTEGER FORMATTING")
	fmt.Println("═══════════════════════════════════════════════════════════\n")

	num := 18
	fmt.Printf("Number: %d\n\n", num)
	fmt.Printf("%%b (binary):     %b\n", num)
	fmt.Printf("%%d (decimal):    %d\n", num)
	fmt.Printf("%%o (octal):      %o\n", num)
	fmt.Printf("%%x (hex lower):  %x\n", num)
	fmt.Printf("%%X (hex upper):  %X\n", num)
	fmt.Printf("%%#x (hex 0x):    %#x\n\n", num)

	// ─────────────────────────────────────────────────────────────────────────

	fmt.Println("═══════════════════════════════════════════════════════════")
	fmt.Println("PART 3: WIDTH & PADDING (Integers)")
	fmt.Println("═══════════════════════════════════════════════════════════\n")

	n := 42
	fmt.Println("Right-justified with spaces:")
	fmt.Printf("%%4d:  '%4d'\n", n)
	fmt.Printf("%%6d:  '%6d'\n\n", n)

	fmt.Println("Right-justified with zeros:")
	fmt.Printf("%%04d: '%04d'\n", n)
	fmt.Printf("%%06d: '%06d'\n\n", n)

	fmt.Println("Left-justified:")
	fmt.Printf("%%-4d: '%-4d'\n", n)
	fmt.Printf("%%-6d: '%-6d'\n\n", n)

	// ─────────────────────────────────────────────────────────────────────────

	fmt.Println("═══════════════════════════════════════════════════════════")
	fmt.Println("PART 4: STRING FORMATTING")
	fmt.Println("═══════════════════════════════════════════════════════════\n")

	text := "World"
	fmt.Printf("String: \"%s\"\n\n", text)
	fmt.Printf("%%s (plain):   %s\n", text)
	fmt.Printf("%%q (quoted):  %q\n", text)
	fmt.Printf("%%x (hex):     %x\n\n", text)

	fmt.Println("String width & justification:")
	fmt.Printf("%%8s:  '%8s' (right)\n", text)
	fmt.Printf("%%-8s: '%-8s' (left)\n\n", text)

	// ─────────────────────────────────────────────────────────────────────────

	fmt.Println("═══════════════════════════════════════════════════════════")
	fmt.Println("PART 5: BOOLEAN & BASIC TYPES")
	fmt.Println("═══════════════════════════════════════════════════════════\n")

	fmt.Printf("%%t (true):  %t\n", true)
	fmt.Printf("%%t (false): %t\n", false)
	fmt.Printf("%%v (true):  %v\n", true)
	fmt.Printf("%%v (false): %v\n\n", false)

	// ─────────────────────────────────────────────────────────────────────────

	fmt.Println("═══════════════════════════════════════════════════════════")
	fmt.Println("PART 6: FLOAT FORMATTING")
	fmt.Println("═══════════════════════════════════════════════════════════\n")

	small := 9.18
	large := 918_000_000.0

	fmt.Printf("Small float (9.18):\n")
	fmt.Printf("%%e (scientific): %e\n", small)
	fmt.Printf("%%f (decimal):    %f\n", small)
	fmt.Printf("%%.2f (precision): %.2f\n", small)
	fmt.Printf("%%g (compact):    %g\n\n", small)

	fmt.Printf("Large float (918,000,000):\n")
	fmt.Printf("%%e (scientific): %e\n", large)
	fmt.Printf("%%f (decimal):    %f\n", large)
	fmt.Printf("%%.2f (precision): %.2f\n", large)
	fmt.Printf("%%g (compact):    %g\n\n", large)

	// ─────────────────────────────────────────────────────────────────────────

	fmt.Println("═══════════════════════════════════════════════════════════")
	fmt.Println("PART 7: WIDTH & PRECISION (Floats)")
	fmt.Println("═══════════════════════════════════════════════════════════\n")

	price := 19.5
	fmt.Printf("Price: %.2f\n\n", price)
	fmt.Printf("%%.0f: %.0f (no decimals)\n", price)
	fmt.Printf("%%.1f: %.1f (1 decimal)\n", price)
	fmt.Printf("%%.2f: %.2f (2 decimals)\n", price)
	fmt.Printf("%%.3f: %.3f (3 decimals)\n\n", price)

	fmt.Println("Width + Precision:")
	fmt.Printf("%%6.2f:  '%6.2f'\n", price)
	fmt.Printf("%%8.2f:  '%8.2f'\n", price)
	fmt.Printf("%%-8.2f: '%-8.2f'\n\n", price)

	// ─────────────────────────────────────────────────────────────────────────

	fmt.Println("═══════════════════════════════════════════════════════════")
	fmt.Println("PART 8: PRACTICAL EXAMPLES")
	fmt.Println("═══════════════════════════════════════════════════════════\n")

	fmt.Println("Example 1: Aligned Table")
	fmt.Printf("%-10s | %8s\n", "Product", "Price")
	fmt.Println("-----------|----------")
	fmt.Printf("%-10s | %8.2f\n", "Laptop", 999.99)
	fmt.Printf("%-10s | %8.2f\n", "Mouse", 29.50)
	fmt.Printf("%-10s | %8.2f\n\n", "Keyboard", 149.00)

	fmt.Println("Example 2: Various Displays")
	age := 25
	name := "Alice"
	score := 95.5

	fmt.Printf("Name:  %s\n", name)
	fmt.Printf("Age:   %d\n", age)
	fmt.Printf("Score: %.1f%%\n\n", score)

	fmt.Println("Example 3: Number Base Display")
	value := 255
	fmt.Printf("Decimal: %d\n", value)
	fmt.Printf("Hex:     0x%X\n", value)
	fmt.Printf("Binary:  0b%b\n", value)
	fmt.Printf("Octal:   0%o\n\n", value)

	// ─────────────────────────────────────────────────────────────────────────

	fmt.Println("═══════════════════════════════════════════════════════════")
	fmt.Println("FORMATTING VERBS REFERENCE TABLE")
	fmt.Println("═══════════════════════════════════════════════════════════\n")

	fmt.Println("┌────────┬──────────────────────┬──────────┬─────────┐")
	fmt.Println("│ Verb   │ Description          │ Example  │ Output  │")
	fmt.Println("├────────┼──────────────────────┼──────────┼─────────┤")
	fmt.Println("│ %v     │ Default value        │ 42       │ 42      │")
	fmt.Println("│ %d     │ Decimal integer      │ 42       │ 42      │")
	fmt.Println("│ %b     │ Binary               │ 5        │ 101     │")
	fmt.Println("│ %x     │ Hexadecimal (lower)  │ 255      │ ff      │")
	fmt.Println("│ %X     │ Hexadecimal (upper)  │ 255      │ FF      │")
	fmt.Println("│ %o     │ Octal                │ 8        │ 10      │")
	fmt.Println("│ %f     │ Float                │ 3.14     │ 3.14    │")
	fmt.Println("│ %e     │ Scientific notation  │ 3.14     │ 3e+00   │")
	fmt.Println("│ %s     │ String               │ \"hi\"    │ hi      │")
	fmt.Println("│ %q     │ Quoted string        │ \"hi\"    │ \"hi\"   │")
	fmt.Println("│ %t     │ Boolean              │ true     │ true    │")
	fmt.Println("│ %T     │ Type                 │ 42       │ int     │")
	fmt.Println("│ %%     │ Literal percent      │ %        │ %       │")
	fmt.Println("└────────┴──────────────────────┴──────────┴─────────┘\n")

	// ─────────────────────────────────────────────────────────────────────────

	fmt.Println("═══════════════════════════════════════════════════════════")
	fmt.Println("KEY CONCEPTS & BEST PRACTICES")
	fmt.Println("═══════════════════════════════════════════════════════════\n")

	fmt.Println(`
BASIC SYNTAX:
%d = decimal integer
%s = string
%f = float (default 6 decimals)
%.2f = 2 decimal places
%5d = width 5, right-justified
%05d = width 5, zero-padded
%-5d = width 5, left-justified
%T = type name
%v = default format

COMMON USES:
✓ Currency: fmt.Printf("Price: $%.2f\n", price)
✓ Alignment: fmt.Printf("%-10s %10.2f\n", name, value)
✓ Debug: fmt.Printf("Type: %T, Value: %v\n", var, var)
✓ Hex codes: fmt.Printf("0x%06X\n", color)
✓ Padding IDs: fmt.Printf("ID: %05d\n", id)

MISTAKES TO AVOID:
✗ Forgetting \n in Printf (newline)
✗ Wrong verb for type (%d for string)
✗ Not specifying float precision (gets 6 decimals)
✗ Width vs precision confusion (5.2f = width 5, 2 decimals)
✗ Using %f for huge/tiny numbers (use %e or %g)

FLOAT FORMAT GUIDE:
%f:  918.000000 (always shows decimal)
%e:  9.18e+02 (scientific notation)
%g:  918 (compact, removes trailing zeros)

WIDTH EXAMPLES:
%5d   = "   42" (spaces)
%05d  = "00042" (zeros)
%-5d  = "42   " (left, spaces)
%8.2f = "   19.50" (width 8, 2 decimals)

Pro Tip: Use underscores in numeric literals for readability
  big := 1_000_000  // easier to read than 1000000
	`)
}
