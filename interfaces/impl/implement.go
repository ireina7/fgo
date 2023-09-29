package impl

type Implement[A, B any] interface {
	Impl(A) B
}
