package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
	"time"
)

// ============================================================================
// 94 Input/Output (io Package)
// ============================================================================
//
// The io package is the "Universal Adapter" of Go. It provides basic interfaces
// for I/O primitives.
//
// Why is it important?
// - Consistency: Use the same methods for Files, Networks, Strings, and Buffers.
// - Interoperability: Pass any data source to any data consumer.
//
// Core Interfaces:
// 1. io.Reader (Read method)
// 2. io.Writer (Write method)
// 3. io.Closer (Close method)
//
// ============================================================================

func main() {
	Demo94IO()
}

// ============================================================================
// PART 1: CORE INTERFACES (Reader & Writer)
// ============================================================================

func Demo94_Part1_CoreInterfaces() {
	fmt.Println("\n=== PART 1: CORE INTERFACES (Reader & Writer) ===")
	fmt.Println()

	fmt.Println("📌 1. Reading from a String (strings.Reader):")
	// strings.NewReader implements io.Reader
	reader := strings.NewReader("Hello, Reader Interface!")

	// Create a "bucket" (byte slice) to hold the data we read
	buffer := make([]byte, 8) // Read 8 bytes at a time

	// Loop until the reader is empty
	for {
		n, err := reader.Read(buffer)
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("  Read %d bytes: %q\n", n, buffer[:n])
	}
	fmt.Println()

	fmt.Println("📌 2. Writing to a Buffer (bytes.Buffer):")
	// bytes.Buffer implements io.Writer
	var writer bytes.Buffer

	data := []byte("Hello, Writer Interface!")
	n, err := writer.Write(data)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("  Wrote %d bytes to buffer.\n", n)
	fmt.Printf("  Buffer content: %s\n", writer.String())
	fmt.Println()
}

// ============================================================================
// PART 2: UTILITY FUNCTIONS (io.Copy & io.MultiReader)
// ============================================================================

func Demo94_Part2_Utilities() {
	fmt.Println("\n=== PART 2: UTILITY FUNCTIONS (Copy & MultiReader) ===")
	fmt.Println()

	fmt.Println("📌 io.Copy (The \"Hose\" Connection):")
	fmt.Println("  Transfers data directly from Source (Reader) to Dest (Writer).")

	src := strings.NewReader("Data transferred via io.Copy")
	dst := new(bytes.Buffer)

	// Magic happens here: No loops, no manual buffer management
	written, err := io.Copy(dst, src)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("  Copied %d bytes: \"%s\"\n", written, dst.String())
	fmt.Println()

	fmt.Println("📌 io.MultiReader (The \"Taped Hose\" Connection):")
	fmt.Println("  Concatenates multiple readers into a single stream.")

	r1 := strings.NewReader("Hello ")
	r2 := strings.NewReader("World ")
	r3 := strings.NewReader("from MultiReader!")

	// Treat r1, r2, and r3 as one continuous reader
	multi := io.MultiReader(r1, r2, r3)

	// Copy the combined stream to stdout to prove it works
	fmt.Print("  Result: ")
	io.Copy(os.Stdout, multi)
	fmt.Println()
	fmt.Println()
}

// ============================================================================
// PART 3: THE PIPE (io.Pipe & Concurrency)
// ============================================================================

func Demo94_Part3_PipeAndConcurrency() {
	fmt.Println("\n=== PART 3: THE PIPE (io.Pipe & Concurrency) ===")
	fmt.Println()

	fmt.Println("📌 Concept: The Synchronous Tunnel")
	fmt.Println("  io.Pipe connects a Reader and a Writer synchronously.")
	fmt.Println("  ⚠️  WARNING: Writing blocks until reading happens (and vice-versa).")
	fmt.Println("  👉 Must use a Go Routine (thread) to prevent deadlock.")
	fmt.Println()

	// Create the pipe
	pr, pw := io.Pipe()

	// 1. Start the WRITER in a background Go Routine
	go func() {
		fmt.Println("  [Go Routine] Starting to write to pipe...")
		defer pw.Close() // Important: Close writer to signal EOF to reader

		// Simulate work
		time.Sleep(500 * time.Millisecond)
		pw.Write([]byte("Data from the underground..."))

		time.Sleep(500 * time.Millisecond)
		pw.Write([]byte(" ...piped to the surface!"))

		fmt.Println("  [Go Routine] Finished writing.")
	}()

	// 2. The MAIN thread acts as the READER
	fmt.Println("  [Main Thread] Waiting to read from pipe...")

	buffer := new(bytes.Buffer)
	buffer.ReadFrom(pr) // Blocks here until data arrives or pipe closes

	fmt.Printf("  [Main Thread] Received: \"%s\"\n", buffer.String())
	fmt.Println()
}

