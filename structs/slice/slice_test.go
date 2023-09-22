package slice

import (
	"fmt"
	"testing"
)

func TestSlice(t *testing.T) {
	xs := []int{1, 2, 3, 4}
	ForEach(From(xs), func(i int, x int) {
		fmt.Println(x)
	})
	fmt.Println(Concat(From([]int{0, 9, 8}), From(xs)))
}
