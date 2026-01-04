package main

import (
	"fmt"
	"regexp"
	"strings"
)

// ============================================================
// REGEX IN GO: COMPREHENSIVE GUIDE
// ============================================================
//
// TL;DR: Go's regexp package uses RE2 syntax for pattern matching.
// Compile patterns once (expensive), execute many times (cheap).
// Use MatchString to check, FindString/FindAllString to extract,
// ReplaceAllString to substitute, and FindAllStringSubmatch for
// capturing groups.
// ============================================================

func main() {
	fmt.Println("=== GO REGEX: Comprehensive Guide ===\n")

	// ============================================================
	// PART 1: BASICS - Does a pattern exist? (MatchString)
	// ============================================================
	part1MatchString()

	// ============================================================
	// PART 2: FINDING - Extract data from text
	// ============================================================
	part2FindingPatterns()

	// ============================================================
	// PART 3: REPLACING - Modify matched patterns
	// ============================================================
	part3ReplacingPatterns()

	// ============================================================
	// PART 4: CAPTURING - Extract specific parts (Groups)
	// ============================================================
	part4CapturingGroups()

	// ============================================================
	// PART 5: PRACTICAL EXAMPLES - Real-world use cases
	// ============================================================
	part5PracticalExamples()

	// ============================================================
	// PART 6: PERFORMANCE & BEST PRACTICES
	// ============================================================
	part6BestPractices()
}

// ============================================================
// PART 1: BASICS - Does a pattern exist? (MatchString)
// ============================================================

func part1MatchString() {
	fmt.Println("\n" + strings.Repeat("=", 70))
	fmt.Println("PART 1: BASICS - Pattern Matching with MatchString")
	fmt.Println(strings.Repeat("=", 70) + "\n")

	fmt.Println("📌 CONCEPT:")
	fmt.Println("===========")
	fmt.Println(`
MatchString checks if a regex pattern exists anywhere in the text.
Returns true if found, false otherwise.

Use case: Validate that a string contains digits, letters, special chars, etc.
`)

	fmt.Println("\n📝 CODE EXAMPLE 1: Check for digits\n")

	code1 := `
pattern := regexp.MustCompile("[0-9]+")

texts := []string{
    "I have 5 apples",      // TRUE (contains digits)
    "No numbers here",      // FALSE (no digits)
    "Year is 2025",         // TRUE (contains digits)
}

for _, text := range texts {
    found := pattern.MatchString(text)
    fmt.Printf("%q has digits: %v\n", text, found)
}
	`
	fmt.Println("```go")
	fmt.Println(code1)
	fmt.Println("```")

	fmt.Println("\n💡 LINE-BY-LINE EXPLANATION:")
	fmt.Println("- Line 1: Compile the pattern [0-9]+ (one or more digits)")
	fmt.Println("- Line 3-6: Define test strings")
	fmt.Println("- Line 8: For each string, call MatchString()")
	fmt.Println("- Line 9: Print whether digits were found")

	fmt.Println("\n🔄 LIVE EXECUTION:\n")

	pattern := regexp.MustCompile("[0-9]+")

	texts := []string{
		"I have 5 apples",
		"No numbers here",
		"Year is 2025",
	}

	for _, text := range texts {
		found := pattern.MatchString(text)
		status := "❌"
		if found {
			status = "✓"
		}
		fmt.Printf("%s %q has digits: %v\n", status, text, found)
	}

	fmt.Println("\n\n📝 CODE EXAMPLE 2: Check for email-like pattern\n")

	code2 := `
emailPattern := regexp.MustCompile("[a-zA-Z0-9]+@[a-zA-Z0-9]+\\.[a-zA-Z]{2,}")

emails := []string{
    "alice@example.com",    // TRUE (valid email pattern)
    "bob@test",             // FALSE (missing .com)
    "charlie@domain.co.uk", // TRUE (valid)
}

for _, email := range emails {
    isValid := emailPattern.MatchString(email)
    fmt.Printf("%q is email-like: %v\n", email, isValid)
}
	`
	fmt.Println("```go")
	fmt.Println(code2)
	fmt.Println("```")

	fmt.Println("\n💡 EXPLANATION:")
	fmt.Println("- [a-zA-Z0-9]+ : One or more alphanumeric characters (username)")
	fmt.Println("- @ : Literal @ symbol")
	fmt.Println("- \\. : Escaped dot (literal period, not wildcard)")
	fmt.Println("- [a-zA-Z]{2,} : At least 2 letters (domain extension like 'com')")

	fmt.Println("\n🔄 LIVE EXECUTION:\n")

	emailPattern := regexp.MustCompile("[a-zA-Z0-9]+@[a-zA-Z0-9]+\\.[a-zA-Z]{2,}")

	emails := []string{
		"alice@example.com",
		"bob@test",
		"charlie@domain.co.uk",
	}

	for _, email := range emails {
		isValid := emailPattern.MatchString(email)
		status := "✓"
		if !isValid {
			status = "❌"
		}
		fmt.Printf("%s %q is email-like: %v\n", status, email, isValid)
	}

	fmt.Println("\n✅ KEY TAKEAWAY:")
	fmt.Println("MatchString is the simplest way to check if a pattern exists. Returns bool.")
}

