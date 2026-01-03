package intermediate

import (
	"fmt"
	"math"
)

// Topic 64: methods
// Combines pragmatic examples with GoBootcamp patterns


func main() {

	fmt.Println("-- 64 Methods --")
	c := Counter64{n: 0}
	c.Inc()
	c.Inc()
	fmt.Println("value:", c.Value())

	// Non-struct method
	var x MyInt = 5
	fmt.Println("MyInt(5).Double():", x.Double())
}
