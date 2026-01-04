package intermediate

import (
	"fmt"
)

// ============================================================================
// Topic 67: GENERICS - The Magic Box Analogy
// ============================================================================
//
// CONCEPT: Building a BLUEPRINT for a Magic Box that can hold *anything*,
// but only one type of thing at a time. Write the logic ONCE, use it for
// EVERY type in the universe.
//
// WHY GENERICS?
// Without Generics: You write IntStack, StringStack, FloatStack (3x code)
// With Generics:   You write Stack[T] once, works for all types
//
// ============================================================================

func main() {
	fmt.Println("=============== 67 Generics - Magic Box Analogy ===============")

	// ========================================================================
	// SECTION 1: The Blueprint - Generic Struct (The Magic Box Definition)
	// ========================================================================
	fmt.Println("\n--- SECTION 1: Stack Data Structure (The Blueprint) ---")

	// Scenario 1: Integer Stack
	intStack := Stack[int]{}
	fmt.Println("\nInteger Stack Example:")
	fmt.Println("Action 1: Push(10)")
	intStack.Push(10)
	fmt.Println("Action 2: Push(20)")
	intStack.Push(20)
	fmt.Println("Action 3: Push(30)")
	intStack.Push(30)
	fmt.Println("Stack contents (after 3 pushes): [10, 20, 30]")

	fmt.Println("\nRemoving items (LIFO - Last In, First Out):")
	val1, ok1 := intStack.Pop()
	fmt.Printf("Pop() -> %v, ok=%v (30 comes out first - Last In)\n", val1, ok1)

	val2, ok2 := intStack.Pop()
	fmt.Printf("Pop() -> %v, ok=%v (20 comes out second)\n", val2, ok2)

	peek, ok := intStack.Peek()
	fmt.Printf("Peek() -> %v, ok=%v (10 is next but stays in stack)\n", peek, ok)

	val3, ok3 := intStack.Pop()
	fmt.Printf("Pop() -> %v, ok=%v (10 comes out last)\n", val3, ok3)

	emptyVal, emptyOk := intStack.Pop()
	fmt.Printf("Pop() on empty -> %v, ok=%v (nothing left)\n", emptyVal, emptyOk)

	// Scenario 2: String Stack
	fmt.Println("\n" + "="*60)
	fmt.Println("String Stack Example:")
	stringStack := Stack[string]{}
	fmt.Println("Action 1: Push(\"Hello\")")
	stringStack.Push("Hello")
	fmt.Println("Action 2: Push(\"World\")")
	stringStack.Push("World")
	fmt.Println("Action 3: Push(\"John\")")
	stringStack.Push("John")
	fmt.Println("Stack contents (after 3 pushes): [\"Hello\", \"World\", \"John\"]")

	fmt.Println("\nRemoving items (LIFO behavior):")
	s1, _ := stringStack.Pop()
	fmt.Printf("Pop() -> \"%v\" (John comes out first - Last In)\n", s1)

	s2, _ := stringStack.Pop()
	fmt.Printf("Pop() -> \"%v\" (World comes out second)\n", s2)

	s3, _ := stringStack.Pop()
	fmt.Printf("Pop() -> \"%v\" (Hello comes out last - First In)\n", s3)

	// ========================================================================
	// SECTION 2: Generic Functions with Type Constraints
	// ========================================================================
	fmt.Println("\n" + "="*60)
	fmt.Println("--- SECTION 2: Generic Functions (Map & Max) ---")

	// MapSlice: Transform one type to another
	fmt.Println("\nMapSlice Example: Convert []int to []string")
	nums := []int{1, 2, 3, 4, 5}
	strs := MapSlice(nums, func(n int) string { return fmt.Sprintf("#%d", n) })
	fmt.Printf("Input:  %v\n", nums)
	fmt.Printf("Output: %v\n", strs)

	// Max: Find maximum of two comparable values
	fmt.Println("\nMax Function Example:")
	max_int := Max(10, 20)
	fmt.Printf("Max(10, 20) = %v (works with int)\n", max_int)

	max_float := Max(3.14, 2.71)
	fmt.Printf("Max(3.14, 2.71) = %.2f (works with float64)\n", max_float)

	max_str := MaxString("apple", "zebra")
	fmt.Printf("MaxString(\"apple\", \"zebra\") = \"%v\" (works with string)\n", max_str)

	// ========================================================================
	// SECTION 3: Understanding the Magic - Type Substitution
	// ========================================================================
	fmt.Println("\n" + "="*60)
	fmt.Println("--- SECTION 3: How the Magic Works (Type Substitution) ---")

	fmt.Println(`
When you write: intStack := Stack[int]{}
Go's compiler does this internally:
  - Replace 'T' with 'int'
  - So 'elements []T' becomes 'elements []int'
  - intStack.Push(10) is valid ✓
  - intStack.Push("hello") is INVALID ✗ (compile error)

When you write: stringStack := Stack[string]{}
Go's compiler does this instead:
  - Replace 'T' with 'string'
  - So 'elements []T' becomes 'elements []string'
  - stringStack.Push("hello") is valid ✓
  - stringStack.Push(10) is INVALID ✗ (compile error)

This is TYPE SAFETY with REUSABLE CODE!
`)

	// ========================================================================
	// SECTION 4: Advanced - Type Constraints
	// ========================================================================
	fmt.Println("--- SECTION 4: Type Constraints (Not All Types Allowed) ---")

	fmt.Println(`
Type constraints limit what types can be used.

Example: Max[T interface{ ~int | ~float64 }](a, b T)
  - 'interface{ ~int | ~float64 }' means "only int or float64 (and their aliases)"
  - Max(10, 20) ✓ (both are int)
  - Max(3.14, 2.71) ✓ (both are float64)
  - Max("a", "b") ✗ (string is not int or float64)

The '~' means: include the type AND any type that has this as underlying type
  - int and type MyInt (type MyInt int) both satisfy ~int
`)

	// ========================================================================
	// SECTION 5: Why NOT to Overuse Generics
	// ========================================================================
	fmt.Println("--- SECTION 5: When to Use Generics ---")

	fmt.Println(`
✓ USE Generics when:
  - You're writing data structures (Stack, Queue, LinkedList)
  - You're writing utility functions (Map, Filter, Sort)
  - The logic is identical for different types
  - You want type safety without code duplication

✗ DON'T Overuse Generics when:
  - The types behave differently (need separate implementations)
  - It makes code harder to read
  - You only have 1-2 concrete types
  - Simple concrete implementations are clearer

RULE OF THUMB: "Make the common case simple, make the complex case possible"
`)
}

