package types

type Unit int

const UnitValue Unit = 0

func MakeUnit() Unit {
	return UnitValue
}

type absurd struct {
	v int
}
type Void absurd

type Nil any
