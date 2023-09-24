package channel

type Chan[A any] chan A

type ChanSend[A any] chan<- A
type ChanRecv[A any] <-chan A

type ChanImplChanSend[A any] struct{}

func (impl *ChanImplChanSend[A]) Impl(ch Chan[A]) ChanSend[A] {
	var chSend chan<- A = ch
	return ChanSend[A](chSend)
}

type ChanImplChanRecv[A any] struct{}

func (impl *ChanImplChanRecv[A]) Impl(ch Chan[A]) ChanRecv[A] {
	var chRecv <-chan A = ch
	return ChanRecv[A](chRecv)
}

func From[A any](ch chan A) Chan[A] {
	return Chan[A](ch)
}

func (ch Chan[A]) Send(a A) Chan[A] {
	ch <- a
	return ch
}
