package intermediate

import (
	"fmt"
	"strconv"
)

// Topic 67: generics
// Combines pragmatic examples with GoBootcamp patterns


func main() {

	fmt.Println("-- 67 Generics --")

	// Map over slice with type parameters
	nums := []int{1, 2, 3}
	strs := MapSlice67(nums, func(n int) string { return fmt.Sprintf("#%d", n) })
	fmt.Println("mapped:", strs)

	// Constrained generic
	max_int := Max67(10, 20)
	max_float := Max67(3.14, 2.71)
	fmt.Println("max(10, 20):", max_int)
	fmt.Println("max(3.14, 2.71):", max_float)
}

func MapSlice67[T any, U any](in []T, fn func(T) U) []U {
	out := make([]U, len(in))
	for i, v := range in {
		out[i] = fn(v)
	}
	return out
}

func Max67[T interface{ ~int | ~float64 }](a, b T) T {
	if a > b {
		return a
	}
	return b
}
