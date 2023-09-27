package examples

import (
	"fmt"
	"testing"

	"github.com/ireina7/fgo/interfaces"
	"github.com/ireina7/fgo/structs/function"
	"github.com/ireina7/fgo/structs/number"
	"github.com/ireina7/fgo/structs/option"
	"github.com/ireina7/fgo/types"
)

type LazyFunctorKind types.Kind

type LazyFunctor[A, B any] struct {
	iter interfaces.Iterator[A]
	f    func(A) B
}

func (LazyFunctor[A, B]) Kind(LazyFunctorKind) {}
func (LazyFunctor[A, B]) ElemType(B)           {}

func (functor LazyFunctor[A, B]) Next() option.Option[B] {
	return option.Map(functor.iter.Next(), functor.f)
}

func (functor LazyFunctor[A, B]) Iter() interfaces.Iterator[B] {
	return functor
}

type LazyFunctorImplIterable[A, B any] struct{}

func (impl *LazyFunctorImplIterable[A, B]) Impl(
	functor types.HKT[LazyFunctorKind, B],
) interfaces.Iterable[B] {
	return functor.(LazyFunctor[A, B])
}

type LazyFunctorFromIterator[A, B any] struct{}

func (self *LazyFunctorFromIterator[A, B]) FromIter(
	iter interfaces.Iterator[B],
) types.HKT[LazyFunctorKind, B] {
	return LazyFunctor[B, B]{
		iter: iter,
		f:    function.Identity[B](),
	}
}

func TestLazyFunctor(t *testing.T) {
	functor := &interfaces.IterFunctor[LazyFunctorKind, int, int]{
		Implement:    &LazyFunctorImplIterable[int, int]{},
		FromIterator: &LazyFunctorFromIterator[int, int]{},
	}
	xs := LazyFunctor[int, int]{
		iter: number.NewEnumInt().Range(0, 7),
		f:    function.Identity[int](),
	}
	ys := functor.Fmap(xs, func(x int) int {
		t.Log("Processing square", x)
		return x * x
	})
	zs := functor.Fmap(ys, func(y int) int {
		t.Log("Processing add1", y)
		return y + 1
	})

	functorS := &interfaces.IterFunctor[LazyFunctorKind, int, string]{
		Implement:    &LazyFunctorImplIterable[int, int]{},
		FromIterator: &LazyFunctorFromIterator[int, string]{},
	}
	ss := functorS.Fmap(zs, func(z int) string {
		t.Log("Processing fmt", z)
		return fmt.Sprintf("{number:%d}", z)
	})

	interfaces.For[string](ss.(LazyFunctor[string, string]), func(x string) {
		t.Log(x)
	})
}
