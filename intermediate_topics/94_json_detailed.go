package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	Demo93Logging()
}

// ============================================================================
// 93 Logging in Go
// ============================================================================
//
// Logging is a critical aspect of software development. It allows developers
// and admins to monitor and debug applications.
//
// Progression:
// 1. Standard library (log) - Simple, robust, but basic.
// 2. Custom Loggers - Creating Info/Warn/Error levels manually.
// 3. File Logging - Persisting logs to disk using os.OpenFile.
// 4. Structured Logging - Using third-party tools (Logrus, Zap) for JSON logs.
//
// ============================================================================

// ============================================================================
// PART 1: THE STANDARD LOG PACKAGE
// ============================================================================

func Demo93_Part1_StandardLog() {
	fmt.Println("\n=== PART 1: THE STANDARD LOG PACKAGE ===")
	fmt.Println()

	exampleBasicPrinting()
	examplePrefixes()
	exampleFlags()
}

func exampleBasicPrinting() {
	fmt.Println("📌 Basic Logging:")
	fmt.Println("   The 'log' package provides robust support for logging.")
	fmt.Println("   It adds a timestamp to every message by default.")
	fmt.Println()

	fmt.Println("   Output:")
	log.Println("This is a simple log message")
	// log.Fatal("This would exit the program") // Commented out to keep demo running
	fmt.Println()
}

func examplePrefixes() {
	fmt.Println("📌 Log Prefixes:")
	fmt.Println("   We can add a prefix (like 'INFO: ') to identify log sources.")
	fmt.Println("   Use log.SetPrefix(\"STRING \")")
	fmt.Println()

	// Store original prefix to restore later
	originalPrefix := log.Prefix()

	log.SetPrefix("INFO: ")
	log.Println("This is an info message")

	log.SetPrefix("WARNING: ")
	log.Println("This is a warning message")

	// Restore original
	log.SetPrefix(originalPrefix)
	fmt.Println()
}

func exampleFlags() {
	fmt.Println("📌 Log Flags (Metadata):")
	fmt.Println("   Flags control the metadata (date, time, file info).")
	fmt.Println("   Use log.SetFlags() with constants like log.Ldate, log.Lshortfile.")
	fmt.Println()

	originalFlags := log.Flags()

	fmt.Println("   1. Date Only (log.Ldate):")
	log.SetFlags(log.Ldate)
	log.Println("Message with date only")

	fmt.Println("   2. Date + Time (log.Ldate | log.Ltime):")
	log.SetFlags(log.Ldate | log.Ltime)
	log.Println("Message with date and time")

	fmt.Println("   3. Short File (log.Lshortfile):")
	fmt.Println("   (Shows filename and line number - useful for debugging)")
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	log.Println("Message with filename")

	// Restore original
	log.SetFlags(originalFlags)
	fmt.Println()
}

// ============================================================================
// PART 2: CUSTOM LOG LEVELS
// ============================================================================

func Demo93_Part2_CustomLogLevels() {
	fmt.Println("\n=== PART 2: CUSTOM LOG LEVELS ===")
	fmt.Println()

	fmt.Println("📌 Problem:")
	fmt.Println("   The standard log package doesn't have built-in levels like Info/Warn/Error.")
	fmt.Println()

	fmt.Println("📌 Solution:")
	fmt.Println("   Create custom loggers using log.New().")
	fmt.Println("   Signature: log.New(writer, prefix, flags)")
	fmt.Println()

	// Creating custom loggers writing to Standard Output
	infoLogger := log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
	warnLogger := log.New(os.Stdout, "WARN: ", log.Ldate|log.Ltime|log.Lshortfile)
	errorLogger := log.New(os.Stdout, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)

	fmt.Println("   Output from Custom Loggers:")
	infoLogger.Println("This is an info message")
	warnLogger.Println("This is a warning message")
	errorLogger.Println("This is an error message")
	fmt.Println()
}

// ============================================================================
// PART 3: LOGGING TO FILES
// ============================================================================

