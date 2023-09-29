package interfaces

import "github.com/ireina7/fgo/interfaces/collection"

type Enum[A any] interface {
	Succ(A) A
	Pred(A) A
	ToEnum(int) A
	Range(A, A) collection.Iterator[A]
}
