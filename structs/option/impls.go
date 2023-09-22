package option

import "github.com/ireina7/fgo/types"

type OptionFunctor[A, B any] func(Option[A], func(A) B) Option[B]

func MakeFunctor[A, B any](
	f func(Option[A], func(A) B) Option[B],
) OptionFunctor[A, B] {
	return OptionFunctor[A, B](f)
}

func (functor OptionFunctor[A, B]) Fmap(
	x types.HKT[OptionKind, A],
	f func(A) B,
) types.HKT[OptionKind, B] {

	return functor(x, f)
}
