package main

// Functions - clearer examples and explanations
// This file demonstrates common function patterns in Go and how
// to think about return types and error handling.

import (
	"errors"
	"fmt"
)

// Single return: the return type follows the parameter list.
func Add(a, b int) int {
	return a + b
}

// Multiple return values: common in Go, especially to return a
// value and an error. Here we return quotient and remainder.
func DivMod(a, b int) (quotient, remainder int) {
	// Named return values used here for readability in short functions.
	quotient = a / b
	remainder = a % b
	return // returns quotient, remainder
}

// Returning an error: when a function can fail, return (T, error).
// Callers check the error before using the value.
func Div(a, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("division by zero") // 0 is the zero-value of int
	}
	return a / b, nil
}

// Variadic functions accept any number of trailing arguments of the same type.
func Sum(vals ...int) int {
	s := 0
	for _, v := range vals {
		s += v
	}
	return s
}

// Functions as values and closures. The returned function captures `base`.
func MakeAdder(base int) func(int) int {
	return func(x int) int { return base + x }
}

func main() {
	fmt.Println("Add:", Add(2, 3))

	q, r := DivMod(13, 5)
	fmt.Printf("13 / 5 => quotient=%d remainder=%d\n", q, r)

	if v, err := Div(10, 2); err == nil {
		fmt.Println("Div(10,2):", v)
	}

	if _, err := Div(10, 0); err != nil {
		fmt.Println("Div error:", err)
	}

	fmt.Println("Sum:", Sum(1, 2, 3, 4, 5))

	add10 := MakeAdder(10)
	fmt.Println("MakeAdder(10)(5):", add10(5))
}
