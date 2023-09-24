package slice

import (
	"fmt"
	"testing"

	"github.com/ireina7/fgo/interfaces"
	"github.com/ireina7/fgo/types"
	"github.com/ireina7/fgo/util"
)

func TestHKT(t *testing.T) {
	var xs types.HKT[SliceKind, int]
	var ys Slice[int]
	xs = ys
	util.Use(xs)
}

func TestSlice(t *testing.T) {
	xs := Make(1, 2, 3, 4, 5)
	xs.ForEach(func(i int, x int) {
		fmt.Println(x)
	})
	fmt.Println(Concat(Make(0, 9, 8), xs))
}

func TestSliceIter(t *testing.T) {
	xs := Make(1, 2, 3, 4, 5)
	for x := range interfaces.Range[int](xs) {
		fmt.Println("Got", x)
	}
}
