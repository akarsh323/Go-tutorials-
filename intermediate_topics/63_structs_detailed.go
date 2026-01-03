package main

import "fmt"

/*
Topic 63: STRUCTS & COMPOSITION

CONCEPT: Structs group related data. Composition (embedding) lets you build
complex types from simpler ones - the Go way instead of inheritance.

STRUCTS:
- Group related fields together
- Create custom data types
- Fields can be exported (Capitalized) or unexported (lowercase)
- Immutable unless using pointers

COMPOSITION (THE LEGO APPROACH):
- Embed structs inside other structs
- "Car HAS an Engine" (not "Car IS A Engine")
- More flexible than inheritance
- Reusable components

FIELD PROMOTION:
- When embedding struct, fields are "promoted" to outer struct
- Access: myCar.Horsepower (promoted) or myCar.Engine.Horsepower (explicit)
- Saves typing, cleaner code

BENEFITS:
✓ Organization: Keeps related code together
✓ Reusability: Embed same struct in multiple types
✓ Clarity: "HAS-A" is clearer than "IS-A"
✓ No hierarchy problems: Can combine components freely
*/

// BASIC STRUCTS
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

// COMPOSITION: EMBEDDED STRUCTS
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

// COMPOSED STRUCT
type Car struct {
	Make   string
	Model  string
	Year   int
	Engine // EMBEDDED (no field name)
	Wheels // EMBEDDED (no field name)
}

// MULTIPLE COMPOSITION
type Employee struct {
	Person     // EMBEDDED
	EmployeeID int
	Department string
	Salary     float64
}

