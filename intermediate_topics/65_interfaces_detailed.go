package main

import (
	"fmt"
	"math"
)

/*
Topic 65: INTERFACES - The Power of Contracts

═══════════════════════════════════════════════════════════════════════════════

WHAT IS AN INTERFACE?

An interface is a "CONTRACT" that specifies what methods a type must have.

Think of it like this:
- A restaurant menu is a "contract" (interface)
- Each chef implements that menu (concrete types)
- A customer doesn't care WHO cooks the food, just that the menu items exist
- The customer works with the "menu interface", not specific chefs

IN GO:
- Interface = A set of method signatures
- Concrete Type = A struct that implements all those methods
- The compiler checks: "Does this type have all required methods?"

═══════════════════════════════════════════════════════════════════════════════

KEY CONCEPT: IMPLICIT IMPLEMENTATION

Go uses IMPLICIT implementation, NOT explicit "implements" keywords.

You don't write:
  type rect struct implements geometry { }  ❌ WRONG

You just write the matching methods:
  func (r rect) area() float64 { }
  func (r rect) perim() float64 { }

If you have these methods, you automatically satisfy the interface!

═══════════════════════════════════════════════════════════════════════════════

PART 1: DECLARING THE INTERFACE AND STRUCTS

═════════════════════════════════════════════════════════════════════════════

THE INTERFACE (The Contract):
This says: "To be a 'geometry', you must have area() and perim() methods"
*/

type geometry interface {
	area() float64
	perim() float64
}

// THE STRUCTS (Concrete Types):
// These are just data structures. They have NO relationship to geometry yet.

type rect struct {
	width, height float64
}

type circle struct {
	radius float64
}

/*
═════════════════════════════════════════════════════════════════════════════

PART 2: IMPLEMENTING THE INTERFACE METHODS

Notice: We write NO "implements" keyword. Just the matching methods.
These implementations are "implicit" - Go checks if methods match the interface.

IMPORTANT: Value Receiver vs Pointer Receiver
- Use value receiver (r rect) if only READING data (no modification)
- Use pointer receiver (r *rect) if MODIFYING data

═════════════════════════════════════════════════════════════════════════════
*/

// RECTANGLE IMPLEMENTATION
func (r rect) area() float64 {
	return r.width * r.height
}

func (r rect) perim() float64 {
	return 2 * (r.width + r.height)
}

// CIRCLE IMPLEMENTATION
func (c circle) area() float64 {
	// math.Pi is capitalized because it's exported (public)
	// If it were math.pi (lowercase), we couldn't access it (private)
	return math.Pi * c.radius * c.radius
}

func (c circle) perim() float64 {
	return 2 * math.Pi * c.radius
}

// EXTRA METHOD (NOT required by the interface)
// Circle has an extra diameter() method. Having extra methods is fine!
// You only need the MINIMUM required methods to satisfy an interface.
func (c circle) diameter() float64 {
	return 2 * c.radius
}

/*
═════════════════════════════════════════════════════════════════════════════

PART 3: POLYMORPHISM - THE MEASURE FUNCTION

This is where interfaces become powerful.

"measure" doesn't ask for a specific rect or circle.
It asks for ANYTHING that behaves like "geometry".

This function is DECOUPLED from concrete types.
It works with ANY shape that satisfies the geometry interface.

═════════════════════════════════════════════════════════════════════════════
*/

func measure(g geometry) {
	fmt.Println("Shape:", g)
	fmt.Println("Area:", g.area())
	fmt.Println("Perimeter:", g.perim())
	fmt.Println()
}

/*
═════════════════════════════════════════════════════════════════════════════

PART 4: THE EMPTY INTERFACE interface{}

The "empty interface" has ZERO required methods.

Since every type has at least zero methods, everything satisfies interface{}.

Common uses:
- function accepts any argument
- Storing different types in one variable/slice
- Generic-like behavior before Go 1.18 generics

In Go 1.18+, interface{} is often aliased as "any"
*/

