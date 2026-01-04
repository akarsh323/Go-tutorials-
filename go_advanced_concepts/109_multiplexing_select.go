package main

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
