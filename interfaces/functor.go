package interfaces

import (
	"github.com/ireina7/fgo/types"
)

type Functor[F_, A, B any] interface {
	Fmap(types.HKT[F_, A], func(A) B) types.HKT[F_, B]
}
