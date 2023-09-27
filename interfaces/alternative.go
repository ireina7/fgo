package interfaces

import "github.com/ireina7/fgo/types"

type Alternative[F_, A any] interface {
	Empty() types.HKT[F_, A]
	Or(types.HKT[F_, A], types.HKT[F_, A]) types.HKT[F_, A]
}
