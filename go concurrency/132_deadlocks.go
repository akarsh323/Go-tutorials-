package main

/*
Deadlocks

A deadlock happens when goroutines wait forever
for each other and no progress can be made
*/

func main() {
    ch := make(chan int)
    // ch <- 1 // Uncommenting causes deadlock
}
