package interfaces

import "github.com/ireina7/fgo/types"

type Applicative[F_, A, B any] interface {
	Apply(types.HKT[F_, func(A) B], types.HKT[F_, A]) types.HKT[F_, B]
}
