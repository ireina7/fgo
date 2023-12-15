package tuple

import "github.com/samber/lo"

type Tuple2[A, B any] lo.Tuple2[A, B]
type Tuple3[A, B, C any] lo.Tuple3[A, B, C]
type Tuple4[A, B, C, D any] lo.Tuple4[A, B, C, D]
type Tuple5[A, B, C, D, E any] lo.Tuple5[A, B, C, D, E]
type Tuple6[A, B, C, D, E, F any] lo.Tuple6[A, B, C, D, E, F]
type Tuple7[A, B, C, D, E, F, G any] lo.Tuple7[A, B, C, D, E, F, G]

// Unpack returns values contained in tuple.
func (t Tuple2[A, B]) Unpack() (A, B) {
	return t.A, t.B
}

func Of2[A, B any](a A, b B) Tuple2[A, B] {
	return Tuple2[A, B](lo.T2(a, b))
}
