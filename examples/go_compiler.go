package main

// Go Compiler — how to build artifacts and common flags.
// Teaches: `go build`, naming, and quick tips for compilation.

import "fmt"

func main() {
	fmt.Println("Run 'go build' to compile this package into an executable.")
	fmt.Println("Example: go build -o gotut_bin go_compiler.go && ./gotut_bin")
	fmt.Println("Tip: For multi-file packages run 'go build' from the package directory.")
}
