package prelude

import (
	"cmp"

	"github.com/ireina7/fgo/structs/maybe"
	"github.com/ireina7/fgo/structs/number"
)

// This will be a max-heap
type PriorityQueue[T cmp.Ordered] struct {
	xs []T
}

func Empty[T cmp.Ordered]() PriorityQueue[T] {
	return PriorityQueue[T]{}
}

func FromSlice[T cmp.Ordered](xs []T) PriorityQueue[T] {
	return PriorityQueue[T]{xs: xs}
}

func (q PriorityQueue[T]) Sift(i int) int {
	return 0
}

func (q PriorityQueue[T]) Push(x T) PriorityQueue[T] {
	return PriorityQueue[T]{}
}

func (q PriorityQueue[T]) Pop(x T) PriorityQueue[T] {
	return PriorityQueue[T]{}
}

func (q PriorityQueue[T]) Peek(x T) maybe.Maybe[T] {
	return maybe.None[T]()
}

// * Get parent of index i
func parent(i number.NonZero) uint {
	return (uint(i) - 1) / 2
}
