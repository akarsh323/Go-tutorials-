package main

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strings"
)

/*
═══════════════════════════════════════════════════════════════════════════════
                         STRUCT TAGS IN GO
═══════════════════════════════════════════════════════════════════════════════

Struct tags are string literals placed between backticks (`...`) after a field type.
They act as metadata for external packages (like encoding/json).

KEY PROBLEM SOLVED:
  • Go uses PascalCase (FirstName) for public fields.
  • JSON APIs often use snake_case (first_name) or camelCase (firstName).
  • Some fields (like passwords) exist in memory but should never leave the server.
  • Some fields (like empty strings) waste bandwidth and should be omitted.

SOLUTION: Struct tags provide a mapping layer between your Go code and external
data formats.

═══════════════════════════════════════════════════════════════════════════════
*/

/*
━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━
  SECTION 1: THE BASICS (RENAME, OMIT, IGNORE)
━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━
*/

// Person represents the struct described in the lecture
type Person struct {
	// 1. Renaming: Maps Go's "FirstName" to JSON's "first_name"
	FirstName string `json:"first_name"`

	// 2. Omit Empty: Maps to "last_name", but excludes field if value is ""
	//    Useful for optional fields.
	LastName string `json:"last_name,omitempty"`

	// 3. Ignoring: "Age" will NEVER appear in the JSON output
	//    Useful for internal calculations or sensitive data.
	Age int `json:"-"`

	// Another example of ignoring sensitive data
	Password string `json:"-"`
}

func Example1_BasicTags() {
	fmt.Println("\n" + strings.Repeat("═", 80))
	fmt.Println("EXAMPLE 1: Renaming, Omitting, and Ignoring")
	fmt.Println(strings.Repeat("═", 80) + "\n")

	// ─────────────────────────────────────────────────────────────────────────
	// Scenario A: All fields populated
	// ─────────────────────────────────────────────────────────────────────────
	fmt.Println("📌 Scenario A: All fields have data")
	fmt.Println(strings.Repeat("─", 80))

	p1 := Person{
		FirstName: "Jane",
		LastName:  "Doe",
		Age:       50,
		Password:  "secret123",
	}

	fmt.Printf("Go Struct: %+v\n", p1)

	jsonData1, _ := json.Marshal(p1)
	fmt.Printf("JSON Out:  %s\n", string(jsonData1))
	fmt.Println("Analysis:  'last_name' is present. 'Age' and 'Password' are hidden.")
	fmt.Println()

	// ─────────────────────────────────────────────────────────────────────────
	// Scenario B: Empty fields (Zero Values)
	// ─────────────────────────────────────────────────────────────────────────
	fmt.Println("📌 Scenario B: Empty fields (Zero Values)")
	fmt.Println(strings.Repeat("─", 80))

	p2 := Person{
		FirstName: "John",
		LastName:  "", // Empty string (Zero Value)
		Age:       25,
		Password:  "password",
	}

	fmt.Printf("Go Struct: %+v\n", p2)

	jsonData2, _ := json.Marshal(p2)
	fmt.Printf("JSON Out:  %s\n", string(jsonData2))
	fmt.Println("Analysis:  'last_name' is COMPLETELY GONE because of omitempty.")
	fmt.Println("           'Age' is still hidden because of hyphen (-).")
}

/*
━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━
  SECTION 2: THE "ZERO VALUE" NUANCE
━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━

WARNING: "omitempty" hides fields when they hold the "Zero Value".
  • String zero value: ""
  • Int zero value:    0
  • Bool zero value:   false

This can cause bugs. If a user sets a setting to "false" or an inventory
count to "0", the field might disappear when you actually wanted to send "0".
━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━
*/

type Wallet struct {
	// If Money is 0, this field disappears. That might be confusing for the frontend.
	Money int `json:"money,omitempty"`

	// If IsActive is false, this field disappears.
	IsActive bool `json:"is_active,omitempty"`
}

func Example2_ZeroValueTrap() {
	fmt.Println("\n" + strings.Repeat("═", 80))
	fmt.Println("EXAMPLE 2: The 'Zero Value' Trap")
	fmt.Println(strings.Repeat("═", 80) + "\n")

	w := Wallet{
		Money:    0,     // User is broke, but the value is valid
		IsActive: false, // User account is explicitly inactive
	}

	fmt.Println("📌 Input: Money=0, IsActive=false")
	jsonData, _ := json.Marshal(w)

	// Result is empty brackets {} because 0 and false are considered "empty"
	fmt.Printf("JSON Out: %s\n", string(jsonData))
	fmt.Println("⚠️  WARNING: The JSON is empty! The frontend won't know the money is 0.")
	fmt.Println("            It might assume the value is missing or undefined.")
}

