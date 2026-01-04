package main

import (
	"fmt"
	"time"
)

/*
Topic: GOROUTINES - A Comprehensive Guide

═══════════════════════════════════════════════════════════════════════════════

DEFINITION:
A Goroutine is a lightweight thread managed by the Go runtime. It is a function
that executes concurrently with other functions. Goroutines are incredibly
efficient, starting with a tiny stack (2KB) compared to OS threads (2MB).

KEY CONCEPTS:
- Lightweight: You can run hundreds of thousands of them simultaneously.
- Non-blocking: The 'go' keyword starts the function and moves to the next line.
- Managed by Go Runtime: Not managed directly by the Operating System.
- $M:N$ Scheduler: Multiplexes $M$ goroutines onto $N$ OS threads.



HOW IT WORKS (The G-M-P Model):
1. G (Goroutine): The code/task to be executed.
2. M (Machine): The actual OS thread.
3. P (Processor): A logical resource that connects G to M.
The Go runtime performs "work-stealing"—if one thread is blocked, it moves
waiting Goroutines to a different thread to keep the CPU busy.



CONCURRENCY VS PARALLELISM:
- Concurrency: Dealing with many things at once (Structure).
- Parallelism: Doing many things at once (Execution on multiple cores).

PITFALLS:
- Goroutine Leaks: Starting a goroutine that never exits, consuming memory.
- Race Conditions: Multiple goroutines accessing shared data simultaneously.
- Main Exit: If the main() function finishes, all other goroutines are killed.

═══════════════════════════════════════════════════════════════════════════════
*/

// sayHello is a simple function to demonstrate goroutine execution
func sayHello() {
	time.Sleep(500 * time.Millisecond) // Simulate a task taking time
	fmt.Println(">> Hello from Goroutine! (Worker finished)")
}

// printNumbers and printLetters demonstrate non-deterministic execution
func printNumbers() {
	for i := 1; i <= 3; i++ {
		fmt.Printf("[Number %d] ", i)
		time.Sleep(100 * time.Millisecond)
	}
}

func printLetters() {
	for _, l := range "ABC" {
		fmt.Printf("[Letter %c] ", l)
		time.Sleep(150 * time.Millisecond)
	}
}

func main() {
	fmt.Println("═══════════════════════════════════════════════════════════")
	fmt.Println("EXAMPLE 1: The 'Main Room' Problem")
	fmt.Println("═══════════════════════════════════════════════════════════\n")

	fmt.Println("Starting a Goroutine...")
	go sayHello() // This starts concurrently

	fmt.Println("The main function continues immediately.")
	fmt.Println("If we don't wait, the program ends before sayHello finishes.")

	// In a real app, we'd use sync.WaitGroup. For now, we sleep.
	time.Sleep(1 * time.Second)
	fmt.Println("\nMain function 'House' is closing.")

	// ─────────────────────────────────────────────────────────────────────────────

	fmt.Println("\n═══════════════════════════════════════════════════════════")
	fmt.Println("EXAMPLE 2: Concurrent Execution (Non-Deterministic)")
	fmt.Println("═══════════════════════════════════════════════════════════\n")

	fmt.Println("Running Numbers and Letters concurrently:")
	go printNumbers()
	go printLetters()

	// Wait long enough for both to finish
	time.Sleep(1 * time.Second)

	fmt.Println("\n\nNotice: The order of Numbers and Letters changes because")
	fmt.Println("the Go Scheduler decides who runs when based on thread availability.")

	// ─────────────────────────────────────────────────────────────────────────────

	fmt.Println("\n═══════════════════════════════════════════════════════════")
	fmt.Println("EXAMPLE 3: Error Propagation in Goroutines")
	fmt.Println("═══════════════════════════════════════════════════════════\n")

	var errResult error

	// We use an anonymous function closure to capture and 'return' an error
	go func() {
		fmt.Println("Goroutine performing work that might fail...")
		time.Sleep(200 * time.Millisecond)
		// Simulating an error
		errResult = fmt.Errorf("database connection failed")
	}()

	// Wait for goroutine to finish its work
	time.Sleep(500 * time.Millisecond)

	if errResult != nil {
		fmt.Println("Caught Error from Goroutine:", errResult)
	}

	// ─────────────────────────────────────────────────────────────────────────────

	fmt.Println("\n═══════════════════════════════════════════════════════════")
	fmt.Println("SUMMARY & BEST PRACTICES")
	fmt.Println("═══════════════════════════════════════════════════════════\n")

	fmt.Println(`
KEY TAKEAWAYS:

1. STACK SIZE:
   OS Threads: ~2MB fixed.
   Goroutines: ~2KB dynamic (can grow to GBs if needed).

2. SCHEDULING (M:N):
   Go manages its own scheduling in 'user space', making context 
   switches much faster than OS-level context switches.

3. CONCURRENCY IS NOT PARALLELISM:
   You can have concurrency on a single core (switching tasks).
   Parallelism requires multiple cores (running tasks simultaneously).

4. GOROUTINE LEAKS:
   A leaking 100,000 Goroutines = ~200MB of wasted RAM.
   Always ensure Goroutines have a clear exit condition.

BEST PRACTICES:

✓ Keep the 'Main' goroutine alive until background work is done.
✓ Use 'sync.WaitGroup' or 'Channels' for proper synchronization.
✓ Avoid shared global variables; prefer passing data through channels.
✓ Use 'context' package to signal cancellation/timeouts to goroutines.
✓ Implement error handling via closure variables or error channels.

Goroutines make high-performance concurrent programming simple and 
accessible in Go!
    `)
}
