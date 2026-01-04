package intermediate

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Topic 80: Buffered I/O (bufio) - Efficient Data Handling
// ========================================================
// This lesson covers the bufio package for efficient reading and writing.
// We learn about buffering concepts, Reader/Writer patterns,
// Scanners, and critical best practices like Flush().

func main() {
	fmt.Println("=== Topic 80: Buffered I/O (bufio) - Efficient Data Handling ===\n")

	lesson1BufferingConcept()
	fmt.Println("\n" + string([]byte{61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 61}) + "\n")

	lesson2ReadingData()
	fmt.Println("\n" + string([]byte{61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 61}) + "\n")

	lesson3WritingData()
	fmt.Println("\n" + string([]byte{61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 61}) + "\n")

	lesson4ScannerPattern()
	fmt.Println("\n" + string([]byte{61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 61}) + "\n")

	lesson5FlushingExplained()
	fmt.Println("\n" + string([]byte{61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 61}) + "\n")

	lesson6BestPractices()
}

// LESSON 1: The Buffering Concept
// ===============================
func lesson1BufferingConcept() {
	fmt.Println("LESSON 1: THE BUFFERING CONCEPT")
	fmt.Println("-------------------------------\n")

	fmt.Println("WHAT IS bufio?")
	fmt.Println("  bufio = Buffered Input/Output")
	fmt.Println("  It wraps around io.Reader and io.Writer")
	fmt.Println("  to improve performance and add convenience.\n")

	fmt.Println("THE PROBLEM WITHOUT BUFFERING:")
	fmt.Println("  Reading data directly from disk/network is SLOW")
	fmt.Println("  • Each read = system call (expensive)")
	fmt.Println("  • Reading 1 byte = 1 system call")
	fmt.Println("  • Reading 1 million bytes = 1 million system calls")
	fmt.Println("  • This is extremely inefficient!\n")

	fmt.Println("THE SOLUTION: BUFFERING")
	fmt.Println("  Instead of reading 1 byte at a time,")
	fmt.Println("  read in chunks (buffers) of multiple bytes")
	fmt.Println("  • 1 system call = read 4KB from disk")
	fmt.Println("  • Your program uses those bytes from RAM")
	fmt.Println("  • Much faster and more efficient!\n")

	fmt.Println("THE STREAMING ANALOGY (Netflix/YouTube):")
	fmt.Println("┌─────────────────────────────────────────────────┐")
	fmt.Println("│ WITHOUT BUFFERING (Bad User Experience)          │")
	fmt.Println("│  Download entire 2-hour 4K movie file            │")
	fmt.Println("│  ↓")
	fmt.Println("│  Wait hours...                                   │")
	fmt.Println("│  ↓")
	fmt.Println("│  Finally watch it                                │")
	fmt.Println("└─────────────────────────────────────────────────┘\n")

	fmt.Println("┌─────────────────────────────────────────────────┐")
	fmt.Println("│ WITH BUFFERING (Good User Experience)            │")
	fmt.Println("│  Download next 10 seconds of video               │")
	fmt.Println("│  ↓")
	fmt.Println("│  Watch current 10 seconds WHILE next downloads   │")
	fmt.Println("│  ↓")
	fmt.Println("│  Seamless viewing experience                     │")
	fmt.Println("└─────────────────────────────────────────────────┘\n")

	fmt.Println("IN GO:")
	fmt.Println("  • bufio.Reader: Reads data in chunks into memory")
	fmt.Println("  • bufio.Writer: Collects data in memory, sends in chunks")
	fmt.Println("  • bufio.Scanner: Even more convenient (reads line-by-line)\n")

	fmt.Println("PERFORMANCE BENEFIT:")
	fmt.Println("  Without bufio: 1,000,000 system calls for 1MB file")
	fmt.Println("  With bufio:    ~250 system calls for 1MB file")
	fmt.Println("  → Can be 100-1000x faster!")
}

