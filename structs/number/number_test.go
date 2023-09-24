package number

import (
	"testing"

	"github.com/ireina7/fgo/interfaces"
)

func TestEnumInt(t *testing.T) {
	var enum interfaces.Enum[int] = &EnumInt{}
	t.Logf("%+v", enum.Range(0, 10))
}