// ============================================================
// PART 2: FINDING - Extract data from text
// ============================================================

func part2FindingPatterns() {
	fmt.Println("\n\n" + strings.Repeat("=", 70))
	fmt.Println("PART 2: FINDING - Extract patterns from text")
	fmt.Println(strings.Repeat("=", 70) + "\n")

	fmt.Println("📌 CONCEPT:")
	fmt.Println("===========")
	fmt.Println(`
Three levels of finding:
1. FindString  → First match (string)
2. FindAllString → All matches ([]string)
3. FindAllStringIndex → All matches with positions

Use case: Extract emails, phone numbers, hashtags, URLs from text.
`)

	fmt.Println("\n📝 CODE EXAMPLE 1: FindString (First match only)\n")

	code1 := `
emailRegex := regexp.MustCompile("[a-zA-Z0-9]+@[a-zA-Z0-9]+\\.[a-zA-Z]{2,}")

text := "Contact alice@example.com or bob@test.org for help"

firstEmail := emailRegex.FindString(text)
fmt.Println("First email found:", firstEmail)  // Output: alice@example.com
	`
	fmt.Println("```go")
	fmt.Println(code1)
	fmt.Println("```")

	fmt.Println("\n💡 EXPLANATION:")
	fmt.Println("- FindString returns the first match as a string")
	fmt.Println("- Returns empty string if no match found")
	fmt.Println("- Useful when you only need one result")

	fmt.Println("\n🔄 LIVE EXECUTION:\n")

	emailRegex := regexp.MustCompile("[a-zA-Z0-9]+@[a-zA-Z0-9]+\\.[a-zA-Z]{2,}")
	text := "Contact alice@example.com or bob@test.org for help"
	firstEmail := emailRegex.FindString(text)
	fmt.Printf("First email found: %s\n", firstEmail)

	fmt.Println("\n\n📝 CODE EXAMPLE 2: FindAllString (All matches)\n")

	code2 := `
emailRegex := regexp.MustCompile("[a-zA-Z0-9]+@[a-zA-Z0-9]+\\.[a-zA-Z]{2,}")

text := "Contact alice@example.com or bob@test.org for help"

// -1 means "find all" (no limit)
allEmails := emailRegex.FindAllString(text, -1)

fmt.Println("All emails found:")
for i, email := range allEmails {
    fmt.Printf("  %d. %s\n", i+1, email)
}
	`
	fmt.Println("```go")
	fmt.Println(code2)
	fmt.Println("```")

	fmt.Println("\n💡 EXPLANATION:")
	fmt.Println("- FindAllString returns []string of ALL matches")
	fmt.Println("- Second parameter is limit: -1 = all, 2 = first 2, etc.")
	fmt.Println("- Returns nil (empty slice) if no matches")

	fmt.Println("\n🔄 LIVE EXECUTION:\n")

	allEmails := emailRegex.FindAllString(text, -1)
	fmt.Println("All emails found:")
	for i, email := range allEmails {
		fmt.Printf("  %d. %s\n", i+1, email)
	}

	fmt.Println("\n\n📝 CODE EXAMPLE 3: FindAllString with limit\n")

	code3 := `
hashtagRegex := regexp.MustCompile("#\\w+")

text := "#golang is great #programming language #go rocks"

// Get only first 2 hashtags
limited := hashtagRegex.FindAllString(text, 2)

fmt.Println("First 2 hashtags:", limited)  // [#golang #programming]
	`
	fmt.Println("```go")
	fmt.Println(code3)
	fmt.Println("```")

	fmt.Println("\n💡 EXPLANATION:")
	fmt.Println("- Limit parameter controls how many matches to return")
	fmt.Println("- Useful for performance (don't need to match entire text)")

	fmt.Println("\n🔄 LIVE EXECUTION:\n")

	hashtagRegex := regexp.MustCompile("#\\w+")
	text2 := "#golang is great #programming language #go rocks"
	limited := hashtagRegex.FindAllString(text2, 2)
	fmt.Printf("First 2 hashtags: %v\n", limited)

	fmt.Println("\n✅ KEY TAKEAWAY:")
	fmt.Println("FindString gets first, FindAllString gets all. Use limit parameter for performance.")
}

