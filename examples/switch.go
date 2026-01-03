package main

// Conditions: Switch — switch statement, fallthrough, type switch,
// and switch without expression.

import "fmt"

func main() {
	// Basic switch
	n := 2
	switch n {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	default:
		fmt.Println("other")
	}

	// Switch with fallthrough
	fmt.Println("\nswitch with fallthrough:")
	x := 1
	switch x {
	case 1:
		fmt.Println("case 1")
		fallthrough
	case 2:
		fmt.Println("case 2 (after fallthrough)")
	default:
		fmt.Println("default")
	}

	// Switch without expression (like if/else if)
	fmt.Println("\nswitch without expression:")
	age := 25
	switch {
	case age < 13:
		fmt.Println("child")
	case age < 18:
		fmt.Println("teen")
	case age < 65:
		fmt.Println("adult")
	default:
		fmt.Println("senior")
	}

	// Type switch
	fmt.Println("\ntype switch:")
	var i interface{} = "hello"
	switch v := i.(type) {
	case int:
		fmt.Println("int:", v)
	case string:
		fmt.Println("string:", v)
	default:
		fmt.Println("unknown type")
	}
}
