package main

// Maps — creation, manipulation, zero value, and iteration.

import "fmt"

func main() {
	// literal map
	m := map[string]int{"alice": 30}
	// insert/update
	m["bob"] = 25
	// lookup with ok idiom
	v, ok := m["carol"]
	fmt.Println("carol present?", ok, "value:", v)

	// iterate over map (order is not guaranteed)
	for k, val := range m {
		fmt.Printf("%s -> %d\n", k, val)
	}

	// delete from map
	delete(m, "alice")
	fmt.Println("after delete:", m)
}
