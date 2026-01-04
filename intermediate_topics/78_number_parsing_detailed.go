package intermediate

import (
	"fmt"
	"strconv"
)

// Topic 78: Number Parsing - Converting Strings to Numbers
// ==========================================================
// This lesson covers converting textual data (strings) into numeric types.
// We learn about Atoi, ParseInt, ParseFloat, different bases,
// and critical error handling patterns.

func main() {
	fmt.Println("=== Topic 78: Number Parsing - String to Number Conversion ===\n")

	lesson1BasicIntegerParsing()
	fmt.Println("\n" + string([]byte{61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 61}) + "\n")

	lesson2AdvancedIntegerParsing()
	fmt.Println("\n" + string([]byte{61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 61}) + "\n")

	lesson3FloatingPointParsing()
	fmt.Println("\n" + string([]byte{61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 61}) + "\n")

	lesson4ParsingDifferentBases()
	fmt.Println("\n" + string([]byte{61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 61}) + "\n")

	lesson5ErrorHandling()
	fmt.Println("\n" + string([]byte{61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 61}) + "\n")

	lesson6PracticalExercise()
}

// LESSON 1: Basic Integer Parsing - strconv.Atoi
// ===============================================
func lesson1BasicIntegerParsing() {
	fmt.Println("LESSON 1: BASIC INTEGER PARSING - strconv.Atoi")
	fmt.Println("----------------------------------------------\n")

	fmt.Println("THE CONCEPT:")
	fmt.Println("  Number parsing converts text strings into numeric types")
	fmt.Println("  that computers can perform arithmetic upon.\n")

	fmt.Println("THE SIMPLEST METHOD: strconv.Atoi")
	fmt.Println("  • Function: strconv.Atoi(string)")
	fmt.Println("  • Returns: (int, error)")
	fmt.Println("  • Atoi = ASCII to Integer\n")

	fmt.Println("WHEN TO USE:")
	fmt.Println("  ✓ Simple base-10 conversions")
	fmt.Println("  ✓ When you don't need custom base or bit size")
	fmt.Println("  ✓ Quick, straightforward conversions\n")

	fmt.Println("SYNTAX:")
	fmt.Println("  num, err := strconv.Atoi(\"12345\")\n")

	fmt.Println("PRACTICAL DEMONSTRATION:\n")

	// Example 1: Simple conversion
	numStr := "12345"
	fmt.Printf("Input string: \"%s\"\n", numStr)

	num, err := strconv.Atoi(numStr)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	fmt.Printf("Parsed integer: %d\n", num)
	fmt.Printf("Data type: int (can do math)\n\n")

	// Example 2: Performing arithmetic
	fmt.Println("We can now perform arithmetic:")
	fmt.Printf("  %d + 1 = %d\n", num, num+1)
	fmt.Printf("  %d * 2 = %d\n", num, num*2)
	fmt.Printf("  %d / 5 = %d\n", num, num/5)
}

