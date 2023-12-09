package maybe

import (
	"reflect"

	"github.com/ireina7/fgo/types"
)

type MaybeKind types.Kind

type Maybe[A any] struct {
	val *A
}

func (Maybe[A]) Kind(MaybeKind) {}
func (Maybe[A]) ElemType(A)     {}

func (o Maybe[A]) IsNone() bool {
	return o.val == nil
}

func (o Maybe[A]) IsSome() bool {
	return o.val != nil
}

func (o Maybe[A]) MustGet() A {
	return *o.val
}

func From[A any](x *A) Maybe[A] {
	if x == nil {
		return None[A]()
	}
	return Some(*x)
}

func FromAny(x any) Maybe[any] {
	if IsNil(x) {
		return None[any]()
	}
	return Some(x)
}

func None[A any]() Maybe[A] {
	return Maybe[A]{}
}

func Some[A any](a A) Maybe[A] {
	return Maybe[A]{val: &a}
}

func IsNone[A any](x Maybe[A]) bool {
	return x.IsNone()
}

func IsSome[A any](x Maybe[A]) bool {
	return x.IsSome()
}

func Get[A any](x Maybe[A]) A {
	return x.MustGet()
}

func Map[A, B any](fa Maybe[A], f func(A) B) Maybe[B] {
	if IsNone(fa) {
		return None[B]()
	}
	return Some(f(Get(fa)))
}

func Map_[A any](fa Maybe[A], f func(A)) {
	if IsNone(fa) {
		return
	}
	f(Get(fa))
}

func For[A any](fa Maybe[A], f func(A)) {
	Map_(fa, f)
}

func FlatMap[A, B any](ma Maybe[A], f func(A) Maybe[B]) Maybe[B] {
	if IsNone(ma) {
		return None[B]()
	}
	a := Get(ma)
	return f(a)
}

// Assert nil without comparing underlying types
func IsNil(x any) bool {
	if x == nil {
		return true
	}
	return reflect.ValueOf(x).Kind() == reflect.Ptr && reflect.ValueOf(x).IsNil()
}
