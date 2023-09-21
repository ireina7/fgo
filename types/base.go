package types

type Unit struct{}

func MakeUnit() Unit {
	return struct{}{}
}
