package intermediate

import (
	"fmt"
	"unicode"
)

// EXAMPLE 3: String to Rune Conversion
// EXAMPLE 5: String Indexing vs Range - The Critical Difference
// EXAMPLE 7: Modifying Strings (Creating New Ones)
// EXAMPLE 9: String Comparison and Ordering
// BONUS: Common Rune Patterns

func stringRuneBonusExample() {
	fmt.Println("\n=== BONUS: Common Rune Patterns ===")
	
	// Pattern 1: Check if string contains specific rune type
	text := "Hello123!"
	hasDigit := hasDigitRune(text)
	fmt.Printf("'%s' contains digit: %v\n", text, hasDigit)
	
	// Pattern 2: Remove specific runes
	fmt.Println("\nRemove punctuation:")
	dirty := "Hello, World!"
	clean := removePunctuation(dirty)
	fmt.Printf("Original: %s\n", dirty)
	fmt.Printf("Cleaned: %s\n", clean)
}

func hasDigitRune(s string) bool {
	for _, r := range s {
		if unicode.IsDigit(r) {
			return true
		}
	}
	return false
}

func removePunctuation(s string) string {
	return strings.Map(func(r rune) rune {
		if unicode.IsPunct(r) {
			return -1 // Remove this rune
		}
		return r
	}, s)
}

// MAIN FUNCTION

func main() {
	fmt.Println("╗")
	fmt.Println("║    STRINGS AND RUNES - DEEP DIVE      ║")
	fmt.Println("")
	
	
	fmt.Println("\n╗")