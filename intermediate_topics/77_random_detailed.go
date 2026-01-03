package intermediate

import (
	"fmt"
	"math/rand"
	"time"
)

// Topic 77: Random Numbers
// Generating random numbers for simulations, games, and sampling

func main() {

	// Seed the random number generator
	rand.Seed(time.Now().UnixNano())

	// Generate random integers
	fmt.Println("Random integers:")
	for i := 0; i < 5; i++ {
		fmt.Println(rand.Intn(100)) // 0-99
	}

	// Generate random float
	fmt.Println("\nRandom float [0.0, 1.0):")
	fmt.Println(rand.Float64())

	// Random from slice
	fruits := []string{"apple", "banana", "cherry"}
	fmt.Println("\nRandom fruit:", fruits[rand.Intn(len(fruits))])

	// Crypto random (more secure)
	fmt.Println("\nRandom with range:")
	fmt.Println(rand.Intn(50) + 1) // 1-50

}
