package nonzero

import (
	"github.com/ireina7/fgo/structs/maybe"
	"github.com/ireina7/fgo/structs/number"
	"golang.org/x/exp/constraints"
)

func From[T constraints.Integer](i T) maybe.Maybe[number.NonZero] {
	if i <= 0 {
		return maybe.None[number.NonZero]()
	}
	return maybe.Some(number.NonZero(i))
}
