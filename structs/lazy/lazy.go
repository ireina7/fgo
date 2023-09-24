package lazy

import (
	"github.com/ireina7/fgo/structs/option"
	"github.com/ireina7/fgo/types"
)

type LazyKind types.Kind

type Lazy[A any] struct {
	val     option.Option[A]
	produce func() A
}

func From[A any](f func() A) *Lazy[A] {
	return &Lazy[A]{
		val:     option.From[A](nil),
		produce: f,
	}
}

func (v *Lazy[A]) Get() A {
	if option.IsNone(v.val) {
		val := v.produce()
		v.val = option.From(&val)
	}
	return option.Get(v.val)
}
