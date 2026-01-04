package main

import (
	"fmt"
	"os"
	"strings"
	"text/template"
)

// Topic 72: text_templates - Complete Breakdown
// =============================================
// Separates CONCEPT (The "What") from CODE (The "How")
//
// Covers:
// Part 1: The Basics - Simple Substitution (Template holes + data)
// Part 2: Logic inside Templates - If/Else with interface{}
// Part 3: Loops - Range with the Shape-Shifting Dot
// Part 4: The CLI Menu App - Architecture with template storage
// Part 5: Key Terms Reference - os.Stdout, bytes.Buffer, FuncMap, {{with}}

func main() {
	fmt.Println("=== 72 TEXT TEMPLATES: Complete Breakdown ===\n")

	// ============================================================
	// PART 1: THE BASICS - Simple Substitution
	// ============================================================
	part1Basics()

	// ============================================================
	// PART 2: LOGIC INSIDE TEMPLATES - If/Else
	// ============================================================
	part2LogicConditionals()

	// ============================================================
	// PART 3: LOOPS - Range with Shape-Shifting Dot
	// ============================================================
	part3Loops()

	// ============================================================
	// PART 4: THE CLI MENU APP - Complete Architecture
	// ============================================================
	part4CLIMenuApp()

	// ============================================================
	// PART 5: KEY TERMS REFERENCE
	// ============================================================
	part5KeyTermsReference()
}

// ============================================================
// PART 1: THE BASICS - Simple Substitution
// ============================================================

func part1Basics() {
	fmt.Println("\n" + strings.Repeat("=", 70))
	fmt.Println("PART 1: THE BASICS - Simple Substitution")
	fmt.Println(strings.Repeat("=", 70) + "\n")

	fmt.Println("📌 THE CONCEPT (The 'What'):")
	fmt.Println("============================")
	fmt.Println(`
You have a TEMPLATE with "holes" in it.
You fill those holes with DATA.

Template (The Blueprint):
  "Hello {{.Name}}!"
   
  The {{.Name}} is the "hole" we will fill.

Data (The Filler):
  {"Name": "Gopher"}
  
  A map where "Name" is the key.

Result (After filling the hole):
  "Hello Gopher!"
`)

	fmt.Println("\n📝 THE CODE (The 'How'):")
	fmt.Println("=========================\n")

	fmt.Println("Step 1: Create the template with the blueprint")
	code1 := `
	tmpl := template.Must(
		template.New("greet").Parse("Hello {{.Name}}!\n"),
	)
	`
	fmt.Println(code1)

	fmt.Println("Step 2: Prepare the data (the filler)")
	code2 := `
	data := map[string]string{"Name": "Gopher"}
	`
	fmt.Println(code2)

	fmt.Println("Step 3: Execute (fill the hole and print)")
	code3 := `
	tmpl.Execute(os.Stdout, data)
	`
	fmt.Println(code3)

	fmt.Println("🔄 LIVE EXECUTION:\n")

	// 1. Create template
	tmpl1 := template.Must(template.New("greet").Parse("Hello {{.Name}}!\n"))

	// 2. Prepare data
	data := map[string]string{"Name": "Gopher"}

	// 3. Execute
	tmpl1.Execute(os.Stdout, data)

	fmt.Println("\n💡 TEACHER'S EXPLANATION:")
	fmt.Println("=========================")
	fmt.Println(`
template.Must(...):
  - A safety wrapper around template parsing
  - If you have a typo in your template, CRASH IMMEDIATELY
  - Better to fail at startup than silently produce wrong output

{{.Name}}:
  - The dot (.) represents YOUR DATA OBJECT
  - .Name means: "Look in the data and get the value of the 'Name' field"
  - If data = {"Name": "Alice"}, then {{.Name}} becomes "Alice"
`)

	fmt.Println("\n✅ KEY TAKEAWAY:")
	fmt.Println("Templates are just text with variables. Dot notation accesses map fields.\n")
}

// ============================================================
// PART 2: LOGIC INSIDE TEMPLATES - If/Else
// ============================================================

