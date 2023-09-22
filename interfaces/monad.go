package interfaces

type Monad[F_, A, B any] interface {
	Applicative[F_, A]
	Bind[F_, A, B]
}
