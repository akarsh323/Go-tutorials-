package intermediate

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"unicode/utf8"
)

// ============================================================================
// Topic 70: STRING FUNCTIONS AND MANIPULATION
// ============================================================================
//
// CORE CONCEPT: In Go, strings are READ-ONLY SLICES OF BYTES.
// This means:
//   • You cannot modify a string in place
//   • Every operation creates a NEW string
//   • Length counts BYTES, not characters
//
// The strings package provides efficient operations for:
//   1. Searching (Contains, Index, HasPrefix, HasSuffix)
//   2. Modifying (ToUpper, ToLower, Replace, Trim)
//   3. Splitting/Joining (Split, Join, Fields)
//   4. Building (strings.Builder for efficiency)
//
// ============================================================================

func main() {
	fmt.Println("=============== 70 String Functions and Manipulation ===============")

	// ========================================================================
	// SECTION 1: Strings Are Immutable (Read-Only Bytes)
	// ========================================================================
	fmt.Println("\n--- SECTION 1: Understanding Strings (Immutability) ---")
	demonstrateImmutability()

	// ========================================================================
	// SECTION 2: Basic Operations (Length, Concatenation, Indexing)
	// ========================================================================
	fmt.Println("\n--- SECTION 2: Basic Operations ---")
	demonstrateBasicOperations()

	// ========================================================================
	// SECTION 3: Searching and Checking Strings
	// ========================================================================
	fmt.Println("\n--- SECTION 3: Searching and Checking ---")
	demonstrateSearching()

	// ========================================================================
	// SECTION 4: Transforming Strings (Case, Trim, Replace)
	// ========================================================================
	fmt.Println("\n--- SECTION 4: Transforming Strings ---")
	demonstrateTransformation()

	// ========================================================================
	// SECTION 5: Splitting and Joining
	// ========================================================================
	fmt.Println("\n--- SECTION 5: Splitting and Joining ---")
	demonstrateSplitJoin()

	// ========================================================================
	// SECTION 6: Type Conversion (strings ↔ numbers)
	// ========================================================================
	fmt.Println("\n--- SECTION 6: Type Conversion ---")
	demonstrateTypeConversion()

	// ========================================================================
	// SECTION 7: Performance - strings.Builder vs Concatenation
	// ========================================================================
	fmt.Println("\n--- SECTION 7: Performance with strings.Builder ---")
	demonstrateBuilder()

	// ========================================================================
	// SECTION 8: Unicode and Runes (Non-English Characters)
	// ========================================================================
	fmt.Println("\n--- SECTION 8: Unicode and Runes ---")
	demonstrateUnicode()

	// ========================================================================
	// SECTION 9: Pattern Matching with Regular Expressions
	// ========================================================================
	fmt.Println("\n--- SECTION 9: Regular Expressions ---")
	demonstrateRegex()

	// ========================================================================
	// SECTION 10: Best Practices Summary
	// ========================================================================
	fmt.Println("\n--- SECTION 10: Best Practices Summary ---")
	fmt.Println(`
✓ DO:
  • Use strings.Builder for loops or multiple concatenations
  • Use strings.TrimSpace() instead of manual trimming
  • Use strings.Fields() for splitting on whitespace
  • Use utf8.RuneCountInString() for accurate character count
  • Use regexp for complex pattern matching
  • Use strconv.Itoa() to convert int to string
  • Remember: ALL string operations return NEW strings

✗ DON'T:
  • Use + operator in loops (very inefficient)
  • Use string(number) to convert numbers to strings
  • Assume len() gives character count (it's bytes!)
  • Forget that strings are immutable
  • Use regex for simple Contains/Index checks
  • Assume ASCII for international text

RULE OF THUMB:
  • Simple operations: Use strings package functions
  • Building strings: Use strings.Builder
  • Complex patterns: Use regexp
  • Non-English text: Use utf8/unicode packages
`)
}

// ============================================================================
// SECTION 1: Immutability - Strings Cannot Change
// ============================================================================

func demonstrateImmutability() {
	str := "Hello"
	fmt.Println(`
Key Fact: Strings are IMMUTABLE (read-only)

This means:
  • You cannot change individual characters
  • Every modification creates a NEW string
  • Original is always left unchanged

Why? Performance and safety:
  • Multiple parts of code can safely share the same string
  • No need for copying/locking
  • Memory efficient

What happens when you "modify" a string:
  str := "Hello"
  str = "World"  ← Not modifying, creating new string!

Example: Concatenation
`)
	fmt.Printf("Original: %s (address: %p)\n", str, &str)

	str2 := str + " World"
	fmt.Printf("After concatenation: %s (address: %p)\n", str2, &str2)

	fmt.Println("Notice: Original string is unchanged!")
	fmt.Println("  str still = \"Hello\"")
	fmt.Println("  str2 is a NEW string = \"Hello World\"")
}

