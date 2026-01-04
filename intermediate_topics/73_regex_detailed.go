package main

import (
	"fmt"
	"log"
	"regexp"
	"strings"
)

// Topic 73: Regular Expressions (Regex)
// ====================================
//
// TL;DR: Go uses RE2 syntax for regex. Compile patterns once (expensive),
// execute many times (cheap). MatchString checks existence, FindString/FindAllString
// extract data, ReplaceAllString modifies, FindAllStringSubmatch captures groups.

func main() {
	fmt.Println("=== 73 REGULAR EXPRESSIONS: Deep Dive ===\n")

	// ============================================================
	// SECTION 1: Pattern Matching (MatchString)
	// ============================================================
	section1PatternMatching()

	// ============================================================
	// SECTION 2: Finding Patterns
	// ============================================================
	section2FindingPatterns()

	// ============================================================
	// SECTION 3: Replacing Patterns
	// ============================================================
	section3ReplacingPatterns()

	// ============================================================
	// SECTION 4: Capturing Groups
	// ============================================================
	section4CapturingGroups()

	// ============================================================
	// SECTION 5: Real-World Examples
	// ============================================================
	section5RealWorldExamples()
}

// ============================================================
// SECTION 1: Pattern Matching (MatchString)
// ============================================================

func section1PatternMatching() {
	fmt.Println(strings.Repeat("=", 70))
	fmt.Println("SECTION 1: Pattern Matching (MatchString)")
	fmt.Println(strings.Repeat("=", 70) + "\n")

	fmt.Println("📌 EXPLANATION:")
	fmt.Println("MatchString checks if a regex pattern exists in text.")
	fmt.Println("Returns true if found, false otherwise.\n")

	fmt.Println("Example 1: Check for digits\n")

	// MatchString: Check if pattern exists
	pattern := regexp.MustCompile("^[a-z]+$")

	testCases := []struct {
		text     string
		expected bool
	}{
		{"hello", true},
		{"hello123", false},
		{"HELLO", false},
		{"test", true},
	}

	fmt.Println("Pattern: ^[a-z]+$ (lowercase letters only)\n")
	for _, tc := range testCases {
		result := pattern.MatchString(tc.text)
		status := "❌"
		if result == tc.expected {
			status = "✓"
		}
		fmt.Printf("%s MatchString(%q) = %v\n", status, tc.text, result)
	}

	fmt.Println("\n\nExample 2: Check for email pattern\n")

	emailPattern := regexp.MustCompile("[a-zA-Z0-9]+@[a-zA-Z0-9]+\\.[a-zA-Z]{2,}")

	emails := []string{
		"alice@example.com",
		"bob@test",
		"charlie@domain.org",
	}

	fmt.Println("Pattern: Email-like [a-zA-Z0-9]+@[a-zA-Z0-9]+\\.[a-zA-Z]{2,}\n")
	for _, email := range emails {
		isValid := emailPattern.MatchString(email)
		status := "❌"
		if isValid {
			status = "✓"
		}
		fmt.Printf("%s %q -> %v\n", status, email, isValid)
	}

	fmt.Println("\n\nExample 3: Check for digits in string\n")

	digitPattern := regexp.MustCompile("[0-9]+")

	strings := []string{
		"I have 5 apples",
		"No numbers here",
		"Year 2025",
	}

	fmt.Println("Pattern: [0-9]+ (one or more digits)\n")
	for _, s := range strings {
		hasDigits := digitPattern.MatchString(s)
		status := "❌"
		if hasDigits {
			status = "✓"
		}
		fmt.Printf("%s %q has digits: %v\n", status, s, hasDigits)
	}

	fmt.Println("\n✅ KEY TAKEAWAY:")
	fmt.Println("MatchString is the simplest: returns bool, checks existence only.\n")
}

// ============================================================
// SECTION 2: Finding Patterns
// ============================================================

