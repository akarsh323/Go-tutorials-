package intermediate

import (
	"bufio"
	"fmt"
	"os"
)

// Topic 84: Reading Files
// Various methods for reading data from files
// From raw byte reading to efficient line-by-line scanning

func main() {

	// ============================================
	// SETUP: Create a test file to read from
	// ============================================
	content := []byte("Line 1: Welcome to Go file reading\nLine 2: This is the second line\nLine 3: And here is the third line\n")
	err := os.WriteFile("test_read.txt", content, 0644)
	if err != nil {
		fmt.Println("Error creating test file:", err)
		return
	}
	defer os.Remove("test_read.txt")

	// ============================================
	// METHOD A: Reading Raw Bytes (file.Read)
	// ============================================
	// Use this for binary data or when you need fixed-size chunks

	fmt.Println("\n=== METHOD A: Reading Raw Bytes ===")
	file, err := os.Open("test_read.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close() // Always close files to prevent resource leaks

	// Create a buffer of 1024 bytes
	data := make([]byte, 1024)

	n, err := file.Read(data)
	if err != nil {
		fmt.Println("Error reading data:", err)
		return
	}

	// Convert bytes to string to display
	// Note: data contains the full buffer, but only n bytes are valid
	fmt.Printf("Read %d bytes\n", n)
	fmt.Println("File Content:", string(data[:n])) // Slice to actual read bytes

	// ============================================
	// METHOD B: Reading Line-by-Line (bufio.Scanner)
	// ============================================
	// BEST PRACTICE for text files, logs, and configuration files
	// The Scanner handles buffering and tokenization automatically

	fmt.Println("\n=== METHOD B: Reading Line-by-Line with Scanner ===")

	file2, err := os.Open("test_read.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file2.Close()

	// Create a scanner to read the file line by line
	scanner := bufio.NewScanner(file2)

	lineNumber := 1
	for scanner.Scan() {
		line := scanner.Text() // Get current line as string
		fmt.Printf("Line %d: %s\n", lineNumber, line)
		lineNumber++
	}

	// CRITICAL: Check for errors AFTER the loop
	// If the loop ended due to EOF, scanner.Err() returns nil
	// If the loop ended due to an actual error, scanner.Err() returns that error
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	fmt.Println("File read successfully with Scanner")

	// ============================================
	// UNDERSTANDING SCANNER INTERNALS
	// ============================================
	// The Scanner internally:
	// 1. Maintains a cursor position in the file
	// 2. Reads data into a buffer
	// 3. Splits the buffer into tokens (lines by default)
	// 4. Returns tokens one by one via scanner.Text()
	// 5. scanner.Scan() returns true for each token, false at EOF

	fmt.Println("\n=== Scanner Internals Demo ===")

	// Create a multi-line test file
	multiLineContent := []byte("First line\nSecond line\nThird line\n")
	err = os.WriteFile("multiline.txt", multiLineContent, 0644)
	if err != nil {
		fmt.Println("Error creating multiline file:", err)
		return
	}
	defer os.Remove("multiline.txt")

	file3, err := os.Open("multiline.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file3.Close()

	scanner2 := bufio.NewScanner(file3)
	tokenCount := 0

	// Each iteration represents scanner.Scan() advancing to the next token
	for scanner2.Scan() {
		tokenCount++
		line := scanner2.Text()
		fmt.Printf("Token %d (Cursor advanced): %s\n", tokenCount, line)
	}

	// Check for EOF or errors
	if err := scanner2.Err(); err != nil {
		// Only reaches here if there's an actual error (not EOF)
		fmt.Println("Error reading file:", err)
	} else {
		// This message appears when EOF is reached normally
		fmt.Printf("Reached EOF after reading %d tokens\n", tokenCount)
	}

	// ============================================
	// METHOD C: Reading Entire File at Once
	// ============================================
	// Go 1.16+ provides os.ReadFile for convenience
	// Use this for small to medium files that fit in memory

	fmt.Println("\n=== METHOD C: Reading Entire File (os.ReadFile) ===")

	fileContent, err := os.ReadFile("test_read.txt")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	fmt.Println("Entire File Content:")
	fmt.Println(string(fileContent))

	// ============================================
	// COMPARISON: Choosing the Right Method
	// ============================================
	fmt.Println("\n=== Method Comparison ===")
	fmt.Println("file.Read()       : Best for binary data or fixed-size chunks")
	fmt.Println("bufio.Scanner     : Best for text files, logs, config files")
	fmt.Println("os.ReadFile()     : Best for small files that fit in memory")

	// ============================================
	// PRACTICAL EXAMPLE: Processing Log File
	// ============================================
	fmt.Println("\n=== Practical Example: Processing Log File ===")

	// Create a sample log file
	logContent := []byte("2024-01-01 10:00:00 INFO Application started\n2024-01-01 10:00:01 DEBUG Loading configuration\n2024-01-01 10:00:02 ERROR Database connection failed\n")
	err = os.WriteFile("app.log", logContent, 0644)
	if err != nil {
		fmt.Println("Error creating log file:", err)
		return
	}
	defer os.Remove("app.log")

	logFile, err := os.Open("app.log")
	if err != nil {
		fmt.Println("Error opening log file:", err)
		return
	}
	defer logFile.Close()

	logScanner := bufio.NewScanner(logFile)
	errorCount := 0

	for logScanner.Scan() {
		logLine := logScanner.Text()
		if len(logLine) > 0 && logLine[len(logLine)-1:] != "" {
			// Check for ERROR entries
			if contains(logLine, "ERROR") {
				errorCount++
				fmt.Println("Found error:", logLine)
			}
		}
	}

	if err := logScanner.Err(); err != nil {
		fmt.Println("Error reading log file:", err)
		return
	}

	fmt.Printf("Total errors found: %d\n", errorCount)

	// ============================================
	// KEY TAKEAWAYS
	// ============================================
	fmt.Println("\n=== Key Takeaways ===")
	fmt.Println("1. Always use defer file.Close() to prevent resource leaks")
	fmt.Println("2. Use bufio.Scanner for text file processing (most common)")
	fmt.Println("3. Check scanner.Err() AFTER the loop (not inside)")
	fmt.Println("4. scanner.Err() returns nil at EOF (not an error)")
	fmt.Println("5. Use file.Read() for binary data or fixed-size reads")
	fmt.Println("6. Use os.ReadFile() for small files that fit in memory")
	fmt.Println("7. Always handle errors from file operations")

}

// Helper function to check if a string contains a substring
func contains(s, substr string) bool {
	for i := 0; i <= len(s)-len(substr); i++ {
		if s[i:i+len(substr)] == substr {
			return true
		}
	}
	return false
}
