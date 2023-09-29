package collection

import (
	"github.com/ireina7/fgo/types"
)

// F[A] should be a FromIterator[A]
type Collector[F_, A any] interface {
	Collect(Iterator[A]) types.HKT[F_, A]
}
