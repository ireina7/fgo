package maybe

import "github.com/ireina7/fgo/types"

type OptionFunctor[A, B any] func(Maybe[A], func(A) B) Maybe[B]

func MakeFunctor[A, B any](
	f func(Maybe[A], func(A) B) Maybe[B],
) OptionFunctor[A, B] {
	return OptionFunctor[A, B](f)
}

func (functor OptionFunctor[A, B]) Fmap(
	x types.HKT[MaybeKind, A],
	f func(A) B,
) types.HKT[MaybeKind, B] {

	return functor(x.(Maybe[A]), f)
}
