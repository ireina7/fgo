package interfaces

import (
	"fmt"
	"testing"

	"github.com/ireina7/fgo/structs/option"
	"github.com/ireina7/fgo/structs/slice"
	"github.com/ireina7/fgo/types"
)

func TestFunctor(t *testing.T) {
	i := 2
	x := option.From[int](&i)
	functor := option.MakeFunctor(option.Map[int, int])
	y := functor.Fmap(x, func(i int) int {
		return i * i
	})
	// t.Logf("%#v", y)
	ShowOption[int](y)
}

func TestFunctorSlice(t *testing.T) {
	xs := slice.From[int]([]int{1, 2, 3, 4, 5})
	ys := slice.Map(xs, func(x int) int {
		return x * x
	})
	t.Log(ys)
}

func ShowOption[A any](x option.Option[A]) {
	fmt.Printf("Showing %#v\n", x)
}

func ForEach[F_, A any](functor Functor[F_, A, types.Unit], fa types.HKT[F_, A], f func(A)) {
	functor.Fmap(fa, func(a A) types.Unit {
		f(a)
		return types.MakeUnit()
	})
}
