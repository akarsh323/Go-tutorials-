package main

import (
	"crypto/md5"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
	"time"
)

/*
═══════════════════════════════════════════════════════════════════════════════
                 TEMPORARY FILES AND DIRECTORIES IN GO
═══════════════════════════════════════════════════════════════════════════════

Temporary storage is used for data that is only needed for a short period—
typically just while the program is running.

COMMON USE CASES:
  • Transient Operations: File uploads, data processing, caching intermediate results
  • Isolation: Keeping test data separate from permanent storage
  • Preprocessing: Temporary transformations before final output
  • Testing: Isolated test data that won't affect the system

WHERE ARE THEY STORED?
  • Linux/macOS:  /tmp or /var/tmp
  • Windows:      C:\Users\<username>\AppData\Local\Temp
  • Go finds this automatically - you don't hardcode paths!

KEY PRINCIPLE:
  Go automatically creates UNIQUE names by replacing '*' with random strings.
  This prevents collisions (conflicts) when multiple programs create temp files.

═══════════════════════════════════════════════════════════════════════════════
*/

/*
━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━
  SECTION 1: CREATING TEMPORARY FILES
━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━

os.CreateTemp(dir, pattern) returns:
  1. A file handle (os.File) ready for reading/writing
  2. An error (if creation failed)

PARAMETERS:
  • dir:     "" = use default system temp directory
             "specific/path" = create temp file in that directory
  • pattern: "prefix_*" = asterisk is replaced with random unique string

NAMING:
  • Pattern:   "cache_*.tmp"
  • Result:    "cache_a1b2c3d4e5f6g7h8.tmp"  (in system temp dir)
━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━
*/

func Example1_CreatingTemporaryFiles() {
	fmt.Println("\n" + strings.Repeat("═", 80))
	fmt.Println("EXAMPLE 1: Creating Temporary Files")
	fmt.Println(strings.Repeat("═", 80) + "\n")

	// ─────────────────────────────────────────────────────────────────────────
	// Basic temporary file creation
	// ─────────────────────────────────────────────────────────────────────────
	fmt.Println("📌 Basic Temporary File Creation")
	fmt.Println(strings.Repeat("─", 80))

	file, err := os.CreateTemp("", "gotut_example_*.txt")
	if err != nil {
		fmt.Printf("Error creating temp file: %v\n", err)
		return
	}

	fmt.Printf("Temp file created at: %q\n", file.Name())
	fmt.Printf("File is ready for:    Reading and Writing\n\n")

	// Write some content to the file
	_, err = file.WriteString("This is temporary content.\n")
	if err != nil {
		fmt.Printf("Error writing to file: %v\n", err)
	}

	// Seek back to beginning to read what we wrote
	file.Seek(0, 0)
	content := make([]byte, 100)
	n, _ := file.Read(content)
	fmt.Printf("Content written:      %q\n\n", string(content[:n]))

	// CLEANUP: Always close the file handle
	file.Close()
	defer os.Remove(file.Name()) // Remove the file from disk

	// ─────────────────────────────────────────────────────────────────────────
	// Multiple temporary files with clear naming
	// ─────────────────────────────────────────────────────────────────────────
	fmt.Println("📌 Creating Multiple Temp Files with Clear Naming")
	fmt.Println(strings.Repeat("─", 80))

	tempFiles := []string{
		"input_*.csv",
		"output_*.json",
		"cache_*.tmp",
	}

	fmt.Println("Creating multiple temp files:\n")

	for _, pattern := range tempFiles {
		f, _ := os.CreateTemp("", pattern)
		fmt.Printf("  Pattern: %-20s → %s\n", pattern, filepath.Base(f.Name()))
		f.Close()
		defer os.Remove(f.Name())
	}
}

/*
━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━
  SECTION 2: CREATING TEMPORARY DIRECTORIES
━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━

os.MkdirTemp(dir, pattern) creates a temporary directory.

ADVANTAGES OVER CreateTemp:
  • Can store multiple related temporary files together
  • Cleaner organization for complex operations
  • Can create subdirectories within the temp space

CLEANUP REQUIREMENT:
  Use os.RemoveAll() to delete the directory AND everything inside it.
  This is critical - forgetting this leads to disk bloat!
━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━
*/

