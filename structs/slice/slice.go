package slice

import (
	"github.com/ireina7/fgo/interfaces"
	"github.com/ireina7/fgo/structs/option"
	"github.com/ireina7/fgo/structs/search/ordered"
)

type SliceKind any

// type SliceType[A any] types.HKT[SliceKind, A]

type Slice[A any] []A

func (s Slice[A]) Kind(SliceKind) {}
func (s Slice[A]) ElemType(A)     {}

func Make[A any](xs ...A) Slice[A] {
	return Slice[A](xs)
}

func From[A any](xs []A) Slice[A] {
	return Slice[A](xs)
}

func Map[A, B any](xs Slice[A], f func(A) B) Slice[B] {
	ys := make([]B, len(xs))
	for i, x := range xs {
		ys[i] = f(x)
	}
	return Slice[B](ys)
}

func Last[A any](xs Slice[A]) option.Option[A] {
	if len(xs) == 0 {
		return option.From[A](nil)
	}
	return option.From[A](&xs[len(xs)-1])
}

func Init[A any](xs Slice[A]) option.Option[Slice[A]] {
	if len(xs) == 0 {
		return option.From[Slice[A]](nil)
	}
	return option.Some[Slice[A]]{Value: xs[0 : len(xs)-1]}
}

func Head[A any](xs Slice[A]) option.Option[A] {
	if len(xs) == 0 {
		return option.From[A](nil)
	}
	return option.Some[A]{Value: xs[0]}
}

func Tail[A any](xs Slice[A]) option.Option[Slice[A]] {
	// xs := list.(slice[A])
	if len(xs) == 0 {
		return option.From[Slice[A]](nil)
	}
	return option.Some[Slice[A]]{Value: xs[1:]}
}

func Concat[A any](xs, ys Slice[A]) Slice[A] {
	zs := make([]A, 0)
	zs = append(zs, xs...)
	zs = append(zs, ys...)
	return Slice[A](zs)
}

func ForEach[A any](ys Slice[A], f func(int, A)) {
	for i, y := range ys {
		f(i, y)
	}
}

func Set[A any](ys Slice[A], index int, x A) Slice[A] {
	ys[index] = x
	return ys
}

func Get[A any](ys Slice[A], index int) option.Option[A] {
	if len(ys) > index {
		return option.From(&ys[index])
	}
	return option.From[A](nil)
}

func Insert[A any](ys Slice[A], index int, x A) option.Option[Slice[A]] {
	if len(ys) < index {
		return option.From[Slice[A]](nil)
	}
	zs := ys[0:index]
	zs = append(zs, x)
	zs = append(zs, ys[index:]...)
	ans := From(zs)
	return option.From[Slice[A]](&ans)
}

func Len[A any](xs Slice[A]) int {
	return len(xs)
}

func ZipBy[A, B, C any](xs Slice[A], ys Slice[B], zip func(A, B) C) Slice[C] {
	len_xs := Len(xs)
	len_ys := Len(ys)
	zs := make([]C, ordered.PreludeMinMax[int]().Min(len_xs, len_ys))
	ForEach(xs, func(i int, x A) {
		option.Map_(Get(ys, i), func(y B) {
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