// ============================================================================
// THE BLUEPRINT: Stack[T any] - The Generic Struct
// ============================================================================
//
// [T any] = "Generic Type Parameter"
// T = placeholder for ANY type (int, string, bool, custom structs, etc.)
// any = constraint saying "accept any type"
//
// This is the MAGIC BOX DEFINITION.

type Stack[T any] struct {
	elements []T // This slice will hold items of type T
}

// ============================================================================
// ACTION A: Push - Adding Items to the Stack
// ============================================================================
//
// func (s *Stack[T]) = receiver is a POINTER to Stack[T]
// The '*' is CRITICAL: it modifies the original stack in memory
//
// Without '*':  You'd modify a copy, original stays empty
// With '*':     You modify the original stack ✓

func (s *Stack[T]) Push(element T) {
	// Simply append the element to the end of the slice
	s.elements = append(s.elements, element)
}

// ============================================================================
// ACTION B: Pop - Removing Items from the Stack (LIFO - Last In, First Out)
// ============================================================================
//
// Returns TWO values:
//   - T:    the element we popped (or zero-value if empty)
//   - bool: true if pop succeeded, false if stack was empty
//
// Why return bool? Because you might pop from empty stack, and we need to
// tell the caller "Hey, that didn't work!"

func (s *Stack[T]) Pop() (T, bool) {
	// STEP 1: Safety Check - Can't remove from empty stack
	if len(s.elements) == 0 {
		// Create a "zero value" for T
		// If T=int, this is 0
		// If T=string, this is ""
		// If T=bool, this is false
		var zero T
		return zero, false // Return zero value and false (failed)
	}

	// STEP 2: Identify the last item (LIFO - Last In, First Out)
	// len(s.elements) - 1 gives us the index of the last item
	index := len(s.elements) - 1
	element := s.elements[index]

	// STEP 3: Shrink the list
	// s.elements[:index] means "keep everything from start up to (not including) index"
	// This effectively removes the last item
	s.elements = s.elements[:index]

	return element, true // Return the element and true (success)
}

