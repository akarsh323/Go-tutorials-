package main

// The Standard Library — overview of essential packages: fmt, time,
// strings, strconv, io, os, and more.

import (
	"fmt"
	"math"
	"strconv"
	"strings"
	"time"
)

func main() {
	fmt.Println("=== Go Standard Library Examples ===")

	// fmt package: printing and formatting
	fmt.Println("\nfmt package:")
	fmt.Printf("  formatted: %s %d %.2f\n", "hello", 42, 3.14159)

	// time package
	fmt.Println("\ntime package:")
	now := time.Now()
	fmt.Println("  now:", now.Format(time.RFC3339))
	fmt.Println("  unix epoch:", now.Unix())

	// strings package
	fmt.Println("\nstrings package:")
	s := "  hello world  "
	fmt.Println("  TrimSpace:", strings.TrimSpace(s))
	fmt.Println("  ToUpper:", strings.ToUpper(s))
	fmt.Println("  Contains(hello):", strings.Contains(s, "hello"))

	// strconv package
	fmt.Println("\nstrconv package:")
	num, _ := strconv.Atoi("42")
	fmt.Println("  Atoi(\"42\"):", num)
	f, _ := strconv.ParseFloat("3.14", 64)
	fmt.Println("  ParseFloat(\"3.14\"):", f)

	// math package
	fmt.Println("\nmath package:")
	fmt.Println("  Sqrt(2):", math.Sqrt(2))
	fmt.Println("  Pow(2,3):", math.Pow(2, 3))
}
