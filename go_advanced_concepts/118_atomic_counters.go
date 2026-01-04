package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

/*
TOPIC: ATOMIC COUNTERS (SYNC/ATOMIC)

CONCEPT:
Atomic operations use low-level hardware instructions to modify memory safely
without the overhead of Mutex locks.

ANALOGY:
- NO SYNC: Two students fight over a marker (Data Corruption).
- MUTEX: One student holds the key; others wait outside (Safe but slower).
- ATOMIC: A magic marker that updates instantly in one un-interruptible motion.

WHEN TO USE:
- Use Atomics for simple counters, flags, or gauges.
- Use Mutexes for complex logic (Maps, Slices, multiple variable updates).
*/

// ---------------------------------------------------------
// Example 1: The "Lost Update" Problem (No Protection)
// ---------------------------------------------------------
// Without atomics or locks, goroutines overwrite each other.
func example1_RaceCondition() {
	fmt.Println("--- Example 1: Race Condition (No Protection) ---")

	var count int64 // Standard int64
	var wg sync.WaitGroup

	// 50 Goroutines * 1000 increments = Expected 50,000
	for i := 0; i < 50; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for j := 0; j < 1000; j++ {
				count++ // DANGER: Not thread-safe
			}
		}()
	}

	wg.Wait()
	fmt.Printf("Final Value: %d (Expected 50000) <-- LIKELY WRONG\n", count)
	fmt.Println("----------------------------------------\n")
}

// ---------------------------------------------------------
// Example 2: The Atomic Solution
// ---------------------------------------------------------
// Using sync/atomic to perform thread-safe math.
type AtomicCounter struct {
	// We use int64 because atomic functions are optimized for fixed-size types
	value int64
}

// Increment safely adds 1.
// CRITICAL: We pass a POINTER to the value (&c.value).
func (c *AtomicCounter) Increment() {
	atomic.AddInt64(&c.value, 1)
}

// GetValue safely reads the current value.
func (c *AtomicCounter) GetValue() int64 {
	return atomic.LoadInt64(&c.value)
}

func example2_AtomicCounter() {
	fmt.Println("--- Example 2: Atomic Counter (Safe) ---")

	var wg sync.WaitGroup
	counter := AtomicCounter{}

	numGoroutines := 50

	fmt.Printf("Starting %d goroutines (1000 increments each)...\n", numGoroutines)
	wg.Add(numGoroutines)

	for i := 0; i < numGoroutines; i++ {
		go func() {
			defer wg.Done()
			for j := 0; j < 1000; j++ {
				counter.Increment()
			}
		}()
	}

	wg.Wait()
	fmt.Printf("Final Value: %d (Expected 50000) <-- CORRECT\n", counter.GetValue())
	fmt.Println("----------------------------------------\n")
}

// ---------------------------------------------------------
// Example 3: Other Atomic Operations
// ---------------------------------------------------------
// Demonstrating Store (Write) and CompareAndSwap (Advanced).
func example3_OtherOperations() {
	fmt.Println("--- Example 3: Store & Swap ---")

	var status int64 = 0 // 0 = OFF, 1 = ON

	// StoreInt64: Safely overwrite value (Write)
	atomic.StoreInt64(&status, 1)
	fmt.Println("Status set to:", atomic.LoadInt64(&status))

	// CompareAndSwapInt64 (CAS):
	// "If value is 1, change it to 0. Otherwise do nothing."
	// Returns true if swap happened.
	swapped := atomic.CompareAndSwapInt64(&status, 1, 0)

	if swapped {
		fmt.Println("CAS Success: Value matched 1, swapped to 0.")
	} else {
		fmt.Println("CAS Failed: Value was not 1.")
	}

	fmt.Println("Current Status:", atomic.LoadInt64(&status))
	fmt.Println("----------------------------------------\n")
}

func main() {
	fmt.Println("═══════════════════════════════════════════════════════════")
	fmt.Println("TOPIC: ATOMIC OPERATIONS (SYNC/ATOMIC)")
	fmt.Println("═══════════════════════════════════════════════════════════\n")

	example1_RaceCondition()

	// Small sleep to separate output logs
	time.Sleep(500 * time.Millisecond)

	example2_AtomicCounter()

	example3_OtherOperations()

	fmt.Println("═══════════════════════════════════════════════════════════")
	fmt.Println("KEY TAKEAWAYS")
	fmt.Println("═══════════════════════════════════════════════════════════")
	fmt.Println(`
1. SPEED: Atomics are much faster than Mutexes for simple counters.
2. POINTERS: Atomic functions ALWAYS require the address of the variable (&val).
3. SCOPE: 
   - Use Atomics for single variable updates (counters, flags).
   - Use Mutexes for updating multiple variables together or complex types (Maps).
	`)
}
