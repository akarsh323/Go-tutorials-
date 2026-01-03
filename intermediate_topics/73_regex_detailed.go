package intermediate

import (
	"fmt"
	"regexp"
	"strings"
)

// Topic 73: regex
// Combines pragmatic examples with GoBootcamp patterns


func main() {

	fmt.Println("-- 73 Regular Expressions --")

	// MatchString: Check if pattern matches
	pattern := "^[a-z]+$"
	regex := regexp.MustCompile(pattern)
	fmt.Println("Match 'hello':", regex.MatchString("hello"))
	fmt.Println("Match 'hello123':", regex.MatchString("hello123"))

	// FindString: Find first match
	emailRegex := regexp.MustCompile(`\w+@\w+\.\w+`)
	text := "Contact: alice@example.com or bob@test.org"
	fmt.Println("First email:", emailRegex.FindString(text))

	// FindAllString: Find all matches (limit with -1 for no limit)
	wordRegex := regexp.MustCompile(`\b\w{5}\b`)
	s := "hello world golang programming"
	matches := wordRegex.FindAllString(s, -1)
	fmt.Println("5-letter words:", matches)

	// FindAllString with limit
	limitMatches := wordRegex.FindAllString(s, 2)
	fmt.Println("First 2 five-letter words:", limitMatches)

	// ReplaceAllString: Replace all matches
	replaceRegex := regexp.MustCompile(`\d+`)
	numbers := "I have 2 apples and 5 oranges"
	result := replaceRegex.ReplaceAllString(numbers, "[NUM]")
	fmt.Println("Replaced numbers:", result)

	// ReplaceAllStringFunc: Replace using a function
	funcReplace := regexp.MustCompile(`[a-z]+`)
	upper := funcReplace.ReplaceAllStringFunc("hello world", strings.ToUpper)
	fmt.Println("Uppercase all words:", upper)

	// Split: Split string by pattern
	splitRegex := regexp.MustCompile(`[,;]`)
	data := "apple,banana;cherry,date"
	parts := splitRegex.Split(data, -1)
	fmt.Println("Split by comma or semicolon:", parts)

	// Grouping and capturing (using FindAllStringSubmatch)
	phoneRegex := regexp.MustCompile(`(\d{3})-(\d{3})-(\d{4})`)
	phone := "Call me at 555-123-4567"
	matches = phoneRegex.FindAllStringSubmatch(phone, -1)
	if len(matches) > 0 {
		fmt.Printf("Phone: %s, Area Code: %s, Prefix: %s, Line: %s\n",
			matches[0][0], matches[0][1], matches[0][2], matches[0][3])
	}

	// MustCompile panics on invalid regex; Compile returns error
	invalidRegex, err := regexp.Compile("[invalid")
	fmt.Println("Invalid regex error:", err)
}

func regexExample1() {
	
	fmt.Println("📚 Find if a pattern exists in text")

	// Create regex pattern
	pattern := regexp.MustCompile(`[0-9]+`)

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
		fmt.Printf("%s %q has numbers: %v\n", status, text, found)
	}
}

func regexExample10() {
	
	fmt.Println(`
When to use Regex:
  ✓ Pattern matching (not exact matches)
  ✓ Data validation (emails, phones, dates)
  ✓ Finding/replacing complex patterns
  ✓ Extracting structured data

When NOT to use Regex:
  ✗ Simple string operations → use strings package
  ✗ Exact substring matching → use strings.Contains
  ✗ When strings package is simpler

Regex Tips:
  • Use ^ and $ to match whole strings
  • \b for word boundaries (word\b)
  • \d for digits, \w for word chars, \s for whitespace
  • + for one or more, * for zero or more, ? for optional
  • {n,m} for specific counts
  • [abc] for character class, [^abc] for NOT
  • | for alternatives (cat|dog)
  • () for capturing groups (use $1 in Replace)
  • Use raw strings (backticks) to avoid escaping

Performance:
  • Compile patterns once, reuse them
  • Use MustCompile for compile-time patterns
  • Use Compile for runtime patterns

Common Mistakes:
  ✗ Forgetting ^ and $ for full string match
  ✗ Using . without considering newlines
  ✗ Over-complex patterns (harder to maintain)
  ✗ Not escaping special characters
	`)
}
