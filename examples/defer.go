package main

// Defer — demonstrates how deferred calls execute in LIFO order and
// are useful for cleanup (closing files, unlocking mutexes, etc.).

import "fmt"

func main() {
	fmt.Println("start")
	defer fmt.Println("defer 1")
	defer fmt.Println("defer 2")
	fmt.Println("end of main body")
}
