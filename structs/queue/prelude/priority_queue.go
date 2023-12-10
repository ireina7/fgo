package prelude

import (
	"cmp"

	"github.com/ireina7/fgo/structs/maybe"
	"github.com/ireina7/fgo/structs/number"
	"github.com/ireina7/fgo/structs/number/checked"
	"github.com/ireina7/fgo/structs/number/nonzero"
	"github.com/ireina7/fgo/types"
)

// This will be a max-heap
type PriorityQueue[T cmp.Ordered] struct {
	xs []T
}

type PriorityQueueKind types.Kind

func (PriorityQueue[T]) KindType()  {}
func (PriorityQueue[T]) ElemType(T) {}

func Empty[T cmp.Ordered]() PriorityQueue[T] {
	return PriorityQueue[T]{}
}

func Room[T cmp.Ordered](size uint) PriorityQueue[T] {
	return PriorityQueue[T]{
		xs: make([]T, 0, size),
	}
}

func (q PriorityQueue[T]) Len() uint {
	return uint(len(q.xs))
}

func (q PriorityQueue[T]) at(i uint) maybe.Maybe[T] {
	if i >= q.Len() {
		return maybe.None[T]()
	}
	return maybe.Some(q.xs[i])
}

func FromSlice[T cmp.Ordered](xs []T) PriorityQueue[T] {
	q := PriorityQueue[T]{xs: xs}
	n := q.Len() - 1
	i := maybe.Map(nonzero.From(n), func(i number.NonZero) uint {
		return parent(i)
	})
	for i.IsSome() {
		q = q.Heapify(i.MustGet())
		i = checked.MinusUnsigned(i.MustGet(), 1)
	}
	return q
}

func (q PriorityQueue[T]) Heapify(i uint) PriorityQueue[T] {
	xs := q.xs
	n := q.Len()
	l := leftChild(i)
	r := rightChild(i)

	// Find the largest index
	largest := i
	if l < n && xs[l] > xs[i] {
		largest = l
	}
	if r < n && xs[r] > xs[largest] {
		largest = r
	}
	if largest != i {
		xs[i], xs[largest] = xs[largest], xs[i]
		return q.Heapify(largest)
	}
	return q
}

func (q PriorityQueue[T]) Lift(i uint) PriorityQueue[T] {
	if i == 0 {
		return q
	}
	xs := q.xs
	j := parent(nonzero.From(i).MustGet())
	for xs[j] < xs[i] {
		xs[j], xs[i] = xs[i], xs[j]
		if j == 0 {
			break
		}
		i = j
		j = parent(nonzero.From(i).MustGet())
	}
	return PriorityQueue[T]{xs: xs}
}

// This is just like slice, changes result to a new queue
func (q PriorityQueue[T]) Push(x T) PriorityQueue[T] {
	n := q.Len()
	xs := q.xs
	xs = append(xs, x)
	return PriorityQueue[T]{xs}.Lift(n)
}

func (q PriorityQueue[T]) Pop() (maybe.Maybe[T], PriorityQueue[T]) {
	ans := q.Peek()
	if ans.IsNone() {
		return ans, q
	}
	xs := q.xs
	n := q.Len()
	// swap
	xs[0] = xs[n-1]
	// heapify
	q = q.Heapify(0)
	return ans, PriorityQueue[T]{xs: xs[:n-1]}
}

func (q PriorityQueue[T]) Peek() maybe.Maybe[T] {
	if q.Len() == 0 {
		return maybe.None[T]()
	}
	return maybe.Some(q.xs[0])
}

// * Get parent of index i
func parent(i number.NonZero) uint {
	return (uint(i) - 1) / 2
}

func leftChild(i uint) uint {
	return i*2 + 1
}

func rightChild(i uint) uint {
	return i*2 + 2
}
