package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

/*
TOPIC: SIGNAL HANDLING (OS/SIGNAL)

CONCEPT:
Signals are messages from the Operating System to your program.
- SIGINT (Ctrl+C): "Please interrupt what you are doing."
- SIGTERM (Kill): "Please terminate cleanly."
- SIGHUP: "Configuration changed, please reload."

THE GOAL:
Prevent the program from crashing instantly. Catch the signal, clean up resources
(close DBs, save files), and THEN exit gracefully.
*/

// ---------------------------------------------------------
// Example 1: Basic Interceptor (Catching Ctrl+C)
// ---------------------------------------------------------
// Simple implementation: Blocks until Ctrl+C is pressed.
func example1_BasicInterceptor() {
	fmt.Println("--- Example 1: Basic Signal Interceptor ---")

	// 1. Create a Buffered Channel
	// PRO TIP: Always buffer signal channels (size 1) so signals aren't dropped
	// if your program is busy processing when the signal arrives.
	sigChan := make(chan os.Signal, 1)

	// 2. Register the channel to receive notifications
	// We only listen for SIGINT (Ctrl+C) here.
	signal.Notify(sigChan, syscall.SIGINT)

	fmt.Println("Program running... Press Ctrl+C to stop.")

	// 3. BLOCKING WAIT
	// The program pauses here until the OS sends the signal.
	sig := <-sigChan

	fmt.Printf("\nReceived Signal: %s\n", sig)
	fmt.Println("Gracefully shutting down...")
	// os.Exit(0) would happen here in a real app
}

// ---------------------------------------------------------
// Example 2: Handling Multiple Signals (Switch)
// ---------------------------------------------------------
// Distinguishing between "Stop" (INT/TERM) and "Reload" (HUP).
// NOTE: To test SIGHUP, you need to use 'kill -SIGHUP [PID]' from another terminal.
func example2_MultipleSignals() {
	fmt.Println("--- Example 2: Multiple Signals (Switch) ---")

	// Get PID so you can send signals manually from another terminal
	fmt.Printf("Process ID: %d\n", os.Getpid())
	fmt.Println("Waiting for signals (INT, TERM, HUP, or USR1)...")

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan,
		syscall.SIGINT,
		syscall.SIGTERM,
		syscall.SIGHUP,
		syscall.SIGUSR1,
	)

	// Loop to handle signals continuously (e.g., allow multiple reloads)
	for {
		sig := <-sigChan

		switch sig {
		case syscall.SIGINT:
			fmt.Println("\nReceived SIGINT (Ctrl+C). Stopping.")
			return // Exit function

		case syscall.SIGTERM:
			fmt.Println("\nReceived SIGTERM (Kubernetes/Docker stop). Stopping.")
			return // Exit function

		case syscall.SIGHUP:
			fmt.Println("\nReceived SIGHUP. Reloading configuration (Mock)...")
			// Do not return; keep running!

		case syscall.SIGUSR1:
			fmt.Println("\nReceived SIGUSR1. Executing custom user logic...")
		}
	}
}

// ---------------------------------------------------------
// Example 3: The Professional Pattern (Worker Coordination)
// ---------------------------------------------------------
// This is how production apps work.
// 1. Main listens for OS Signal.
// 2. Main tells Worker to stop (via 'done' channel).
// 3. Worker cleans up and stops.
// 4. Main exits.

func worker(done chan bool) {
	fmt.Println("[Worker] Starting infinite loop...")
	for {
		select {
		case <-done:
			// 3. Handle Shutdown
			fmt.Println("[Worker] Stop signal received! Cleaning up...")
			time.Sleep(500 * time.Millisecond) // Simulate cleanup time
			fmt.Println("[Worker] Cleanup complete. Bye!")
			return // Exit the loop

		default:
			// 4. Do Normal Work
			fmt.Print(".")
			time.Sleep(500 * time.Millisecond)
		}
	}
}

func example3_GracefulShutdown() {
	fmt.Println("--- Example 3: Graceful Worker Shutdown ---")

	// Setup Channels
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	done := make(chan bool)

	// Start Worker
	go worker(done)

	fmt.Println("[Main] Program running. Press Ctrl+C to trigger graceful shutdown.")

	// Block until signal
	sig := <-sigChan
	fmt.Printf("\n[Main] Received %s. Notifying worker...\n", sig)

	// Notify Worker
	done <- true

	// Allow time for worker to print its cleanup message
	// In production, you would use a WaitGroup here instead of sleep!
	time.Sleep(1 * time.Second)
	fmt.Println("[Main] Exiting.")
}

func main() {
	fmt.Println("═══════════════════════════════════════════════════════════")
	fmt.Println("TOPIC: OS SIGNAL HANDLING")
	fmt.Println("═══════════════════════════════════════════════════════════")
	fmt.Println("NOTE: Since these examples involve blocking/stopping,")
	fmt.Println("uncomment the specific example you want to run in main().")
	fmt.Println("═══════════════════════════════════════════════════════════\n")

	// UNCOMMENT ONE OF THE BELOW:

	// example1_BasicInterceptor()

	// example2_MultipleSignals()

	example3_GracefulShutdown()

	fmt.Println("\n═══════════════════════════════════════════════════════════")
	fmt.Println("KEY TAKEAWAYS")
	fmt.Println("═══════════════════════════════════════════════════════════")
	fmt.Println(`
1. BUFFERING: Always use 'make(chan os.Signal, 1)'.
2. NOTIFY: Use 'signal.Notify' to bind the channel to specific signals.
3. SELECT: Use the 'select' statement in workers to listen for the
   stop signal while performing work.
4. GRACEFUL: Do not just os.Exit(). Close DBs and files first.
	`)
}