func printType(i interface{}) {
	// i.(type) is a TYPE SWITCH - only valid inside switch statements
	// It dynamically checks what type of data 'i' is holding at runtime
	switch i.(type) {
	case int:
		fmt.Println("This is an integer:", i)
	case string:
		fmt.Println("This is a string:", i)
	case bool:
		fmt.Println("This is a boolean:", i)
	default:
		fmt.Println("Unknown type:", i)
	}
}

/*
═════════════════════════════════════════════════════════════════════════════

PART 5: TYPE ASSERTION (Getting the actual value)

Type switch tells you WHAT type something is.
Type assertion lets you EXTRACT and USE the actual value.

═════════════════════════════════════════════════════════════════════════════
*/

func processValue(i interface{}) {
	// Type assertion: i.(int) tries to convert i to int
	// The second return value is a boolean indicating success
	if num, ok := i.(int); ok {
		fmt.Printf("Got integer: %d, doubled: %d\n", num, num*2)
	} else if str, ok := i.(string); ok {
		fmt.Printf("Got string: %s, length: %d\n", str, len(str))
	}
}

/*
═════════════════════════════════════════════════════════════════════════════

PART 6: INCOMPLETE IMPLEMENTATION (ERROR SCENARIO)

What happens if a type only PARTIALLY satisfies an interface?
═════════════════════════════════════════════════════════════════════════════
*/

type rect1 struct {
	width, height float64
}

// rect1 ONLY implements area(), NOT perim()
func (r rect1) area() float64 {
	return r.width * r.height
}

// If we tried to use rect1 as geometry:
// r1 := rect1{width: 10, height: 10}
// measure(r1)  ❌ COMPILE ERROR: "rect1 does not implement geometry (missing method perim)"

/*
═════════════════════════════════════════════════════════════════════════════

PART 7: VARIADIC INTERFACE (Multiple arguments of any type)

The "..." makes a parameter variadic - accept as many arguments as you want.
Combined with interface{}, you can pass ANY number of ANY types.
═════════════════════════════════════════════════════════════════════════════
*/

func printMany(items ...interface{}) {
	for i, item := range items {
		fmt.Printf("Item %d: %v (Type: %T)\n", i+1, item, item)
	}
}

/*
═════════════════════════════════════════════════════════════════════════════

PART 8: BENEFITS OF INTERFACES

═════════════════════════════════════════════════════════════════════════════
*/

