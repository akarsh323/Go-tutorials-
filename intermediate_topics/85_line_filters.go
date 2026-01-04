package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

/*
═══════════════════════════════════════════════════════════════════════════════
                        LINE FILTERING IN GO
═══════════════════════════════════════════════════════════════════════════════

LINE FILTERING is the process of reading text line-by-line, analyzing each line
against specific criteria, and then processing, modifying, or discarding that
line accordingly.

It's fundamental to:
  • Data cleaning and validation
  • Log analysis and monitoring
  • File transformation and ETL pipelines
  • Text processing workflows

═══════════════════════════════════════════════════════════════════════════════
                          CORE PROCESS
═══════════════════════════════════════════════════════════════════════════════

1. READ:    Open a file and create a bufio.Scanner to read line-by-line
2. CHECK:   Examine the current line (scanner.Text()) against criteria
3. PROCESS: If criteria is met, perform an action (print, modify, save, discard)

This three-step pattern is the foundation of all line filtering operations.

═══════════════════════════════════════════════════════════════════════════════
*/

/*
━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━
  EXAMPLE 1: BASIC FILTERING & TRANSFORMATION
━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━

Filter for lines containing a keyword, replace the keyword, and add line numbers.

PROCESS FLOW:
  • Read each line from the file
  • Check if it contains the keyword "important"
  • If yes: replace "important" with "necessary", add line number, print it
  • If no: skip the line (discard)

KEY TECHNIQUES:
  • strings.Contains(line, keyword) → Boolean check
  • strings.ReplaceAll(source, old, new) → Transformation
  • Manual counter for line numbering
━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━
*/

