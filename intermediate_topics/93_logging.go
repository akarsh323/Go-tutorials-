package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

// ============================================================================
// 93 Logging in Go
// ============================================================================
//
// Logging is how your program communicates what it's doing. Go provides:
// 1. Standard library log package - Simple and lightweight
// 2. File logging - Persist logs to disk
// 3. Third-party packages - Structured logging (Logrus, Zap, etc.)
//
// Key Progression:
// - Basic output (fmt) → Simple timestamps (log) → Files → Structured JSON
//
// ============================================================================

// ============================================================================
// PART 1: THE STANDARD log PACKAGE
// ============================================================================

func Demo93_Part1_StandardLog() {
	fmt.Println("\n=== PART 1: THE STANDARD log PACKAGE ===")
	fmt.Println()

	exampleBasicLogging()
	examplePrefixes()
	exampleFlags()
}

func exampleBasicLogging() {
	fmt.Println("📌 Basic Logging Functions:")
	fmt.Println("   log.Println() - Print with automatic timestamp")
	fmt.Println("   log.Printf()  - Formatted print with timestamp")
	fmt.Println("   log.Print()   - Print without newline")
	fmt.Println("   log.Fatal()   - Print and exit program (os.Exit(1))")
	fmt.Println("   log.Panic()   - Print and panic")
	fmt.Println()

	fmt.Println("   Example output:")
	log.Println("Application started")
	log.Printf("Processing user %d\n", 42)
	fmt.Println()
}

func examplePrefixes() {
	fmt.Println("📌 Adding Prefixes to Messages:")
	fmt.Println("   log.SetPrefix(\"PREFIX: \") adds text to start of every message")
	fmt.Println()

	// Save original prefix
	originalPrefix := log.Prefix()

	log.SetPrefix("INFO: ")
	log.Println("This has the INFO prefix")

	log.SetPrefix("ERROR: ")
	log.Println("This has the ERROR prefix")

	log.SetPrefix("WARN: ")
	log.Println("This has the WARN prefix")

	// Restore
	log.SetPrefix(originalPrefix)
	fmt.Println()
}

func exampleFlags() {
	fmt.Println("📌 Controlling Log Metadata (Flags):")
	fmt.Println("   Flags control what information appears in each log:")
	fmt.Println()

	originalFlags := log.Flags()

	// Show different flag combinations
	flagExamples := []struct {
		name  string
		flags int
	}{
		{"No flags", 0},
		{"Date only", log.Ldate},
		{"Time only", log.Ltime},
		{"Date + Time", log.Ldate | log.Ltime},
		{"Microseconds", log.Ldate | log.Ltime | log.Lmicroseconds},
		{"Short file (filename:line)", log.Ldate | log.Ltime | log.Lshortfile},
		{"Long file (full path)", log.Ldate | log.Ltime | log.Llongfile},
	}

	for _, example := range flagExamples {
		log.SetFlags(example.flags)
		fmt.Printf("   [%s]:\n", example.name)
		log.Println("   Sample message")
	}

	log.SetFlags(originalFlags)
	fmt.Println()
}

// ============================================================================
// PART 2: CUSTOM LOG LEVELS
// ============================================================================

func Demo93_Part2_CustomLogLevels() {
	fmt.Println("\n=== PART 2: CUSTOM LOG LEVELS ===")
	fmt.Println()

	fmt.Println("📌 Problem: Standard Library Lacks Log Levels")
	fmt.Println("   log package has no Info(), Warn(), Error() functions.")
	fmt.Println("   We need to create them manually using log.New()")
	fmt.Println()

	fmt.Println("📌 Solution: Create Custom Loggers")
	fmt.Println("   log.New(Writer, Prefix, Flags) returns a custom logger")
	fmt.Println()

	// Create custom loggers
	infoLogger := log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime)
	warnLogger := log.New(os.Stdout, "WARN: ", log.Ldate|log.Ltime)
	errorLogger := log.New(os.Stdout, "ERROR: ", log.Ldate|log.Ltime)

	fmt.Println("   Usage:")
	infoLogger.Println("Application initialized")
	warnLogger.Println("High memory usage detected")
	errorLogger.Println("Failed to connect to database")
	fmt.Println()

	fmt.Println("📌 Benefits of Custom Loggers:")
	fmt.Println("   ✓ Different prefixes for different severity levels")
	fmt.Println("   ✓ Can route to different outputs (stdout, file, etc.)")
	fmt.Println("   ✓ Consistent formatting across application")
	fmt.Println()
}

// ============================================================================
// PART 3: LOGGING TO FILES
// ============================================================================

