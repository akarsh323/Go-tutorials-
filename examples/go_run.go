package main

// Go Run — quick way to run Go programs without building.
// Teaches: `go run` for single files and multiple files, and when to use it.

import "fmt"

func main() {
	fmt.Println("Run with: go run go_run.go")
	fmt.Println("You can also run multiple files: go run file1.go file2.go")
	fmt.Println("Note: 'go run' compiles to a temporary binary and executes it.")
}