// ============================================================================
// ACTION C: Peek - Look at the Top Without Removing
// ============================================================================
//
// Similar to Pop, but doesn't remove the element.
// Useful when you want to see what's next without taking it out.

func (s *Stack[T]) Peek() (T, bool) {
	if len(s.elements) == 0 {
		var zero T
		return zero, false
	}
	index := len(s.elements) - 1
	return s.elements[index], true
}

// ============================================================================
// GENERIC FUNCTION 1: MapSlice - Transform one type to another
// ============================================================================
//
// This is a HIGHER-ORDER FUNCTION that works with ANY types.
//
// [T any, U any] = Two type parameters
// - T = input slice element type
// - U = output slice element type
// fn func(T) U = a function that transforms T to U
//
// Example: MapSlice([]int{1,2,3}, func(n int) string { return fmt.Sprintf("#%d", n) })
//   - T = int, U = string
//   - Transforms [1, 2, 3] to ["#1", "#2", "#3"]

func MapSlice[T any, U any](in []T, fn func(T) U) []U {
	out := make([]U, len(in)) // Create output slice with same length
	for i, v := range in {
		out[i] = fn(v) // Apply transformation function
	}
	return out
}

// ============================================================================
// GENERIC FUNCTION 2: Max - Find maximum of two values
// ============================================================================
//
// [T interface{ ~int | ~float64 }] = Type Constraint
// This says: "T can ONLY be int or float64 (or aliases thereof)"
//
// Why the constraint? Because you can't compare all types with >
// You can't do: Max("apple", "banana") with this constraint
// (well, you CAN, but we didn't allow it)
//
// The ~ symbol means: "include the type AND any alias of it"
// For example, these would all work:
//   - type CustomInt int
//   - type MyFloat float64

func Max[T interface{ ~int | ~float64 }](a, b T) T {
	if a > b {
		return a
	}
	return b
}

// ============================================================================
// GENERIC FUNCTION 3: MaxString - Constrained to Strings Only
// ============================================================================
//
// [T interface{ ~string }] = Only string type (and its aliases)
// Go doesn't allow strings with the ~int | ~float64 constraint above.

func MaxString[T interface{ ~string }](a, b T) T {
	if a > b {
		return a
	}
	return b
}

// ============================================================================
// EXPLANATION: Why This Is Better
// ============================================================================
//
// SCENARIO 1: WITHOUT Generics (The Old Way)
//
//   type IntStack struct {
//       elements []int
//   }
//   func (s *IntStack) Push(e int) { s.elements = append(s.elements, e) }
//   func (s *IntStack) Pop() (int, bool) { /* same logic */ }
//
//   type StringStack struct {
//       elements []string
//   }
//   func (s *StringStack) Push(e string) { s.elements = append(s.elements, e) }
//   func (s *StringStack) Pop() (string, bool) { /* same logic */ }
//
//   type FloatStack struct {
//       elements []float64
//   }
//   func (s *FloatStack) Push(e float64) { s.elements = append(s.elements, e) }
//   func (s *FloatStack) Pop() (float64, bool) { /* same logic */ }
//
// Total: 3 structs + 6 methods = LOTS OF DUPLICATED CODE
//
// ============================================================================
//
// SCENARIO 2: WITH Generics (This Way)
//
//   type Stack[T any] struct {
//       elements []T
//   }
//   func (s *Stack[T]) Push(e T) { s.elements = append(s.elements, e) }
//   func (s *Stack[T]) Pop() (T, bool) { /* same logic */ }
//
// Total: 1 struct + 2 methods = SINGLE BLUEPRINT, WORKS FOR ALL TYPES
// Use it: Stack[int]{}, Stack[string]{}, Stack[float64]{}, etc.
//
// ============================================================================
