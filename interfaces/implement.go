package interfaces

type Implement[A, B any] interface {
	To(A) B
}
