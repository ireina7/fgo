package slice

import (
	"github.com/ireina7/fgo/interfaces"
	"github.com/ireina7/fgo/structs/option"
	"github.com/ireina7/fgo/structs/search/ordered"
	"github.com/ireina7/fgo/types"
)

type SliceKind types.Kind

// type SliceType[A any] types.HKT[SliceKind, A]

// Slice
type Slice[A any] []A

func (s Slice[A]) Kind(SliceKind) {}
func (s Slice[A]) ElemType(A)     {}

func Make[A any](xs ...A) Slice[A] {
	return Slice[A](xs)
}

func From[A any](xs []A) Slice[A] {
	return Slice[A](xs)
}

func Room[A any](length int) Slice[A] {
	return make([]A, length)
}

func Empty[A any]() Slice[A] {
	return Room[A](0)
}

func Container[A any](capacity int) Slice[A] {
	return make([]A, 0, capacity)
}

func Cons[A any](x A, xs Slice[A]) Slice[A] {
	ys := Make(x)
	return ys.Append(xs...)
}

func Map[A, B any](xs Slice[A], f func(A) B) Slice[B] {
	ys := make([]B, len(xs))
	for i, x := range xs {
		ys[i] = f(x)
	}
	return Slice[B](ys)
}

func (xs Slice[A]) Filter(f func(A) bool) Slice[A] {
	ys := Make[A]()
	xs.ForEach(func(_ int, x A) {
		if f(x) {
			ys = ys.Append(x)
		}
	})
	return ys
}

func FlatMap[A, B any](xs Slice[A], f func(A) Slice[B]) Slice[B] {
	ys := Make[B]()
	xs.ForEach(func(_ int, x A) {
		bs := f(x)
		ys = ys.Append(bs...)
	})
	return ys
}

func Flatten[A any](xxs Slice[Slice[A]]) Slice[A] {
	ys := Make[A]()
	xxs.ForEach(func(_ int, xs Slice[A]) {
		ys = ys.Append(xs...)
	})
	return ys
}

func (xs Slice[A]) Last() option.Option[A] {
	if len(xs) == 0 {
		return option.From[A](nil)
	}
	return option.From[A](&xs[len(xs)-1])
}

func (xs Slice[A]) Init() option.Option[Slice[A]] {
	if len(xs) == 0 {
		return option.From[Slice[A]](nil)
	}
	return option.Some[Slice[A]]{Value: xs[0 : len(xs)-1]}
}

func (xs Slice[A]) Head() option.Option[A] {
	if len(xs) == 0 {
		return option.From[A](nil)
	}
	return option.Some[A]{Value: xs[0]}
}

func (xs Slice[A]) Tail() option.Option[Slice[A]] {
	if len(xs) == 0 {
		return option.From[Slice[A]](nil)
	}
	return option.Some[Slice[A]]{Value: xs[1:]}
}

func (xs Slice[A]) Slice(start, end int) Slice[A] {
	return From(xs[start:end])
}

func Concat[A any](xs, ys Slice[A]) Slice[A] {
	zs := make([]A, 0)
	zs = append(zs, xs...)
	zs = append(zs, ys...)
	return Slice[A](zs)
}

func (ys Slice[A]) ForEach(f func(int, A)) {
	for i, y := range ys {
		f(i, y)
	}
}

func (ys Slice[A]) Set(index int, x A) Slice[A] {
	ys[index] = x
	return ys
}

func (ys Slice[A]) Get(index int) option.Option[A] {
	if len(ys) > index {
		return option.From(&ys[index])
	}
	return option.From[A](nil)
}

func (ys Slice[A]) Insert(index int, x A) option.Option[Slice[A]] {
	if len(ys) < index {
		return option.From[Slice[A]](nil)
	}
	zs := ys[0:index]
	zs = append(zs, x)
	zs = append(zs, ys[index:]...)
	ans := From(zs)
	return option.From[Slice[A]](&ans)
}

func (xs Slice[A]) Append(x ...A) Slice[A] {
	ys := append(xs, x...)
	return From(ys)
}

func (xs Slice[A]) Len() int {
	return len(xs)
}

func ZipBy[A, B, C any](xs Slice[A], ys Slice[B], zip func(A, B) C) Slice[C] {
	zs := make([]C, ordered.PreludeMinMax[int]().Min(xs.Len(), ys.Len()))
	xs.ForEach(func(i int, x A) {
		option.Map_(ys.Get(i), func(y B) {
			zs[i] = zip(x, y)
		})
	})
	return From(zs)
}

type sliceIter[A any] struct {
	s []A
}

func (iter *sliceIter[A]) Next() option.Option[A] {
	if len(iter.s) == 0 {
		return option.From[A](nil)
	}
	res := option.From[A](&iter.s[0])
	iter.s = iter.s[1:]
	return res
}

func (s Slice[A]) Iter() interfaces.Iterator[A] {
	return &sliceIter[A]{
		s: s[:],
	}
}

type Distinct[A comparable] struct{}

func (self *Distinct[A]) Distinct(xs Slice[A]) Slice[A] {
	ys := Empty[A]()
	memo := map[A]struct{}{}
	for _, x := range xs {
		_, exist := memo[x]
		if exist {
			continue
		}
		memo[x] = struct{}{}
		ys = append(ys, x)
	}
	return ys
}

type SliceFromIter[A any] struct{}

func (self *SliceFromIter[A]) FromIter(iter interfaces.Iterator[A]) types.HKT[SliceKind, A] {
	xs := Empty[A]()
	for x := iter.Next(); !option.IsNone(x); x = iter.Next() {
		xs = xs.Append(option.Get(x))
	}
	return xs
}

type SliceCollector[A any] struct {
	interfaces.FromIterator[SliceKind, A]
}

func NewSliceCollector[A any]() *SliceCollector[A] {
	return &SliceCollector[A]{
		FromIterator: &SliceFromIter[A]{},
	}
}

func (co *SliceCollector[A]) Collect(iter interfaces.Iterator[A]) types.HKT[SliceKind, A] {
	return co.FromIter(iter)
}