func main() {
	fmt.Println("═══════════════════════════════════════════════════════════")
	fmt.Println("PART 1: BASIC STRUCT DECLARATION")
	fmt.Println("═══════════════════════════════════════════════════════════\n")

	// Zero value (all fields get default values)
	var p1 Person
	fmt.Println("Zero value:", p1)

	// Named field initialization
	p2 := Person{
		Name:  "Alice",
		Age:   30,
		Email: "alice@example.com",
	}
	fmt.Printf("Named fields: %+v\n", p2)

	// Positional initialization (order matters!)
	p3 := Person{"Bob", 25, "bob@example.com"}
	fmt.Printf("Positional: %+v\n\n", p3)

	// ─────────────────────────────────────────────────────────────────────────

	fmt.Println("═══════════════════════════════════════════════════════════")
	fmt.Println("PART 2: ACCESSING & MODIFYING FIELDS")
	fmt.Println("═══════════════════════════════════════════════════════════\n")

	person := Person{Name: "Charlie", Age: 35, Email: "charlie@example.com"}
	fmt.Printf("Original: Name=%s, Age=%d\n", person.Name, person.Age)

	// Modify fields
	person.Age = 36
	person.Email = "charlie.new@example.com"
	fmt.Printf("Modified: Name=%s, Age=%d\n", person.Name, person.Age)
	fmt.Printf("          Email=%s\n\n", person.Email)

	// ─────────────────────────────────────────────────────────────────────────

	fmt.Println("═══════════════════════════════════════════════════════════")
	fmt.Println("PART 3: POINTERS TO STRUCTS")
	fmt.Println("═══════════════════════════════════════════════════════════\n")

	original := Person{Name: "David", Age: 40, Email: "david@example.com"}
	fmt.Println("Original:", original)

	// Create pointer
	ptr := &original
	fmt.Printf("Pointer: %p\n", ptr)

	// Go auto-dereferences for struct fields!
	ptr.Name = "Dave"
	ptr.Age = 41

	fmt.Println("After modification through pointer:")
	fmt.Println("Original struct changed:", original)
	fmt.Println("(Pointer points to original, changes affect it)\n")

	// ─────────────────────────────────────────────────────────────────────────

	fmt.Println("═══════════════════════════════════════════════════════════")
	fmt.Println("PART 4: COMPOSITION - THE LEGO APPROACH")
	fmt.Println("═══════════════════════════════════════════════════════════\n")

	fmt.Println("Creating a Car with embedded Engine and Wheels:\n")

	myCar := Car{
		Make:  "Tesla",
		Model: "Model S",
		Year:  2024,
		Engine: Engine{
			Horsepower: 1020,
			Type:       "Electric",
			Cylinders:  0,
		},
		Wheels: Wheels{
			Count: 4,
			Brand: "Michelin",
			Size:  "19 inch",
		},
	}

	fmt.Printf("Car: %d %s %s\n", myCar.Year, myCar.Make, myCar.Model)
	fmt.Printf("Engine: %d HP, %s, %d cylinders\n", myCar.Engine.Horsepower, myCar.Engine.Type, myCar.Engine.Cylinders)
	fmt.Printf("Wheels: %d wheels, %s, %s\n\n", myCar.Wheels.Count, myCar.Wheels.Brand, myCar.Wheels.Size)

	// ─────────────────────────────────────────────────────────────────────────

	fmt.Println("═══════════════════════════════════════════════════════════")
	fmt.Println("PART 5: FIELD PROMOTION (THE SHORTCUT)")
	fmt.Println("═══════════════════════════════════════════════════════════\n")

	fmt.Println("Explicit way (through embedded struct):")
	fmt.Printf("myCar.Engine.Horsepower: %d HP\n", myCar.Engine.Horsepower)
	fmt.Printf("myCar.Wheels.Brand:      %s\n\n", myCar.Wheels.Brand)

	fmt.Println("Promoted way (Go's convenience feature):")
	fmt.Printf("myCar.Horsepower (promoted): %d HP\n", myCar.Horsepower)
	fmt.Printf("myCar.Brand (promoted):      %s\n\n", myCar.Brand)

	fmt.Println("✓ Both ways work! Promoted fields are cleaner.")
	fmt.Println("✓ Promoted: direct access without nesting\n")

	// ─────────────────────────────────────────────────────────────────────────

	fmt.Println("═══════════════════════════════════════════════════════════")
	fmt.Println("PART 6: MODIFYING PROMOTED FIELDS")
	fmt.Println("═══════════════════════════════════════════════════════════\n")

	fmt.Printf("Original Horsepower: %d\n", myCar.Horsepower)
	myCar.Horsepower = 900
	fmt.Printf("Modified Horsepower: %d\n", myCar.Horsepower)
	fmt.Printf("Verified (explicit):  %d\n", myCar.Engine.Horsepower)
	fmt.Println("(Both refer to the same field!)\n")

	// ─────────────────────────────────────────────────────────────────────────

	fmt.Println("═══════════════════════════════════════════════════════════")
	fmt.Println("PART 7: MULTIPLE COMPOSITION (EMPLOYEE)")
	fmt.Println("═══════════════════════════════════════════════════════════\n")

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
	fmt.Printf("Name (promoted):       %s\n", emp.Name)
	fmt.Printf("Age (promoted):        %d\n", emp.Age)
	fmt.Printf("Email (promoted):      %s\n", emp.Email)
	fmt.Printf("Employee ID:           %d\n", emp.EmployeeID)
	fmt.Printf("Department:            %s\n", emp.Department)
	fmt.Printf("Salary:                $%.2f\n\n", emp.Salary)

	// ─────────────────────────────────────────────────────────────────────────

	fmt.Println("═══════════════════════════════════════════════════════════")
	fmt.Println("COMPOSITION VS INHERITANCE COMPARISON")
	fmt.Println("═══════════════════════════════════════════════════════════\n")

	fmt.Println(`
INHERITANCE (Traditional):
- "IS A" relationship (Car IS A Vehicle)
- Rigid hierarchy
- Diamond problem (multiple inheritance conflicts)
- All methods from parent inherited
- Difficult to customize

COMPOSITION (Go's Approach):
- "HAS A" relationship (Car HAS AN Engine)
- Flexible and modular
- Mix and match components freely
- Only include what you need
- Easy to understand and maintain

WHY COMPOSITION IS BETTER:

1. FLEXIBILITY:
   Inheritance = fixed hierarchy
   Composition = mix any components

2. REUSABILITY:
   Inheritance = tight coupling to parent
   Composition = Engine used in Car, Truck, Motorcycle

3. SIMPLICITY:
   Inheritance = complex rules and conflicts
   Composition = clear "HAS A" structure

4. EXTENSIBILITY:
   Inheritance = override methods (complexity)
   Composition = add new components (simplicity)

REAL EXAMPLE:
  Car HAS Engine + HAS Wheels
  Truck HAS Engine + HAS Wheels + HAS Bed
  Motorcycle HAS Engine + HAS Wheels

  No need for inheritance hierarchy!
  Just compose what you need!
	`)

	// ─────────────────────────────────────────────────────────────────────────

	fmt.Println("═══════════════════════════════════════════════════════════")
	fmt.Println("KEY CONCEPTS & BEST PRACTICES")
	fmt.Println("═══════════════════════════════════════════════════════════\n")

	fmt.Println(`
STRUCT BASICS:
✓ Group related fields together
✓ Exported (Capitalized) = public, Unexported (lowercase) = private
✓ Immutable unless using pointers
✓ Multiple initialization methods (zero, named, positional)

COMPOSITION:
✓ Embed struct inside another (no field name, just type)
✓ "Car HAS Engine" not "Car IS Engine"
✓ More flexible than inheritance
✓ Highly reusable components

FIELD PROMOTION:
✓ Promoted fields accessible directly on outer struct
✓ Explicit access also always works
✓ Saves typing and improves readability
✓ Go automatically finds promoted fields

POINTERS TO STRUCTS:
✓ Use & to get address of struct
✓ Go auto-dereferences for field access (ptr.Field works)
✓ Necessary for large structs (avoid copying)
✓ Required if you want to modify struct

BEST PRACTICES:
✓ Use composition over inheritance
✓ Keep related data in structs
✓ Export only fields that need to be public
✓ Use pointers for large structs
✓ Embed structs for code reuse
✓ Document which fields are important
✓ Initialize all fields (or use zero values)
✓ Group related fields together logically

NAMING CONVENTIONS:
✓ Type names: Capitalized (Person, Car)
✓ Exported fields: Capitalized (Name, Age)
✓ Unexported fields: lowercase (privateData)
✓ Constructors: NewTypeName() (NewCar())
✓ Methods: normal function names

The beauty of composition: build complex systems from simple, reusable parts!
	`)
}