func part2LogicConditionals() {
	fmt.Println("\n" + strings.Repeat("=", 70))
	fmt.Println("PART 2: LOGIC INSIDE TEMPLATES - If/Else")
	fmt.Println(strings.Repeat("=", 70) + "\n")

	fmt.Println("📌 THE CONCEPT (The 'What'):")
	fmt.Println("============================")
	fmt.Println(`
Templates aren't just text substitution; they can make DECISIONS.

The Decision:
  "If someone's age is >= 18, call them an adult. Otherwise, a minor."

In Go, you can't write: if .Age >= 18 (that's not template syntax)
Instead, you write: if ge .Age 18 (function-call style)

Why "ge"? It's a FUNCTION, not a SYMBOL.
  - ge = "greater than or equal"
  - Think of it as: ge(.Age, 18) → returns true or false
`)

	fmt.Println("\n📝 THE CODE (The 'How'):")
	fmt.Println("=========================\n")

	fmt.Println("Step 1: Create a template with conditional logic")
	code1 := `
	text := "{{.Name}} is {{if ge .Age 18}}an adult{{else}}a minor{{end}}"
	tmpl := template.Must(template.New("age").Parse(text))
	`
	fmt.Println(code1)

	fmt.Println("Step 2: Prepare data with MIXED TYPES (String + Integer)")
	code2 := `
	// We use map[string]interface{} because our map needs:
	//   - "Name" (a String)
	//   - "Age" (an Integer)
	// A normal map[string]string can only hold Strings!
	
	user := map[string]interface{}{
		"Name": "Bob",
		"Age":  25,
	}
	`
	fmt.Println(code2)

	fmt.Println("Step 3: Execute")
	code3 := `
	tmpl.Execute(os.Stdout, user)
	`
	fmt.Println(code3)

	fmt.Println("🔄 LIVE EXECUTION:\n")

	tmpl2 := template.Must(template.New("age").Parse(
		"{{.Name}} is {{if ge .Age 18}}an adult{{else}}a minor{{end}}\n"))

	user1 := map[string]interface{}{
		"Name": "Bob",
		"Age":  25,
	}

	user2 := map[string]interface{}{
		"Name": "Charlie",
		"Age":  10,
	}

	tmpl2.Execute(os.Stdout, user1)
	tmpl2.Execute(os.Stdout, user2)

	fmt.Println("\n💡 TEACHER'S EXPLANATION:")
	fmt.Println("=========================")
	fmt.Println(`
Why map[string]interface{} and not map[string]string?

map[string]string = "All values must be strings"
  {"Name": "Bob", "Age": "25"}  ✓ Works (Age is a string "25")
  
map[string]interface{} = "Values can be ANYTHING"
  {"Name": "Bob", "Age": 25}  ✓ Works (Age is an integer 25)
  {"Name": "Bob", "Age": "25"}  ✓ Also works (Age is a string "25")

INTERFACE{} = "I don't care what type you are, just be something"

Why does this matter in templates?
  - When comparing with ge, we need the ACTUAL NUMBER 25, not the string "25"
  - The template engine is smarter when you give it the real type

Template Comparison Functions (NOT math symbols):
  eq  → equal
  ne  → not equal
  lt  → less than
  le  → less or equal
  gt  → greater than
  ge  → greater or equal
  and → logical AND
  or  → logical OR
`)

	fmt.Println("Comparison Table:")
	fmt.Println("=================")
	fmt.Println("Normal Code     | Template Code   | Reads As")
	fmt.Println("=============== | =============== | ================")
	fmt.Println("if Age == 18    | if eq .Age 18   | eq(Age, 18)")
	fmt.Println("if Age != 18    | if ne .Age 18   | ne(Age, 18)")
	fmt.Println("if Age < 18     | if lt .Age 18   | lt(Age, 18)")
	fmt.Println("if Age <= 18    | if le .Age 18   | le(Age, 18)")
	fmt.Println("if Age > 18     | if gt .Age 18   | gt(Age, 18)")
	fmt.Println("if Age >= 18    | if ge .Age 18   | ge(Age, 18)")

	fmt.Println("\n✅ KEY TAKEAWAY:")
	fmt.Println("Use map[string]interface{} for mixed types. Conditionals use function-call syntax.\n")
}

// ============================================================
// PART 3: LOOPS - Range with Shape-Shifting Dot
// ============================================================

