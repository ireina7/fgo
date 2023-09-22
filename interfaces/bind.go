package interfaces

import "github.com/ireina7/fgo/types"

type Bind[F_, A, B any] interface {
	FlatMap(types.HKT[F_, A], func(A) types.HKT[F_, B]) types.HKT[F_, B]
}