// ============================================================
// PART 3: REPLACING - Modify matched patterns
// ============================================================

func part3ReplacingPatterns() {
	fmt.Println("\n\n" + strings.Repeat("=", 70))
	fmt.Println("PART 3: REPLACING - Modify matched patterns")
	fmt.Println(strings.Repeat("=", 70) + "\n")

	fmt.Println("📌 CONCEPT:")
	fmt.Println("===========")
	fmt.Println(`
Three ways to replace:
1. ReplaceAllString → Replace with static string
2. ReplaceAllStringFunc → Replace using a function (dynamic)
3. ReplaceAllLiteral → Replace treating replacement as literal (no special chars)

Use case: Censor data, format text, remove patterns, normalize input.
`)

	fmt.Println("\n📝 CODE EXAMPLE 1: ReplaceAllString (Static replacement)\n")

	code1 := `
// Replace all digits with [NUM]
digitRegex := regexp.MustCompile("[0-9]+")

original := "I have 2 apples and 5 oranges"
modified := digitRegex.ReplaceAllString(original, "[NUM]")

fmt.Println("Original:", original)
fmt.Println("Modified:", modified)  // I have [NUM] apples and [NUM] oranges
	`
	fmt.Println("```go")
	fmt.Println(code1)
	fmt.Println("```")

	fmt.Println("\n💡 EXPLANATION:")
	fmt.Println("- ReplaceAllString replaces ALL matches with a static string")
	fmt.Println("- Original string is not modified (strings are immutable)")
	fmt.Println("- Returns new string with replacements")

	fmt.Println("\n🔄 LIVE EXECUTION:\n")

	digitRegex := regexp.MustCompile("[0-9]+")
	original := "I have 2 apples and 5 oranges"
	modified := digitRegex.ReplaceAllString(original, "[NUM]")
	fmt.Printf("Original: %s\n", original)
	fmt.Printf("Modified: %s\n", modified)

	fmt.Println("\n\n📝 CODE EXAMPLE 2: ReplaceAllStringFunc (Dynamic replacement)\n")

	code2 := `
// Replace all words with UPPERCASE version
wordRegex := regexp.MustCompile("[a-z]+")

text := "hello world golang"

// Function receives each match and returns replacement
replaced := wordRegex.ReplaceAllStringFunc(text, func(match string) string {
    return strings.ToUpper(match)
})

fmt.Println("Original:", text)
fmt.Println("Modified:", replaced)  // HELLO WORLD GOLANG
	`
	fmt.Println("```go")
	fmt.Println(code2)
	fmt.Println("```")

	fmt.Println("\n💡 EXPLANATION:")
	fmt.Println("- ReplaceAllStringFunc calls a function for EACH match")
	fmt.Println("- Function receives matched string, returns replacement")
	fmt.Println("- Allows dynamic, context-aware replacements")

	fmt.Println("\n🔄 LIVE EXECUTION:\n")

	wordRegex := regexp.MustCompile("[a-z]+")
	text := "hello world golang"
	replaced := wordRegex.ReplaceAllStringFunc(text, func(match string) string {
		return strings.ToUpper(match)
	})
	fmt.Printf("Original: %s\n", text)
	fmt.Printf("Modified: %s\n", replaced)

	fmt.Println("\n✅ KEY TAKEAWAY:")
	fmt.Println("ReplaceAllString for static, ReplaceAllStringFunc for dynamic replacements.")
}

