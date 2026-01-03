package main

// init function — demonstration of package initialization order.
// init runs before main and can be used for setup. Avoid heavy work in init.

import "fmt"

func init() {
	fmt.Println("init: runs before main (package initialization)")
}

func main() {
	fmt.Println("main: runs after init")
}
