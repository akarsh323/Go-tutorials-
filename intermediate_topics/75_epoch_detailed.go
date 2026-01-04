package intermediate

import (
	"fmt"
	"time"
)

// Topic 75: Unix Epoch - Fundamental Time Concept
// ================================================
// This lesson breaks down Unix Epoch like a computer science class.
// It covers: what the epoch is, why we use it, practical Go implementation,
// and key concepts for working with time in production systems.

func main() {
	fmt.Println("=== Topic 75: Unix Epoch - The Foundation of Computer Time ===\n")

	lesson1WhatIsEpoch()
	fmt.Println("\n" + string([]byte{61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 61}) + "\n")

	lesson2PracticalCodingInGo()
	fmt.Println("\n" + string([]byte{61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 61}) + "\n")

	lesson3AllUnitsComparison()
	fmt.Println("\n" + string([]byte{61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 61}) + "\n")

	lesson4KeyConcepts()
	fmt.Println("\n" + string([]byte{61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 61}) + "\n")

	lesson5TeachersInsiight()
}

// LESSON 1: What is the Epoch?
// =============================
// Imagine a giant stopwatch that started ticking at a specific moment in history.
func lesson1WhatIsEpoch() {
	fmt.Println("LESSON 1: WHAT IS THE EPOCH?")
	fmt.Println("------------------------------\n")

	fmt.Println("Imagine a giant stopwatch that started ticking at a very specific moment in history.\n")

	fmt.Println("THE STARTING LINE:")
	fmt.Println("  • 00:00:00 UTC on January 1, 1970")
	fmt.Println("  • This is the reference point for all Unix timestamps\n")

	fmt.Println("THE COUNT:")
	fmt.Println("  • We simply count the SECONDS that have ticked by since that moment")
	fmt.Println("  • This allows computers to store time as a simple, single integer\n")

	fmt.Println("THE NUMBER SYSTEM:")
	fmt.Println("  • Positive Numbers = Dates AFTER Jan 1, 1970")
	fmt.Println("  • Negative Numbers = Dates BEFORE Jan 1, 1970 (rare in practice)\n")

	fmt.Println("IMPORTANT NOTE:")
	fmt.Println("  • Unix time does NOT count \"leap seconds\" (Earth's rotation adjustments)")
	fmt.Println("  • This keeps the math simple and consistent across all systems\n")

	// Practical demonstration
	now := time.Now()
	epochZero := time.Unix(0, 0).UTC()

	fmt.Println("DEMONSTRATION:")
	fmt.Printf("  Epoch Zero:        %v\n", epochZero)
	fmt.Printf("  Current Time:      %v\n", now)
	fmt.Printf("  Current Timestamp: %d seconds since 1970\n", now.Unix())
}

// LESSON 2: Practical Coding in Go
// =================================
func lesson2PracticalCodingInGo() {
	fmt.Println("LESSON 2: PRACTICAL CODING IN GO")
	fmt.Println("---------------------------------\n")

	fmt.Println("The basic workflow:")
	fmt.Println("  1. Get the current time")
	fmt.Println("  2. Turn it into a Unix timestamp (a big number)")
	fmt.Println("  3. Turn that number back into a readable time")
	fmt.Println("  4. Format it nicely\n")

	// STEP 1: Getting the Unix Timestamp
	fmt.Println("STEP 1: Getting the Unix Timestamp")
	fmt.Println("-----------------------------------")
	now := time.Now()
	unixTime := now.Unix()

	fmt.Printf("Current Time Object:      %v\n", now)
	fmt.Printf("Unix Timestamp (Seconds): %d\n", unixTime)
	fmt.Println("\nWhat is happening here?")
	fmt.Println("  • time.Now() gives us a complex object with date, time, and timezone")
	fmt.Println("  • .Unix() strips all complexity and gives us just the seconds")
	fmt.Println("  • This big number is perfect for storing in databases\n")

	// STEP 2: Converting Unix Back to Human Time
	fmt.Println("STEP 2: Converting Unix Back to Human Time")
	fmt.Println("-------------------------------------------")

	// Demonstrate with a known timestamp (2021-01-01 00:00:00 UTC)
	demonstrationTimestamp := int64(1609459200)
	t := time.Unix(demonstrationTimestamp, 0).UTC()

	fmt.Printf("Timestamp (from database): %d\n", demonstrationTimestamp)
	fmt.Printf("Converted to Time object:  %v\n", t)
	fmt.Println("\nWhy the second argument (0)?")
	fmt.Println("  • time.Unix(seconds, nanoseconds) allows extreme precision")
	fmt.Println("  • If we only have seconds, we pass 0 for nanoseconds")
	fmt.Println("  • This is the standard way to handle second-level timestamps\n")

	// STEP 3: Formatting the Date
	fmt.Println("STEP 3: Formatting the Date (Human Readable)")
	fmt.Println("---------------------------------------------")

	formattedTime := t.Format("2006-01-02")
	detailedFormat := t.Format("Monday, January 2, 2006 at 15:04:05 MST")

	fmt.Printf("Simple format (YYYY-MM-DD):     %s\n", formattedTime)
	fmt.Printf("Detailed format:                %s\n", detailedFormat)
	fmt.Println("\nGo's unique formatting system:")
	fmt.Println("  • Reference layout: Mon Jan 2 15:04:05 MST 2006")
	fmt.Println("  • This specific date/time is used to define the format")
	fmt.Println("  • Example: \"2006-01-02\" means Year-Month-Day")
	fmt.Println("  • Example: \"15:04:05\" means Hour:Minute:Second (24-hour)")
}

