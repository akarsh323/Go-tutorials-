package intermediate

import (
	"fmt"
	"time"
)

// Topic 75: epoch
// Combines pragmatic examples with GoBootcamp patterns


func main() {

	fmt.Println("-- 75 Epoch --")

	// Current time in Unix seconds (seconds since Jan 1, 1970)
	now := time.Now()
	unixSec := now.Unix()
	fmt.Println("Unix timestamp (seconds):", unixSec)

	// Unix nanoseconds (seconds * 1e9 + nanosecond offset)
	unixNano := now.UnixNano()
	fmt.Println("Unix timestamp (nanoseconds):", unixNano)

	// UnixMilli: milliseconds since epoch (Go 1.17+)
	unixMilli := now.UnixMilli()
	fmt.Println("Unix timestamp (milliseconds):", unixMilli)

	// UnixMicro: microseconds since epoch (Go 1.17+)
	unixMicro := now.UnixMicro()
	fmt.Println("Unix timestamp (microseconds):", unixMicro)

	// Unix epoch zero point
	epochZero := time.Unix(0, 0).UTC()
	fmt.Println("Unix epoch zero:", epochZero)

	// Converting Unix timestamp back to time
	timestamp := int64(1609459200) // 2021-01-01 00:00:00 UTC
	t := time.Unix(timestamp, 0).UTC()
	fmt.Println("Timestamp", timestamp, "converts to:", t)

	// Unix timestamp with nanosecond precision
	secs := int64(1609459200)
	nanos := int64(500000000) // 0.5 seconds
	preciseTime := time.Unix(secs, nanos).UTC()
	fmt.Println("Time with nanoseconds:", preciseTime)

	// Common use case: calculating duration from timestamp
	pastTime := time.Unix(unixSec-3600, 0) // 1 hour ago
	duration := now.Sub(pastTime)
	fmt.Println("Time elapsed:", duration)

	// Comparison of different units
	fmt.Printf("Same time in different units:\n")
	fmt.Printf("  Seconds: %d\n", now.Unix())
	fmt.Printf("  Milliseconds: %d\n", now.UnixMilli())
	fmt.Printf("  Microseconds: %d\n", now.UnixMicro())
	fmt.Printf("  Nanoseconds: %d\n", now.UnixNano())
}

func epochExample1() {
	
	fmt.Println("📚 Epoch = seconds since Jan 1, 1970 UTC")

	// Get current time and its epoch
	now := time.Now()
	epochSeconds := now.Unix()
	epochNanos := now.UnixNano()

	fmt.Printf("Current time: %v\n", now.Format("2006-01-02 15:04:05"))
	fmt.Printf("Unix epoch (seconds): %d\n", epochSeconds)
	fmt.Printf("Unix epoch (nanoseconds): %d\n", epochNanos)

	// The epoch reference
	epoch := time.Date(1970, 1, 1, 0, 0, 0, 0, time.UTC)
	fmt.Printf("\nEpoch reference: %v\n", epoch)

	// Seconds since epoch
	fmt.Printf("Seconds since epoch: %d\n", epochSeconds)
}

func epochExample10() {
	
	fmt.Println(`
Epoch Key Points:

What is Epoch:
  • Jan 1, 1970, 00:00:00 UTC = time 0
  • Used to represent any moment as a number
  • Same epoch = same moment everywhere

Units:
  time.Unix()      → seconds since epoch
  time.UnixMilli() → milliseconds (JavaScript)
  time.UnixMicro() → microseconds (high precision)
  time.UnixNano()  → nanoseconds (benchmarking)

Conversion:
  To epoch:        t.Unix(), t.UnixMilli(), etc.
  From epoch:      time.Unix(seconds, nanoseconds)

Why Use Epoch:
  ✓ Simple to store and transmit
  ✓ Easy to calculate time differences
  ✓ Works across timezones
  ✓ Standard in APIs and databases

When to Use:
  ✓ Storing in databases (efficient)
  ✓ APIs and JSON responses
  ✓ Log files and timestamps
  ✓ Calculating durations

Display:
  ✓ Always store as epoch in databases
  ✓ Convert to readable time for display
  ✓ Don't mix storage and display formats

Important Facts:
  • Epoch is ALWAYS UTC
  • Different units (seconds/millis) for different needs
  • Leap seconds are usually ignored in practice
	`)
}