func Example2_CreatingTemporaryDirectories() {
	fmt.Println("\n" + strings.Repeat("═", 80))
	fmt.Println("EXAMPLE 2: Creating Temporary Directories")
	fmt.Println(strings.Repeat("═", 80) + "\n")

	// ─────────────────────────────────────────────────────────────────────────
	// Basic temporary directory creation
	// ─────────────────────────────────────────────────────────────────────────
	fmt.Println("📌 Basic Temporary Directory Creation")
	fmt.Println(strings.Repeat("─", 80))

	tempDir, err := os.MkdirTemp("", "gotut_workspace_*")
	if err != nil {
		fmt.Printf("Error creating temp directory: %v\n", err)
		return
	}

	fmt.Printf("Temp directory created at: %q\n", tempDir)
	fmt.Printf("Ready to create files in:  %q\n\n", tempDir)

	// Create some files inside the temp directory
	files := []string{"input.txt", "output.txt", "metadata.json"}
	fmt.Println("Creating files inside temp directory:")

	for _, fname := range files {
		fullPath := filepath.Join(tempDir, fname)
		f, _ := os.Create(fullPath)
		f.WriteString("Temporary content for: " + fname)
		f.Close()
		fmt.Printf("  ✓ Created: %s\n", fname)
	}

	fmt.Println()

	// List contents
	entries, _ := os.ReadDir(tempDir)
	fmt.Printf("Directory contains %d items\n\n", len(entries))

	// CLEANUP: Remove entire directory tree
	defer os.RemoveAll(tempDir)

	// ─────────────────────────────────────────────────────────────────────────
	// Nested temporary directories
	// ─────────────────────────────────────────────────────────────────────────
	fmt.Println("📌 Nested Temporary Directories")
	fmt.Println(strings.Repeat("─", 80))

	nestedTemp, _ := os.MkdirTemp("", "gotut_project_*")
	defer os.RemoveAll(nestedTemp)

	// Create subdirectories
	subdirs := []string{"data", "logs", "cache"}
	for _, subdir := range subdirs {
		os.Mkdir(filepath.Join(nestedTemp, subdir), 0755)
	}

	fmt.Printf("Created temp directory: %s\n", filepath.Base(nestedTemp))
	fmt.Println("Structure:")

	for _, subdir := range subdirs {
		fmt.Printf("  📁 %s/\n", subdir)
	}
}

/*
━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━
  SECTION 3: CLEANUP PATTERNS - THE CRITICAL ASPECT
━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━

IMPORTANT: Temporary files/directories MUST be cleaned up!

Forgetting cleanup leads to:
  • Disk bloat (accumulation of unused files)
  • Security issues (sensitive temp data left on disk)
  • Performance degradation (filesystem fills up)

BEST PRACTICE: Use defer immediately after creation to guarantee cleanup
even if the function crashes or returns early.

CLEANUP STRATEGIES:
  1. os.Remove(path)     - Delete single file or empty directory
  2. os.RemoveAll(path)  - Delete directory and everything inside
  3. Always use defer    - Guarantees cleanup happens
━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━
*/

func Example3_CleanupPatterns() {
	fmt.Println("\n" + strings.Repeat("═", 80))
	fmt.Println("EXAMPLE 3: Cleanup Patterns (Critical for Safety)")
	fmt.Println(strings.Repeat("═", 80) + "\n")

	// ─────────────────────────────────────────────────────────────────────────
	// PATTERN 1: Cleanup with defer (RECOMMENDED)
	// ─────────────────────────────────────────────────────────────────────────
	fmt.Println("📌 PATTERN 1: Using defer for guaranteed cleanup")
	fmt.Println(strings.Repeat("─", 80))

	file, _ := os.CreateTemp("", "cleanup_defer_*.txt")
	filename := file.Name()

	// IMMEDIATELY after creation, set up cleanup with defer
	defer os.Remove(filename)
	defer file.Close()

	file.WriteString("This will be cleaned up properly")

	fmt.Printf("Created file: %s\n", filepath.Base(filename))
	fmt.Println("defer statements registered:")
	fmt.Println("  1. defer file.Close()")
	fmt.Println("  2. defer os.Remove(filename)")
	fmt.Println("  → Cleanup GUARANTEED even if function panics!\n")

	// ─────────────────────────────────────────────────────────────────────────
	// PATTERN 2: Nested cleanup (files inside temp directory)
	// ─────────────────────────────────────────────────────────────────────────
	fmt.Println("📌 PATTERN 2: Cleanup nested structure")
	fmt.Println(strings.Repeat("─", 80))

	tempDir, _ := os.MkdirTemp("", "nested_cleanup_*")

	// Register cleanup for entire tree FIRST
	defer os.RemoveAll(tempDir)

	// Then create files inside
	for i := 1; i <= 3; i++ {
		filepath := filepath.Join(tempDir, fmt.Sprintf("file_%d.txt", i))
		os.WriteFile(filepath, []byte(fmt.Sprintf("content %d", i)), 0644)
	}

	fmt.Printf("Created temp directory: %s\n", filepath.Base(tempDir))
	fmt.Println("  ├── file_1.txt")
	fmt.Println("  ├── file_2.txt")
	fmt.Println("  └── file_3.txt")
	fmt.Println("defer os.RemoveAll() will clean entire tree\n")

	// ─────────────────────────────────────────────────────────────────────────
	// PATTERN 3: Manual cleanup vs defer
	// ─────────────────────────────────────────────────────────────────────────
	fmt.Println("📌 PATTERN 3: Comparing manual vs defer cleanup")
	fmt.Println(strings.Repeat("─", 80))

	fmt.Println("❌ DON'T DO THIS (manual cleanup can be missed):")
	fmt.Println(`
  f, _ := os.CreateTemp("", "bad_*.txt")
  f.WriteString("content")
  f.Close()
  os.Remove(f.Name())    // What if we forget this line?
                         // What if function returns early?
	`)

	fmt.Println("\n✓ DO THIS INSTEAD (defer guarantees cleanup):")
	fmt.Println(`
  f, _ := os.CreateTemp("", "good_*.txt")
  defer f.Close()
  defer os.Remove(f.Name())
  
  f.WriteString("content")
  // Even if we return early, defer will execute!
	`)
}

