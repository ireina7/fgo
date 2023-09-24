package types

type Unit struct{}

var UnitValue Unit = Unit(struct{}{})

func MakeUnit() Unit {
	return UnitValue
}

type absurd struct {
	v int
}
type Void absurd

type Nil any
