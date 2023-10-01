package result

import "github.com/ireina7/fgo/types"

type ResultKind types.Kind

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

func IsOK[A any](res Result[A]) bool {
	_, ok := res.(Ok[A])
	return ok
}

func IsErr[A any](res Result[A]) bool {
	return !IsOK(res)
}

func Get[A any](res Result[A]) A {
	return res.(Ok[A]).Value
}

func GetErr[A any](res Result[A]) error {
	return res.(Err[A]).Error
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

func FlatMap[A, B any](res Result[A], f func(A) Result[B]) Result[B] {
	if IsErr(res) {
		return FromErr[B](res.(Err[A]).Error)
	}
	return f(Get(res))
}

func AndThen[A, B any](res Result[A], f func(A) Result[B]) Result[B] {
	return FlatMap(res, f)
}

func Unpack[A any](res Result[A]) (A, error) {
	var a A
	var err error
	switch x := res.(type) {
	case Ok[A]:
		a = x.Value
	case Err[A]:
		err = x.Error
	default:
	}
	return a, err
}

type resultFunctor[A, B any] struct{}

func (functor *resultFunctor[A, B]) Fmap(
	xs types.HKT[ResultKind, A],
	f func(A) B,
) types.HKT[ResultKind, B] {

	return Map(xs.(Result[A]), f)
}
