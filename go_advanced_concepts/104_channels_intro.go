package main

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
