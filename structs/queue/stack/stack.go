package stack

import "github.com/ireina7/fgo/structs/maybe"

// First-in-last-out stack
type Stack[T any] []T

func Empty[T any]() Stack[T] {
	return []T{}
}

func FromSlice[T any](xs []T) Stack[T] {
	return Stack[T](xs)
}

func (s Stack[T]) Len() uint {
	return uint(len(s))
}

func (s Stack[T]) Push(x T) Stack[T] {
	return FromSlice(append(s, x))
}

func (s Stack[T]) Pop() (maybe.Maybe[T], Stack[T]) {
	if s.Len() == 0 {
		return maybe.None[T](), s
	}
	return maybe.Some(s[s.Len()-1]), FromSlice(s[0 : s.Len()-1])
}

func (s Stack[T]) Peek() maybe.Maybe[T] {
	if s.Len() == 0 {
		return maybe.None[T]()
	}
	return maybe.Some(s[s.Len()-1])
}