// ============================================================================
// SECTION 2: Basic Operations
// ============================================================================

func demonstrateBasicOperations() {
	s := "Go is awesome"

	fmt.Println(`
OPERATION 1: Length (counts BYTES, not characters)
`)
	fmt.Printf("  String: %q\n", s)
	fmt.Printf("  Length: %d bytes\n", len(s))
	fmt.Println("  Note: For English text, bytes = characters")
	fmt.Println("  But for Unicode (emojis, etc.), bytes > characters")

	fmt.Println(`
OPERATION 2: Indexing (access individual byte)
`)
	fmt.Printf("  s[0] = %v (%c) — ASCII value 71 = 'G'\n", s[0], s[0])
	fmt.Printf("  s[1] = %v (%c) — ASCII value 111 = 'o'\n", s[1], s[1])
	fmt.Println("  Note: Indexing returns the BYTE value (integer)")

	fmt.Println(`
OPERATION 3: Slicing [start:end] (end is exclusive)
`)
	fmt.Printf("  String: %q (indices: 0=G, 1=o, 2=space, 3=i, 4=s)\n", s)
	fmt.Printf("  s[0:2] = %q (from 0 up to but not including 2)\n", s[0:2])
	fmt.Printf("  s[3:5] = %q (from 3 up to but not including 5)\n", s[3:5])
	fmt.Printf("  s[6:]  = %q (from 6 to end)\n", s[6:])

	fmt.Println(`
OPERATION 4: Concatenation (+ operator)
`)
	result := "Hello" + " " + "World"
	fmt.Printf("  \"Hello\" + \" \" + \"World\" = %q\n", result)
	fmt.Println("  WARNING: Each + creates a new string in memory!")
	fmt.Println("  In loops, this is very inefficient (use strings.Builder instead)")
}

// ============================================================================
// SECTION 3: Searching and Checking
// ============================================================================

func demonstrateSearching() {
	s := "Go is awesome"

	fmt.Println("Function: strings.Contains()")
	fmt.Printf("  Contains 'is': %v\n", strings.Contains(s, "is"))
	fmt.Printf("  Contains 'Go': %v\n", strings.Contains(s, "Go"))
	fmt.Printf("  Contains 'Python': %v\n", strings.Contains(s, "Python"))

	fmt.Println("\nFunction: strings.Index() — Find position")
	fmt.Printf("  Index of 'is': %d (position in string)\n", strings.Index(s, "is"))
	fmt.Printf("  Index of 'awesome': %d\n", strings.Index(s, "awesome"))
	fmt.Printf("  Index of 'notfound': %d (-1 means not found)\n", strings.Index(s, "notfound"))

	fmt.Println("\nFunction: strings.Count() — Count occurrences")
	text := "Go Go Go"
	fmt.Printf("  String: %q\n", text)
	fmt.Printf("  Count('Go'): %d\n", strings.Count(text, "Go"))

	fmt.Println("\nFunction: strings.HasPrefix / HasSuffix()")
	fmt.Printf("  HasPrefix('Go'): %v\n", strings.HasPrefix(s, "Go"))
	fmt.Printf("  HasSuffix('awesome'): %v\n", strings.HasSuffix(s, "awesome"))
	fmt.Printf("  HasSuffix('.pdf'): %v\n", strings.HasSuffix(s, ".pdf"))
}

// ============================================================================
// SECTION 4: Transforming Strings
// ============================================================================

func demonstrateTransformation() {
	s := "  Go is Awesome  "

	fmt.Println("Function: strings.TrimSpace() — Remove leading/trailing whitespace")
	fmt.Printf("  Before: %q (with spaces)\n", s)
	fmt.Printf("  After:  %q\n", strings.TrimSpace(s))

	fmt.Println("\nFunction: strings.ToUpper / ToLower()")
	fmt.Printf("  Original: %q\n", strings.TrimSpace(s))
	fmt.Printf("  ToUpper:  %q\n", strings.ToUpper(strings.TrimSpace(s)))
	fmt.Printf("  ToLower:  %q\n", strings.ToLower(strings.TrimSpace(s)))

	fmt.Println("\nFunction: strings.Replace(s, old, new, count)")
	clean := strings.TrimSpace(s)
	fmt.Printf("  Original: %q\n", clean)
	fmt.Printf("  Replace('is', 'was', 1): %q (only 1st occurrence)\n",
		strings.Replace(clean, "is", "was", 1))
	fmt.Printf("  ReplaceAll('is', 'was'): %q (all occurrences)\n",
		strings.ReplaceAll(clean, "is", "was"))

	fmt.Println("\nFunction: strings.Repeat(s, count)")
	fmt.Printf("  Repeat('ab', 3): %q\n", strings.Repeat("ab", 3))
}