// LESSON 2: Advanced Integer Parsing - strconv.ParseInt
// ======================================================
func lesson2AdvancedIntegerParsing() {
	fmt.Println("LESSON 2: ADVANCED INTEGER PARSING - strconv.ParseInt")
	fmt.Println("-----------------------------------------------------\n")

	fmt.Println("THE LIMITATION OF ATOI:")
	fmt.Println("  • Atoi only works with base-10 (decimal)")
	fmt.Println("  • Atoi always returns a standard int\n")

	fmt.Println("SOLUTION: strconv.ParseInt")
	fmt.Println("  • Function: strconv.ParseInt(string, base, bitSize)")
	fmt.Println("  • Returns: (int64, error)\n")

	fmt.Println("THE THREE PARAMETERS:\n")

	fmt.Println("1. String: The text to parse")
	fmt.Println("   Example: \"12345\", \"FF\", \"1010\"\n")

	fmt.Println("2. Base: The numbering system")
	fmt.Println("   • 10 = Decimal (0-9)")
	fmt.Println("   • 2  = Binary (0-1)")
	fmt.Println("   • 16 = Hexadecimal (0-9, A-F)\n")

	fmt.Println("3. BitSize: Desired integer precision")
	fmt.Println("   • 8  = int8 (-128 to 127)")
	fmt.Println("   • 16 = int16 (-32768 to 32767)")
	fmt.Println("   • 32 = int32")
	fmt.Println("   • 64 = int64 (largest, used most often)\n")

	fmt.Println("IMPORTANT: ParseInt ALWAYS returns int64")
	fmt.Println("  The bitSize parameter controls range checking, not the return type.\n")

	fmt.Println("EXAMPLE 1: Standard Base-10 Parsing with int64")
	numStr := "12345"
	fmt.Printf("  Input: \"%s\" (base 10, 64-bit)\n", numStr)

	num64, err := strconv.ParseInt(numStr, 10, 64)
	if err != nil {
		fmt.Printf("  Error: %v\n", err)
		return
	}

	fmt.Printf("  Result: %d (type: int64)\n", num64)
	fmt.Println()

	fmt.Println("EXAMPLE 2: Parsing into a smaller bit size (32-bit)")
	num32, err := strconv.ParseInt(numStr, 10, 32)
	if err != nil {
		fmt.Printf("  Error: %v\n", err)
		return
	}

	fmt.Printf("  Result: %d (still int64, but values are bounded by 32-bit range)\n", num32)
	fmt.Println()

	fmt.Println("EXAMPLE 3: Casting int64 to int (if needed)")
	castedNum := int(num64)
	fmt.Printf("  int64(%d) cast to int: %d\n", num64, castedNum)
}

// LESSON 3: Floating Point Parsing - strconv.ParseFloat
// =====================================================
func lesson3FloatingPointParsing() {
	fmt.Println("LESSON 3: FLOATING POINT PARSING - strconv.ParseFloat")
	fmt.Println("----------------------------------------------------\n")

	fmt.Println("THE NEED FOR FLOATS:")
	fmt.Println("  Integers can't represent fractional values (3.14, 2.5, etc.)")
	fmt.Println("  For decimal numbers, we use ParseFloat.\n")

	fmt.Println("THE FUNCTION:")
	fmt.Println("  strconv.ParseFloat(string, bitSize)")
	fmt.Println("  Returns: (float64, error)\n")

	fmt.Println("BITSIZE PARAMETER:")
	fmt.Println("  • 32 = float32 (single precision, less accurate)")
	fmt.Println("  • 64 = float64 (double precision, more accurate)\n")

	fmt.Println("GENERAL RULE: Use 64 for better precision\n")

	fmt.Println("EXAMPLE 1: Parsing 3.14 with full precision")
	floatStr := "3.14159265359"
	fmt.Printf("  Input: \"%s\"\n", floatStr)

	floatVal, err := strconv.ParseFloat(floatStr, 64)
	if err != nil {
		fmt.Printf("  Error: %v\n", err)
		return
	}

	fmt.Printf("  Result: %.11f (full precision)\n", floatVal)
	fmt.Println()

	fmt.Println("EXAMPLE 2: Formatting output to limited decimal places")
	fmt.Printf("  %.2f (2 decimal places)\n", floatVal)
	fmt.Printf("  %.4f (4 decimal places)\n", floatVal)
	fmt.Println()

	fmt.Println("EXAMPLE 3: Parsing different float formats")
	examples := []string{"3.14", "0.5", "100.0", "-2.5", "1.23e-4"}

	for _, example := range examples {
		val, err := strconv.ParseFloat(example, 64)
		if err == nil {
			fmt.Printf("  %-15s → %f\n", example, val)
		}
	}
	fmt.Println()

	fmt.Println("NOTE: Scientific Notation is supported!")
	sci := "1.23e-4"
	sciVal, _ := strconv.ParseFloat(sci, 64)
	fmt.Printf("  \"%s\" = %.10f\n", sci, sciVal)
}

