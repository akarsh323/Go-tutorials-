package intermediate

import (
	"bytes"
	"fmt"
	"os"
	"strings"
	"text/template"
)

// Topic 72: text_templates
// Deep dive into Why Go Templates look the way they do.
// Covers:
// Part 1: Basic templates with dot notation
// Part 2: Conditionals - Function-call syntax (Why "ge" instead of ">=")
// Part 3: Loops and the Shape-Shifting Dot
// Part 4: CLI Architecture - Vending Machine Pattern

func main() {
	fmt.Println("=== 72 TEXT TEMPLATES: Deep Dive ===\n")

	// ============================================================
	// PART 1: Basics - Template Variables
	// ============================================================
	fmt.Println("PART 1: Basics - Template Variables")
	fmt.Println("=====================================\n")

	// Basic template with field access
	tmpl1 := template.Must(template.New("greet").Parse("Hello {{.Name}}!\n"))
	tmpl1.Execute(os.Stdout, map[string]string{"Name": "Gopher"})

	// Template with multiple fields
	tmpl2 := template.Must(template.New("person").Parse(
		"Name: {{.Name}}, Age: {{.Age}}, City: {{.City}}\n"))
	person := map[string]interface{}{
		"Name": "Alice",
		"Age":  30,
		"City": "NYC",
	}
	tmpl2.Execute(os.Stdout, person)

	// ============================================================
	// PART 2: Conditionals - Why "ge .Age 18" instead of ".Age >= 18"?
	// ============================================================
	fmt.Println("\n\nPART 2: Conditionals - Function-Call Syntax")
	fmt.Println("==============================================\n")

	fmt.Println("📌 The Analogy: Function Calls")
	fmt.Println("In most languages: if Age >= 18 { ... }")
	fmt.Println("In Go Templates:   {{if ge .Age 18}} ... {{end}}")
	fmt.Println()
	fmt.Println("Think of 'ge' not as a SYMBOL, but as a FUNCTION NAME.")
	fmt.Println("ge(Age, 18) is how the template engine reads it.\n")

	// Visual comparison table
	fmt.Println("📊 Code Translation Table:")
	fmt.Println("Normal Logic       | Go Template Logic | Meaning")
	fmt.Println("================== | ================= | ===================")
	fmt.Println("if Age == 18       | {{if eq .Age 18}} | eq = Equal")
	fmt.Println("if Age != 18       | {{if ne .Age 18}} | ne = Not Equal")
	fmt.Println("if Age < 18        | {{if lt .Age 18}} | lt = Less Than")
	fmt.Println("if Age <= 18       | {{if le .Age 18}} | le = Less or Equal")
	fmt.Println("if Age > 18        | {{if gt .Age 18}} | gt = Greater Than")
	fmt.Println("if Age >= 18       | {{if ge .Age 18}} | ge = Greater or Equal")
	fmt.Println()

	// Live example of conditionals
	fmt.Println("🔄 Live Example of Conditionals:\n")
	tmpl3 := template.Must(template.New("age").Parse(
		"{{.Name}} is {{if ge .Age 18}}an adult{{else}}a minor{{end}}\n"))
	tmpl3.Execute(os.Stdout, map[string]interface{}{
		"Name": "Bob",
		"Age":  25,
	})
	tmpl3.Execute(os.Stdout, map[string]interface{}{
		"Name": "Charlie",
		"Age":  10,
	})

	fmt.Println("\nHow it works:")
	fmt.Println("  - ge is a FUNCTION: ge(25, 18) → true")
	fmt.Println("  - ge is a FUNCTION: ge(10, 18) → false")
	fmt.Println("  - Template engine reads left-to-right")
	fmt.Println("  - Grabs .Age, grabs 18, calls ge() with those arguments\n")

	// ============================================================
	// PART 3: Loops - The Shape-Shifting Dot (The Spotlight Analogy)
	// ============================================================
	fmt.Println("\n\nPART 3: Loops - The Shape-Shifting Dot")
	fmt.Println("=========================================\n")

	fmt.Println("🎯 The Spotlight Analogy:")
	fmt.Println("Imagine the Dot '.' is a spotlight that shines on your data.\n")

	fmt.Println("BEFORE the loop:")
	fmt.Println("  The spotlight shines on the ENTIRE MAP.")
	fmt.Println("  You can access: {{.Items}}, {{.User}}, {{.Details}}")
	fmt.Println()

	fmt.Println("INSIDE the loop:")
	fmt.Println("  The spotlight MOVES!")
	fmt.Println("  Now it shines ONLY on the CURRENT ITEM.")
	fmt.Println("  {{.}} refers to the individual item, NOT the map.")
	fmt.Println()

	fmt.Println("WHY THIS MATTERS:")
	fmt.Println("  If you tried {{.Items}} INSIDE the loop, it would crash!")
	fmt.Println("  Why? Because 'apple' (the current item) has no 'Items' field.")
	fmt.Println("  The spotlight has moved away from the outer map.\n")

	// Visual diagram
	fmt.Println("📊 Visual Diagram of the Spotlight:\n")
	fmt.Println(`
Data Structure:
  {
    "Name": "Alice",
    "Items": ["Apple", "Banana", "Cherry"]
  }

Outside the loop (spotlight on entire map):
  {{.Name}}         → "Alice" ✓
  {{.Items}}        → ["Apple", "Banana", "Cherry"] ✓

{{range .Items}}
  Inside the loop (spotlight on current item):
  {{.}}             → "Apple" (or "Banana", or "Cherry") ✓
  {{.Items}}        → ✗ CRASH! String has no "Items" field
  {{.Name}}         → ✗ CRASH! String has no "Name" field
{{end}}

After the loop (spotlight returns to map):
  {{.Name}}         → "Alice" ✓
  {{.Items}}        → ["Apple", "Banana", "Cherry"] ✓
`)

	// Live example of loops
	fmt.Println("🔄 Live Example of Loops:\n")
	tmpl4 := template.Must(template.New("list").Parse(
		`Items: {{range .Items}}{{.}} {{end}}`))
	fmt.Println("Output: ")
	tmpl4.Execute(os.Stdout, map[string][]string{
		"Items": {"apple", "banana", "cherry"},
	})
	fmt.Println("\n")

	// Advanced example: nested data with dot context
	fmt.Println("Advanced Example: Nested Data\n")
	tmpl7 := template.Must(template.New("nested").Parse(
		`User: {{.User}}, {{with .Details}}Email: {{.Email}}, Phone: {{.Phone}}{{end}}`))
	data := map[string]interface{}{
		"User": "Dave",
		"Details": map[string]string{
			"Email": "dave@example.com",
			"Phone": "555-1234",
		},
	}
	fmt.Println("Output: ")
	tmpl7.Execute(os.Stdout, data)
	fmt.Println("\n")

	fmt.Println("What happened:")
	fmt.Println("  1. {{.User}} accesses the outer map → 'Dave'")
	fmt.Println("  2. {{with .Details}} moves the spotlight to the Details map")
	fmt.Println("  3. {{.Email}} now accesses 'dave@example.com' from Details")
	fmt.Println("  4. {{with}} ends, spotlight returns to outer map\n")

	// ============================================================
	// PART 4: CLI Architecture - The Vending Machine Pattern
	// ============================================================
	fmt.Println("\n\nPART 4: CLI Architecture - The Vending Machine Pattern")
	fmt.Println("=========================================================\n")

	fmt.Println("The Final Example: Why Structure Templates in Maps?\n")

	fmt.Println("🏪 Analogy: A Vending Machine\n")

	fmt.Println("STEP 1: The Inventory (Pre-load Templates)")
	fmt.Println("  - At startup, the machine loads ALL possible items into memory")
	fmt.Println("  - It parses (compiles) them ONCE, not every time someone buys")
	fmt.Println("  - Think of it as: map[string]*template = {\"1\": welcome, \"2\": goodbye}")
	fmt.Println()

	fmt.Println("STEP 2: The Keypad (User Presses a Button)")
	fmt.Println("  - User types their choice: \"1\" or \"2\"")
	fmt.Println("  - This is captured as a string variable")
	fmt.Println()

	fmt.Println("STEP 3: The Dispenser (Look Up & Execute)")
	fmt.Println("  - Machine looks up templates[\"2\"]")
	fmt.Println("  - Fills it with data (like adding milk to coffee)")
	fmt.Println("  - Pushes it out to the user's screen")
	fmt.Println()

	fmt.Println("WHY THIS MATTERS:")
	fmt.Println("  ✓ Parsing is EXPENSIVE (slow)")
	fmt.Println("  ✓ Execution is CHEAP (fast)")
	fmt.Println("  ✓ Do parsing once at startup")
	fmt.Println("  ✓ When user clicks, just execute (instant)\n")

	// Live example of the vending machine pattern
	fmt.Println("🔄 Live Example: Vending Machine Pattern\n")

	// STEP 1: Inventory (Pre-load templates)
	templates := make(map[string]*template.Template)

	templates["welcome"] = template.Must(template.New("welcome").Parse(
		"🎉 Welcome, {{.Name}}! Your total: ${{.Total}}\n"))

	templates["goodbye"] = template.Must(template.New("goodbye").Parse(
		"👋 Goodbye {{.Name}}! Thanks for ${{.Amount}}. See you soon!\n"))

	templates["error"] = template.Must(template.New("error").Parse(
		"⚠️  Error: {{.Message}}\n"))

	fmt.Println("Step 1: Inventory loaded (3 templates parsed & ready)\n")

	// STEP 2: User input (simulated)
	choices := []struct {
		choice string
		data   interface{}
	}{
		{"welcome", map[string]interface{}{"Name": "Alice", "Total": 42.50}},
		{"goodbye", map[string]interface{}{"Name": "Bob", "Amount": 25.00}},
		{"error", map[string]interface{}{"Message": "Invalid card"}},
	}

	fmt.Println("Step 2 & 3: User presses buttons & templates execute instantly\n")

	// STEP 3: Dispenser (execute the selected template)
	for _, c := range choices {
		fmt.Printf("User pressed: '%s'\n", c.choice)
		fmt.Print("Output: ")
		if tmpl, exists := templates[c.choice]; exists {
			tmpl.Execute(os.Stdout, c.data)
		} else {
			fmt.Printf("Template '%s' not found\n", c.choice)
		}
		fmt.Println()
	}

	fmt.Println("Performance Note:")
	fmt.Println("  - Parsed template: 3 templates, parsed ONCE")
	fmt.Println("  - Executed template: Millions of times instantly")
	fmt.Println("  - This is why web servers use this pattern!")
	fmt.Println()

	// ============================================================
	// Summary
	// ============================================================
	fmt.Println("\n=== Summary ===\n")
	fmt.Println("✅ Part 1: Dot notation accesses fields in your data")
	fmt.Println("✅ Part 2: Conditionals use function-call syntax (ge, eq, lt, etc.)")
	fmt.Println("✅ Part 3: The dot changes meaning in loops (spotlight moves)")
	fmt.Println("✅ Part 4: Pre-parse templates, execute them fast (vending machine)")
	fmt.Println()
}

