package main

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
)

/*
═══════════════════════════════════════════════════════════════════════════════
                      WORKING WITH DIRECTORIES IN GO
═══════════════════════════════════════════════════════════════════════════════

Directory operations are primarily handled by:
  • os package:       Basic operations (Mkdir, Chdir, ReadDir, Remove)
  • path/filepath:    Advanced navigation (WalkDir for recursive traversal)

This module covers:
  1. Creating directories (single and nested)
  2. Navigating directories (checking location, changing location)
  3. Reading directory contents (listing files)
  4. Walking directory trees (recursive exploration)
  5. Deleting directories (safe and unsafe methods)
  6. Practical patterns (finding files, generating reports)

═══════════════════════════════════════════════════════════════════════════════
                      CORE CONCEPTS
═══════════════════════════════════════════════════════════════════════════════

CURRENT WORKING DIRECTORY (CWD):
  Your program has a "cursor" position in the file system.
  • Default: The directory where the program is executed from
  • Check:   os.Getwd() returns the current directory
  • Change:  os.Chdir(path) moves the cursor to a new location

PERMISSIONS (0755):
  Linux-style file permissions represented as octal:
  • First digit (7):  Owner can read, write, execute
  • Second digit (5): Group can read and execute
  • Third digit (5):  Public can read and execute
  → 0755 is the standard safe permission for directories

═══════════════════════════════════════════════════════════════════════════════
*/

/*
━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━
  SECTION 1: CREATING DIRECTORIES
━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━
*/

func Example1_CreatingDirectories() {
	fmt.Println("\n" + strings.Repeat("═", 80))
	fmt.Println("EXAMPLE 1: Creating Directories")
	fmt.Println(strings.Repeat("═", 80) + "\n")

	// ─────────────────────────────────────────────────────────────────────────
	// os.Mkdir(): Create a single directory
	// ─────────────────────────────────────────────────────────────────────────
	fmt.Println("📌 os.Mkdir() - Create a single directory")
	fmt.Println(strings.Repeat("─", 80))

	singleDir := "demo_single_dir"
	fmt.Printf("Creating: %q\n", singleDir)

	err := os.Mkdir(singleDir, 0755)
	if err != nil {
		fmt.Printf("  ✗ Error: %v\n", err)
		fmt.Println("  (Directory might already exist or parent doesn't exist)\n")
	} else {
		fmt.Printf("  ✓ Successfully created\n\n")
		defer os.Remove(singleDir) // Clean up after demonstration
	}

	// ─────────────────────────────────────────────────────────────────────────
	// os.MkdirAll(): Create nested directories (RECOMMENDED)
	// ─────────────────────────────────────────────────────────────────────────
	fmt.Println("📌 os.MkdirAll() - Create nested directory structure")
	fmt.Println(strings.Repeat("─", 80))

	nestedPath := "demo_nested/level1/level2/level3"
	fmt.Printf("Creating: %q\n", nestedPath)

	err = os.MkdirAll(nestedPath, 0755)
	if err != nil {
		fmt.Printf("  ✗ Error: %v\n", err)
	} else {
		fmt.Printf("  ✓ Successfully created entire path\n")
		fmt.Printf("  (Creates all parent directories automatically)\n\n")
		defer os.RemoveAll("demo_nested") // Clean up entire tree
	}

	// ─────────────────────────────────────────────────────────────────────────
	// os.MkdirTemp(): Create temporary directory
	// ─────────────────────────────────────────────────────────────────────────
	fmt.Println("📌 os.MkdirTemp() - Create temporary directory")
	fmt.Println(strings.Repeat("─", 80))

	tempDir, err := os.MkdirTemp("", "gotut_demo_")
	if err != nil {
		fmt.Printf("  ✗ Error: %v\n", err)
	} else {
		fmt.Printf("  ✓ Created temp directory: %q\n", tempDir)
		fmt.Println("  (OS manages cleanup, typically in /tmp or %TEMP%)")
		defer os.RemoveAll(tempDir)
	}
}

/*
━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━
  SECTION 2: NAVIGATING DIRECTORIES
━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━
*/

