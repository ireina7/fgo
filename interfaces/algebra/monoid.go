package algebra

import (
	"github.com/ireina7/fgo/structs/maybe"
	"github.com/ireina7/fgo/types"
)

type Monoid[A any] interface {
	Semigroup[A]
	Empty() A
}

type HKTMonoid struct{}

func (m *HKTMonoid) Empty() types.HKT[maybe.MaybeKind, string] {
	return maybe.Some[string]("")
}

func (m *HKTMonoid) Combine(
	a, b types.HKT[maybe.MaybeKind, string],
) types.HKT[maybe.MaybeKind, string] {
	return maybe.Map[string, string](a.(maybe.Maybe[string]), func(a string) string {
		z := maybe.Map[string, string](b.(maybe.Maybe[string]), func(b string) string {
			return a + b
		})
		if z.IsNone() {
			return ""
		}
		return z.MustGet()
	})
}
