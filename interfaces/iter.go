package interfaces

import (
	"github.com/ireina7/fgo/structs/option"
)

type Iterator[A any] interface {
	Next() option.Option[A] //Stateful interface!
}

type Iterable[A any] interface {
	Iter() Iterator[A]
}

func For[M Iterable[A], A any](xs M, f func(A)) {
	iter := xs.Iter()
	for x := iter.Next(); !option.IsNone(x); x = iter.Next() {
		option.Map_(x, func(x A) {
			f(x)
		})
	}
}

// for x := range Range(customContainer) {}
func Range[M Iterable[A], A any](xs M) <-chan A {
	ch := make(chan A)
	iter := xs.Iter()
	go func() {
		for x := iter.Next(); !option.IsNone(x); x = iter.Next() {
			option.Map_(x, func(x A) {
				ch <- x
			})
		}
	}()
	return ch
}
