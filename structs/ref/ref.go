package ref

import (
	"github.com/ireina7/fgo/structs/maybe"
	"github.com/ireina7/fgo/types"
	"github.com/ireina7/fgo/utils"
)

type RefKind types.Kind

// * Reference type
// Ensure not nil
type ptr[A any] *A

type Ref[A any] struct {
	ptr ptr[A]
}

func (Ref[A]) Kind(RefKind) {}
func (Ref[A]) ElemType(A)   {}

func Refer[A any](a A) Ref[A] {
	return Ref[A]{ptr: &a}
}

func Make[A any](a A) Ref[A] {
	return Refer(a)
}

func MustFromPtr[A any](p *A) Ref[A] {
	utils.Assert(p != nil)
	return Ref[A]{ptr: p}
}

func FromPtr[A any](p *A) maybe.Maybe[Ref[A]] {
	if p == nil {
		return maybe.None[Ref[A]]()
	}
	return maybe.Some(Ref[A]{ptr: p})
}

func (ref Ref[A]) Deref() A {
	return *ref.ptr
}

func (ref Ref[A]) Modify(modify func(*A)) {
	modify(ref.ptr)
}

func (ref Ref[A]) For(f func(A)) {
	f(ref.Deref())
}

func (ref Ref[A]) UnsafePtr() *A {
	return ref.ptr
}

func GetBy[A, B any](ref Ref[A], f func(A) B) maybe.Maybe[B] {
	return func() (ans maybe.Maybe[B]) {
		defer func() {
			if r := recover(); r != nil {
				ans = maybe.None[B]()
			}
		}()
		return maybe.Some(f(ref.Deref()))
	}()
}

func MustGetBy[A, B any](ref Ref[A], f func(A) B) B {
	return f(ref.Deref())
}

func Map[A, B any](ref Ref[A], f func(A) B) Ref[B] {
	b := f(ref.Deref())
	return Refer(b)
}

func FlatMap[A, B any](ref Ref[A], f func(A) Ref[B]) Ref[B] {
	return f(ref.Deref())
}
