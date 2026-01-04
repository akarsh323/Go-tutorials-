package main

import (
	"context"
	"fmt"
	"time"
)

// TOPIC: Context
// Explanation: The 'context' package is used to carry deadlines, cancellation signals,
// and other request-scoped values across API boundaries and between processes.
// It is critical for controlling cancellation of goroutines.

func main() {
	// Create a context that times out after 2 seconds
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	// Always defer cancel to release resources
	defer cancel()

	select {
	// Simulate work waiting for 5 seconds
	case <-time.After(5 * time.Second):
		fmt.Println("overslept")
	
	// This case triggers when the Context timeout expires
	case <-ctx.Done():
		fmt.Println("context error:", ctx.Err()) // prints "context deadline exceeded"
	}
}
