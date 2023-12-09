package channel

type Chan[A any] chan A

type Send[A any] chan<- A
type Recv[A any] <-chan A

type ChanImplChanSend[A any] struct{}

func (impl *ChanImplChanSend[A]) Impl(ch Chan[A]) Send[A] {
	var chSend chan<- A = ch
	return Send[A](chSend)
}

type ChanImplChanRecv[A any] struct{}

func (impl *ChanImplChanRecv[A]) Impl(ch Chan[A]) Recv[A] {
	var chRecv <-chan A = ch
	return Recv[A](chRecv)
}

func From[A any](ch chan A) Chan[A] {
	return Chan[A](ch)
}

func (ch Chan[A]) Send(a A) Chan[A] {
	ch <- a
	return ch
}
