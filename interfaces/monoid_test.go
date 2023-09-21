package interfaces

import (
	"testing"

	"github.com/ireina7/fgo/structs/option"
	"github.com/ireina7/fgo/types"
)

func TestMonoid(t *testing.T) {
	monoid := &HKTMonoid{}
	xs := option.Some[string]{Value: "1"}
	ys := option.Some[string]{Value: "2"}
	// ys := option.None[string]{}
	zs := foldl[option.OptionKind, string](monoid, ys, xs)
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
