package types

type Unit struct{}

func MakeUnit() Unit {
	return struct{}{}
}

type Void struct{}

type Nil any
