package util

import "fmt"

func TODO[A any](msg string, args ...any) A {
	panic(fmt.Errorf(msg, args...))
}

func Use(...any) {}
