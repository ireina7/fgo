package result

import "github.com/ireina7/fgo/types"

type ResultKind any

type Result[A any] types.HKT[ResultKind, A]

type Ok[A any] struct {
	Value A
}

func (Ok[A]) Kind(ResultKind) {}
func (Ok[A]) ElemType(A)      {}

type Err[A any] struct {
	Error error
}

func (Err[A]) Kind(ResultKind) {}
func (Err[A]) ElemType(A)      {}

func From[A any](x A) Result[A] {
	return Ok[A]{Value: x}
}

func FromErr[A any](err error) Result[A] {
	return Err[A]{Error: err}
}

func Map[A, B any](res Result[A], f func(A) B) Result[B] {
	var ans Result[B]
	switch x := res.(type) {
	case Ok[A]:
		ans = Ok[B]{Value: f(x.Value)}
	case Err[A]:
		ans = Err[B]{Error: x.Error}
	}
	return ans
}

func MapErr[A any](res Result[A], f func(error) error) Result[A] {
	var ans Result[A]
	switch x := res.(type) {
	case Ok[A]:
		ans = Ok[A]{Value: x.Value}
	case Err[A]:
		ans = Err[A]{Error: f(x.Error)}
	}
	return ans
}
