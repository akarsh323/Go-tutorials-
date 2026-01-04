package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

/*
═══════════════════════════════════════════════════════════════════════════════
                         FILE PATHS IN GO
═══════════════════════════════════════════════════════════════════════════════

The path/filepath package is essential for handling file system interactions
reliably across different operating systems (Windows vs. Unix/Linux/Mac).

KEY PROBLEM SOLVED:
  • Unix systems use forward slashes (/)
  • Windows uses backslashes (\)
  • Hardcoding separators causes bugs when moving code between OSs

SOLUTION: filepath package automatically detects the OS and applies the
correct separator, ensuring PORTABLE code.

═══════════════════════════════════════════════════════════════════════════════
                    CORE CONCEPTS: ABSOLUTE vs. RELATIVE PATHS
═══════════════════════════════════════════════════════════════════════════════

ABSOLUTE PATH: Specifies complete location from the root
  • Windows: C:\Users\Akarsh\Docs
  • Linux/Mac: /home/user/docs
  → Can be used from any current working directory

RELATIVE PATH: Specifies location relative to Current Working Directory (CWD)
  • . (dot):   Current directory
  • .. (dots):  Parent directory (one level up)
  • file.txt:   File in current directory
  • subdir/file.txt: File in subdirectory
  → Only works correctly if CWD is expected location

═══════════════════════════════════════════════════════════════════════════════
*/

/*
━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━
  SECTION 1: CONSTRUCTING & NORMALIZING PATHS
━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━
*/

func Example1_JoinAndClean() {
	fmt.Println("\n" + strings.Repeat("═", 80))
	fmt.Println("EXAMPLE 1: Constructing & Normalizing Paths")
	fmt.Println(strings.Repeat("═", 80) + "\n")

	// ─────────────────────────────────────────────────────────────────────────
	// filepath.Join(): Join path elements with correct OS separator
	// ─────────────────────────────────────────────────────────────────────────
	fmt.Println("📌 filepath.Join() - Combine path elements")
	fmt.Println(strings.Repeat("─", 80))

	path1 := filepath.Join("home", "user", "docs", "file.txt")
	fmt.Printf("Join('home', 'user', 'docs', 'file.txt')\n")
	fmt.Printf("  Result: %q\n\n", path1)

	// Real-world example: Build config file path
	homeDir := os.Getenv("HOME") // Get home directory
	if homeDir != "" {
		configPath := filepath.Join(homeDir, ".config", "myapp", "config.json")
		fmt.Printf("Config path: %q\n", configPath)
	} else {
		fmt.Println("HOME env var not set")
	}

	fmt.Println()

	// ─────────────────────────────────────────────────────────────────────────
	// filepath.Clean(): Normalize and simplify paths
	// ─────────────────────────────────────────────────────────────────────────
	fmt.Println("📌 filepath.Clean() - Simplify and normalize paths")
	fmt.Println(strings.Repeat("─", 80))

	messyPath1 := "data/../data/file.txt" // Goes out, then back in
	cleanedPath1 := filepath.Clean(messyPath1)
	fmt.Printf("Input:  %q\n", messyPath1)
	fmt.Printf("Output: %q\n\n", cleanedPath1)

	messyPath2 := "./documents//folder///file.txt" // Multiple slashes and dot
	cleanedPath2 := filepath.Clean(messyPath2)
	fmt.Printf("Input:  %q\n", messyPath2)
	fmt.Printf("Output: %q\n\n", cleanedPath2)

	messyPath3 := "a/b/c/../../d/file.txt" // Complex navigation
	cleanedPath3 := filepath.Clean(messyPath3)
	fmt.Printf("Input:  %q\n", messyPath3)
	fmt.Printf("Output: %q\n", cleanedPath3)
}

/*
━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━
  SECTION 2: DECONSTRUCTING PATHS (SPLITTING)
━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━

When working with file paths, you often need to extract components:
  • The directory containing the file
  • The filename itself
  • The file extension

These functions make deconstruction easy and portable.
━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━
*/

