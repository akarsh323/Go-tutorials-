package main

// Constants — show untyped vs typed constants and iota.

import "fmt"

const Pi = 3.14159 // untyped constant

const (
	// iota example: successive integer constants
	_  = iota
	KB = 1 << (10 * iota) // 1 << 10
	MB
	GB
)

func main() {
	const greeting = "hello"
	fmt.Println(greeting, "Pi=", Pi)
	fmt.Println("KB, MB, GB:", KB, MB, GB)
}