// LESSON 2: Reading Data with bufio.Reader
// ========================================
func lesson2ReadingData() {
	fmt.Println("LESSON 2: READING DATA WITH bufio.Reader")
	fmt.Println("----------------------------------------\n")

	fmt.Println("THE CONCEPT:")
	fmt.Println("  bufio.Reader wraps any io.Reader")
	fmt.Println("  It adds methods for convenient reading patterns.\n")

	fmt.Println("CREATING A READER:")
	fmt.Println("  reader := bufio.NewReader(source)")
	fmt.Println("  source can be:")
	fmt.Println("    • strings.NewReader(\"text\")")
	fmt.Println("    • os.Stdin (keyboard input)")
	fmt.Println("    • os.Open(\"file.txt\") (file handle)")
	fmt.Println("    • net.Dial() (network connection)\n")

	fmt.Println("EXAMPLE 1: Basic Setup")
	source := strings.NewReader("Hello bufio package\nSecond line\n")
	reader := bufio.NewReader(source)
	fmt.Println("  source := strings.NewReader(\"Hello bufio package\\nSecond line\\n\")")
	fmt.Println("  reader := bufio.NewReader(source)\n")

	fmt.Println("METHOD A: Read() - Fill a byte buffer")
	fmt.Println("  Syntax: n, err := reader.Read(buffer)")
	fmt.Println("  Returns: (bytes read, error)")
	fmt.Println("  Use: When you want to control memory with a fixed-size buffer\n")

	fmt.Println("  Example:")
	buffer := make([]byte, 10) // Can read up to 10 bytes
	n, err := reader.Read(buffer)
	if err != nil {
		fmt.Printf("  Error: %v\n\n", err)
	} else {
		fmt.Printf("  Read %d bytes: \"%s\"\n\n", n, string(buffer[:n]))
	}

	fmt.Println("METHOD B: ReadString() - Read until delimiter")
	fmt.Println("  Syntax: line, err := reader.ReadString('\\n')")
	fmt.Println("  Returns: (string up to delimiter, error)")
	fmt.Println("  Use: When you want to read until a specific character\n")

	fmt.Println("  Example:")
	source2 := strings.NewReader("First line\nSecond line\nThird line\n")
	reader2 := bufio.NewReader(source2)

	for i := 1; i <= 3; i++ {
		line, err := reader2.ReadString('\n')
		if err != nil {
			break
		}
		fmt.Printf("  Line %d: \"%s\"\n", i, strings.TrimSpace(line))
	}
	fmt.Println()

	fmt.Println("METHOD C: ReadLine() - Read until newline (more complex)")
	fmt.Println("  Syntax: line, isPrefix, err := reader.ReadLine()")
	fmt.Println("  Less common - usually use ReadString or Scanner instead\n")

	fmt.Println("KEY INSIGHT:")
	fmt.Println("  ✓ Read() for fixed-size chunks")
	fmt.Println("  ✓ ReadString() for delimiter-based reading")
	fmt.Println("  ✓ Most people use Scanner (coming next) - simpler and safer")
}

