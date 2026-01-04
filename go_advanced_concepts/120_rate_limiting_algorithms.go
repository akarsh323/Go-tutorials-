package main

import (
	"fmt"
	"time"
)

// TOPIC: Rate Limiting Algorithms (Conceptual)
// Covers: Token Bucket, Fixed Window, Leaky Bucket
//
// Explanation:
// 1. Token Bucket: Tokens are added to a bucket at a fixed rate. Requests consume tokens.
//    If the bucket is empty, the request waits or is dropped. Good for bursts.
// 2. Leaky Bucket: Requests enter a queue (bucket) and are processed at a constant fixed rate.
//    Good for smoothing out traffic (no bursts).

func main() {
	// --- TOKEN BUCKET EXAMPLE ---
	fmt.Println("--- Token Bucket (Bursty) ---")

	// 1. Create a channel to hold tokens (buffer size 3 = burst capability)
	burstyLimiter := make(chan time.Time, 3)

	// 2. Fill the bucket initially
	for i := 0; i < 3; i++ {
		burstyLimiter <- time.Now()
	}

	// 3. Add a new token every 200ms
	go func() {
		for t := range time.Tick(200 * time.Millisecond) {
			// Non-blocking send (if bucket is full, drop the token)
			select {
			case burstyLimiter <- t:
			default:
			}
		}
	}()

	// 4. Simulate 5 incoming requests
	burstyRequests := make(chan int, 5)
	for i := 1; i <= 5; i++ {
		burstyRequests <- i
	}
	close(burstyRequests)

	// Process requests
	for req := range burstyRequests {
		<-burstyLimiter // Consume token
		fmt.Println("request", req, time.Now())
	}
}
package main

import (
	"fmt"
	"time"
)

// --- 1. The Rate Limiter Definition ---

// RateLimiter holds our bucket (channel) and timing logic.
type RateLimiter struct {
	tokens     chan struct{} // The bucket of tokens (0 bytes per token for efficiency)
	refillTime time.Duration // How often we add a new token
}

// NewRateLimiter creates the bucket and fills it up initially.
// rateLimit: The maximum number of tokens the bucket can hold (buffer size).
// refillTime: The duration to wait before adding a new token.
func NewRateLimiter(rateLimit int, refillTime time.Duration) *RateLimiter {
	// Create the instance
	rl := &RateLimiter{
		// Create a buffered channel. The buffer size IS the rate limit.
		tokens:     make(chan struct{}, rateLimit),
		refillTime: refillTime,
	}

	// Initial Fill: The bucket starts full.
	// We allow 'bursts' immediately upon creation.
	for i := 0; i < rateLimit; i++ {
		rl.tokens <- struct{}{}
	}

	// Start the background refill worker (The "Producer")
	go rl.startRefill()

	return rl
}

// --- 2. The Refill Mechanism (The Producer) ---

// startRefill runs in the background forever, acting like the arcade owner.
func (rl *RateLimiter) startRefill() {
	// Create a ticker that ticks at our refill interval
	ticker := time.NewTicker(rl.refillTime)
	defer ticker.Stop()

	for {
		select {
		// Wait for the ticker to "tick"
		case <-ticker.C:
			// Try to add a token to the bucket
			select {
			case rl.tokens <- struct{}{}:
				// Success! Token added.
			default:
				// The bucket is full (channel buffer is full).
				// We drop this token and do nothing.
			}
		}
	}
}

// --- 3. The Allow Method (The Consumer) ---

// Allow attempts to take a token from the bucket.
// Returns true if a token was taken (request allowed), false otherwise.
func (rl *RateLimiter) Allow() bool {
	select {
	case <-rl.tokens:
		// We successfully took a token from the bucket
		return true
	default:
		// The bucket is empty! Request denied.
		return false
	}
}

// --- 4. The Main Simulation ---

func main() {
	// Setup: Max 5 tokens, refill 1 token every 1 second
	// This allows a burst of 5 immediately, then throttles to 1 request/sec.
	limiter := NewRateLimiter(5, 1*time.Second)

	fmt.Println("--- Starting Token Bucket Simulation ---")
	fmt.Println("Limit: 5 Burst, Refill: 1 per sec")

	// Simulate 10 incoming requests
	for i := 1; i <= 10; i++ {
		// Ask the limiter for permission
		if limiter.Allow() {
			fmt.Printf("Request %d : Allowed ✅\n", i)
		} else {
			fmt.Printf("Request %d : Denied ❌\n", i)
		}

		// Wait 200ms before the next request
		// Note: Because 5 requests take 1000ms (5 * 200ms), the 6th request
		// often sneaks in right as the refill ticker fires at the 1-second mark!
		time.Sleep(200 * time.Millisecond)
	}
}