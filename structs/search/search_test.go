package search

import (
	"testing"

	"github.com/ireina7/fgo/interfaces"
	"github.com/ireina7/fgo/structs/slice"
	"github.com/ireina7/fgo/types"
)

type sliceImpl[A any] struct{}

func (impl *sliceImpl[A]) To(xs types.HKT[slice.SliceKind, A]) interfaces.Iterable[A] {
	return xs.(slice.Slice[A])
}

func TestLinear(t *testing.T) {
	xs := slice.Make(1, 2, 3, 4, 5)
	find := &SearchLinear[slice.SliceKind, int]{
		Implement: &sliceImpl[int]{},
	}
	ans := find.SearchBy(xs, func(x int) bool {
		return x > 2
	})
	t.Log(*ans)
}