// LESSON 3: Writing Data with bufio.Writer
// ========================================
func lesson3WritingData() {
	fmt.Println("LESSON 3: WRITING DATA WITH bufio.Writer")
	fmt.Println("---------------------------------------\n")

	fmt.Println("THE CONCEPT:")
	fmt.Println("  bufio.Writer buffers data in memory")
	fmt.Println("  Then sends it to the destination in chunks\n")

	fmt.Println("CREATING A WRITER:")
	fmt.Println("  writer := bufio.NewWriter(destination)")
	fmt.Println("  destination can be:")
	fmt.Println("    • os.Stdout (screen terminal)")
	fmt.Println("    • os.Stderr (error output)")
	fmt.Println("    • os.Create(\"file.txt\") (file handle)")
	fmt.Println("    • net connection\n")

	fmt.Println("WRITING METHODS:")
	fmt.Println("  1. Write([]byte)     - Write raw bytes")
	fmt.Println("  2. WriteString(str)  - Write a string (simpler, recommended)\n")

	fmt.Println("EXAMPLE 1: Using WriteString()")
	fmt.Println("  writer := bufio.NewWriter(os.Stdout)")
	fmt.Println("  writer.WriteString(\"Hello World\\n\")")
	fmt.Println("  writer.Flush() // Critical!\n")

	// Demonstrate with actual output
	fmt.Println("  Output: (This message prints below)")
	writer := bufio.NewWriter(os.Stdout)
	writer.WriteString("  ✓ This text was buffered and flushed!\n")
	writer.Flush()
	fmt.Println()

	fmt.Println("EXAMPLE 2: Multiple writes before flush")
	fmt.Println("  (All data sits in memory until Flush())\n")

	writer2 := bufio.NewWriter(os.Stdout)
	writer2.WriteString("  Line 1\n")
	writer2.WriteString("  Line 2\n")
	writer2.WriteString("  Line 3\n")
	fmt.Println("  (Before Flush - nothing on screen yet)")
	writer2.Flush()
	fmt.Println("  (After Flush - all lines appear)\n")

	fmt.Println("KEY INSIGHT:")
	fmt.Println("  ✓ WriteString is easier than Write() (no type conversion)")
	fmt.Println("  ✓ Data stays in RAM until you Flush()")
	fmt.Println("  ✓ Multiple writes are efficient (batch operation)")
}

// LESSON 4: The Scanner Pattern (Easiest)
// =======================================
func lesson4ScannerPattern() {
	fmt.Println("LESSON 4: THE SCANNER PATTERN (EASIEST)")
	fmt.Println("--------------------------------------\n")

	fmt.Println("PROBLEM WITH Read() and ReadString():")
	fmt.Println("  • Need error handling after each call")
	fmt.Println("  • Manual buffering management")
	fmt.Println("  • Easy to make mistakes\n")

	fmt.Println("SOLUTION: bufio.Scanner")
	fmt.Println("  Higher-level abstraction for reading")
	fmt.Println("  Automatically handles buffering and errors\n")

	fmt.Println("SYNTAX:")
	fmt.Println("  scanner := bufio.NewScanner(reader)")
	fmt.Println("  for scanner.Scan() {")
	fmt.Println("      line := scanner.Text()")
	fmt.Println("      // Process line")
	fmt.Println("  }\n")

	fmt.Println("EXAMPLE: Reading lines\n")

	input := "apple\nbanana\ncherry\ndate\n"
	fmt.Printf("Input text:\n%s\n", input)

	fmt.Println("Code:")
	fmt.Println("  scanner := bufio.NewScanner(strings.NewReader(input))")
	fmt.Println("  for scanner.Scan() {")
	fmt.Println("      fmt.Println(scanner.Text())")
	fmt.Println("  }\n")

	fmt.Println("Output:")
	scanner := bufio.NewScanner(strings.NewReader(input))
	for scanner.Scan() {
		fmt.Printf("  %s\n", scanner.Text())
	}
	fmt.Println()

	fmt.Println("WHY SCANNER IS BETTER:")
	fmt.Println("  ✓ Cleaner syntax (for loop)")
	fmt.Println("  ✓ Automatic error handling")
	fmt.Println("  ✓ Default behavior (line-by-line) is most common")
	fmt.Println("  ✓ Can customize split function if needed\n")

	fmt.Println("CUSTOM SPLIT FUNCTION:")
	fmt.Println("  By default, Scanner splits by lines")
	fmt.Println("  You can change this with .Split():\n")

	fmt.Println("  scanner.Split(bufio.ScanWords)  // Split by whitespace")
	fmt.Println("  scanner.Split(bufio.ScanBytes)  // Split by individual bytes")
	fmt.Println("  scanner.Split(customFunc)       // Custom split logic\n")

	fmt.Println("EXAMPLE: Scanning words")
	words := "apple banana cherry date"
	fmt.Printf("Input: \"%s\"\n", words)
	fmt.Println("Code:")
	fmt.Println("  scanner := bufio.NewScanner(strings.NewReader(words))")
	fmt.Println("  scanner.Split(bufio.ScanWords)")
	fmt.Println("  for scanner.Scan() { fmt.Println(scanner.Text()) }\n")

	fmt.Println("Output:")
	wordScanner := bufio.NewScanner(strings.NewReader(words))
	wordScanner.Split(bufio.ScanWords)
	for wordScanner.Scan() {
		fmt.Printf("  %s\n", wordScanner.Text())
	}
}

