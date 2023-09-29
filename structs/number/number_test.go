package number_test

import (
	"testing"

	"github.com/ireina7/fgo/interfaces"
	"github.com/ireina7/fgo/interfaces/collection"
	"github.com/ireina7/fgo/structs/number"
)

func TestEnumInt(t *testing.T) {
	var enum interfaces.Enum[int] = number.NewEnumInt()
	collection.For[int](collection.FromIter(enum.Range(0, 10)), func(x int) {
		t.Log("Got", x)
	})
}
