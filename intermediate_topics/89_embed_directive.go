package main

import (
	"fmt"
	"strings"
)

// ============================================================================
// 89 EMBED DIRECTIVE
// ============================================================================
// Introduced in Go 1.16, the embed package allows you to include static files
// and directories directly into your Go binary at compile time.
//
// Key Benefits:
// 1. Simplicity & Deployment: Single standalone executable (no external files)
// 2. Security: Files are baked into binary, harder to tamper with
// 3. Efficiency: No external filesystem dependencies at runtime
//
// IMPORTANT RULES:
// - The //go:embed directive must start at package level (before var declarations)
// - Format: //go:embed <file_or_folder_name>
// - Must be immediately followed by a var declaration
// - For string/[]byte: use blank import (_ "embed")
// - For embed.FS: use regular import (embed is used explicitly)
// ============================================================================

// Example 1: Embedding a Single File into a String
// Note: This example shows the syntax. In practice, you'd have actual files.
// To use this: create a file named "example.txt" in the same directory
// Uncomment the code below when you have the file:
/*
//go:embed example.txt
var exampleContent string

func Demo89EmbedString() {
	fmt.Println("-- Example 1: Embedding String Content --")
	fmt.Println("Content:", exampleContent)
}
*/

// Example 2: Embedding a Directory into embed.FS
// To use this: create a folder named "assets" with files inside
// Uncomment the code below when you have the folder:
/*
//go:embed assets
var assets embed.FS

func Demo89EmbedDirectory() {
	fmt.Println("-- Example 2: Embedding Directory (embed.FS) --")

	// Read a specific file from embedded directory
	data, err := assets.ReadFile("assets/config.txt")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}
	fmt.Println("Config content:", string(data))
}
*/

// Example 3: Walking Through Embedded Files
// This demonstrates how to traverse all files in an embedded directory
// Uncomment and use with an actual embedded directory:
/*
func Demo89WalkEmbedded() {
	fmt.Println("-- Example 3: Walking Embedded Directory --")

	fs.WalkDir(assets, "assets", func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if !d.IsDir() {
			fmt.Println("Found embedded file:", path)
		}
		return nil
	})
}
*/

// DEMO: Syntax Examples (Instructional)
func Demo89EmbedDirective() {
	fmt.Println("-- 89 Embed Directive --")
	fmt.Println()

	// Syntax Overview
	fmt.Println("📋 SYNTAX EXAMPLES:")
	fmt.Println()

	example1 := `// Embed single file as string (requires blank import):
import (
    _ "embed"
    "fmt"
)

//go:embed example.txt
var content string

func main() {
    fmt.Println(content) // File content is loaded at compile time
}
`
	fmt.Println(example1)

	example2 := `// Embed directory (uses embed.FS, regular import):
import (
    "embed"
    "fmt"
)

//go:embed assets
var assets embed.FS

func main() {
    data, _ := assets.ReadFile("assets/config.txt")
    fmt.Println(string(data))
}
`
	fmt.Println(example2)

	example3 := `// Walk through embedded directory:
import (
    "embed"
    "io/fs"
)

//go:embed assets
var assets embed.FS

func main() {
    fs.WalkDir(assets, "assets", func(path string, d fs.DirEntry, err error) error {
        if err == nil && !d.IsDir() {
            fmt.Println("File:", path)
        }
        return nil
    })
}
`
	fmt.Println(example3)

	// Key Considerations
	fmt.Println()
	fmt.Println("🔑 KEY CONSIDERATIONS:")
	fmt.Println(strings.Repeat("-", 70))
	considerations := []string{
		"✓ Read-Only: Embedded files are immutable at runtime",
		"✓ Binary Size: Large files increase executable size significantly",
		"✓ Compile Time: Embedding happens at build time, not runtime",
		"✓ Blank Import: Required for string/[]byte, not for embed.FS",
		"✓ Pattern Matching: Can use wildcards like //go:embed static/*",
		"✓ Recursive: //go:embed dir/* embeds folder and all subfolders",
	}
	for _, item := range considerations {
		fmt.Println(item)
	}

	fmt.Println()
	fmt.Println("💡 PRACTICAL USE CASES:")
	fmt.Println(strings.Repeat("-", 70))
	useCases := []string{
		"• Web servers: Serve HTML/CSS/JS files without external filesystem",
		"• Configuration: Embed default config files in binary",
		"• Templates: Package text/html templates with the application",
		"• Static assets: Include images, JSON data, or documentation",
		"• Standalone CLI tools: No missing files on user's machine",
	}
	for _, useCase := range useCases {
		fmt.Println(useCase)
	}

	fmt.Println()
	fmt.Println("⚠️  REMEMBER:")
	fmt.Println(strings.Repeat("-", 70))
	fmt.Println("To use embed in your code:")
	fmt.Println("1. Create actual files/folders to embed")
	fmt.Println("2. Add //go:embed directive before var declaration")
	fmt.Println("3. Use 'go build' (compile time) to embed files into binary")

	fmt.Println()
	fmt.Println("🌐 PRACTICAL EXAMPLE: Web Server with Embedded Static Files")
	fmt.Println(strings.Repeat("-", 70))

	webServerExample := `// To serve HTML/CSS from a Go web server using embed:

package main

import (
	"embed"
	"log"
	"net/http"
)

// Create a 'static' folder with index.html, style.css, app.js files
// Then embed the entire folder:
//go:embed static
var staticFiles embed.FS

func main() {
	// Serve embedded files on /static/ path
	fs := http.FileServer(http.FS(staticFiles))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	// Serve HTML from root
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		data, _ := staticFiles.ReadFile("static/index.html")
		w.Header().Set("Content-Type", "text/html")
		w.Write(data)
	})

	log.Println("Server running on :8080")
	http.ListenAndServe(":8080", nil)
}

// Directory structure:
// project/
//   main.go
//   static/
//     index.html
//     style.css
//     app.js
//
// Build: go build -o myapp main.go
// Run:   ./myapp
// Result: Single executable file containing all HTML/CSS/JS!
`
	fmt.Println(webServerExample)

	fmt.Println()
	fmt.Println("📦 ADVANCED: Pattern Matching with Embed")
	fmt.Println(strings.Repeat("-", 70))

	patternExample := `// Embed multiple patterns:

//go:embed static/html/*.html
var htmlFiles embed.FS

//go:embed static/css/*.css
var cssFiles embed.FS

//go:embed static templates/*
var allStaticFiles embed.FS

// Wildcards:
// *.ext         - All files with extension in current folder
// */*.ext       - Files in subfolders
// dir/*         - All contents of dir (recursive)
// dir/**/*.ext  - Nested pattern matching
`
	fmt.Println(patternExample)

	fmt.Println()
	fmt.Println("✅ COMPLETE CHECKLIST:")
	fmt.Println(strings.Repeat("-", 70))
	checklist := []string{
		"☐ Understand embed is compile-time, not runtime",
		"☐ Know when to use blank import (_ \"embed\")",
		"☐ Know when to use regular import (embed.FS)",
		"☐ Remember embedded files are read-only",
		"☐ Use fs.FS interface for flexibility",
		"☐ Consider binary size for large files",
		"☐ Use http.FS for web server integration",
		"☐ Test with 'go build' not 'go run' for actual embedding",
	}
	for _, item := range checklist {
		fmt.Println(item)
	}
}
