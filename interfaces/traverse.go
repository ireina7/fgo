package interfaces

import (
	"github.com/ireina7/fgo/structs/function"
	"github.com/ireina7/fgo/types"
)

type Traversable[F_, G_, A, B any] interface {
	// Apply[F_, A, B]
	Traverse(
		types.HKT[G_, A],
		func(A) types.HKT[F_, B],
	) types.HKT[F_, types.HKT[G_, B]]
}

type Sequence[F_, G_, B any] struct {
	Traversable[F_, G_, types.HKT[F_, B], B]
}

func (dsl *Sequence[F_, G_, B]) Sequence(
	container types.HKT[G_, types.HKT[F_, B]],
) types.HKT[F_, types.HKT[G_, B]] {

	return dsl.Traverse(
		container,
		function.Identity[types.HKT[F_, B]](),
	)
}
