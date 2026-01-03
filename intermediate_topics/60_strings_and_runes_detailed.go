package main

import (
	"fmt"
	"unicode/utf8"
)

/*
Topic 60: STRINGS AND RUNES

CONCEPT: Strings are immutable sequences of bytes. Runes are individual
Unicode characters (type int32).

STRINGS:
- Declared with double quotes: "hello" (supports escape sequences)
- Or backticks: `raw string` (no escape sequences)
- Immutable - create new strings to "modify"
- Compared lexicographically (dictionary order)

ESCAPE SEQUENCES:
\n = newline, \t = tab, \r = carriage return, \\ = backslash, \" = quote

RUNES:
- Type alias for int32 (represents Unicode code point)
- Declared with single quotes: 'A', '日', '😊'
- Support any language and emoji
- Essential for international applications

UTF-8:
- Variable-length encoding (1-4 bytes per character)
- len() counts bytes, not characters
- Use utf8.RuneCountInString() for character count

KEY DIFFERENCE: Rune (go) = Unicode support, Char (C) = ASCII only
*/

func main() {
	fmt.Println("═══════════════════════════════════════════════════════════")
	fmt.Println("PART 1: STRING DECLARATION AND BASICS")
	fmt.Println("═══════════════════════════════════════════════════════════\n")

	// Standard string with escape sequences
	msg1 := "Hello\nGo"
	fmt.Println("String with \\n (newline):")
	fmt.Println(msg1)

	// Raw string (no escape sequences)
	msg2 := `Hello\nGo`
	fmt.Println("\nRaw string (backticks, \\n treated literally):")
	fmt.Println(msg2)

	// With tab
	msg3 := "Hello\tGo"
	fmt.Println("\nString with \\t (tab):")
	fmt.Println(msg3)
	fmt.Println()

	// ─────────────────────────────────────────────────────────────────────────

	fmt.Println("═══════════════════════════════════════════════════════════")
	fmt.Println("STRING LENGTH & INDEXING")
	fmt.Println("═══════════════════════════════════════════════════════════\n")

	str := "Hello"
	fmt.Printf("String: \"%s\"\n", str)
	fmt.Printf("Length: %d\n", len(str))
	fmt.Printf("Index 0: %c (ASCII/byte value: %d)\n", str[0], str[0])
	fmt.Printf("Index 1: %c (ASCII/byte value: %d)\n\n", str[1], str[1])

	// ─────────────────────────────────────────────────────────────────────────

	fmt.Println("═══════════════════════════════════════════════════════════")
	fmt.Println("STRING CONCATENATION & COMPARISON")
	fmt.Println("═══════════════════════════════════════════════════════════\n")

	greeting := "Hello"
	name := "Alice"

	fmt.Printf("Concatenation (no space): \"%s\" + \"%s\" = \"%s\"\n", greeting, name, greeting+name)
	fmt.Printf("With manual space: \"%s\"\n", greeting+" "+name)
	fmt.Printf("Using Print (auto space): ")
	fmt.Println(greeting, name)

	fmt.Println("\nLexicographical comparison:")
	fmt.Printf("\"Apple\" < \"apple\": %v (uppercase < lowercase)\n", "Apple" < "apple")
	fmt.Printf("\"app\" < \"apple\": %v (prefix is smaller)\n", "app" < "apple")
	fmt.Printf("\"banana\" > \"apple\": %v (b > a)\n\n", "banana" > "apple")

	// ─────────────────────────────────────────────────────────────────────────

	fmt.Println("═══════════════════════════════════════════════════════════")
	fmt.Println("STRING ITERATION & IMMUTABILITY")
	fmt.Println("═══════════════════════════════════════════════════════════\n")

	text := "Hello"
	fmt.Printf("Iterating over \"%s\" with range:\n", text)
	for idx, char := range text {
		fmt.Printf("  Index %d: '%c' (rune: %v)\n", idx, char, char)
	}

	fmt.Println("\nImmutability (creating new string):")
	original := "Hello"
	modified := original + " World"
	fmt.Printf("Original: \"%s\"\n", original)
	fmt.Printf("Modified: \"%s\"\n", modified)
	fmt.Println("Original unchanged - strings are immutable!\n")

	// ─────────────────────────────────────────────────────────────────────────

	fmt.Println("═══════════════════════════════════════════════════════════")
	fmt.Println("PART 2: RUNES - UNICODE SUPPORT")
	fmt.Println("═══════════════════════════════════════════════════════════\n")

	var ch rune = 'A'
	var japanese rune = '日' // "day" in Japanese
	var emoji rune = '😊'    // smiley emoji

	fmt.Println("Rune values (as integers):")
	fmt.Printf("'A': %d (rune/int32)\n", ch)
	fmt.Printf("'日': %d\n", japanese)
	fmt.Printf("'😊': %d\n\n", emoji)

	fmt.Println("Rune values (as characters):")
	fmt.Printf("'A': %c\n", ch)
	fmt.Printf("'日': %c\n", japanese)
	fmt.Printf("'😊': %c\n\n", emoji)

	// ─────────────────────────────────────────────────────────────────────────

	fmt.Println("═══════════════════════════════════════════════════════════")
	fmt.Println("INTERNATIONAL TEXT & UNICODE")
	fmt.Println("═══════════════════════════════════════════════════════════\n")

	english := "Hello"
	japanese_hello := "こんにちは"
	spanish := "Hola"
	arabic := "مرحبا"
	chinese := "你好"

	fmt.Println("Multiple languages:")
	fmt.Printf("English:   %s\n", english)
	fmt.Printf("Japanese:  %s\n", japanese_hello)
	fmt.Printf("Spanish:   %s\n", spanish)
	fmt.Printf("Arabic:    %s\n", arabic)
	fmt.Printf("Chinese:   %s\n\n", chinese)

	// ─────────────────────────────────────────────────────────────────────────

	fmt.Println("═══════════════════════════════════════════════════════════")
	fmt.Println("BYTES VS CHARACTERS (UTF-8)")
	fmt.Println("═══════════════════════════════════════════════════════════\n")

	mixedText := "Hi 日本"
	byteCount := len(mixedText)
	runeCount := utf8.RuneCountInString(mixedText)

	fmt.Printf("Text: \"%s\"\n", mixedText)
	fmt.Printf("len(string):            %d bytes\n", byteCount)
	fmt.Printf("utf8.RuneCountInString: %d characters\n\n", runeCount)

	fmt.Println("Iterating runes in international text:")
	for idx, rune := range mixedText {
		fmt.Printf("  Position %d: %c (U+%04X)\n", idx, rune, rune)
	}
	fmt.Println()

	// ─────────────────────────────────────────────────────────────────────────

	fmt.Println("═══════════════════════════════════════════════════════════")
	fmt.Println("KEY CONCEPTS & BEST PRACTICES")
	fmt.Println("═══════════════════════════════════════════════════════════\n")

	fmt.Println(`
STRINGS:
✓ Immutable sequences of bytes
✓ Double quotes support escape sequences
✓ Backticks = raw strings (no escapes)
✓ Concatenate with + operator
✓ Lexicographical comparison (a < b < c, A < a)
✓ Iterate with range (returns runes, not bytes)

RUNES:
✓ Type int32 for Unicode code points
✓ Single quotes: 'A', '日', '😊'
✓ Handle any language naturally
✓ Essential for internationalization

UTF-8 ENCODING:
✓ Variable-length: 1-4 bytes per character
✓ len() = byte count (not character count!)
✓ utf8.RuneCountInString() = actual characters
✓ range gives correct rune iteration

KEY DIFFERENCES:
Go Rune:         C Char:
- int32 (Unicode) - byte (ASCII only)
- Handles all languages - limited to 255 values
- Perfect for i18n - needs external libraries

BEST PRACTICES:
✓ Use double quotes for normal strings
✓ Use backticks for raw strings (regex, file paths)
✓ Remember: range over strings gives runes, not bytes
✓ Use utf8.RuneCountInString() for character count
✓ Leverage Go's Unicode support for global apps
✓ Test with international characters if global
✓ Be aware UTF-8 uses variable byte lengths

Go's built-in Unicode support is one of its best features!
	`)
}
