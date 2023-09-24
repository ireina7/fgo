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

func For[A any](xs Iterable[A], f func(A)) {
	iter := xs.Iter()
	for x := iter.Next(); !option.IsNone(x); x = iter.Next() {
		option.Map_(x, func(x A) {
			f(x)
		})
	}
}

// func Map[A, B any](xs Iterable[A], f func(A) B) Iterable[B] {

// }

// for x := range Range(customContainer) {}
func Range[A any](xs Iterable[A]) <-chan A {
	ch := make(chan A)
	iter := xs.Iter()
	go func() {
		defer close(ch)
		for x := iter.Next(); !option.IsNone(x); x = iter.Next() {
			option.Map_(x, func(x A) {
				ch <- x
			})
		}
	}()
	return ch
}
