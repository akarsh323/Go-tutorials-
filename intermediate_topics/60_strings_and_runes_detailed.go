package main

import (
	"fmt"
	"unicode/utf8"
)

/*
Topic 60: STRINGS AND RUNES - A Comprehensive Guide

═══════════════════════════════════════════════════════════════════════════════

PART 1: STRINGS

DEFINITION:
A string is a sequence of bytes that often represent text.
Strings are IMMUTABLE - once created, their values cannot be changed.

KEY CHARACTERISTICS:
- Immutable: Cannot be modified after creation
- Sequences of bytes: Represent text using byte values
- Can be declared with double quotes or backticks
- Support escape sequences (when using double quotes)
- Can be concatenated, compared, and indexed
- Have a length that can be calculated

STRING DECLARATION & INITIALIZATION:

1. Using double quotes (standard strings):
   message := "Hello\nGo"
   - Supports escape sequences: \n, \t, \r, etc.

2. Using backticks (raw string literals):
   rawMessage := `Hello\nGo`
   - Does NOT interpret escape sequences
   - Treats everything literally as-is

ESCAPE SEQUENCES:
- \n = newline (moves to next line, first position)
- \t = tab (horizontal whitespace)
- \r = carriage return (cursor to first position, DOES NOT create new line)
- \\ = backslash
- \" = double quote

HISTORICAL NOTE:
Old typewriters had two operations:
1. Line feed (\n) - move to next line
2. Carriage return (\r) - move cursor to first position
Modern systems combine these, so \n alone is sufficient.

═══════════════════════════════════════════════════════════════════════════════

PART 2: RUNES

DEFINITION:
A rune is an alias for int32 that represents a Unicode code point.
Runes are used to represent individual characters in a string.

KEY CHARACTERISTICS:
- Type: int32 (4 bytes of memory)
- Represents Unicode code points
- Declared with single quotes: 'A', 'é', '😊'
- Support characters from multiple languages
- Can represent Emoji, Japanese, Chinese, etc.
- Essential for internationalization (i18n)

RUNES VS CHARACTERS (C's char):

SIMILARITIES:
- Both represent individual characters
- Both occupy fixed memory

DIFFERENCES:
- Rune: int32 (4 bytes) - supports Unicode
- Char (C): typically 1 byte - ASCII only
- Rune: built-in Unicode support
- Char (C): requires external libraries for Unicode

═══════════════════════════════════════════════════════════════════════════════

LEXICOGRAPHICAL COMPARISON (Dictionary Order):
When comparing strings, Go uses ASCII/Unicode values of characters.
- Uppercase letters have LOWER values than lowercase (A=65, a=97)
- Shorter string is considered smaller if it's a prefix
- Example: "app" < "apple" < "apple2"

═══════════════════════════════════════════════════════════════════════════════
*/