func Demo93_Part3_LogToFile() {
	fmt.Println("\n=== PART 3: LOGGING TO FILES ===")
	fmt.Println()

	exampleOpeningFiles()
	exampleFileLogging()
}

func exampleOpeningFiles() {
	fmt.Println("📌 Opening Files with os.OpenFile:")
	fmt.Println("   os.OpenFile(filename, flags, permissions)")
	fmt.Println()

	fmt.Println("   Key Flags:")
	fmt.Println("   - os.O_CREATE:  Create file if it doesn't exist")
	fmt.Println("   - os.O_WRONLY:  Open for writing only")
	fmt.Println("   - os.O_APPEND:  Add to end (DON'T overwrite)")
	fmt.Println("   - os.O_TRUNC:   Truncate (empty) file on open")
	fmt.Println("   - os.O_RDONLY:  Open for reading")
	fmt.Println()

	fmt.Println("   Permissions (0666 = read/write for all):")
	fmt.Println("   - 0644 = user read/write, others read")
	fmt.Println("   - 0666 = everyone read/write")
	fmt.Println("   - 0755 = user read/write/execute, others read/execute")
	fmt.Println()

	fmt.Println("📌 For Logging, Always Use os.O_APPEND")
	fmt.Println("   ✓ CORRECT:   os.O_APPEND|os.O_CREATE|os.O_WRONLY")
	fmt.Println("   ✗ WRONG:     os.O_TRUNC (would erase previous logs!)")
	fmt.Println()
}