// LESSON 4: Parsing Different Bases
// =================================
func lesson4ParsingDifferentBases() {
	fmt.Println("LESSON 4: PARSING DIFFERENT BASES - BINARY & HEXADECIMAL")
	fmt.Println("------------------------------------------------------\n")

	fmt.Println("WHY DIFFERENT BASES?")
	fmt.Println("  • Binary (Base 2): Used in low-level programming, flags, bit operations")
	fmt.Println("  • Hexadecimal (Base 16): Used in colors (#FF00AA), memory addresses")
	fmt.Println("  • Decimal (Base 10): Human default (0-9)\n")

	fmt.Println("BASE CONVERSION TABLE:")
	fmt.Println("  Decimal  Binary    Hexadecimal")
	fmt.Println("  -------  ------    -----------")
	fmt.Println("  0        0         0")
	fmt.Println("  10       1010      A")
	fmt.Println("  15       1111      F")
	fmt.Println("  16       10000     10")
	fmt.Println("  255      11111111  FF\n")

	fmt.Println("EXAMPLE 1: Parsing Binary (Base 2)")
	fmt.Println("  Convert \"1010\" (binary) to decimal")
	binaryStr := "1010"
	fmt.Printf("  Input: \"%s\" (binary)\n", binaryStr)

	binaryVal, err := strconv.ParseInt(binaryStr, 2, 64)
	if err != nil {
		fmt.Printf("  Error: %v\n", err)
		return
	}

	fmt.Printf("  Result: %d (decimal)\n", binaryVal)
	fmt.Println("  Breakdown: 1×8 + 0×4 + 1×2 + 0×1 = 10\n")

	fmt.Println("EXAMPLE 2: Parsing Hexadecimal (Base 16)")
	fmt.Println("  Convert \"FF\" (hex) to decimal")
	hexStr := "FF"
	fmt.Printf("  Input: \"%s\" (hexadecimal)\n", hexStr)

	hexVal, err := strconv.ParseInt(hexStr, 16, 64)
	if err != nil {
		fmt.Printf("  Error: %v\n", err)
		return
	}

	fmt.Printf("  Result: %d (decimal)\n", hexVal)
	fmt.Println("  Breakdown: F×16 + F×1 = 15×16 + 15×1 = 255\n")

	fmt.Println("EXAMPLE 3: Real-world hex - Web Color Codes")
	fmt.Println("  Web colors use hex: #RRGGBB")
	fmt.Println("  Example: #FF00AA (magenta)")

	redHex := "FF"
	greenHex := "00"
	blueHex := "AA"

	red, _ := strconv.ParseInt(redHex, 16, 64)
	green, _ := strconv.ParseInt(greenHex, 16, 64)
	blue, _ := strconv.ParseInt(blueHex, 16, 64)

	fmt.Printf("  Red:   %s = %d\n", redHex, red)
	fmt.Printf("  Green: %s = %d\n", greenHex, green)
	fmt.Printf("  Blue:  %s = %d\n", blueHex, blue)
	fmt.Printf("  RGB value: (%d, %d, %d)\n\n", red, green, blue)

	fmt.Println("EXAMPLE 4: Batch conversion from different bases")
	conversions := []struct {
		str  string
		base int
		name string
	}{
		{"1010", 2, "Binary"},
		{"12345", 10, "Decimal"},
		{"1A", 16, "Hex"},
		{"77", 8, "Octal"},
	}

	for _, conv := range conversions {
		val, _ := strconv.ParseInt(conv.str, conv.base, 64)
		fmt.Printf("  %s %-8s → Decimal: %d\n", conv.name, conv.str, val)
	}
}

