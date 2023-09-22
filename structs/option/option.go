package option

import "github.com/ireina7/fgo/types"

type OptionKind any

type Option[A any] types.HKT[OptionKind, A]

type None[A any] struct{}

func (None[A]) Kind(OptionKind) {}
func (None[A]) ElemType(A)      {}

type Some[A any] struct {
	Value A
}

func (Some[A]) Kind(OptionKind) {}
func (Some[A]) ElemType(A)      {}

func From[A any](x *A) Option[A] {
	if x == nil {
		return None[A]{}
	}
	return Some[A]{Value: *x}
}

func Map[A, B any](fa Option[A], f func(A) B) Option[B] {
	var y B
	switch x := fa.(type) {
	case Some[A]:
		y = f(x.Value)
	}
	return From(&y)
}
