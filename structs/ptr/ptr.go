package ptr

import (
	"github.com/ireina7/fgo/structs/option"
	"github.com/ireina7/fgo/types"
)

type PtrKind types.Kind

type Ptr[A any] struct {
	ptr *A
}

func Point[A any](a A) Ptr[A] {
	return Ptr[A]{ptr: &a}
}

func From[A any](pa *A) Ptr[A] {
	return Ptr[A]{ptr: pa}
}

func (Ptr[A]) Kind(PtrKind) {}
func (Ptr[A]) ElemType(A)   {}

func (ptr Ptr[A]) IsNil() bool {
	return ptr.ptr == nil
}

func (ptr Ptr[A]) Deref() option.Option[A] {
	if ptr.IsNil() {
		return option.From[A](nil)
	}
	return option.From[A](ptr.ptr)
}

func Map[A, B any](ptr Ptr[A], f func(A) B) Ptr[B] {
	if ptr.IsNil() {
		return From[B](nil)
	}
	b := f(*ptr.ptr)
	return Point(b)
}
