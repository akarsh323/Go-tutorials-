package main

import (
	"encoding/xml"
	"fmt"
)

// 96 XML
func Demo96XML() {
	fmt.Println("-- 96 XML --")
	type Node struct {
		XMLName xml.Name `xml:"node"`
		Value   string   `xml:",chardata"`
	}
	n := Node{Value: "hello"}
	b, _ := xml.Marshal(n)
	fmt.Println("xml:", string(b))
}
