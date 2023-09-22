package interfaces

type Ordering int

const (
	LessThan  Ordering = -1
	Equal     Ordering = 0
	GreatThan Ordering = 1
)

type Order[A any] interface {
	Compare(A, A) Ordering
}
