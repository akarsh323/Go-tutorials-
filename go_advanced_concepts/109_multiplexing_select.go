package main

import (
	"fmt"
	"time"
)

/*
TOPIC: MULTIPLEXING WITH SELECT

CONCEPT:
The 'select' statement is the Traffic Controller of Go concurrency.
It allows a single Goroutine to wait on multiple channel operations simultaneously.
It blocks until ONE of its cases is ready to run.

KEY FEATURES:
1. NON-BLOCKING: Use 'default' to act immediately if no channels are ready.
2. RACE HANDLING: If multiple channels are ready, 'select' picks one at random.
3. TIMEOUTS: Use 'time.After' to prevent hanging forever.
4. SAFETY: Use 'val, ok := <-ch' to detect if a channel is closed.
*/

// ---------------------------------------------------------
// Example 1: The "Naive" Approach (Blocking/Deadlock)
// ---------------------------------------------------------
// Without 'select', reading from an empty channel blocks the program.
// If no one is sending, this causes a Deadlock crash.
func example1_NaiveBlocking() {
	fmt.Println("--- Example 1: The Problem (Blocking) ---")

	// Uncommenting the code below will crash the program!
	/*
		ch1 := make(chan int)
		ch2 := make(chan int)

		// BLOCKS FOREVER -> DEADLOCK
		fmt.Println(<-ch1)
		fmt.Println(<-ch2)
	*/
	fmt.Println("(Skipped to prevent crash. See code comments.)\n")
}

// ---------------------------------------------------------
// Example 2: Non-Blocking Select with 'default'
// ---------------------------------------------------------
// The 'default' case runs immediately if no other channels are ready.
func example2_NonBlocking() {
	fmt.Println("--- Example 2: Non-Blocking with 'default' ---")

	ch1 := make(chan int)
	ch2 := make(chan int)

	select {
	case msg1 := <-ch1:
		fmt.Println("Received from ch1:", msg1)
	case msg2 := <-ch2:
		fmt.Println("Received from ch2:", msg2)
	default:
		// Because ch1 and ch2 are empty, this runs INSTANTLY.
		fmt.Println("No channels ready yet! Moving on...")
	}
	fmt.Println("-----------------------------------------------\n")
}

// ---------------------------------------------------------
// Example 3: The Race (Handling Multiple Channels)
// ---------------------------------------------------------
// We start two timers. 'select' waits for whichever finishes first.
func example3_TheRace() {
	fmt.Println("--- Example 3: The Race (Multiple Channels) ---")

	ch1 := make(chan string)
	ch2 := make(chan string)

	// Worker 1: 1 second
	go func() {
		time.Sleep(1 * time.Second)
		ch1 <- "Turtle (1s)"
	}()

	// Worker 2: 2 seconds
	go func() {
		time.Sleep(2 * time.Second)
		ch2 <- "Rabbit (2s)"
	}()

	fmt.Println("Waiting for winners...")

	// We want to hear from both, so we loop twice.
	for i := 0; i < 2; i++ {
		select {
		case msg1 := <-ch1:
			fmt.Println("Winner:", msg1)
		case msg2 := <-ch2:
			fmt.Println("Winner:", msg2)
		}
	}
	fmt.Println("-----------------------------------------------\n")
}

// ---------------------------------------------------------
// Example 4: The Timeout Pattern
// ---------------------------------------------------------
// Critical for production. Don't wait forever for a slow server.
func example4_Timeout() {
	fmt.Println("--- Example 4: The Timeout Pattern ---")

	ch := make(chan string)

	// Simulate a slow server taking 3 seconds
	go func() {
		time.Sleep(3 * time.Second)
		ch <- "Server Response"
	}()

	fmt.Println("Waiting for server (Timeout set to 1s)...")

	select {
	case res := <-ch:
		fmt.Println("Success:", res)
	// time.After returns a channel that fires after duration
	case <-time.After(1 * time.Second):
		fmt.Println("Error: Operation timed out!")
	}
	fmt.Println("-----------------------------------------------\n")
}

// ---------------------------------------------------------
// Example 5: The "Comma Ok" Idiom (Graceful Shutdown)
// ---------------------------------------------------------
// How to stop listening when the sender closes the channel.
func example5_GracefulShutdown() {
	fmt.Println("--- Example 5: Graceful Shutdown (Comma Ok) ---")

	dataChan := make(chan int)

	// Producer
	go func() {
		for i := 1; i <= 3; i++ {
			dataChan <- i
			time.Sleep(200 * time.Millisecond)
		}
		close(dataChan) // Signal completion
		fmt.Println("[Sender] Channel Closed.")
	}()

	// Consumer
	looping := true
	for looping {
		select {
		case msg, ok := <-dataChan:
			if !ok {
				// ok is FALSE if channel is closed & empty
				fmt.Println("[Receiver] Channel closed. Stopping loop.")
				looping = false
			} else {
				fmt.Println("[Receiver] Received:", msg)
			}

		// We can still use a safety timeout
		case <-time.After(2 * time.Second):
			fmt.Println("Timeout safety trigger.")
			looping = false
		}
	}
	fmt.Println("-----------------------------------------------\n")
}

func main() {
	fmt.Println("═══════════════════════════════════════════════════════════")
	fmt.Println("TOPIC: MULTIPLEXING WITH SELECT")
	fmt.Println("═══════════════════════════════════════════════════════════\n")

	example1_NaiveBlocking()
	example2_NonBlocking()
	example3_TheRace()
	example4_Timeout()
	example5_GracefulShutdown()

	fmt.Println("═══════════════════════════════════════════════════════════")
	fmt.Println("KEY TAKEAWAYS")
	fmt.Println("═══════════════════════════════════════════════════════════")
	fmt.Println(`
1. TRAFFIC CONTROL: 'select' lets one Goroutine manage multiple channels.
2. DEFAULT: Use 'default' to avoid blocking if you don't want to wait.
3. TIMEOUTS: Always use 'time.After' in production to prevent hanging.
4. CLOSING: Always check 'ok' (msg, ok := <-ch) to handle closed channels
   gracefully instead of receiving infinite zero-values.
	`)
}