// LESSON 3: All Units Comparison
// ===============================
func lesson3AllUnitsComparison() {
	fmt.Println("LESSON 3: UNIX TIMESTAMP UNITS")
	fmt.Println("-------------------------------\n")

	fmt.Println("Different situations require different precision levels:\n")

	now := time.Now()

	fmt.Printf("Current moment: %v\n\n", now)

	fmt.Printf("Unix Seconds:        %d\n", now.Unix())
	fmt.Println("  → Standard for databases and APIs")
	fmt.Println("  → Sufficient for most use cases\n")

	fmt.Printf("Unix Milliseconds:   %d\n", now.UnixMilli())
	fmt.Println("  → Used by JavaScript (it counts in milliseconds)")
	fmt.Println("  → Useful for high-precision timestamps\n")

	fmt.Printf("Unix Microseconds:   %d\n", now.UnixMicro())
	fmt.Println("  → Even higher precision (1 millionth of a second)")
	fmt.Println("  → Used in scientific and timing applications\n")

	fmt.Printf("Unix Nanoseconds:    %d\n", now.UnixNano())
	fmt.Println("  → Maximum precision (1 billionth of a second)")
	fmt.Println("  → Rarely used but available for benchmarking\n")

	fmt.Println("Unit Relationships:")
	fmt.Println("  • 1 second       = 1,000 milliseconds")
	fmt.Println("  • 1 millisecond  = 1,000 microseconds")
	fmt.Println("  • 1 microsecond  = 1,000 nanoseconds")
}

// LESSON 4: Key Concepts Summary
// ===============================
func lesson4KeyConcepts() {
	fmt.Println("LESSON 4: KEY CONCEPTS SUMMARY")
	fmt.Println("-------------------------------\n")

	fmt.Println("CONCEPT TABLE:")
	fmt.Println("+----------------------+---------------------------------------------+")
	fmt.Println("| Concept              | Explanation                                 |")
	fmt.Println("+----------------------+---------------------------------------------+")
	fmt.Println("| Unix Epoch           | The \"zero\" start time: Jan 1, 1970, UTC   |")
	fmt.Println("| Standard Unit        | Usually SECONDS (Unix Timestamp)             |")
	fmt.Println("| Alternative Units    | Milliseconds (JavaScript), Microseconds      |")
	fmt.Println("| Universal            | Works the same on Mac, Windows, Linux, etc.  |")
	fmt.Println("| Leap Seconds         | Ignored by Unix time for simplicity         |")
	fmt.Println("| Storage              | Simple to store and transmit                |")
	fmt.Println("| Time Differences     | Easy math for calculating durations         |")
	fmt.Println("| Timezones            | Epoch is ALWAYS UTC (no timezone issues)    |")
	fmt.Println("+----------------------+---------------------------------------------+\n")

	fmt.Println("WHY USE EPOCH:")
	fmt.Println("  ✓ Simple to store and transmit")
	fmt.Println("  ✓ Easy to calculate time differences")
	fmt.Println("  ✓ Works across timezones")
	fmt.Println("  ✓ Standard in APIs and databases")
	fmt.Println("  ✓ No human interpretation needed by computers\n")

	fmt.Println("WHEN TO USE EPOCH:")
	fmt.Println("  ✓ Storing in databases (most efficient)")
	fmt.Println("  ✓ APIs and JSON responses")
	fmt.Println("  ✓ Log files and server timestamps")
	fmt.Println("  ✓ Calculating durations between two times")
	fmt.Println("  ✓ Any time-based comparisons or sorting\n")

	fmt.Println("BEST PRACTICE:")
	fmt.Println("  ✓ Always STORE as epoch (numbers) in databases")
	fmt.Println("  ✓ Always CONVERT to readable time for user display")
	fmt.Println("  ✓ Never mix storage and display formats")
}

// LESSON 5: Teacher's Insight - The Core Philosophy
// ==================================================
func lesson5TeachersInsiight() {
	fmt.Println("LESSON 5: TEACHER'S INSIGHT - THE CORE PHILOSOPHY")
	fmt.Println("--------------------------------------------------\n")

	fmt.Println("THE FUNDAMENTAL TRUTH:")
	fmt.Println("  Computers prefer NUMBERS, but humans prefer DATES.\n")

	fmt.Println("THE WORKFLOW:")
	fmt.Println("  WHEN SAVING/CALCULATING:")
	fmt.Println("    1. Convert time to UNIX (a number)")
	fmt.Println("    2. Store/calculate/transmit")
	fmt.Println("    3. Result is always a number\n")

	fmt.Println("  WHEN DISPLAYING:")
	fmt.Println("    1. Get the number from database/API")
	fmt.Println("    2. Convert to Time object with time.Unix()")
	fmt.Println("    3. Format with .Format() for human readers\n")

	fmt.Println("REAL-WORLD EXAMPLE:")
	now := time.Now()
	oneHourAgo := time.Unix(now.Unix()-3600, 0)
	duration := now.Sub(oneHourAgo)

	fmt.Printf("  Current time:      %v (Unix: %d)\n", now, now.Unix())
	fmt.Printf("  1 hour ago:        %v (Unix: %d)\n", oneHourAgo, oneHourAgo.Unix())
	fmt.Printf("  Duration between:  %v\n\n", duration)

	fmt.Println("Why this matters:")
	fmt.Println("  • Database stores both as numbers: 1735950000 and 1735946400")
	fmt.Println("  • Subtraction is simple: 1735950000 - 1735946400 = 3600 seconds")
	fmt.Println("  • Convert to readable format ONLY when showing to user")
	fmt.Println("  • This pattern scales to millions of records efficiently")
}



timrexample2(0{

}