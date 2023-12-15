package ref

// * Thread safe reference
type Cell[A any] struct {
	ref Ref[A]
}
