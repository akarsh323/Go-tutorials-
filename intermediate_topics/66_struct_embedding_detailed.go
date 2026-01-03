package intermediate

import "fmt"

// Topic 66: struct_embedding
// Combines pragmatic examples with GoBootcamp patterns


func main() {

	fmt.Println("-- 66 Struct Embedding --")

	// Embedded struct: fields promoted
	u := User{Name: "Anna", Contact: Contact{Email: "anna@example.com"}}
	fmt.Println("name:", u.Name)
	fmt.Println("email (promoted):", u.Email)

	// Embedded interface
	var lr LogReader
	fmt.Println("embedded interface Read():", lr.Read())
}

func (fr FileReader) Read() string { return "file contents" }

type LogReader struct {
	FileReader // embedded interface
	log        string
}
