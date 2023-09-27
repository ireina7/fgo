package interfaces

import "github.com/ireina7/fgo/types"

// * Natural Transformation
type Transformation[F_, G_, A any] interface {
	Transform(types.HKT[F_, A]) types.HKT[G_, A]
}
