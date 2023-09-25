package interfaces

import "fmt"

type Debugging[A any] interface {
	Debug(A) string
}

type DefaultDebug[A any] struct{}

func (dbg *DefaultDebug[A]) Debug(a A) string {
	return fmt.Sprintf("%#v", a)
}
