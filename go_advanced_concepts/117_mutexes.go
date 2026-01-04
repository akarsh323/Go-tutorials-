package main

import (
	"fmt"
	"sync"
	"time"
)

/*
TOPIC: SYNC.MUTEX (MUTUAL EXCLUSION)

CONCEPT:
A Mutex is like a "Key" to a shared resource.
- Lock(): "I am holding the key. No one else can enter."
- Unlock(): "I am returning the key."

THE PROBLEM (RACE CONDITION):
If multiple Goroutines try to write to the same variable at the same time,
they overwrite each other's work, leading to corrupted data.

THE RULE:
Code between Lock() and Unlock() is called the "Critical Section".
Only ONE Goroutine can be in the Critical Section at a time.
*/

// ---------------------------------------------------------
// Example 1: The Broken Counter (Race Condition)
// ---------------------------------------------------------
// This function demonstrates what happens WITHOUT a Mutex.
// The result will be inconsistent and wrong (e.g., 3452 instead of 5000).
func example1_BrokenCounter() {
	fmt.Println("--- Example 1: Broken Counter (Race Condition) ---")

	var count int
	var wg sync.WaitGroup

	// Launch 5 goroutines
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for j := 0; j < 1000; j++ {
				// RACE CONDITION:
				// Read (count) -> Modify (+1) -> Write (count)
				// Overlap here causes lost updates.
				count++
			}
		}()
	}

	wg.Wait()
	fmt.Printf("Final Value (Expected 5000): %d  <-- LIKELY WRONG!\n", count)
	fmt.Println("----------------------------------------\n")
}

// ---------------------------------------------------------
// Example 2: The Fixed Counter (With sync.Mutex)
// ---------------------------------------------------------
// This uses a Mutex to ensure only one Goroutine writes at a time.
func example2_FixedCounter() {
	fmt.Println("--- Example 2: Fixed Counter (With Mutex) ---")

	var count int
	var wg sync.WaitGroup
	var mu sync.Mutex // The "Key"

	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for j := 0; j < 1000; j++ {
				// 1. LOCK (Take Key)
				mu.Lock()

				// 2. CRITICAL SECTION (Modify Data)
				count++

				// 3. UNLOCK (Return Key)
				mu.Unlock()
			}
		}()
	}

	wg.Wait()
	fmt.Printf("Final Value (Expected 5000): %d  <-- CORRECT!\n", count)
	fmt.Println("----------------------------------------\n")
}

// ---------------------------------------------------------
// Example 3: Best Practice (Struct + Defer)
// ---------------------------------------------------------
// Bundle the data and the mutex together in a struct.
// Use 'defer' to ensure unlocking happens even if code panics.

// SafeCounter is thread-safe.
type SafeCounter struct {
	mu    sync.Mutex
	value int
}

// Increment adds 1 to the counter safely.
// CRITICAL: Must use Pointer Receiver (*SafeCounter) so we share the Mutex.
func (c *SafeCounter) Increment() {
	c.mu.Lock()
	defer c.mu.Unlock() // Ensures unlock happens at function exit

	c.value++
}

// GetValue reads the counter safely.
func (c *SafeCounter) GetValue() int {
	c.mu.Lock()
	defer c.mu.Unlock()

	// We lock reads too, to prevent reading half-written data.
	return c.value
}

func example3_StructBestPractice() {
	fmt.Println("--- Example 3: Best Practice (Struct + Defer) ---")

	counter := SafeCounter{}
	var wg sync.WaitGroup

	// Launch 10 goroutines
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for j := 0; j < 1000; j++ {
				counter.Increment()
			}
		}()
	}

	wg.Wait()
	fmt.Printf("Final Counter Value: %d\n", counter.GetValue())
	fmt.Println("----------------------------------------\n")
}

// ---------------------------------------------------------
// Example 4: Conceptual Deadlock (What NOT to do)
// ---------------------------------------------------------
func example4_DeadlockDemo() {
	fmt.Println("--- Example 4: Deadlock Explanation ---")
	fmt.Println("If you call Lock() twice without Unlocking, the program freezes.")

	/*
		var mu sync.Mutex
		mu.Lock()
		fmt.Println("Locked once.")

		mu.Lock() // DEADLOCK: Waiting for myself to unlock!
		fmt.Println("Locked twice (Never prints).")
	*/

	fmt.Println("(Skipped actual execution to prevent crash)")
}

func main() {
	fmt.Println("═══════════════════════════════════════════════════════════")
	fmt.Println("TOPIC: SYNC.MUTEX")
	fmt.Println("═══════════════════════════════════════════════════════════\n")

	example1_BrokenCounter()

	// Small sleep just to separate output clearly
	time.Sleep(500 * time.Millisecond)

	example2_FixedCounter()

	example3_StructBestPractice()

	example4_DeadlockDemo()

	fmt.Println("═══════════════════════════════════════════════════════════")
	fmt.Println("KEY TAKEAWAYS")
	fmt.Println("═══════════════════════════════════════════════════════════")
	fmt.Println(`
1. CRITICAL SECTION: The code between Lock() and Unlock(). Keep it small.
2. DEFER: Use 'defer mu.Unlock()' to prevent deadlocks if code returns early.
3. NO COPYING: Always pass structs with Mutexes by POINTER (*Struct).
   Copying a struct copies the mutex (creating a new, separate lock), which breaks protection.
	`)
}
