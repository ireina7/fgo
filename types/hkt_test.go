package types_test

import (
	"testing"

	option "github.com/ireina7/fgo/structs/maybe"
	"github.com/ireina7/fgo/structs/result"
	"github.com/ireina7/fgo/types"
	"github.com/ireina7/fgo/types/hkt"
)

func TestWrappedHKT(t *testing.T) {
	var xs types.HKT[result.ResultKind, types.HKT[option.MaybeKind, int]]
	xs = result.From[types.HKT[option.MaybeKind, int]](option.Some(7))
	t.Logf("%#v", xs.(result.Result[types.HKT[option.MaybeKind, int]]))
}

type Functor[F_, A, B any] interface {
	Fmap(types.HKT[F_, A], func(A) B) types.HKT[F_, B]
}

type RawFunctor[F_ any] interface {
	Fmap(types.HKT[F_, any], func(any) any) types.HKT[F_, any]
}

type FromUnsafe[F_, A, B any] struct {
	hkt.Pipe[F_, A, B]
	unsafe RawFunctor[F_]
}

func (self FromUnsafe[F_, A, B]) Fmap(ma types.HKT[F_, A], f func(A) B) types.HKT[F_, B] {
	mb := self.unsafe.Fmap(self.Boxed(ma), func(x any) any {
		return f(x.(A))
	})
	return *self.Unboxed(mb)
}
