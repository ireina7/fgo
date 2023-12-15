package hkt

import "github.com/ireina7/fgo/types"

type Box[F_ types.Kind, A any] interface {
	Boxed(types.HKT[F_, A]) types.HKT[F_, any]
}

type Unbox[F_ types.Kind, A any] interface {
	Unboxed(types.HKT[F_, any]) *types.HKT[F_, A] // may be nil
}

type Boxing[F_ types.Kind, A any] interface {
	Box[F_, A]
	Unbox[F_, A]
}

type Pipe[F_ types.Kind, A, B any] interface {
	Box[F_, A]
	Unbox[F_, B]
}
