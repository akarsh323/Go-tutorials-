package main

import (
	"fmt"
	"sync"
	"time"
)

/*
TOPIC: SYNC.COND (CONDITION VARIABLES)

CONCEPT:
A Condition Variable is a mechanism to block a Goroutine until a specific condition
is met (e.g., "Wait until buffer is not full").

WHY NOT BUSY WAIT?
Instead of spinning in a loop (consuming 100% CPU) checking 'if len < 5',
we go to sleep using Cond.Wait(). We are woken up only when signaled.

THE PATTERN:
1. Lock Mutex.
2. Check Condition in a Loop (for !condition { cond.Wait() }).
3. Do Work.
4. Signal others (cond.Signal()).
5. Unlock Mutex.
*/

// Buffer Configuration
const BufferSize = 5

// Buffer is a thread-safe slice of integers.
type Buffer struct {
	items []int      // The shared data
	mu    sync.Mutex // The lock protecting the data
	cond  *sync.Cond // The signaling mechanism
}

// NewBuffer initializes the struct.
func NewBuffer() *Buffer {
	b := &Buffer{
		items: make([]int, 0, BufferSize),
	}
	// CRITICAL: sync.Cond must be linked to the specific Mutex protecting the data.
	b.cond = sync.NewCond(&b.mu)
	return b
}

// ---------------------------------------------------------
// The Producer (Adds items)
// ---------------------------------------------------------
func (b *Buffer) Produce(item int) {
	// 1. ACQUIRE LOCK
	b.mu.Lock()
	defer b.mu.Unlock()

	// 2. CHECK CONDITION (Loop is mandatory!)
	// "If buffer is full, I must wait."
	for len(b.items) == BufferSize {
		fmt.Println("[Producer] Buffer Full. Waiting...")

		// WAIT:
		// 1. Atomically Unlocks b.mu
		// 2. Suspends execution (Sleeps)
		// 3. When woken, Locks b.mu again before returning
		b.cond.Wait()
	}

	// 3. ACT (Critical Section)
	b.items = append(b.items, item)
	fmt.Printf("   -> Produced: %d (Size: %d)\n", item, len(b.items))

	// 4. SIGNAL
	// Wake up a waiting consumer (if any).
	// "Hey, the buffer is no longer empty!"
	b.cond.Signal()
}

// ---------------------------------------------------------
// The Consumer (Removes items)
// ---------------------------------------------------------
func (b *Buffer) Consume() int {
	// 1. ACQUIRE LOCK
	b.mu.Lock()
	defer b.mu.Unlock()

	// 2. CHECK CONDITION
	// "If buffer is empty, I must wait."
	for len(b.items) == 0 {
		fmt.Println("[Consumer] Buffer Empty. Waiting...")
		b.cond.Wait()
	}

	// 3. ACT (FIFO Removal)
	item := b.items[0]
	b.items = b.items[1:]
	fmt.Printf("<- Consumed: %d (Size: %d)\n", item, len(b.items))

	// 4. SIGNAL
	// Wake up a waiting producer (if any).
	// "Hey, the buffer is no longer full!"
	b.cond.Signal()

	return item
}

// ---------------------------------------------------------
// Execution
// ---------------------------------------------------------
func main() {
	fmt.Println("═══════════════════════════════════════════════════════════")
	fmt.Println("TOPIC: SYNC.COND (PRODUCER / CONSUMER)")
	fmt.Println("═══════════════════════════════════════════════════════════\n")

	buffer := NewBuffer()
	var wg sync.WaitGroup

	// 1. Start Producer
	// It produces faster (100ms) than the consumer consumes.
	// It will eventually hit the limit (5) and block.
	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 1; i <= 10; i++ {
			buffer.Produce(i)
			time.Sleep(100 * time.Millisecond)
		}
	}()

	// 2. Start Consumer
	// It consumes slower (300ms).
	wg.Add(1)
	go func() {
		defer wg.Done()
		// Give the producer a head start to prove the wait logic works
		time.Sleep(500 * time.Millisecond)

		for i := 1; i <= 10; i++ {
			buffer.Consume()
			time.Sleep(300 * time.Millisecond)
		}
	}()

	wg.Wait()
	fmt.Println("\nDone. All items processed.")

	fmt.Println("═══════════════════════════════════════════════════════════")
	fmt.Println("KEY TAKEAWAYS")
	fmt.Println("═══════════════════════════════════════════════════════════")
	fmt.Println(`
1. THE LOOP: Always wrap cond.Wait() in a 'for' loop, never an 'if'.
   Why? Because when you wake up, the condition might have changed again
   before you grabbed the lock (Spurious Wakeups).

2. THE MAGIC UNLOCK: cond.Wait() automatically unlocks the mutex while sleeping
   and re-locks it when waking up. This allows other goroutines to enter and
   change the state.

3. SIGNAL VS BROADCAST:
   - Signal(): Wakes 1 waiter (efficient for 1-in-1-out logic).
   - Broadcast(): Wakes ALL waiters (useful for 'Shutdown' signals).
	`)
}
