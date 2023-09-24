package number

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

func (enum *EnumInt) Range(i, j int) []int {
	if j <= i {
		return []int{}
	}
	ans := make([]int, 0, j-i)
	for k := i; k < j; k++ {
		ans = append(ans, k)
	}
	return ans
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

func (enum *EnumNat) Range(i, j Nat) []Nat {
	if j <= i {
		return []Nat{}
	}
	ans := make([]Nat, 0, j-i)
	var k Nat
	for k = i; k < j; k++ {
		ans = append(ans, k)
	}
	return ans
}