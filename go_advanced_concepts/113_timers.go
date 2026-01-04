package main

import (
	"fmt"
	"time"
)

/*
TOPIC: TIMERS

CONCEPT:
A Timer is like a kitchen egg timer. You set it for a duration.
When the time is up, it sends a value (the current time) on its channel 'C'.

DIFFERENCE FROM TIME.SLEEP:
- time.Sleep: Freezes the current Goroutine completely. (The "Zombie" method).
- time.Timer: Runs in the background. You can do other work while waiting,
  or cancel it before it rings. (The "Smart" method).
*/

// ---------------------------------------------------------
// Example 1: The Basic Timer (Blocking Wait)
// ---------------------------------------------------------
// This looks like time.Sleep, but uses a channel.
// We wait for the "beep" (signal) from the timer.
func example1_BasicTimer() {
	fmt.Println("--- Example 1: Basic Timer ---")

	// 1. Create a timer for 2 seconds
	// 'timer' is a struct containing a channel named 'C'
	timer := time.NewTimer(2 * time.Second)

	fmt.Println("[Main] Timer started. Blocking until it fires...")

	// 2. BLOCKING WAIT
	// We pause here, waiting for a value to appear in timer.C
	expirationTime := <-timer.C

	fmt.Println("[Main] Timer fired at:", expirationTime)
	fmt.Println("----------------------------------------\n")
}

// ---------------------------------------------------------
// Example 2: Stopping and Resetting
// ---------------------------------------------------------
// Unlike Sleep, we can stop a Timer before it finishes.
// This is useful for resource cleanup or changing deadlines.
func example2_StopAndReset() {
	fmt.Println("--- Example 2: Stopping & Resetting ---")

	timer := time.NewTimer(5 * time.Second)

	// Launch a Goroutine to wait for it
	go func() {
		<-timer.C
		fmt.Println("[Goroutine] Timer fired! (This won't happen)")
	}()

	// STOP the timer immediately
	// Returns true if stopped, false if already fired
	stopped := timer.Stop()
	if stopped {
		fmt.Println("[Main] Timer stopped before it could fire.")
	}

	// RESET the timer
	// We reuse the same timer for a new duration (1 second)
	fmt.Println("[Main] Resetting timer for 1 second...")
	timer.Reset(1 * time.Second)

	// Wait for the new duration
	<-timer.C
	fmt.Println("[Main] Reset timer fired successfully.")
	fmt.Println("----------------------------------------\n")
}

// ---------------------------------------------------------
// Example 3: The Timeout Pattern (time.After)
// ---------------------------------------------------------
// This is the #1 use case in production.
// "Try to do X, but if it takes too long, stop."
func example3_TimeoutPattern() {
	fmt.Println("--- Example 3: The Timeout Pattern ---")

	// Channel to signal work completion
	done := make(chan string)

	// Start a "slow" worker (3 seconds)
	go func() {
		time.Sleep(3 * time.Second)
		done <- "Work Complete"
	}()

	fmt.Println("[Main] Race Started: Work (3s) vs Timeout (2s)")

	// The Race: Wait for Work OR Timer
	select {
	case res := <-done:
		fmt.Println("[Main] Success:", res)

	// time.After creates a temporary timer and returns the channel directly
	case <-time.After(2 * time.Second):
		fmt.Println("[Main] Error: Operation timed out!")
	}
	fmt.Println("----------------------------------------\n")
}

// ---------------------------------------------------------
// Example 4: Fire-and-Forget (Background Delay)
// ---------------------------------------------------------
// Schedule a task for the future without stopping the Main thread.
func example4_BackgroundDelay() {
	fmt.Println("--- Example 4: Background Execution ---")

	timer := time.NewTimer(2 * time.Second)

	// Schedule task in background
	go func() {
		// This blocks ONLY this goroutine
		<-timer.C
		fmt.Println(">>> [Background Job] I am running 2 seconds later!")
	}()

	fmt.Println("[Main] I am continuing to do other work...")
	fmt.Println("[Main] (Sleeping 3s to let background job finish)")
	time.Sleep(3 * time.Second)
	fmt.Println("----------------------------------------\n")
}

func main() {
	fmt.Println("═══════════════════════════════════════════════════════════")
	fmt.Println("TOPIC: TIMERS IN GO")
	fmt.Println("═══════════════════════════════════════════════════════════\n")

	example1_BasicTimer()
	example2_StopAndReset()
	example3_TimeoutPattern()
	example4_BackgroundDelay()

	fmt.Println("═══════════════════════════════════════════════════════════")
	fmt.Println("KEY TAKEAWAYS")
	fmt.Println("═══════════════════════════════════════════════════════════")
	fmt.Println(`
1. TIME.AFTER: Use this inside 'select' for simple timeouts.
   It's a clean one-liner: case <-time.After(2 * time.Second):

2. TIME.NEWTIMER: Use this when you need a reference to the timer, specifically
   if you need to STOP it or RESET it.

3. CLEANUP: If you create a NewTimer but finish work early, call timer.Stop()
   to release resources immediately.

4. NON-BLOCKING: Timers allow your code to wait for time *concurrently* with
   other events (like user input or network requests).
	`)
}
