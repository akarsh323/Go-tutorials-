package main

import (
	"fmt"
	"sort"
)

/*
TOPIC: SORTING DATA

CONCEPT:
Sorting in Go is done "In-Place". This means we do not return a new slice;
we modify the memory of the existing slice directly.
Time Complexity: Generally O(n log n).

STRATEGIES:
1. PRIMITIVES: Helper functions for ints, strings, float64s.
2. INTERFACE: Implementing Len, Less, and Swap for custom types.
3. FUNCTIONAL: Creating reusable sorters for complex logic.
4. SORT.SLICE: The modern, closure-based approach (easiest).
*/

// Shared Struct for our examples
type Person struct {
	Name string
	Age  int
}

// ---------------------------------------------------------
// Example 1: The Basics (Primitives)
// ---------------------------------------------------------
func example1_Primitives() {
	fmt.Println("--- Example 1: Primitives (In-Place Mutation) ---")

	// 1. Sorting Integers
	numbers := []int{5, 3, 4, 1, 2}
	sort.Ints(numbers) // Modifies 'numbers' directly
	fmt.Println("Sorted numbers:", numbers)

	// 2. Sorting Strings
	// Strings are sorted lexicographically (A-Z)
	names := []string{"Walter", "Anthony", "Steve", "Victor", "John"}
	sort.Strings(names)
	fmt.Println("Sorted strings:", names)
	fmt.Println("----------------------------------------\n")
}

// ---------------------------------------------------------
// Example 2: The Foundation (sort.Interface)
// ---------------------------------------------------------
// To sort via Interface, we need a specific type that implements 3 methods.

type ByAge []Person

// 1. Len: How many items?
func (a ByAge) Len() int { return len(a) }

// 2. Swap: How to switch items?
func (a ByAge) Swap(i, j int) { a[i], a[j] = a[j], a[i] }

// 3. Less: Logic for "Is i smaller than j?"
func (a ByAge) Less(i, j int) bool { return a[i].Age < a[j].Age }

func example2_Interface() {
	fmt.Println("--- Example 2: The sort.Interface ---")

	people := []Person{
		{"Alice", 30},
		{"Bob", 25},
		{"Anna", 35},
	}

	// We cast 'people' to our 'ByAge' type so it gains the methods
	sort.Sort(ByAge(people))

	fmt.Println("Sorted by Age:", people)
	fmt.Println("----------------------------------------\n")
}

// ---------------------------------------------------------
// Example 3: The Professional Pattern (Functional)
// ---------------------------------------------------------
// This allows us to define the sorting logic on the fly without
// creating a new struct type (like ByAge, ByName, ByCity) for every field.

// The Function Type
type By func(p1, p2 *Person) bool

// The Sort Method attached to the function type
func (by By) Sort(people []Person) {
	ps := &personSorter{
		people: people,
		by:     by,
	}
	sort.Sort(ps)
}

// The Helper Struct
type personSorter struct {
	people []Person
	by     func(p1, p2 *Person) bool
}

// Implement Interface for the helper
func (s *personSorter) Len() int      { return len(s.people) }
func (s *personSorter) Swap(i, j int) { s.people[i], s.people[j] = s.people[j], s.people[i] }
func (s *personSorter) Less(i, j int) bool {
	// Delegate logic to the function passed in
	return s.by(&s.people[i], &s.people[j])
}

func example3_Functional() {
	fmt.Println("--- Example 3: Functional Pattern ---")

	people := []Person{
		{"Alice", 30},
		{"Bob", 25},
		{"Anna", 35},
	}

	// Logic: Sort by Name Length (Shortest name first)
	nameLenLogic := func(p1, p2 *Person) bool {
		return len(p1.Name) < len(p2.Name)
	}

	// Apply the logic using our custom type
	By(nameLenLogic).Sort(people)

	fmt.Println("Sorted by Name Length:", people)
	fmt.Println("----------------------------------------\n")
}

// ---------------------------------------------------------
// Example 4: The Modern Approach (sort.Slice)
// ---------------------------------------------------------
// No structs needed. Just pass a closure function.

func example4_SortSlice() {
	fmt.Println("--- Example 4: sort.Slice (Modern) ---")

	fruits := []string{"banana", "apple", "cherry", "grapes", "guava"}

	// Task: Sort by the LAST character of the word.
	sort.Slice(fruits, func(i, j int) bool {
		lastCharI := fruits[i][len(fruits[i])-1]
		lastCharJ := fruits[j][len(fruits[j])-1]

		if lastCharI != lastCharJ {
			return lastCharI < lastCharJ
		}
		// Tie-breaker: normal string sort
		return fruits[i] < fruits[j]
	})

	fmt.Println("Sorted by last char:", fruits)
	// Explanation: 'a' (banana/guava/apple), then 's' (grapes), then 'y' (cherry)
	fmt.Println("----------------------------------------\n")
}

func main() {
	fmt.Println("═══════════════════════════════════════════════════════════")
	fmt.Println("TOPIC: SORTING STRATEGIES")
	fmt.Println("═══════════════════════════════════════════════════════════\n")

	example1_Primitives()
	example2_Interface()
	example3_Functional()
	example4_SortSlice()

	fmt.Println("═══════════════════════════════════════════════════════════")
	fmt.Println("KEY TAKEAWAYS")
	fmt.Println("═══════════════════════════════════════════════════════════")
	fmt.Println(`
1. IN-PLACE: Sorting modifies the original slice. It does not return a new one.
2. STABLE: sort.Slice is stable (preserves original order of equal elements).
3. COMPLEXITY: Go uses efficient algorithms (Quicksort/Mergesort variants)
   running in O(n log n).
4. BEST PRACTICE: Use sort.Slice for quick tasks. Use the Interface pattern
   if you are building a library or need strict reusable types.
	`)
}
