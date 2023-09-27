package number

import (
	"github.com/ireina7/fgo/interfaces"
	"github.com/ireina7/fgo/structs/option"
	"github.com/ireina7/fgo/structs/slice"
)

type EnumInt struct{}

func (enum *EnumInt) Succ(x int) int {
	return x + 1
}

func (enum *EnumInt) Pred(x int) int {
	return x - 1
}

func (enum *EnumInt) ToEnum(i int) int {
	return i
}

func (enum *EnumInt) Range(i, j int) interfaces.Iterator[int] {
	if j <= i {
		return interfaces.EmptyIter[int]()
	}
	ans := make([]int, 0, j-i)
	for k := i; k < j; k++ {
		ans = append(ans, k)
	}
	return slice.From(ans).Iter()
}

func NewEnumInt() *EnumInt {
	return &EnumInt{}
}

type Nat uint

type EnumNat struct{}

func (enum *EnumNat) Succ(x Nat) Nat {
	return x + 1
}

func (enum *EnumNat) Pred(x Nat) Nat {
	if x == 0 {
		return 0
	}
	return x - 1
}

func (enum *EnumNat) ToEnum(i int) Nat {
	if i < 0 {
		return 0
	}
	return Nat(uint(i))
}

func (enum *EnumNat) Range(i, j Nat) interfaces.Iterator[Nat] {
	if j <= i {
		return interfaces.EmptyIter[Nat]()
	}
	ans := make([]Nat, 0, j-i)
	var k Nat
	for k = i; k < j; k++ {
		ans = append(ans, k)
	}
	return slice.From(ans).Iter()
}

type IntIter struct {
	curr int
}

func (IntIter) From(start int) *IntIter {
	return &IntIter{
		curr: start,
	}
}

func (iter *IntIter) Next() option.Option[int] {
	num := iter.curr
	iter.curr += 1
	return option.From(&num)
}
