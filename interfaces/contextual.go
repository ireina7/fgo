package interfaces

import "context"

type Contextual[A any] interface {
	Context(context.Context) A
}
