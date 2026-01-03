package intermediate

import (
	"bytes"
	"fmt"
	"os"
	"text/template"
)

// Topic 72: text_templates
// Combines pragmatic examples with GoBootcamp patterns


func main() {

	fmt.Println("-- 72 Text Templates --")

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

	// Template with conditionals
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

	// Template with loops (range)
	tmpl4 := template.Must(template.New("list").Parse(
		"Items: {{range .Items}}{{.}} {{end}}\n"))
	tmpl4.Execute(os.Stdout, map[string][]string{
		"Items": {"apple", "banana", "cherry"},
	})

	// Template with helper functions
	funcMap := template.FuncMap{
		"upper": func(s string) string { return strings.ToUpper(s) },
	}
	tmpl5 := template.Must(template.New("func").Funcs(funcMap).Parse(
		"Uppercase: {{upper .Text}}\n"))
	tmpl5.Execute(os.Stdout, map[string]string{
		"Text": "hello world",
	})

	// Template with buffer output (instead of directly to stdout)
	var buf bytes.Buffer
	tmpl6 := template.Must(template.New("buffer").Parse(
		"Result: {{.Value}}\n"))
	tmpl6.Execute(&buf, map[string]int{"Value": 42})
	fmt.Print("Buffer output: ", buf.String())

	// Template with with statement (context change)
	tmpl7 := template.Must(template.New("nested").Parse(
		"User: {{.User}}, {{with .Details}}Email: {{.Email}}, Phone: {{.Phone}}{{end}}\n"))
	data := map[string]interface{}{
		"User": "Dave",
		"Details": map[string]string{
			"Email": "dave@example.com",
			"Phone": "555-1234",
		},
	}
	tmpl7.Execute(os.Stdout, data)
}

func templateExample1() {
	
	fmt.Println("📚 Use {{.FieldName}} to insert values")

	// Define template string
	templateString := `Hello {{.Name}}!
You are {{.Age}} years old.`

	// Parse template
	tmpl, err := template.New("greeting").Parse(templateString)
	if err != nil {
		log.Fatal(err)
	}

	// Data to fill in
	data := map[string]interface{}{
		"Name": "Alice",
		"Age":  30,
	}

	// Execute template
	var output bytes.Buffer
	tmpl.Execute(&output, data)

	fmt.Println("Output:")
	fmt.Println(output.String())
}

func templateExample10() {
	
	fmt.Println(`
Template Syntax:

Variables:
  {{.FieldName}}           → Access field
  {{.}}                    → Current value
  {{.Person.Name}}         → Nested field

Loops:
  {{range .Items}}         → Loop over items
    {{.}}                  → Current item
  {{end}}

Conditionals:
  {{if .Condition}}        → If true
  {{else if}}              → Else if
  {{else}}                 → Else
  {{end}}

Functions:
  {{len .Slice}}           → Built-in functions
  {{.Name | lower}}        → Pipe operator
  {{printf "%d" .Count}}   → Format functions

Comparison Functions (use in if):
  eq  → equal
  ne  → not equal
  lt  → less than
  le  → less than or equal
  gt  → greater than
  ge  → greater than or equal
  and → logical AND
  or  → logical OR

  {{if gt .Age 18}}Adult{{end}}
  {{if and .IsActive .IsAdmin}}Active Admin{{end}}

Best Practices:
  ✓ Separate template strings to external files
  ✓ Define custom functions for business logic
  ✓ Use descriptive field names
  ✓ Keep templates simple
  ✓ Test template output
	`)
}
