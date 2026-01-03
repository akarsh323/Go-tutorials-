package main

// Exit — demonstrates how to exit with a status code.
// Note: calling os.Exit will terminate the program immediately and
// deferred functions will not run, so leave commented for interactive
// examples.

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("program would exit with code 1 (call os.Exit(1))")
	// Uncomment to use in real CLI programs:
	// os.Exit(1)
	// If you need cleanup, return from main or call cleanup before os.Exit.
	_ = os.Stderr
}
