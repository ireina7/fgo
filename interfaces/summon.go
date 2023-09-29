package interfaces

import (
	"github.com/ireina7/fgo/structs/option"
	"github.com/ireina7/fgo/types"
)

// * For https://github.com/ireina7/summoner
type Summoner[A any] interface {
	Summon() option.Option[A]
	Given(A)
}

func Instance[A any]() Summoner[A] {
	return nil
}

func Inject() {
	Instance[int]().Given(0)
	Instance[types.HKT[option.OptionKind, int]]().Given(option.Just(7))
}
