package interfaces

import "github.com/ireina7/fgo/types"

type Applicative[F_, A any] interface {
	Pure(A) types.HKT[F_, A]
}
