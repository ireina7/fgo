package interfaces

type Enum[A any] interface {
	Succ(A) A
	Pred(A) A
	ToEnum(int) A
	Range(A, A) []A
}