// LESSON 5: Error Handling - Critical for Robust Code
// ==================================================
func lesson5ErrorHandling() {
	fmt.Println("LESSON 5: ERROR HANDLING - CRITICAL FOR ROBUST CODE")
	fmt.Println("--------------------------------------------------\n")

	fmt.Println("WHY ERROR HANDLING MATTERS:")
	fmt.Println("  Parsing can fail if:")
	fmt.Println("    • String contains letters: \"456ABC\"")
	fmt.Println("    • String contains invalid characters: \"12.34.56\"")
	fmt.Println("    • Number is out of range for the specified bitSize")
	fmt.Println("    • String is empty\n")

	fmt.Println("WITHOUT ERROR HANDLING:")
	fmt.Println("  Program crashes or uses zero-values silently\n")

	fmt.Println("WITH ERROR HANDLING:")
	fmt.Println("  You catch the error and handle it gracefully\n")

	fmt.Println("PATTERN:")
	fmt.Println("  value, err := strconv.Atoi(str)")
	fmt.Println("  if err != nil {")
	fmt.Println("      // Handle the error (log, return, retry, etc.)")
	fmt.Println("  }\n")

	fmt.Println("EXAMPLE 1: Valid input (no error)")
	validStr := "12345"
	val, err := strconv.Atoi(validStr)
	if err != nil {
		fmt.Printf("  Error: %v\n", err)
	} else {
		fmt.Printf("  ✓ Successfully parsed: %d\n", val)
	}
	fmt.Println()

	fmt.Println("EXAMPLE 2: Invalid input (error case)")
	invalidStr := "456ABC"
	fmt.Printf("  Input: \"%s\"\n", invalidStr)

	_, err = strconv.Atoi(invalidStr)
	if err != nil {
		fmt.Printf("  ✗ Error caught: %v\n", err)
		fmt.Println("  (Program continues safely instead of crashing)")
	}
	fmt.Println()

	fmt.Println("EXAMPLE 3: Number out of range (bitSize error)")
	largeNum := "999999999999999999999999"
	fmt.Printf("  Input: \"%s\"\n", largeNum)
	fmt.Printf("  Parsing as 16-bit (max 32767)...\n")

	_, err = strconv.ParseInt(largeNum, 10, 16)
	if err != nil {
		fmt.Printf("  ✗ Error: %v\n", err)
		fmt.Println("  (Number is too large for 16-bit storage)")
	}
	fmt.Println()

	fmt.Println("EXAMPLE 4: Safe parsing with default value")
	userInput := "invalid"
	defaultValue := 100

	parsed, err := strconv.Atoi(userInput)
	if err != nil {
		parsed = defaultValue
		fmt.Printf("  Parsing failed, using default: %d\n", parsed)
	}
}

// LESSON 6: Practical Exercise - Mixed Transaction Parsing
// ========================================================
func lesson6PracticalExercise() {
	fmt.Println("LESSON 6: PRACTICAL EXERCISE - PARSING MIXED TRANSACTION DATA")
	fmt.Println("-----------------------------------------------------------\n")

	fmt.Println("SCENARIO:")
	fmt.Println("  You receive transaction data from various sources.")
	fmt.Println("  Some are valid, some are corrupted or invalid.")
	fmt.Println("  You need to parse and validate each one.\n")

	transactions := []string{
		"12345",   // Valid
		"500",     // Valid
		"invalid", // Invalid: letters
		"123.45",  // Invalid: not an integer
		"0",       // Valid: edge case
		"",        // Invalid: empty
		"-1000",   // Valid: negative
		"999ABC",  // Invalid: mixed
	}

	fmt.Println("TRANSACTIONS TO PARSE:\n")

	validCount := 0
	invalidCount := 0
	totalAmount := 0

	for i, txn := range transactions {
		fmt.Printf("Transaction %d: \"%s\"\n", i+1, txn)

		amount, err := strconv.Atoi(txn)
		if err != nil {
			fmt.Printf("  ✗ INVALID: %v\n", err)
			invalidCount++
		} else {
			fmt.Printf("  ✓ VALID: Amount = %d\n", amount)
			totalAmount += amount
			validCount++
		}
		fmt.Println()
	}

	fmt.Println("SUMMARY:")
	fmt.Printf("  Total transactions: %d\n", len(transactions))
	fmt.Printf("  Valid: %d\n", validCount)
	fmt.Printf("  Invalid: %d\n", invalidCount)
	fmt.Printf("  Total amount: %d\n", totalAmount)
}
