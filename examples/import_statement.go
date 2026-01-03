package main

// Import Statement — single and grouped imports, aliasing, and blank imports.

import (
	"fmt"
	_ "net/http/pprof" // blank import example (registers init side-effects)
	strings2 "strings" // alias example
)

func main() {
	// Use alias for clarity in examples
	s := strings2.ToUpper("import statement example")
	fmt.Println(s)
	fmt.Println("Note: Use a blank import when you need init side-effects from a package.")
}
