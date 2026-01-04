package intermediate

import (
	"bufio"
	"fmt"
	"os"
)

// Topic 83: Writing Files
// Various methods for writing data to files
// Best practices for file writing in Go

func main() {

	// Using os.WriteFile (Go 1.16+)
	// This is the simplest and most common way
	content := []byte("Hello, World!\nThis is written to a file.\n")
	err := os.WriteFile("output.txt", content, 0644)
	if err != nil {
		fmt.Println("Error writing file:", err)
		return
	}
	fmt.Println("File written successfully with os.WriteFile")
	defer os.Remove("output.txt")

	// Using os.Create
	file, err := os.Create("example.txt")
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer file.Close()

	// Write to file
	_, err = file.WriteString("Hello from Create\n")
	if err != nil {
		fmt.Println("Error writing:", err)
		return
	}
	fmt.Println("File written successfully with os.Create")
	defer os.Remove("example.txt")

	// Writing bytes
	data := []byte("Binary data\n")
	_, err = file.Write(data)
	if err != nil {
		fmt.Println("Error writing bytes:", err)
		return
	}

	// Appending to file
	file2, err := os.OpenFile("append.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file2.Close()

	_, err = file2.WriteString("Appended line 1\n")
	if err != nil {
		fmt.Println("Error appending:", err)
		return
	}

	_, err = file2.WriteString("Appended line 2\n")
	if err != nil {
		fmt.Println("Error appending:", err)
		return
	}
	fmt.Println("File appended successfully")
	defer os.Remove("append.txt")

	// Formatting and writing
	file3, err := os.Create("formatted.txt")
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer file3.Close()

	fmt.Fprintf(file3, "Name: %s\n", "John")
	fmt.Fprintf(file3, "Age: %d\n", 25)
	fmt.Fprintf(file3, "Score: %.2f\n", 95.5)
	fmt.Println("Formatted file written successfully")
	defer os.Remove("formatted.txt")

	// ============================================
	// BUFFERED I/O - BEST PRACTICE FOR PERFORMANCE
	// ============================================
	// For writing large amounts of data, use bufio.Writer
	// It batches writes in memory before flushing to disk

	file4, err := os.Create("buffered.txt")
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer file4.Close()

	// Create a buffered writer wrapping the file
	writer := bufio.NewWriter(file4)

	// Write multiple lines efficiently
	for i := 1; i <= 5; i++ {
		_, err := writer.WriteString(fmt.Sprintf("Line %d: This is buffered output\n", i))
		if err != nil {
			fmt.Println("Error writing:", err)
			return
		}
	}

	// CRITICAL: Flush the buffer to ensure all data is written to disk
	err = writer.Flush()
	if err != nil {
		fmt.Println("Error flushing buffer:", err)
		return
	}
	fmt.Println("Buffered file written successfully")
	defer os.Remove("buffered.txt")

	// ============================================
	// BUFFERED I/O WITH LARGE DATASET
	// ============================================
	// Demonstrating performance benefits of buffering

	file5, err := os.Create("large_dataset.txt")
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer file5.Close()

	writer2 := bufio.NewWriter(file5)

	// Write 100 lines with buffering (efficient)
	for i := 1; i <= 100; i++ {
		_, err := writer2.WriteString(fmt.Sprintf("Record %d: Data entry\n", i))
		if err != nil {
			fmt.Println("Error writing:", err)
			return
		}
	}

	// Always flush at the end
	err = writer2.Flush()
	if err != nil {
		fmt.Println("Error flushing buffer:", err)
		return
	}
	fmt.Println("Large dataset written successfully with buffering")
	defer os.Remove("large_dataset.txt")

	// ============================================
	// os.WriteFile - CONVENIENT FOR SMALL FILES
	// ============================================
	// Use os.WriteFile for small files (Go 1.16+)
	// This is a one-liner that handles everything

	content := []byte("Hello, World!\nThis is a simple file.\nCreated with os.WriteFile\n")
	err = os.WriteFile("simple.txt", content, 0644)
	if err != nil {
		fmt.Println("Error writing file:", err)
		return
	}
	fmt.Println("Simple file written successfully with os.WriteFile")
	defer os.Remove("simple.txt")

	// ============================================
	// KEY TAKEAWAYS
	// ============================================
	// 1. Always use defer file.Close() to prevent resource leaks
	// 2. Use bufio.Writer for writing large amounts of data
	// 3. Don't forget to call writer.Flush() when using buffered I/O
	// 4. Use os.WriteFile for quick, small file writes
	// 5. Always check and handle errors from file operations
	// 6. Explicitly add \n for newlines - Go won't do it automatically

}
