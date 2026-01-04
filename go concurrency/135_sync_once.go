package main

import (
	"fmt"
	"sync"
	"time"
)

/*
TOPIC: SYNC.ONCE (SINGLETON / LAZY INIT)

CONCEPT:
Sometimes you have a critical setup task (like connecting to a DB or loading a config)
that must happen EXACTLY ONCE.
If 100 Goroutines start at the same time and all try to connect, you don't want
100 connections. You want the first one to connect, and the other 99 to wait
and use that connection.

THE TOOL:
sync.Once is a struct with a single method: .Do(func)
- It tracks if the function has run.
- If not, it runs it (thread-safe).
- If yes, it skips it immediately.
*/

var (
	// The Guard
	once sync.Once

	// A WaitGroup to coordinate our workers
	wg sync.WaitGroup
)

// The Critical Task (Initialization)
// We want this to print ONLY ONE time.
func loadDatabase() {
	fmt.Println("   >>> [System] INITIALIZING DATABASE (Expensive Operation)...")
	time.Sleep(500 * time.Millisecond) // Simulate work
	fmt.Println("   >>> [System] DATABASE CONNECTED.")
}

// ---------------------------------------------------------
// Example 1: The Safe Initialization
// ---------------------------------------------------------
func main() {
	fmt.Println("═══════════════════════════════════════════════════════════")
	fmt.Println("TOPIC: SYNC.ONCE")
	fmt.Println("═══════════════════════════════════════════════════════════\n")

	// We spin up 5 workers simultaneously
	numWorkers := 5
	fmt.Printf("[Main] Spawning %d workers. They will all try to init the DB.\n\n", numWorkers)

	for i := 1; i <= numWorkers; i++ {
		wg.Add(1)

		go func(id int) {
			defer wg.Done()

			fmt.Printf("Goroutine %d: Requesting DB connection...\n", id)

			// THE MAGIC:
			// We wrap the function call in once.Do()
			// Only the VERY FIRST goroutine to reach this line executes 'loadDatabase'.
			// All others block until the first one finishes, then they proceed (skipping the init).
			once.Do(loadDatabase)

			fmt.Printf("Goroutine %d: Connection ready. Working...\n", id)
		}(i)
	}

	wg.Wait()
	fmt.Println("\n═══════════════════════════════════════════════════════════")
	fmt.Println("KEY TAKEAWAYS")
	fmt.Println("═══════════════════════════════════════════════════════════")
	fmt.Println(`
1. IDEMPOTENCY: 'sync.Once' guarantees the function runs exactly once.
   Even if called 1,000 times, the code executes only the first time.

2. THREAD SAFETY: It handles all the locking for you. You don't need manual Mutexes.

3. ONE-SHOT: You cannot "reset" a sync.Once. Once it fires, it is done forever.
   If you need to re-initialize later, you must create a new sync.Once instance.

4. USAGE: Perfect for Lazy Loading (connecting to DB only when the first user asks).
	`)
}