// ============================================================
// PART 4: CAPTURING - Extract specific parts (Groups)
// ============================================================

func part4CapturingGroups() {
	fmt.Println("\n\n" + strings.Repeat("=", 70))
	fmt.Println("PART 4: CAPTURING - Extract specific parts (Groups)")
	fmt.Println(strings.Repeat("=", 70) + "\n")

	fmt.Println("📌 CONCEPT:")
	fmt.Println("===========")
	fmt.Println(`
Use parentheses () in regex to create capturing groups.
Extract specific parts of a match, not just the whole thing.

Use case: Parse structured data (phone numbers, dates, names).
`)

	fmt.Println("\n📝 CODE EXAMPLE 1: Phone number parsing\n")

	code1 := `
// Pattern: (555)-123-4567
// Groups: (1) area code  (2) prefix  (3) line number
phoneRegex := regexp.MustCompile("(\\d{3})-(\\d{3})-(\\d{4})")

text := "Call me at 555-123-4567"

// FindAllStringSubmatch returns: [["555-123-4567", "555", "123", "4567"]]
//   [0] = entire match
//   [1] = first group (area code)
//   [2] = second group (prefix)
//   [3] = third group (line number)
matches := phoneRegex.FindAllStringSubmatch(text, -1)

if len(matches) > 0 {
    fullNumber := matches[0][0]
    areaCode := matches[0][1]
    prefix := matches[0][2]
    lineNum := matches[0][3]
    
    fmt.Printf("Full: %s, Area: %s, Prefix: %s, Line: %s\n",
        fullNumber, areaCode, prefix, lineNum)
}
	`
	fmt.Println("```go")
	fmt.Println(code1)
	fmt.Println("```")

	fmt.Println("\n💡 LINE-BY-LINE EXPLANATION:")
	fmt.Println("- (\\d{3}) = Group 1: exactly 3 digits")
	fmt.Println("- (\\d{3}) = Group 2: exactly 3 digits")
	fmt.Println("- (\\d{4}) = Group 3: exactly 4 digits")
	fmt.Println("- FindAllStringSubmatch returns [][]string (matrix of matches and groups)")
	fmt.Println("- matches[0][0] = entire match")
	fmt.Println("- matches[0][1] = first capturing group")

	fmt.Println("\n🔄 LIVE EXECUTION:\n")

	phoneRegex := regexp.MustCompile("(\\d{3})-(\\d{3})-(\\d{4})")
	text := "Call me at 555-123-4567"
	matches := phoneRegex.FindAllStringSubmatch(text, -1)

	if len(matches) > 0 {
		fullNumber := matches[0][0]
		areaCode := matches[0][1]
		prefix := matches[0][2]
		lineNum := matches[0][3]

		fmt.Printf("Full: %s, Area: %s, Prefix: %s, Line: %s\n",
			fullNumber, areaCode, prefix, lineNum)
	}

	fmt.Println("\n\n📝 CODE EXAMPLE 2: Date parsing (YYYY-MM-DD)\n")

	code2 := `
dateRegex := regexp.MustCompile("(\\d{4})-(\\d{2})-(\\d{2})")

text := "Born on 1990-05-15 in NYC"

matches := dateRegex.FindAllStringSubmatch(text, 1)

if len(matches) > 0 {
    year := matches[0][1]
    month := matches[0][2]
    day := matches[0][3]
    
    fmt.Printf("Year: %s, Month: %s, Day: %s\n", year, month, day)
}
	`
	fmt.Println("```go")
	fmt.Println(code2)
	fmt.Println("```")

	fmt.Println("\n🔄 LIVE EXECUTION:\n")

	dateRegex := regexp.MustCompile("(\\d{4})-(\\d{2})-(\\d{2})")
	text = "Born on 1990-05-15 in NYC"
	matches = dateRegex.FindAllStringSubmatch(text, 1)

	if len(matches) > 0 {
		year := matches[0][1]
		month := matches[0][2]
		day := matches[0][3]

		fmt.Printf("Year: %s, Month: %s, Day: %s\n", year, month, day)
	}

	fmt.Println("\n✅ KEY TAKEAWAY:")
	fmt.Println("Use FindAllStringSubmatch for groups. matches[i][0] is full match, [i][j] is group j.")
}

