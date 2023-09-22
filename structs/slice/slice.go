package slice

import (
	"github.com/ireina7/fgo/structs/option"
	"github.com/ireina7/fgo/types"
)

type SliceKind any

type Slice[A any] types.HKT[SliceKind, A]

type slice[A any] []A

func (s slice[A]) Kind(SliceKind) {}
func (s slice[A]) ElemType(A)     {}

func MakeSlice[A any](xs []A, x ...A) Slice[A] {
	xs = append(xs, x...)
	return slice[A](xs)
}

func From[A any](xs []A) Slice[A] {
	return slice[A](xs)
}

func Map[A, B any](list Slice[A], f func(A) B) Slice[B] {
	xs := list.(slice[A])
	ys := make([]B, len(xs))
	for i, x := range xs {
		ys[i] = f(x)
	}
	return slice[B](ys)
}

func Last[A any](list Slice[A]) option.Option[A] {
	xs := list.(slice[A])
	if len(xs) == 0 {
		return option.None[A]{}
	}
	return option.Some[A]{Value: xs[len(xs)-1]}
}

func Init[A any](list Slice[A]) option.Option[Slice[A]] {
	xs := list.(slice[A])
	if len(xs) == 0 {
		return option.None[Slice[A]]{}
	}
	return option.Some[Slice[A]]{Value: xs[0 : len(xs)-1]}
}

func Head[A any](list Slice[A]) option.Option[A] {
	xs := list.(slice[A])
	if len(xs) == 0 {
		return option.None[A]{}
	}
	return option.Some[A]{Value: xs[0]}
}

func Tail[A any](list Slice[A]) option.Option[Slice[A]] {
	xs := list.(slice[A])
	if len(xs) == 0 {
		return option.None[Slice[A]]{}
	}
	return option.Some[Slice[A]]{Value: xs[1:]}
}

func Concat[A any](list0, list1 Slice[A]) Slice[A] {
	xs := list0.(slice[A])
	ys := list1.(slice[A])
	zs := make([]A, 0)
	zs = append(zs, xs...)
	zs = append(zs, ys...)
	return slice[A](zs)
}

func ForEach[A any](xs Slice[A], f func(int, A)) {
	ys := xs.(slice[A])
	for i, y := range ys {
		f(i, y)
	}
}

func Set[A any](xs Slice[A], index int, x A) Slice[A] {
	ys := xs.(slice[A])
	ys[index] = x
	return ys
}

func Get[A any](xs Slice[A], index int) option.Option[A] {
	ys := xs.(slice[A])
	if len(ys) > index {
		return option.From(&ys[index])
	}
	return option.From[A](nil)
}

func Len[A any](xs Slice[A]) int {
	return len(xs.(slice[A]))
}
