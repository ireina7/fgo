package linklist

import (
	"testing"

	"github.com/ireina7/fgo/interfaces"
)

func TestListMake(t *testing.T) {
	xs := Make(1, 2, 3, 4, 5)
	interfaces.For[int](xs, func(x int) {
		t.Log(x)
	})
}