func Example2_NavigatingDirectories() {
	fmt.Println("\n" + strings.Repeat("═", 80))
	fmt.Println("EXAMPLE 2: Navigating Directories")
	fmt.Println(strings.Repeat("═", 80) + "\n")

	// ─────────────────────────────────────────────────────────────────────────
	// os.Getwd(): Get current working directory
	// ─────────────────────────────────────────────────────────────────────────
	fmt.Println("📌 os.Getwd() - Get current working directory")
	fmt.Println(strings.Repeat("─", 80))

	cwd, err := os.Getwd()
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	} else {
		fmt.Printf("Current location: %q\n\n", cwd)
	}

	// ─────────────────────────────────────────────────────────────────────────
	// os.Chdir(): Change current working directory
	// ─────────────────────────────────────────────────────────────────────────
	fmt.Println("📌 os.Chdir() - Change working directory")
	fmt.Println(strings.Repeat("─", 80))

	// Create a test directory to navigate to
	testDir := "demo_nav_test"
	os.Mkdir(testDir, 0755)
	defer os.RemoveAll(testDir)

	fmt.Printf("Original location: %q\n", cwd)

	err = os.Chdir(testDir)
	if err != nil {
		fmt.Printf("Error changing directory: %v\n", err)
	} else {
		newCwd, _ := os.Getwd()
		fmt.Printf("Changed to:       %q\n", newCwd)

		// Change back to original location
		os.Chdir(cwd)
		backCwd, _ := os.Getwd()
		fmt.Printf("Changed back to:  %q\n\n", backCwd)
	}
}

/*
━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━
  SECTION 3: READING DIRECTORY CONTENTS
━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━

os.ReadDir() returns a slice of os.DirEntry objects.
Each DirEntry is lightweight and can tell you:
  • Name():     The file/folder name
  • IsDir():    Whether it's a directory (true) or file (false)
  • Type():     The file mode/permissions
  • Info():     Full FileInfo structure with size, modification time, etc.
━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━
*/

func Example3_ReadingDirectoryContents() {
	fmt.Println("\n" + strings.Repeat("═", 80))
	fmt.Println("EXAMPLE 3: Reading Directory Contents")
	fmt.Println(strings.Repeat("═", 80) + "\n")

	// Create a test directory structure
	testDir := "demo_read_test"
	os.MkdirAll(testDir+"/subdir", 0755)
	os.Create(testDir + "/file1.txt")
	os.Create(testDir + "/file2.go")
	defer os.RemoveAll(testDir)

	// ─────────────────────────────────────────────────────────────────────────
	// os.ReadDir(): List directory contents
	// ─────────────────────────────────────────────────────────────────────────
	fmt.Println("📌 os.ReadDir() - List directory contents")
	fmt.Println(strings.Repeat("─", 80))

	fmt.Printf("Contents of %q:\n\n", testDir)

	entries, err := os.ReadDir(testDir)
	if err != nil {
		fmt.Printf("Error reading directory: %v\n", err)
		return
	}

	for i, entry := range entries {
		fmt.Printf("%d. Name: %-20s ", i+1, entry.Name())

		if entry.IsDir() {
			fmt.Print("Type: [DIRECTORY]  ")
		} else {
			fmt.Print("Type: [FILE]       ")
		}

		// Get more detailed info
		info, _ := entry.Info()
		fmt.Printf("Size: %d bytes\n", info.Size())
	}

	fmt.Println()

	// ─────────────────────────────────────────────────────────────────────────
	// Filtering: Only directories or specific file types
	// ─────────────────────────────────────────────────────────────────────────
	fmt.Println("📌 Filtering Results - Only directories")
	fmt.Println(strings.Repeat("─", 80))

	for _, entry := range entries {
		if entry.IsDir() {
			fmt.Printf("  📁 %s\n", entry.Name())
		}
	}

	fmt.Println("\n📌 Filtering Results - Only files")
	fmt.Println(strings.Repeat("─", 80))

	for _, entry := range entries {
		if !entry.IsDir() {
			fmt.Printf("  📄 %s\n", entry.Name())
		}
	}
}