// ============================================================================
// PART 4: MEMORY ALLOCATION (Stack vs Heap)
// ============================================================================

func Demo94_Part4_MemoryAllocation() {
	fmt.Println("\n=== PART 4: MEMORY ALLOCATION (Stack vs Heap) ===")
	fmt.Println()

	fmt.Println("📌 1. Stack Allocation (Value):")
	fmt.Println("  var buf bytes.Buffer")
	var stackBuf bytes.Buffer
	stackBuf.WriteString("I am on the stack")
	fmt.Printf("  Address: %p (Value type)\n", &stackBuf)
	fmt.Println("  (Copies the whole buffer if passed by value)")
	fmt.Println()

	fmt.Println("📌 2. Heap Allocation (Pointer):")
	fmt.Println("  buf := new(bytes.Buffer)")
	heapBuf := new(bytes.Buffer)
	heapBuf.WriteString("I am on the heap")
	fmt.Printf("  Address: %p (Pointer type)\n", heapBuf)
	fmt.Println("  (Efficient: Passes only the pointer/key around)")
	fmt.Println()
}

// ============================================================================
// PART 5: IO vs BUFIO (Performance & Features)
// ============================================================================

func Demo94_Part5_BufioVsIO() {
	fmt.Println("\n=== PART 5: IO vs BUFIO ===")
	fmt.Println()

	fmt.Println("📌 The Difference:")
	fmt.Println("  - io: Raw, direct access (Good for simple copies)")
	fmt.Println("  - bufio: Adds a \"Waiting Room\" (Buffer) to reduce system calls")
	fmt.Println()

	data := "Line 1\nLine 2\nLine 3"
	baseReader := strings.NewReader(data)

	fmt.Println("📌 bufio.Scanner (Line-by-Line):")
	// bufio wraps the base reader
	scanner := bufio.NewScanner(baseReader)

	lineCount := 0
	for scanner.Scan() {
		lineCount++
		// scanner.Text() gives us the line without the \n
		fmt.Printf("  Line %d: %s\n", lineCount, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	fmt.Println()
}

// ============================================================================
// PART 6: TYPE CONVERSION & INTERFACES (Under the Hood)
// ============================================================================

func Demo94_Part6_TypeConversion() {
	fmt.Println("\n=== PART 6: TYPE CONVERSION & INTERFACES ===")
	fmt.Println()

	fmt.Println("📌 Explicit Interface Conversion:")

	// Create a file (OS specific struct)
	fileName := "io_demo.txt"
	file, err := os.Create(fileName)
	if err != nil {
		log.Fatal(err)
	}

	// We ensure the file is closed using the io.Closer interface pattern
	// os.File implements Close(), so it satisfies io.Closer
	defer closeResource(file)

	// Explicitly converting *os.File to io.Writer
	// This works because *os.File implements the Write() method
	var writer io.Writer = file

	fmt.Printf("  Type of 'file':   %T\n", file)
	fmt.Printf("  Type of 'writer': %T (Interface holding *os.File)\n", writer)

	// We can now write to it using the interface
	writer.Write([]byte("Hello from Type Conversion!"))

	fmt.Println("  ✓ Wrote to file via io.Writer interface")
	fmt.Println()
}

// Helper function accepting the io.Closer interface
// This can close Files, Pipes, Network Connections, etc.
func closeResource(c io.Closer) {
	fmt.Println("  [Helper] Closing resource...")
	err := c.Close()
	if err != nil {
		fmt.Printf("  Error closing resource: %v\n", err)
	} else {
		fmt.Println("  [Helper] Resource closed successfully.")
	}
}

// ============================================================================
// MAIN DEMO FUNCTION
// ============================================================================

func Demo94IO() {
	fmt.Println("-- 94 Input/Output (io Package) --")
	fmt.Println()
	fmt.Println("Based on the transcript, this demo covers:")
	fmt.Println("1. Core Interfaces (Reader/Writer)")
	fmt.Println("2. Copying & MultiReading")
	fmt.Println("3. Pipes & Concurrency")
	fmt.Println("4. Buffering (bufio)")
	fmt.Println()

	Demo94_Part1_CoreInterfaces()
	Demo94_Part2_Utilities()
	Demo94_Part3_PipeAndConcurrency()
	Demo94_Part4_MemoryAllocation()
	Demo94_Part5_BufioVsIO()
	Demo94_Part6_TypeConversion()

	fmt.Println("\n=== END OF DEMO ===")
}