// ============================================================
// REFERENCE SECTION: Template Syntax Cheat Sheet
// ============================================================

// Variable Syntax
// {{.FieldName}}           → Access field in data
// {{.}}                    → Current value (especially in loops)
// {{.Person.Name}}         → Nested field access

// Loops
// {{range .Items}}         → Loop over items
//   {{.}}                  → Current item
// {{end}}

// Conditionals with Function-Call Syntax
// {{if eq .Value 5}}       → Equal
// {{if ne .Value 5}}       → Not Equal
// {{if lt .Value 5}}       → Less Than
// {{if le .Value 5}}       → Less or Equal
// {{if gt .Value 5}}       → Greater Than
// {{if ge .Value 5}}       → Greater or Equal
// {{if and .A .B}}         → Logical AND
// {{if or .A .B}}          → Logical OR

// Template Functions
// {{len .Slice}}           → Built-in: length
// {{.Name | upper}}        → Pipe operator (custom function)
// {{printf "%d" .Count}}   → Format function

// Context Switching
// {{with .Details}}        → Change dot context to .Details
//   {{.Email}}             → Access fields of Details
// {{end}}                  → Context returns to outer data

// Best Practices:
// ✓ Pre-parse templates at startup (parse once, execute many)
// ✓ Use maps to store multiple templates for fast lookup
// ✓ Keep templates in separate files when possible
// ✓ Define custom functions for business logic
// ✓ Use descriptive field names for clarity
// ✓ Remember: the Dot (.) changes meaning in loops and {{with}} blocks
