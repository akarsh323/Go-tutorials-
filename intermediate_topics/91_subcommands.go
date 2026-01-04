package main

import (
	"flag"
	"fmt"
	"os"
)

// ============================================================================
// 91 Command Line Arguments, Flags & Subcommands
// ============================================================================
//
// Go provides TWO MAIN approaches to handle command line input:
// 1. RAW ARGUMENTS (os.Args) - Direct access to all input
// 2. FLAGS (flag package) - Named parameters with types and defaults
// 3. SUBCOMMANDS - Command routing (e.g., git commit, git push)
//
// ============================================================================

// ============================================================================
// PART 1: RAW ARGUMENTS (os.Args)
// ============================================================================

func Demo91_Part1_RawArguments() {
	fmt.Println("\n=== PART 1: RAW ARGUMENTS (os.Args) ===")
	fmt.Println("os.Args is a slice containing all command line input.")
	fmt.Println()

	// Example 1: Access all arguments
	fmt.Println("📌 All Arguments:")
	fmt.Println("os.Args contains all input including the program name.")
	for i, arg := range os.Args {
		fmt.Printf("  os.Args[%d]: %v\n", i, arg)
	}
	fmt.Println()

	// Example 2: Understanding os.Args[0]
	fmt.Println("📌 os.Args[0] - The Program Name/Path:")
	fmt.Printf("  When running 'go run file.go', os.Args[0] is a temporary path.\n")
	fmt.Printf("  Actual value: %s\n", os.Args[0])
	fmt.Println()

	// Example 3: Safe access to arguments
	fmt.Println("📌 Safe Access (checking length):")
	if len(os.Args) > 1 {
		fmt.Printf("  First argument: %s\n", os.Args[1])
	} else {
		fmt.Println("  No arguments provided!")
	}
	fmt.Println()
}

// ============================================================================
// PART 2: FLAGS (flag Package)
// ============================================================================

func Demo91_Part2_Flags() {
	fmt.Println("\n=== PART 2: FLAGS (flag Package) ===")
	fmt.Println("Flags provide named parameters with types, defaults, and descriptions.")
	fmt.Println()

	// Example 1: Simple flags
	exampleSimpleFlags()

	// Example 2: Flag types
	exampleFlagTypes()

	// Example 3: Boolean flags (short hand)
	exampleBooleanFlags()

	// Example 4: Positional arguments after flags
	examplePositionalArgs()
}

func exampleSimpleFlags() {
	fmt.Println("📌 Simple String and Int Flags:")
	fmt.Println("   Usage: go run file.go -name James -age 30")
	fmt.Println()

	// Create a new flag set for this example
	fs := flag.NewFlagSet("simple", flag.ContinueOnError)
	var name string
	var age int

	fs.StringVar(&name, "name", "John", "Name of the user")
	fs.IntVar(&age, "age", 18, "Age of the user")

	// Simulate command: go run file.go -name James -age 30
	args := []string{"-name", "James", "-age", "30"}
	fs.Parse(args)

	fmt.Printf("   Result: Name=%s, Age=%d\n", name, age)
	fmt.Println()
}

func exampleFlagTypes() {
	fmt.Println("📌 Different Flag Types (String, Int, Bool, Float64):")
	fmt.Println("   Usage: go run file.go -name Alex -score 95.5 -active=true")
	fmt.Println()

	fs := flag.NewFlagSet("types", flag.ContinueOnError)
	var name string
	var score float64
	var active bool
	var count int

	fs.StringVar(&name, "name", "Guest", "Name")
	fs.Float64Var(&score, "score", 0.0, "Score")
	fs.BoolVar(&active, "active", false, "Is active?")
	fs.IntVar(&count, "count", 1, "Count")

	args := []string{"-name", "Alex", "-score", "95.5", "-active=true", "-count", "5"}
	fs.Parse(args)

	fmt.Printf("   String: name=%s\n", name)
	fmt.Printf("   Float:  score=%.1f\n", score)
	fmt.Printf("   Bool:   active=%v\n", active)
	fmt.Printf("   Int:    count=%d\n", count)
	fmt.Println()
}

func exampleBooleanFlags() {
	fmt.Println("📌 Boolean Flags (Multiple Syntax Options):")
	fmt.Println("   All of these are equivalent:")
	fmt.Println("   - go run file.go -verbose=true")
	fmt.Println("   - go run file.go -verbose")
	fmt.Println("   - go run file.go -v (short form, needs separate definition)")
	fmt.Println()

	fs := flag.NewFlagSet("bool", flag.ContinueOnError)
	var verbose bool

	fs.BoolVar(&verbose, "verbose", false, "Enable verbose output")

	// Test different boolean flag formats
	fmt.Println("   Testing different formats:")
	for _, args := range [][]string{
		{"-verbose=true"},
		{"-verbose"},
		{"-verbose=false"},
	} {
		fs.Parse(args)
		fmt.Printf("   Input: %v => verbose=%v\n", args, verbose)
	}
	fmt.Println()
}

