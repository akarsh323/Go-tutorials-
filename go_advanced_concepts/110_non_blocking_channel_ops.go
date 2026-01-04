package main

import (
	"fmt"
	"time"
)

/*
TOPIC: NON-BLOCKING CHANNEL OPERATIONS

CONCEPT:
By default, sending and receiving on channels is BLOCKING.
- Receiving (<-ch) waits until data is sent.
- Sending (ch <- val) waits until a receiver is ready (or buffer has space).

THE SOLUTION:
We use the 'select' statement with a 'default' case.
Logic: "Is the channel ready right now? If yes, proceed. If no, execute 'default' and move on."

USE CASES:
1. Polling: Checking for data without stopping other work.
2. Dropping Data: If a receiver is too slow, just drop the message instead of halting.
3. Event Loops: Handling multiple signals (Data, Quit, Idle) in one loop.
*/

// ---------------------------------------------------------
// Example 1: Non-Blocking Receive (The "Peek" Strategy)
// ---------------------------------------------------------
// Scenario: We look into the mailbox. If it's empty, we leave immediately.
// Without 'default', this code would deadlock.
func example1_NonBlockingReceive() {
	fmt.Println("--- Example 1: Non-Blocking Receive ---")

	ch := make(chan int) // Unbuffered, empty channel

	// We attempt to read from 'ch'
	select {
	case msg := <-ch:
		// This runs ONLY if data is ready to be read.
		fmt.Println("Received message:", msg)
	default:
		// This runs IMMEDIATELY if 'ch' is empty.
		fmt.Println("Channel is empty. I'm not waiting!")
	}
	fmt.Println("----------------------------------------\n")
}

// ---------------------------------------------------------
// Example 2: Non-Blocking Send (The "Drop and Go" Strategy)
// ---------------------------------------------------------
// Scenario: Trying to hand off a package. If no one takes it, we drop it.
// Without 'default', this code would block forever (deadlock).
func example2_NonBlockingSend() {
	fmt.Println("--- Example 2: Non-Blocking Send ---")

	ch := make(chan string) // Unbuffered channel (needs active receiver)

	msg := "Urgent Message"

	// We attempt to send to 'ch'
	select {
	case ch <- msg:
		// This runs ONLY if a receiver is waiting right now.
		fmt.Println("Message sent successfully.")
	default:
		// This runs if sending would block.
		fmt.Println("Receiver not ready. Dropping message:", msg)
	}
	fmt.Println("----------------------------------------\n")
}

// ---------------------------------------------------------
// Example 3: Real-World Scenario (The Smart Worker)
// ---------------------------------------------------------
// A worker that processes data, listens for a quit signal,
// and does "maintenance work" when idle, without freezing.
func example3_SmartWorker() {
	fmt.Println("--- Example 3: The 'Smart Consumer' Event Loop ---")

	dataChan := make(chan int)
	quitChan := make(chan bool)

	// Start the Worker
	go func() {
		for {
			select {
			// Case A: Process Data (High Priority)
			case d := <-dataChan:
				fmt.Printf("\t[Worker] ✅ Processed Job #%d\n", d)

			// Case B: Handle Quit Signal (High Priority)
			case <-quitChan:
				fmt.Println("\t[Worker] 🛑 Quit signal received. Shutting down.")
				return // Exit the loop and the Goroutine

			// Case C: Idle / Maintenance (Runs when A & B are not ready)
			default:
				fmt.Println("\t[Worker] ... Idle (Performing background tasks) ...")

				// CRITICAL: Prevent "Busy Waiting"
				// If we don't sleep here, the loop runs millions of times per second,
				// spiking CPU usage to 100%.
				time.Sleep(500 * time.Millisecond)
			}
		}
	}()

	// Producer Simulation
	fmt.Println("[Main] Sending 3 jobs over 3 seconds...")
	for i := 1; i <= 3; i++ {
		dataChan <- i
		time.Sleep(1 * time.Second) // Wait between sends
	}

	// Send shutdown signal
	fmt.Println("[Main] Work complete. Sending Quit signal.")
	quitChan <- true

	// Give the worker a moment to print the shutdown message
	time.Sleep(1 * time.Second)
	fmt.Println("----------------------------------------\n")
}

func main() {
	fmt.Println("═══════════════════════════════════════════════════════════")
	fmt.Println("TOPIC: NON-BLOCKING CHANNEL OPERATIONS")
	fmt.Println("═══════════════════════════════════════════════════════════\n")

	// 1. Run Peek Strategy
	example1_NonBlockingReceive()

	// 2. Run Drop Strategy
	example2_NonBlockingSend()

	// 3. Run Smart Worker
	example3_SmartWorker()

	fmt.Println("═══════════════════════════════════════════════════════════")
	fmt.Println("KEY TAKEAWAYS & BEST PRACTICES")
	fmt.Println("═══════════════════════════════════════════════════════════")
	fmt.Println(`
1. THE DEFAULT CASE: Adding 'default' inside a 'select' makes it non-blocking.
   It creates an "else" branch for when no channels are ready.

2. AVOID DEADLOCKS: This is the safest way to check a channel without risking
   your program hanging forever if the other side isn't ready.

3. BUSY WAITING: In an infinite loop (like Example 3), ALWAYS include a 
   'time.Sleep' inside the default case. 
   Without sleep, a 'default' loop will burn 100% of your CPU spinning uselessly.
	`)
}