// ============================================================================
// SECTION 5: Splitting and Joining
// ============================================================================

func demonstrateSplitJoin() {
	fmt.Println("Function: strings.Split(s, sep) — Break apart by separator")
	data := "Apple,Orange,Banana"
	fmt.Printf("  String: %q\n", data)
	parts := strings.Split(data, ",")
	fmt.Printf("  Split by ',': %v\n", parts)

	fmt.Println("\nFunction: strings.Fields(s) — Split on whitespace")
	text := "Go   is   awesome"
	fmt.Printf("  String: %q\n", text)
	fmt.Printf("  Fields: %v\n", strings.Fields(text))
	fmt.Println("  Note: Automatically handles multiple spaces")

	fmt.Println("\nFunction: strings.Join(slice, sep) — Put together")
	countries := []string{"Germany", "France", "Italy"}
	fmt.Printf("  Array: %v\n", countries)
	fmt.Printf("  Join with ', ': %q\n", strings.Join(countries, ", "))
	fmt.Printf("  Join with ' | ': %q\n", strings.Join(countries, " | "))

	fmt.Println("\nPattern: CSV parsing")
	csv := "name,age,city\nAlice,30,NYC\nBob,25,LA"
	lines := strings.Split(csv, "\n")
	for i, line := range lines {
		fields := strings.Split(line, ",")
		fmt.Printf("  Line %d: %v\n", i+1, fields)
	}
}

// ============================================================================
// SECTION 6: Type Conversion (strings ↔ numbers)
// ============================================================================

func demonstrateTypeConversion() {
	fmt.Println("Converting NUMBER to STRING:")
	fmt.Println("  WRONG: str := string(42)  ← This creates a non-printable character!")
	fmt.Println("  RIGHT: str := strconv.Itoa(42)")

	num := 42
	str := strconv.Itoa(num)
	fmt.Printf("  Result: %q (type: %T)\n", str, str)

	fmt.Println("\nConverting STRING to NUMBER:")
	fmt.Println("  Use strconv.Atoi(string) → (int, error)")

	numStr := "123"
	result, err := strconv.Atoi(numStr)
	if err == nil {
		fmt.Printf("  Result: %v (type: %T)\n", result, result)
	}

	fmt.Println("\nOther conversions:")
	fmt.Println("  strconv.FormatInt(num, 10)  ← Convert int64 to string (base 10)")
	fmt.Println("  strconv.ParseInt(str, 10, 64) ← Convert string to int64")
	fmt.Println("  strconv.FormatFloat(f, 'f', 2, 64) ← Format float with precision")
}

// ============================================================================
// SECTION 7: Performance - strings.Builder
// ============================================================================

func demonstrateBuilder() {
	fmt.Println("Problem: Concatenation in loops is inefficient")
	fmt.Println("  result := \"\"")
	fmt.Println("  for i := 0; i < 10000; i++ {")
	fmt.Println("    result += \"x\"  ← Creates 10000 new strings!")
	fmt.Println("  }")
	fmt.Println()
	fmt.Println("Solution: Use strings.Builder")
	fmt.Println()

	var builder strings.Builder

	fmt.Println("Method 1: WriteString()")
	builder.WriteString("Hello")
	builder.WriteString(" ")
	builder.WriteString("World")
	fmt.Printf("  Result: %q\n", builder.String())

	fmt.Println("\nMethod 2: WriteRune() — Add individual character")
	builder.Reset() // Clear the builder
	builder.WriteRune('H')
	builder.WriteRune('i')
	builder.WriteRune('!')
	fmt.Printf("  Result: %q\n", builder.String())

	fmt.Println("\nMethod 3: Building efficiently in a loop")
	builder.Reset()
	for i := 1; i <= 5; i++ {
		builder.WriteString(fmt.Sprintf("Item%d ", i))
	}
	fmt.Printf("  Result: %q\n", builder.String())

	fmt.Println("\nWhy is Builder efficient?")
	fmt.Println("  • Single allocation of memory")
	fmt.Println("  • Appends to that allocation")
	fmt.Println("  • Only creates final string when you call String()")
	fmt.Println("  • vs + operator: creates new string EVERY time")
}

