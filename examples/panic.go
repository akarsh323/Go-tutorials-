package main

// Panic — the panic builtin causes runtime errors. Deferred functions
// still run before panic terminates. Use recover() to catch panics.

import "fmt"

func mayPanic() {
	fmt.Println("about to panic")
	panic("something went wrong")
	fmt.Println("this will never run")
}

func withRecover() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("recovered from:", r)
		}
	}()
	mayPanic()
}

func main() {
	fmt.Println("starting program")

	// Example 1: panic without recover (commented to keep running)
	// deferred after panic will still run
	// defer fmt.Println("deferred runs even during panic")
	// mayPanic()

	// Example 2: panic with recover
	withRecover()
	fmt.Println("program continues after recover")
}