func Example2_DeconstructingPaths() {
	fmt.Println("\n" + strings.Repeat("═", 80))
	fmt.Println("EXAMPLE 2: Deconstructing Paths (Splitting)")
	fmt.Println(strings.Repeat("═", 80) + "\n")

	path := "/users/akarsh/documents/report_2025.pdf"

	// ─────────────────────────────────────────────────────────────────────────
	// filepath.Split(): Separate directory from filename
	// ─────────────────────────────────────────────────────────────────────────
	fmt.Println("📌 filepath.Split() - Separate directory and filename")
	fmt.Println(strings.Repeat("─", 80))

	dir, file := filepath.Split(path)
	fmt.Printf("Path:      %q\n", path)
	fmt.Printf("Directory: %q\n", dir)
	fmt.Printf("Filename:  %q\n\n", file)

	// ─────────────────────────────────────────────────────────────────────────
	// filepath.Dir(): Get only the directory
	// ─────────────────────────────────────────────────────────────────────────
	fmt.Println("📌 filepath.Dir() - Extract directory path only")
	fmt.Println(strings.Repeat("─", 80))

	dirOnly := filepath.Dir(path)
	fmt.Printf("Path:          %q\n", path)
	fmt.Printf("Directory:     %q\n\n", dirOnly)

	// ─────────────────────────────────────────────────────────────────────────
	// filepath.Base(): Get only the filename
	// ─────────────────────────────────────────────────────────────────────────
	fmt.Println("📌 filepath.Base() - Extract filename only")
	fmt.Println(strings.Repeat("─", 80))

	filename := filepath.Base(path)
	fmt.Printf("Path:     %q\n", path)
	fmt.Printf("Filename: %q\n\n", filename)

	// ─────────────────────────────────────────────────────────────────────────
	// filepath.Ext(): Get the file extension
	// ─────────────────────────────────────────────────────────────────────────
	fmt.Println("📌 filepath.Ext() - Extract file extension")
	fmt.Println(strings.Repeat("─", 80))

	extension := filepath.Ext(path)
	fmt.Printf("Path:      %q\n", path)
	fmt.Printf("Extension: %q\n", extension)
}

/*
━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━
  SECTION 3: EXTENSION HANDLING (AUDIO CONVERSION PATTERN)
━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━

SCENARIO: You receive a .wav audio file, encode it to .mp3, and need to
save it with the same name but different extension.

PATTERN:
  1. Get the filename and extension
  2. Remove the extension
  3. Append the new extension
  4. Save with the new path

This demonstrates a practical, real-world use case for extension handling.
━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━
*/

func Example3_ExtensionHandling() {
	fmt.Println("\n" + strings.Repeat("═", 80))
	fmt.Println("EXAMPLE 3: Extension Handling (Audio Conversion Pattern)")
	fmt.Println(strings.Repeat("═", 80) + "\n")

	wavFile := "recordings/podcast_episode_42.wav"

	fmt.Println("SCENARIO: Convert .wav audio to .mp3")
	fmt.Println(strings.Repeat("─", 80))
	fmt.Printf("Original file: %q\n\n", wavFile)

	// Step 1: Get extension
	extension := filepath.Ext(wavFile)
	fmt.Printf("Step 1 - Get extension:     %q\n", extension)

	// Step 2: Remove extension to get base name
	nameOnly := strings.TrimSuffix(wavFile, extension)
	fmt.Printf("Step 2 - Remove extension:  %q\n", nameOnly)

	// Step 3: Add new extension
	mp3File := nameOnly + ".mp3"
	fmt.Printf("Step 3 - Add .mp3 ext:      %q\n\n", mp3File)

	fmt.Printf("✓ Converted path: %q\n\n", mp3File)

	// ─────────────────────────────────────────────────────────────────────────
	// More complex example: Handle files with multiple extensions
	// ─────────────────────────────────────────────────────────────────────────

	fmt.Println("EDGE CASE: File with multiple extensions")
	fmt.Println(strings.Repeat("─", 80))

	archiveFile := "backup.tar.gz"
	fmt.Printf("Original: %q\n", archiveFile)

	ext1 := filepath.Ext(archiveFile) // ".gz"
	fmt.Printf("First extension:  %q\n", ext1)

	nameBeforeGz := strings.TrimSuffix(archiveFile, ext1) // "backup.tar"
	ext2 := filepath.Ext(nameBeforeGz)                    // ".tar"
	fmt.Printf("Second extension: %q\n", ext2)

	fullName := strings.TrimSuffix(nameBeforeGz, ext2) // "backup"
	fmt.Printf("File name:        %q\n", fullName)
}

