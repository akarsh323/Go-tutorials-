package main

import (
	"fmt"
	"time"
)

/*
TOPIC: TICKERS (REPEATING TIMERS)

CONCEPT:
A Ticker is a clock that ticks repeatedly at a regular interval.
It sends the current time on its channel 'C' every X duration.

DIFFERENCE FROM TIMERS:
- Timer: Fires ONCE. (Kitchen Timer)
- Ticker: Fires REPEATEDLY. (Metronome / Heartbeat)

CRITICAL RULE:
You MUST stop a Ticker when you are done.
A Timer eventually stops itself (when it fires), but a Ticker runs forever
and will cause a "Memory Leak" if not stopped manually using ticker.Stop().
*/

// ---------------------------------------------------------
// Example 1: The Basic Ticker
// ---------------------------------------------------------
// This demonstrates how to iterate over a ticker using a loop.
func example1_BasicTicker() {
	fmt.Println("--- Example 1: Basic Ticker ---")

	// 1. Create a Ticker that fires every 500ms
	ticker := time.NewTicker(500 * time.Millisecond)

	// Best Practice: Defer stop to ensure cleanup
	defer ticker.Stop()

	fmt.Println("[Main] Ticker started (3 ticks)...")

	// 2. Loop to consume ticks
	// We'll manually count 3 ticks then exit
	count := 0
	for t := range ticker.C {
		fmt.Println("Tick at:", t.Format("15:04:05.000"))
		count++
		if count == 3 {
			fmt.Println("[Main] Stopping loop.")
			break // Break the loop, defer will stop the ticker
		}
	}
	fmt.Println("----------------------------------------\n")
}

// ---------------------------------------------------------
// Example 2: The Heartbeat Pattern (Select)
// ---------------------------------------------------------
// Using a Ticker inside a 'select' loop to perform a task periodically
// while still listening for other events (like a stop signal).
func example2_Heartbeat() {
	fmt.Println("--- Example 2: The Heartbeat (Select) ---")

	ticker := time.NewTicker(1 * time.Second)
	done := make(chan bool)

	// Start a worker that runs for 3.5 seconds
	go func() {
		time.Sleep(3500 * time.Millisecond)
		done <- true
	}()

	fmt.Println("[System] Monitoring System Health...")

	// LOOP: Wait for Tick OR Done signal
	running := true
	for running {
		select {
		case t := <-ticker.C:
			fmt.Println("♥ Heartbeat check at:", t.Format("15:04:05"))

		case <-done:
			fmt.Println("[System] Shutdown signal received!")
			running = false
		}
	}

	ticker.Stop() // Cleanup
	fmt.Println("[System] Stopped.")
	fmt.Println("----------------------------------------\n")
}

// ---------------------------------------------------------
// Example 3: The Rate Limiter (Simulation)
// ---------------------------------------------------------
// Tickers are often used to limit how fast a program does something.
// Here we process a queue of "requests" but only 1 per 200ms.
func example3_RateLimiter() {
	fmt.Println("--- Example 3: Rate Limiter ---")

	// A queue of 5 requests
	requests := make(chan int, 5)
	for i := 1; i <= 5; i++ {
		requests <- i
	}
	close(requests) // Close channel so range loop knows when to stop

	// A limiter that ticks every 200ms
	limiter := time.NewTicker(200 * time.Millisecond)
	defer limiter.Stop()

	fmt.Println("[Worker] Processing requests (Max 5/sec)...")

	for req := range requests {
		// BLOCKING WAIT:
		// We wait for the tick before processing the next item.
		// This forces the loop to slow down.
		<-limiter.C
		fmt.Printf("  -> Processed Request #%d\n", req)
	}
	fmt.Println("----------------------------------------\n")
}

// ---------------------------------------------------------
// Example 4: The Trap (time.Tick vs NewTicker)
// ---------------------------------------------------------
// time.Tick() is a wrapper that creates a ticker but gives you NO WAY to stop it.
// It is "Fire and Forget".
// ONLY use this in main() or infinite loops. Never in short-lived functions.
func example4_TimeTickTrap() {
	fmt.Println("--- Example 4: time.Tick (The Trap) ---")

	fmt.Println("Be careful with 'time.Tick(d)'!")
	fmt.Println("It returns a channel, not the Ticker struct.")
	fmt.Println("You cannot call .Stop() on it.")
	fmt.Println("If used in a function that runs many times, it causes a MEMORY LEAK.")

	// Safe usage: Quick scripts or Main function logic
	// c := time.Tick(1 * time.Second)
	// for now := range c { ... }

	fmt.Println("(Skipping execution to avoid waiting)")
	fmt.Println("----------------------------------------\n")
}

func main() {
	fmt.Println("═══════════════════════════════════════════════════════════")
	fmt.Println("TOPIC: TICKERS IN GO")
	fmt.Println("═══════════════════════════════════════════════════════════\n")

	example1_BasicTicker()
	example2_Heartbeat()
	example3_RateLimiter()
	example4_TimeTickTrap()

	fmt.Println("═══════════════════════════════════════════════════════════")
	fmt.Println("KEY TAKEAWAYS")
	fmt.Println("═══════════════════════════════════════════════════════════")
	fmt.Println(`
1. REPETITION: Use Tickers for events that repeat (Heartbeats, Polling, Animation).
2. CLEANUP: ALWAYS call 'defer ticker.Stop()' to free resources. Tickers do not
   garbage collect automatically if the function ends!
3. SELECTION: Use Tickers inside 'select' to handle repeating tasks without
   blocking the entire program.
4. RATE LIMITING: Use a Ticker to throttle fast loops (e.g., API requests).
	`)
}
