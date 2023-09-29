package interfaces

import (
	"github.com/ireina7/fgo/interfaces/collection"
	"github.com/ireina7/fgo/interfaces/impl"
	"github.com/ireina7/fgo/structs/option"
	"github.com/ireina7/fgo/types"
)

type SeqInterface[F_, A any] interface {
	impl.Implement[types.HKT[F_, A], collection.Iterable[A]]

	Get(types.HKT[F_, A], int) option.Option[A]
	Len(types.HKT[F_, A]) int
}

type Seq[F_, A any] interface {
	SeqInterface[F_, A]
	Append(types.HKT[F_, A], A) types.HKT[F_, A]
	Concat(types.HKT[F_, A], types.HKT[F_, A]) types.HKT[F_, A]
	Count(types.HKT[F_, A], func(A) bool) int
	Find(types.HKT[F_, A], func(A) bool) option.Option[A]
}

type SeqDefault[F_, A any] struct {
	SeqInterface[F_, A]
}

func (dsl *SeqDefault[F_, A]) Append(types.HKT[F_, A], A) types.HKT[F_, A] {
	return nil
}