func exampleFileLogging() {
	fmt.Println("📌 Writing Logs to File:")

	// Create temp log file for demonstration
	logFile := "/tmp/demo_app.log"
	file, err := os.OpenFile(logFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// Create file logger
	fileLogger := log.New(file, "[APP] ", log.Ldate|log.Ltime|log.Lshortfile)

	fmt.Printf("   Writing to: %s\n", logFile)
	fileLogger.Println("Server started")
	fileLogger.Println("Listening on port 8080")
	fileLogger.Println("Request received from 192.168.1.1")

	fmt.Println("   ✓ Logs written to file")
	fmt.Println()

	fmt.Println("📌 Multi-Writer (Log to Both File and Stdout):")
	multiWriter := io.MultiWriter(os.Stdout, file)
	dualLogger := log.New(multiWriter, "[DUAL] ", log.Ltime)
	dualLogger.Println("This appears in both terminal and file")
	fmt.Println()
}

// ============================================================================
// PART 4: STRUCTURED LOGGING (THIRD-PARTY)
// ============================================================================

func Demo93_Part4_StructuredLogging() {
	fmt.Println("\n=== PART 4: STRUCTURED LOGGING (Third-Party) ===")
	fmt.Println()

	exampleStructuredConcept()
	exampleLogrusPattern()
	exampleZapPattern()
}

func exampleStructuredConcept() {
	fmt.Println("📌 What is Structured Logging?")
	fmt.Println()

	fmt.Println("   Text Log (Hard to parse):")
	fmt.Println("   2024-01-04 10:30:45 User john logged in from 192.168.1.1")
	fmt.Println()

	fmt.Println("   JSON Log (Machine-readable):")
	fmt.Println("   {\"timestamp\":\"2024-01-04T10:30:45Z\",\"message\":\"User logged in\",")
	fmt.Println("    \"username\":\"john\",\"ip\":\"192.168.1.1\",\"level\":\"info\"}")
	fmt.Println()

	fmt.Println("   Benefits of JSON:")
	fmt.Println("   ✓ Easy to parse with software")
	fmt.Println("   ✓ Consistent structure")
	fmt.Println("   ✓ Easily ingested into ELK Stack, Grafana, Datadog, etc.")
	fmt.Println("   ✓ Supports key-value pairs naturally")
	fmt.Println()
}

func exampleLogrusPattern() {
	fmt.Println("📌 Logrus (Popular Structured Logger):")
	fmt.Println()

	fmt.Println("   Installation: go get github.com/sirupsen/logrus")
	fmt.Println()

	fmt.Println("   Basic Usage Pattern:")
	fmt.Println("   ```go")
	fmt.Println("   import \"github.com/sirupsen/logrus\"")
	fmt.Println()
	fmt.Println("   logger := logrus.New()")
	fmt.Println("   logger.SetFormatter(&logrus.JSONFormatter{})")
	fmt.Println()
	fmt.Println("   // Structured fields as JSON")
	fmt.Println("   logger.WithFields(logrus.Fields{")
	fmt.Println("       \"username\": \"john\",")
	fmt.Println("       \"method\":   \"GET\",")
	fmt.Println("       \"status\":   200,")
	fmt.Println("   }).Info(\"User logged in\")")
	fmt.Println("   ```")
	fmt.Println()

	fmt.Println("   Output (formatted JSON):")
	fmt.Println("   {")
	fmt.Println("     \"level\": \"info\",")
	fmt.Println("     \"msg\": \"User logged in\",")
	fmt.Println("     \"username\": \"john\",")
	fmt.Println("     \"method\": \"GET\",")
	fmt.Println("     \"status\": 200,")
	fmt.Println("     \"time\": \"2024-01-04T10:30:45Z\"")
	fmt.Println("   }")
	fmt.Println()
}

func exampleZapPattern() {
	fmt.Println("📌 Zap (High-Performance Logger by Uber):")
	fmt.Println()

	fmt.Println("   Installation: go get go.uber.org/zap")
	fmt.Println()

	fmt.Println("   Why Zap?")
	fmt.Println("   ✓ Extremely fast (optimized for performance)")
	fmt.Println("   ✓ Strongly typed fields (type-safe)")
	fmt.Println("   ✓ Less memory allocation")
	fmt.Println()

	fmt.Println("   Basic Usage Pattern:")
	fmt.Println("   ```go")
	fmt.Println("   import \"go.uber.org/zap\"")
	fmt.Println()
	fmt.Println("   logger, _ := zap.NewProduction()")
	fmt.Println("   defer logger.Sync() // IMPORTANT: Flush before exit")
	fmt.Println()
	fmt.Println("   logger.Info(\"User logged in\",")
	fmt.Println("       zap.String(\"username\", \"john\"),")
	fmt.Println("       zap.String(\"method\", \"GET\"),")
	fmt.Println("       zap.Int(\"status\", 200),")
	fmt.Println("   )")
	fmt.Println("   ```")
	fmt.Println()

	fmt.Println("   Key Differences from Logrus:")
	fmt.Println("   ✓ Type-specific functions: zap.String(), zap.Int(), zap.Bool()")
	fmt.Println("   ✓ Requires defer logger.Sync() to flush buffered logs")
	fmt.Println("   ✓ Faster performance due to less allocations")
	fmt.Println()
}

// ============================================================================
// PART 5: IMPORTANT CONCEPTS
// ============================================================================

func Demo93_Part5_ImportantConcepts() {
	fmt.Println("\n=== PART 5: IMPORTANT CONCEPTS ===")
	fmt.Println()

	fmt.Println("📌 Concept 1: Log Levels (Severity)")
	fmt.Println("   DEBUG:  Detailed diagnostic info (development only)")
	fmt.Println("   INFO:   General informational messages")
	fmt.Println("   WARN:   Warning messages (something unexpected)")
	fmt.Println("   ERROR:  Error messages (something failed)")
	fmt.Println("   FATAL:  Critical error (program must exit)")
	fmt.Println()

	fmt.Println("📌 Concept 2: Log Rotation")
	fmt.Println("   Problem: Log files grow indefinitely → disk space exhausted")
	fmt.Println()
	fmt.Println("   Solution: Rotate logs periodically")
	fmt.Println("   - By size: Start new file when current reaches 100MB")
	fmt.Println("   - By time: Start new file daily/weekly")
	fmt.Println("   - Compress: Archive old logs (app.log.1.gz)")
	fmt.Println()
	fmt.Println("   Use: gopkg.in/natefinch/lumberjack.v2 for rotation")
	fmt.Println()

	fmt.Println("📌 Concept 3: Log Aggregation")
	fmt.Println("   In production, collect logs from multiple servers:")
	fmt.Println("   - ELK Stack: Elasticsearch + Logstash + Kibana")
	fmt.Println("   - Grafana Loki: Lightweight log aggregation")
	fmt.Println("   - Splunk: Enterprise log management")
	fmt.Println("   - Datadog: Hosted log management")
	fmt.Println()

	fmt.Println("📌 Concept 4: Context Logging")
	fmt.Println("   Add context to traces request flows:")
	fmt.Println("   - Request ID: Unique ID for each request")
	fmt.Println("   - User ID: Who made the request")
	fmt.Println("   - Trace ID: Track request through multiple services")
	fmt.Println()
}

// ============================================================================
// PART 6: BEST PRACTICES
// ============================================================================

func Demo93_Part6_BestPractices() {
	fmt.Println("\n=== PART 6: BEST PRACTICES ===")
	fmt.Println()

	fmt.Println("📌 Best Practice 1: Use Structured Logging in Production")
	fmt.Println("   ✓ Use Logrus or Zap with JSON format")
	fmt.Println("   ✗ Avoid plain text logs (hard to parse)")
	fmt.Println()

	fmt.Println("📌 Best Practice 2: Always Handle Errors")
	fmt.Println("   ✗ WRONG:")
	fmt.Println("      logger, _ := zap.NewProduction() // Ignores error")
	fmt.Println()
	fmt.Println("   ✓ RIGHT:")
	fmt.Println("      logger, err := zap.NewProduction()")
	fmt.Println("      if err != nil {")
	fmt.Println("          panic(err)")
	fmt.Println("      }")
	fmt.Println()

	fmt.Println("📌 Best Practice 3: Don't Log Sensitive Data")
	fmt.Println("   ✗ logger.Info(\"User password\", zap.String(\"pwd\", password))")
	fmt.Println("   ✓ logger.Info(\"User login successful\", zap.String(\"username\", user))")
	fmt.Println()

	fmt.Println("📌 Best Practice 4: Use Log Levels Appropriately")
	fmt.Println("   DEBUG → Development debugging (enabled only in dev)")
	fmt.Println("   INFO → Normal operations")
	fmt.Println("   WARN → Unexpected but recoverable")
	fmt.Println("   ERROR → Critical failures")
	fmt.Println()

	fmt.Println("📌 Best Practice 5: Centralize Logger Configuration")
	fmt.Println("   Create a single logger instance")
	fmt.Println("   Pass it to functions that need logging")
	fmt.Println("   Don't create new loggers in every function")
	fmt.Println()

	fmt.Println("📌 Best Practice 6: Use Defer for File Cleanup")
	fmt.Println("   file, _ := os.OpenFile(...)")
	fmt.Println("   defer file.Close()  // Guarantees cleanup")
	fmt.Println()
}

// ============================================================================
// PART 7: COMPLETE EXAMPLE
// ============================================================================

func Demo93_Part7_CompleteExample() {
	fmt.Println("\n=== PART 7: COMPLETE EXAMPLE ===")
	fmt.Println()

	fmt.Println("📌 Application Logger Setup Pattern:")
	fmt.Println()

	// Simulate production logger setup
	logFile := "/tmp/app_example.log"
	file, err := os.OpenFile(logFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal("Failed to open log file:", err)
	}
	defer file.Close()

	// Multi-destination logging
	multiWriter := io.MultiWriter(os.Stdout, file)

	// Create different loggers for different levels
	infoLogger := log.New(multiWriter, "[INFO] ", log.Ldate|log.Ltime|log.Lshortfile)
	warnLogger := log.New(multiWriter, "[WARN] ", log.Ldate|log.Ltime|log.Lshortfile)
	errorLogger := log.New(multiWriter, "[ERROR] ", log.Ldate|log.Ltime|log.Lshortfile)

	fmt.Println("   Logger Configuration:")
	fmt.Printf("   - Log file: %s\n", logFile)
	fmt.Println("   - Output: Both terminal and file")
	fmt.Println("   - Format: Timestamp + Log level + Message")
	fmt.Println()

	fmt.Println("   Simulating Application Lifecycle:")
	infoLogger.Println("Application starting...")
	infoLogger.Println("Loading configuration from config.yaml")
	infoLogger.Println("Connecting to database")
	warnLogger.Println("Database response slow (2500ms)")
	infoLogger.Println("Cache initialized")
	infoLogger.Println("Server listening on :8080")
	errorLogger.Println("Failed to send email notification")
	infoLogger.Println("Application shutdown")

	fmt.Println()
}

// ============================================================================
// MAIN DEMO FUNCTION
// ============================================================================

func Demo93Logging() {
	fmt.Println("-- 93 Logging in Go --")
	fmt.Println()
	fmt.Println("Logging is how your program communicates state and issues.")
	fmt.Println("Progression: fmt → log → files → structured JSON")
	fmt.Println()

	Demo93_Part1_StandardLog()
	Demo93_Part2_CustomLogLevels()
	Demo93_Part3_LogToFile()
	Demo93_Part4_StructuredLogging()
	Demo93_Part5_ImportantConcepts()
	Demo93_Part6_BestPractices()
	Demo93_Part7_CompleteExample()

	fmt.Println("\n=== SUMMARY ===")
	fmt.Println("✓ log package: Simple, built-in logging with timestamps")
	fmt.Println("✓ log.New(): Create custom loggers with different outputs")
	fmt.Println("✓ File logging: Use os.OpenFile with O_APPEND flag")
	fmt.Println("✓ Structured logging: Use Logrus or Zap for JSON output")
	fmt.Println("✓ Log rotation: Prevent log files from consuming all disk space")
	fmt.Println("✓ Best practice: Don't log sensitive data, use log levels")
}
