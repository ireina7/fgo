package option

import "github.com/ireina7/fgo/types"

type OptionKind any

type Option[A any] types.HKT[OptionKind, A]

type None[A any] struct{}

func (None[A]) Kind(OptionKind) {}
func (None[A]) Elem(A)          {}

type Some[A any] struct {
	Value A
}

func (Some[A]) Kind(OptionKind) {}
func (Some[A]) Elem(A)          {}
