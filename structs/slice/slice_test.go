package slice_test

import (
	"fmt"
	"testing"

	"github.com/ireina7/fgo/interfaces"
	"github.com/ireina7/fgo/interfaces/collection"
	"github.com/ireina7/fgo/structs/maybe"
	"github.com/ireina7/fgo/structs/number"
	"github.com/ireina7/fgo/structs/slice"
	"github.com/ireina7/fgo/structs/tuple"
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
	iter := collection.Zip[int, int](indexIter, sliceIter)
	for x := iter.Next(); !maybe.IsNone(x); x = iter.Next() {
		y := maybe.Get(x)
		fmt.Println(y.A, y.B)
	}
}

func TestForEach(t *testing.T) {
	xs := slice.Make(1, 2, 3, 4, 5, 6)
	collection.ForEach[int](xs, func(i int, x int) {
		fmt.Println(i, x)
	})
}

func TestCollect(t *testing.T) {
	xs := slice.Make(1, 2, 3, 4, 5, -7)
	p := xs.Iter()
	var collector collection.Collector[slice.SliceKind, int] = slice.NewSliceCollector[int]()
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

func TestSequence(t *testing.T) {
	xs := slice.Make(
		maybe.Some(1),
		maybe.Some(3),
		maybe.None[int](),
		maybe.Some(7),
	)
	// sequence := &slice.SequenceSlice[option.OptionKind, int]{
	// 	Functor: &optionFunctor[int, types.Unit]{},
	// 	Pure:    &optionApplicative[slice.Slice[int]]{},
	// }
	sequence := &slice.SequenceSlice[maybe.MaybeKind, int]{
		TraverseSlice: &slice.TraverseSlice[maybe.MaybeKind, types.HKT[maybe.MaybeKind, int], int]{
			Pure:    &optionPure[slice.Slice[int]]{},
			Apply:   &optionApply[slice.Slice[int], int]{},
			Functor: &optionFunctor[tuple.Tuple2[slice.Slice[int], int], slice.Slice[int]]{},
		},
	}
	ys := sequence.Sequence(
		slice.Map(xs, func(x maybe.Maybe[int]) types.HKT[maybe.MaybeKind, int] {
			return x
		}),
	)
	t.Logf("ys: %#v", ys)
}

type optionFunctor[A, B any] struct{}

func (functor *optionFunctor[A, B]) Fmap(
	xs types.HKT[maybe.MaybeKind, A], f func(A) B,
) types.HKT[maybe.MaybeKind, B] {
	return maybe.Map(xs.(maybe.Maybe[A]), f)
}

type optionPure[A any] struct{}

func (self *optionPure[A]) Pure(a A) types.HKT[maybe.MaybeKind, A] {
	return maybe.Some(a)
}

type optionApply[A, B any] struct{}

func (self *optionApply[A, B]) Product(
	fa types.HKT[maybe.MaybeKind, A],
	fb types.HKT[maybe.MaybeKind, B],
) types.HKT[maybe.MaybeKind, tuple.Tuple2[A, B]] {
	a := fa.(maybe.Maybe[A])
	b := fb.(maybe.Maybe[B])
	return maybe.FlatMap(a, func(a A) maybe.Maybe[tuple.Tuple2[A, B]] {
		return maybe.Map(b, func(b B) tuple.Tuple2[A, B] {
			return tuple.Tuple2[A, B]{
				A: a,
				B: b,
			}
		})
	})
}
