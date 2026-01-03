package main

import (
	"fmt"
	"io/ioutil"
)

/*
REMAINING INTERMEDIATE TOPICS QUICK REFERENCE

This file references topics 84-93 which follow similar patterns:

84. READ_FILE - Read file contents (ioutil.ReadFile)
85. LINE_FILTERS - Process files line by line (bufio.Scanner)
86. FILE_PATHS - Path manipulation (filepath package)
87. DIRECTORIES - Directory operations (os.ReadDir, os.Mkdir)
88. TEMP_FILES - Create temporary files (ioutil.TempFile)
89. EMBED_DIRECTIVE - Embed files in binary (//go:embed)
90. CMD_ARGS - Command-line arguments (os.Args, flag package)
91. SUBCOMMANDS - Handle CLI subcommands
92. ENV_VARS - Environment variables (os.Getenv, os.LookupEnv)
93. LOGGING - Structured logging (log package)
*/

// ============================================================================
// EXAMPLE 1: File Reading Quick Reference
// ============================================================================

func quickRef84() {
	fmt.Println("\n=== 84. FILE READING ===")
	fmt.Println("📚 Read entire file: ioutil.ReadFile")
	fmt.Println("📚 Read line by line: bufio.Scanner")

	// Quick example
	if data, err := ioutil.ReadFile("example.txt"); err == nil {
		fmt.Printf("Read %d bytes\n", len(data))
	}
}

// ============================================================================
// EXAMPLE 2: Line Filters
// ============================================================================

func quickRef85() {
	fmt.Println("\n=== 85. LINE FILTERS ===")
	fmt.Println("📚 Process file line by line with bufio.Scanner")
	fmt.Println("  • Filter lines matching pattern")
	fmt.Println("  • Count specific lines")
	fmt.Println("  • Transform content")
}

// ============================================================================
// EXAMPLE 3: File Paths
// ============================================================================

func quickRef86() {
	fmt.Println("\n=== 86. FILE PATHS ===")
	fmt.Println("📚 Use filepath package for path manipulation")

	examples := []struct {
		func_name string
		example   string
	}{
		{"filepath.Join", "/home + user + documents = /home/user/documents"},
		{"filepath.Dir", "/home/user/file.txt = /home/user"},
		{"filepath.Base", "/home/user/file.txt = file.txt"},
		{"filepath.Ext", "/home/user/file.txt = .txt"},
		{"filepath.IsAbs", "Check if path is absolute"},
		{"filepath.Abs", "Get absolute path from relative"},
	}

	for _, ex := range examples {
		fmt.Printf("  %s: %s\n", ex.func_name, ex.example)
	}
}

// ============================================================================
// EXAMPLE 4: Directories
// ============================================================================

func quickRef87() {
	fmt.Println("\n=== 87. DIRECTORIES ===")
	fmt.Println("📚 Directory operations")

	fmt.Println("  os.ReadDir() - List directory contents")
	fmt.Println("  os.Mkdir() - Create single directory")
	fmt.Println("  os.MkdirAll() - Create directory tree")
	fmt.Println("  os.Remove() - Delete file or empty dir")
	fmt.Println("  os.RemoveAll() - Delete directory recursively")
}

// ============================================================================
// EXAMPLE 5: Temporary Files
// ============================================================================

func quickRef88() {
	fmt.Println("\n=== 88. TEMPORARY FILES ===")
	fmt.Println("📚 Create temp files safely")

	// f, _ := ioutil.TempFile("", "prefix_*.txt")
	// defer os.Remove(f.Name())
	// f.WriteString("temp content")

	fmt.Println("  ioutil.TempFile(\"\", \"prefix_*.txt\") - Create temp file")
	fmt.Println("  ioutil.TempDir(\"\", \"prefix_\") - Create temp directory")
	fmt.Println("  Always defer Remove() for cleanup")
}

// ============================================================================
// EXAMPLE 6: Embed Directive
// ============================================================================

func quickRef89() {
	fmt.Println("\n=== 89. EMBED DIRECTIVE ===")
	fmt.Println("📚 Embed files in Go binary (Go 1.16+)")

	fmt.Println(`
  //go:embed path/to/file
  var content []byte
  
  //go:embed path/to/directory/*
  var embedFS embed.FS
  
  Advantages:
  • Files included in binary
  • No external dependencies
  • Single executable deployment
  • Zero runtime cost (already compiled)
	`)
}

// ============================================================================
// EXAMPLE 7: Command-line Arguments
// ============================================================================

func quickRef90() {
	fmt.Println("\n=== 90. COMMAND-LINE ARGUMENTS ===")
	fmt.Println("📚 Parse command-line arguments")

	fmt.Println("  os.Args[0] - Program name")
	fmt.Println("  os.Args[1:] - Arguments")
	fmt.Println("  flag.String() - Define string flag")
	fmt.Println("  flag.Bool() - Define bool flag")
	fmt.Println("  flag.Parse() - Parse arguments")
	fmt.Println("")
	fmt.Println("Example:")
	fmt.Println("  $ myapp -name Alice -verbose")
	fmt.Println("  name := flag.String(\"name\", \"Guest\", \"Your name\")")
}

