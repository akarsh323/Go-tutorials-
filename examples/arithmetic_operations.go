package main

// Arithmetic Operations — improved teaching example.
// This file demonstrates:
// - basic arithmetic operators (+, -, *, /, %)
// - integer vs floating-point division
// - operator precedence and mixing types
// - bitwise operators (brief)

import (
	"fmt"
)

func main() {
	// Integers: division truncates toward zero
	a, b := 7, 3
	fmt.Printf("%d + %d = %d\n", a, b, a+b)
	fmt.Printf("%d - %d = %d\n", a, b, a-b)
	fmt.Printf("%d * %d = %d\n", a, b, a*b)
	fmt.Printf("%d / %d = %d (integer division)\n", a, b, a/b)
	fmt.Printf("%d %% %d = %d (remainder)\n", a, b, a%b)

	// Floating point division: convert one operand to float
	af := float64(a)
	bf := float64(b)
	fmt.Printf("%d / %d = %.6f (float division)\n", a, b, af/bf)

	// Mixing types: must convert explicitly
	var x int = 5
	var y float64 = 2.0
	// fmt.Println(x + y) // compile error: mismatched types
	fmt.Printf("mixed: %d + %.1f => %.1f\n", x, y, float64(x)+y)

	// Operator precedence: *, /, % before +, -
	fmt.Println("precedence: 1 + 2*3 =", 1+2*3)

	// Bitwise ops (integers): &, |, ^, <<, >>
	p, q := 6, 3 // binary: 110 and 011
	fmt.Printf("bitwise: %d & %d = %d\n", p, q, p&q)
	fmt.Printf("bitwise: %d | %d = %d\n", p, q, p|q)
	fmt.Printf("bitwise: %d ^ %d = %d\n", p, q, p^q)
}
