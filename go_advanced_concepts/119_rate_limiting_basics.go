package main

import (
	"fmt"
	"time"
)

/*
TOPIC: BASIC RATE LIMITING

CONCEPT:
We use channels and tickers to control the speed of execution.
1. STRICT: Uses a Ticker. Forces a wait between EVERY request.
2. BURSTY: Uses a Buffered Channel. Allows a small group of requests to execute
   instantly (up to buffer size) before slowing down.
*/

// ---------------------------------------------------------
// Example 1: Strict Limiting (The Steady Drip)
// ---------------------------------------------------------
// Scenario: We want to process requests, but NEVER faster than
// one request every 200 milliseconds.
func strictLimiter() {
	fmt.Println("--- Example 1: Strict Rate Limiting ---")

	// 1. Create a queue of 5 requests
	requests := make(chan int, 5)
	for i := 1; i <= 5; i++ {
		requests <- i
	}
	close(requests)

	// 2. Create a Ticker (The regulator)
	// This ticks exactly once every 200ms.
	limiter := time.NewTicker(200 * time.Millisecond)
	defer limiter.Stop()

	fmt.Println("[System] Processing requests (1 per 200ms)...")

	for req := range requests {
		// BLOCKING WAIT:
		// We wait for the tick here. This forces the delay.
		<-limiter.C

		fmt.Printf("Request %d processed at %v\n", req, time.Now().Format("15:04:05.000"))
	}
	fmt.Println("----------------------------------------\n")
}

// ---------------------------------------------------------
// Example 2: Bursty Limiting (Token Bucket)
// ---------------------------------------------------------
// Scenario: We want to allow users to do 3 things instantly (Burst),
// but then slow them down to 1 request every 200ms.
func burstyLimiter() {
	fmt.Println("--- Example 2: Bursty Rate Limiting ---")

	// 1. Create a queue of 10 requests (More than the burst limit)
	requests := make(chan int, 10)
	for i := 1; i <= 10; i++ {
		requests <- i
	}
	close(requests)

	// 2. Create the "Token Bucket" (Buffered Channel)
	// Capacity of 3 means we can hold 3 tokens for a burst.
	burstLimit := 3
	limiter := make(chan time.Time, burstLimit)

	// 3. Fill the bucket initially (Give the user their starting tokens)
	for i := 0; i < burstLimit; i++ {
		limiter <- time.Now()
	}

	// 4. Start a Refiller Goroutine
	// Adds a new token every 200ms to refill what was used.
	go func() {
		for t := range time.Tick(200 * time.Millisecond) {
			// Try to add a token. If full, just ignore (non-blocking).
			select {
			case limiter <- t:
			default:
			}
		}
	}()

	fmt.Println("[System] Processing requests (Burst 3, then 1 per 200ms)...")

	for req := range requests {
		// Wait for a token.
		// The first 3 will happen INSTANTLY because the channel is full.
		// The rest will wait for the refiller.
		<-limiter

		fmt.Printf("Request %d processed at %v\n", req, time.Now().Format("15:04:05.000"))
	}
	fmt.Println("----------------------------------------\n")
}

func main() {
	fmt.Println("═══════════════════════════════════════════════════════════")
	fmt.Println("TOPIC: BASICS OF RATE LIMITING")
	fmt.Println("═══════════════════════════════════════════════════════════\n")

	strictLimiter()

	burstyLimiter()

	fmt.Println("═══════════════════════════════════════════════════════════")
	fmt.Println("KEY TAKEAWAYS")
	fmt.Println("═══════════════════════════════════════════════════════════")
	fmt.Println(`
1. TICKER: Simplest way to limit. Put '<-ticker.C' before your work.
   It creates a hard pause between every action.

2. BUFFERED CHANNEL: Use this to allow Bursts. Pre-fill the channel with "tokens".
   The channel size determines the max burst (e.g., size 3 = 3 instant requests).
	`)
}
