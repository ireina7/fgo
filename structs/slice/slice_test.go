package slice_test

import (
	"fmt"
	"testing"

	"github.com/ireina7/fgo/interfaces"
	"github.com/ireina7/fgo/interfaces/iter"
	"github.com/ireina7/fgo/structs/number"
	"github.com/ireina7/fgo/structs/option"
	"github.com/ireina7/fgo/structs/slice"
	"github.com/ireina7/fgo/types"
	"github.com/ireina7/fgo/util"
)

func TestHKT(t *testing.T) {
	var xs types.HKT[slice.SliceKind, int]
	var ys slice.Slice[int]
	xs = ys
	util.Use(xs)
}

func TestSlice(t *testing.T) {
	xs := slice.Make(1, 2, 3, 4, 5)
	xs.ForEach(func(i int, x int) {
		fmt.Println(x)
	})
	fmt.Println(slice.Concat(slice.Make(0, 9, 8), xs))
}

func TestSliceIter(t *testing.T) {
	xs := slice.Make(1, 2, 3, 4, 5)
	for _, x := range xs {
		fmt.Println("Got", x)
	}
}

func TestSliceIterZip(t *testing.T) {
	xs := slice.Make(1, 2, 3, 4, 5, 6, -7)
	sliceIter := xs.Iter()
	indexIter := number.IntIter{}.From(0)
	iter := iter.Zip[int, int](indexIter, sliceIter)
	for x := iter.Next(); !option.IsNone(x); x = iter.Next() {
		y := option.Get(x)
		fmt.Println(y.A, y.B)
	}
}

func TestForEach(t *testing.T) {
	xs := slice.Make(1, 2, 3, 4, 5, 6)
	iter.ForEach[int](xs, func(i int, x int) {
		fmt.Println(i, x)
	})
}

func TestCollect(t *testing.T) {
	xs := slice.Make(1, 2, 3, 4, 5, -7)
	p := xs.Iter()
	var collector iter.Collector[slice.SliceKind, int] = slice.NewSliceCollector[int]()
	ys := collector.Collect(p)
	t.Log("Collected", ys)
}

func TestOperation(t *testing.T) {
	xs := slice.Make(1, 2, 3, 4, 5, -7)
	ys := slice.Map(xs, func(x int) int {
		return x * 2
	})
	t.Log(ys)
}

func TestGrouping(t *testing.T) {
	xs := slice.Make(1, 2, 3, 4, 5, 6, 7, 8)
	grouper := slice.Grouping[int, int]{
		Eq:   &interfaces.Prelude[int]{},
		Hash: &interfaces.HashInt{},
	}
	hm := grouper.GroupBy(xs, func(x int) int {
		return x % 2
	})
	t.Logf("%#v", hm.Get(1))
}
