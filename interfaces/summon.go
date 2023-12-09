package interfaces

import (
	"github.com/ireina7/fgo/structs/maybe"
	"github.com/ireina7/fgo/types"
)

// * For https://github.com/ireina7/summoner
type Summoner[A any] interface {
	Summon() maybe.Maybe[A]
	Given(A)
}

func Instance[A any]() Summoner[A] {
	return nil
}

func Inject() {
	Instance[int]().Given(0)
	Instance[types.HKT[maybe.MaybeKind, int]]().Given(maybe.Some(7))
}