/*
━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━
  SECTION 4: PRACTICAL PATTERN - FILE UPLOAD PROCESSING
━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━

Real-world scenario: User uploads a file, we need to:
  1. Save it to a temp location
  2. Validate/scan it
  3. Process it
  4. Clean up the temp file

This demonstrates a complete workflow using temporary files.
━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━
*/

func Example4_FileUploadProcessing() {
	fmt.Println("\n" + strings.Repeat("═", 80))
	fmt.Println("EXAMPLE 4: Practical Pattern - File Upload Processing")
	fmt.Println(strings.Repeat("═", 80) + "\n")

	// Simulate file upload
	uploadedContent := []byte("User uploaded file content here")

	fmt.Println("SCENARIO: File upload workflow")
	fmt.Println(strings.Repeat("─", 80))
	fmt.Println("Step 1: Create temp file to store uploaded data")

	// Step 1: Create temp file for the upload
	tempFile, err := os.CreateTemp("", "upload_*.tmp")
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	defer tempFile.Close()
	defer os.Remove(tempFile.Name())

	fmt.Printf("  ✓ Created: %s\n", filepath.Base(tempFile.Name()))

	// Step 2: Write upload to temp file
	fmt.Println("\nStep 2: Write uploaded content to temp file")
	_, err = tempFile.Write(uploadedContent)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	fmt.Printf("  ✓ Wrote %d bytes\n", len(uploadedContent))

	// Step 3: Validate the file
	fmt.Println("\nStep 3: Validate uploaded file")

	// Calculate hash for validation
	tempFile.Seek(0, 0)
	hash := md5.New()
	io.Copy(hash, tempFile)
	checksum := fmt.Sprintf("%x", hash.Sum(nil))

	fmt.Printf("  ✓ Calculated MD5: %s\n", checksum[:16]+"...")

	// Step 4: Process the file
	fmt.Println("\nStep 4: Process the file")
	tempFile.Seek(0, 0)
	content := make([]byte, len(uploadedContent))
	tempFile.Read(content)
	fmt.Printf("  ✓ Read content: %q\n", string(content[:20])+"...")

	// Step 5: Report results
	fmt.Println("\nStep 5: Processing complete")
	fmt.Printf("  ✓ Temp file automatically cleaned by defer: {tempFile.Close(), os.Remove()}")
	fmt.Println("\n  (Deferred cleanup executes when function returns)\n")
}

/*
━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━
  SECTION 5: BATCH PROCESSING WITH TEMP DIRECTORY
━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━

Real-world scenario: Process multiple files together
  1. Create temp directory
  2. Extract/decompress files into it
  3. Process all files
  4. Clean up entire workspace

This demonstrates organizing multiple temp files together.
━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━
*/

