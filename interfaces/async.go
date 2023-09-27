package interfaces

type Async[A any] interface {
	Async(func() A) <-chan A
}

type LocalAsync[A any] struct{}

func (async *LocalAsync[A]) Async(f func() A) <-chan A {
	ch := make(chan A)
	go func() {
		ch <- f()
	}()
	return ch
}
