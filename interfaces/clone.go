package interfaces

type Clone[A any] interface {
	Clone(A) A
}
