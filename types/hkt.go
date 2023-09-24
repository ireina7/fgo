package types

type Kind interface {
	// Apply()
}

type KindType[F_ any] interface {
	Kind(F_)
}

type ElemType[T any] interface {
	ElemType(T)
}

type HKT[F_ Kind, A any] interface {
	KindType[F_]
	ElemType[A]
}
