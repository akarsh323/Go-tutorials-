package main

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