// ============================================================
// PART 5: PRACTICAL EXAMPLES - Real-world use cases
// ============================================================

func part5PracticalExamples() {
	fmt.Println("\n\n" + strings.Repeat("=", 70))
	fmt.Println("PART 5: PRACTICAL EXAMPLES - Real-world use cases")
	fmt.Println(strings.Repeat("=", 70) + "\n")

	fmt.Println("📌 EXAMPLE 1: Validate a password\n")

	code1 := `
// Password must have: uppercase, lowercase, digit, 8+ chars
passwordRegex := regexp.MustCompile("^(?=.*[A-Z])(?=.*[a-z])(?=.*[0-9]).{8,}$")

passwords := []string{
    "ValidPass123",   // ✓ has upper, lower, digit, 8+ chars
    "invalid",        // ✗ no upper, no digit
    "Pass1",          // ✗ too short
}

for _, pwd := range passwords {
    valid := passwordRegex.MatchString(pwd)
    status := "❌"
    if valid {
        status = "✓"
    }
    fmt.Printf("%s %q is valid\n", status, pwd)
}
	`
	fmt.Println("```go")
	fmt.Println(code1)
	fmt.Println("```")

	fmt.Println("\n🔄 LIVE EXECUTION:\n")

	passwordRegex := regexp.MustCompile("^(?=.*[A-Z])(?=.*[a-z])(?=.*[0-9]).{8,}$")
	passwords := []string{
		"ValidPass123",
		"invalid",
		"Pass1",
	}

	for _, pwd := range passwords {
		valid := passwordRegex.MatchString(pwd)
		status := "❌"
		if valid {
			status = "✓"
		}
		fmt.Printf("%s %q is valid\n", status, pwd)
	}

	fmt.Println("\n\n📌 EXAMPLE 2: Extract URLs from text\n")

	code2 := `
urlRegex := regexp.MustCompile("https?://[a-zA-Z0-9.-]+")

text := "Visit https://golang.org or http://github.com for code"

urls := urlRegex.FindAllString(text, -1)

fmt.Println("URLs found:")
for i, url := range urls {
    fmt.Printf("  %d. %s\n", i+1, url)
}
	`
	fmt.Println("```go")
	fmt.Println(code2)
	fmt.Println("```")

	fmt.Println("\n🔄 LIVE EXECUTION:\n")

	urlRegex := regexp.MustCompile("https?://[a-zA-Z0-9.-]+")
	urlText := "Visit https://golang.org or http://github.com for code"
	urls := urlRegex.FindAllString(urlText, -1)

	fmt.Println("URLs found:")
	for i, url := range urls {
		fmt.Printf("  %d. %s\n", i+1, url)
	}

	fmt.Println("\n\n📌 EXAMPLE 3: Censor sensitive data (phone numbers)\n")

	code3 := `
phoneRegex := regexp.MustCompile("\\d{3}-\\d{3}-\\d{4}")

text := "Call: 555-123-4567 or 555-987-6543"

// Replace all phone numbers with *** (using ReplaceAllString)
censored := phoneRegex.ReplaceAllString(text, "***")

fmt.Println("Original:", text)
fmt.Println("Censored:", censored)
	`
	fmt.Println("```go")
	fmt.Println(code3)
	fmt.Println("```")

	fmt.Println("\n🔄 LIVE EXECUTION:\n")

	phoneRegex := regexp.MustCompile("\\d{3}-\\d{3}-\\d{4}")
	phoneText := "Call: 555-123-4567 or 555-987-6543"
	censored := phoneRegex.ReplaceAllString(phoneText, "***")

	fmt.Printf("Original: %s\n", phoneText)
	fmt.Printf("Censored: %s\n", censored)

	fmt.Println("\n✅ KEY TAKEAWAY:")
	fmt.Println("Regex solves password validation, URL extraction, and data censoring elegantly.")
}

// ============================================================
// PART 6: PERFORMANCE & BEST PRACTICES
// ============================================================

