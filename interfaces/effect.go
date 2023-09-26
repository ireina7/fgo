package interfaces

type Exception interface {
	Throw(error)
}

type Coroutine[A, B any] interface {
	Resume(A) B
	Return(B)
}

type Effect[A, B any] interface {
	Coroutine[A, B]
	Exception
	Do(func()) A
}

type chanEffectHandler[A, B any] struct {
	resumeCh chan A
	throwCh  chan error
	returnCh chan B
}

func NewDefaultEffect[A, B any]() *chanEffectHandler[A, B] {
	return &chanEffectHandler[A, B]{
		resumeCh: make(chan A),
		returnCh: make(chan B),
		throwCh:  make(chan error),
	}
}

func (eff *chanEffectHandler[A, B]) Do(f func()) A {
	go f()
	select {
	case a := <-eff.resumeCh:
		return a
	case err := <-eff.throwCh:
		panic(err)
	}
}

func (eff *chanEffectHandler[A, B]) Resume(a A) B {
	eff.resumeCh <- a
	b := <-eff.returnCh
	return b
}

func (eff *chanEffectHandler[A, B]) Throw(err error) {
	eff.throwCh <- err
}

func (eff *chanEffectHandler[A, B]) Return(b B) {
	eff.returnCh <- b
}

type EffMonad[A any] struct {
}
