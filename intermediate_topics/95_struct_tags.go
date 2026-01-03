package main

import "fmt"

// 95 Struct Tags
type DemoTagged95 struct {
	Field string `json:"field" xml:"field"`
}

func Demo95StructTags() {
	fmt.Println("-- 95 Struct Tags --")
	fmt.Println("Tags are used by encoding packages (json/xml) and reflection.")
}
