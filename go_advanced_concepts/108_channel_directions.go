package main

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
