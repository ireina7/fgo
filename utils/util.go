package utils

import "fmt"

func TODO[A any](msg string, args ...any) A {
	panic(fmt.Errorf(msg, args...))
}

func Use(...any) {}

func Assert(cond bool) {
	if !cond {
		panic(fmt.Errorf("assertion failed"))
	}
}
