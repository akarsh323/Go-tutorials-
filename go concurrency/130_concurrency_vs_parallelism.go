package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

/*
TOPIC: CONCURRENCY VS PARALLELISM

CONCEPT:
1. CONCURRENCY: Structure. Breaking a program into independently executing pieces.
   Analogy: One barista juggling 3 customers (Context Switching).

2. PARALLELISM: Execution. Running those pieces at the exact same time.
   Analogy: 3 baristas serving 3 customers simultaneously.

   NOTE: You can have Concurrency without Parallelism (on a single core CPU).
*/

// ---------------------------------------------------------
// Example 1: Concurrency (The Illusion of Simultaneity)
// ---------------------------------------------------------
// Here, we rely on the Go Scheduler to switch between tasks rapidly.
// Even on a single core, this looks like they run together.

func printNumbers(wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 1; i <= 5; i++ {
		fmt.Printf("%d ", i)
		// Sleep allows the scheduler to say "Pause this, run the other one"
		time.Sleep(200 * time.Millisecond)
	}
}

func printLetters(wg *sync.WaitGroup) {
	defer wg.Done()
	letters := []rune{'a', 'b', 'c', 'd', 'e'}
	for _, l := range letters {
		fmt.Printf("%c ", l)
		time.Sleep(200 * time.Millisecond)
	}
}

func example1_Concurrency() {
	fmt.Println("--- Example 1: Concurrency (Interleaving) ---")

	// Optional: Force 1 CPU to prove concurrency works without parallelism
	runtime.GOMAXPROCS(1)

	var wg sync.WaitGroup
	wg.Add(2)

	go printNumbers(&wg)
	go printLetters(&wg)

	wg.Wait()
	fmt.Println("\n----------------------------------------\n")
}

// ---------------------------------------------------------
// Example 2: Parallelism (True Simultaneity)
// ---------------------------------------------------------
// Here, we unlock multiple CPU cores. We run a "Heavy" CPU task.
// If run on multi-core, these will finish at almost the exact same time.

func heavyTask(id int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("[Task %d] Starting crunching numbers...\n", id)

	// Simulate heavy CPU work (No Sleep!)
	// This forces the CPU core to stay busy.
	for i := 0; i < 100_000_000; i++ {
		// complex calculation...
	}

	// Log the finish time with high precision
	fmt.Printf("[Task %d] FINISHED at %s\n", id, time.Now().Format("15:04:05.000000"))
}

func example2_Parallelism() {
	fmt.Println("--- Example 2: Parallelism (Multi-Core) ---")

	// 1. UNLOCK HARDWARE
	// Tell Go to use 4 logical CPUs (or all available)
	numCPU := runtime.NumCPU()
	runtime.GOMAXPROCS(numCPU)
	fmt.Printf("System has %d CPUs. Using all of them.\n", numCPU)

	var wg sync.WaitGroup
	numTasks := 4 // Launch 4 tasks to utilize the cores

	// 2. Launch Tasks
	for i := 1; i <= numTasks; i++ {
		wg.Add(1)
		go heavyTask(i, &wg)
	}

	wg.Wait()
	fmt.Println("----------------------------------------\n")
}

func main() {
	fmt.Println("═══════════════════════════════════════════════════════════")
	fmt.Println("TOPIC: CONCURRENCY VS PARALLELISM")
	fmt.Println("═══════════════════════════════════════════════════════════\n")

	// 1. Run the Concurrent Example
	// Watch the output interleave (1 a 2 b...)
	example1_Concurrency()

	// 2. Run the Parallel Example
	// Watch the timestamps. They should be nearly identical.
	example2_Parallelism()

	fmt.Println("═══════════════════════════════════════════════════════════")
	fmt.Println("KEY TAKEAWAYS")
	fmt.Println("═══════════════════════════════════════════════════════════")
	fmt.Println(`
1. CONCURRENCY is about DESIGN. We break the program into 'printNumbers' and 
   'printLetters' so they can be managed independently.

2. PARALLELISM is about EXECUTION. By using 'runtime.GOMAXPROCS', we utilize
   physical hardware to run the code at the exact same instant.

3. THE TRAP: 'time.Sleep' yields the processor (Concurrency friendly). 
   Busy loops (like example 2) hog the processor.
	`)
}
