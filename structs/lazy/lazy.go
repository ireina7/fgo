package lazy

import (
	"github.com/ireina7/fgo/structs/maybe"
	"github.com/ireina7/fgo/types"
)

type LazyKind types.Kind

type Lazy[A any] struct {
	val     maybe.Maybe[A]
	produce func() A
}

func From[A any](f func() A) *Lazy[A] {
	return &Lazy[A]{
		val:     maybe.From[A](nil),
		produce: f,
	}
}

func (v *Lazy[A]) Get() A {
	if maybe.IsNone(v.val) {
		val := v.produce()
		v.val = maybe.From(&val)
	}
	return maybe.Get(v.val)
}