// LESSON 5: Flushing Explained - The Critical Concept
// ==================================================
func lesson5FlushingExplained() {
	fmt.Println("LESSON 5: FLUSHING EXPLAINED - THE CRITICAL CONCEPT")
	fmt.Println("--------------------------------------------------\n")

	fmt.Println("THE PROBLEM:")
	fmt.Println("  When you call writer.WriteString(), data goes to:")
	fmt.Println("    ✗ NOT the screen/file immediately")
	fmt.Println("    ✓ YES the bufio internal memory (RAM)\n")

	fmt.Println("  If you forget to Flush() and your program exits:")
	fmt.Println("    • Data in RAM is lost")
	fmt.Println("    • Nothing appears on screen or in file\n")

	fmt.Println("THE SOLUTION: Flush()")
	fmt.Println("  writer.Flush() pushes buffered data to destination\n")

	fmt.Println("VISUALIZATION:\n")

	fmt.Println("Step 1: WriteString (data in RAM):")
	fmt.Println("  writer.WriteString(\"Hello\")")
	fmt.Println("  ┌──────────────┐      ┌─────────────┐")
	fmt.Println("  │  Write Call  │ -->  │ RAM Buffer  │")
	fmt.Println("  └──────────────┘      └─────────────┘")
	fmt.Println("                        \"Hello\" sitting here\n")

	fmt.Println("Step 2: Before Flush (still in RAM):")
	fmt.Println("  ┌─────────────┐      ┌──────────────┐")
	fmt.Println("  │ RAM Buffer  │      │ Screen/File  │")
	fmt.Println("  │ \"Hello\"     │  ✗   │ Empty        │")
	fmt.Println("  └─────────────┘      └──────────────┘\n")

	fmt.Println("Step 3: After Flush (data reaches destination):")
	fmt.Println("  writer.Flush()")
	fmt.Println("  ┌─────────────┐      ┌──────────────┐")
	fmt.Println("  │ RAM Buffer  │  --> │ Screen/File  │")
	fmt.Println("  │ Empty       │      │ \"Hello\"      │")
	fmt.Println("  └─────────────┘      └──────────────┘\n")

	fmt.Println("WHEN IS Flush() CALLED AUTOMATICALLY?")
	fmt.Println("  • When buffer is full (typically 4KB)")
	fmt.Println("  • When you explicitly call Flush()")
	fmt.Println("  • When the program exits gracefully\n")

	fmt.Println("BEST PRACTICE: Always Flush explicitly!")
	fmt.Println("  Don't rely on automatic flushing\n")

	fmt.Println("EXAMPLE: Buffer sizes\n")

	fmt.Println("Code that relies on auto-flush (NOT SAFE):")
	fmt.Println("  writer := bufio.NewWriter(os.Stdout)")
	fmt.Println("  writer.WriteString(\"Short message\")")
	fmt.Println("  // Program exits - data might be lost!\n")

	fmt.Println("Code with explicit Flush (SAFE):")
	fmt.Println("  writer := bufio.NewWriter(os.Stdout)")
	fmt.Println("  writer.WriteString(\"Short message\")")
	fmt.Println("  writer.Flush() // Guaranteed delivery\n")

	fmt.Println("CRITICAL RULE:")
	fmt.Println("┌─────────────────────────────────────────────────┐")
	fmt.Println("│ ALWAYS CALL Flush() before program ends!        │")
	fmt.Println("│ Or use defer: defer writer.Flush()              │")
	fmt.Println("└─────────────────────────────────────────────────┘")
}

