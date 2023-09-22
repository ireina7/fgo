package interfaces

import (
	"github.com/ireina7/fgo/structs/option"
	"github.com/ireina7/fgo/types"
)

type Functor[F_, A, B any] interface {
	Fmap(types.HKT[F_, A], func(A) B) types.HKT[F_, B]
}

type OptionFunctor[A, B any] func(option.Option[A], func(A) B) option.Option[B]

func Make[A, B any](
	f func(option.Option[A], func(A) B) option.Option[B],
) OptionFunctor[A, B] {
	return OptionFunctor[A, B](f)
}

func (functor OptionFunctor[A, B]) Fmap(
	x types.HKT[option.OptionKind, A],
	f func(A) B,
) types.HKT[option.OptionKind, B] {

	return functor(x, f)
}