func main() {
	fmt.Println("════════════════════════════════════════════════════════════")
	fmt.Println("PART 1: BASIC INTERFACE & POLYMORPHISM")
	fmt.Println("════════════════════════════════════════════════════════════\n")

	// Create instances of rect and circle
	r := rect{width: 3, height: 4}
	c := circle{radius: 5}

	fmt.Println("Rectangle (3 × 4):")
	measure(r)

	fmt.Println("Circle (radius 5):")
	measure(c)

	fmt.Println("✓ Both types work in 'measure' function even though they're different!")
	fmt.Println("  'measure' accepts the geometry interface, not specific types.\n")

	// ─────────────────────────────────────────────────────────────────────────

	fmt.Println("════════════════════════════════════════════════════════════")
	fmt.Println("PART 2: EXTRA METHODS (Not part of interface)")
	fmt.Println("════════════════════════════════════════════════════════════\n")

	fmt.Printf("Circle diameter (extra method): %.2f\n", c.diameter())
	fmt.Println("✓ Circle has this extra method. That's fine!")
	fmt.Println("  To satisfy an interface, you need the MINIMUM required methods.\n")

	// ─────────────────────────────────────────────────────────────────────────

	fmt.Println("════════════════════════════════════════════════════════════")
	fmt.Println("PART 3: STORING INTERFACE TYPES IN A SLICE")
	fmt.Println("════════════════════════════════════════════════════════════\n")

	// Create a slice of geometry interface
	shapes := []geometry{r, c}

	fmt.Println("Measuring all shapes in a single slice:")
	for i, shape := range shapes {
		fmt.Printf("\nShape %d:\n", i+1)
		measure(shape)
	}

	fmt.Println("✓ Different types stored in the same slice!")
	fmt.Println("  This is the power of interfaces - they unify different types.\n")

	// ─────────────────────────────────────────────────────────────────────────

	fmt.Println("════════════════════════════════════════════════════════════")
	fmt.Println("PART 4: EMPTY INTERFACE & TYPE SWITCHING")
	fmt.Println("════════════════════════════════════════════════════════════\n")

	fmt.Println("Testing different types with empty interface:")
	printType(42)
	printType("Hello, Go!")
	printType(3.14)
	fmt.Println()

	// ─────────────────────────────────────────────────────────────────────────

	fmt.Println("════════════════════════════════════════════════════════════")
	fmt.Println("PART 5: TYPE ASSERTION (Extracting actual values)")
	fmt.Println("════════════════════════════════════════════════════════════\n")

	processValue(100)
	processValue("Go Language")
	fmt.Println()

	// ─────────────────────────────────────────────────────────────────────────

	fmt.Println("════════════════════════════════════════════════════════════")
	fmt.Println("PART 6: VARIADIC INTERFACE (Multiple items of any type)")
	fmt.Println("════════════════════════════════════════════════════════════\n")

	printMany(42, "text", 3.14, true, []int{1, 2, 3})
	fmt.Println()

	// ─────────────────────────────────────────────────────────────────────────

	fmt.Println("════════════════════════════════════════════════════════════")
	fmt.Println("PART 7: KEY CONCEPTS SUMMARY")
	fmt.Println("════════════════════════════════════════════════════════════\n")

	fmt.Println(`
WHAT MAKES GO INTERFACES POWERFUL:

1. IMPLICIT IMPLEMENTATION
   - No "implements" keyword needed
   - Just write the matching methods
   - Go checks automatically at compile time

2. DECOUPLING
   - Functions work with interfaces, not concrete types
   - Add new types without changing existing code
   - Follow the "Open/Closed Principle"

3. POLYMORPHISM
   - Same function, different behaviors
   - measure(rect) vs measure(circle) - same function call
   - Different execution based on actual type

4. COMPOSITION
   - Types satisfy multiple interfaces (if they implement all methods)
   - A type can implement many different interfaces
   - More flexible than traditional inheritance

5. EMPTY INTERFACE interface{}
   - Accepts ANY type (has zero required methods)
   - Enables dynamic typing when needed
   - Type switches and type assertions provide type safety

6. DUCK TYPING
   - "If it walks like a duck and quacks like a duck, it's a duck"
   - Go doesn't care about the name or inheritance chain
   - Only cares about the methods (behavior)

═════════════════════════════════════════════════════════════════════════════

PRACTICAL EXAMPLE: IO Package

Go's io package uses interfaces heavily:

type Reader interface {
    Read(p []byte) (n int, err error)
}

type Writer interface {
    Write(p []byte) (n int, err error)
}

Now:
- File implements Reader and Writer
- Socket implements Reader and Writer
- Buffer implements Reader and Writer
- Network connection implements Reader and Writer

A single function can work with ANY of these:

func Copy(dst Writer, src Reader) { }

This works with files, sockets, buffers, and anything else that implements
Reader/Writer. The function doesn't know or care about the concrete type!

═════════════════════════════════════════════════════════════════════════════

WHEN TO USE INTERFACES:

✓ When you want to decouple code from concrete types
✓ When multiple types need to behave the same way
✓ When you want to accept different types in functions
✓ When building packages/libraries for others to extend
✓ When you need to mock types for testing
✓ When handling unknown types (empty interface)

═════════════════════════════════════════════════════════════════════════════

NEXT STEP: Methods

Now that you understand interfaces, the next topic is METHODS, which are
the building blocks of interfaces. Methods bind behavior to types, and
interfaces define what behaviors are required.

After interfaces and methods come STRUCT EMBEDDING and more advanced patterns!
`)
}
