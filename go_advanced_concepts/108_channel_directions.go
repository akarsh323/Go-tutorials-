package main

import (
	"fmt"
	"time"
)

/*
TOPIC: CHANNEL DIRECTIONS & PRODUCER-CONSUMER PATTERN

CONCEPT:
By default, channels are Bidirectional (you can read & write).
To build robust systems, we restrict channels when passing them to functions.
This enforces a "One-Way Flow" of data, preventing bugs.

SYNTAX:
1. Send-Only (chan<- int): Arrow points INTO the channel.
   - You can ONLY send.
   - Reading causes a compile-time error.

2. Receive-Only (<-chan int): Arrow points OUT OF the channel.
   - You can ONLY receive.
   - Sending causes a compile-time error.

PATTERN: PRODUCER-CONSUMER
- Producer: Generates data -> Sends to channel -> Closes channel.
- Consumer: Ranges over channel -> Processes data -> Stops when closed.
*/

// ---------------------------------------------------------
// 1. The Producer (The Chef)
// ---------------------------------------------------------
// DIRECTION: chan<- int (Send Only)
// The compiler guarantees this function CANNOT read from 'out'.
func producer(out chan<- int) {
	// BEST PRACTICE: The Sender is responsible for closing the channel.
	// We use defer to ensure it closes even if the function crashes.
	defer close(out)

	for i := 1; i <= 5; i++ {
		fmt.Printf("[Chef] Cooking order #%d...\n", i)
		time.Sleep(500 * time.Millisecond) // Simulate work

		// SENDING data into the channel
		out <- i
	}
	fmt.Println("[Chef] All orders cooked. Kitchen closed.")
}

// ---------------------------------------------------------
// 2. The Consumer (The Waiter)
// ---------------------------------------------------------
// DIRECTION: <-chan int (Receive Only)
// The compiler guarantees this function CANNOT send to 'in'.
func consumer(in <-chan int) {
	fmt.Println("[Waiter] Waiting for orders...")

	// RANGE LOOP:
	// Automatically receives values until the channel is closed.
	// It handles the "blocking" logic for us.
	for order := range in {
		fmt.Printf("\t[Waiter] Served order #%d\n", order)
	}

	fmt.Println("[Waiter] Shift over. No more orders.")
}

// ---------------------------------------------------------
// 3. Compile-Time Safety Demo (Commented Out)
// ---------------------------------------------------------
/*
func unsafeAction(in <-chan int, out chan<- int) {
	// ERROR: Cannot send to a receive-only channel
	// in <- 10

	// ERROR: Cannot receive from a send-only channel
	// val := <-out
}
*/

func main() {
	fmt.Println("═══════════════════════════════════════════════════════════")
	fmt.Println("TOPIC: CHANNEL DIRECTIONS (PRODUCER / CONSUMER)")
	fmt.Println("═══════════════════════════════════════════════════════════\n")

	// 1. Create a Bidirectional Channel
	// We MUST create it as bidirectional initially.
	// make(chan<- int) is useless because you could never read from it!
	kitchenPipe := make(chan int)

	// 2. Start the Producer
	// GO MAGIC: Implicit Conversion
	// We pass 'kitchenPipe' (bidirectional).
	// Go converts it to 'chan<- int' (send-only) for the producer function.
	go producer(kitchenPipe)

	// 3. Start the Consumer
	// We pass the SAME 'kitchenPipe'.
	// Go converts it to '<-chan int' (receive-only) for the consumer function.
	consumer(kitchenPipe)

	fmt.Println("\n═══════════════════════════════════════════════════════════")
	fmt.Println("KEY TAKEAWAYS")
	fmt.Println("═══════════════════════════════════════════════════════════")
	fmt.Println(`
1. IMPLICIT CONVERSION:
   You create a two-way channel in 'main', but pass it as a one-way channel.
   This "locks" the door functionality for that specific function.

2. SAFETY:
   If the 'consumer' tries to send data back, the code won't compile.
   This prevents logic bugs where consumers accidentally act like producers.

3. CLOSING:
   Always close in the Producer (Sender).
   If the Consumer closes the channel, the Producer might try to send 
   to a closed channel, causing a PANIC.
	`)
}
