package types

type Unit struct{}

var UnitValue Unit = Unit(struct{}{})

func MakeUnit() Unit {
	return UnitValue
}

type absurd interface {
	neverImplMe()
}

type Void absurd

func IsVoid(x any) bool {
	return x == nil
}

type Nil any
