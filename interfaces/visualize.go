package interfaces

import "github.com/ireina7/fgo/types"

type Visualize[F_, A any, M types.HKT[F_, A]] struct {
}

func (v *Visualize[F_, A, M]) Visualize(xs types.HKT[F_, A]) M {
	return xs.(M)
}
