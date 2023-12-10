package queue

import "github.com/ireina7/fgo/structs/maybe"

// First-in-first-out queue
type Queue[T any] struct {
	xs []T
}

func Empty[T any]() Queue[T] {
	return Queue[T]{}
}

func FromSlice[T any](xs []T) Queue[T] {
	return Queue[T]{xs: xs}
}

func (q Queue[T]) Len() uint {
	return uint(len(q.xs))
}

func (q Queue[T]) Push(x T) Queue[T] {
	return FromSlice(append(q.xs, x))
}

func (q Queue[T]) Pop() (maybe.Maybe[T], Queue[T]) {
	if q.Len() == 0 {
		return maybe.None[T](), q
	}
	return maybe.Some(q.xs[0]), FromSlice(q.xs[1:])
}

func (q Queue[T]) Peek() maybe.Maybe[T] {
	if q.Len() == 0 {
		return maybe.None[T]()
	}
	return maybe.Some(q.xs[0])
}
