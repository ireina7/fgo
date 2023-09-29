package interfaces

import "github.com/ireina7/fgo/interfaces/iter"

type Enum[A any] interface {
	Succ(A) A
	Pred(A) A
	ToEnum(int) A
	Range(A, A) iter.Iterator[A]
}