func main() {
	fmt.Println("════════════════════════════════════════════════════════════")
	fmt.Println("PART 1: STRINGS - DECLARATION AND BASICS")
	fmt.Println("════════════════════════════════════════════════════════════\n")

	// STRING DECLARATION
	message := "Hello\nGo"        // Standard string with escape sequences
	rawMessage := `Hello\nGo`     // Raw string literal (no escape sequences)
	messageWithTab := "Hello\tGo" // Tab character

	fmt.Println("Standard string (with escape sequences):")
	fmt.Println(message)
	fmt.Println("\nRaw string literal (no escape sequences):")
	fmt.Println(rawMessage)
	fmt.Println("\nString with tab:")
	fmt.Println(messageWithTab)

	// Carriage Return Example
	messageWithCR := "Hello\rWorld"
	fmt.Println("\nCarriage Return example (\\r only, no newline):")
	fmt.Println(messageWithCR)
	fmt.Println("(Notice: 'World' overwrites 'Hello' from the start)")

	// ─────────────────────────────────────────────────────────────────────────

	fmt.Println("\n════════════════════════════════════════════════════════════")
	fmt.Println("STRING LENGTH")
	fmt.Println("════════════════════════════════════════════════════════════\n")

	msg1 := "Hello\nGo"
	msg2 := "HelloGo"
	msg3 := `Hello\nGo`

	fmt.Printf("Length of \"%s\" = %d\n", "Hello\\nGo", len(msg1))
	fmt.Printf("Length of \"%s\" = %d\n", "HelloGo", len(msg2))
	fmt.Printf("Length of \"%s\" = %d (backticks treat \\n as 2 chars)\n", `Hello\nGo`, len(msg3))

	// ─────────────────────────────────────────────────────────────────────────

	fmt.Println("\n════════════════════════════════════════════════════════════")
	fmt.Println("STRING INDEXING (Returns Byte Values)")
	fmt.Println("════════════════════════════════════════════════════════════\n")

	str := "Hello"
	fmt.Printf("String: \"%s\"\n", str)
	fmt.Printf("First character at index 0: %d (ASCII value of 'H')\n", str[0])
	fmt.Printf("Second character at index 1: %d (ASCII value of 'e')\n", str[1])
	fmt.Println("\nNote: Indexing returns byte/ASCII values, not characters!")

	// ─────────────────────────────────────────────────────────────────────────

	fmt.Println("\n════════════════════════════════════════════════════════════")
	fmt.Println("STRING CONCATENATION")
	fmt.Println("════════════════════════════════════════════════════════════\n")

	greeting := "Hello"
	name := "Alice"

	// Direct concatenation (no space added automatically)
	result1 := greeting + name
	fmt.Printf("Direct concatenation: \"%s\"\n", result1)
	fmt.Println("(No space added automatically)")

	// Concatenation with space
	result2 := greeting + " " + name
	fmt.Printf("With manual space: \"%s\"\n", result2)

	// Using Print with commas (adds space automatically)
	fmt.Println("\nUsing Print with commas (adds space automatically):")
	fmt.Println(greeting, name)

	// ─────────────────────────────────────────────────────────────────────────

	fmt.Println("\n════════════════════════════════════════════════════════════")
	fmt.Println("STRING COMPARISON (Lexicographical Ordering)")
	fmt.Println("════════════════════════════════════════════════════════════\n")

	str1 := "Apple"
	str2 := "apple"
	str3 := "app"

	fmt.Printf("Comparing \"%s\" and \"%s\":\n", str1, str2)
	fmt.Printf("  \"%s\" < \"%s\": %v\n", str1, str2, str1 < str2)
	fmt.Println("  (Uppercase 'A' has lower ASCII value than lowercase 'a')")

	fmt.Printf("\nComparing \"%s\" and \"%s\":\n", str3, str1)
	fmt.Printf("  \"%s\" < \"%s\": %v\n", str3, str1, str3 < str1)
	fmt.Println("  (Shorter string is smaller if it's a prefix)")

	fmt.Printf("\nComparing \"%s\" and \"%s\":\n", "banana", "apple")
	fmt.Printf("  \"banana\" > \"apple\": %v\n", "banana" > "apple")
	fmt.Println("  ('b' > 'a' in ASCII values)")

	// ─────────────────────────────────────────────────────────────────────────

	fmt.Println("\n════════════════════════════════════════════════════════════")
	fmt.Println("STRING ITERATION")
	fmt.Println("════════════════════════════════════════════════════════════\n")

	text := "Hello"
	fmt.Printf("Iterating over string: \"%s\"\n\n", text)

	fmt.Println("Method 1: Using range (gets rune values):")
	for index, char := range text {
		fmt.Printf("  Index %d: character '%c' (rune value: %v)\n", index, char, char)
	}

	fmt.Println("\nMethod 2: Hexadecimal values:")
	for _, char := range text {
		fmt.Printf("  Character: %c, Hex: 0x%x\n", char, char)
	}

	// ─────────────────────────────────────────────────────────────────────────

	fmt.Println("\n════════════════════════════════════════════════════════════")
	fmt.Println("STRING IMMUTABILITY")
	fmt.Println("════════════════════════════════════════════════════════════\n")

	original := "Hello"
	fmt.Printf("Original string: \"%s\"\n", original)

	// To modify, create a new string
	modified := original + " World"
	fmt.Printf("Modified (new string): \"%s\"\n", modified)
	fmt.Printf("Original unchanged: \"%s\"\n", original)
	fmt.Println("\nNote: Strings are IMMUTABLE - operations create new strings!")

	// ─────────────────────────────────────────────────────────────────────────

	fmt.Println("\n════════════════════════════════════════════════════════════")
	fmt.Println("PART 2: RUNES - INDIVIDUAL CHARACTERS")
	fmt.Println("════════════════════════════════════════════════════════════\n")

	// RUNE DECLARATION
	var ch rune = 'A'
	var japanese rune = '日' // Japanese character for "day"
	var emoji rune = '😊'    // Smiley emoji

	fmt.Println("Rune values (as integers):")
	fmt.Printf("  'A' = %d\n", ch)
	fmt.Printf("  '日' (Japanese 'day') = %d\n", japanese)
	fmt.Printf("  '😊' (emoji) = %d\n", emoji)

	fmt.Println("\nRune values (as characters using %c):")
	fmt.Printf("  Rune ch: %c\n", ch)
	fmt.Printf("  Rune japanese: %c\n", japanese)
	fmt.Printf("  Rune emoji: %c\n", emoji)

	// ─────────────────────────────────────────────────────────────────────────

	fmt.Println("\n════════════════════════════════════════════════════════════")
	fmt.Println("RUNE TO STRING CONVERSION")
	fmt.Println("════════════════════════════════════════════════════════════\n")

	runeChar := 'G'
	convertedStr := string(runeChar)

	fmt.Printf("Rune: %c (type: rune/int32)\n", runeChar)
	fmt.Printf("Converted to string: \"%s\" (type: string)\n", convertedStr)
	fmt.Printf("Type check using %%T format verb: %T\n", convertedStr)

	// ─────────────────────────────────────────────────────────────────────────

	fmt.Println("\n════════════════════════════════════════════════════════════")
	fmt.Println("UNICODE AND INTERNATIONAL TEXT")
	fmt.Println("════════════════════════════════════════════════════════════\n")

	englishHello := "Hello"
	japaneseHello := "こんにちは" // Japanese "hello"
	spanishHello := "Hola"
	arabicHello := "مرحبا" // Arabic "hello"
	chineseHello := "你好"   // Chinese "hello"

	fmt.Println("Supporting multiple languages:")
	fmt.Printf("English:  %s\n", englishHello)
	fmt.Printf("Japanese: %s\n", japaneseHello)
	fmt.Printf("Spanish:  %s\n", spanishHello)
	fmt.Printf("Arabic:   %s\n", arabicHello)
	fmt.Printf("Chinese:  %s\n", chineseHello)

	fmt.Println("\nGo has NATIVE Unicode support - perfect for global applications!")

	// ─────────────────────────────────────────────────────────────────────────

	fmt.Println("\n════════════════════════════════════════════════════════════")
	fmt.Println("ITERATING OVER RUNES IN INTERNATIONAL TEXT")
	fmt.Println("════════════════════════════════════════════════════════════\n")

	japaneseWord := "こんにちは" // "hello" in Japanese
	fmt.Printf("Iterating over Japanese text: \"%s\"\n\n", japaneseWord)

	for index, runeValue := range japaneseWord {
		fmt.Printf("Position %d: %c (Unicode: U+%04X, Decimal: %d)\n",
			index, runeValue, runeValue, runeValue)
	}

	// ─────────────────────────────────────────────────────────────────────────

	fmt.Println("\n════════════════════════════════════════════════════════════")
	fmt.Println("COUNTING RUNES VS BYTES")
	fmt.Println("════════════════════════════════════════════════════════════\n")

	mixedText := "Hello こんにちは"
	byteCount := len(mixedText)
	runeCount := utf8.RuneCountInString(mixedText)

	fmt.Printf("Text: \"%s\"\n", mixedText)
	fmt.Printf("Byte count (len()): %d bytes\n", byteCount)
	fmt.Printf("Rune count (utf8.RuneCountInString()): %d characters\n", runeCount)
	fmt.Println("\nNote: UTF-8 uses multiple bytes for non-ASCII characters!")

	// ─────────────────────────────────────────────────────────────────────────

	fmt.Println("\n════════════════════════════════════════════════════════════")
	fmt.Println("COMPARISON: RUNES vs CHARACTERS (from C)")
	fmt.Println("════════════════════════════════════════════════════════════\n")

	fmt.Println(`
RUNES (Go):                          CHARACTERS (C):
- Type: int32 (4 bytes)              - Type: char (1 byte)
- Range: 0 to 1,114,111             - Range: 0 to 255
- Supports Unicode globally          - ASCII only (mostly)
- Can handle any language            - Limited to ASCII/extended ASCII
- Built-in support for i18n          - Needs external libraries
- Efficient for multilingual apps    - Difficult for non-ASCII text

EXAMPLE DIFFERENCES:
┌─────────────────┬──────────┬──────────┐
│ Character       │ Go Rune  │ C char   │
├─────────────────┼──────────┼──────────┤
│ 'A'             │ int32: 65│ int8: 65 │
│ 'é' (e-acute)   │ int32:233│ ❌ Error │
│ '日' (Japanese) │ int32:26085│ ❌ Error│
│ '😊' (emoji)    │ int32:128522│ ❌ Error│
└─────────────────┴──────────┴──────────┘

Go's rune design reflects its philosophy:
- Simplicity: No need for external libraries
- Efficiency: Built-in Unicode support
- Internationalization: Handle global text naturally
	`)

	// ─────────────────────────────────────────────────────────────────────────

	fmt.Println("\n════════════════════════════════════════════════════════════")
	fmt.Println("SUMMARY & BEST PRACTICES")
	fmt.Println("════════════════════════════════════════════════════════════\n")

	fmt.Println(`
KEY TAKEAWAYS:

STRINGS:
✓ Immutable sequences of bytes
✓ Created with double quotes (with escape sequences) or backticks (raw)
✓ Concatenate using + operator
✓ Compare using relational operators (lexicographical order)
✓ Iterate using range (returns runes, not bytes)
✓ Get length with len() function

RUNES:
✓ Type alias for int32
✓ Represent Unicode code points
✓ Declared with single quotes: 'A', '日', '😊'
✓ Essential for international/multilingual text
✓ Iterate using range over strings
✓ Convert to string using string(rune)

INTERNATIONALIZATION:
✓ Go has native, built-in Unicode support
✓ No external libraries needed for international text
✓ Easily handle text from any language
✓ Perfect for global applications

STRING OPERATIONS:
✓ Strings are immutable - create new strings for modifications
✓ Use + for concatenation or string() for conversion
✓ Remember: len() counts BYTES, not UTF-8 characters
✓ Use utf8.RuneCountInString() to count actual characters

BEST PRACTICES:
✓ Use double quotes for normal strings, backticks for raw strings
✓ Be aware of escape sequences when using double quotes
✓ Remember: range over strings gives runes, not individual bytes
✓ For non-ASCII text, use runes and range iteration
✓ Count characters with utf8.RuneCountInString(), not len()
✓ Leverage Go's Unicode support for global applications
✓ Test with international characters if your app is global

PERFORMANCE:
✓ Strings are efficient and immutable
✓ Use strings.Builder for efficient concatenation in loops
✓ Remember: UTF-8 encoded strings use variable byte lengths

Understanding strings and runes is crucial for writing robust Go applications,
especially those that need to handle text from multiple languages and regions!
	`)
}
