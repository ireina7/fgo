package interfaces

import "github.com/ireina7/fgo/types"

type Traversable[F_, G_, A, B any] interface {
	Traverse(
		func(A) types.HKT[F_, B],
		types.HKT[G_, A],
		types.HKT[F_, types.HKT[G_, B]],
	)
}
