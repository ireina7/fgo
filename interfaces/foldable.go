package interfaces

import (
	"github.com/ireina7/fgo/types"
)

type Foldable[F_, A, B any] interface {
	FoldLeft(types.HKT[F_, A], func(B, A) B, B) B
}
