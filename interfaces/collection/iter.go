package collection

import (
	"github.com/ireina7/fgo/interfaces/impl"
	"github.com/ireina7/fgo/structs/maybe"
	"github.com/ireina7/fgo/structs/tuple"
	"github.com/ireina7/fgo/types"
)

type Iterator[A any] interface {
	Next() maybe.Maybe[A] //Stateful interface!
}

type Iterable[A any] interface {
	Iter() Iterator[A]
}

type FromIterator[F_, A any] interface {
	FromIter(Iterator[A]) types.HKT[F_, A]
}

func Iterate[A any](iter Iterator[A], f func(A)) {
	for x := iter.Next(); !maybe.IsNone(x); x = iter.Next() {
		f(maybe.Get(x))
	}
}

type emptyIter[A any] struct{}

func (iter *emptyIter[A]) Next() maybe.Maybe[A] {
	return maybe.None[A]()
}

func EmptyIter[A any]() Iterator[A] {
	return &emptyIter[A]{}
}

func For[A any](xs Iterable[A], f func(A)) {
	iter := xs.Iter()
	for x := iter.Next(); !maybe.IsNone(x); x = iter.Next() {
		maybe.Map_(x, func(x A) {
			f(x)
		})
	}
}

func ForEach[A any](xs Iterable[A], f func(int, A)) {
	iter := xs.Iter()
	for i, x := 0, iter.Next(); !maybe.IsNone(x); x = iter.Next() {
		maybe.Map_(x, func(x A) {
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
		for x := iter.Next(); !maybe.IsNone(x); x = iter.Next() {
			maybe.Map_(x, func(x A) {
				ch <- x
			})
		}
	}()
	return ch
}

type IteratorIsIterable[A any] struct {
	iter Iterator[A]
}

func (iter *IteratorIsIterable[A]) Iter() Iterator[A] {
	return iter.iter
}

func FromIter[A any](iter Iterator[A]) Iterable[A] {
	return &IteratorIsIterable[A]{iter: iter}
}

type zipByIter[A, B, C any] struct {
	xs Iterator[A]
	ys Iterator[B]
	f  func(A, B) C
}

func (iter *zipByIter[A, B, C]) Next() maybe.Maybe[C] {
	x := iter.xs.Next()
	y := iter.ys.Next()
	return maybe.FlatMap(x, func(x A) maybe.Maybe[C] {
		return maybe.FlatMap(y, func(y B) maybe.Maybe[C] {
			return maybe.Some[C](iter.f(x, y))
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

type IterFunctor[F_, A, B any] struct {
	impl.Implement[types.HKT[F_, A], Iterable[A]] //Well, we need summoner!
	FromIterator[F_, B]
}

type IterMap[A, B any] struct {
	iter Iterator[A]
	f    func(A) B
}

func (iter *IterMap[A, B]) Next() maybe.Maybe[B] {
	return maybe.Map(iter.iter.Next(), iter.f)
}

func (functor *IterFunctor[F_, A, B]) Fmap(
	xs types.HKT[F_, A],
	f func(A) B,
) types.HKT[F_, B] {
	as := functor.Impl(xs)
	iter := &IterMap[A, B]{
		iter: as.Iter(),
		f:    f,
	}
	return functor.FromIter(iter)
}