func examplePositionalArgs() {
	fmt.Println("📌 Positional Arguments (After Flags):")
	fmt.Println("   Usage: go run file.go -name James file1.txt file2.txt")
	fmt.Println()

	fs := flag.NewFlagSet("positional", flag.ContinueOnError)
	var name string

	fs.StringVar(&name, "name", "John", "Name")

	// Simulate: -name James file1.txt file2.txt
	args := []string{"-name", "James", "file1.txt", "file2.txt"}
	fs.Parse(args)

	fmt.Printf("   Flag: name=%s\n", name)
	fmt.Printf("   Positional args (fs.Args()): %v\n", fs.Args())
	fmt.Println()
}

// ============================================================================
// PART 3: SUBCOMMANDS
// ============================================================================

func Demo91_Part3_Subcommands() {
	fmt.Println("\n=== PART 3: SUBCOMMANDS ===")
	fmt.Println("Subcommands allow command routing like: git commit, git push, git clone")
	fmt.Println()

	exampleSimpleSubcommands()
	exampleSubcommandsWithFlags()
}

func exampleSimpleSubcommands() {
	fmt.Println("📌 Simple Subcommand Routing:")
	fmt.Println("   Usage: go run file.go create filename.txt")
	fmt.Println("          go run file.go delete filename.txt")
	fmt.Println()

	// Simulate commands
	commands := [][]string{
		{"create", "myfile.txt"},
		{"delete", "oldfile.txt"},
		{"list"},
	}

	for _, cmdArgs := range commands {
		fmt.Printf("   Command: %v\n", cmdArgs)

		if len(cmdArgs) == 0 {
			fmt.Println("   → No command provided")
			continue
		}

		switch cmdArgs[0] {
		case "create":
			if len(cmdArgs) > 1 {
				fmt.Printf("   → Creating file: %s\n", cmdArgs[1])
			}
		case "delete":
			if len(cmdArgs) > 1 {
				fmt.Printf("   → Deleting file: %s\n", cmdArgs[1])
			}
		case "list":
			fmt.Println("   → Listing all files")
		default:
			fmt.Printf("   → Unknown command: %s\n", cmdArgs[0])
		}
	}
	fmt.Println()
}

func exampleSubcommandsWithFlags() {
	fmt.Println("📌 Subcommands with Flags:")
	fmt.Println("   Usage: go run file.go deploy -environment production -force")
	fmt.Println("          go run file.go build -output bin/app -verbose")
	fmt.Println()

	// Example: deploy command with flags
	deployArgs := []string{"-environment", "production", "-force"}
	deployFS := flag.NewFlagSet("deploy", flag.ContinueOnError)
	var env string
	var force bool
	deployFS.StringVar(&env, "environment", "staging", "Deployment environment")
	deployFS.BoolVar(&force, "force", false, "Force deployment")
	deployFS.Parse(deployArgs)

	fmt.Println("   Deploy command:")
	fmt.Printf("   → Environment: %s\n", env)
	fmt.Printf("   → Force: %v\n", force)
	fmt.Println()

	// Example: build command with flags
	buildArgs := []string{"-output", "bin/app", "-verbose"}
	buildFS := flag.NewFlagSet("build", flag.ContinueOnError)
	var output string
	var verbose bool
	buildFS.StringVar(&output, "output", "a.out", "Output file")
	buildFS.BoolVar(&verbose, "verbose", false, "Verbose output")
	buildFS.Parse(buildArgs)

	fmt.Println("   Build command:")
	fmt.Printf("   → Output: %s\n", output)
	fmt.Printf("   → Verbose: %v\n", verbose)
	fmt.Println()
}

// ============================================================================
// PART 4: IMPORTANT CONCEPTS & BEST PRACTICES
// ============================================================================

