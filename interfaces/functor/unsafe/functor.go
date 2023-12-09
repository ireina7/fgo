package unsafe

import (
	"github.com/ireina7/fgo/structs/slice"
	"github.com/ireina7/fgo/types"
)

type UnsafeSliceFunctor struct{}

func (*UnsafeSliceFunctor) Fmap(
	ma types.HKT[slice.SliceKind, any], f func(any) any,
) types.HKT[slice.SliceKind, any] {

	xs := ma.(slice.Slice[any])
	ys := slice.Room[any](xs.Len())
	for i := 0; i < xs.Len(); i++ {
		ys.Set(i, f(xs[i]))
	}
	return ys
}
