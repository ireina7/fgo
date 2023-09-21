package interfaces

import (
	"github.com/ireina7/fgo/structs/option"
	"github.com/ireina7/fgo/types"
)

type Monoid[A any] interface {
	Empty() A
	Combine(A, A) A
}

type HKTMonoid struct{}

func (m *HKTMonoid) Empty() types.HKT[option.OptionKind, string] {
	return option.Some[string]{Value: ""}
}

func (m *HKTMonoid) Combine(
	a, b types.HKT[option.OptionKind, string],
) types.HKT[option.OptionKind, string] {
	return mapOption[string, string](a, func(a string) string {
		z := mapOption[string, string](b, func(b string) string {
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

func mapOption[A, B any](fa option.Option[A], f func(A) B) option.Option[B] {
	switch x := fa.(type) {
	case option.Some[A]:
		return option.Some[B]{Value: f(x.Value)}
	default:
		return option.None[B](struct{}{})
	}
}