func section2FindingPatterns() {
	fmt.Println("\n" + strings.Repeat("=", 70))
	fmt.Println("SECTION 2: Finding Patterns (FindString / FindAllString)")
	fmt.Println(strings.Repeat("=", 70) + "\n")

	fmt.Println("📌 EXPLANATION:")
	fmt.Println("FindString returns the FIRST match.")
	fmt.Println("FindAllString returns ALL matches (with optional limit).\n")

	// FindString: Find first match
	fmt.Println("Example 1: FindString (first match only)\n")

	emailRegex := regexp.MustCompile("[a-zA-Z0-9]+@[a-zA-Z0-9]+\\.[a-zA-Z]{2,}")
	text := "Contact alice@example.com or bob@test.org"

	first := emailRegex.FindString(text)
	fmt.Printf("Text: %q\n", text)
	fmt.Printf("FindString: %q\n\n", first)

	// FindAllString: Find all matches (limit with -1 for no limit)
	fmt.Println("Example 2: FindAllString (all matches, limit 2)\n")

	wordRegex := regexp.MustCompile(`\b\w{5}\b`)
	text = "hello world golang programming language"
	matches := wordRegex.FindAllString(text, 2)

	fmt.Printf("Text: %q\n", text)
	fmt.Printf("Pattern: \\b\\w{5}\\b (5-letter words, word boundary)\n")
	fmt.Printf("Limit: 2 (only first 2 matches)\n")
	fmt.Printf("Result: %v\n\n", matches)

	// FindAllString with no limit
	fmt.Println("Example 3: FindAllString (all matches, no limit)\n")

	allMatches := wordRegex.FindAllString(text, -1)
	fmt.Printf("Text: %q\n", text)
	fmt.Printf("Pattern: \\b\\w{5}\\b\n")
	fmt.Printf("Limit: -1 (all matches)\n")
	fmt.Printf("Result: %v\n\n", allMatches)

	// Finding numbers
	fmt.Println("Example 4: Extract all numbers\n")

	numberRegex := regexp.MustCompile("[0-9]+")
	text = "Order 123 for item 456 with price 789"

	numbers := numberRegex.FindAllString(text, -1)
	fmt.Printf("Text: %q\n", text)
	fmt.Printf("Pattern: [0-9]+ (digits)\n")
	fmt.Printf("Numbers found: %v\n\n", numbers)

	fmt.Println("✅ KEY TAKEAWAY:")
	fmt.Println("FindString = first only. FindAllString = all (with limit option).\n")
}

// ============================================================
// SECTION 3: Replacing Patterns
// ============================================================

func section3ReplacingPatterns() {
	fmt.Println("\n" + strings.Repeat("=", 70))
	fmt.Println("SECTION 3: Replacing Patterns")
	fmt.Println(strings.Repeat("=", 70) + "\n")

	fmt.Println("📌 EXPLANATION:")
	fmt.Println("ReplaceAllString: Replace with static string.")
	fmt.Println("ReplaceAllStringFunc: Replace using a function (dynamic).\n")

	// ReplaceAllString: Replace all matches with static string
	fmt.Println("Example 1: ReplaceAllString (static replacement)\n")

	replaceRegex := regexp.MustCompile("[0-9]+")
	original := "I have 2 apples and 5 oranges"
	result := replaceRegex.ReplaceAllString(original, "[NUM]")

	fmt.Printf("Original:  %s\n", original)
	fmt.Printf("Modified:  %s\n\n", result)

	// ReplaceAllStringFunc: Replace using a function
	fmt.Println("Example 2: ReplaceAllStringFunc (dynamic replacement with function)\n")

	funcReplace := regexp.MustCompile("[a-z]+")
	text := "hello world golang"
	upper := funcReplace.ReplaceAllStringFunc(text, strings.ToUpper)

	fmt.Printf("Original:  %s\n", text)
	fmt.Printf("Modified:  %s\n\n", upper)

	// ReplaceAllStringFunc with custom logic
	fmt.Println("Example 3: ReplaceAllStringFunc (custom logic)\n")

	wordRegex := regexp.MustCompile("[a-z]+")
	text = "apple banana cherry"

	capitalized := wordRegex.ReplaceAllStringFunc(text, func(match string) string {
		// Capitalize first letter of each word
		return strings.ToUpper(match[:1]) + match[1:]
	})

	fmt.Printf("Original:      %s\n", text)
	fmt.Printf("Capitalized:   %s\n\n", capitalized)

	// Censor sensitive data
	fmt.Println("Example 4: Censor phone numbers\n")

	phoneRegex := regexp.MustCompile("[0-9]{3}-[0-9]{3}-[0-9]{4}")
	text = "Call: 555-123-4567 or 555-987-6543"
	censored := phoneRegex.ReplaceAllString(text, "***")

	fmt.Printf("Original:  %s\n", text)
	fmt.Printf("Censored:  %s\n\n", censored)

	fmt.Println("✅ KEY TAKEAWAY:")
	fmt.Println("ReplaceAllString for static. ReplaceAllStringFunc for dynamic logic.\n")
}

