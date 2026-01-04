package main

import (
	"math/rand"
	"testing"
	"time"
)

/*
TOPIC: TESTING, BENCHMARKING & PROFILING

INSTRUCTIONS:
1. Save this file as 'main_test.go'.
2. Open your terminal in the same folder.
3. Run Unit Tests:      go test -v
4. Run Benchmarks:      go test -bench=.
5. Run Memory Profile:  go test -bench=. -memprofile=mem.pprof

CONCEPT:
Go has a built-in 'testing' package.
- Tests must start with 'TestXxx' and take '(t *testing.T)'.
- Benchmarks must start with 'BenchmarkXxx' and take '(b *testing.B)'.
*/

// ---------------------------------------------------------
// PART 1: THE CODE TO BE TESTED
// ---------------------------------------------------------
// Usually, this would be in 'main.go', but we put it here
// so you can see everything in one file.

// Add is a simple function we want to test.
func Add(a, b int) int {
	return a + b
}

// SumSlice sums a list of integers.
// We will use this for Benchmarking.
func SumSlice(data []int) int {
	sum := 0
	for _, n := range data {
		sum += n
	}
	return sum
}

// GenerateRandomSlice simulates an "Expensive Setup".
func GenerateRandomSlice(size int) []int {
	rand.Seed(time.Now().UnixNano())
	data := make([]int, size)
	for i := 0; i < size; i++ {
		data[i] = rand.Intn(100)
	}
	return data
}

// ---------------------------------------------------------
// PART 2: UNIT TESTS
// ---------------------------------------------------------

// 1. THE BASICS
// A simple test to verify 2 + 3 = 5.
func TestAddBasic(t *testing.T) {
	result := Add(2, 3)
	expected := 5

	if result != expected {
		// t.Errorf reports a failure but continues the test
		t.Errorf("Add(2, 3) = %d; want %d", result, expected)
	}
}

// 2. TABLE-DRIVEN TESTS (The Pro Move)
// Instead of writing 10 if-statements, we use a slice of structs.
func TestAddTableDriven(t *testing.T) {
	// Define the table
	tests := []struct {
		name     string // Name of the test case
		a, b     int    // Inputs
		expected int    // Expected Output
	}{
		{"Positive Numbers", 2, 3, 5},
		{"Zeros", 0, 0, 0},
		{"Negative Numbers", -1, 1, 0},
		{"Large Numbers", 100, 200, 300},
	}

	// Loop over the table
	for _, tc := range tests {
		// 3. SUBTESTS (t.Run)
		// Creates a named sub-test in the logs.
		// Helpful for debugging specific failed cases.
		t.Run(tc.name, func(t *testing.T) {
			result := Add(tc.a, tc.b)
			if result != tc.expected {
				t.Errorf("Add(%d, %d) = %d; want %d", tc.a, tc.b, result, tc.expected)
			}
		})
	}
}

// ---------------------------------------------------------
// PART 3: BENCHMARKING
// ---------------------------------------------------------
// Benchmarks test PERFORMANCE (Speed/Memory).
// Go runs the loop b.N times, adjusting N automatically until stable.

// 1. Basic Benchmark
func BenchmarkAdd(b *testing.B) {
	// The loop runs b.N times
	for i := 0; i < b.N; i++ {
		Add(2, 3)
	}
}

// 2. Benchmarking with Expensive Setup (ResetTimer)
// We want to measure SumSlice, but NOT the random generation.
func BenchmarkSumSlice(b *testing.B) {
	// HEAVY SETUP (Don't measure this time)
	// Create a slice of 10,000 items
	data := GenerateRandomSlice(10000)

	// RESET THE CLOCK
	// Tells Go: "Start the stopwatch NOW."
	b.ResetTimer()

	// MEASURED LOOP
	for i := 0; i < b.N; i++ {
		SumSlice(data)
	}
}

// ---------------------------------------------------------
// PART 4: PROFILING (INSTRUCTIONS)
// ---------------------------------------------------------
/*
   PROFILING IS THE "X-RAY" FOR YOUR CODE.

   How to run a Memory Profile:
   1. Run command:
      go test -bench=. -memprofile=mem.pprof

   2. Analyze the file created (mem.pprof) using the tool:
      go tool pprof mem.pprof

   3. Inside the interactive tool, type:
      > top      (Shows functions using most RAM)
      > list Add (Shows exact lines of code in the 'Add' function)
      > web      (Generates a visual graph - requires Graphviz installed)
*/