/*
━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━
  SECTION 4: ADVANCED NAVIGATION (Abs and Rel)
━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━

Sometimes you need to:
  • Convert a relative path to absolute (Abs)
  • Calculate the relative path between two locations (Rel)

These functions handle complex path calculations without manual string work.
━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━
*/

func Example4_AdvancedNavigation() {
	fmt.Println("\n" + strings.Repeat("═", 80))
	fmt.Println("EXAMPLE 4: Advanced Navigation (Abs and Rel)")
	fmt.Println(strings.Repeat("═", 80) + "\n")

	// ─────────────────────────────────────────────────────────────────────────
	// filepath.Abs(): Convert relative to absolute path
	// ─────────────────────────────────────────────────────────────────────────
	fmt.Println("📌 filepath.Abs() - Convert relative path to absolute")
	fmt.Println(strings.Repeat("─", 80))

	relativePath := "./data/config.json"
	fmt.Printf("Relative path: %q\n", relativePath)

	absolutePath, err := filepath.Abs(relativePath)
	if err != nil {
		fmt.Printf("ERROR: %v\n", err)
	} else {
		fmt.Printf("Absolute path: %q\n\n", absolutePath)
	}

	// ─────────────────────────────────────────────────────────────────────────
	// filepath.Rel(): Calculate relative path between two locations
	// ─────────────────────────────────────────────────────────────────────────
	fmt.Println("📌 filepath.Rel() - Calculate relative path between locations")
	fmt.Println(strings.Repeat("─", 80))

	// Scenario: You are in /home/user/projects/app and want to reference /home/user/docs/readme.txt
	basePath := "/home/user/projects/app"
	targetPath := "/home/user/docs/readme.txt"

	fmt.Printf("Current location (base):  %q\n", basePath)
	fmt.Printf("Target file location:     %q\n", targetPath)

	relPath, err := filepath.Rel(basePath, targetPath)
	if err != nil {
		fmt.Printf("ERROR: %v\n", err)
	} else {
		fmt.Printf("Relative path needed:     %q\n", relPath)
		fmt.Printf("(Go up 2 levels: .., then into docs, access readme.txt)\n\n")
	}

	// Another example within the same directory tree
	fmt.Println("ANOTHER EXAMPLE:")
	base2 := "/var/www/html"
	target2 := "/var/www/assets/images/logo.png"

	relPath2, err := filepath.Rel(base2, target2)
	if err != nil {
		fmt.Printf("ERROR: %v\n", err)
	} else {
		fmt.Printf("From:     %q\n", base2)
		fmt.Printf("To:       %q\n", target2)
		fmt.Printf("Relative: %q\n", relPath2)
	}
}

/*
━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━
  SECTION 5: FILEPROCESSOR INTERFACE PATTERN
━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━

Challenge: Create a FileProcessor interface with implementations for:
  • MP3Converter: Changes file extension to .mp3
  • BackupMover: Moves file to /backup directory

This combines interfaces (from earlier lessons) with path manipulation.
━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━
*/

// FileProcessor interface defines how different file operations transform paths
type FileProcessor interface {
	// GetNewPath takes the original file path and returns the transformed path
	GetNewPath(originalPath string) string
	// GetDescription returns a human-readable description of this processor
	GetDescription() string
}

// ─────────────────────────────────────────────────────────────────────────────
// MP3Converter: Changes file extension to .mp3
// ─────────────────────────────────────────────────────────────────────────────