// ============================================================
// SECTION 4: Capturing Groups
// ============================================================

func section4CapturingGroups() {
	fmt.Println("\n" + strings.Repeat("=", 70))
	fmt.Println("SECTION 4: Capturing Groups")
	fmt.Println(strings.Repeat("=", 70) + "\n")

	fmt.Println("📌 EXPLANATION:")
	fmt.Println("Use parentheses () to capture specific parts of a match.")
	fmt.Println("FindAllStringSubmatch returns matrix: [match][group].\n")

	// Capturing phone number parts
	fmt.Println("Example 1: Parse phone number (XXX)-XXX-XXXX\n")

	phoneRegex := regexp.MustCompile("(\\d{3})-(\\d{3})-(\\d{4})")
	phone := "Call me at 555-123-4567"

	matches := phoneRegex.FindAllStringSubmatch(phone, -1)
	if len(matches) > 0 {
		fullNumber := matches[0][0]
		areaCode := matches[0][1]
		prefix := matches[0][2]
		lineNum := matches[0][3]

		fmt.Printf("Phone: %s\n", phone)
		fmt.Printf("Pattern: (\\d{3})-(\\d{3})-(\\d{4})\n\n")
		fmt.Printf("  [0] Full match:      %s\n", fullNumber)
		fmt.Printf("  [1] Area code:       %s\n", areaCode)
		fmt.Printf("  [2] Prefix:          %s\n", prefix)
		fmt.Printf("  [3] Line number:     %s\n\n", lineNum)
	}

	// Capturing date parts
	fmt.Println("Example 2: Parse date (YYYY-MM-DD)\n")

	dateRegex := regexp.MustCompile("(\\d{4})-(\\d{2})-(\\d{2})")
	text := "Born on 1990-05-15"

	matches = dateRegex.FindAllStringSubmatch(text, -1)
	if len(matches) > 0 {
		year := matches[0][1]
		month := matches[0][2]
		day := matches[0][3]

		fmt.Printf("Text: %s\n", text)
		fmt.Printf("Pattern: (\\d{4})-(\\d{2})-(\\d{2})\n\n")
		fmt.Printf("  Year:   %s\n", year)
		fmt.Printf("  Month:  %s\n", month)
		fmt.Printf("  Day:    %s\n\n", day)
	}

	// Capturing name parts
	fmt.Println("Example 3: Parse full name (FirstName LastName)\n")

	nameRegex := regexp.MustCompile("([A-Za-z]+) ([A-Za-z]+)")
	text = "My name is Alice Johnson"

	matches = nameRegex.FindAllStringSubmatch(text, 1)
	if len(matches) > 0 {
		firstName := matches[0][1]
		lastName := matches[0][2]

		fmt.Printf("Text: %s\n", text)
		fmt.Printf("Pattern: ([A-Za-z]+) ([A-Za-z]+)\n\n")
		fmt.Printf("  First name: %s\n", firstName)
		fmt.Printf("  Last name:  %s\n\n", lastName)
	}

	fmt.Println("✅ KEY TAKEAWAY:")
	fmt.Println("FindAllStringSubmatch returns [][]string. matches[i][0] = full, [i][j] = group j.\n")
}