/*
━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━
  SECTION 4: WALKING DIRECTORY TREES (RECURSIVE)
━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━

filepath.WalkDir() is the modern, efficient way to explore a directory tree
recursively. It automatically handles:
  • Traversing subdirectories
  • Calling your callback function for every file and folder
  • Handling errors gracefully

SIGNATURE:
  func WalkDir(root string, fn fs.WalkDirFunc) error

CALLBACK FUNCTION:
  func(path string, d fs.DirEntry, err error) error

RETURN VALUE:
  • Return nil to continue walking
  • Return filepath.SkipDir to skip this directory and its contents
  • Return any other error to stop walking
━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━
*/

func Example4_WalkingDirectoryTrees() {
	fmt.Println("\n" + strings.Repeat("═", 80))
	fmt.Println("EXAMPLE 4: Walking Directory Trees Recursively")
	fmt.Println(strings.Repeat("═", 80) + "\n")

	// Create a nested directory structure for testing
	testRoot := "demo_walk_test"
	os.MkdirAll(testRoot+"/src/pkg1", 0755)
	os.MkdirAll(testRoot+"/src/pkg2", 0755)
	os.MkdirAll(testRoot+"/docs", 0755)
	os.Create(testRoot + "/README.md")
	os.Create(testRoot + "/src/main.go")
	os.Create(testRoot + "/src/pkg1/file1.go")
	os.Create(testRoot + "/src/pkg2/file2.go")
	os.Create(testRoot + "/docs/guide.txt")
	defer os.RemoveAll(testRoot)

	// ─────────────────────────────────────────────────────────────────────────
	// Basic WalkDir: List all files and folders
	// ─────────────────────────────────────────────────────────────────────────
	fmt.Println("📌 Basic WalkDir - List all items in tree")
	fmt.Println(strings.Repeat("─", 80))

	fileCount := 0
	dirCount := 0

	err := filepath.WalkDir(testRoot, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		// Indent based on depth
		depth := strings.Count(path, string(os.PathSeparator))
		indent := strings.Repeat("  ", depth)

		if d.IsDir() {
			fmt.Printf("%s📁 %s/\n", indent, d.Name())
			dirCount++
		} else {
			fmt.Printf("%s📄 %s\n", indent, d.Name())
			fileCount++
		}

		return nil
	})

	if err != nil {
		fmt.Printf("Error walking directory: %v\n", err)
	}

	fmt.Printf("\nSummary: Found %d directories and %d files\n\n", dirCount, fileCount)

	// ─────────────────────────────────────────────────────────────────────────
	// Filtered WalkDir: Only .go files
	// ─────────────────────────────────────────────────────────────────────────
	fmt.Println("📌 Filtered WalkDir - Find all .go files")
	fmt.Println(strings.Repeat("─", 80))

	goCount := 0

	filepath.WalkDir(testRoot, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		// Skip directories in the filter
		if !d.IsDir() && strings.HasSuffix(d.Name(), ".go") {
			fmt.Printf("  ✓ %s\n", path)
			goCount++
		}

		return nil
	})

	fmt.Printf("Found %d Go files\n\n", goCount)

	// ─────────────────────────────────────────────────────────────────────────
	// Advanced: Skip directories
	// ─────────────────────────────────────────────────────────────────────────
	fmt.Println("📌 Advanced - Skip specific directories")
	fmt.Println(strings.Repeat("─", 80))

	fmt.Println("Walking tree (but skipping 'docs' directory):")

	filepath.WalkDir(testRoot, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		// Skip the 'docs' directory entirely
		if d.IsDir() && d.Name() == "docs" {
			fmt.Printf("  ⏭️  Skipping: %s\n", path)
			return filepath.SkipDir // Don't descend into this directory
		}

		if !d.IsDir() {
			fmt.Printf("  ✓ %s\n", path)
		}

		return nil
	})
}

/*
━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━
  SECTION 5: DELETING DIRECTORIES
━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━

Two deletion strategies with very different behaviors:

os.Remove(path):
  • Deletes ONE item (file OR empty directory only)
  • Fails if directory contains anything
  • SAFE: Limited scope of damage if misused

os.RemoveAll(path):
  • Deletes the path AND everything inside it recursively
  • Equivalent to: rm -rf on Linux
  • DANGEROUS: Can delete large directory trees
  • Use with caution!
━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━
*/

