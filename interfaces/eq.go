package interfaces

type Eq[A any] interface {
	Equal(A, A) bool
}