type MP3Converter struct {
	targetExtension string
}

func (m *MP3Converter) GetNewPath(originalPath string) string {
	// Get the file extension
	ext := filepath.Ext(originalPath)

	// Remove the original extension
	nameOnly := strings.TrimSuffix(originalPath, ext)

	// Add .mp3 extension
	return nameOnly + m.targetExtension
}

func (m *MP3Converter) GetDescription() string {
	return fmt.Sprintf("Audio Converter (converts to %s)", m.targetExtension)
}

// ─────────────────────────────────────────────────────────────────────────────
// BackupMover: Moves file to a /backup directory
// ─────────────────────────────────────────────────────────────────────────────

type BackupMover struct {
	backupDir string
}

func (b *BackupMover) GetNewPath(originalPath string) string {
	// Extract just the filename (without directory path)
	filename := filepath.Base(originalPath)

	// Join with backup directory
	return filepath.Join(b.backupDir, filename)
}

func (b *BackupMover) GetDescription() string {
	return fmt.Sprintf("Backup Mover (moves to %s)", b.backupDir)
}

// ─────────────────────────────────────────────────────────────────────────────
// LogRotator: Appends timestamp to filename before rotating
// ─────────────────────────────────────────────────────────────────────────────

type LogRotator struct {
	timestamp string
}

func (l *LogRotator) GetNewPath(originalPath string) string {
	dir := filepath.Dir(originalPath)
	base := filepath.Base(originalPath)
	ext := filepath.Ext(originalPath)

	// Create new filename with timestamp: app_2025-01-04.log
	nameOnly := strings.TrimSuffix(base, ext)
	newFilename := nameOnly + "_" + l.timestamp + ext

	return filepath.Join(dir, newFilename)
}

func (l *LogRotator) GetDescription() string {
	return "Log Rotator (appends timestamp to filename)"
}

func Example5_FileProcessorInterface() {
	fmt.Println("\n" + strings.Repeat("═", 80))
	fmt.Println("EXAMPLE 5: FileProcessor Interface Pattern")
	fmt.Println(strings.Repeat("═", 80) + "\n")

	originalPath := "music/recordings/podcast_episode_42.wav"

	fmt.Printf("Original file path: %q\n\n", originalPath)
	fmt.Println(strings.Repeat("─", 80))

	// Create different processors
	processors := []FileProcessor{
		&MP3Converter{targetExtension: ".mp3"},
		&BackupMover{backupDir: "/backup"},
		&LogRotator{timestamp: "2025-01-04"},
	}

	// Demonstrate each processor
	for i, processor := range processors {
		newPath := processor.GetNewPath(originalPath)
		fmt.Printf("\n%d. %s\n", i+1, processor.GetDescription())
		fmt.Printf("   New path: %q\n", newPath)
	}

	fmt.Println()
}

/*
═══════════════════════════════════════════════════════════════════════════════
                    BEST PRACTICES & SECURITY
═══════════════════════════════════════════════════════════════════════════════

1. ALWAYS USE filepath.Join()
   ✗ DON'T: path := "documents/" + filename
   ✓ DO:    path := filepath.Join("documents", filename)
   → Ensures correct separators on all operating systems

2. ERROR HANDLING FOR Abs() AND Rel()
   Both functions return errors. Always check for them:

   abs, err := filepath.Abs(userPath)
   if err != nil {
       // Handle error - path might not exist or be invalid
       log.Fatal(err)
   }

3. PATH SANITIZATION & SECURITY
   DANGER: Directory Traversal attacks using ".."

   ✗ Vulnerable:
       userInput := "../../etc/passwd"
       path := filepath.Join("/home/user", userInput)
       // Becomes: /etc/passwd  <-- DANGEROUS!

   ✓ Secure:
       safePath := filepath.Clean(filepath.Join("/home/user", userInput))
       // Check if path is still within base directory
       absPath, _ := filepath.Abs(safePath)
       if !strings.HasPrefix(absPath, "/home/user") {
           return fmt.Errorf("path traversal attack detected")
       }

4. USE filepath.HasPrefix() FOR VALIDATION
   When validating that a path is within an expected directory:

   if !strings.HasPrefix(path, allowedDir) {
       return fmt.Errorf("path is outside allowed directory")
   }

5. CLEAN BEFORE VALIDATING
   Always call filepath.Clean() before checking paths:

   cleanPath := filepath.Clean(userPath)
   // Now safe to validate or use

═══════════════════════════════════════════════════════════════════════════════
*/

