package main

import (
	"fmt"
	"sync"
	"time"
)

/*
TOPIC: RACE CONDITIONS

CONCEPT:
A Race Condition happens when multiple Goroutines try to access (read/write)
shared memory at the same time without synchronization.
The result becomes unpredictable (Non-deterministic).

THE DETECTOR:
Go has a built-in tool to catch this. Run your code with the -race flag.
Usage: go run -race main.go
*/

// ---------------------------------------------------------
// Example 1: The Buggy Counter (Race Condition)
// ---------------------------------------------------------
// This function fails to protect the 'counter' variable.
// If run with -race, it will panic/warn.
func example1_Buggy() {
	fmt.Println("--- Example 1: Buggy Counter (Run with -race) ---")

	var counter int
	var wg sync.WaitGroup

	// We launch 1000 goroutines. We expect 1000 as the result.
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()

			// CRITICAL SECTION (UNSAFE)
			// Read -> Modify -> Write happens here.
			// Two goroutines can read the same value (e.g., 5) at once,
			// both add 1, and both write 6. We lost an update.
			counter++
		}()
	}

	wg.Wait()
	fmt.Printf("Final Counter: %d (Expected 1000)\n", counter)

	if counter != 1000 {
		fmt.Println("Result: FAILURE (Data Race occurred)")
	} else {
		fmt.Println("Result: SUCCESS (But might be luck!)")
	}
	fmt.Println("----------------------------------------\n")
}

// ---------------------------------------------------------
// Example 2: The Fixed Counter (With Mutex)
// ---------------------------------------------------------
// This uses sync.Mutex to "lock the door" to the variable.
func example2_Fixed() {
	fmt.Println("--- Example 2: Fixed Counter (Mutex) ---")

	var counter int
	var wg sync.WaitGroup
	var mu sync.Mutex // The Lock

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()

			// LOCK
			// Only one goroutine can pass this line at a time.
			mu.Lock()

			// SAFE ZONE
			counter++

			// UNLOCK
			// Release the lock for the next goroutine.
			mu.Unlock()
		}()
	}

	wg.Wait()
	fmt.Printf("Final Counter: %d (Expected 1000)\n", counter)
	fmt.Println("Result: SUCCESS (Thread Safe)")
	fmt.Println("----------------------------------------\n")
}

func main() {
	fmt.Println("═══════════════════════════════════════════════════════════")
	fmt.Println("TOPIC: RACE CONDITIONS")
	fmt.Println("To see the detector work, run: go run -race main.go")
	fmt.Println("═══════════════════════════════════════════════════════════\n")

	// 1. Run the buggy version
	example1_Buggy()

	// Small pause to separate logs
	time.Sleep(500 * time.Millisecond)

	// 2. Run the fixed version
	example2_Fixed()

	fmt.Println("═══════════════════════════════════════════════════════════")
	fmt.Println("KEY TAKEAWAYS")
	fmt.Println("═══════════════════════════════════════════════════════════")
	fmt.Println(`
1. DATA RACES: Occur when concurrent tasks touch shared data without locks.
2. -race FLAG: Always use this during development/testing. It finds hidden bugs.
3. MUTEX: Use sync.Mutex to create a "Critical Section" where only one
   goroutine can enter at a time.
	`)
}
