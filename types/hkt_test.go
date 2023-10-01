package types_test

import (
	"testing"

	"github.com/ireina7/fgo/structs/option"
	"github.com/ireina7/fgo/structs/result"
	"github.com/ireina7/fgo/types"
)

func TestWrappedHKT(t *testing.T) {
	var xs types.HKT[result.ResultKind, types.HKT[option.OptionKind, int]]
	xs = result.From[types.HKT[option.OptionKind, int]](option.Just(7))
	t.Logf("%#v", xs.(result.Result[types.HKT[option.OptionKind, int]]))
}
