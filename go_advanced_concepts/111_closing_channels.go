package main

import (
	"fmt"
	"time"
)

/*
TOPIC: CLOSING CHANNELS

CONCEPT:
Closing a channel is like turning off a valve. It signals to the receiver that
"no more data will ever be sent."

THE GOLDEN RULE:
Only the SENDER should close the channel.
- If a Receiver closes it, the Sender might panic.
- If you close it twice, the program panics.

READING RULES:
1. Open Channel: Blocks until data arrives.
2. Closed Channel (Data left): Returns the data.
3. Closed Channel (Empty): Returns "Zero Value" immediately (0, "", nil).
   It does NOT panic to read from a closed channel.
*/

// ---------------------------------------------------------
// Example 1: The "Comma Ok" Idiom
// ---------------------------------------------------------
// How to manually check if a channel is open or closed.
// Syntax: val, ok := <-ch
func example1_CommaOk() {
	fmt.Println("--- Example 1: The 'Comma Ok' Idiom ---")

	ch := make(chan int)

	// Sender
	go func() {
		for i := 1; i <= 3; i++ {
			fmt.Printf("[Sender] Sending %d\n", i)
			ch <- i
		}
		close(ch) // Signal completion
		fmt.Println("[Sender] Channel Closed.")
	}()

	// Receiver
	for {
		// 'val' receives the data
		// 'ok' is true if open, false if closed
		val, ok := <-ch

		if !ok {
			fmt.Println("[Receiver] Channel is closed! (ok == false)")
			break // Break the loop safely
		}
		fmt.Printf("[Receiver] Received: %d\n", val)
	}
	fmt.Println("----------------------------------------\n")
}

// ---------------------------------------------------------
// Example 2: The Range Shortcut (Preferred)
// ---------------------------------------------------------
// 'for range' automatically handles the "comma ok" check for you.
// It loops while the channel is open and breaks when it closes.
func example2_RangeLoop() {
	fmt.Println("--- Example 2: The Range Loop ---")

	ch := make(chan string)

	go func() {
		time.Sleep(100 * time.Millisecond)
		ch <- "Data A"
		ch <- "Data B"
		ch <- "Data C"

		// CRITICAL: If we forget to close here, the loop below
		// will wait forever (Deadlock).
		close(ch)
	}()

	// Loop continues until 'ch' is CLOSED and EMPTY
	for msg := range ch {
		fmt.Println("[Receiver] Range got:", msg)
	}
	fmt.Println("[Receiver] Loop finished automatically.")
	fmt.Println("----------------------------------------\n")
}

// ---------------------------------------------------------
// Example 3: The Pipeline Pattern (Producer -> Filter -> Main)
// ---------------------------------------------------------
// Demonstrates how closing propagates down a chain of workers.
// 1. Producer generates numbers -> Closes ch1
// 2. Filter reads ch1 -> Filters evens -> Sends to ch2 -> Closes ch2
// 3. Main reads ch2 -> Prints results

// Stage 1: Producer (Writes to 'out')
func producer(out chan<- int) {
	for i := 0; i < 10; i++ {
		out <- i
	}
	close(out) // Done producing
}

// Stage 2: Filter (Reads 'in', Writes to 'out')
func filter(in <-chan int, out chan<- int) {
	// Loop stops when 'in' is closed by Producer
	for val := range in {
		if val%2 == 0 {
			out <- val // Pass only even numbers
		}
	}
	close(out) // Done filtering
}

func example3_Pipeline() {
	fmt.Println("--- Example 3: The Pipeline Pattern ---")

	// Create the pipes
	ch1 := make(chan int)
	ch2 := make(chan int)

	// Start the background workers
	go producer(ch1)
	go filter(ch1, ch2)

	// Start the consumer (Main)
	// Range stops when 'filter' closes 'ch2'
	for val := range ch2 {
		fmt.Printf("[Main] Received Even Number: %d\n", val)
	}
	fmt.Println("----------------------------------------\n")
}

// ---------------------------------------------------------
// Example 4: The Panics (What NOT to do)
// ---------------------------------------------------------
// Uncomment the lines inside to see the crashes.
func example4_CommonMistakes() {
	fmt.Println("--- Example 4: Common Panics (Conceptual) ---")

	/*
		ch := make(chan int)
		close(ch)

		// MISTAKE A: Closing twice
		// panic: close of closed channel
		close(ch)

		// MISTAKE B: Sending to closed channel
		// panic: send on closed channel
		ch <- 1
	*/

	fmt.Println("1. Never close a channel twice.")
	fmt.Println("2. Never send to a closed channel.")
	fmt.Println("3. Always let the Sender close the channel.")
}

func main() {
	fmt.Println("═══════════════════════════════════════════════════════════")
	fmt.Println("TOPIC: CLOSING CHANNELS")
	fmt.Println("═══════════════════════════════════════════════════════════\n")

	example1_CommaOk()
	example2_RangeLoop()
	example3_Pipeline()
	example4_CommonMistakes()

	fmt.Println("═══════════════════════════════════════════════════════════")
	fmt.Println("KEY TAKEAWAYS")
	fmt.Println("═══════════════════════════════════════════════════════════")
	fmt.Println(`
1. SIGNALING: Closing is the standard way to say "I am done."
2. SAFETY: Reading from a closed channel is SAFE (returns zero-value).
   Sending to a closed channel is FATAL (panic).
3. PIPELINES: In a chain of Goroutines, each stage closes its output
   channel when it finishes processing.
	`)
}