// ============================================================================
// EXAMPLE 8: Subcommands
// ============================================================================

func quickRef91() {
	fmt.Println("\n=== 91. SUBCOMMANDS ===")
	fmt.Println("📚 CLI tools with subcommands (like git)")

	fmt.Println(`
  git commit -m "message"
  git push origin main
  git pull
  
  Pattern:
  1. Check first argument (subcommand)
  2. Create flag.FlagSet for each subcommand
  3. Parse remaining arguments
  4. Execute appropriate subcommand
	`)
}

// ============================================================================
// EXAMPLE 9: Environment Variables
// ============================================================================

func quickRef92() {
	fmt.Println("\n=== 92. ENVIRONMENT VARIABLES ===")
	fmt.Println("📚 Read environment variables")

	fmt.Println("  os.Getenv(key) - Get variable (empty if not set)")
	fmt.Println("  os.LookupEnv(key) - Get with bool (ok=false if not set)")
	fmt.Println("  os.Setenv(key, value) - Set variable")
	fmt.Println("  os.Environ() - Get all variables")
	fmt.Println("")
	fmt.Println("Example:")
	fmt.Println("  dbURL := os.Getenv(\"DATABASE_URL\")")
	fmt.Println("  apiKey, ok := os.LookupEnv(\"API_KEY\")")
}

// ============================================================================
// EXAMPLE 10: Logging
// ============================================================================

func quickRef93() {
	fmt.Println("\n=== 93. LOGGING ===")
	fmt.Println("📚 Built-in logging (log package)")

	fmt.Println("  log.Print() - Simple logging")
	fmt.Println("  log.Printf() - Formatted logging")
	fmt.Println("  log.Fatal() - Log and exit")
	fmt.Println("  log.Panic() - Log and panic")
	fmt.Println("  log.SetOutput() - Change output destination")
	fmt.Println("")
	fmt.Println("For production:")
	fmt.Println("  • Use structured logging (slog, logrus, zap)")
	fmt.Println("  • Include context (user ID, request ID)")
	fmt.Println("  • Use log levels (DEBUG, INFO, WARN, ERROR)")
}

// ============================================================================
// TOPICS 95-102 SUMMARY
// ============================================================================

func summaryTopics95to102() {
	fmt.Println("\n╔════════════════════════════════════════╗")
	fmt.Println("║       TOPICS 95-102 SUMMARY           ║")
	fmt.Println("╚════════════════════════════════════════╝")

	topics := map[string]string{
		"95. STRUCT TAGS":      "JSON tags, encoding, metadata: `json:\"name\" xml:\"name\"`",
		"96. XML":              "xml.Marshal/Unmarshal: encoding/xml package",
		"97. TYPE CONVERSIONS": "strconv package, type assertion, casting",
		"98. IO PACKAGE":       "Reader/Writer interfaces: io.Copy, io.ReadAll",
		"99. MATH PACKAGE":     "Math functions: math.Sqrt, math.Sin, math.Max",
		"100. MATH EXAMPLES":   "Practical math operations and algorithms",
		"101. SECTION SUMMARY": "Review key concepts from 57-100",
		"102. PROJECT IDEAS":   "Build applications using intermediate concepts",
	}

	for topic, description := range topics {
		fmt.Printf("\n%s:\n  %s\n", topic, description)
	}
}

// ============================================================================
// MAIN
// ============================================================================

func main() {
	fmt.Println("╔════════════════════════════════════════╗")
	fmt.Println("║  INTERMEDIATE TOPICS 84-102 REFERENCE ║")
	fmt.Println("╚════════════════════════════════════════╝")

	quickRef84()
	quickRef85()
	quickRef86()
	quickRef87()
	quickRef88()
	quickRef89()
	quickRef90()
	quickRef91()
	quickRef92()
	quickRef93()

	summaryTopics95to102()

	fmt.Println("\n╔════════════════════════════════════════╗")
	fmt.Println("║     CREATING FULL TEACHING FILES      ║")
	fmt.Println("╚════════════════════════════════════════╝")
	fmt.Println(`
For complete teaching materials, create individual files:

For each topic 84-93, create a file following the pattern:
  • 10 detailed examples
  • Progressive difficulty
  • Practical use cases
  • Error handling
  • Best practices
  • Teaching checkpoints

Topics to expand:
  84_read_file_detailed.go
  85_line_filters_detailed.go
  86_file_paths_detailed.go
  87_directories_detailed.go
  88_temp_files_dirs_detailed.go
  89_embed_directive_detailed.go
  90_cmd_args_flags_detailed.go
  91_subcommands_detailed.go
  92_env_vars_detailed.go
  93_logging_detailed.go
  
Plus summary and projects:
  95_struct_tags_detailed.go
  96_xml_detailed.go
  97_type_conversions_detailed.go
  98_io_package_detailed.go
  99_math_package_detailed.go
  100_math_examples_detailed.go
  101_section_summary.go
  102_project_ideas.go
	`)
}
