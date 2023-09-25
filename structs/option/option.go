package option

import (
	"github.com/ireina7/fgo/types"
)

type OptionKind types.Kind

type Option[A any] types.HKT[OptionKind, A]

type None[A any] struct{}

func (None[A]) Kind(OptionKind) {}
func (None[A]) ElemType(A)      {}

type Some[A any] struct {
	Value A
}

func (Some[A]) Kind(OptionKind) {}
func (Some[A]) ElemType(A)      {}

func From[A any](x *A) Option[A] {
	if x == nil {
		return None[A](struct{}{})
	}
	return Some[A]{Value: *x}
}

func IsNone[A any](x Option[A]) bool {
	switch x.(type) {
	case None[A]:
		return true
	}
	return false
}

func Get[A any](x Option[A]) A {
	return x.(Some[A]).Value
}

func Map[A, B any](fa Option[A], f func(A) B) Option[B] {
	var y B
	switch x := fa.(type) {
	case Some[A]:
		y = f(x.Value)
		return From(&y)
	default:
		return From[B](nil)
	}
}

func Map_[A any](fa Option[A], f func(A)) {
	switch x := fa.(type) {
	case Some[A]:
		f(x.Value)
	default:
	}
}

func For[A any](fa Option[A], f func(A)) {
	Map_(fa, f)
}

func FlatMap[A, B any](ma Option[A], f func(A) Option[B]) Option[B] {
	if IsNone(ma) {
		return None[B]{}
	}
	a := Get(ma)
	return f(a)
}
