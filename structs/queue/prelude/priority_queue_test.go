package prelude

import (
	"fmt"
	"testing"

	"github.com/ireina7/fgo/structs/maybe"
	"github.com/ireina7/fgo/structs/number"
	"github.com/ireina7/fgo/structs/number/nonzero"
)

func TestParent(t *testing.T) {
	i := 5
	maybe.For(nonzero.From(uint(i)), func(n number.NonZero) {
		fmt.Println(parent(n))
	})
}
