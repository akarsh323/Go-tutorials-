package main

import (
	"fmt"
	"io"
	"os/exec"
	"strings"
	"time"
)

/*
TOPIC: PROCESS SPAWNING (OS/EXEC)

CONCEPT:
Go can launch external Operating System processes (like Python scripts,
FFmpeg, or CLI tools) using the 'os/exec' package.

KEY METHODS:
1. cmd.Output(): Runs the command and waits for it to finish. Returns STDOUT.
2. cmd.Start(): Starts the command asynchronously (Non-blocking).
3. cmd.Wait(): Waits for a started command to finish (Releases resources).
4. cmd.Process.Kill(): Forcefully stops a running process.
*/

// ---------------------------------------------------------
// Example 1: The "Hello World" of Exec
// ---------------------------------------------------------
// Runs a simple command and captures the output.
func example1_BasicExec() {
	fmt.Println("--- Example 1: Basic Execution (.Output) ---")

	// 1. Prepare the command
	// We are asking the OS to run: echo "Hello World"
	cmd := exec.Command("echo", "Hello from External Process!")

	// 2. Execute
	// .Output() blocks until the command finishes.
	outputBytes, err := cmd.Output()
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		return
	}

	// 3. Result
	fmt.Printf("Result: %s", string(outputBytes))
	fmt.Println("----------------------------------------\n")
}

// ---------------------------------------------------------
// Example 2: Feeding Input (Stdin Pipe)
// ---------------------------------------------------------
// Piping data INTO a command (simulating: echo "data" | grep "foo")
func example2_StdinPipe() {
	fmt.Println("--- Example 2: Feeding Stdin (Grep) ---")

	// We want to run: grep "foo"
	cmd := exec.Command("grep", "foo")

	// 1. Prepare Input Data
	inputData := "foo bar\nbar bas\nfood is good\nhello world\n"
	reader := strings.NewReader(inputData)

	// 2. Attach Reader to Stdin
	cmd.Stdin = reader

	// 3. Execute
	// Note: If grep finds nothing, it returns exit code 1 (which Go sees as an error)
	output, err := cmd.Output()
	if err != nil {
		fmt.Println("Grep finished (might be no matches or error):", err)
	}

	fmt.Printf("Grep Results:\n%s", string(output))
	fmt.Println("----------------------------------------\n")
}

// ---------------------------------------------------------
// Example 3: Lifecycle (Start, Wait, Kill)
// ---------------------------------------------------------
// Managing a long-running background process.
func example3_Lifecycle() {
	fmt.Println("--- Example 3: Lifecycle (Start -> Kill -> Wait) ---")

	// A command that sleeps for 10 seconds
	cmd := exec.Command("sleep", "10")

	// 1. START (Non-blocking)
	if err := cmd.Start(); err != nil {
		fmt.Println("Failed to start:", err)
		return
	}
	fmt.Printf("Process started. PID: %d\n", cmd.Process.Pid)
	fmt.Println("Main: I am doing other work while process runs...")

	// 2. TIMEOUT LOGIC
	time.Sleep(2 * time.Second)
	fmt.Println("Main: Process taking too long. Killing it!")

	// 3. KILL
	if err := cmd.Process.Kill(); err != nil {
		fmt.Println("Failed to kill:", err)
	} else {
		fmt.Println("Main: Sent Kill signal.")
	}

	// 4. WAIT (Cleanup)
	// Essential to prevent "Zombie Processes"
	err := cmd.Wait()
	fmt.Printf("Final Status: %v\n", err) // Expected: "signal: killed"
	fmt.Println("----------------------------------------\n")
}

// ---------------------------------------------------------
// Example 4: Advanced Piping (io.Pipe)
// ---------------------------------------------------------
// Streaming dynamic data from a Goroutine into an external process.
func example4_StreamingPipe() {
	fmt.Println("--- Example 4: Streaming Data (io.Pipe) ---")

	// Create a pipe:
	// pr (PipeReader) -> Read by Grep
	// pw (PipeWriter) -> Written to by Go routine
	pr, pw := io.Pipe()

	cmd := exec.Command("grep", "priority")
	cmd.Stdin = pr

	// 1. Writer Routine
	go func() {
		defer pw.Close() // Close triggers EOF for grep

		// Simulate dynamic data generation
		lines := []string{
			"log: task started",
			"log: priority task running",
			"log: normal task running",
			"log: priority error occurred",
		}

		for _, line := range lines {
			fmt.Fprintln(pw, line)
			time.Sleep(100 * time.Millisecond) // Simulate delay
		}
	}()

	// 2. Execute
	output, _ := cmd.Output()
	fmt.Printf("Filtered Stream:\n%s", string(output))
	fmt.Println("----------------------------------------\n")
}

func main() {
	fmt.Println("═══════════════════════════════════════════════════════════")
	fmt.Println("TOPIC: PROCESS SPAWNING (OS/EXEC)")
	fmt.Println("═══════════════════════════════════════════════════════════\n")

	// Check if we are on a supported OS (rudimentary check)
	// These examples rely on Unix tools (echo, grep, sleep)
	_, err := exec.LookPath("grep")
	if err != nil {
		fmt.Println("WARNING: 'grep' command not found.")
		fmt.Println("These examples require a Unix-like environment (Linux, Mac, WSL).")
		fmt.Println("If you are on Windows CMD/PowerShell, some commands may fail.")
		fmt.Println("═══════════════════════════════════════════════════════════\n")
	}

	example1_BasicExec()
	example2_StdinPipe()
	example3_Lifecycle()
	example4_StreamingPipe()

	fmt.Println("═══════════════════════════════════════════════════════════")
	fmt.Println("KEY TAKEAWAYS")
	fmt.Println("═══════════════════════════════════════════════════════════")
	fmt.Println(`
1. OUTPUT: Use cmd.Output() for simple run-and-wait commands.
2. INPUT: Use cmd.Stdin to pipe data (io.Reader) into the process.
3. CONTROL: Use cmd.Start() and cmd.Wait() for background tasks.
4. CLEANUP: Always call cmd.Wait() if you used cmd.Start(), otherwise the OS 
   cannot release the process resources (Zombie processes).
	`)
}
