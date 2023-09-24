package ordered

import "golang.org/x/exp/constraints"

type LessThan[A any] interface {
	LessThan(A, A) bool
}

type MinMax[A any] struct {
	Less LessThan[A]
}

func (MinMax[A]) Make(less LessThan[A]) *MinMax[A] {
	return &MinMax[A]{Less: less}
}

func PreludeMinMax[A constraints.Ordered]() *MinMax[A] {
	return MinMax[A]{}.Make(&PreludeLessThan[A]{})
}

func (self *MinMax[A]) Min(a, b A) A {
	if self.Less.LessThan(a, b) {
		return a
	}
	return b
}

func (self *MinMax[A]) Max(a, b A) A {
	if self.Less.LessThan(a, b) {
		return b
	}
	return a
}
