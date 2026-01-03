package intermediate

import "fmt"

// Topic 65: interfaces
// Combines pragmatic examples with GoBootcamp patterns


func main() {

	fmt.Println("-- 65 Interfaces --")

	// Different types satisfying same interface
	var speakers []Speaker
	speakers = append(speakers, Dog{Name: "Rex"})
	speakers = append(speakers, Cat{Name: "Whiskers"})

	for _, s := range speakers {
		fmt.Println(s.Speak())
	}

	// interface{} accepts anything
	var anything interface{} = 42
	fmt.Println("anything:", anything)

	// Type assertion
	if num, ok := anything.(int); ok {
		fmt.Println("it's an int:", num)
	}
}
