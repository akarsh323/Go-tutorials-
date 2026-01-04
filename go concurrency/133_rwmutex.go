package main

import (
	"fmt"
	"sync"
	"time"
)

/*
TOPIC: SYNC.RWMUTEX (READ-WRITE LOCK)

CONCEPT:
A standard Mutex (sync.Mutex) is a "Block Everyone" lock.
A RWMutex (sync.RWMutex) is smarter:
1. RLock (Read Lock): Allows MULTIPLE readers at the same time.
   "We are just looking, not touching."
2. Lock (Write Lock): Exclusive access. Blocks ALL readers and writers.
   "I am changing the data, everyone stay out."

WHEN TO USE:
Use this when you have MANY readers and FEW writers (e.g., Caches, Configs).
*/

var (
	// The shared resource
	counter int

	// The Smarter Lock
	// We use RWMutex instead of Mutex
	rwMu sync.RWMutex
)

// ---------------------------------------------------------
// The Reader (The Crowd)
// ---------------------------------------------------------
// Uses RLock() / RUnlock()
// Multiple Goroutines can run this function simultaneously!
func readCounter(wg *sync.WaitGroup, id int) {
	defer wg.Done()

	// 1. ACQUIRE READ LOCK
	// "I want to read. If a writer is working, I wait.
	// If other readers are working, I join them."
	rwMu.RLock()

	fmt.Printf("   [Reader %d] Seeing value: %d\n", id, counter)
	time.Sleep(1 * time.Millisecond) // Simulate reading time

	// 2. RELEASE READ LOCK
	rwMu.RUnlock()
}

// ---------------------------------------------------------
// The Writer (The Librarian)
// ---------------------------------------------------------
// Uses Lock() / Unlock()
// Only ONE Goroutine can run this. It blocks everyone else.
func writeCounter(wg *sync.WaitGroup, newValue int) {
	defer wg.Done()

	// 1. ACQUIRE WRITE LOCK
	// "I want to write. Wait until all active readers are done.
	// Then block everyone else until I am finished."
	rwMu.Lock()

	fmt.Printf("\n--- [Writer] Updating counter to %d (Blocking everyone) ---\n", newValue)
	counter = newValue
	time.Sleep(200 * time.Millisecond) // Simulate slow write
	fmt.Println("--- [Writer] Update complete. Releasing lock. ---\n")

	// 2. RELEASE WRITE LOCK
	rwMu.Unlock()
}

func main() {
	fmt.Println("═══════════════════════════════════════════════════════════")
	fmt.Println("TOPIC: SYNC.RWMUTEX")
	fmt.Println("═══════════════════════════════════════════════════════════\n")

	var wg sync.WaitGroup

	// 1. Launch Batch 1 of Readers (Fast, Concurrent)
	fmt.Println("[Main] Launching first batch of 5 Readers...")
	for i := 1; i <= 5; i++ {
		wg.Add(1)
		go readCounter(&wg, i)
	}

	// Small sleep to let readers start
	time.Sleep(10 * time.Millisecond)

	// 2. Launch the Writer (Slow, Exclusive)
	// This will wait for the first batch to finish, then lock the door.
	wg.Add(1)
	go writeCounter(&wg, 100)

	// 3. Launch Batch 2 of Readers
	// These will be BLOCKED until the Writer finishes.
	fmt.Println("[Main] Launching second batch of 5 Readers...")
	for i := 6; i <= 10; i++ {
		wg.Add(1)
		go readCounter(&wg, i)
	}

	wg.Wait()
	fmt.Println("All operations complete.")

	fmt.Println("═══════════════════════════════════════════════════════════")
	fmt.Println("KEY TAKEAWAYS")
	fmt.Println("═══════════════════════════════════════════════════════════")
	fmt.Println(`
1. CONCURRENCY: Notice how the first batch of readers prints almost instantly.
   They don't block each other.
   
2. EXCLUSIVITY: Notice how the second batch waits. They cannot start until the
   Writer prints "Releasing lock".

3. PERFORMANCE: If you used a standard Mutex here, every single reader would
   have to wait in a single-file line. RWMutex creates a "highway" for readers.
	`)
}
