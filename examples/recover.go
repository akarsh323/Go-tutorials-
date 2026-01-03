package main

// Recover — catches panics using recover() in a deferred function.
// Recover only works inside defer; once called, panic stops propagating.

import "fmt"

func safe() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("recovered:", r)
		}
	}()
	fmt.Println("about to panic")
	panic("error in safe()")
	fmt.Println("this never runs")
}

func riskyWithoutRecover() {
	panic("unrecovered panic")
}

func main() {
	fmt.Println("starting")

	// Example 1: with recover
	safe()
	fmt.Println("continues after safe() recover")

	// Example 2: without recover (would crash)
	// Uncomment to see difference:
	// riskyWithoutRecover()
	// fmt.Println("this never runs")
}
