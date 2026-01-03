package intermediate

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Topic 80: Buffered I/O (bufio)
// Efficient reading and writing with buffers

func main() {

	// Reading with bufio.Scanner
	reader := bufio.NewReader(strings.NewReader("Hello\nWorld\nGo\n"))

	fmt.Println("Reading lines with Scanner:")
	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println("Line:", line)
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error:", err)
	}

	// Reading words
	reader2 := bufio.NewReader(strings.NewReader("Hello World Go"))
	fmt.Println("\nReading words:")
	for {
		word, err := reader2.ReadString(' ')
		if err != nil {
			break
		}
		fmt.Println("Word:", strings.TrimSpace(word))
	}

	// Writing with bufio.Writer
	writer := bufio.NewWriter(os.Stdout)
	fmt.Fprintln(writer, "\nWriting with buffered writer:")
	writer.WriteString("This is buffered.\n")
	writer.WriteString("This will be flushed.\n")
	writer.Flush()

	// Scanning by delimiter
	scanner2 := bufio.NewScanner(strings.NewReader("apple,banana,cherry"))
	scanner2.Split(bufio.ScanLines)
	fmt.Println("\nScanning by custom delimiter:")
	for scanner2.Scan() {
		fmt.Println("Item:", scanner2.Text())
	}

}