func Example5_DeletingDirectories() {
	fmt.Println("\n" + strings.Repeat("═", 80))
	fmt.Println("EXAMPLE 5: Deleting Directories")
	fmt.Println(strings.Repeat("═", 80) + "\n")

	// ─────────────────────────────────────────────────────────────────────────
	// os.Remove(): Delete empty directory only
	// ─────────────────────────────────────────────────────────────────────────
	fmt.Println("📌 os.Remove() - Delete empty directory")
	fmt.Println(strings.Repeat("─", 80))

	emptyDir := "demo_empty_dir"
	os.Mkdir(emptyDir, 0755)

	fmt.Printf("Attempting to delete empty directory: %q\n", emptyDir)
	err := os.Remove(emptyDir)
	if err != nil {
		fmt.Printf("  ✗ Error: %v\n", err)
	} else {
		fmt.Printf("  ✓ Successfully deleted\n\n")
	}

	// ─────────────────────────────────────────────────────────────────────────
	// os.Remove() on non-empty directory: FAILS
	// ─────────────────────────────────────────────────────────────────────────
	fmt.Println("📌 os.Remove() on non-empty directory - FAILS (as expected)")
	fmt.Println(strings.Repeat("─", 80))

	nonEmptyDir := "demo_nonempty_dir"
	os.Mkdir(nonEmptyDir, 0755)
	os.Create(nonEmptyDir + "/file.txt")

	fmt.Printf("Directory: %q (contains 1 file)\n", nonEmptyDir)
	fmt.Println("Attempting os.Remove()...")

	err = os.Remove(nonEmptyDir)
	if err != nil {
		fmt.Printf("  ✗ Failed (expected): %v\n\n", err)
	}

	os.RemoveAll(nonEmptyDir) // Clean up properly

	// ─────────────────────────────────────────────────────────────────────────
	// os.RemoveAll(): Delete directory and ALL contents
	// ─────────────────────────────────────────────────────────────────────────
	fmt.Println("📌 os.RemoveAll() - Delete directory and all contents")
	fmt.Println(strings.Repeat("─", 80))

	treeDir := "demo_tree_delete"
	os.MkdirAll(treeDir+"/level1/level2", 0755)
	os.Create(treeDir + "/file1.txt")
	os.Create(treeDir + "/level1/file2.txt")
	os.Create(treeDir + "/level1/level2/file3.txt")

	fmt.Printf("Directory structure created:\n")
	fmt.Printf("  %s/\n", treeDir)
	fmt.Printf("    ├── file1.txt\n")
	fmt.Printf("    └── level1/\n")
	fmt.Printf("        ├── file2.txt\n")
	fmt.Printf("        └── level2/\n")
	fmt.Printf("            └── file3.txt\n\n")

	fmt.Println("Executing os.RemoveAll()...")
	err = os.RemoveAll(treeDir)
	if err != nil {
		fmt.Printf("  ✗ Error: %v\n", err)
	} else {
		fmt.Printf("  ✓ Entire tree deleted\n")
	}
}

/*
━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━
  SECTION 6: PRACTICAL PATTERN - FINDING FILES BY EXTENSION
━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━

Real-world scenario: Find all .jpg image files in a massive directory tree
and generate a report.

This combines:
  • filepath.WalkDir() for recursive traversal
  • filepath.Ext() to check file extensions
  • Collecting statistics during the walk
━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━
*/

