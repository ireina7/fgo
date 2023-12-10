// * functions about number pverflow
package checked

import (
	"github.com/ireina7/fgo/structs/maybe"
	"golang.org/x/exp/constraints"
)

func MinusUnsigned[T constraints.Unsigned](a, b T) maybe.Maybe[T] {
	if a < b {
		return maybe.None[T]()
	}
	return maybe.Some(a - b)
}
