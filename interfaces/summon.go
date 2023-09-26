package interfaces

import "github.com/ireina7/fgo/structs/option"

// * For https://github.com/ireina7/summoner
type Summoner[A any] interface {
	Summon() option.Option[A]
	Given(A)
}
