package examples

import (
	"fmt"
	"sync"
	"testing"

	"github.com/ireina7/fgo/structs/slice"
)

type concurrentSlice[A, B any] struct{}

func (s *concurrentSlice[A, B]) fmap(xs slice.Slice[A], f func(A) B) slice.Slice[B] {
	ys := slice.From(make([]B, slice.Len(xs)))
	var wg sync.WaitGroup
	slice.ForEach(xs, func(i int, x A) {
		wg.Add(1)
		go func() {
			defer wg.Done()
			y := f(x)
			ys = slice.Set(ys, i, y)
		}()
	})
	wg.Wait()
	return ys
}

func TestMe(t *testing.T) {
	functor := &concurrentSlice[int, int]{}
	xs := slice.From([]int{1, 2, 3, 4, 5, 6, 7})
	ys := functor.fmap(xs, func(x int) int {
		fmt.Println("Processing", x)
		return x * x
	})
	t.Log(ys)
}
