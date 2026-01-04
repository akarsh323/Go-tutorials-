package main

import (
	"fmt"
	"reflect"
)

// TOPIC: Reflect
// Explanation: Reflection allows a program to examine its own structure, 
// strictly checking types at runtime. It's powerful but should be used sparingly
// as it is slower and complex.

func main() {
	// Basic Reflection
	var x float64 = 3.4
	
	fmt.Println("type:", reflect.TypeOf(x))
	fmt.Println("value:", reflect.ValueOf(x))

	// Modifying values via reflection
	// We must pass the pointer to allow modification
	p := reflect.ValueOf(&x) 
	
	// Elem() gets the value the pointer points to
	v := p.Elem() 
	
	if v.CanSet() {
		v.SetFloat(7.1)
	}
	fmt.Println("new value of x:", x)
}
