package interfaces

type Async[A any] interface {
	Async(func() A) <-chan A
}
