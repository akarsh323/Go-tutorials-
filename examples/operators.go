package main

// Operators — comparison, arithmetic, logical, bitwise, and assignment.

import "fmt"

func main() {
	a := 10
	b := 3

	// Comparison operators
	fmt.Println("comparison:")
	fmt.Println("a == b", a == b)
	fmt.Println("a != b", a != b)
	fmt.Println("a < b", a < b)
	fmt.Println("a > b", a > b)
	fmt.Println("a <= b", a <= b)
	fmt.Println("a >= b", a >= b)

	// Arithmetic operators
	fmt.Println("\narithmetic:")
	fmt.Println("a + b =", a+b)
	fmt.Println("a - b =", a-b)
	fmt.Println("a * b =", a*b)
	fmt.Println("a / b =", a/b)
	fmt.Println("a % b =", a%b)

	// Logical operators
	fmt.Println("\nlogical:")
	fmt.Println("true && false =", true && false)
	fmt.Println("true || false =", true || false)
	fmt.Println("!true =", !true)

	// Bitwise operators
	fmt.Println("\nbitwise:")
	fmt.Println("a & b =", a&b)   // AND
	fmt.Println("a | b =", a|b)   // OR
	fmt.Println("a ^ b =", a^b)   // XOR
	fmt.Println("a << 1 =", a<<1) // left shift
	fmt.Println("a >> 1 =", a>>1) // right shift

	// Assignment operators
	c := 5
	c += 2
	fmt.Println("\nassignment: c += 2 results in", c)
}
