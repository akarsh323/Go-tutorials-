package main

import (
	"fmt"
	"time"
)

/*
TOPIC: CHANNELS & BLOCKING LOGIC

CONCEPT: Channels are the pipes that connect concurrent Goroutines.
They allow you to send values from one Goroutine to another.

TYPES:
1. UNBUFFERED (make(chan int)):
   - "Synchronous" / "Batton Handoff"
   - Capacity = 0.
   - The Sender BLOCKS until a Receiver is ready.
   - The Receiver BLOCKS until a Sender is ready.
   - Guaranteed synchronization.

2. BUFFERED (make(chan int, 3)):
   - "Asynchronous" / "Mailbox"
   - Capacity > 0.
   - Sender only BLOCKS if the buffer is FULL.
   - Receiver only BLOCKS if the buffer is EMPTY.

THE "BLOCKING" MECHANISM:
When a Goroutine blocks on a channel, it pauses execution. It does not burn CPU.
It goes to sleep until the Go Runtime wakes it up (when data arrives).
*/

// ---------------------------------------------------------
// Example 1: Unbuffered Channel (Scenario A - Successful Wait)
// ---------------------------------------------------------
// This demonstrates the "High-Five" or "Batton Pass".
// The Main function (Receiver) will wait for the Goroutine (Sender).
func scenarioA_SuccessfulWait() {
	fmt.Println("--- Scenario A: The Successful Wait (Unbuffered) ---")

	// 1. Create Unbuffered Channel
	ch := make(chan string)

	// 2. Start the Sender (Goroutine)
	go func() {
		fmt.Println("[Goroutine] I am going to sleep for 2 seconds...")
		time.Sleep(2 * time.Second) // Simulate work

		fmt.Println("[Goroutine] I woke up! Sending data now.")
		// BLOCKING POINT: If 'main' wasn't ready, this would wait here.
		ch <- "Hello from the other side!"
		fmt.Println("[Goroutine] Handshake complete. I am done.")
	}()

	fmt.Println("[Main] I am waiting at the channel (Blocking)...")

	// 3. BLOCKING POINT
	// The Main thread freezes here. It waits for the batton.
	msg := <-ch

	fmt.Println("[Main] Received message:", msg)
	fmt.Println("----------------------------------------------------\n")
}

// ---------------------------------------------------------
// Example 2: Buffered Channel (The Mailbox)
// ---------------------------------------------------------
// This demonstrates storage. The Sender does NOT block if there is space.
func bufferedChannelDemo() {
	fmt.Println("--- Example 2: Buffered Channel (The Mailbox) ---")

	// 1. Create Buffered Channel with capacity of 2
	mailbox := make(chan string, 2)

	fmt.Println("[Main] Sending letter 1...")
	mailbox <- "Letter A" // Won't block (Box has space)

	fmt.Println("[Main] Sending letter 2...")
	mailbox <- "Letter B" // Won't block (Box has space)

	fmt.Println("[Main] Both letters sent without a receiver present!")
	fmt.Println("[Main] Now reading them back...")

	fmt.Println("Received:", <-mailbox)
	fmt.Println("Received:", <-mailbox)
	fmt.Println("----------------------------------------------------\n")
}

// ---------------------------------------------------------
// Example 3: The Deadlock (Scenario B - The Ghost Wait)
// ---------------------------------------------------------
// WARNING: This function crashes the program if run.
// Logic: Receiver waits for data, but the Sender dies before sending.
func scenarioB_Deadlock() {
	fmt.Println("--- Scenario B: The Deadlock (Simulation) ---")

	ch := make(chan int)

	go func() {
		fmt.Println("[Goroutine] Running... but I won't send anything.")
		time.Sleep(1 * time.Second)
		fmt.Println("[Goroutine] Bye! (Dying without sending)")
		// The goroutine exits here.
	}()

	fmt.Println("[Main] Waiting for data...")

	// DEADLOCK HAPPENS HERE:
	// Go Runtime sees that 'ch' is empty.
	// It sees NO other goroutines are alive to send data.
	// It panics: "fatal error: all goroutines are asleep - deadlock!"
	val := <-ch

	fmt.Println("Received", val) // This line is never reached
}

// ---------------------------------------------------------
// Example 4: The Fix (Closing Channels)
// ---------------------------------------------------------
// How to prevent the "Ghost Wait"? The Sender must close the channel
// to say "I am done, no more data is coming."
func preventingDeadlock() {
	fmt.Println("--- Example 4: The Fix (Closing Channels) ---")

	ch := make(chan int)

	go func() {
		fmt.Println("[Goroutine] Sending 10...")
		ch <- 10
		fmt.Println("[Goroutine] I have no more data. Closing channel.")
		close(ch) // This tells the receiver: "Don't wait anymore."
	}()

	// We can loop over the channel. The loop breaks automatically when closed.
	for val := range ch {
		fmt.Printf("[Main] Received: %d\n", val)
	}
	fmt.Println("[Main] Channel closed. Loop finished. No Deadlock!")
	fmt.Println("----------------------------------------------------\n")
}

func main() {
	fmt.Println("═══════════════════════════════════════════════════════════")
	fmt.Println("TOPIC: CHANNELS & RUNTIME BLOCKING")
	fmt.Println("═══════════════════════════════════════════════════════════\n")

	// 1. Run Scenario A (Unbuffered / Synchronous)
	scenarioA_SuccessfulWait()

	// 2. Run Buffered Example (Asynchronous)
	bufferedChannelDemo()

	// 3. Run The Fix (How to avoid indefinite waiting)
	preventingDeadlock()

	// 4. SCENARIO B (DEADLOCK)
	// Uncomment the line below to see the crash!
	// scenarioB_Deadlock()

	fmt.Println("Note: Uncomment 'scenarioB_Deadlock()' in code to see the crash.")
}
