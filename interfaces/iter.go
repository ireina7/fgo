package interfaces

import (
	"github.com/ireina7/fgo/structs/option"
	"github.com/ireina7/fgo/structs/tuple"
)

type Iterator[A any] interface {
	Next() option.Option[A] //Stateful interface!
}

type Iterable[A any] interface {
	Iter() Iterator[A]
}

func For[A any](xs Iterable[A], f func(A)) {
	iter := xs.Iter()
	for x := iter.Next(); !option.IsNone(x); x = iter.Next() {
		option.Map_(x, func(x A) {
			f(x)
		})
	}
}

func ForEach[A any](xs Iterable[A], f func(int, A)) {
	iter := xs.Iter()
	for i, x := 0, iter.Next(); !option.IsNone(x); x = iter.Next() {
		option.Map_(x, func(x A) {
			f(i, x)
		})
		i += 1
	}
}

// for x := range Range(customContainer) {}
func Range[A any](xs Iterable[A]) <-chan A {
	ch := make(chan A)
	iter := xs.Iter()
	go func() {
		defer close(ch)
		for x := iter.Next(); !option.IsNone(x); x = iter.Next() {
			option.Map_(x, func(x A) {
				ch <- x
			})
		}
	}()
	return ch
}

type zipByIter[A, B, C any] struct {
	xs Iterator[A]
	ys Iterator[B]
	f  func(A, B) C
}

func (iter *zipByIter[A, B, C]) Next() option.Option[C] {
	x := iter.xs.Next()
	y := iter.ys.Next()
	return option.FlatMap(x, func(x A) option.Option[C] {
		return option.FlatMap(y, func(y B) option.Option[C] {
			return option.Some[C]{Value: iter.f(x, y)}
		})
	})
}

func ZipBy[A, B, C any](xs Iterator[A], ys Iterator[B], f func(A, B) C) Iterator[C] {
	return &zipByIter[A, B, C]{
		xs: xs,
		ys: ys,
		f:  f,
	}
}

func Zip[A, B any](xs Iterator[A], ys Iterator[B]) Iterator[tuple.Tuple2[A, B]] {
	return ZipBy(xs, ys, func(x A, y B) tuple.Tuple2[A, B] {
		return tuple.Tuple2[A, B]{A: x, B: y}
	})
}
