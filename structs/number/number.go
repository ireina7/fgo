package number

import (
	"github.com/ireina7/fgo/interfaces/collection"
	"github.com/ireina7/fgo/structs/maybe"
	"github.com/ireina7/fgo/structs/slice"
)

// * Enum for int
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

func (enum *EnumInt) Range(i, j int) collection.Iterator[int] {
	if j <= i {
		return collection.EmptyIter[int]()
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

// * Natural numbers
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

func (enum *EnumNat) Range(i, j Nat) collection.Iterator[Nat] {
	if j <= i {
		return collection.EmptyIter[Nat]()
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

func (iter *IntIter) Next() maybe.Maybe[int] {
	num := iter.curr
	iter.curr += 1
	return maybe.From(&num)
}

// * Non-zero Integers
type NonZero uint

// * Infinite precision Integers
type Integer struct {
	raw int64
}

// * Infinite precision floating numbers
type Number struct {
	raw float64
}
