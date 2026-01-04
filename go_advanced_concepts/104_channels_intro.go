package main

import (
	"fmt"
	"time"
)

/*
Topic: GO CHANNELS - A Comprehensive Guide

═══════════════════════════════════════════════════════════════════════════════

DEFINITION:
Channels are the "pipes" that connect concurrent Goroutines. While Goroutines
allow independent execution (Concurrency), Channels allow those functions to
communicate (Data Exchange) and coordinate timing (Synchronization).

ANALOGY:
Think of a Channel as a Relay Race Baton.
- Runner A (Goroutine 1) has the baton.
- Runner B (Goroutine 2) is waiting.
- Runner B cannot start until Runner A hands over the baton.
- The handover point is the "Synchronization" event.

KEY CONCEPTS:
- Typed: A channel transports data of a specific type (e.g., chan string).
- Blocking: By default, sends and receives block execution until the other side is ready.
- Thread-Safe: You can safely access channels from multiple Goroutines without locks.

SYNTAX:
1. Creation:  ch := make(chan int)
2. Send:      ch <- 10        (Arrow points INTO the variable)
3. Receive:   val := <-ch     (Arrow points OUT OF the variable)

THE "DEADLOCK" RISK:
If a Goroutine tries to send data on a channel but no one is listening, it blocks forever.
If this happens in the main thread with no other active routines, Go crashes with:
"fatal error: all goroutines are asleep - deadlock!"

═══════════════════════════════════════════════════════════════════════════════
*/

// demonstrateBasicChannel shows the simplest send/receive pattern
func demonstrateBasicChannel() {
	// 1. Create the channel
	messages := make(chan string)

	// 2. Launch a Goroutine to SEND
	// We MUST do this in a separate Goroutine. If we did this in main,
	// main would block waiting for a receiver that hasn't been created yet.
	go func() {
		fmt.Println("   [Sender] Putting 'ping' into the pipe...")
		messages <- "ping" // This BLOCKS until 'main' is ready to receive
		fmt.Println("   [Sender] delivery complete!")
	}()

	// 3. Main Goroutine RECEIVES
	fmt.Println("   [Receiver] Waiting for message...")
	msg := <-messages // This BLOCKS until the Goroutine sends
	fmt.Println("   [Receiver] Received:", msg)
}

// demonstrateDeadlock explains why code crashes without Goroutines
func demonstrateDeadlock() {
	fmt.Println("   Code: ch <- \"hello\" (Blocking Main...)")
	fmt.Println("   Error: fatal error: all goroutines are asleep - deadlock!")
	fmt.Println("   Reason: Main paused to send, but no other routine exists to receive.")
}

// streamWeatherData mimics the API example from the transcript
func streamWeatherData() {
	// Channel for weather updates
	updates := make(chan string)

	// Producer: The Weather API (Background Goroutine)
	go func() {
		weatherData := []string{"Sunny 25C", "Cloudy 23C", "Rain 20C"}
		for _, data := range weatherData {
			fmt.Printf("   [API] Sending update: %s\n", data)
			updates <- data                    // Block until Main receives
			time.Sleep(200 * time.Millisecond) // Simulate delay between updates
		}
		close(updates) // Important: Tell receiver no more data is coming
	}()

	// Consumer: The Main App
	// We use 'range' to loop over the channel until it is closed
	for update := range updates {
		fmt.Printf("   [App] Displaying: %s\n", update)
	}
}

func main() {
	fmt.Println("═══════════════════════════════════════════════════════════")
	fmt.Println("EXAMPLE 1: Basic Synchronization (The 'Baton Pass')")
	fmt.Println("═══════════════════════════════════════════════════════════\n")

	demonstrateBasicChannel()

	// ─────────────────────────────────────────────────────────────────────────────

	fmt.Println("\n═══════════════════════════════════════════════════════════")
	fmt.Println("EXAMPLE 2: The Deadlock Mistake")
	fmt.Println("═══════════════════════════════════════════════════════════\n")

	// We are not running the actual code here because it would crash the program.
	// Instead, we describe the scenario explained in the transcript.
	demonstrateDeadlock()

	// ─────────────────────────────────────────────────────────────────────────────

	fmt.Println("\n═══════════════════════════════════════════════════════════")
	fmt.Println("EXAMPLE 3: Continuous Streams (Range over Channel)")
	fmt.Println("═══════════════════════════════════════════════════════════\n")

	fmt.Println("Simulating Weather API Stream...")
	streamWeatherData()

	// ─────────────────────────────────────────────────────────────────────────────

	fmt.Println("\n═══════════════════════════════════════════════════════════")
	fmt.Println("SUMMARY & BEHAVIOR TABLE")
	fmt.Println("═══════════════════════════════════════════════════════════\n")

	fmt.Println(`
KEY TAKEAWAYS:

1. CHANNELS ARE BLOCKING:
   Unbuffered channels synchronize execution.
   - Sender blocks until Receiver is ready.
   - Receiver blocks until Sender is ready.

2. THE "MAIN" GOROUTINE:
   Main is also a Goroutine. It must coordinate with background routines.
   If Main finishes (returns), all background routines die instantly.

3. SOLVING DEADLOCKS:
   Never send to an unbuffered channel in the SAME Goroutine that is 
   supposed to receive it later. Always spawn a 'go func()' for one side 
   of the transaction.

SUMMARY TABLE:

| Operation | Syntax         | Behavior (Unbuffered)                  |
|-----------|----------------|----------------------------------------|
| Send      | ch <- val      | Blocks until something receives.       |
| Receive   | val := <-ch    | Blocks until something sends.          |
| Deadlock  | N/A            | Occurs if waiting without a partner.   |

NEXT STEP:
Consider replacing 'time.Sleep' hacks with 'sync.WaitGroup' for robust
synchronization in complex apps.
    `)
}