func Example6_FindingFilesByExtension() {
	fmt.Println("\n" + strings.Repeat("═", 80))
	fmt.Println("EXAMPLE 6: Practical Pattern - Finding Files by Extension")
	fmt.Println(strings.Repeat("═", 80) + "\n")

	// Create a sample directory structure with various files
	testDir := "demo_image_search"
	os.MkdirAll(testDir+"/photos/vacation", 0755)
	os.MkdirAll(testDir+"/documents", 0755)
	os.MkdirAll(testDir+"/code", 0755)

	// Create sample files
	files := []string{
		testDir + "/photo1.jpg",
		testDir + "/photo2.JPG", // uppercase extension
		testDir + "/photos/sunset.jpg",
		testDir + "/photos/vacation/beach.png",
		testDir + "/photos/vacation/people.jpg",
		testDir + "/documents/report.pdf",
		testDir + "/documents/image.jpg",
		testDir + "/code/main.go",
		testDir + "/code/utils.go",
	}

	for _, f := range files {
		os.Create(f)
	}
	defer os.RemoveAll(testDir)

	// ─────────────────────────────────────────────────────────────────────────
	// Search for .jpg files
	// ─────────────────────────────────────────────────────────────────────────
	fmt.Println("📌 Searching for all .jpg image files")
	fmt.Println(strings.Repeat("─", 80))

	targetExt := ".jpg"
	var jpgFiles []string
	totalSize := int64(0)

	err := filepath.WalkDir(testDir, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		if !d.IsDir() {
			// Check extension (case-insensitive)
			if strings.EqualFold(filepath.Ext(path), targetExt) {
				jpgFiles = append(jpgFiles, path)

				// Get file size
				info, _ := d.Info()
				totalSize += info.Size()
			}
		}

		return nil
	})

	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	// Display results
	fmt.Printf("Found %d %s files:\n\n", len(jpgFiles), targetExt)
	for i, file := range jpgFiles {
		fmt.Printf("  %d. %s\n", i+1, file)
	}

	fmt.Printf("\nTotal size: %d bytes\n", totalSize)
}

/*
═══════════════════════════════════════════════════════════════════════════════
                    BEST PRACTICES SUMMARY
═══════════════════════════════════════════════════════════════════════════════

1. ERROR HANDLING IS MANDATORY
   All file system operations can fail. Always check errors:

   if err != nil {
       log.Fatal(err)  // or handle appropriately
   }

2. USE MkdirAll() BY DEFAULT
   It's safer than Mkdir() because:
   • Creates entire path automatically
   • Doesn't fail if directory already exists
   • Handles nested structures in one call

3. PREFER WalkDir() FOR RECURSION
   Modern and efficient:
   • Lower memory footprint than Walk()
   • Lazy evaluation - doesn't load entire tree
   • Better error handling per entry

4. ALWAYS USE filepath.Join() FOR PATHS
   ✗ DON'T:  path := dir + "/" + file
   ✓ DO:     path := filepath.Join(dir, file)

5. BE VERY CAREFUL WITH RemoveAll()
   Think of it as "rm -rf":
   • Deletes recursively without confirmation
   • No way to undo
   • Verify path carefully before calling
   • Consider adding a safety check

6. USE DEFER FOR CLEANUP
   Ensures resources are cleaned up even if function panics:

   tempDir, _ := os.MkdirTemp("", "prefix")
   defer os.RemoveAll(tempDir)
   // tempDir will be cleaned up when function exits

═══════════════════════════════════════════════════════════════════════════════
*/

// Demo function matching the original style
func Demo87Directories() {
	fmt.Println("\n" + strings.Repeat("═", 80))
	fmt.Println("87 DIRECTORIES - COMPLETE GUIDE")
	fmt.Println(strings.Repeat("═", 80))

	Example1_CreatingDirectories()
	Example2_NavigatingDirectories()
	Example3_ReadingDirectoryContents()
	Example4_WalkingDirectoryTrees()
	Example5_DeletingDirectories()
	Example6_FindingFilesByExtension()

	fmt.Println("\n" + strings.Repeat("═", 80))
	fmt.Println("KEY TAKEAWAYS:")
	fmt.Println("  • Use MkdirAll() for creating nested directories")
	fmt.Println("  • Use filepath.WalkDir() for recursive traversal")
	fmt.Println("  • Always handle errors from file operations")
	fmt.Println("  • Be careful with RemoveAll() - it's irreversible!")
	fmt.Println("  • Use filepath.Join() for portable path construction")
	fmt.Println(strings.Repeat("═", 80) + "\n")
}