func Demo93_Part3_LogToFile() {
	fmt.Println("\n=== PART 3: LOGGING TO FILES ===")
	fmt.Println()

	fmt.Println("📌 Opening Files:")
	fmt.Println("   To log to a file, we must open it using os.OpenFile.")
	fmt.Println("   We use specific flags to ensure we APPEND to the log, not overwrite it.")
	fmt.Println()

	fmt.Println("   Key Flags:")
	fmt.Println("   - os.O_APPEND: Add new logs to the end")
	fmt.Println("   - os.O_CREATE: Create file if missing")
	fmt.Println("   - os.O_WRONLY: Write only mode")
	fmt.Println("   - Permission: 0666 (Read/Write for everyone)")
	fmt.Println()

	// 1. Open/Create the file
	fileName := "app.log"
	file, err := os.OpenFile(fileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close() // Always defer close!

	fmt.Printf("   Successfully opened %s\n", fileName)

	// 2. Create a logger that writes to this file
	fileLogger := log.New(file, "FILE_LOG: ", log.Ldate|log.Ltime|log.Lshortfile)

	// 3. Write to file
	fileLogger.Println("This message goes to the file, not the console")
	fileLogger.Println("Application started successfully")

	fmt.Println("   ✓ Log messages written to app.log")

	// 4. Multi-Writer (Bonus from transcript)
	// often you want to log to BOTH file and console.
	fmt.Println("   ✓ Configuring MultiWriter (File + Console)...")
	multiWriter := io.MultiWriter(os.Stdout, file)
	dualLogger := log.New(multiWriter, "DUAL: ", log.Ldate|log.Ltime)
	dualLogger.Println("This appears in both places!")
	fmt.Println()
}

// ============================================================================
// PART 4: STRUCTURED LOGGING (THIRD-PARTY)
// ============================================================================

func Demo93_Part4_StructuredLogging() {
	fmt.Println("\n=== PART 4: STRUCTURED LOGGING (THIRD-PARTY) ===")
	fmt.Println()

	fmt.Println("📌 Why Third-Party Packages?")
	fmt.Println("   Standard logging produces text. Modern systems (ELK, Splunk) prefer JSON.")
	fmt.Println("   Packages like Logrus and Zap provide Structured Logging.")
	fmt.Println()

	exampleLogrusPattern()
	exampleZapPattern()
}

func exampleLogrusPattern() {
	fmt.Println("📌 1. Logrus (User-Friendly Structured Logging):")
	fmt.Println("   (Requires 'go get github.com/sirupsen/logrus')")
	fmt.Println()
	fmt.Println("   Key Features:")
	fmt.Println("   - JSON Formatter: logger.SetFormatter(&logrus.JSONFormatter{})")
	fmt.Println("   - Fields: Add context with log.WithFields(...)")
	fmt.Println()
	fmt.Println("   Code Example:")
	fmt.Println("   -------------------------------------------------")
	fmt.Println("   log.WithFields(logrus.Fields{")
	fmt.Println("       \"username\": \"John Doe\",")
	fmt.Println("       \"method\":   \"GET\",")
	fmt.Println("   }).Info(\"User logged in\")")
	fmt.Println("   -------------------------------------------------")
	fmt.Println("   Output -> {\"level\":\"info\", \"msg\":\"User logged in\", \"username\":\"John Doe\"...}")
	fmt.Println()
}

func exampleZapPattern() {
	fmt.Println("📌 2. Zap by Uber (High Performance):")
	fmt.Println("   (Requires 'go get go.uber.org/zap')")
	fmt.Println()
	fmt.Println("   Key Features:")
	fmt.Println("   - Zero-allocation (extremely fast)")
	fmt.Println("   - Strongly typed fields (zap.String, zap.Int)")
	fmt.Println("   - Must flush buffer: defer logger.Sync()")
	fmt.Println()
	fmt.Println("   Code Example:")
	fmt.Println("   -------------------------------------------------")
	fmt.Println("   logger, _ := zap.NewProduction()")
	fmt.Println("   defer logger.Sync()")
	fmt.Println("   logger.Info(\"User logged in\",")
	fmt.Println("       zap.String(\"username\", \"John Doe\"),")
	fmt.Println("       zap.Int(\"status\", 200),")
	fmt.Println("   )")
	fmt.Println("   -------------------------------------------------")
	fmt.Println()
}

// ============================================================================
// PART 5: IMPORTANT CONCEPTS
// ============================================================================

func Demo93_Part5_ImportantConcepts() {
	fmt.Println("\n=== PART 5: IMPORTANT CONCEPTS ===")
	fmt.Println()

	fmt.Println("📌 1. Structured Logging (JSON)")
	fmt.Println("   - Key-Value pairs make logs machine-readable.")
	fmt.Println("   - Easier to parse in tools like Kibana or Datadog.")
	fmt.Println("   - Better than Regex parsing plain text logs.")
	fmt.Println()

	fmt.Println("📌 2. Log Rotation")
	fmt.Println("   - Problem: Log files grow infinitely and fill the disk.")
	fmt.Println("   - Solution: Rotate logs (archive old ones, start fresh).")
	fmt.Println("   - Can be based on Size (e.g., every 10MB) or Time (Daily).")
	fmt.Println("   - Tools: 'lumberjack' is a popular Go package for this.")
	fmt.Println()

	fmt.Println("📌 3. Contextual Logging")
	fmt.Println("   - Attaching context (RequestID, UserID) to logs.")
	fmt.Println("   - Helps trace a specific request through the entire system.")
	fmt.Println()
}

// ============================================================================
// PART 6: BEST PRACTICES
// ============================================================================

func Demo93_Part6_BestPractices() {
	fmt.Println("\n=== PART 6: BEST PRACTICES ===")
	fmt.Println()

	fmt.Println("   ✓ 1. Use Log Levels:")
	fmt.Println("     Always filter logs (Info, Warn, Error). Don't log everything in Prod.")
	fmt.Println()

	fmt.Println("   ✓ 2. Handle Errors:")
	fmt.Println("     When opening log files, always check if err != nil.")
	fmt.Println("     Example: if err != nil { log.Fatal(err) }")
	fmt.Println()

	fmt.Println("   ✓ 3. Centralized Aggregation:")
	fmt.Println("     In production, send logs to a central system (ELK, Loki).")
	fmt.Println("     Don't just leave them on the server disk.")
	fmt.Println()

	fmt.Println("   ✓ 4. Don't Panic (usually):")
	fmt.Println("     Avoid log.Panic or log.Fatal in libraries. Let the main app decide when to exit.")
	fmt.Println()
}

// ============================================================================
// MAIN DEMO FUNCTION
// ============================================================================

func Demo93Logging() {
	fmt.Println("-- 93 Logging in Go --")
	fmt.Println()
	fmt.Println("Based on the transcript, this demo covers:")
	fmt.Println("1. Standard Library Basics")
	fmt.Println("2. File I/O for Logs")
	fmt.Println("3. Structured Logging Concepts")
	fmt.Println()

	Demo93_Part1_StandardLog()
	Demo93_Part2_CustomLogLevels()
	Demo93_Part3_LogToFile()
	Demo93_Part4_StructuredLogging()
	Demo93_Part5_ImportantConcepts()
	Demo93_Part6_BestPractices()

	fmt.Println("\n=== END OF DEMO ===")
	fmt.Println("Check the directory for 'app.log' to see the file output.")
}
