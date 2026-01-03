package intermediate

import "fmt"

// Topic 63: structs
// Combines pragmatic examples with GoBootcamp patterns


func main() {

	fmt.Println("-- 63 Structs --")

	// Struct literal
	p := DemoPerson63{Name: "Gopher", Age: 5}
	fmt.Println("person:", p.Name, p.Age)

	// Embedded struct
	e := Employee{
		DemoPerson63: DemoPerson63{Name: "Alice", Age: 30},
		Department:   "Engineering",
	}
	fmt.Println("employee:", e.Name, e.Department)

	// Pointers to structs
	person_ptr := &p
	person_ptr.Name = "Updated"
	fmt.Println("after update:", p.Name)
}
