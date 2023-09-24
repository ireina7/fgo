package ordered

import "golang.org/x/exp/constraints"

type PreludeLessThan[A constraints.Ordered] struct{}

func (prelude *PreludeLessThan[A]) LessThan(a, b A) bool {
	return a < b
}
