package main

import "fmt"

/*
Topic 66: STRUCT EMBEDDING - Composition Through Embedding

═══════════════════════════════════════════════════════════════════════════════

WHAT IS STRUCT EMBEDDING?

Struct embedding is a mechanism in Go that allows one struct to include another
struct type without giving it a field name.

Key Features:
- One struct is nested inside another (ANONYMOUS field)
- Fields from inner struct are PROMOTED to outer struct
- Methods from inner struct are also PROMOTED
- Provides code reuse and cleaner syntax
- Go's approach to "composition" and "reusability"

DIFFERENCE FROM TRADITIONAL FIELD:

Named Field (Regular):
  type employee struct {
      info person     // Named field - must use emp.info.name
  }

Anonymous Field (Embedding):
  type employee struct {
      person          // No name - can use emp.name directly
  }

═══════════════════════════════════════════════════════════════════════════════

KEY CONCEPT: FIELD PROMOTION

When you embed a struct without a name, its fields are "promoted" to the
outer struct. You can access them directly without typing the intermediate
struct name.

EXAMPLE:
  type person struct {
      name string
      age  int
  }

  type employee struct {
      person              // EMBEDDED (anonymous)
      employeeID string
  }

  emp := employee{...}
  fmt.Println(emp.name)   ✓ Works! (promoted field)
  fmt.Println(emp.age)    ✓ Works! (promoted field)

This saves typing and makes the code cleaner!

═══════════════════════════════════════════════════════════════════════════════

PART 1: BASIC STRUCT EMBEDDING & FIELD PROMOTION

═════════════════════════════════════════════════════════════════════════════

THE INNER STRUCT (Base Component):
*/

type person struct {
	name string
	age  int
}

// THE OUTER STRUCT (Composite):
// Embeds 'person' as an ANONYMOUS field (no name, just the type)
type employee struct {
	person     // EMBEDDED (anonymous field)
	employeeID string
	salary     float64
}

/*
═════════════════════════════════════════════════════════════════════════════

PART 2: METHOD INHERITANCE (Promoted Methods)

Methods defined on the embedded struct are also promoted to the outer struct.

═════════════════════════════════════════════════════════════════════════════
*/

// Method on person struct
func (p person) introduce() string {
	return fmt.Sprintf("I am %s and I am %d years old.", p.name, p.age)
}

func (p person) getAge() int {
	return p.age
}

/*
═════════════════════════════════════════════════════════════════════════════

PART 3: METHOD OVERRIDING

If the outer struct defines a method with the SAME NAME as the inner struct,
the outer struct's method takes precedence. But you can still call the inner
struct's method explicitly.

═════════════════════════════════════════════════════════════════════════════
*/

// Override the introduce method for employee
func (e employee) introduce() string {
	return fmt.Sprintf("I am %s (ID: %s) and I earn $%.2f per year.",
		e.name, e.employeeID, e.salary)
}

// New method specific to employee
func (e employee) getEmployeeInfo() string {
	return fmt.Sprintf("Name: %s, Age: %d, ID: %s, Salary: $%.2f",
		e.name, e.age, e.employeeID, e.salary)
}

/*
═════════════════════════════════════════════════════════════════════════════

PART 4: NAMED FIELDS vs. ANONYMOUS FIELDS

IMPORTANT: If you give the embedded struct a NAME, field promotion STOPS!

Named Field Example:
  type employee struct {
      info person       // Named field - NO promotion
      employeeID string
  }
  emp.info.name  ✓ Must use full path
  emp.name       ✗ ERROR - field is NOT promoted

═════════════════════════════════════════════════════════════════════════════
*/

// Example of NAMED field (field promotion disabled)
type employee2 struct {
	personInfo person // Named field - NOT anonymous
	employeeID string
	department string
}

/*
═════════════════════════════════════════════════════════════════════════════

PART 5: MULTIPLE EMBEDDING

A struct can embed multiple other structs. This creates a powerful composition.

═════════════════════════════════════════════════════════════════════════════
*/

type address struct {
	street string
	city   string
	state  string
	zip    string
}

// Manager embeds both person and address
type manager struct {
	person
	address
	managedTeamSize int
}

/*
═════════════════════════════════════════════════════════════════════════════

PART 6: SHADOWING (Name Conflicts)

If both the inner and outer struct have fields with the same name, accessing
that name on the outer struct will access the OUTER struct's field.
This is called "shadowing."

To access the shadowed inner field, use the explicit path.

═════════════════════════════════════════════════════════════════════════════
*/

type person2 struct {
	name string
	age  int
	city string // Same field as below
}

type employee3 struct {
	person2
	city string // This SHADOWS person2.city
}

