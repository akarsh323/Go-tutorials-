package main

import (
	"fmt"
	"time"
)

/*
TOPIC: WORKER POOLS

CONCEPT:
A Worker Pool is a pattern to manage concurrency resources.
Instead of spawning 1,000 Goroutines for 1,000 tasks (which can crash a system),
we hire a fixed number of "Workers" (e.g., 3).

ARCHITECTURE:
1. JOBS CHANNEL: A buffered channel holding tasks waiting to be processed.
2. WORKERS: Fixed number of Goroutines that loop over the Jobs channel.
3. RESULTS CHANNEL: A channel where workers send their finished output.

BENEFITS:
- Controls resource usage (CPU/Memory).
- Distributes work evenly (Fast workers take more jobs).
*/

// ---------------------------------------------------------
// Example 1: The Basic "Integer Doubler" Pool
// ---------------------------------------------------------

// Step 1: The Worker Function
// id: To identify the worker.
// jobs: Receive-Only (<-chan) to read tasks.
// results: Send-Only (chan<-) to send answers.
func intWorker(id int, jobs <-chan int, results chan<- int) {
	// The range loop waits for jobs.
	// It stops automatically when 'jobs' is CLOSED and EMPTY.
	for j := range jobs {
		fmt.Printf("[Worker %d] Started job %d\n", id, j)

		// Simulate expensive processing (e.g., calculation)
		time.Sleep(1 * time.Second)

		fmt.Printf("[Worker %d] Finished job %d\n", id, j)

		// Send result
		results <- j * 2
	}
}

func example1_IntegerPool() {
	fmt.Println("--- Example 1: Integer Doubler Pool ---")

	const numJobs = 5
	const numWorkers = 3

	// 1. Create Buffered Channels
	jobs := make(chan int, numJobs)
	results := make(chan int, numJobs)

	// 2. Spawn Workers
	fmt.Printf("[Main] Spawning %d workers...\n", numWorkers)
	for w := 1; w <= numWorkers; w++ {
		go intWorker(w, jobs, results)
	}

	// 3. Send Jobs
	fmt.Printf("[Main] Sending %d jobs...\n", numJobs)
	for j := 1; j <= numJobs; j++ {
		jobs <- j
	}
	// CRITICAL: Close the jobs channel to signal "No more work".
	// If we forget this, workers will wait forever (Deadlock).
	close(jobs)

	// 4. Collect Results
	// We know how many jobs we sent, so we loop exactly that many times.
	for a := 1; a <= numJobs; a++ {
		res := <-results
		fmt.Printf("[Main] Collected Result: %d\n", res)
	}
	fmt.Println("----------------------------------------\n")
}

// ---------------------------------------------------------
// Example 2: Real-World "Ticket Processing" System
// ---------------------------------------------------------

// The Data Structure
type TicketRequest struct {
	PersonID   int
	NumTickets int
	Cost       int
}

// The Worker
func ticketProcessor(id int, requests <-chan TicketRequest, confirmations chan<- string) {
	for req := range requests {
		// Log the work
		fmt.Printf("[Worker %d] Processing Order: Person %d wants %d tickets ($%d)\n",
			id, req.PersonID, req.NumTickets, req.Cost)

		// Simulate API call / Database transaction
		time.Sleep(500 * time.Millisecond)

		// Send Confirmation Message
		msg := fmt.Sprintf("Order confirmed for Person %d", req.PersonID)
		confirmations <- msg
	}
}

func example2_TicketSystem() {
	fmt.Println("--- Example 2: Ticket Processing System ---")

	const totalOrders = 6
	const workerCount = 2 // Fewer workers than orders
	const pricePerTicket = 10

	// Channels
	requestChan := make(chan TicketRequest, totalOrders)
	confirmChan := make(chan string, totalOrders)

	// 1. Start Workers
	fmt.Println("[System] Opening Ticket Counters (Workers)...")
	for w := 1; w <= workerCount; w++ {
		go ticketProcessor(w, requestChan, confirmChan)
	}

	// 2. Generate Orders
	fmt.Println("[System] Queuing Orders...")
	for i := 1; i <= totalOrders; i++ {
		req := TicketRequest{
			PersonID:   i,
			NumTickets: i * 2,
			Cost:       (i * 2) * pricePerTicket,
		}
		requestChan <- req
	}
	close(requestChan) // Queue is full, close input.

	// 3. Wait for Confirmations
	for i := 1; i <= totalOrders; i++ {
		confirmation := <-confirmChan
		fmt.Println("[System] " + confirmation)
	}
	fmt.Println("----------------------------------------\n")
}

func main() {
	fmt.Println("═══════════════════════════════════════════════════════════")
	fmt.Println("TOPIC: WORKER POOLS")
	fmt.Println("═══════════════════════════════════════════════════════════\n")

	// Run Basic Example
	example1_IntegerPool()

	// Run Real-World Example
	example2_TicketSystem()

	fmt.Println("═══════════════════════════════════════════════════════════")
	fmt.Println("KEY TAKEAWAYS")
	fmt.Println("═══════════════════════════════════════════════════════════")
	fmt.Println(`
1. RESOURCE CONTROL: We processed 6 orders with only 2 workers. This prevents
   overwhelming the system if 1 million orders came in.

2. QUEUEING: The 'jobs' channel acts as a queue. Workers pick tasks as soon 
   as they are free.

3. CLOSING: You MUST close the 'jobs' channel when you are done sending.
   This tells the workers to stop their 'range' loops and shut down gracefully.
	`)
}
