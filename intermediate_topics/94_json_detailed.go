package intermediate

import (
	"encoding/json"
	"fmt"
)

// Topic 94: json
// Combines pragmatic examples with GoBootcamp patterns


func main() {

	fmt.Println("-- 94 JSON --")
	p := DemoPerson94{Name: "Gopher", Age: 5}
	b, _ := json.Marshal(p)
	fmt.Println("json:", string(b))
}

func bonusJSONPractices() {
	fmt.Println("\n=== BONUS: JSON Best Practices ===")
	fmt.Println(`
1. Use struct tags for JSON key names
2. Use omitempty for optional fields
3. Use - to hide sensitive fields
4. Keep JSON keys lowercase (convention)
5. Validate JSON before unmarshaling
6. Handle parsing errors gracefully
7. Use json.Number for large numbers
8. Consider custom MarshalJSON for complex logic
9. Document expected JSON structure
10. Test with various JSON inputs
	`)
}
