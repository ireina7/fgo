package generic

import (
	"testing"

	"github.com/ireina7/fgo/interfaces"
	"github.com/ireina7/fgo/interfaces/collection"
	"github.com/ireina7/fgo/structs/tuple"
)

func TestHashMap(t *testing.T) {
	hm := Make[string, int](
		&interfaces.Prelude[string]{},
		&interfaces.HashString{},
		10,
	)
	hm.Set("first", 1)
	hm.Set("second", 2)
	hm.Set("third", 3)
	collection.For[tuple.Tuple2[string, int]](hm, func(xs tuple.Tuple2[string, int]) {
		t.Log(xs.A, xs.B)
	})
}