func part6BestPractices() {
	fmt.Println("\n\n" + strings.Repeat("=", 70))
	fmt.Println("PART 6: PERFORMANCE & BEST PRACTICES")
	fmt.Println(strings.Repeat("=", 70) + "\n")

	fmt.Println("📌 BEST PRACTICE 1: Compile once, reuse many times\n")

	fmt.Println(`
WRONG (Inefficient):
  for i := 0; i < 1000000; i++ {
    pattern := regexp.MustCompile("[0-9]+")  // Compiles EVERY iteration! SLOW!
    pattern.MatchString(data[i])
  }

RIGHT (Efficient):
  pattern := regexp.MustCompile("[0-9]+")    // Compile ONCE
  for i := 0; i < 1000000; i++ {
    pattern.MatchString(data[i])             // Reuse many times. FAST!
  }

Why? Compilation (parsing regex syntax) is expensive. Matching is cheap.
`)

	fmt.Println("📌 BEST PRACTICE 2: Use raw strings (backticks) to avoid escaping\n")

	fmt.Println(`
WRONG (Hard to read):
  pattern := regexp.MustCompile("\\d{3}-\\d{3}-\\d{4}")

RIGHT (Cleaner):
  pattern := regexp.MustCompile("` + "`" + `\\d{3}-\\d{3}-\\d{4}` + "`" + `")

With raw strings, you only need ONE backslash instead of two.
`)

	fmt.Println("\n📌 BEST PRACTICE 3: Use ^ and $ for full-string matching\n")

	fmt.Println(`
Pattern: "[0-9]+"
Text: "abc123def"
MatchString returns TRUE (found digits somewhere)

Pattern: "^[0-9]+$"
Text: "abc123def"
MatchString returns FALSE (not ALL digits)

^ = start of string
$ = end of string
`)

	fmt.Println("\n📌 BEST PRACTICE 4: Understand anchors and word boundaries\n")

	fmt.Println(`
\\b = word boundary (between word and non-word char)
\\B = NOT a word boundary

Pattern: \\bgo\\b
Text: "golang golang go"
Matches: Only the standalone "go" (not in "golang")
`)

	fmt.Println("\n📌 BEST PRACTICE 5: Use Compile() for runtime patterns (returns error)\n")

	fmt.Println(`
MustCompile:
  pattern := regexp.MustCompile("[invalid")  // PANICS if invalid
  
Compile:
  pattern, err := regexp.Compile("[invalid")  // Returns error
  if err != nil {
    log.Fatal(err)
  }

Use Compile() when the pattern comes from user input!
Use MustCompile() for compile-time patterns (string literals).
`)

	fmt.Println("\n\n📌 COMMON REGEX PATTERNS (Copy-Paste Ready)\n")

	fmt.Println(`
Email (simple):
  [a-zA-Z0-9]+@[a-zA-Z0-9]+\\.[a-zA-Z]{2,}

Phone (XXX-XXX-XXXX):
  \\d{3}-\\d{3}-\\d{4}

Date (YYYY-MM-DD):
  \\d{4}-\\d{2}-\\d{2}

URL (http/https):
  https?://[a-zA-Z0-9.-]+

Digits only:
  ^[0-9]+$

Lowercase letters only:
  ^[a-z]+$

Uppercase letters only:
  ^[A-Z]+$

Word characters (letters, digits, underscore):
  ^\\w+$

Whitespace:
  \\s (single space)
  \\s+ (one or more spaces)

Hashtag:
  #\\w+

Mention (@username):
  @\\w+
`)

	fmt.Println("\n\n" + strings.Repeat("=", 70))
	fmt.Println("FINAL SUMMARY")
	fmt.Println(strings.Repeat("=", 70) + "\n")

	fmt.Println("✅ PART 1: Use MatchString to check if a pattern exists")
	fmt.Println("✅ PART 2: Use FindString/FindAllString to extract patterns")
	fmt.Println("✅ PART 3: Use ReplaceAllString/ReplaceAllStringFunc to modify")
	fmt.Println("✅ PART 4: Use FindAllStringSubmatch to extract groups")
	fmt.Println("✅ PART 5: Apply to real-world problems (validation, parsing, censoring)")
	fmt.Println("✅ PART 6: Compile once, reuse many. Use raw strings. Use anchors.")
	fmt.Println("\n🎯 Master these 6 parts, and you master Go regex.\n")
}
