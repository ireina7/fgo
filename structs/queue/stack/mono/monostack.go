package stack

type MonoStack[T any] struct {
	xs []T
}

func Empty[T any]() MonoStack[T] {
	return MonoStack[T]{}
}
