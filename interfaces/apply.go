package interfaces

import (
	"github.com/ireina7/fgo/types"
	"github.com/samber/lo"
)

type Apply[F_, A, B any] interface {
	Unit() types.HKT[F_, types.Unit]
	Product(types.HKT[F_, A], types.HKT[F_, B]) types.HKT[F_, lo.Tuple2[A, B]]
}
