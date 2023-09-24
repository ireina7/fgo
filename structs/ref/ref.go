package ref

type RefKind any

// * Reference type
// Ensure not nil
type ptr[A any] struct {
	ptr *A
}

type Ref[A any] struct {
	ptr ptr[A]
}

func (Ref[A]) Kind(RefKind) {}
func (Ref[A]) ElemType(A)   {}

func Refer[A any](a A) Ref[A] {
	return Ref[A]{ptr: ptr[A]{ptr: &a}}
}

func Make[A any](a A) Ref[A] {
	return Refer(a)
}

func (ref Ref[A]) Deref() A {
	return *ref.ptr.ptr
}

func Map[A, B any](ref Ref[A], f func(A) B) Ref[B] {
	b := f(ref.Deref())
	return Refer(b)
}
