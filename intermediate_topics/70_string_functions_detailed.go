package intermediate

import (
	"fmt"
	"strings"
)

// Topic 70: string_functions
// Combines pragmatic examples with GoBootcamp patterns


func main() {

	fmt.Println("-- 70 String Functions --")
	s := "Go is awesome"

	// String tests
	fmt.Println("HasPrefix(Go):", strings.HasPrefix(s, "Go"))
	fmt.Println("HasSuffix(awesome):", strings.HasSuffix(s, "awesome"))
	fmt.Println("Contains(is):", strings.Contains(s, "is"))
	fmt.Println("Index(is):", strings.Index(s, "is"))

	// Transformations
	fmt.Println("ToUpper:", strings.ToUpper(s))
	fmt.Println("ToLower:", strings.ToLower(s))
	fmt.Println("Title:", strings.Title(s))

	// Split and join
	parts := strings.Split(s, " ")
	fmt.Println("Split by space:", parts)
	joined := strings.Join(parts, "-")
	fmt.Println("Join with dash:", joined)

	// Replace
	replaced := strings.ReplaceAll(s, "awesome", "wonderful")
	fmt.Println("ReplaceAll awesome->wonderful:", replaced)

	// Trim
	messy := "  hello world  "
	fmt.Println("TrimSpace:", strings.TrimSpace(messy))
}

func stringFunctionsExample10() {
	
	fmt.Println(`
Most Common Functions:

Search & Check:
  Contains(s, substr)      → bool      (is substring in string?)
  Index(s, substr)         → int       (where is substring?)
  Count(s, substr)         → int       (how many occurrences?)
  HasPrefix/HasSuffix      → bool      (starts/ends with?)

Modify:
  ToUpper/ToLower          → string    (change case)
  Trim/TrimSpace           → string    (remove edges)
  Replace/ReplaceAll       → string    (substitute)
  Repeat                   → string    (duplicate)

Split & Join:
  Split(s, sep)            → []string  (break apart)
  Fields(s)                → []string  (split on whitespace)
  Join(slice, sep)         → string    (put together)

Best Practices:
  ✓ Use strings.Builder for multiple concatenations
  ✓ Use TrimSpace not manual trim
  ✓ Use HasPrefix/HasSuffix for type checking
  ✓ Use Fields instead of Split for whitespace
  ✓ Remember: strings are immutable, so these return new strings
	`)
}
