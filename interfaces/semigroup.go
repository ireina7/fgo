package interfaces

type Semigroup[A any] interface {
	Combine(A, A) A
}
