package search

import (
	"github.com/ireina7/fgo/interfaces"
	"github.com/ireina7/fgo/structs/option"
	"github.com/ireina7/fgo/types"
)

type Searching[F_, A any] interface {
	SearchBy(types.HKT[F_, A], func(A) bool) *A
}

type SearchLinear[F_, A any] struct {
	interfaces.Implement[types.HKT[F_, A], interfaces.Iterable[A]]
}

func (s *SearchLinear[F_, A]) SearchBy(
	xs types.HKT[F_, A], hit func(A) bool,
) *A {
	var ans *A
	ys := s.To(xs)
	iter := ys.Iter()
	for x := iter.Next(); !option.IsNone(x); x = iter.Next() {
		found := false
		option.Map_(x, func(x A) {
			if hit(x) {
				ans = &x
				found = true
			}
		})
		if found {
			break
		}
	}
	return ans
}