/*
━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━
  SECTION 3: MULTIPLE TAGS (JSON + DB + XML)
━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━

Struct tags are just strings. Go's reflection API allows you to put multiple
instructions in one tag line, separated by spaces.
━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━
*/

type MultiSystemData struct {
	// This single field works for JSON APIs, SQL Databases, and XML files
	Identifier string `json:"user_id" db:"id_col" xml:"IdNode"`
}

func Example3_MultipleTags() {
	fmt.Println("\n" + strings.Repeat("═", 80))
	fmt.Println("EXAMPLE 3: Multiple Tags (Reading Metadata)")
	fmt.Println(strings.Repeat("═", 80) + "\n")

	data := MultiSystemData{Identifier: "A100"}

	// 1. Prove JSON works
	jsonBytes, _ := json.Marshal(data)
	fmt.Printf("JSON Usage: %s\n", string(jsonBytes))

	// 2. Use Reflection to see the other tags (how DB libraries see them)
	// Get the type of the struct
	t := reflect.TypeOf(data)

	// Get the specific field "Identifier"
	field, _ := t.FieldByName("Identifier")

	fmt.Println("\n📌 Reflection Analysis (How libraries read tags):")
	fmt.Println(strings.Repeat("─", 80))
	fmt.Printf("Field Name: %s\n", field.Name)
	fmt.Printf("Full Tag:   %s\n", field.Tag)
	fmt.Printf("JSON key:   %s\n", field.Tag.Get("json"))
	fmt.Printf("DB column:  %s\n", field.Tag.Get("db"))
	fmt.Printf("XML node:   %s\n", field.Tag.Get("xml"))
}

/*
━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━
  SECTION 4: COMPLEX NESTED STRUCTURES
━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━
  Real-world JSON is rarely flat. It contains arrays and nested objects.
  Struct tags apply recursively.
━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━
*/

type Item struct {
	Name  string  `json:"item_name"`
	Price float64 `json:"cost"`
}

type ShoppingCart struct {
	User  string  `json:"user"`
	Items []Item  `json:"cart_items,omitempty"` // Omit if slice is nil/empty
	Total float64 `json:"-"`                    // Ignore, we calculate this on the fly
}

func Example4_NestedStructures() {
	fmt.Println("\n" + strings.Repeat("═", 80))
	fmt.Println("EXAMPLE 4: Nested Structures (Shopping Cart)")
	fmt.Println(strings.Repeat("═", 80) + "\n")

	cart := ShoppingCart{
		User: "akarsh_blr",
		Items: []Item{
			{Name: "MacBook Pro", Price: 2000.00},
			{Name: "Go Language Book", Price: 45.50},
		},
		Total: 2045.50, // This will be ignored
	}

	jsonData, _ := json.MarshalIndent(cart, "", "  ")
	fmt.Println("📌 JSON Output (Pretty Printed):")
	fmt.Println(string(jsonData))
}

/*
═══════════════════════════════════════════════════════════════════════════════
                        QUICK REFERENCE TABLE
═══════════════════════════════════════════════════════════════════════════════

TAG SYNTAX           | EFFECT                                | EXAMPLE
─────────────────────|───────────────────────────────────────|───────────────────
`json:"name"`        | Renames field to "name"               | Field `json:"id"`
`json:"name,omitempty"`| Renames + hides if zero value (0,"")| Field `json:"id,omitempty"`
`json:"-"`           | Field is ALWAYS hidden                | Password `json:"-"`
`json:"-,"`          | Field is literally named "-"          | rare case
`json:",string"`     | Encodes int/bool/float as string      | ID int `json:",string"`

BEST PRACTICES:
  1. Always use snake_case for JSON keys (`first_name`) per industry standard.
  2. Use `omitempty` for optional fields to save bandwidth.
  3. Be careful using `omitempty` with Booleans (false) and Ints (0).
     If 0 is a valid value, do NOT use omitempty.

═══════════════════════════════════════════════════════════════════════════════
*/

// Main Entry Point
func DemoStructTags() {
	fmt.Println("\n" + strings.Repeat("═", 80))
	fmt.Println("STRUCT TAGS - COMPLETE GUIDE")
	fmt.Println(strings.Repeat("═", 80))

	Example1_BasicTags()
	Example2_ZeroValueTrap()
	Example3_MultipleTags()
	Example4_NestedStructures()

	fmt.Println("\n" + strings.Repeat("═", 80))
	fmt.Println("KEY TAKEAWAY:")
	fmt.Println("  Struct tags are the bridge between Go's strict typing")
	fmt.Println("  and the flexible world of JSON/Databases.")
	fmt.Println(strings.Repeat("═", 80) + "\n")
}

func main() {
	DemoStructTags()
}
