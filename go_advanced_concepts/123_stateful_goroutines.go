**Topic: Stateful Goroutines**

This file implements the **Stateful Worker** pattern discussed.
It demonstrates how to maintain internal state (memory) inside a Goroutine safely using channels, without needing Mutex locks.

You can save this code in a file named `main.go` and run it using `go run main.go`.

```go
package main

import (
	"fmt"
	"time"
)

/*
TOPIC: STATEFUL GOROUTINES

CONCEPT:
A Stateful Goroutine is a process that remembers data (history) across multiple calls.
Instead of using Global Variables or Mutexes to protect this data, we use
a single Goroutine that "owns" the data.

THE PATTERN:
1. The State: Kept inside a struct (e.g., 'count').
2. The Guard: A background Goroutine that loops forever.
3. The Input: A channel. The loop reads from the channel to update state.

WHY IT WORKS:
Channels naturally serialize data. Even if 100 functions call .Send() at once,
the Goroutine receives them one-by-one, preventing Race Conditions perfectly.
*/

// ---------------------------------------------------------
// 1. The Blueprint (Encapsulation)
// ---------------------------------------------------------

// StatefulWorker holds the state and the communication line.
type StatefulWorker struct {
	count int      // INTERNAL STATE: Only the background goroutine touches this.
	ch    chan int // INPUT: The only way to talk to the worker.
}

// NewStatefulWorker is a constructor to initialize the worker safely.
func NewStatefulWorker() *StatefulWorker {
	return &StatefulWorker{
		count: 0,
		ch:    make(chan int),
	}
}

// ---------------------------------------------------------
// 2. The Engine (Start Method)
// ---------------------------------------------------------

// Start initializes the background process.
func (w *StatefulWorker) Start() {
	fmt.Println("[Worker] Engine Starting...")

	// Launch the "Guard" Goroutine
	go func() {
		// Infinite loop keeps the state alive
		for {
			// Wait for data to arrive on the channel
			select {
			case value := <-w.ch:
				// --- CRITICAL SECTION (Implicit) ---
				// No Mutex needed! Only this single goroutine runs this code.
				
				// Calculate new state based on history
				prev := w.count
				w.count += value
				
				fmt.Printf("   -> Received: %d | (Old: %d + New: %d) = Total: %d\n", 
					value, prev, value, w.count)
			}
		}
	}()
}

// ---------------------------------------------------------
// 3. The Interface (Send Method)
// ---------------------------------------------------------

// Send is a helper method. It hides the complexity of channels from the user.
// The user just calls .Send(5), they don't need to know about "<-w.ch".
func (w *StatefulWorker) Send(value int) {
	w.ch <- value
}

// ---------------------------------------------------------
// 4. Execution (Main)
// ---------------------------------------------------------

func main() {
	fmt.Println("═══════════════════════════════════════════════════════════")
	fmt.Println("TOPIC: STATEFUL GOROUTINES (THE ACCUMULATOR)")
	fmt.Println("═══════════════════════════════════════════════════════════\n")

	// 1. Initialize
	streetWorker := NewStatefulWorker()

	// 2. Start the Engine
	streetWorker.Start()

	// 3. Interaction Loop
	// We simulate a stream of data coming in (0, 1, 2, 3, 4)
	fmt.Println("[Main] Starting data stream...")
	
	for i := 0; i < 5; i++ {
		fmt.Printf("[Main] Sending value: %d\n", i)
		streetWorker.Send(i)
		
		// Sleep purely to make the output readable and distinct
		time.Sleep(500 * time.Millisecond)
	}

	fmt.Println("\n[Main] Stream finished.")
	
	// Wait a moment before exiting so the final print can finish
	time.Sleep(100 * time.Millisecond)
	
	fmt.Println("═══════════════════════════════════════════════════════════")
	fmt.Println("KEY TAKEAWAYS")
	fmt.Println("═══════════════════════════════════════════════════════════")
	fmt.Println(`
1. NO MUTEX NEEDED: Since only ONE goroutine accesses 'w.count', 
   we cannot have a race condition.
   
2. SEQUENTIAL ACCESS: The channel forces data to line up single-file.
   Even if Main sent 100 items instantly, the Worker processes them one by one.

3. PERSISTENCE: The 'count' variable lives as long as the for-loop runs.
   This is how we give "memory" to a concurrent process.
	`)
}

```