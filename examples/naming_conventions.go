package main

// Naming Conventions — exported (Capital) vs unexported (lowercase),
// CamelCase, package names, and constants.

import "fmt"

// Exported type (visible outside package)
type Person struct {
	Name string // exported field (Capital letter)
	Age  int    // exported field
	age  int    // unexported field (lowercase, hidden)
}

// Exported function (visible outside package)
func (p Person) String() string {
	return fmt.Sprintf("%s (%d)", p.Name, p.Age)
}

// unexported function (hidden, only usable in package)
func privateHelper() string {
	return "hidden from outside"
}

// Constants: follow same rule (exported vs unexported)
const MaxConnections = 100
const defaultTimeout = 30

func main() {
	p := Person{Name: "Akarsh", Age: 30, age: 99}
	fmt.Println("Person:", p)
	fmt.Println("MaxConnections:", MaxConnections)
	fmt.Println("private helper:", privateHelper())
}
