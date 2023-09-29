package generic

import (
	"testing"

	"github.com/ireina7/fgo/interfaces"
)

func TestHashMap(t *testing.T) {
	hm := Make[string, int](
		&interfaces.Prelude[string]{},
		&interfaces.HashString{},
		10,
	)
	hm.Set("first", 1)
	hm.Set("second", 2)
	t.Log(hm.Get("first"))
}
