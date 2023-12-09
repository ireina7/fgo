package types_test

import (
	"fmt"
	"testing"

	option "github.com/ireina7/fgo/structs/maybe"
	"github.com/ireina7/fgo/structs/result"
	"github.com/ireina7/fgo/types"
)

func TestWrappedHKT(t *testing.T) {
	var xs types.HKT[result.ResultKind, types.HKT[option.MaybeKind, int]]
	xs = result.From[types.HKT[option.MaybeKind, int]](option.Some(7))
	t.Logf("%#v", xs.(result.Result[types.HKT[option.MaybeKind, int]]))
}

func TestAbsurd(t *testing.T) {
	var x types.Void
	fmt.Println(x)
}
