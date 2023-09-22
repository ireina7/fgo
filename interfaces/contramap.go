package interfaces

import "github.com/ireina7/fgo/types"

type ContraFunctor[F_, A, B any] interface {
	ContraMap(types.HKT[F_, A], func(B) A) types.HKT[F_, B]
}
