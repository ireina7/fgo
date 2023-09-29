package linklist

import (
	"testing"

	"github.com/ireina7/fgo/interfaces/collection"
)

func TestListMake(t *testing.T) {
	xs := Make(1, 2, 3, 4, 5)
	collection.For[int](xs, func(x int) {
		t.Log(x)
	})
}