/*
═════════════════════════════════════════════════════════════════════════════

PART 7: EMBEDDING INTERFACES (Advanced)

Go allows embedding interfaces inside structs. This is a powerful pattern
for creating flexible, loosely-coupled designs.

═════════════════════════════════════════════════════════════════════════════
*/

// Define an interface
type speaker interface {
	speak() string
}

// A type implementing the interface
type dog struct {
	name string
}

func (d dog) speak() string {
	return fmt.Sprintf("%s barks: Woof!", d.name)
}

// Embed the interface in a struct
type pack struct {
	speaker  // Embedded interface - satisfies the interface implicitly
	packSize int
}

/*
═════════════════════════════════════════════════════════════════════════════

PART 8: PRACTICAL PATTERNS

═════════════════════════════════════════════════════════════════════════════
*/

func main() {
	fmt.Println("════════════════════════════════════════════════════════════")
	fmt.Println("PART 1: BASIC EMBEDDING & FIELD PROMOTION")
	fmt.Println("════════════════════════════════════════════════════════════\n")

	// Create an employee
	emp := employee{
		person: person{
			name: "John Doe",
			age:  30,
		},
		employeeID: "E001",
		salary:     75000.0,
	}

	fmt.Println("Employee Information (using promoted fields):")
	fmt.Printf("Name: %s\n", emp.name)     // Promoted field
	fmt.Printf("Age: %d\n", emp.age)       // Promoted field
	fmt.Printf("ID: %s\n", emp.employeeID) // Own field
	fmt.Printf("Salary: $%.2f\n\n", emp.salary)

	// Can also access through explicit path
	fmt.Println("Same data (using explicit path):")
	fmt.Printf("Name: %s\n", emp.person.name) // Explicit access
	fmt.Printf("Age: %d\n\n", emp.person.age) // Explicit access

	// ─────────────────────────────────────────────────────────────────────────

	fmt.Println("════════════════════════════════════════════════════════════")
	fmt.Println("PART 2: METHOD INHERITANCE (Promoted Methods)")
	fmt.Println("════════════════════════════════════════════════════════════\n")

	fmt.Println("Calling promoted method from person:")
	fmt.Println(emp.getAge())              // Promoted method
	fmt.Println(emp.person.getAge(), "\n") // Explicit call

	// ─────────────────────────────────────────────────────────────────────────

	fmt.Println("════════════════════════════════════════════════════════════")
	fmt.Println("PART 3: METHOD OVERRIDING")
	fmt.Println("════════════════════════════════════════════════════════════\n")

	fmt.Println("Original person method:")
	p := person{name: "Jane Smith", age: 28}
	fmt.Println(p.introduce())

	fmt.Println("\nOverridden employee method (same name, different output):")
	fmt.Println(emp.introduce())

	fmt.Println("\nCalling inner method explicitly:")
	fmt.Println(emp.person.introduce())

	fmt.Println("\nEmployee-specific method:")
	fmt.Println(emp.getEmployeeInfo(), "\n")

	// ─────────────────────────────────────────────────────────────────────────

	fmt.Println("════════════════════════════════════════════════════════════")
	fmt.Println("PART 4: NAMED vs. ANONYMOUS FIELDS")
	fmt.Println("════════════════════════════════════════════════════════════\n")

	fmt.Println("ANONYMOUS FIELD (with promotion):")
	emp1 := employee{
		person:     person{name: "Alice", age: 25},
		employeeID: "E002",
		salary:     60000.0,
	}
	fmt.Printf("emp1.name = %s (promoted)\n", emp1.name)
	fmt.Printf("emp1.person.name = %s (explicit)\n", emp1.person.name)

	fmt.Println("\nNAMED FIELD (without promotion):")
	emp2 := employee2{
		personInfo: person{name: "Bob", age: 35},
		employeeID: "E003",
		department: "Engineering",
	}
	// fmt.Printf("emp2.name = %s\n", emp2.name) // ❌ ERROR - not promoted
	fmt.Printf("emp2.personInfo.name = %s (must use full path)\n", emp2.personInfo.name)
	fmt.Printf("emp2.personInfo.age = %d\n\n", emp2.personInfo.age)

	// ─────────────────────────────────────────────────────────────────────────

	fmt.Println("════════════════════════════════════════════════════════════")
	fmt.Println("PART 5: MULTIPLE EMBEDDING")
	fmt.Println("════════════════════════════════════════════════════════════\n")

	mgr := manager{
		person: person{
			name: "Carol Johnson",
			age:  45,
		},
		address: address{
			street: "123 Main St",
			city:   "San Francisco",
			state:  "CA",
			zip:    "94102",
		},
		managedTeamSize: 12,
	}

	fmt.Println("Manager Information (multiple embedded structs):")
	fmt.Printf("Name: %s\n", mgr.name)                   // From person
	fmt.Printf("Age: %d\n", mgr.age)                     // From person
	fmt.Printf("City: %s\n", mgr.city)                   // From address
	fmt.Printf("State: %s\n", mgr.state)                 // From address
	fmt.Printf("Team Size: %d\n\n", mgr.managedTeamSize) // Own field

	// ─────────────────────────────────────────────────────────────────────────

	fmt.Println("════════════════════════════════════════════════════════════")
	fmt.Println("PART 6: FIELD SHADOWING")
	fmt.Println("════════════════════════════════════════════════════════════\n")

	emp3 := employee3{
		person2: person2{
			name: "David",
			age:  40,
			city: "New York", // Inner struct's city
		},
		city: "Los Angeles", // Outer struct's city (shadows inner)
	}

	fmt.Println("When both inner and outer have same field name:")
	fmt.Printf("emp3.city = %s (outer struct's field - SHADOWED)\n", emp3.city)
	fmt.Printf("emp3.person2.city = %s (inner struct's field - explicit)\n\n", emp3.person2.city)

	// ─────────────────────────────────────────────────────────────────────────

	fmt.Println("════════════════════════════════════════════════════════════")
	fmt.Println("PART 7: EMBEDDING INTERFACES")
	fmt.Println("════════════════════════════════════════════════════════════\n")

	d := dog{name: "Buddy"}
	myPack := pack{
		speaker:  d, // Embedded interface holds a dog
		packSize: 5,
	}

	fmt.Println("Pack with embedded speaker interface:")
	fmt.Println(myPack.speak())         // Promoted method from interface
	fmt.Println(myPack.speaker.speak()) // Explicit call
	fmt.Printf("Pack Size: %d\n\n", myPack.packSize)

	// ─────────────────────────────────────────────────────────────────────────

	fmt.Println("════════════════════════════════════════════════════════════")
	fmt.Println("PART 8: BEST PRACTICES & KEY CONCEPTS")
	fmt.Println("════════════════════════════════════════════════════════════\n")

	fmt.Println(`
STRUCT EMBEDDING SUMMARY:

┌─────────────────────┬──────────────────────────┬──────────────────────┐
│ Aspect              │ Anonymous Field (Embed)  │ Named Field          │
├─────────────────────┼──────────────────────────┼──────────────────────┤
│ Field Promotion     │ YES ✓                    │ NO ✗                 │
│ Method Promotion    │ YES ✓                    │ NO ✗                 │
│ Syntax              │ type X struct { Y }      │ type X struct {y Y}  │
│ Access Fields       │ x.field                  │ x.y.field            │
│ Use Case            │ Inheritance-like reuse   │ Regular composition  │
└─────────────────────┴──────────────────────────┴──────────────────────┘

WHEN TO USE STRUCT EMBEDDING:

✓ Reusing fields and methods from another type
✓ Creating "inheritance-like" relationships (Go's style)
✓ Building complex types from simpler components
✓ Keeping code DRY (Don't Repeat Yourself)
✓ Creating flexible, composable designs

WHEN NOT TO USE:

✗ If you need to control access (use named fields)
✗ If embedding creates ambiguity
✗ If the relationship doesn't truly fit "inheritance"

KEY RULES:

1. FIELD PROMOTION works only with ANONYMOUS fields
2. METHODS are promoted from embedded types
3. SHADOWING occurs when inner and outer have same field/method names
4. OVERRIDING is achieved by defining a method with the same name
5. EXPLICIT ACCESS always works: outer.inner.field or outer.inner.method()

WHY EMBEDDING OVER INHERITANCE?

Go doesn't have traditional inheritance. Instead, it provides embedding
because:

✓ Simpler: No complex hierarchies to manage
✓ Flexible: A type can embed multiple types
✓ Clear: "X has Y" is clearer than "X is Y"
✓ Composable: Easy to add/remove embedded types
✓ Testable: Easier to mock and test

ADVANCED PATTERN: EMBEDDING INTERFACES

Embedding an interface in a struct is powerful:

  type myType struct {
      reader io.Reader       // Embedded interface
  }

This allows myType to act as a Reader without explicitly implementing it!
Very useful for wrappers and decorators.

═════════════════════════════════════════════════════════════════════════════

NEXT STEPS:

Now that you understand embedding, you can build more sophisticated
composition patterns. The next topics explore:
- Generics (parameterized types)
- Custom Errors
- Advanced interface patterns
`)
}