func Example5_BatchProcessingWithTempDir() {
	fmt.Println("\n" + strings.Repeat("═", 80))
	fmt.Println("EXAMPLE 5: Batch Processing with Temp Directory")
	fmt.Println(strings.Repeat("═", 80) + "\n")

	fmt.Println("SCENARIO: Process batch of files in isolated temp workspace")
	fmt.Println(strings.Repeat("─", 80))

	// Create temp directory for batch processing
	batchDir, err := os.MkdirTemp("", "batch_processing_*")
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	defer os.RemoveAll(batchDir)

	fmt.Printf("Created batch workspace: %s\n\n", filepath.Base(batchDir))

	// Create subdirectories for organization
	inputDir := filepath.Join(batchDir, "input")
	outputDir := filepath.Join(batchDir, "output")
	logsDir := filepath.Join(batchDir, "logs")

	os.Mkdir(inputDir, 0755)
	os.Mkdir(outputDir, 0755)
	os.Mkdir(logsDir, 0755)

	fmt.Println("Directory structure:")
	fmt.Printf("  %s/\n", filepath.Base(batchDir))
	fmt.Println("  ├── input/")
	fmt.Println("  ├── output/")
	fmt.Println("  └── logs/\n")

	// Simulate processing multiple files
	fmt.Println("Processing workflow:")

	files := []string{"data_1.csv", "data_2.csv", "data_3.csv"}

	// Step 1: Create input files
	fmt.Println("\n  Step 1: Create input files")
	for _, fname := range files {
		path := filepath.Join(inputDir, fname)
		os.WriteFile(path, []byte("Sample data for "+fname), 0644)
		fmt.Printf("    ✓ Created: %s\n", fname)
	}

	// Step 2: Process each file
	fmt.Println("\n  Step 2: Process each file")
	for _, fname := range files {
		inputPath := filepath.Join(inputDir, fname)
		outputPath := filepath.Join(outputDir, "processed_"+fname)

		// Read input
		content, _ := os.ReadFile(inputPath)

		// Process (simulate uppercase transformation)
		processed := strings.ToUpper(string(content))

		// Write output
		os.WriteFile(outputPath, []byte(processed), 0644)
		fmt.Printf("    ✓ Processed: %s\n", fname)
	}

	// Step 3: Create log file
	fmt.Println("\n  Step 3: Create processing log")
	logPath := filepath.Join(logsDir, "processing.log")
	timestamp := time.Now().Format("2006-01-02 15:04:05")
	logEntry := fmt.Sprintf("[%s] Batch processing completed\n", timestamp)
	os.WriteFile(logPath, []byte(logEntry), 0644)
	fmt.Println("    ✓ Log created: processing.log")

	// Summary
	fmt.Println("\n  Step 4: Summary")
	fmt.Printf("    Input files:   %d\n", len(files))
	fmt.Println("    Output files:  3")
	fmt.Println("    Logs created:  1")
	fmt.Println("\n  Entire workspace will be cleaned up by defer os.RemoveAll()\n")
}

/*
━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━
  SECTION 6: SECURITY AND BEST PRACTICES
━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━

SECURITY CONSIDERATIONS:

1. SENSITIVE DATA IN TEMP FILES
   ❌ DON'T store passwords, tokens, or credentials in temp files
   ❌ DON'T assume temp directory is private (depends on OS and permissions)
   ✓ DO use secure deletion if handling sensitive data
   ✓ DO encrypt if necessary

2. UNIQUE NAMING PREVENTS COLLISIONS
   The '*' in the pattern is replaced with random unique string.
   This prevents multiple programs from overwriting each other's files.

3. ALWAYS USE DEFER FOR CLEANUP
   ✓ Guarantees cleanup even if exceptions occur
   ✓ Follows Go's idiomatic resource management pattern
   ✓ Makes code intentions clear

4. CHECK ERRORS
   File operations can fail for many reasons:
   • Disk full
   • Permissions denied
   • File system errors
   Always check err != nil

5. USE CLEAR NAMING PATTERNS
   ✓ "myapp_upload_*"    (easy to identify source)
   ✓ "batch_worker_*.tmp"
   ✗ "tmp_*"             (too generic)
   ✗ "x_*"               (unclear purpose)

6. CONSIDER ISOLATION
   Create a subdirectory for related temp files:
   tempDir := os.MkdirTemp("", "myapp_session_*")
   This keeps all session files together and cleanup is simple.

═══════════════════════════════════════════════════════════════════════════════
*/

