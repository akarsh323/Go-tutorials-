package intermediate

import (
	"fmt"
	"unsafe"
)

// Topic 59: pointers
// Combines pragmatic examples with GoBootcamp patterns


func main() {

	fmt.Println("-- 59 Pointers --")
	x := 10
	p := &x // p is *int
	fmt.Println("initial x:", x, "address p:", p)
	*p = 20
	fmt.Println("x after *p = 20:", x)

	// Passing pointer to a function to modify caller's value
	increment := func(ptr *int) {
		*ptr++
	}
	increment(&x)
	fmt.Println("x after increment(&x):", x)
}

func pointerBonusExample() {
	fmt.Println("\n=== BONUS: Pointer Best Practices ===")
	fmt.Println(`
1. Use & to get address: ptr := &variable
2. Use * to dereference: value := *ptr
3. Always check nil before dereferencing: if ptr != nil { ... }
4. Pointer receivers (*T) can modify structs
5. Use pointers for large data structures
6. Don't return pointers to local variables
7. Slice and map internals already use pointers
8. Strings and ints are usually passed by value
9. Use pointers for optional struct fields
10. Method chaining is enabled with pointer receivers
	`)
}