func Example1_BasicFilteringAndTransformation() {
	fmt.Println("\n=== EXAMPLE 1: Basic Filtering & Transformation ===\n")

	// SETUP: Open file and create scanner
	file, err := os.Open("example.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// CONFIGURATION
	scanner := bufio.NewScanner(file)
	keyword := "important"
	lineNumber := 1 // Counter for matched lines only

	fmt.Printf("Filtering for keyword: '%s'\n", keyword)
	fmt.Println("Replacing with: 'necessary'")
	fmt.Println(strings.Repeat("─", 60))

	// FILTERING LOOP
	for scanner.Scan() {
		line := scanner.Text() // Step 1: READ the current line

		// Step 2: CHECK if the line contains the keyword
		if strings.Contains(line, keyword) {
			// Step 3: PROCESS (transform and output)
			updatedLine := strings.ReplaceAll(line, keyword, "necessary")
			fmt.Printf("Line %d: %s\n", lineNumber, updatedLine)
			lineNumber++
		}
		// If the condition is false, the loop implicitly continues to the next line
		// This is equivalent to "skipping" or "discarding" lines that don't match
	}

	// ERROR HANDLING: Check for non-EOF errors
	if err := scanner.Err(); err != nil {
		fmt.Println("Error scanning file:", err)
	}
}

/*
━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━
  EXAMPLE 2: MULTI-CRITERIA FILTERING
━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━

Filter lines using multiple criteria with logical operators (AND, OR, NOT).

PROCESS FLOW:
  • Read each line
  • Check multiple conditions simultaneously:
    - Contains "ERROR" (required)
    - Does NOT contain "DEBUG" (exclusion)
    - Length > 10 characters (constraint)
  • If all conditions are met: process the line
  • Otherwise: skip it

This demonstrates more sophisticated filtering logic used in log analysis,
where you might want to find errors that aren't debug messages, for example.
━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━
*/

func Example2_MultiCriteriaFiltering() {
	fmt.Println("\n=== EXAMPLE 2: Multi-Criteria Filtering ===\n")

	file, err := os.Open("example.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	lineCount := 0
	matchedLines := 0

	fmt.Println("Filtering criteria:")
	fmt.Println("  ✓ Must contain: 'ERROR'")
	fmt.Println("  ✗ Must NOT contain: 'DEBUG'")
	fmt.Println("  ✓ Length must exceed: 10 characters")
	fmt.Println(strings.Repeat("─", 60))

	for scanner.Scan() {
		line := scanner.Text()
		lineCount++

		// MULTI-CRITERIA CHECK: All conditions must be true (AND logic)
		containsError := strings.Contains(line, "ERROR")
		noDebugFlag := !strings.Contains(line, "DEBUG") // NOT operator
		sufficientLength := len(line) > 10

		if containsError && noDebugFlag && sufficientLength {
			fmt.Printf("[Line %d] %s\n", lineCount, line)
			matchedLines++
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error scanning file:", err)
	}

	fmt.Printf("\nMatched %d out of %d lines\n", matchedLines, lineCount)
}

/*
━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━
  EXAMPLE 3: FILTERED OUTPUT TO FILE
━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━

Read from one file, filter lines, and write the results to a new file.

PROCESS FLOW:
  • Open source file for reading
  • Open destination file for writing
  • For each line:
    - Check if it matches criteria
    - If yes: write to destination file
    - If no: skip it
  • Close both files

This pattern is essential for data cleaning pipelines and log extraction.
━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━
*/

func Example3_FilteredOutputToFile() {
	fmt.Println("\n=== EXAMPLE 3: Filtered Output to File ===\n")

	// STEP 1: OPEN SOURCE FILE FOR READING
	sourceFile, err := os.Open("example.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer sourceFile.Close()

	// STEP 2: OPEN DESTINATION FILE FOR WRITING
	outputFile, err := os.Create("filtered_output.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer outputFile.Close()

	scanner := bufio.NewScanner(sourceFile)
	writer := bufio.NewWriter(outputFile) // Buffered writer for efficiency
	lineCount := 0
	writtenLines := 0

	fmt.Println("Reading from: example.txt")
	fmt.Println("Writing to: filtered_output.txt")
	fmt.Println("Criteria: Lines containing 'CRITICAL'")
	fmt.Println(strings.Repeat("─", 60))

	// FILTERING LOOP: Read, check, and write
	for scanner.Scan() {
		line := scanner.Text()
		lineCount++

		// Check if line matches our criteria
		if strings.Contains(line, "CRITICAL") {
			// Write the line to the output file (with newline)
			_, err := writer.WriteString(line + "\n")
			if err != nil {
				log.Fatal(err)
			}
			writtenLines++
		}
	}

	// STEP 3: FLUSH BUFFERED WRITER
	// Important: Always flush to ensure all data is written to disk
	err = writer.Flush()
	if err != nil {
		log.Fatal(err)
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error scanning file:", err)
	}

	fmt.Printf("\nProcessed %d lines, wrote %d matching lines\n", lineCount, writtenLines)
	fmt.Println("✓ Results saved to filtered_output.txt")
}

/*
━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━
  EXAMPLE 4: PATTERN MATCHING WITH STRINGS.FIELDS
━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━

Filter lines by analyzing individual words (fields) within each line.

PROCESS FLOW:
  • Read each line
  • Split line into fields (words) using strings.Fields()
  • Check if ANY field matches a keyword
  • If yes: process the line
  • If no: skip it

This is useful for:
  • Extracting lines containing specific terms (regardless of position)
  • Processing structured data (CSV, space-separated values)
  • Log analysis where keywords can appear at any position
━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━
*/

func Example4_FieldBasedFiltering() {
	fmt.Println("\n=== EXAMPLE 4: Pattern Matching with strings.Fields ===\n")

	file, err := os.Open("example.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	targetKeyword := "FAIL"
	matchedLines := 0

	fmt.Printf("Looking for lines containing the field: '%s'\n", targetKeyword)
	fmt.Println(strings.Repeat("─", 60))

	for scanner.Scan() {
		line := scanner.Text()

		// Split the line into individual fields (words)
		fields := strings.Fields(line)

		// Check if ANY field matches the keyword
		for _, field := range fields {
			if field == targetKeyword {
				fmt.Printf("✓ %s\n", line)
				matchedLines++
				break // Break inner loop; we've found a match for this line
			}
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error scanning file:", err)
	}

	fmt.Printf("\nFound %d matching lines\n", matchedLines)
}

/*
━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━
  EXAMPLE 5: COMPLEX TRANSFORMATIONS
━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━

Filter AND transform lines: normalize whitespace, convert to uppercase, etc.

PROCESS FLOW:
  • Read each line
  • Apply multiple transformations:
    1. Trim leading/trailing whitespace
    2. Convert to uppercase
    3. Replace multiple spaces with single space
  • Check if transformed line meets criteria
  • Output the transformed line
  • Track statistics (original vs transformed)

This pattern is essential for data normalization and preprocessing.
━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━
*/

func Example5_ComplexTransformations() {
	fmt.Println("\n=== EXAMPLE 5: Complex Transformations ===\n")

	file, err := os.Open("example.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	totalLines := 0
	transformedLines := 0

	fmt.Println("Applying transformations:")
	fmt.Println("  1. Trim whitespace")
	fmt.Println("  2. Convert to uppercase")
	fmt.Println("  3. Normalize internal spaces")
	fmt.Println("  4. Filter lines with length > 5 characters")
	fmt.Println(strings.Repeat("─", 60))

	for scanner.Scan() {
		line := scanner.Text()
		totalLines++

		// STEP 1: Trim leading and trailing whitespace
		normalized := strings.TrimSpace(line)

		// STEP 2: Convert to uppercase
		normalized = strings.ToUpper(normalized)

		// STEP 3: Normalize internal spaces (replace multiple spaces with single space)
		normalizedFields := strings.Fields(normalized)   // Splits on any whitespace
		normalized = strings.Join(normalizedFields, " ") // Joins with single space

		// STEP 4: Filter based on transformed content
		if len(normalized) > 5 {
			fmt.Printf("BEFORE: %q\n", line)
			fmt.Printf("AFTER:  %q\n\n", normalized)
			transformedLines++
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error scanning file:", err)
	}

	fmt.Printf("Processed %d lines, output %d lines (%.1f%% pass rate)\n",
		totalLines, transformedLines, float64(transformedLines)/float64(totalLines)*100)
}

/*
━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━
  EXAMPLE 6: STATISTICS AND AGGREGATION
━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━

Filter lines while collecting statistics about the filtered data.

PROCESS FLOW:
  • Initialize counters and aggregators
  • For each line:
    - Check if it matches criteria
    - If yes: collect statistics (count, length sum, specific patterns found)
    - If no: skip it
  • After loop: calculate and display aggregated statistics

This pattern is vital for log analysis, performance monitoring, and
data quality reports.
━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━
*/

func Example6_StatisticsAndAggregation() {
	fmt.Println("\n=== EXAMPLE 6: Statistics and Aggregation ===\n")

	file, err := os.Open("example.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	// INITIALIZE STATISTICS TRACKERS
	totalLines := 0
	matchedLines := 0
	totalMatchedLength := 0
	shortestMatch := -1
	longestMatch := 0

	keyword := "error"

	fmt.Printf("Analyzing lines containing: '%s'\n", keyword)
	fmt.Println(strings.Repeat("─", 60))

	for scanner.Scan() {
		line := scanner.Text()
		totalLines++

		if strings.Contains(strings.ToLower(line), keyword) {
			matchedLines++
			lineLength := len(line)
			totalMatchedLength += lineLength

			if shortestMatch == -1 || lineLength < shortestMatch {
				shortestMatch = lineLength
			}
			if lineLength > longestMatch {
				longestMatch = lineLength
			}
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error scanning file:", err)
	}

	// CALCULATE AVERAGES
	averageLength := 0.0
	if matchedLines > 0 {
		averageLength = float64(totalMatchedLength) / float64(matchedLines)
	}

	// DISPLAY STATISTICS
	fmt.Println("\nFILTERING STATISTICS:")
	fmt.Printf("  Total lines analyzed:     %d\n", totalLines)
	fmt.Printf("  Matched lines:            %d\n", matchedLines)
	fmt.Printf("  Match percentage:         %.1f%%\n", float64(matchedLines)/float64(totalLines)*100)
	fmt.Printf("  Shortest matched line:    %d characters\n", shortestMatch)
	fmt.Printf("  Longest matched line:     %d characters\n", longestMatch)
	fmt.Printf("  Average matched length:   %.1f characters\n", averageLength)
}

/*
═══════════════════════════════════════════════════════════════════════════════
                    KEY CONCEPTS & BEST PRACTICES
═══════════════════════════════════════════════════════════════════════════════

1. FILTERING vs. SKIPPING
   The "filtering" happens implicitly through conditional logic. When a condition
   is false, the loop naturally continues to the next line—effectively "dropping"
   or "discarding" that line.

2. BUFFERED I/O FOR PERFORMANCE
   Always use bufio.Scanner for line-by-line reading. For writing, use
   bufio.Writer. This reduces the number of expensive disk interactions,
   which is crucial when processing large files (e.g., gigabyte-sized logs).

   Example performance improvement:
   • Unbuffered: Each line = 1 system call = slow
   • Buffered (bufio): ~4KB chunks = fewer system calls = fast

3. ERROR HANDLING PATTERN
   After the scanning loop exits (either at EOF or due to error), ALWAYS call
   scanner.Err() to check for non-EOF errors. The scanner doesn't panic; it
   sets an internal error flag.

   for scanner.Scan() { ... }  // Loop until EOF or error
   if err := scanner.Err(); err != nil {
       // Handle the error
   }

4. MEMORY EFFICIENCY
   scanner.Text() returns a new string each iteration. This is acceptable for
   line filtering because lines are typically small. For extremely large lines,
   use scanner.Bytes() which doesn't allocate a new string.

5. MULTI-CRITERIA FILTERING
   Combine conditions using boolean operators:
   • AND: All must be true       → if cond1 && cond2 && cond3 { ... }
   • OR:  At least one is true   → if cond1 || cond2 || cond3 { ... }
   • NOT: Condition is false     → if !condition { ... }

6. TRANSFORMATION PIPELINE
   Apply transformations in order: normalize → validate → transform → output
   Each step builds on the previous, allowing incremental improvements to data.

7. WRITER FLUSHING
   When using bufio.Writer, ALWAYS call Flush() before closing the file to
   ensure all buffered data is written to disk. Forgetting this is a common
   source of data loss bugs.

═══════════════════════════════════════════════════════════════════════════════
*/

func main() {
	fmt.Println("\n" + strings.Repeat("═", 80))
	fmt.Println("                 GO LINE FILTERING - COMPREHENSIVE GUIDE")
	fmt.Println(strings.Repeat("═", 80))

	// NOTE: These examples require "example.txt" to exist in the same directory
	// For demonstration, only show the structure (don't actually run to avoid file errors)

	fmt.Println("\n📖 EXAMPLES STRUCTURE:")
	fmt.Println("  1. Basic Filtering & Transformation")
	fmt.Println("  2. Multi-Criteria Filtering (AND, OR, NOT logic)")
	fmt.Println("  3. Filtered Output to File (data pipeline)")
	fmt.Println("  4. Field-Based Filtering (word-level analysis)")
	fmt.Println("  5. Complex Transformations (normalization)")
	fmt.Println("  6. Statistics & Aggregation (reporting)")

	fmt.Println("\n💡 KEY TAKEAWAY:")
	fmt.Println("  Line filtering = Read → Check → Process (or Skip)")
	fmt.Println("  Always use bufio.Scanner for efficiency on large files!")

	fmt.Println("\n" + strings.Repeat("═", 80) + "\n")
}
