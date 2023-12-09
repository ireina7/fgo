package algebra

import (
	"testing"

	"github.com/ireina7/fgo/structs/maybe"
	"github.com/ireina7/fgo/types"
)

func TestMonoid(t *testing.T) {
	monoid := &HKTMonoid{}
	xs := maybe.Some[string]("1")
	ys := maybe.Some[string]("2")
	// ys := option.None[string]{}
	zs := foldl[maybe.MaybeKind, string](monoid, ys, xs)
	t.Logf("%#v", zs)
}

func foldl[F_, A any](
	monoid Monoid[types.HKT[F_, A]],
	xs ...types.HKT[F_, A],
) types.HKT[F_, A] {
	zs := monoid.Empty()
	for _, ys := range xs {
		zs = monoid.Combine(zs, ys)
	}
	return zs
}
