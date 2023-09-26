package interfaces

type Eq[A any] interface {
	Equal(A, A) bool
}

type Prelude[T comparable] struct{}

func (eq *Prelude[T]) Equal(a, b T) bool {
	return a == b
}

type ToComparable[A any, B comparable] interface {
	ToComparable(A) B
}
