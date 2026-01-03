package intermediate

import (
	"fmt"
	"strconv"
)

// Topic 78: Number Parsing
// Converting strings to numbers and handling errors

func main() {

	numStr := "12345"
	num, err := strconv.Atoi(numStr)
	if err != nil {
		fmt.Println("Error parsing the value:", err)
	}
	fmt.Println("Parsed Integer:", num)
	fmt.Println("Parsed Integer:", num+1)

	numistr, err := strconv.ParseInt(numStr, 10, 32)
	if err != nil {
		fmt.Println("Error parsing the value:", err)
	}

	fmt.Println("Parsed Integer:", numistr)

	floatstr := "3.14"
	floatval, err := strconv.ParseFloat(floatstr, 64)
	if err != nil {
		fmt.Println("Error parsing value:", err)
	}
	fmt.Printf("Parsed float: %.2f\n", floatval)

	binaryStr := "1010" // 0 + 2 + 0 + 8 = 10
	decimal, err := strconv.ParseInt(binaryStr, 2, 64)
	if err != nil {
		fmt.Println("Error parsing binary value:", err)
		return
	}
	fmt.Println("Parsed binary to decimal:", decimal)

}
