package interfaces

import "github.com/ireina7/fgo/types"

type Turn[F_, A, B any] interface {
	Turn(types.HKT[F_, A]) types.HKT[F_, B]
}
