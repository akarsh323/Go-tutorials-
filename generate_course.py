import os

# The directory where files will be created
output_dir = "go_advanced_concepts"

# Dictionary containing filename and the source code content with explanations
course_content = {
    "103_goroutines.go": r"""package main

import (
	"fmt"
	"time"
)

// TOPIC: Goroutines
// Explanation: A goroutine is a lightweight thread managed by the Go runtime.
// To start a goroutine, use the keyword 'go' before a function call.

func speak(phrase string) {
	for i := 0; i < 3; i++ {
		fmt.Println(phrase)
		time.Sleep(100 * time.Millisecond)
	}
}

func main() {
	// This runs synchronously (blocking)
	speak("Direct Call: Hello")

	// This runs in a Goroutine (concurrently)
	// The main function continues immediately without waiting for this to finish
	go speak("Goroutine: World")

	// We sleep here to prevent the 'main' function from exiting 
	// before the goroutine has a chance to execute.
	time.Sleep(1 * time.Second)
	fmt.Println("Main function finished")
}
""",

    "104_channels_intro.go": r"""package main

import "fmt"

// TOPIC: Channels - Introduction
// Explanation: Channels are the pipes that connect concurrent goroutines.
// You can send values into channels from one goroutine and receive those 
// values into another goroutine.

func main() {
	// Create a new channel with make(chan val-type)
	messages := make(chan string)

	// Send a value into a channel using the `channel <-` syntax.
	// We do this in a goroutine because sending into an unbuffered channel blocks
	// until there is a receiver.
	go func() {
		messages <- "ping" 
	}()

	// The `<-channel` syntax receives a value from the channel.
	msg := <-messages
	fmt.Println(msg)
}
""",

    "105_unbuffered_channels.go": r"""package main

import (
	"fmt"
	"time"
)

// TOPIC: Unbuffered Channels and Runtime Mechanism
// Explanation: By default, channels are unbuffered. This means that a sender
// will BLOCK until the receiver is ready to receive. This allows for synchronization
// between goroutines without explicit locks.

func main() {
	done := make(chan bool)

	go func() {
		fmt.Print("Working...")
		time.Sleep(time.Second)
		fmt.Println("done")
		
		// Send a signal that work is done
		// This line blocks if no one is listening, but main is listening below.
		done <- true 
	}()

	// The main function blocks here until it receives a value from 'done'
	<-done
	fmt.Println("Main received signal")
}
""",

    "106_buffered_channels.go": r"""package main

import "fmt"

// TOPIC: Buffered Channels
// Explanation: Buffered channels accept a limited number of values without
// a corresponding receiver for those values. Blocking only occurs when
// the buffer is full (sending) or empty (receiving).

func main() {
	// Create a channel that can hold 2 strings
	messages := make(chan string, 2)

	// Because this channel is buffered, we can send these values
	// without a corresponding concurrent receive.
	messages <- "buffered"
	messages <- "channel"

	// If we tried to send a 3rd value here, it would deadlock/block 
	// because the buffer size is 2.

	fmt.Println(<-messages)
	fmt.Println(<-messages)
}
""",

    "107_channel_synchronization.go": r"""package main

import (
	"fmt"
	"time"
)

// TOPIC: Channel Synchronization
// Explanation: We can use channels to synchronize execution across goroutines.
// This example uses a blocking receive to wait for a goroutine to finish.

func worker(done chan bool) {
	fmt.Print("Worker starting...")
	time.Sleep(time.Second)
	fmt.Println("Worker done")

	// Notify that we are done
	done <- true
}

func main() {
	// Start a worker goroutine, giving it the channel to notify on
	done := make(chan bool, 1)
	go worker(done)

	// Block until we receive a notification from the worker on the channel.
	<-done
}
""",

    "108_channel_directions.go": r"""package main

import "fmt"

// TOPIC: Channel Directions
// Explanation: When using channels as function parameters, you can
// specify if a channel is meant to only send or only receive values.
// This increases type-safety.

// ping only accepts a channel for sending values (chan<-)
func ping(pings chan<- string, msg string) {
	pings <- msg
}

// pong accepts one channel for receives (<-chan) and one for sends (chan<-)
func pong(pings <-chan string, pongs chan<- string) {
	msg := <-pings
	pongs <- msg
}

func main() {
	pings := make(chan string, 1)
	pongs := make(chan string, 1)

	ping(pings, "passed message")
	pong(pings, pongs)

	fmt.Println(<-pongs)
}
""",

    "109_multiplexing_select.go": r"""package main

import (
	"fmt"
	"time"
)

// TOPIC: Multiplexing using Select
// Explanation: The 'select' statement lets a goroutine wait on multiple
// communication operations. It acts like a switch statement for channels.

func main() {
	c1 := make(chan string)
	c2 := make(chan string)

	go func() {
		time.Sleep(1 * time.Second)
		c1 <- "one"
	}()
	go func() {
		time.Sleep(2 * time.Second)
		c2 <- "two"
	}()

	// We use a loop to receive both values
	for i := 0; i < 2; i++ {
		// Select will block until ONE of the cases is ready
		select {
		case msg1 := <-c1:
			fmt.Println("received", msg1)
		case msg2 := <-c2:
			fmt.Println("received", msg2)
		}
	}
}
""",

    "110_non_blocking_channel_ops.go": r"""package main

import "fmt"

// TOPIC: Non-blocking Channel Operations
// Explanation: Basic sends and receives on channels are blocking.
// However, we can use 'select' with a 'default' clause to implement
// non-blocking sends, receives, and even non-blocking multi-way selects.

func main() {
	messages := make(chan string)
	signals := make(chan bool)

	// Non-blocking receive
	select {
	case msg := <-messages:
		fmt.Println("received message", msg)
	default:
		fmt.Println("no message received")
	}

	// Non-blocking send
	msg := "hi"
	select {
	case messages <- msg:
		fmt.Println("sent message", msg)
	default:
		fmt.Println("no message sent")
	}

	// Multi-way non-blocking select
	select {
	case msg := <-messages:
		fmt.Println("received message", msg)
	case sig := <-signals:
		fmt.Println("received signal", sig)
	default:
		fmt.Println("no activity")
	}
}
""",

    "111_closing_channels.go": r"""package main

import "fmt"

// TOPIC: Closing Channels
// Explanation: Closing a channel indicates that no more values will be sent on it.
// This is useful to communicate completion to the receivers.

func main() {
	jobs := make(chan int, 5)
	done := make(chan bool)

	go func() {
		// Looping over a channel receives values until the channel is closed
		for j := range jobs {
			fmt.Println("received job", j)
		}
		fmt.Println("received all jobs")
		done <- true
	}()

	for j := 1; j <= 3; j++ {
		jobs <- j
		fmt.Println("sent job", j)
	}
	// Close the channel to signal no more jobs are coming
	close(jobs)
	fmt.Println("sent all jobs")

	<-done
}
""",

    "112_context.go": r"""package main

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
""",

    "113_timers.go": r"""package main

import (
	"fmt"
	"time"
)

// TOPIC: Timers
// Explanation: Timers represent a single event in the future. You tell the timer
// how long you want to wait, and it provides a channel that will be notified
// at that time.

func main() {
	// Timer waits 2 seconds
	timer1 := time.NewTimer(2 * time.Second)

	// Blocks here until timer fires
	<-timer1.C
	fmt.Println("Timer 1 fired")

	timer2 := time.NewTimer(time.Second)
	go func() {
		<-timer2.C
		fmt.Println("Timer 2 fired")
	}()

	// You can cancel a timer before it fires
	stop2 := timer2.Stop()
	if stop2 {
		fmt.Println("Timer 2 stopped")
	}

	time.Sleep(2 * time.Second)
}
""",

    "114_tickers.go": r"""package main

import (
	"fmt"
	"time"
)

// TOPIC: Tickers
// Explanation: Tickers are for when you want to do something repeatedly at regular
// intervals. Like timers, they use a channel to send values.

func main() {
	// Ticker triggers every 500ms
	ticker := time.NewTicker(500 * time.Millisecond)
	done := make(chan bool)

	go func() {
		for {
			select {
			case <-done:
				return
			case t := <-ticker.C:
				fmt.Println("Tick at", t)
			}
		}
	}()

	// Let it tick for 1600ms
	time.Sleep(1600 * time.Millisecond)
	
	// Tickers must be stopped to release resources
	ticker.Stop()
	done <- true
	fmt.Println("Ticker stopped")
}
""",

    "115_worker_pools.go": r"""package main

import (
	"fmt"
	"time"
)

// TOPIC: Worker Pools
// Explanation: Worker pools allows you to run a fixed number of goroutines
// to process jobs from a shared channel. This is a common pattern for 
// limiting concurrency.

// The worker function processes jobs
func worker(id int, jobs <-chan int, results chan<- int) {
	for j := range jobs {
		fmt.Println("worker", id, "started  job", j)
		time.Sleep(time.Second) // Simulate expensive work
		fmt.Println("worker", id, "finished job", j)
		results <- j * 2
	}
}

func main() {
	const numJobs = 5
	jobs := make(chan int, numJobs)
	results := make(chan int, numJobs)

	// Start 3 workers (goroutines)
	// Only 3 jobs can run simultaneously
	for w := 1; w <= 3; w++ {
		go worker(w, jobs, results)
	}

	// Send 5 jobs
	for j := 1; j <= numJobs; j++ {
		jobs <- j
	}
	close(jobs)

	// Collect results
	for a := 1; a <= numJobs; a++ {
		<-results
	}
}
""",

    "116_wait_groups.go": r"""package main

import (
	"fmt"
	"sync"
	"time"
)

// TOPIC: WaitGroups
// Explanation: To wait for multiple goroutines to finish, we use a WaitGroup.
// It acts like a counter: Add() increments it, Done() decrements it, and Wait()
// blocks until the counter is zero.

func worker(id int, wg *sync.WaitGroup) {
	// On return, notify the WaitGroup that we're done
	defer wg.Done()

	fmt.Printf("Worker %d starting\n", id)
	time.Sleep(time.Second)
	fmt.Printf("Worker %d done\n", id)
}

func main() {
	var wg sync.WaitGroup

	for i := 1; i <= 5; i++ {
		// Increment the WaitGroup counter
		wg.Add(1)
		
		// Pass the pointer to the WaitGroup to the goroutine
		go worker(i, &wg)
	}

	// Block until the WaitGroup counter goes back to 0
	wg.Wait()
	fmt.Println("All workers done")
}
""",

    "117_mutexes.go": r"""package main

import (
	"fmt"
	"sync"
)

// TOPIC: Mutexes
// Explanation: For more complex state management where multiple goroutines
// access specific data, we use a Mutex (mutual exclusion) to safely access data.
// It ensures only one goroutine can access a variable at a time.

type Container struct {
	mu       sync.Mutex
	counters map[string]int
}

// Helper to increment safely
func (c *Container) inc(name string) {
	// Lock the mutex before accessing counters
	c.mu.Lock()
	// Unlock when the function exits
	defer c.mu.Unlock()
	
	c.counters[name]++
}

func main() {
	c := Container{
		counters: map[string]int{"a": 0, "b": 0},
	}

	var wg sync.WaitGroup

	// Helper function to spin up goroutines
	doIncrement := func(name string, n int) {
		for i := 0; i < n; i++ {
			c.inc(name)
		}
		wg.Done()
	}

	wg.Add(3)
	go doIncrement("a", 10000)
	go doIncrement("a", 10000)
	go doIncrement("b", 10000)

	wg.Wait()
	fmt.Println(c.counters)
}
""",

    "118_atomic_counters.go": r"""package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

// TOPIC: Atomic Counters
// Explanation: For simple counters, Mutexes can be overkill and slow.
// The sync/atomic package provides low-level atomic memory primitives.

func main() {
	var ops uint64
	var wg sync.WaitGroup

	// Start 50 goroutines that each increment the counter 1000 times
	for i := 0; i < 50; i++ {
		wg.Add(1)
		go func() {
			for c := 0; c < 1000; c++ {
				// Atomic Add ensures no race conditions without using a heavy lock
				atomic.AddUint64(&ops, 1)
			}
			wg.Done()
		}()
	}

	wg.Wait()
	fmt.Println("Ops:", ops) // Should be 50,000
}
""",

    "119_rate_limiting_basics.go": r"""package main

import (
	"fmt"
	"time"
)

// TOPIC: Rate Limiting (Basics)
// Explanation: Rate limiting controls the frequency of events.
// Go supports this elegantly using Tickers and Channels.

func main() {
	// Create a stream of requests
	requests := make(chan int, 5)
	for i := 1; i <= 5; i++ {
		requests <- i
	}
	close(requests)

	// The limiter channel will receive a value every 200 milliseconds.
	limiter := time.Tick(200 * time.Millisecond)

	for req := range requests {
		// By blocking on the limiter channel, we limit the loop execution speed
		<-limiter
		fmt.Println("request", req, time.Now())
	}
}
""",

    "120_rate_limiting_algorithms.go": r"""package main

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
""",

    "123_stateful_goroutines.go": r"""package main

import (
	"fmt"
	"math/rand"
	"sync/atomic"
	"time"
)

// TOPIC: Stateful Goroutines
// Explanation: Instead of using Mutexes to protect shared data, we can use
// Go's philosophy: "Do not communicate by sharing memory; share memory by communicating."
// A single goroutine owns the state, and others communicate with it via channels.

type readOp struct {
	key  int
	resp chan int
}
type writeOp struct {
	key  int
	val  int
	resp chan bool
}

func main() {
	var readOps uint64
	var writeOps uint64

	reads := make(chan readOp)
	writes := make(chan writeOp)

	// This is the Stateful Goroutine. It owns the 'state' map.
	go func() {
		var state = make(map[int]int)
		for {
			select {
			case read := <-reads:
				read.resp <- state[read.key]
			case write := <-writes:
				state[write.key] = write.val
				write.resp <- true
			}
		}
	}()

	// Start 100 goroutines requesting reads
	for r := 0; r < 100; r++ {
		go func() {
			for {
				read := readOp{
					key:  rand.Intn(5),
					resp: make(chan int)}
				reads <- read
				<-read.resp
				atomic.AddUint64(&readOps, 1)
				time.Sleep(time.Millisecond)
			}
		}()
	}

	// Start 10 goroutines requesting writes
	for w := 0; w < 10; w++ {
		go func() {
			for {
				write := writeOp{
					key:  rand.Intn(5),
					val:  rand.Intn(100),
					resp: make(chan bool)}
				writes <- write
				<-write.resp
				atomic.AddUint64(&writeOps, 1)
				time.Sleep(time.Millisecond)
			}
		}()
	}

	time.Sleep(time.Second)
	fmt.Println("readOps:", atomic.LoadUint64(&readOps))
	fmt.Println("writeOps:", atomic.LoadUint64(&writeOps))
}
""",

    "124_sorting.go": r"""package main

import (
	"fmt"
	"sort"
)

// TOPIC: Sorting
// Explanation: Go's 'sort' package implements sorting for built-ins and user-defined types.

// ByFunction implements sort.Interface for []string based on the length of the string
type ByLength []string

func (s ByLength) Len() int {
	return len(s)
}
func (s ByLength) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
func (s ByLength) Less(i, j int) bool {
	return len(s[i]) < len(s[j])
}

func main() {
	// Sorting specific types
	ints := []int{7, 2, 4}
	sort.Ints(ints)
	fmt.Println("Ints:   ", ints)

	// Checking if sorted
	s := sort.IntsAreSorted(ints)
	fmt.Println("Sorted: ", s)

	// Custom sorting
	fruits := []string{"peach", "banana", "kiwi"}
	sort.Sort(ByLength(fruits))
	fmt.Println("By Len: ", fruits)
}
""",

    "125_testing_test.go": r"""package main

import (
	"fmt"
	"testing"
)

// TOPIC: Testing / Benchmarking
// Explanation: Go has a built-in 'testing' package.
// Files must end in '_test.go'. Functions must start with 'Test'.
// RUN THIS USING: go test -v

// Simple function to test
func IntMin(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// The Test Function
func TestIntMinBasic(t *testing.T) {
	ans := IntMin(2, -2)
	if ans != -2 {
		// t.Errorf reports a failure but continues execution
		t.Errorf("IntMin(2, -2) = %d; want -2", ans)
	}
}

// Table-driven tests (Common Go pattern)
func TestIntMinTableDriven(t *testing.T) {
	var tests = []struct {
		a, b int
		want int
	}{
		{0, 1, 0},
		{1, 0, 0},
		{2, -2, -2},
		{0, -1, -1},
		{-1, 0, -1},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("%d,%d", tt.a, tt.b)
		t.Run(testname, func(t *testing.T) {
			ans := IntMin(tt.a, tt.b)
			if ans != tt.want {
				t.Errorf("got %d, want %d", ans, tt.want)
			}
		})
	}
}

// Benchmarking Function
// Run using: go test -bench=.
func BenchmarkIntMin(b *testing.B) {
	for i := 0; i < b.N; i++ {
		IntMin(1, 2)
	}
}
""",

    "126_executing_processes.go": r"""package main

import (
	"fmt"
	"io/ioutil"
	"os/exec"
)

// TOPIC: Executing Processes
// Explanation: We can spawn external processes from Go using the 'os/exec' package.

func main() {
	// Simple command with no arguments
	dateCmd := exec.Command("date")
	
	// Output returns standard output
	dateOut, err := dateCmd.Output()
	if err != nil {
		panic(err)
	}
	fmt.Println("> date")
	fmt.Println(string(dateOut))

	// Command with arguments
	lsCmd := exec.Command("ls", "-a", "-l", "-h")
	lsOut, err := lsCmd.Output()
	if err != nil {
		panic(err)
	}
	fmt.Println("> ls -alh")
	fmt.Println(string(lsOut))
}
""",

    "127_signals.go": r"""package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

// TOPIC: Signals
// Explanation: Signals allow us to handle external interrupts (like CTRL+C).
// This is essential for Graceful Shutdowns of servers.

func main() {
	// Create a channel to receive signals
	sigs := make(chan os.Signal, 1)
	done := make(chan bool, 1)

	// Register the channel to receive SIGINT (Ctrl+C) and SIGTERM
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	// This goroutine executes a blocking receive for signals
	go func() {
		sig := <-sigs
		fmt.Println()
		fmt.Println(sig)
		done <- true
	}()

	fmt.Println("awaiting signal (Press CTRL+C)")
	// The program waits here until it gets the signal
	<-done
	fmt.Println("exiting")
}
""",

    "128_reflect.go": r"""package main

import (
	"fmt"
	"reflect"
)

// TOPIC: Reflect
// Explanation: Reflection allows a program to examine its own structure, 
// strictly checking types at runtime. It's powerful but should be used sparingly
// as it is slower and complex.

func main() {
	// Basic Reflection
	var x float64 = 3.4
	
	fmt.Println("type:", reflect.TypeOf(x))
	fmt.Println("value:", reflect.ValueOf(x))

	// Modifying values via reflection
	// We must pass the pointer to allow modification
	p := reflect.ValueOf(&x) 
	
	// Elem() gets the value the pointer points to
	v := p.Elem() 
	
	if v.CanSet() {
		v.SetFloat(7.1)
	}
	fmt.Println("new value of x:", x)
}
"""
}

# Create directory
if not os.path.exists(output_dir):
    os.makedirs(output_dir)

# Write files
for filename, content in course_content.items():
    file_path = os.path.join(output_dir, filename)
    with open(file_path, "w") as f:
        f.write(content)
    print(f"Created {filename}")

print("\nSuccess! All files created in the 'go_advanced_concepts' directory.")