package main

import (
	"fmt"
	"strings"
)

// ============================================================================
// 90 COMMAND LINE ARGUMENTS / FLAGS
// ============================================================================
// Command line arguments allow you to pass dynamic parameters to your program
// from the terminal (e.g., `go run main.go hello world`).
//
// Two approaches:
// 1. Raw Arguments (os.Args): Simple, untyped slice of strings
// 2. Flags (flag package): Named parameters with types, defaults, and help text
// ============================================================================

// DEMO: Comprehensive Overview
func Demo90CmdArgsFlags() {
	fmt.Println("-- 90 Command Line Arguments / Flags --")
	fmt.Println()

	// Part 1: os.Args Explanation
	fmt.Println("📋 PART 1: RAW ARGUMENTS (os.Args)")
	fmt.Println(strings.Repeat("-", 70))
	fmt.Println()

	osArgsExplanation := `os.Args is a slice (list) of strings from the flag package.

KEY RULE:
- os.Args[0] = Program name/path (always present)
- os.Args[1], [2], etc. = Actual arguments passed by user

EXAMPLE:
Command: go run main.go Hello World 42

os.Args:
  [0]: /tmp/go-build.../exe  (temporary program executable)
  [1]: Hello                 (first argument)
  [2]: World                 (second argument)
  [3]: 42                    (third argument)

COMPLETE CODE EXAMPLE:

package main

import (
    "fmt"
    "os"
)

func main() {
    fmt.Println("Program name:", os.Args[0])
    fmt.Println("Total arguments:", len(os.Args)-1) // -1 to exclude program name
    
    if len(os.Args) > 1 {
        for i, arg := range os.Args[1:] {
            fmt.Printf("Arg %d: %s\n", i+1, arg)
        }
    } else {
        fmt.Println("No arguments provided")
    }
}

EXECUTION:
$ go run main.go apple banana cherry
Program name: /tmp/go-build.../exe
Total arguments: 3
Arg 1: apple
Arg 2: banana
Arg 3: cherry
`
	fmt.Println(osArgsExplanation)

	// Part 2: Flag Package Explanation
	fmt.Println()
	fmt.Println("🚩 PART 2: FLAGS (flag Package) - RECOMMENDED")
	fmt.Println(strings.Repeat("-", 70))
	fmt.Println()

	flagExplanation := `The flag package provides a cleaner way to handle arguments:
- Named parameters (e.g., -name, -age, -verbose)
- Type-safe (string, int, bool, float64, etc.)
- Default values for missing flags
- Built-in help text

THREE STEPS:
1. Declare variables to hold flag values
2. Bind flags using flag.StringVar, flag.IntVar, etc.
3. Call flag.Parse() to process command line input

COMPLETE CODE EXAMPLE:

package main

import (
    "flag"
    "fmt"
)

func main() {
    // Step 1: Declare variables
    var name string
    var age int
    var verbose bool

    // Step 2: Bind flags to variables
    // Syntax: flag.TYPE(&variable, "flag-name", defaultValue, "description")
    flag.StringVar(&name, "name", "Guest", "User's name")
    flag.IntVar(&age, "age", 18, "User's age")
    flag.BoolVar(&verbose, "verbose", false, "Enable verbose output")

    // Step 3: Parse the flags
    flag.Parse()

    // Use the values
    fmt.Printf("Name: %s, Age: %d, Verbose: %v\n", name, age, verbose)
}

EXECUTION:
$ go run main.go -name=Alice -age=30 -verbose
Name: Alice, Age: 30, Verbose: true

$ go run main.go -name=Bob
Name: Bob, Age: 18, Verbose: false

$ go run main.go -help
Usage of main:
  -age int
    	User's age (default 18)
  -name string
    	User's name (default "Guest")
  -verbose
    	Enable verbose output
`
	fmt.Println(flagExplanation)

	// Part 3: Important Nuances
	fmt.Println()
	fmt.Println("⚠️  IMPORTANT NUANCES")
	fmt.Println(strings.Repeat("-", 70))

	nuances := []string{
		"📌 SPACES IN ARGUMENTS:",
		"   • Without quotes: go run main.go -name James Doe",
		"     Result: name = \"James\", \"Doe\" = loose argument",
		"   • With quotes:    go run main.go -name \"James Doe\"",
		"     Result: name = \"James Doe\" ✓",
		"",
		"📌 TEMPORARY FILES:",
		"   • go run creates temp binary in /tmp/go-build...",
		"   • That's why os.Args[0] looks like a long path",
		"   • Use 'go build' to create a real executable",
		"",
		"📌 PARSING ORDER:",
		"   • Flags must come BEFORE positional arguments",
		"   • Once flag.Parse() is called, remaining args are in flag.Args()",
		"   • flag.Args() returns slice of non-flag arguments",
		"",
		"📌 POINTER REQUIREMENT:",
		"   • flag.StringVar(&name, ...) - Note the &",
		"   • The & gives the memory address (pointer)",
		"   • flag package needs it to update the original variable",
	}
	for _, item := range nuances {
		fmt.Println(item)
	}

	// Part 4: Practical Examples
	fmt.Println()
	fmt.Println("💡 PRACTICAL EXAMPLE: File Size CLI Tool")
	fmt.Println(strings.Repeat("-", 70))

	fileSizeExample := `package main

import (
    "flag"
    "fmt"
    "os"
)

func main() {
    // Define flags
    var filepath string
    var verbose bool

    flag.StringVar(&filepath, "file", "", "Path to the file")
    flag.BoolVar(&verbose, "verbose", false, "Show detailed info")
    flag.Parse()

    if filepath == "" {
        fmt.Println("Error: -file flag required")
        return
    }

    // Get file info
    info, err := os.Stat(filepath)
    if err != nil {
        fmt.Printf("Error: %v\n", err)
        return
    }

    fmt.Printf("File: %s\n", info.Name())
    fmt.Printf("Size: %d bytes\n", info.Size())
    
    if verbose {
        fmt.Printf("Modified: %s\n", info.ModTime())
        fmt.Printf("Is Dir: %v\n", info.IsDir())
    }
}

USAGE:
$ go run main.go -file=data.txt
File: data.txt
Size: 1024 bytes

$ go run main.go -file=data.txt -verbose
File: data.txt
Size: 1024 bytes
Modified: 2024-01-04 10:30:45
Is Dir: false
`
	fmt.Println(fileSizeExample)

	// Part 5: Advanced Flag Features
	fmt.Println()
	fmt.Println("🔧 ADVANCED FLAG FEATURES")
	fmt.Println(strings.Repeat("-", 70))

	advancedFeatures := []string{
		"1. MULTIPLE FLAG STYLES:",
		"   -name=value  OR  -name value",
		"   -v           (shorthand flags)",
		"",
		"2. BOOLEAN FLAGS:",
		"   -verbose     (true if present)",
		"   -verbose=true  OR  -verbose=false (explicit)",
		"",
		"3. REMAINING ARGUMENTS:",
		"   var remaining []string = flag.Args()",
		"   // Gets all arguments after flags",
		"",
		"4. BUILT-IN HELP:",
		"   -help or -h automatically shows all flags",
		"   // flag package generates this automatically",
		"",
		"5. FLAG TYPES:",
		"   flag.StringVar() - strings",
		"   flag.IntVar()    - integers",
		"   flag.BoolVar()   - booleans",
		"   flag.Float64Var() - floating point",
		"   flag.DurationVar() - time.Duration",
	}
	for _, item := range advancedFeatures {
		fmt.Println(item)
	}

	// Part 6: Best Practices
	fmt.Println()
	fmt.Println("✅ BEST PRACTICES")
	fmt.Println(strings.Repeat("-", 70))

	bestPractices := []string{
		"☐ Prefer flag package over os.Args for production code",
		"☐ Always provide default values for flags",
		"☐ Add clear descriptions for each flag (visible in -help)",
		"☐ Validate flag values immediately after flag.Parse()",
		"☐ Use descriptive flag names (-verbose not -v for clarity)",
		"☐ Sanitize/validate user input before using",
		"☐ Test with actual 'go build' executable, not 'go run'",
		"☐ Use flag.Args() for remaining positional arguments",
		"☐ Check len(flag.Args()) > 0 before accessing",
		"☐ Consider using external packages (cobra, urfave/cli) for complex CLIs",
	}
	for _, practice := range bestPractices {
		fmt.Println(practice)
	}

	// Part 7: os.Args vs flag Package Comparison
	fmt.Println()
	fmt.Println("🔀 COMPARISON: os.Args vs flag Package")
	fmt.Println(strings.Repeat("-", 70))

	comparison := `
os.Args Approach:
  Pros:  Simple, minimal dependencies
  Cons:  No type safety, no defaults, manual parsing, fragile

flag Package:
  Pros:  Type-safe, defaults, validation, help text, robust
  Cons:  Slightly more setup code (worth it!)

RECOMMENDATION: Use flag package for any real CLI tool!
`
	fmt.Println(comparison)
}
