package main

import (
	"fmt"
	"time"
)

/*
TOPIC: CHANNEL SYNCHRONIZATION

CONCEPT:
In real-world apps, the Main function (or parent Goroutine) must often wait
for background workers to finish before exiting or moving to the next step.
Without synchronization, the program might exit while workers are still running.

PATTERNS:
1. "DONE" SIGNAL: Blocking until a single worker reports completion.
2. WORKER POOL: launching N workers and waiting for N signals.
3. STREAMING: Using 'range' to read until a channel is closed.
*/

// ---------------------------------------------------------
// Pattern 1: Synchronizing a Single Goroutine
// ---------------------------------------------------------
// ANALOGY: A manager waiting for one employee to return to the desk
// before going home.
func syncSingleGoroutine() {
	fmt.Println("--- Pattern 1: Single Goroutine Sync ---")

	// 1. Create a channel for the signal.
	// We use struct{} (empty struct) because it uses 0 bytes of memory.
	// We don't care about the *value*, only that a send occurred.
	done := make(chan struct{})

	// 2. Start the background worker
	go func() {
		fmt.Println("[Worker] Started job (Heavy calculation)...")
		time.Sleep(1 * time.Second) // Simulate work

		fmt.Println("[Worker] Job finished. Sending 'done' signal.")

		// SIGNAL: Notify main that we are finished
		done <- struct{}{}
	}()

	fmt.Println("[Main] Waiting for worker...")

	// 3. BLOCKING POINT
	// The code pauses here explicitly. It cannot proceed until
	// the worker sends a value into 'done'.
	<-done

	fmt.Println("[Main] Signal received. Continuing.")
	fmt.Println("----------------------------------------\n")
}

// ---------------------------------------------------------
// Pattern 2: Synchronizing Multiple Goroutines (Worker Pool)
// ---------------------------------------------------------
// ANALOGY: A teacher waiting for 3 students to return to the bus.
// The bus doesn't leave until the count of students returned is 3.
func syncMultipleGoroutines() {
	fmt.Println("--- Pattern 2: Multiple Workers Sync ---")

	numWorkers := 3

	// Buffered channel avoids blocking the workers if Main isn't ready immediately.
	results := make(chan int, numWorkers)

	fmt.Printf("[Main] Spawning %d workers...\n", numWorkers)

	// 1. Launch multiple goroutines
	for i := 1; i <= numWorkers; i++ {
		go func(id int) {
			// Simulate variable work time
			fmt.Printf("\t[Worker %d] Starting...\n", id)
			time.Sleep(500 * time.Millisecond)

			fmt.Printf("\t[Worker %d] Done.\n", id)

			// Send completion signal (ID)
			results <- id
		}(i)
	}

	// 2. WAIT LOOP
	// Critical: We know we have 3 workers, so we must receive 3 times.
	for i := 0; i < numWorkers; i++ {
		// This line blocks until *any* worker sends a result
		id := <-results
		fmt.Printf("[Main] Acknowledged completion of Worker %d\n", id)
	}

	fmt.Println("[Main] All workers accounted for. Exiting.")
	fmt.Println("----------------------------------------\n")
}

// ---------------------------------------------------------
// Pattern 3: Data Exchange & Range Loops (Streaming)
// ---------------------------------------------------------
// ANALOGY: A News Ticker. The sender keeps pushing news.
// The receiver keeps reading. Reading stops only when the station goes off-air.
func syncDataStream() {
	fmt.Println("--- Pattern 3: Data Streaming (Range & Close) ---")

	// Unbuffered channel for streaming data
	stream := make(chan string)

	// 1. Start the Data Generator
	go func() {
		// CRITICAL: Ensure channel is closed when function finishes.
		// If we forget this, the Main loop will wait forever (Deadlock).
		defer close(stream)

		for i := 1; i <= 5; i++ {
			msg := fmt.Sprintf("Data Packet #%d", i)
			stream <- msg // Send data
			time.Sleep(200 * time.Millisecond)
		}
		fmt.Println("[Generator] All data sent. Closing channel.")
	}()

	fmt.Println("[Main] Listening for stream...")

	// 2. RANGE LOOP
	// Automatically receives values.
	// Automatically BREAKS when the channel is closed.
	for packet := range stream {
		fmt.Printf("\t[Receiver] Processed: %s\n", packet)
	}

	fmt.Println("[Main] Channel closed. Stream finished.")
	fmt.Println("----------------------------------------\n")
}

func main() {
	fmt.Println("═══════════════════════════════════════════════════════════")
	fmt.Println("TOPIC: CHANNEL SYNCHRONIZATION")
	fmt.Println("═══════════════════════════════════════════════════════════\n")

	// Run Pattern 1
	syncSingleGoroutine()

	// Run Pattern 2
	syncMultipleGoroutines()

	// Run Pattern 3
	syncDataStream()

	fmt.Println("═══════════════════════════════════════════════════════════")
	fmt.Println("KEY TAKEAWAYS")
	fmt.Println("═══════════════════════════════════════════════════════════")
	fmt.Println(`
1. ALWAYS WAIT: Main does not wait for Goroutines by default. You must block it.
2. SIGNALING: Use 'struct{}' for pure signals (0 memory usage).
3. MATCHING: If you launch N workers, ensure you receive N signals.
   - Too few receives = Memory Leak (Orphaned workers)
   - Too many receives = Deadlock (Main waits forever)
4. CLOSING: When streaming data with 'range', the SENDER must close the channel
   to let the receiver know the stream is finished.
	`)
}
