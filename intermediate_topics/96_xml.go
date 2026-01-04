package main

import (
	"encoding/xml"
	"fmt"
	"strings"
)

/*
═══════════════════════════════════════════════════════════════════════════════
                         XML PROCESSING IN GO
═══════════════════════════════════════════════════════════════════════════════

Just like JSON, Go handles XML using struct tags. However, XML is more complex
because it distinguishes between "Attributes" (<tag attr="val">) and
"Child Elements" (<tag><child>val</child></tag>).

KEY CONCEPTS:
  • XMLName: A special field to define the root element name.
  • `xml:"name"`: Maps a field to a child element <name>.
  • `xml:"name,attr"`: Maps a field to an attribute inside the tag.
  • `xml:",chardata"`: Reads/writes raw text content inside a tag.

═══════════════════════════════════════════════════════════════════════════════
*/

/*
━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━
  SECTION 1: THE BASICS (ELEMENTS, ATTRIBUTES, IGNORING)
━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━
*/

// Book represents a standard XML structure
type Book struct {
	// 1. XMLName: Special field. Defines the name of the root tag for this struct.
	//    If omitted, Go uses the struct name "Book".
	XMLName xml.Name `xml:"catalog_item"`

	// 2. Attribute: Lives INSIDE the opening tag <catalog_item id="101">
	ID string `xml:"id,attr"`

	// 3. Child Element: Lives BETWEEN tags <title>Go Programming</title>
	Title string `xml:"title"`

	// 4. Omit Empty: Won't appear if value is 0 or ""
	Price float64 `xml:"price,omitempty"`

	// 5. Ignoring: Internal field, never exported to XML
	InternalCode string `xml:"-"`
}

func Example1_BasicXML() {
	fmt.Println("\n" + strings.Repeat("═", 80))
	fmt.Println("EXAMPLE 1: Elements vs Attributes")
	fmt.Println(strings.Repeat("═", 80) + "\n")

	b := Book{
		ID:           "BK-2025",
		Title:        "Mastering XML in Go",
		Price:        49.99,
		InternalCode: "SECRET_SKU_99",
	}

	fmt.Println("📌 Marshaling a struct to XML:")

	// We use MarshalIndent to make it human-readable (pretty-printed)
	xmlBytes, _ := xml.MarshalIndent(b, "", "  ")

	fmt.Printf("%s\n", string(xmlBytes))
	fmt.Println("\nAnalysis:")
	fmt.Println("  • Root is <catalog_item> (from XMLName)")
	fmt.Println("  • 'id' is inside the tag (Attribute)")
	fmt.Println("  • 'title' is its own tag (Child Element)")
	fmt.Println("  • 'InternalCode' is missing (-)")
}

/*
━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━
  SECTION 2: NESTED STRUCTURES & LISTS
━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━
  XML is hierarchical. Go handles this by embedding structs or using slices.
  Slices automatically become a list of repeated XML tags.
━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━
*/

type Address struct {
	City  string `xml:"city"`
	State string `xml:"state"`
}

type Employee struct {
	XMLName xml.Name `xml:"employee"`
	Name    string   `xml:"full_name"`
	// Nested Struct: Appears as children inside <employee>
	Address Address `xml:"address"`
}

type Company struct {
	XMLName xml.Name `xml:"company"`
	// Slice: Creates multiple <employee> tags one after another
	Staff []Employee `xml:"staff>employee"`
	// The "staff>employee" tag syntax is a shorthand to wrap them in a <staff> parent!
}

func Example2_NestedAndLists() {
	fmt.Println("\n" + strings.Repeat("═", 80))
	fmt.Println("EXAMPLE 2: Nested Structs & Slices")
	fmt.Println(strings.Repeat("═", 80) + "\n")

	comp := Company{
		Staff: []Employee{
			{Name: "Akarsh", Address: Address{City: "Bangalore", State: "KA"}},
			{Name: "John", Address: Address{City: "London", State: "UK"}},
		},
	}

	xmlBytes, _ := xml.MarshalIndent(comp, "", "    ")
	fmt.Println(string(xmlBytes))

	fmt.Println("\nAnalysis:")
	fmt.Println("  • The 'staff>employee' tag created a wrapping <staff> tag automatically.")
	fmt.Println("  • Inside that, it generated a list of <employee> items.")
}

/*
━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━
  SECTION 3: UNMARSHALLING (READING XML)
━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━
  Reading XML is strictly about mapping the shape of the data to the shape
  of your struct.
━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━
*/

func Example3_Unmarshalling() {
	fmt.Println("\n" + strings.Repeat("═", 80))
	fmt.Println("EXAMPLE 3: Unmarshalling (Parsing XML string)")
	fmt.Println(strings.Repeat("═", 80) + "\n")

	// Raw XML string (simulating a file or API response)
	rawXML := `
	<server status="active">
		<ip>192.168.1.1</ip>
		<location>DataCenter-1</location>
	</server>`

	// We define a struct that matches the shape of the XML
	type ServerConfig struct {
		XMLName  xml.Name `xml:"server"`
		Status   string   `xml:"status,attr"` // Read the attribute!
		IP       string   `xml:"ip"`
		Location string   `xml:"location"`
	}

	var config ServerConfig

	// Convert string to bytes and unmarshal
	err := xml.Unmarshal([]byte(rawXML), &config)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("📌 Parsed Data into Go Struct:")
	fmt.Printf("  • Status (Attr): %s\n", config.Status)
	fmt.Printf("  • IP (Child):    %s\n", config.IP)
	fmt.Printf("  • Location:      %s\n", config.Location)
}

/*
═══════════════════════════════════════════════════════════════════════════════
                         QUICK REFERENCE TABLE
═══════════════════════════════════════════════════════════════════════════════

TAG SYNTAX           | EFFECT                                | EXAMPLE OUTPUT
─────────────────────|───────────────────────────────────────|───────────────────
`xml:"name"`         | Child element                         | <name>Value</name>
`xml:"name,attr"`    | Attribute of parent                   | <parent name="Value">
`xml:",chardata"`    | The raw text content of the tag       | <tag>Value</tag>
`xml:"a>b"`          | Nested XML elements                   | <a><b>Value</b></a>
`xml:"-"`            | Ignore field                          | (Hidden)
`xml:",omitempty"`   | Hide if Zero Value (0, "", nil)       | (Hidden)

BEST PRACTICES:
  1. Use `XMLName` if you need to control the root element name strictly.
  2. Use `,attr` for simple metadata (IDs, codes) and Child Elements for actual data.
  3. XML is case-sensitive! Ensure your struct tags match the XML exactly.

═══════════════════════════════════════════════════════════════════════════════
*/

func main() {
	fmt.Println("XML ENCODING DEMO - GO")
	Example1_BasicXML()
	Example2_NestedAndLists()
	Example3_Unmarshalling()
}
