package interfaces

type Implement[A, B any] interface {
	Impl(A) B
}
