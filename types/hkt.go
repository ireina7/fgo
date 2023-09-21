package types

type HKT[F_, A any] interface {
	Kind(F_)
	Elem(A)
}