// LESSON 6: Best Practices and Common Patterns
// ============================================
func lesson6BestPractices() {
	fmt.Println("LESSON 6: BEST PRACTICES AND COMMON PATTERNS")
	fmt.Println("------------------------------------------\n")

	fmt.Println("PRACTICE #1: Always Check Errors")
	fmt.Println("  I/O operations can fail:")
	fmt.Println("    • File not found")
	fmt.Println("    • Permission denied")
	fmt.Println("    • Disk full")
	fmt.Println("    • Network disconnected\n")

	fmt.Println("  Pattern:")
	fmt.Println("    line, err := reader.ReadString('\\n')")
	fmt.Println("    if err != nil {")
	fmt.Println("        log.Fatal(\"Failed to read:\", err)")
	fmt.Println("    }\n")

	fmt.Println("PRACTICE #2: Use defer for Flush()")
	fmt.Println("  Ensures Flush is called even if panic occurs\n")

	fmt.Println("  Pattern:")
	fmt.Println("    writer := bufio.NewWriter(os.Stdout)")
	fmt.Println("    defer writer.Flush() // Always executes at end")
	fmt.Println("    writer.WriteString(\"Data\")\n")

	fmt.Println("PRACTICE #3: Scanner vs Reader")
	fmt.Println("  Use Scanner for:")
	fmt.Println("    ✓ Line-by-line reading")
	fmt.Println("    ✓ Word-by-word reading")
	fmt.Println("    ✓ Simple use cases\n")

	fmt.Println("  Use Reader for:")
	fmt.Println("    ✓ Fixed-size chunks")
	fmt.Println("    ✓ Advanced control")
	fmt.Println("    ✓ Performance-critical code\n")

	fmt.Println("PRACTICE #4: Chain Operations")
	fmt.Println("  Don't store intermediate values if not needed\n")

	fmt.Println("  Example:")
	fmt.Println("    // Instead of:")
	fmt.Println("    reader := bufio.NewReader(os.Stdin)")
	fmt.Println("    line, _ := reader.ReadString('\\n')")
	fmt.Println("")
	fmt.Println("    // Can do:")
	fmt.Println("    scanner := bufio.NewScanner(os.Stdin)")n")
	fmt.Println("    for scanner.Scan() { /* process */ }\n")

	fmt.Println("PRACTICE #5: Performance Characteristics")
	fmt.Println("  Default buffer size: 4KB (4096 bytes)")
	fmt.Println("  • Good for most text files")
	fmt.Println("  • Can customize with bufio.NewReaderSize()\n")

	fmt.Println("REAL-WORLD EXAMPLE: Reading Config File\n")

	configText := "host=localhost\nport=8080\ndebug=true\n"
	fmt.Println("Config file content:")
	fmt.Println(configText)

	fmt.Println("Code:")
	fmt.Println("  reader := bufio.NewReader(strings.NewReader(configText))")
	fmt.Println("  config := make(map[string]string)")
	fmt.Println("  for {")
	fmt.Println("      line, err := reader.ReadString('\\n')")
	fmt.Println("      if err != nil { break }")
	fmt.Println("      parts := strings.Split(strings.TrimSpace(line), \"=\")")
	fmt.Println("      if len(parts) == 2 {")
	fmt.Println("          config[parts[0]] = parts[1]")
	fmt.Println("      }")
	fmt.Println("  }\n")

	fmt.Println("Result:")
	reader := bufio.NewReader(strings.NewReader(configText))
	config := make(map[string]string)
	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			break
		}
		parts := strings.Split(strings.TrimSpace(line), "=")
		if len(parts) == 2 {
			config[parts[0]] = parts[1]
		}
	}

	for key, value := range config {
		fmt.Printf("  %s: %s\n", key, value)
	}
	fmt.Println()

	fmt.Println("SUMMARY CHECKLIST:")
	fmt.Println("  ☐ Always use bufio for I/O (performance)")
	fmt.Println("  ☐ Use Scanner for line/word reading (easiest)")
	fmt.Println("  ☐ Use Reader for advanced control (harder)")
	fmt.Println("  ☐ Always call Flush() before program ends")
	fmt.Println("  ☐ Use defer Flush() to ensure it runs")
	fmt.Println("  ☐ Always check errors from I/O operations")
	fmt.Println("  ☐ Choose WriteString over Write for text")
}
