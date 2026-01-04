package main

import (
	"fmt"
	"sync"
	"time"
)

/*
TOPIC: SYNC.WAITGROUP

CONCEPT:
A WaitGroup is like a roll call sheet.
1. Add(n): "n people are going on the trip." (Increment counter)
2. Done(): "I have returned." (Decrement counter)
3. Wait(): "Bus driver waits until counter is 0." (Block)

CRITICAL RULES:
1. ALWAYS pass the WaitGroup by pointer (*sync.WaitGroup). If you pass by value,
   the worker gets a photocopy of the sheet, and the main function never sees the updates.
2. ALWAYS call Add() *before* starting the Goroutine to prevent race conditions.
*/

// ---------------------------------------------------------
// Example 1: The Basic WaitGroup (No Channels)
// ---------------------------------------------------------
func basicWorker(id int, wg *sync.WaitGroup) {
	// Defer ensures Done() is called even if the function crashes.
	defer wg.Done()

	fmt.Printf("[Worker %d] Starting...\n", id)
	time.Sleep(500 * time.Millisecond) // Simulate work
	fmt.Printf("[Worker %d] Finished.\n", id)
}

func example1_Basic() {
	fmt.Println("--- Example 1: Basic WaitGroup ---")

	var wg sync.WaitGroup
	numWorkers := 3

	// 1. ADD: Set the counter
	wg.Add(numWorkers)

	for i := 1; i <= numWorkers; i++ {
		// 2. Pass the memory address (&wg)
		go basicWorker(i, &wg)
	}

	// 3. WAIT: Block until counter is 0
	fmt.Println("[Main] Waiting for workers...")
	wg.Wait()

	fmt.Println("[Main] All workers finished.")
	fmt.Println("----------------------------------------\n")
}

// ---------------------------------------------------------
// Example 2: WaitGroup + Channels (The Monitor Pattern)
// ---------------------------------------------------------
// Problem: Range loops block forever if the channel isn't closed.
// Solution: Use a background Goroutine to Wait() then Close().
func channelWorker(id int, results chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()

	// Simulate calculation
	time.Sleep(500 * time.Millisecond)
	res := id * 2
	results <- res
}

func example2_WithChannels() {
	fmt.Println("--- Example 2: WaitGroup + Channels ---")

	var wg sync.WaitGroup
	numWorkers := 3
	results := make(chan int, numWorkers)

	wg.Add(numWorkers)

	for i := 1; i <= numWorkers; i++ {
		go channelWorker(i, results, &wg)
	}

	//
	// THE MONITOR GOROUTINE
	// This runs in the background. It monitors the WaitGroup.
	go func() {
		wg.Wait()      // Wait for all workers
		close(results) // Then close the channel safely
		fmt.Println("[Monitor] All workers done. Channel closed.")
	}()

	// Main thread consumes the results
	// This loop breaks automatically when 'results' is closed
	fmt.Println("[Main] Collecting results...")
	for res := range results {
		fmt.Printf("  -> Result: %d\n", res)
	}
	fmt.Println("----------------------------------------\n")
}

// ---------------------------------------------------------
// Example 3: Real World (Structs & Incremental Add)
// ---------------------------------------------------------
// Using a method on a struct and calling wg.Add(1) inside the loop.

type ConstructionWorker struct {
	ID   int
	Task string
}

func (w *ConstructionWorker) PerformTask(wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("[Worker %d] Started %s\n", w.ID, w.Task)
	time.Sleep(200 * time.Millisecond)
}

func example3_Structs() {
	fmt.Println("--- Example 3: Structs & Loops ---")

	var wg sync.WaitGroup
	tasks := []string{"Digging", "Laying Bricks", "Painting"}

	for i, task := range tasks {
		w := ConstructionWorker{
			ID:   i + 1,
			Task: task,
		}

		// INCREMENTAL ADD
		// It is safe to call Add(1) inside the loop, AS LONG AS
		// it happens before the 'go' keyword.
		wg.Add(1)

		go w.PerformTask(&wg)
	}

	wg.Wait()
	fmt.Println("[Main] Construction site clear.")
	fmt.Println("----------------------------------------\n")
}

// ---------------------------------------------------------
// Example 4: Conceptual Pitfalls (Explanation)
// ---------------------------------------------------------
func example4_Pitfalls() {
	fmt.Println("--- Example 4: Common Pitfalls (Read Only) ---")

	fmt.Println("1. THE NEGATIVE COUNTER:")
	fmt.Println("   If you call Done() more times than Add(), Go panics.")

	fmt.Println("2. THE DEADLOCK:")
	fmt.Println("   If you Add(5) but only 4 workers call Done(), Wait() blocks forever.")

	fmt.Println("3. THE POINTER MISTAKE:")
	fmt.Println("   func work(wg sync.WaitGroup) {} // WRONG! Copies the lock.")
	fmt.Println("   func work(wg *sync.WaitGroup) {} // CORRECT! Uses shared lock.")
}

func main() {
	fmt.Println("═══════════════════════════════════════════════════════════")
	fmt.Println("TOPIC: SYNC.WAITGROUP")
	fmt.Println("═══════════════════════════════════════════════════════════\n")

	example1_Basic()
	example2_WithChannels()
	example3_Structs()
	example4_Pitfalls()

	fmt.Println("═══════════════════════════════════════════════════════════")
	fmt.Println("KEY TAKEAWAYS")
	fmt.Println("═══════════════════════════════════════════════════════════")
	fmt.Println(`
1. POINTERS: Always pass *sync.WaitGroup (pointer) to workers.
2. ADD LOCATION: Call Add() in the parent thread (main), not inside the goroutine.
3. MONITOR PATTERN: When using channels, spawn a separate goroutine to Wait() 
   and Close() so the main thread can Range over results.
	`)
}
