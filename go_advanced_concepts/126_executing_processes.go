package main

import (
	"fmt"
	"io/ioutil"
	"os/exec"
)

// TOPIC: Executing Processes
// Explanation: We can spawn external processes from Go using the 'os/exec' package.

func main() {
	// Simple command with no arguments
	dateCmd := exec.Command("date")
	
	// Output returns standard output
	dateOut, err := dateCmd.Output()
	if err != nil {
		panic(err)
	}
	fmt.Println("> date")
	fmt.Println(string(dateOut))

	// Command with arguments
	lsCmd := exec.Command("ls", "-a", "-l", "-h")
	lsOut, err := lsCmd.Output()
	if err != nil {
		panic(err)
	}
	fmt.Println("> ls -alh")
	fmt.Println(string(lsOut))
}
