package main

import (
	"context"
	"fmt"
	"time"
)

/*
TOPIC: THE CONTEXT PACKAGE

CONCEPT:
Think of a Context as a "Box" passed from function to function.
It carries three things:
1. VALUES: Request IDs, User Auth tokens.
2. DEADLINES: Timeouts (e.g., "Stop if this takes > 2 seconds").
3. SIGNALS: Cancellation (e.g., "The user closed the browser, stop working").

ROOT CONTEXTS:
- context.Background(): The standard empty starting point. Use in main() or HTTP handlers.
- context.TODO(): A placeholder when you are unsure what context to use.
*/

// ---------------------------------------------------------
// Example 1: Carrying Values (WithValue)
// ---------------------------------------------------------
// Scenario: We want to attach a "Request ID" to the context so
// every function down the chain knows which request it is handling.
func example_WithValue() {
	fmt.Println("--- Example 1: Context WithValue ---")

	// 1. Start with a clean slate
	rootCtx := context.Background()

	// 2. Add a value to the "Box"
	// Key: "requestID", Value: "12345-ABC"
	// Note: In real apps, use custom types for keys to avoid collisions.
	ctxWithValue := context.WithValue(rootCtx, "requestID", "12345-ABC")

	// 3. Pass the box to the next function
	processRequest(ctxWithValue)
	fmt.Println("----------------------------------------\n")
}

func processRequest(ctx context.Context) {
	// 4. Retrieve value inside the function
	val := ctx.Value("requestID")

	if val != nil {
		fmt.Printf("Processing request with ID: %v\n", val)
	} else {
		fmt.Println("No Request ID found")
	}
}

// ---------------------------------------------------------
// Example 2: Handling Timeouts (WithTimeout)
// ---------------------------------------------------------
// Scenario: Fetching data from a slow database.
// We set a hard deadline. If the task is too slow, we abort.
func example_WithTimeout() {
	fmt.Println("--- Example 2: Context WithTimeout ---")

	bg := context.Background()

	// Create a context that dies automatically after 1 second
	// Returns: ctx (the box) and cancel (a cleanup function)
	ctx, cancel := context.WithTimeout(bg, 1*time.Second)

	// BEST PRACTICE: Always defer cancel() to clean up resources,
	// even if the function finishes early.
	defer cancel()

	fmt.Println("[Main] Requesting slow task (Limit: 1s, Task Needs: 2s)...")

	// We ask for a task that takes 2 seconds (Too long!)
	result := performSlowTask(ctx, 2*time.Second)

	fmt.Println("[Main] Result:", result)
	fmt.Println("----------------------------------------\n")
}

func performSlowTask(ctx context.Context, duration time.Duration) string {
	// We use 'select' to wait for EITHER the work to finish OR the context to die
	select {
	case <-time.After(duration):
		// This simulates the work finishing successfully
		return "Work completed successfully"

	case <-ctx.Done():
		// This channel closes automatically when the timeout is hit
		// ctx.Err() tells us why (Canceled or DeadlineExceeded)
		return fmt.Sprintf("Error: %s", ctx.Err())
	}
}

// ---------------------------------------------------------
// Example 3: Manual Cancellation (WithCancel)
// ---------------------------------------------------------
// Scenario: A "Kill Switch". The user clicks "Cancel" on the UI,
// and we need to stop the background calculation immediately.
func example_WithCancel() {
	fmt.Println("--- Example 3: Context WithCancel (Kill Switch) ---")

	bg := context.Background()

	// Create a context we can cancel manually (no timer)
	ctx, cancel := context.WithCancel(bg)

	// Start a worker in the background
	go func() {
		// Simulate a long running job
		performHeavyWork(ctx)
	}()

	// Simulate Main program running for 2 seconds
	time.Sleep(2 * time.Second)
	fmt.Println("[Main] I've seen enough. Shutting down worker...")

	// PULL THE KILL SWITCH
	cancel()

	// Wait a moment to see the worker output
	time.Sleep(1 * time.Second)
	fmt.Println("----------------------------------------\n")
}

func performHeavyWork(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			// We received the kill signal
			fmt.Println("\t[Worker] I am stopping my work now!")
			fmt.Println("\t[Worker] Reason:", ctx.Err())
			return // Exit function
		default:
			// Simulating ongoing work
			fmt.Println("\t[Worker] ... crunching numbers ...")
			time.Sleep(500 * time.Millisecond)
		}
	}
}

// ---------------------------------------------------------
// Real-World Pattern: Contextual Logging
// ---------------------------------------------------------
func logWithContext(ctx context.Context, msg string) {
	reqID := ctx.Value("requestID")
	if reqID == nil {
		reqID = "unknown"
	}
	fmt.Printf("LOG [ReqID: %v] %s\n", reqID, msg)
}

func main() {
	fmt.Println("═══════════════════════════════════════════════════════════")
	fmt.Println("TOPIC: THE CONTEXT PACKAGE")
	fmt.Println("═══════════════════════════════════════════════════════════\n")

	// 1. Carrying Values
	example_WithValue()

	// 2. Timeouts
	example_WithTimeout()

	// 3. Manual Cancellation
	example_WithCancel()

	fmt.Println("═══════════════════════════════════════════════════════════")
	fmt.Println("TEACHER'S RULES (BEST PRACTICES)")
	fmt.Println("═══════════════════════════════════════════════════════════")
	fmt.Println(`
1. FIRST PARAMETER: Always pass 'ctx context.Context' as the first argument.
   Example: func DoSomething(ctx context.Context, arg int) ...

2. NO GLOBALS: Never store Context in a struct or global variable. 
   Pass it explicitly through the function chain.

3. SCOPED VALUES: Only use WithValue for request-scoped data (IDs, Auth).
   Do not use it for optional function parameters.

4. ALWAYS CANCEL: If you use WithTimeout or WithCancel, always 'defer cancel()'
   in the same function that created the context.
	`)
}
