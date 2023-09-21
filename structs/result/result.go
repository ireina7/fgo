package result

import "github.com/ireina7/fgo/types"

type ResultKind any

type Result[A any] types.HKT[ResultKind, A]

type Ok[A any] struct {
	Value A
}

func (Ok[A]) Kind(ResultKind) {}
func (Ok[A]) Elem(A)          {}

type Err[A any] struct {
	Error error
}

func (Err[A]) Kind(ResultKind) {}
func (Err[A]) Elem(A)          {}

func fmap[A, B any](res Result[A], f func(A) B) Result[B] {
	var ans Result[B]
	switch x := res.(type) {
	case Ok[A]:
		ans = Ok[B]{Value: f(x.Value)}
	case Err[A]:
		ans = Err[B]{Error: x.Error}
	}
	return ans
}
