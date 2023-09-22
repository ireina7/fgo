package interfaces

import (
	"github.com/ireina7/fgo/structs/option"
	"github.com/ireina7/fgo/types"
)

type Monoid[A any] interface {
	Semigroup[A]
	Empty() A
}

type HKTMonoid struct{}

func (m *HKTMonoid) Empty() types.HKT[option.OptionKind, string] {
	return option.Some[string]{Value: ""}
}

func (m *HKTMonoid) Combine(
	a, b types.HKT[option.OptionKind, string],
) types.HKT[option.OptionKind, string] {
	return option.Map[string, string](a, func(a string) string {
		z := option.Map[string, string](b, func(b string) string {
			return a + b
		})
		switch x := z.(type) {
		case option.Some[string]:
			return x.Value
		default:
			return ""
		}
	})
}
