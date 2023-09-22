package types

type Kind[F_ any] interface {
	Kind(F_)
}

type ElemType[T any] interface {
	ElemType(T)
}

type HKT[F_, A any] interface {
	Kind[F_]
	ElemType[A]
}
