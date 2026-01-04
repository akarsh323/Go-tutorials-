package main

import (
	"fmt"
	"testing"
)

// TOPIC: Testing / Benchmarking
// Explanation: Go has a built-in 'testing' package.
// Files must end in '_test.go'. Functions must start with 'Test'.
// RUN THIS USING: go test -v

// Simple function to test
func IntMin(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// The Test Function
func TestIntMinBasic(t *testing.T) {
	ans := IntMin(2, -2)
	if ans != -2 {
		// t.Errorf reports a failure but continues execution
		t.Errorf("IntMin(2, -2) = %d; want -2", ans)
	}
}

// Table-driven tests (Common Go pattern)
func TestIntMinTableDriven(t *testing.T) {
	var tests = []struct {
		a, b int
		want int
	}{
		{0, 1, 0},
		{1, 0, 0},
		{2, -2, -2},
		{0, -1, -1},
		{-1, 0, -1},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("%d,%d", tt.a, tt.b)
		t.Run(testname, func(t *testing.T) {
			ans := IntMin(tt.a, tt.b)
			if ans != tt.want {
				t.Errorf("got %d, want %d", ans, tt.want)
			}
		})
	}
}

// Benchmarking Function
// Run using: go test -bench=.
func BenchmarkIntMin(b *testing.B) {
	for i := 0; i < b.N; i++ {
		IntMin(1, 2)
	}
}