// ============================================================
// SECTION 5: Real-World Examples
// ============================================================

func section5RealWorldExamples() {
	fmt.Println("\n" + strings.Repeat("=", 70))
	fmt.Println("SECTION 5: Real-World Examples")
	fmt.Println(strings.Repeat("=", 70) + "\n")

	// Validate email
	fmt.Println("Example 1: Validate email address\n")

	emailRegex := regexp.MustCompile("^[a-zA-Z0-9]+@[a-zA-Z0-9]+\\.[a-zA-Z]{2,}$")

	emails := []string{
		"alice@example.com",
		"bob@test",
		"charlie@domain.org",
		"invalid@",
	}

	fmt.Println("Pattern: ^[a-zA-Z0-9]+@[a-zA-Z0-9]+\\.[a-zA-Z]{2,}$\n")
	for _, email := range emails {
		valid := emailRegex.MatchString(email)
		status := "❌"
		if valid {
			status = "✓"
		}
		fmt.Printf("%s %q\n", status, email)
	}

	// Extract URLs
	fmt.Println("\n\nExample 2: Extract URLs from text\n")

	urlRegex := regexp.MustCompile("https?://[a-zA-Z0-9.-]+")
	text := "Visit https://golang.org or http://github.com for code"

	urls := urlRegex.FindAllString(text, -1)
	fmt.Printf("Text: %s\n\n", text)
	fmt.Println("URLs found:")
	for i, url := range urls {
		fmt.Printf("  %d. %s\n", i+1, url)
	}

	// Extract hashtags
	fmt.Println("\n\nExample 3: Extract hashtags from tweet\n")

	hashtagRegex := regexp.MustCompile("#\\w+")
	text = "#golang is #awesome for #programming"

	hashtags := hashtagRegex.FindAllString(text, -1)
	fmt.Printf("Text: %s\n\n", text)
	fmt.Println("Hashtags found:")
	for i, tag := range hashtags {
		fmt.Printf("  %d. %s\n", i+1, tag)
	}

	// Validate password
	fmt.Println("\n\nExample 4: Validate password strength\n")

	// Must have: uppercase, lowercase, digit, 8+ chars
	pwdRegex := regexp.MustCompile("^(?=.*[A-Z])(?=.*[a-z])(?=.*[0-9]).{8,}$")

	passwords := []string{
		"ValidPass123",
		"weakpass",
		"NoDigits",
		"Short1",
	}

	fmt.Println("Requirements: Uppercase, lowercase, digit, 8+ chars\n")
	for _, pwd := range passwords {
		valid := pwdRegex.MatchString(pwd)
		status := "❌"
		if valid {
			status = "✓"
		}
		fmt.Printf("%s %q\n", status, pwd)
	}

	// Split by pattern
	fmt.Println("\n\nExample 5: Split string by pattern\n")

	splitRegex := regexp.MustCompile("[,;]")
	data := "apple,banana;cherry,date"

	parts := splitRegex.Split(data, -1)
	fmt.Printf("Original: %s\n", data)
	fmt.Printf("Pattern: [,;] (comma or semicolon)\n\n")
	fmt.Println("Parts after split:")
	for i, part := range parts {
		fmt.Printf("  %d. %q\n", i, part)
	}

	// Error handling: Invalid regex
	fmt.Println("\n\nExample 6: Error handling with Compile\n")

	// MustCompile panics on invalid regex
	// Compile returns an error instead
	_, err := regexp.Compile("[invalid")
	if err != nil {
		fmt.Printf("Error compiling regex: %v\n", err)
	}

	// Use valid regex
	validRegex, err := regexp.Compile("[0-9]+")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Successfully compiled: %v\n", validRegex)
}
