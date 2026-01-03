package main

import "fmt"

/*
Topic 63: STRUCTS - A Comprehensive Guide

═══════════════════════════════════════════════════════════════════════════════

WHAT IS A STRUCT?

A struct is a composite data type that groups multiple data fields together.
It allows you to create custom data types by combining different types.

Think of a struct like a "blueprint" for creating objects:
- A house blueprint describes walls, doors, windows, and roof
- A Person struct describes name, age, email, and address

KEY CONCEPTS:

1. FIELDS: Variables inside a struct (also called member variables or attributes)
2. STRUCT TYPE: The definition/blueprint itself
3. STRUCT VALUE: An instance created from the blueprint
4. VISIBILITY: Exported (Capitalized) or Unexported (lowercase) fields

═══════════════════════════════════════════════════════════════════════════════

PART 1: BASIC STRUCT DEFINITION & INITIALIZATION

═════════════════════════════════════════════════════════════════════════════

SYNTAX:
type StructName struct {
    FieldName1 FieldType1
    FieldName2 FieldType2
    // ...
}

INITIALIZATION:
1. Zero Value (all fields get default values):
   var p Person

2. With Field Names (struct literal):
   p := Person{Name: "John", Age: 25}

3. With Position (order matters!):
   p := Person{"John", 25}

4. Partial Initialization:
   p := Person{Name: "John"}
   // Age gets its zero value (0)

═════════════════════════════════════════════════════════════════════════════

PART 2: POINTERS TO STRUCTS

═════════════════════════════════════════════════════════════════════════════

Why use pointers to structs?
- Structs can be large (many fields)
- Copying large structs is inefficient
- Use pointers to avoid unnecessary copying
- Pointers allow modification of original struct

POINTER SYNTAX:
var ptr *StructName = &instance
// OR
ptr := &instance

ACCESSING FIELDS THROUGH POINTER:
Go automatically dereferences pointers when accessing struct fields!
- ptr.Field (Go's syntactic sugar)
- (*ptr).Field (explicit dereferencing, rarely needed)

═════════════════════════════════════════════════════════════════════════════

PART 3: COMPOSITION ("HAS-A" RELATIONSHIP) & FIELD PROMOTION

═════════════════════════════════════════════════════════════════════════════

THE PROBLEM WITH TRADITIONAL INHERITANCE:

In Java, Python, C++:
- "Car IS A Vehicle" (inheritance)
- Car inherits all methods and fields automatically
- Problem: Rigid hierarchy, doesn't work well in practice

THE GO SOLUTION: COMPOSITION

In Go:
- "Car HAS AN Engine" (composition)
- You embed one struct inside another
- You get the fields you need, nothing more
- More flexible and compositional

THE "LEGO" ANALOGY:

Imagine building with Legos:
- A simple Lego brick = struct with few fields
- A complex Lego structure = struct with embedded structs
- You "snap together" smaller pieces to make larger ones
- This is COMPOSITION

EXAMPLE: Car has an Engine and Wheels

```
    Car
    ├── Model (string)
    ├── Make (string)
    ├── Engine (EMBEDDED)
    │   ├── Horsepower (int)
    │   └── Type (string)
    └── Wheels (EMBEDDED)
        ├── Count (int)
        └── Brand (string)
```

═════════════════════════════════════════════════════════════════════════════

FIELD PROMOTION (The "Shortcut")

When you embed a struct, Go automatically "promotes" its fields.

WITHOUT PROMOTION (The Long Way):
myCar.Engine.Horsepower ✓ Valid, but tedious

WITH PROMOTION (The Promoted Way):
myCar.Horsepower ✓ Go finds it automatically!

This works because:
- Engine struct is EMBEDDED (no field name, just the type)
- Go treats Engine's fields as if they belong to Car directly
- Saves typing long chains like: car.Engine.Parts.Bolts.Size

IMPORTANT: If Car also had a field named Horsepower, shadowing occurs
and you must use the explicit way: myCar.Engine.Horsepower

═════════════════════════════════════════════════════════════════════════════

BENEFITS OF COMPOSITION:

1. ORGANIZATION:
   - Keep Engine logic separate from Car logic
   - Reusable components (Engine can be used in other vehicles)
   - Clear separation of concerns

2. CONVENIENCE:
   - Field promotion saves typing
   - Don't need deep nesting
   - Access promoted fields directly

3. FLEXIBILITY:
   - Easy to add, remove, or change embedded structs
   - No rigid hierarchy
   - True composition of parts

═════════════════════════════════════════════════════════════════════════════
*/