// ============================================================================
// SECTION 8: Unicode and Runes
// ============================================================================

func demonstrateUnicode() {
	fmt.Println("Challenge: Non-English text")
	fmt.Println()

	str := "Hello 😊"
	fmt.Printf("String: %q\n", str)
	fmt.Printf("len(str): %d (counting BYTES)\n", len(str))
	fmt.Printf("utf8.RuneCountInString(str): %d (counting CHARACTERS)\n", utf8.RuneCountInString(str))

	fmt.Println("\nWhy the difference?")
	fmt.Println("  • 'H', 'e', 'l', 'l', 'o', ' ' = 6 bytes (ASCII)")
	fmt.Println("  • '😊' emoji = 4 bytes (UTF-8 encoded)")
	fmt.Println("  • Total: 10 bytes")
	fmt.Println("  • But visually: 7 characters")

	fmt.Println("\nCorrect way to iterate over runes:")
	str2 := "Hello"
	fmt.Printf("String: %q\n", str2)
	fmt.Print("Runes: ")
	for i, r := range str2 {
		fmt.Printf("[%d]=%q ", i, r)
	}
	fmt.Println()

	fmt.Println("\nUtility functions:")
	foreign := "こんにちは" // Japanese: "Hello"
	fmt.Printf("String: %q\n", foreign)
	fmt.Printf("len(): %d bytes\n", len(foreign))
	fmt.Printf("RuneCountInString(): %d characters\n", utf8.RuneCountInString(foreign))
}

// ============================================================================
// SECTION 9: Regular Expressions
// ============================================================================

func demonstrateRegex() {
	fmt.Println("Use regexp for COMPLEX pattern matching")
	fmt.Println()

	fmt.Println("Pattern: \\d+ (one or more digits)")
	text := "I have 2 apples and 15 oranges"
	fmt.Printf("Text: %q\n", text)

	re := regexp.MustCompile(`\d+`)
	matches := re.FindAllString(text, -1)
	fmt.Printf("Numbers found: %v\n", matches)

	fmt.Println("\nPattern: Email validation (simplified)")
	emailPattern := `^[a-zA-Z0-9]+@[a-zA-Z0-9]+\.[a-zA-Z]{2,}$`
	re2 := regexp.MustCompile(emailPattern)

	emails := []string{"user@example.com", "invalid.email", "test@domain.co.uk"}
	for _, email := range emails {
		fmt.Printf("  %q is valid: %v\n", email, re2.MatchString(email))
	}

	fmt.Println("\nPattern: Extract parts")
	urlPattern := `https://([a-zA-Z0-9.]+)/([a-zA-Z0-9]+)`
	re3 := regexp.MustCompile(urlPattern)
	url := "https://github.com/golang"
	parts := re3.FindStringSubmatch(url)
	fmt.Printf("URL: %q\n", url)
	fmt.Printf("Parts: %v\n", parts)
	if len(parts) > 1 {
		fmt.Printf("  Domain: %q\n", parts[1])
		fmt.Printf("  Path: %q\n", parts[2])
	}
}

// ============================================================================
// COMPREHENSIVE REFERENCE TABLE
// ============================================================================
//
// SEARCH & CHECK:
//   Contains(s, substr) bool
//   Index(s, substr) int
//   Count(s, substr) int
//   HasPrefix(s, prefix) bool
//   HasSuffix(s, suffix) bool
//
// MODIFY:
//   ToUpper(s) string
//   ToLower(s) string
//   TrimSpace(s) string
//   Trim(s, cutset) string
//   Replace(s, old, new, n) string
//   ReplaceAll(s, old, new) string
//   Repeat(s, count) string
//
// SPLIT & JOIN:
//   Split(s, sep) []string
//   Fields(s) []string
//   Join(a, sep) string
//
// BUILD:
//   strings.Builder (most efficient for multiple operations)
//
// CONVERT:
//   strconv.Itoa(i) string
//   strconv.Atoi(s) (int, error)
//
// UNICODE:
//   utf8.RuneCountInString(s) int
//   range s ← iterates over runes
//
// PATTERN:
//   regexp.MustCompile(pattern) *regexp.Regexp
//   re.FindAllString(s, n) []string
//   re.MatchString(s) bool
//
// ============================================================================
