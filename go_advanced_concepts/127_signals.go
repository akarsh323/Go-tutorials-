package main

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
