package search

import (
	"testing"

	"github.com/ireina7/fgo/interfaces/collection"
	"github.com/ireina7/fgo/structs/option"
	"github.com/ireina7/fgo/structs/ref"
	"github.com/ireina7/fgo/structs/slice"
	"github.com/ireina7/fgo/types"
)

type sliceImpl[A any] struct{}

func (impl *sliceImpl[A]) Impl(xs types.HKT[slice.SliceKind, A]) collection.Iterable[A] {
	return xs.(slice.Slice[A])
}

func TestLinear(t *testing.T) {
	xs := slice.Make(1, 2, 3, 4, 5)
	find := &SearchLinear[slice.SliceKind, int]{
		Implement: &sliceImpl[int]{},
	}
	ans := find.SearchBy(xs, func(x int) bool {
		return x == -1
	})
	t.Logf("%#v", option.Map(ans, func(ref ref.Ref[int]) int {
		return ref.Deref()
	}))
}
