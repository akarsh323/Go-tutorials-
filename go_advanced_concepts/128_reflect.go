**Topic: The Select Statement**

This file provides a comprehensive guide to the `select` statement in Go.
It covers the "Race" between channels, the essential Timeout pattern, and Non-Blocking operations.

You can save this code in a file named `main.go` and run it using `go run main.go`.

```go
package main

import (
	"fmt"
	"time"
)

/*
TOPIC: THE SELECT STATEMENT

CONCEPT:
The 'select' statement is the Traffic Controller for Go concurrency.
It allows a single Goroutine to wait on MULTIPLE channels at once.
- If one channel is ready, it runs that case.
- If multiple are ready, it picks one at random.
- If none are ready, it BLOCKS (waits), unless a 'default' case exists.

USE CASES:
1. Racing multiple operations (First one wins).
2. Timeouts (Operation vs. Clock).
3. Non-Blocking checks (Peek and move on).
*/

// ---------------------------------------------------------
// Example 1: The Basic Race (Traffic Controller)
// ---------------------------------------------------------
// We have two workers. We wait for the FIRST one to finish.
func example1_BasicRace() {
	fmt.Println("--- Example 1: The Basic Race ---")

	fastChan := make(chan string)
	slowChan := make(chan string)

	// Worker 1 (Fast)
	go func() {
		time.Sleep(1 * time.Second)
		fastChan <- "Fast Worker Done"
	}()

	// Worker 2 (Slow)
	go func() {
		time.Sleep(3 * time.Second)
		slowChan <- "Slow Worker Done"
	}()

	fmt.Println("[Main] Waiting for the first response...")

	// SELECT BLOCK
	// This freezes until EITHER fastChan OR slowChan receives data.
	select {
	case msg1 := <-fastChan:
		fmt.Println("Winner:", msg1)
	case msg2 := <-slowChan:
		fmt.Println("Winner:", msg2)
	}
	
	fmt.Println("[Main] Race finished. We ignored the slower worker.")
	fmt.Println("----------------------------------------\n")
}

// ---------------------------------------------------------
// Example 2: The Timeout Pattern (Critical for Production)
// ---------------------------------------------------------
// Never let a Goroutine wait forever. We race the operation against a Timer.
func example2_Timeout() {
	fmt.Println("--- Example 2: The Timeout Pattern ---")

	c := make(chan string)

	// Simulate a slow database query (5 seconds)
	go func() {
		fmt.Println("[DB] Query started...")
		time.Sleep(5 * time.Second) 
		c <- "Query Result"
	}()

	fmt.Println("[Main] Waiting for DB (Timeout set to 2s)...")

	// The Race: Query vs Clock
	select {
	case res := <-c:
		fmt.Println("Success:", res)
	
	// time.After returns a channel that sends the time after the duration.
	// If 2 seconds pass, this channel receives a value, and this case wins.
	case <-time.After(2 * time.Second):
		fmt.Println("Error: Operation timed out!")
	}
	fmt.Println("----------------------------------------\n")
}

// ---------------------------------------------------------
// Example 3: Non-Blocking Select
// ---------------------------------------------------------
// Sometimes you want to check a channel but NOT wait if it's empty.
// We use 'default' for this.
func example3_NonBlocking() {
	fmt.Println("--- Example 3: Non-Blocking Select ---")

	messages := make(chan string)
	// Note: We never sent anything to 'messages', so it is empty.

	select {
	case msg := <-messages:
		fmt.Println("Received:", msg)
	default:
		// If 'messages' is not ready, Go jumps here IMMEDIATELY.
		fmt.Println("No messages waiting. Moving on to other work...")
	}
	fmt.Println("----------------------------------------\n")
}

func main() {
	fmt.Println("═══════════════════════════════════════════════════════════")
	fmt.Println("TOPIC: THE SELECT STATEMENT")
	fmt.Println("═══════════════════════════════════════════════════════════\n")

	example1_BasicRace()
	
	// Small sleep to separate output visually
	time.Sleep(500 * time.Millisecond)
	
	example2_Timeout()
	
	time.Sleep(500 * time.Millisecond)

	example3_NonBlocking()

	fmt.Println("═══════════════════════════════════════════════════════════")
	fmt.Println("KEY TAKEAWAYS")
	fmt.Println("═══════════════════════════════════════════════════════════")
	fmt.Println(`
1. FLOW CONTROL: 'select' is a switch statement for channels.
2. TIMEOUTS: Use 'case <-time.After(d):' to prevent deadlocks in production.
3. NON-BLOCKING: Use 'default:' to peek at a channel without stopping execution.
4. RANDOMNESS: If multiple channels are ready at once, 'select' chooses randomly.
	`)
}