func part3Loops() {
	fmt.Println("\n" + strings.Repeat("=", 70))
	fmt.Println("PART 3: LOOPS - Range with Shape-Shifting Dot")
	fmt.Println(strings.Repeat("=", 70) + "\n")

	fmt.Println("📌 THE CONCEPT (The 'What'):")
	fmt.Println("============================")
	fmt.Println(`
If you have a LIST of items (like a shopping cart), you need to LOOP through each one.

The Spotlight Analogy:
  The Dot (.) is like a spotlight shining on your data.

BEFORE the loop:
  The spotlight shines on the ENTIRE DATA MAP.
  You can access: {{.Items}}, {{.User}}, {{.Details}}

INSIDE the loop ({{range .Items}}):
  The spotlight MOVES!
  Now it shines ONLY on the CURRENT ITEM.
  {{.}} refers to the individual item (e.g., "apple"), NOT the map.

AFTER the loop ({{end}}):
  The spotlight returns to the outer map.
  You can access: {{.Items}}, {{.User}}, {{.Details}} again
`)

	fmt.Println("\n📝 THE CODE (The 'How'):")
	fmt.Println("=========================\n")

	fmt.Println("Step 1: Create a template with a loop")
	code1 := `
	text := "Items: {{range .Items}}{{.}} {{end}}"
	tmpl := template.Must(template.New("list").Parse(text))
	`
	fmt.Println(code1)

	fmt.Println("Step 2: Prepare data with a SLICE")
	code2 := `
	data := map[string][]string{
		"Items": {"apple", "banana", "cherry"},
	}
	`
	fmt.Println(code2)

	fmt.Println("Step 3: Execute")
	code3 := `
	tmpl.Execute(os.Stdout, data)
	`
	fmt.Println(code3)

	fmt.Println("🔄 LIVE EXECUTION:\n")

	tmpl3 := template.Must(template.New("list").Parse(
		"Items: {{range .Items}}{{.}} {{end}}\n"))

	data3 := map[string][]string{
		"Items": {"apple", "banana", "cherry"},
	}

	tmpl3.Execute(os.Stdout, data3)

	fmt.Println("\n💡 TEACHER'S EXPLANATION:")
	fmt.Println("=========================")
	fmt.Println(`
{{range .Items}} - Start the loop
  For every single item in the Items list, do what's inside this block.

{{.}} - The current item (spotlight focus)
  - First iteration: {{.}} = "apple"
  - Second iteration: {{.}} = "banana"
  - Third iteration: {{.}} = "cherry"

{{end}} - Stop the loop
  When the loop is finished, restore the original spotlight position.

⚠️  CRITICAL: The Spotlight Moves (and gets lost if you're not careful)

Visual Diagram:

Data Structure:
{
  "Name": "Alice",
  "Items": ["Apple", "Banana", "Cherry"]
}

OUTSIDE the loop (spotlight on entire map):
  {{.Name}}         → "Alice"       ✓
  {{.Items}}        → [entire list] ✓

{{range .Items}}    ← Spotlight moves HERE
  INSIDE the loop (spotlight on current item):
  {{.}}             → "Apple"       ✓ (the string itself)
  {{.Items}}        → ✗ CRASH! String "Apple" has no field .Items
  {{.Name}}         → ✗ CRASH! String "Apple" has no field .Name
{{end}}             ← Spotlight moves back

AFTER the loop (spotlight back on map):
  {{.Name}}         → "Alice"       ✓
  {{.Items}}        → [entire list] ✓
`)

	fmt.Println("Advanced Example: Nested Data with {{with}}\n")

	fmt.Println("Step 1: Template with {{with}} (temporary spotlight movement)")
	code4 := `
	text := "User: {{.User}}, {{with .Details}}Email: {{.Email}}, Phone: {{.Phone}}{{end}}"
	tmpl := template.Must(template.New("nested").Parse(text))
	`
	fmt.Println(code4)

	fmt.Println("Step 2: Data with nested structure")
	code5 := `
	data := map[string]interface{}{
		"User": "Dave",
		"Details": map[string]string{
			"Email": "dave@example.com",
			"Phone": "555-1234",
		},
	}
	`
	fmt.Println(code5)

	fmt.Println("🔄 LIVE EXECUTION:\n")

	tmpl4 := template.Must(template.New("nested").Parse(
		"User: {{.User}}, {{with .Details}}Email: {{.Email}}, Phone: {{.Phone}}{{end}}\n"))

	data4 := map[string]interface{}{
		"User": "Dave",
		"Details": map[string]string{
			"Email": "dave@example.com",
			"Phone": "555-1234",
		},
	}

	tmpl4.Execute(os.Stdout, data4)

	fmt.Println("\n💡 What Happened:")
	fmt.Println("==================")
	fmt.Println(`
1. {{.User}} → Accessed outer map → 'Dave'
2. {{with .Details}} → Moved spotlight to the Details map
3. {{.Email}} and {{.Phone}} → Accessed fields of Details
4. {{end}} → Spotlight returns to outer map
5. Result: User: Dave, Email: dave@example.com, Phone: 555-1234
`)

	fmt.Println("\n✅ KEY TAKEAWAY:")
	fmt.Println("The dot (.) changes context. In loops, it becomes the current item. Use {{with}} to temporarily move context.\n")
}