func Demo91_Part4_BestPractices() {
	fmt.Println("\n=== PART 4: KEY CONCEPTS & BEST PRACTICES ===")
	fmt.Println()

	fmt.Println("📌 Concept 1: os.Args[0] is NOT 'main.go'")
	fmt.Println("   When running 'go run main.go', Go creates a temporary executable.")
	fmt.Println("   os.Args[0] will be something like: /tmp/go-build123/exe/a.out")
	fmt.Println("   NOT: main.go (unless you compile to a binary)")
	fmt.Println()

	fmt.Println("📌 Concept 2: Spaces in Arguments Require Quotes")
	fmt.Println("   Input:  go run file.go -name James Doe")
	fmt.Println("   Result: name='James', 'Doe' treated as separate argument")
	fmt.Println()
	fmt.Println("   Input:  go run file.go -name \"James Doe\"")
	fmt.Println("   Result: name='James Doe' (quotes prevent splitting)")
	fmt.Println()

	fmt.Println("📌 Concept 3: Flag Parsing Order")
	fmt.Println("   - Flags must come BEFORE positional arguments")
	fmt.Println("   - After flag.Parse(), remaining args are in fs.Args()")
	fmt.Println("   - flag.Parse() stops at first non-flag argument")
	fmt.Println()

	fmt.Println("📌 Concept 4: Pointer (&) is Required with FlagVar")
	fmt.Println("   ✗ WRONG:  flag.StringVar(name, \"name\", \"default\", \"...\")")
	fmt.Println("   ✓ RIGHT:  flag.StringVar(&name, \"name\", \"default\", \"...\")")
	fmt.Println("   Reason: flag package needs to modify the variable's value")
	fmt.Println()

	fmt.Println("📌 Concept 5: Default Values and Validation")
	fmt.Println("   - Flags have default values that apply if not provided")
	fmt.Println("   - Always validate/sanitize user input before using it")
	fmt.Println("   - Check os.Args length before accessing indices")
	fmt.Println()
}

// ============================================================================
// PART 5: COMPLETE EXAMPLE - Simple CLI Tool
// ============================================================================

func Demo91_Part5_CompleteExample() {
	fmt.Println("\n=== PART 5: COMPLETE CLI TOOL EXAMPLE ===")
	fmt.Println()

	fmt.Println("📌 File Manager CLI Tool Simulation")
	fmt.Println("   This demonstrates a real-world pattern for CLI tools.")
	fmt.Println()

	// Simulate: filemanager -action delete -path /tmp/test.txt -force
	args := []string{"-action", "delete", "-path", "/tmp/test.txt", "-force"}

	fs := flag.NewFlagSet("filemanager", flag.ContinueOnError)
	var action string
	var path string
	var force bool

	fs.StringVar(&action, "action", "list", "Action: create, delete, list, info")
	fs.StringVar(&path, "path", "", "File path to operate on")
	fs.BoolVar(&force, "force", false, "Force operation without confirmation")

	fs.Parse(args)

	fmt.Printf("   Input: %v\n", args)
	fmt.Println()
	fmt.Printf("   Parsed values:\n")
	fmt.Printf("   - action: %s\n", action)
	fmt.Printf("   - path:   %s\n", path)
	fmt.Printf("   - force:  %v\n", force)
	fmt.Println()

	// Execute based on action
	fmt.Println("   Execution:")
	switch action {
	case "delete":
		if path == "" {
			fmt.Println("   ✗ Error: path is required for delete")
		} else if force {
			fmt.Printf("   ✓ Force deleting: %s\n", path)
		} else {
			fmt.Printf("   ? Confirm delete: %s (y/n)\n", path)
		}
	case "create":
		if path == "" {
			fmt.Println("   ✗ Error: path is required for create")
		} else {
			fmt.Printf("   ✓ Creating file: %s\n", path)
		}
	case "list":
		fmt.Println("   ✓ Listing files in current directory")
	default:
		fmt.Printf("   ✗ Unknown action: %s\n", action)
	}
	fmt.Println()
}

// ============================================================================
// MAIN DEMO FUNCTION
// ============================================================================

func Demo91SubCommands() {
	fmt.Println("-- 91 Command Line Arguments, Flags & Subcommands --")
	fmt.Println()
	fmt.Println("This guide covers how to handle user input from the command line.")
	fmt.Println()

	Demo91_Part1_RawArguments()
	Demo91_Part2_Flags()
	Demo91_Part3_Subcommands()
	Demo91_Part4_BestPractices()
	Demo91_Part5_CompleteExample()

	fmt.Println("\n=== SUMMARY ===")
	fmt.Println("✓ os.Args: Raw access to all command line input")
	fmt.Println("✓ flag package: Type-safe named parameters with defaults")
	fmt.Println("✓ Subcommands: Route commands to different handlers")
	fmt.Println("✓ Always validate input and check argument lengths")
	fmt.Println("✓ Use quotes for arguments with spaces")
}