func Example6_SecurityAndValidation() {
	fmt.Println("\n" + strings.Repeat("═", 80))
	fmt.Println("EXAMPLE 6: Security & Validation")
	fmt.Println(strings.Repeat("═", 80) + "\n")

	baseDir := "/home/user/documents"

	testPaths := []string{
		"reports/2025.pdf",            // SAFE: Normal relative path
		"../../etc/passwd",            // DANGEROUS: Directory traversal
		"./././reports/file.txt",      // SUSPICIOUS: Redundant dot notation
		"reports/../reports/file.txt", // SUSPICIOUS: Unnecessary navigation
	}

	fmt.Printf("Base directory: %q\n\n", baseDir)
	fmt.Println(strings.Repeat("─", 80))

	for _, testPath := range testPaths {
		fmt.Printf("\nUser input: %q\n", testPath)

		// Step 1: Clean the path
		cleanedPath := filepath.Clean(filepath.Join(baseDir, testPath))
		fmt.Printf("Cleaned:    %q\n", cleanedPath)

		// Step 2: Check if it's still within base directory
		isValid := strings.HasPrefix(cleanedPath, baseDir)

		if isValid {
			fmt.Printf("Status:     ✓ SAFE\n")
		} else {
			fmt.Printf("Status:     ✗ BLOCKED (path traversal attempt)\n")
		}
	}
}

/*
═══════════════════════════════════════════════════════════════════════════════
                        QUICK REFERENCE TABLE
═══════════════════════════════════════════════════════════════════════════════

FUNCTION            | PURPOSE                      | EXAMPLE
────────────────────|──────────────────────────────|─────────────────────────
Join()              | Combine path elements        | Join("a", "b", "c")
Clean()             | Normalize/simplify path      | Clean("a/../b//c")
Split()             | Separate dir & filename      | Split("/a/b/file.txt")
Dir()               | Get directory only           | Dir("/a/b/file.txt")
Base()              | Get filename only            | Base("/a/b/file.txt")
Ext()               | Get extension only           | Ext("file.txt")
Abs()               | Relative → Absolute          | Abs("./file.txt")
Rel()               | Calculate relative path      | Rel("/a", "/a/b/c")

COMMON PATTERNS:

Change extension:
  newPath := strings.TrimSuffix(file, filepath.Ext(file)) + ".new"

Move to different directory:
  newPath := filepath.Join(newDir, filepath.Base(file))

Ensure absolute path:
  abs, _ := filepath.Abs(userPath)

═══════════════════════════════════════════════════════════════════════════════
*/

// Demo function matching the original style
func Demo86FilePaths() {
	fmt.Println("\n" + strings.Repeat("═", 80))
	fmt.Println("86 FILE PATHS - COMPLETE GUIDE")
	fmt.Println(strings.Repeat("═", 80))

	Example1_JoinAndClean()
	Example2_DeconstructingPaths()
	Example3_ExtensionHandling()
	Example4_AdvancedNavigation()
	Example5_FileProcessorInterface()
	Example6_SecurityAndValidation()

	fmt.Println("\n" + strings.Repeat("═", 80))
	fmt.Println("KEY TAKEAWAY:")
	fmt.Println("  Always use filepath.Join() instead of string concatenation")
	fmt.Println("  Always validate user-provided paths for security")
	fmt.Println("  Always check errors from Abs() and Rel() functions")
	fmt.Println(strings.Repeat("═", 80) + "\n")
}
