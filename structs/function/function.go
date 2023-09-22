package function

import "github.com/ireina7/fgo/types"

type ReaderKind any

type Reader[A any] types.HKT[ReaderKind, A]

type Func[A, B any] func(A) B

func (f Func[A, B]) Kind(ReaderKind) {}
func (f Func[A, B]) Elem(B)          {}

func From[A, B any](f func(A) B) Func[A, B] {
	return Func[A, B](f)
}

func Unpack[A, B any](f Func[A, B]) func(A) B {
	return f
}

func Identity[A any]() func(A) A {
	return func(a A) A {
		return a
	}
}

func Compose[A, B, C any](f func(A) B, g func(B) C) func(A) C {
	return func(a A) C {
		return g(f(a))
	}
}

func Flip[A, B, C any](f func(A, B) C) func(B, A) C {
	return func(b B, a A) C {
		return f(a, b)
	}
}

func Curry2[A, B, C any](f func(A, B) C) func(A) func(B) C {
	return func(a A) func(b B) C {
		return func(b B) C {
			return f(a, b)
		}
	}
}

func Curry3[A, B, C, D any](f func(A, B, C) D) func(A) func(B) func(C) D {
	return func(a A) func(b B) func(c C) D {
		return Curry2(func(b B, c C) D {
			return f(a, b, c)
		})
	}
}

func Curry4[A, B, C, D, E any](f func(A, B, C, D) E) func(A) func(B) func(C) func(D) E {
	return func(a A) func(b B) func(c C) func(d D) E {
		return Curry3(func(b B, c C, d D) E {
			return f(a, b, c, d)
		})
	}
}

func Curry5[A, B, C, D, E, F any](f func(A, B, C, D, E) F) func(A) func(B) func(C) func(D) func(E) F {
	return func(a A) func(b B) func(c C) func(d D) func(E) F {
		return Curry4(func(b B, c C, d D, e E) F {
			return f(a, b, c, d, e)
		})
	}
}
