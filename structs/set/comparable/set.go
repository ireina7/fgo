package comparable

import (
	"fmt"
	"strings"

	"github.com/ireina7/fgo/types"
)

type SetKind any

type Set[A comparable] map[A]types.Unit

func (Set[A]) Kind(SetKind) {}
func (Set[A]) ElemType(A)   {}

func Empty[A comparable]() Set[A] {
	return map[A]types.Unit{}
}

func Make[A comparable](xs ...A) Set[A] {
	ss := Empty[A]()
	for _, x := range xs {
		ss[x] = types.MakeUnit()
	}
	return ss
}

func (ss Set[A]) String() string {
	xs := Map(ss, func(a A) string {
		return fmt.Sprintf("%v", a)
	})
	ys := make([]string, 0)
	for x := range xs {
		ys = append(ys, x)
	}
	content := strings.Join(ys, ", ")
	return fmt.Sprintf("Set{%s}", content)
}

func (ss Set[A]) Has(a A) bool {
	_, has := ss[a]
	return has
}

func (ss Set[A]) Add(a A) {
	ss[a] = types.MakeUnit()
}

func (ss Set[A]) Delete(a A) {
	delete(ss, a)
}

func (ss Set[A]) Len() int {
	return len(ss)
}

func Map[A, B comparable](ss Set[A], f func(A) B) Set[B] {
	rs := Empty[B]()
	for s := range ss {
		rs.Add(f(s))
	}
	return rs
}
