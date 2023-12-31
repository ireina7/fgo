package search

import (
	"github.com/ireina7/fgo/interfaces/collection"
	"github.com/ireina7/fgo/interfaces/impl"
	"github.com/ireina7/fgo/structs/maybe"
	"github.com/ireina7/fgo/structs/ref"
	"github.com/ireina7/fgo/types"
)

type Searching[F_, A any] interface {
	SearchBy(types.HKT[F_, A], func(A) bool) maybe.Maybe[ref.Ref[A]]
}

type SearchLinear[F_, A any] struct {
	impl.Implement[types.HKT[F_, A], collection.Iterable[A]]
}

func (s *SearchLinear[F_, A]) SearchBy(
	xs types.HKT[F_, A], hit func(A) bool,
) maybe.Maybe[ref.Ref[A]] {
	ans := maybe.From[ref.Ref[A]](nil)
	ys := s.Impl(xs)
	iter := ys.Iter()
	for x := iter.Next(); !maybe.IsNone(x); x = iter.Next() {
		found := false
		maybe.Map_(x, func(x A) {
			if hit(x) {
				r := ref.Refer(x)
				ans = maybe.From[ref.Ref[A]](&r)
				found = true
			}
		})
		if found {
			return ans
		}
	}
	return ans
}