// ============================================================
// PART 4: THE CLI MENU APP - Complete Architecture
// ============================================================

func part4CLIMenuApp() {
	fmt.Println("\n" + strings.Repeat("=", 70))
	fmt.Println("PART 4: THE CLI MENU APP - Complete Architecture")
	fmt.Println(strings.Repeat("=", 70) + "\n")

	fmt.Println("📌 THE CONCEPT (The 'What'):")
	fmt.Println("============================")
	fmt.Println(`
A menu-driven application has THREE PARTS:

1. THE SETUP (Storing Templates)
   - Pre-load ALL possible templates into memory
   - Parse them ONCE at startup (parsing is slow)
   - Store them in a map by name for instant lookup

2. THE USER INPUT (The Listener)
   - Listen to keyboard input from the user
   - Read what they type until they press Enter
   - Clean up the input (remove invisible Enter character)

3. THE SWITCH (The Traffic Cop)
   - Look at what the user chose
   - Find the matching template in the storage
   - Execute it with the user's data
   - Print the result

Why structure it this way?
  ✓ EFFICIENCY: Parse once (slow), execute many times (fast)
  ✓ SCALABILITY: Add 100 templates? Just add to the map.
  ✓ MAINTAINABILITY: All templates in one place.
`)

	fmt.Println("\n📝 THE CODE (The 'How'):")
	fmt.Println("=========================\n")

	fmt.Println("STEP 1: THE SETUP - Pre-load templates into a storage map")
	fmt.Println("=========================================================\n")

	code1 := `
	// Create an empty map to store templates
	// Key = template name (string)
	// Value = the compiled template (*template.Template)
	parsedTemplates := make(map[string]*template.Template)

	// Fill the map with templates
	parsedTemplates["welcome"] = template.Must(
		template.New("welcome").Parse(
			"🎉 Welcome, {{.Name}}! Your total: ${{.Total}}\n"))

	parsedTemplates["goodbye"] = template.Must(
		template.New("goodbye").Parse(
			"👋 Goodbye {{.Name}}! Thanks for ${{.Amount}}.\n"))

	parsedTemplates["error"] = template.Must(
		template.New("error").Parse(
			"⚠️  Error: {{.Message}}\n"))
	`
	fmt.Println(code1)

	// LIVE: Set up templates
	parsedTemplates := make(map[string]*template.Template)

	parsedTemplates["welcome"] = template.Must(
		template.New("welcome").Parse(
			"🎉 Welcome, {{.Name}}! Your total: ${{.Total}}\n"))

	parsedTemplates["goodbye"] = template.Must(
		template.New("goodbye").Parse(
			"👋 Goodbye {{.Name}}! Thanks for ${{.Amount}}.\n"))

	parsedTemplates["error"] = template.Must(
		template.New("error").Parse(
			"⚠️  Error: {{.Message}}\n"))

	fmt.Println("✅ Step 1 Complete: 3 templates are now loaded in memory (parsed & ready)\n")

	fmt.Println("\nSTEP 2: THE USER INPUT - Listen to keyboard (Simulated)")
	fmt.Println("=======================================================\n")

	code2 := `
	// In a REAL app, you would do:
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter your name: ")
	name, _ := reader.ReadString('\n')
	name = strings.TrimSpace(name)  // Remove the invisible Enter character

	// For this demo, we'll simulate the input
	`
	fmt.Println(code2)

	// LIVE: Simulate user input
	simulatedName := "Alice"
	fmt.Printf("(Simulated) User entered: %s\n\n", simulatedName)

	fmt.Println("\nSTEP 3: THE SWITCH - Traffic Cop routing")
	fmt.Println("========================================\n")

	code3 := `
	// Simulate user choosing different options
	choices := []struct {
		choice string
		data   interface{}
	}{
		{"welcome", map[string]interface{}{"Name": "Alice", "Total": 42.50}},
		{"goodbye", map[string]interface{}{"Name": "Bob", "Amount": 25.00}},
		{"error", map[string]interface{}{"Message": "Invalid card"}},
	}

	// For each choice, look it up and execute
	for _, c := range choices {
		if tmpl, exists := parsedTemplates[c.choice]; exists {
			tmpl.Execute(os.Stdout, c.data)
		}
	}
	`
	fmt.Println(code3)

	// LIVE: Execute each template
	fmt.Println("🔄 LIVE EXECUTION:\n")

	choices := []struct {
		choice string
		data   interface{}
	}{
		{"welcome", map[string]interface{}{"Name": "Alice", "Total": 42.50}},
		{"goodbye", map[string]interface{}{"Name": "Bob", "Amount": 25.00}},
		{"error", map[string]interface{}{"Message": "Invalid card"}},
	}

	for _, c := range choices {
		fmt.Printf("User pressed: '%s'\n", c.choice)
		fmt.Print("Output: ")
		if tmpl, exists := parsedTemplates[c.choice]; exists {
			tmpl.Execute(os.Stdout, c.data)
		} else {
			fmt.Printf("Template '%s' not found\n", c.choice)
		}
		fmt.Println()
	}

	fmt.Println("\n💡 PERFORMANCE EXPLANATION:")
	fmt.Println("============================")
	fmt.Println(`
Parsing vs Execution:
  - Parsing: Converting template text to instructions (SLOW)
  - Execution: Using pre-parsed template with data (FAST)

In this app:
  - Parse once: 3 templates, parsed at startup (happens 3 times total)
  - Execute many: Could run the same template 1,000,000 times instantly

Why parse at startup?
  - User requests happen continuously (potentially 1000s/sec on a web server)
  - Parsing every time = 1000s of slow operations
  - Pre-parse = parse once, serve 1000s of fast requests

This is why EVERY WEB SERVER (Python, Go, Ruby, etc.) does this pattern!
`)

	fmt.Println("\n✅ KEY TAKEAWAY:")
	fmt.Println("Structure: Setup (parse once) → Input (listen) → Switch (execute fast). This is a universal pattern.\n")
}

