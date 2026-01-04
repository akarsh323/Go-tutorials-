package main

import "fmt"

// 98 Type Conversions
func Demo98TypeConversions() {
	fmt.Println("-- 98 Type Conversions --")
	var i int = 42
	var f float64 = float64(i)
	fmt.Println(i, f)
}




Based on the structure of the logging demo you provided and the content of the type conversion transcript, here is the complete code for **Demo Type Conversion**.

It includes the numeric conversions, string/byte manipulations, restrictions, and the philosophical advice mentioned in the lecture.

```go
package main

import (
	"fmt"
)

func main() {
	DemoTypeConversion()
}

// ============================================================================
// Type Conversion in Go
// ============================================================================
//
// Go requires EXPLICIT type conversion (casting). There is no implicit
// conversion (e.g., you cannot automatically treat an int as a float).
//
// Key Concepts:
// 1. Syntax: Type(Value) -> e.g., int(3.14)
// 2. Numeric: converting ints to floats and vice-versa.
// 3. Precision: converting float to int causes truncation (loss of data).
// 4. Strings/Bytes: converting strings to []byte and back (Unicode/ASCII).
// 5. Restrictions: fundamental types (like bool vs int) cannot convert.
//
// ============================================================================

// ============================================================================
// PART 1: NUMERIC CONVERSIONS
// ============================================================================

func Demo_Part1_NumericConversion() {
	fmt.Println("\n=== PART 1: NUMERIC CONVERSIONS ===")
	fmt.Println()

	fmt.Println("📌 Int to Float:")
	var a int32 = 10
	// Syntax: TargetType(Value)
	// Note: This is NOT a function call; it is a conversion expression built into the compiler.
	var b float64 = float64(a)

	fmt.Printf("   Original: %v (Type: %T)\n", a, a)
	fmt.Printf("   Converted: %.2f (Type: %T)\n", b, b)
	fmt.Println()

	fmt.Println("📌 Float to Int (Truncation):")
	var pi float64 = 3.14159
	// When converting float to int, Go drops the decimal completely.
	// It does NOT round; it truncates.
	var piInt int = int(pi)

	fmt.Printf("   Original Float: %v\n", pi)
	fmt.Printf("   Converted Int:  %v (Precision lost!)\n", piInt)
	fmt.Println()
}

// ============================================================================
// PART 2: RESTRICTIONS (THE "SAD PATH")
// ============================================================================

func Demo_Part2_Restrictions() {
	fmt.Println("\n=== PART 2: RESTRICTIONS & ERRORS ===")
	fmt.Println()

	fmt.Println("📌 Incompatible Types:")
	fmt.Println("   Go only allows conversion between 'related' types.")
	fmt.Println("   Fundamental types like 'bool' and 'int' are NOT related.")
	fmt.Println()

	// ---------------------------------------------------------
	// CODE THAT WOULD CAUSE A COMPILER ERROR (Squiggly Lines):
	// ---------------------------------------------------------
	// var myBool bool = true
	// var myInt int = int(myBool) // ERROR: cannot convert myBool (type bool) to type int
	//
	// var zero int = 0
	// var falseVal bool = bool(zero) // ERROR: cannot convert zero (type int) to type bool

	fmt.Println("   (Refer to code comments: You cannot convert int <-> bool in Go.)")
	fmt.Println()
}

// ============================================================================
// PART 3: STRINGS AND BYTE SLICES
// ============================================================================

func Demo_Part3_StringsAndBytes() {
	fmt.Println("\n=== PART 3: STRINGS AND BYTE SLICES ===")
	fmt.Println()

	fmt.Println("📌 String to []byte:")
	// Strings are essentially sequences of characters (Runes/Bytes).
	// We can convert them to a slice of bytes to see their numeric (ASCII/Unicode) values.
	greeting := "Hello"
	byteSlice := []byte(greeting)

	fmt.Printf("   String: \"%s\"\n", greeting)
	fmt.Printf("   Bytes:  %v\n", byteSlice) // [72 101 108 108 111]
	fmt.Println()

	fmt.Println("📌 Special Characters (Unicode):")
	// Go handles UTF-8 handles multi-byte characters naturally
	complexStr := "Hello @ 🇯🇵"
	complexBytes := []byte(complexStr)

	fmt.Printf("   String: \"%s\"\n", complexStr)
	fmt.Printf("   Bytes:  %v\n", complexBytes)
	fmt.Println()

	fmt.Println("📌 []byte to String:")
	// We can convert those numbers back into text.
	// 72 = 'H', 255 is outside standard ASCII but valid in byte slices
	rawBytes := []byte{72, 101, 108, 108, 111}
	restoredString := string(rawBytes)

	fmt.Printf("   Original Bytes: %v\n", rawBytes)
	fmt.Printf("   Restored String: \"%s\"\n", restoredString)
	fmt.Println()
}

// ============================================================================
// PART 4: DATA OVERFLOW (BYTE LIMITS)
// ============================================================================

func Demo_Part4_Overflow() {
	fmt.Println("\n=== PART 4: DATA OVERFLOW CONCEPTS ===")
	fmt.Println()

	fmt.Println("📌 Byte Limits (uint8):")
	fmt.Println("   A byte is an alias for uint8 (0 to 255).")
	fmt.Println("   If we try to put a number > 255 explicitly into a byte slice,")
	fmt.Println("   the compiler might catch it, or it might wrap around if calculated.")
	fmt.Println()

	// Example from transcript:
	// We cannot put 256 or 1000 into a byte slice directly using literals in a way that fits.
	// validBytes := []byte{255, 1000} // This would cause a compile error: "1000 overflows byte"

	validBytes := []byte{255, 72} // 255 is the max, 72 is 'H'
	fmt.Printf("   Valid Bytes: %v\n", validBytes)
	fmt.Printf("   As String:   %s (Note: 255 might render as a placeholder)\n", string(validBytes))
	fmt.Println()
}

// ============================================================================
// PART 5: MENTAL MODEL & PHILOSOPHY
// ============================================================================

func Demo_Part5_Philosophy() {
	fmt.Println("\n=== PART 5: LEARNING PHILOSOPHY ===")
	fmt.Println()

	fmt.Println("📌 The \"Analysis Paralysis\" Trap:")
	fmt.Println("   As we move to complex topics (IO, OS, File Systems), you might ask:")
	fmt.Println("   \"Why does it work this way?\" or \"What if I do X unlikely thing?\"")
	fmt.Println()

	fmt.Println("   > INSTRUCTOR ADVICE:")
	fmt.Println("   1. Focus on the WHAT first (Functionality).")
	fmt.Println("   2. Learn how to use the tool to build the application.")
	fmt.Println("   3. Don't wander into \"unnecessary possibilities\" immediately.")
	fmt.Println("   4. Understand the source code (control-click), but stay focused on the objective.")
	fmt.Println()
}

// ============================================================================
// MAIN DEMO FUNCTION
// ============================================================================

func DemoTypeConversion() {
	fmt.Println("-- Go Type Conversion Demo --")
	fmt.Println("Based on the transcript, this demo covers:")
	fmt.Println("1. Numeric Conversions (Int/Float)")
	fmt.Println("2. String & Byte Slice Conversions")
	fmt.Println("3. Incompatible Types")
	fmt.Println("4. Learning Mindset")

	Demo_Part1_NumericConversion()
	Demo_Part2_Restrictions()
	Demo_Part3_StringsAndBytes()
	Demo_Part4_Overflow()
	Demo_Part5_Philosophy()

	fmt.Println("\n=== END OF DEMO ===")
}

```