func Example6_SecurityAndBestPractices() {
	fmt.Println("\n" + strings.Repeat("═", 80))
	fmt.Println("EXAMPLE 6: Security and Best Practices")
	fmt.Println(strings.Repeat("═", 80) + "\n")

	fmt.Println("📌 Security Best Practices")
	fmt.Println(strings.Repeat("─", 80))

	fmt.Println("\n1. UNIQUE NAMING PREVENTS COLLISIONS")
	fmt.Println("   Pattern: 'myapp_cache_*'")

	for i := 0; i < 3; i++ {
		f, _ := os.CreateTemp("", "myapp_cache_*")
		fmt.Printf("   Result: %s\n", filepath.Base(f.Name()))
		f.Close()
		defer os.Remove(f.Name())
	}

	fmt.Println("\n   (Each file has a unique random suffix - no collisions!)\n")

	fmt.Println("2. NAMING PATTERNS FOR IDENTIFICATION")
	fmt.Println("   ✓ GOOD:  'gotut_upload_*.tmp'  (clear source)")
	fmt.Println("   ✓ GOOD:  'batch_worker_*.log'  (clear purpose)")
	fmt.Println("   ✗ BAD:   'tmp_*'               (too generic)")
	fmt.Println("   ✗ BAD:   'x_*'                 (unclear)\n")

	fmt.Println("3. PROPER ERROR HANDLING")
	fmt.Println("   ALWAYS check for errors:")

	// Demonstrate error handling
	_, err := os.CreateTemp("/invalid/path", "*.txt")
	if err != nil {
		fmt.Printf("   Error caught: %v\n\n", err)
	}

	fmt.Println("4. STRUCTURED TEMP DIRECTORIES")
	fmt.Println("   Create dedicated temp directory for related files:")

	tempDir, _ := os.MkdirTemp("", "session_*")
	defer os.RemoveAll(tempDir)

	fmt.Printf("   Main dir:  %s/\n", filepath.Base(tempDir))
	fmt.Println("   Sub-dirs:")
	fmt.Println("     ├── cache/")
	fmt.Println("     ├── uploads/")
	fmt.Println("     └── processing/")
	fmt.Println("\n   Single os.RemoveAll() cleans everything!\n")

	fmt.Println("5. DEFER FOR GUARANTEED CLEANUP")
	fmt.Println("   Pattern:")
	fmt.Println(`
    resource, err := os.CreateTemp(...)
    if err != nil { handle(err) }
    
    defer resource.Close()        // Cleanup: file handle
    defer os.Remove(path)         // Cleanup: file from disk
    
    // Use resource here
    // Cleanup ALWAYS happens when function returns!
	`)
}

/*
═══════════════════════════════════════════════════════════════════════════════
                        QUICK REFERENCE TABLE
═══════════════════════════════════════════════════════════════════════════════

FUNCTION           | PURPOSE                    | EXAMPLE
─────────────────────|──────────────────────────|────────────────────────
CreateTemp         | Create temp file           | CreateTemp("", "cache_*.tmp")
MkdirTemp          | Create temp directory      | MkdirTemp("", "work_*")
                   |                            |
file.Close()       | Close file handle          | defer file.Close()
os.Remove()        | Delete single file/dir     | os.Remove(path)
os.RemoveAll()     | Delete dir + contents      | defer os.RemoveAll(tempDir)

CLEANUP PATTERN:
  resource, err := Create[Temp|Temp]()
  if err != nil { ... }
  defer cleanup()  // ALWAYS use defer!
  
  // Use resource

STORAGE LOCATIONS (automatic):
  Linux/macOS:      /tmp or /var/tmp
  Windows:          C:\Users\<username>\AppData\Local\Temp

═══════════════════════════════════════════════════════════════════════════════
*/

// Demo function matching the original style
func Demo88TempFilesDirs() {
	fmt.Println("\n" + strings.Repeat("═", 80))
	fmt.Println("88 TEMPORARY FILES AND DIRECTORIES - COMPLETE GUIDE")
	fmt.Println(strings.Repeat("═", 80))

	Example1_CreatingTemporaryFiles()
	Example2_CreatingTemporaryDirectories()
	Example3_CleanupPatterns()
	Example4_FileUploadProcessing()
	Example5_BatchProcessingWithTempDir()
	Example6_SecurityAndBestPractices()

	fmt.Println("\n" + strings.Repeat("═", 80))
	fmt.Println("KEY TAKEAWAYS:")
	fmt.Println("  • Always use defer for cleanup - it's the Go idiom")
	fmt.Println("  • Use os.RemoveAll() for directories, os.Remove() for files")
	fmt.Println("  • Unique naming (with *) prevents collisions automatically")
	fmt.Println("  • Clear naming patterns help identify temp files in system")
	fmt.Println("  • Missing cleanup leads to disk bloat and security issues!")
	fmt.Println(strings.Repeat("═", 80) + "\n")
}