// PART 1: BASIC STRUCT DEFINITION

type Person struct {
	Name  string
	Age   int
	Email string
}

type Address struct {
	Street string
	City   string
	State  string
	ZIP    string
}

// PART 2: COMPOSITION - Inner Structs (Parts)

type Engine struct {
	Horsepower int
	Type       string
	Cylinders  int
}

type Wheels struct {
	Count int
	Brand string
	Size  string
}

// PART 3: COMPOSITION - Outer Struct (Whole)
// This struct EMBEDS Engine and Wheels
type Car struct {
	Make   string
	Model  string
	Year   int
	Engine // EMBEDDED (no field name, just type)
	Wheels // EMBEDDED (no field name, just type)
}

// Another example: Employee with embedded Person
type Employee struct {
	Person     // EMBEDDED
	EmployeeID int
	Department string
	Salary     float64
}

func main() {
	fmt.Println("════════════════════════════════════════════════════════════")
	fmt.Println("PART 1: BASIC STRUCT DECLARATION & INITIALIZATION")
	fmt.Println("════════════════════════════════════════════════════════════\n")

	// Method 1: Zero value (all fields get default values)
	fmt.Println("Method 1: Zero Value Initialization")
	var person1 Person
	fmt.Printf("Zero value: %+v\n\n", person1)

	// Method 2: Named field initialization
	fmt.Println("Method 2: Named Field Initialization")
	person2 := Person{
		Name:  "John Doe",
		Age:   30,
		Email: "john@example.com",
	}
	fmt.Printf("person2: %+v\n\n", person2)

	// Method 3: Positional initialization (order matters!)
	fmt.Println("Method 3: Positional Initialization")
	person3 := Person{"Alice", 28, "alice@example.com"}
	fmt.Printf("person3: %+v\n\n", person3)

	// Method 4: Partial initialization
	fmt.Println("Method 4: Partial Initialization")
	person4 := Person{Name: "Bob", Email: "bob@example.com"}
	// Age gets zero value (0)
	fmt.Printf("person4: %+v\n\n", person4)

	// ─────────────────────────────────────────────────────────────────────────

	fmt.Println("════════════════════════════════════════════════════════════")
	fmt.Println("PART 2: ACCESSING & MODIFYING STRUCT FIELDS")
	fmt.Println("════════════════════════════════════════════════════════════\n")

	person := Person{Name: "Jane", Age: 25, Email: "jane@example.com"}

	fmt.Println("Original struct:")
	fmt.Printf("Name: %s, Age: %d, Email: %s\n\n", person.Name, person.Age, person.Email)

	// Modify fields
	person.Age = 26
	person.Email = "jane.updated@example.com"

	fmt.Println("After modification:")
	fmt.Printf("Name: %s, Age: %d, Email: %s\n\n", person.Name, person.Age, person.Email)

	// ─────────────────────────────────────────────────────────────────────────

	fmt.Println("════════════════════════════════════════════════════════════")
	fmt.Println("PART 3: POINTERS TO STRUCTS")
	fmt.Println("════════════════════════════════════════════════════════════\n")

	original := Person{Name: "Charlie", Age: 35, Email: "charlie@example.com"}
	fmt.Println("Original:", original)

	// Create a pointer to the struct
	ptr := &original
	fmt.Printf("Pointer value: %p\n", ptr)
	fmt.Printf("Dereferenced: %+v\n\n", *ptr)

	// Modify through pointer (Go auto-dereferences for fields!)
	fmt.Println("Modifying through pointer:")
	ptr.Name = "Charles"
	ptr.Age = 36

	fmt.Println("Original after modification:", original)
	fmt.Println("(Original was modified because pointer points to it)\n")

	// ─────────────────────────────────────────────────────────────────────────

	fmt.Println("════════════════════════════════════════════════════════════")
	fmt.Println("PART 4: COMPOSITION - THE LEGO APPROACH")
	fmt.Println("════════════════════════════════════════════════════════════\n")

	fmt.Println("CONCEPT: Car HAS-AN Engine and HAS Wheels\n")

	// Create a Car with embedded Engine and Wheels
	myCar := Car{
		Make:  "Tesla",
		Model: "Model S",
		Year:  2024,
		Engine: Engine{
			Horsepower: 1020,
			Type:       "Electric",
			Cylinders:  0, // Electric cars have no cylinders
		},
		Wheels: Wheels{
			Count: 4,
			Brand: "Michelin",
			Size:  "19 inch",
		},
	}

	fmt.Printf("Car created: %d %s %s\n\n", myCar.Year, myCar.Make, myCar.Model)

	// ─────────────────────────────────────────────────────────────────────────

	fmt.Println("════════════════════════════════════════════════════════════")
	fmt.Println("PART 5: FIELD PROMOTION (THE SHORTCUT)")
	fmt.Println("════════════════════════════════════════════════════════════\n")

	fmt.Println("Accessing fields through embedded structs:\n")

	// THE EXPLICIT WAY (always valid)
	fmt.Println("EXPLICIT WAY (accessing through the embedded struct):")
	fmt.Printf("Engine Horsepower (explicit): %d HP\n", myCar.Engine.Horsepower)
	fmt.Printf("Engine Type (explicit): %s\n", myCar.Engine.Type)
	fmt.Printf("Wheels Count (explicit): %d\n", myCar.Wheels.Count)
	fmt.Printf("Wheels Brand (explicit): %s\n\n", myCar.Wheels.Brand)

	// THE PROMOTED WAY (Go's convenience feature)
	fmt.Println("PROMOTED WAY (field promotion - Go's syntactic sugar):")
	fmt.Printf("Horsepower (promoted): %d HP\n", myCar.Horsepower)
	fmt.Printf("Type (promoted): %s\n", myCar.Type)
	fmt.Printf("Count (promoted): %d\n", myCar.Count)
	fmt.Printf("Brand (promoted): %s\n\n", myCar.Brand)

	fmt.Println("✓ Both ways work! Go automatically promotes embedded fields.")
	fmt.Println("  The promoted way is more convenient and cleaner.\n")

	// ─────────────────────────────────────────────────────────────────────────

	fmt.Println("════════════════════════════════════════════════════════════")
	fmt.Println("PART 6: MODIFYING PROMOTED FIELDS")
	fmt.Println("════════════════════════════════════════════════════════════\n")

	fmt.Println("Original Horsepower:", myCar.Horsepower)

	// Modify using promoted field
	myCar.Horsepower = 900
	fmt.Println("After modification:", myCar.Horsepower)
	fmt.Println("Verified through explicit access:", myCar.Engine.Horsepower)
	fmt.Println("(Both refer to the same field!)\n")

	// ─────────────────────────────────────────────────────────────────────────

	fmt.Println("════════════════════════════════════════════════════════════")
	fmt.Println("PART 7: MULTIPLE COMPOSITION - EMPLOYEE EXAMPLE")
	fmt.Println("════════════════════════════════════════════════════════════\n")

	emp := Employee{
		Person: Person{
			Name:  "Diana",
			Age:   32,
			Email: "diana@company.com",
		},
		EmployeeID: 12345,
		Department: "Engineering",
		Salary:     95000.50,
	}

	fmt.Println("Employee Information:")
	fmt.Printf("Name (promoted): %s\n", emp.Name)
	fmt.Printf("Age (promoted): %d\n", emp.Age)
	fmt.Printf("Email (promoted): %s\n", emp.Email)
	fmt.Printf("Employee ID: %d\n", emp.EmployeeID)
	fmt.Printf("Department: %s\n", emp.Department)
	fmt.Printf("Salary: $%.2f\n\n", emp.Salary)

	fmt.Println("Comparison:")
	fmt.Printf("Explicit way: emp.Person.Name = %s\n", emp.Person.Name)
	fmt.Printf("Promoted way: emp.Name = %s\n", emp.Name)
	fmt.Println("(Same field, cleaner syntax!)\n")

	// ─────────────────────────────────────────────────────────────────────────

	fmt.Println("════════════════════════════════════════════════════════════")
	fmt.Println("PART 8: COMPOSITION BENEFITS VISUALIZATION")
	fmt.Println("════════════════════════════════════════════════════════════\n")

	fmt.Println(`
TRADITIONAL INHERITANCE (Java, C++, Python):
Car
  ├── Vehicle (inherited)
  │   ├── Speed
  │   ├── Direction
  │   └── Drive() method
  ├── Engine (inherited)
  └── ...
Problem: Rigid hierarchy, multiple inheritance conflicts, "Fragile Base Class"

GO COMPOSITION (Lego Blocks):
Car
  ├── Make (string)
  ├── Model (string)
  ├── Engine (EMBEDDED)
  │   ├── Horsepower (int)
  │   ├── Type (string)
  │   └── Methods on Engine...
  └── Wheels (EMBEDDED)
      ├── Count (int)
      ├── Brand (string)
      └── Methods on Wheels...

Benefits:
✓ Organization: Engine logic separate from Car logic
✓ Reusability: Engine can be embedded in Truck, Motorcycle, etc.
✓ Flexibility: Easy to add/remove/change components
✓ Clarity: "Car HAS an Engine" is clearer than "Car IS-A Engine"
✓ No conflicts: Easily handle multiple similar components
`)

	// ─────────────────────────────────────────────────────────────────────────

	fmt.Println("════════════════════════════════════════════════════════════")
	fmt.Println("PART 9: BEST PRACTICES & SUMMARY TABLE")
	fmt.Println("════════════════════════════════════════════════════════════\n")

	fmt.Println(`
STRUCT DEFINITION GUIDELINES:

✓ Use descriptive field names (not x, y, z unless context is clear)
✓ Export fields that need to be accessible outside package (Capitalize)
✓ Keep unexported fields for internal use only (lowercase)
✓ Group related fields together
✓ Use composition over inheritance
✓ Embed structs for code reuse
✓ Use pointers for large structs (avoid copying)

WHEN TO USE STRUCTS:

✓ Grouping related data together
✓ Creating custom data types
✓ Building complex objects from simpler ones (composition)
✓ Organizing code into logical units
✓ Preparing data for JSON marshaling
✓ Creating domain models

COMPOSITION VS INHERITANCE:

┌────────────────┬──────────────────────────────┬──────────────────────────┐
│ Aspect         │ Inheritance (Traditional)    │ Composition (Go)         │
├────────────────┼──────────────────────────────┼──────────────────────────┤
│ Relationship   │ IS-A (Car IS-A Vehicle)      │ HAS-A (Car HAS-A Engine) │
│ Flexibility    │ Rigid hierarchy              │ Flexible building blocks │
│ Complexity     │ Can be complex               │ Simple and clear         │
│ Reusability    │ Limited (tied to hierarchy)  │ High (use anywhere)      │
│ Conflicts      │ Diamond problem, ambiguity   │ None, explicit           │
│ Go Support     │ Not directly supported       │ Primary approach         │
└────────────────┴──────────────────────────────┴──────────────────────────┘

FIELD PROMOTION RULES:

1. Only works with EMBEDDED structs (no field name)
2. Promoted fields are ACCESSIBLE but not SHADOWED
3. If conflict exists, explicit access is required
4. Saves typing for commonly accessed fields
5. Makes code more readable and cleaner

MEMORY CONSIDERATIONS:

- Structs are VALUE types (not references)
- Assignment copies entire struct
- Use pointers for large structs
- Embedded structs increase memory footprint (union-like, not inheritance)
- Pointers reduce copying overhead

NEXT STEP:

The next logical progression:
Basic Structs (storing data) → Composition (organizing data) → Methods (adding behavior)

After methods, you'll learn about INTERFACES, which is where Go becomes truly powerful
by allowing different types to interact based on their behavior, not their inheritance chain.
`)
}
