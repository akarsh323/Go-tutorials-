package main

import (
	"fmt"
	"time"
)

/*
TOPIC: BUFFERED CHANNELS

CONCEPT: Unlike unbuffered channels (which have 0 capacity), Buffered Channels
have a specific "Storage Capacity".

SYNTAX: make(chan Type, capacity)
Example: make(chan string, 3) // Holds 3 strings

KEY DIFFERENCES:
1. ASYNCHRONOUS SEND: You can send values even if no one is listening,
   as long as the buffer isn't full.
2. BLOCKING RULES:
   - SENDING blocks only when the buffer is FULL.
   - RECEIVING blocks only when the buffer is EMPTY.

ANALOGY: The "Mailbox".
You can drop letters in the box and walk away (Asynchronous).
You only have to wait if the box is stuffed full (Blocking on Full).
*/

// ---------------------------------------------------------
// Example 1: Basic Storage (Non-Blocking Send)
// ---------------------------------------------------------
// Demonstrates that we can send data without an active receiver
// because the channel has storage space.
func example1_BasicStorage() {
	fmt.Println("--- Example 1: Basic Storage (No Blocking) ---")

	// Create a Buffered Channel with capacity of 2
	mailbox := make(chan string, 2)

	fmt.Println("[Main] Dropping Letter 1...")
	mailbox <- "Letter A" // Won't block (0/2 full)

	fmt.Println("[Main] Dropping Letter 2...")
	mailbox <- "Letter B" // Won't block (1/2 full)

	fmt.Println("[Main] Letters sent! No receiver was ready yet.")

	// Now we read them
	fmt.Println("[Main] Reading:", <-mailbox)
	fmt.Println("[Main] Reading:", <-mailbox)
	fmt.Println("-----------------------------------------------\n")
}

// ---------------------------------------------------------
// Example 2: Blocking When FULL
// ---------------------------------------------------------
// What happens if we try to put a 3rd item into a 2-item buffer?
// The Main thread will BLOCK until space is made.
func example2_BlockingOnFull() {
	fmt.Println("--- Example 2: Blocking When FULL ---")

	ch := make(chan int, 2) // Capacity = 2

	// Fill the buffer
	ch <- 1
	ch <- 2
	fmt.Println("[Main] Buffer is now FULL (2/2).")

	// Start a worker to free up space after 2 seconds
	go func() {
		fmt.Println("[Goroutine] Sleeping for 2 seconds...")
		time.Sleep(2 * time.Second)

		val := <-ch // RECEIVE: This removes 1 item, freeing a slot
		fmt.Printf("[Goroutine] I took value '%d' out. Space available!\n", val)
	}()

	fmt.Println("[Main] Trying to send value 3 (Blocking)...")

	// BLOCKING POINT:
	// This line freezes main until the Goroutine above runs '<-ch'
	ch <- 3

	fmt.Println("[Main] Success! Value 3 sent.")
	fmt.Println("-----------------------------------------------\n")
}

// ---------------------------------------------------------
// Example 3: Blocking When EMPTY
// ---------------------------------------------------------
// Just like unbuffered channels, if the bucket is empty,
// the Receiver must wait for data.
func example3_BlockingOnEmpty() {
	fmt.Println("--- Example 3: Blocking When EMPTY ---")

	ch := make(chan string, 2) // Empty initially

	// Start a slow sender
	go func() {
		fmt.Println("[Goroutine] Preparing data (1 sec)...")
		time.Sleep(1 * time.Second)
		ch <- "Data Packet"
		fmt.Println("[Goroutine] Data sent.")
	}()

	fmt.Println("[Main] Attempting to read (Buffer Empty)...")

	// BLOCKING POINT:
	// Waits for the Goroutine to send data
	msg := <-ch

	fmt.Println("[Main] Received:", msg)
	fmt.Println("-----------------------------------------------\n")
}

// ---------------------------------------------------------
// Example 4: The Deadlock (Overfilling)
// ---------------------------------------------------------
// WARNING: This crashes if run.
// If you overfill a buffer and no one ever reads from it,
// the program deadlocks.
func example4_DeadlockSimulation() {
	fmt.Println("--- Example 4: Deadlock Simulation ---")

	ch := make(chan int, 2)

	ch <- 1 // OK
	ch <- 2 // OK

	// DEADLOCK: Buffer is full. Main waits for space.
	// No other Goroutines are running to make space.
	// Crash!
	ch <- 3
}

func main() {
	fmt.Println("═══════════════════════════════════════════════════════════")
	fmt.Println("TOPIC: BUFFERED CHANNELS")
	fmt.Println("═══════════════════════════════════════════════════════════\n")

	// 1. Basic Storage
	example1_BasicStorage()

	// 2. Blocking on Full
	example2_BlockingOnFull()

	// 3. Blocking on Empty
	example3_BlockingOnEmpty()

	// 4. Deadlock (Uncomment to see crash)
	// example4_DeadlockSimulation()

	fmt.Println("Note: Uncomment 'example4_DeadlockSimulation()' to see the crash.")
}
