package interfaces

import (
	"github.com/ireina7/fgo/structs/tuple"
	"github.com/ireina7/fgo/types"
)

type Pure[F_, A any] interface {
	Pure(A) types.HKT[F_, A]
}

type Apply[F_, A, B any] interface {
	Product(types.HKT[F_, A], types.HKT[F_, B]) types.HKT[F_, tuple.Tuple2[A, B]]
}