// ============================================================
// PART 5: KEY TERMS REFERENCE
// ============================================================

func part5KeyTermsReference() {
	fmt.Println("\n" + strings.Repeat("=", 70))
	fmt.Println("PART 5: KEY TERMS REFERENCE")
	fmt.Println(strings.Repeat("=", 70) + "\n")

	fmt.Println("Summary of Important Concepts:\n")

	fmt.Println("📌 os.Stdout (Where output goes)")
	fmt.Println("================================")
	fmt.Println(`
What it means: The computer screen (your Terminal window)

When you write:
  tmpl.Execute(os.Stdout, data)
  
You're saying: "Print the result to the terminal screen"

Other options:
  file, _ := os.Create("output.txt")
  tmpl.Execute(file, data)          ← Write to a file instead
  
  var buf bytes.Buffer
  tmpl.Execute(&buf, data)           ← Write to memory instead
`)

	fmt.Println("\n📌 bytes.Buffer (Temporary holding area in memory)")
	fmt.Println("===================================================")
	fmt.Println(`
What it means: A place to store text in memory instead of printing it

Example: You want to capture output in a variable first
  var buf bytes.Buffer
  tmpl.Execute(&buf, data)        ← Store in memory
  
  result := buf.String()           ← Convert to string
  fmt.Println("Result:", result)   ← Now you can use it
  
When to use:
  ✓ Sending output via email (capture it first)
  ✓ Logging (save to file after generation)
  ✓ Testing (check if output is correct before using)
`)

	fmt.Println("\n📌 FuncMap (Teaching templates new functions)")
	fmt.Println("==============================================")
	fmt.Println(`
What it means: A way to teach templates custom functions they don't know by default

Example: Templates don't have a .Upper() function by default
  funcMap := template.FuncMap{
    "upper": func(s string) string { 
      return strings.ToUpper(s) 
    },
  }
  
  tmpl := template.New("test").Funcs(funcMap).Parse(
    "Uppercase: {{upper .Name}}")
  
  tmpl.Execute(os.Stdout, map[string]string{"Name": "alice"})
  // Output: Uppercase: ALICE

Common custom functions:
  "title"     → "alice smith" → "Alice Smith"
  "lower"     → "HELLO" → "hello"
  "reverse"   → "abc" → "cba"
  "repeat"    → "a" × 5 → "aaaaa"
`)

	fmt.Println("\n📌 {{with .Field}} (Spotlight shortcut)")
	fmt.Println("========================================")
	fmt.Println(`
What it means: Temporarily move the spotlight to focus on one field

Instead of:
  {{.Details.Email}} {{.Details.Phone}} {{.Details.Address}}
  
You can write:
  {{with .Details}}
    {{.Email}} {{.Phone}} {{.Address}}
  {{end}}

Why use it?
  ✓ Cleaner, easier to read
  ✓ Avoid repeating .Details over and over
  ✓ Like a "temporary zoom" into that section

Example:
  {{with .User}}
    Name: {{.Name}}
    Email: {{.Email}}
  {{end}}
  
This is equivalent to:
  Name: {{.User.Name}}
  Email: {{.User.Email}}
`)

	fmt.Println("\n📌 map[string]interface{} (Mixed types)")
	fmt.Println("========================================")
	fmt.Println(`
What it means: A map that can hold ANY type of value

Limitation of map[string]string:
  m := map[string]string{
    "name": "Alice",     ✓ string
    "age": "25",         ✓ (must be a string, even though it's a number)
    "age": 25,           ✗ ERROR! Only strings allowed
  }

Power of map[string]interface{}:
  m := map[string]interface{}{
    "name": "Alice",     ✓ string
    "age": 25,           ✓ integer
    "scores": []int{90}, ✓ slice
    "valid": true,       ✓ boolean
  }

Why templates care:
  - When comparing numbers with ge, lt, etc., the engine needs the REAL type
  - "25" (string) and 25 (int) behave differently in comparisons
  - Use interface{} to give templates the full picture
`)

	fmt.Println("\n📌 bufio.NewReader (Reading keyboard input)")
	fmt.Println("============================================")
	fmt.Println(`
What it means: A way to read user input from the keyboard

Example:
  reader := bufio.NewReader(os.Stdin)
  fmt.Print("Enter your name: ")
  name, _ := reader.ReadString('\n')  ← Read until Enter key
  name = strings.TrimSpace(name)      ← Remove the invisible Enter
  
Why TrimSpace()?
  When you type "Alice" and press Enter, Go reads it as "Alice\n"
  The \n is an invisible character representing the Enter key
  TrimSpace() removes it, leaving just "Alice"

What is os.Stdin?
  - Stdin = Standard Input (your keyboard)
  - Stdout = Standard Output (your screen)
  - Stderr = Standard Error (also your screen, for error messages)
`)

	fmt.Println("\n" + strings.Repeat("=", 70))
	fmt.Println("FINAL SUMMARY")
	fmt.Println(strings.Repeat("=", 70) + "\n")

	fmt.Println("✅ PART 1: Templates are text with {{.Field}} holes. Dots access map fields.")
	fmt.Println("✅ PART 2: Conditionals use function-call syntax (ge, eq, lt). Use interface{} for mixed types.")
	fmt.Println("✅ PART 3: The dot (.) changes meaning in loops. Use {{range}} and {{with}} carefully.")
	fmt.Println("✅ PART 4: Pre-parse templates, execute fast. This is the universal web server pattern.")
	fmt.Println("✅ PART 5: os.Stdout, bytes.Buffer, FuncMap, {{with}} are powerful tools.")
	fmt.Println("\n🎯 Master these 5 parts, and you master Go text templates.\n")
}
