package interfaces

import "fmt"

type Debugger[A any] interface {
	Debug(A) string
}

type Debugging interface {
	Debug() string
}

type DefaultDebug[A any] struct{}

func (dbg *DefaultDebug[A]) Debug(a A) string {
	return fmt.Sprintf("%#v", a)
}
