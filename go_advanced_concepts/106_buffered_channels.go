package main

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
