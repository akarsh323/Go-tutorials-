package intermediate

import (
	"fmt"
	"os"
)

// Topic 83: Writing Files
// Various methods for writing data to files

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

}
