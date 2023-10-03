package interfaces

type Monad[F_, A, B any] interface {
	Apply[F_, A, B]
	Bind[F_, A, B]
}
