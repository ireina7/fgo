package prelude

import (
	"fmt"
	"testing"

	"github.com/ireina7/fgo/structs/maybe"
	"github.com/ireina7/fgo/structs/number"
	"github.com/ireina7/fgo/structs/number/nonzero"
)

func TestParent(t *testing.T) {
	i := 5
	maybe.For(nonzero.From(uint(i)), func(n number.NonZero) {
		fmt.Println(parent(n))
	})
}

func TestBuild(t *testing.T) {
	q := FromSlice([]int{1, 2, 31, 4, 5})
	fmt.Println(q)

	q = q.Push(9)
	fmt.Println(q)
	q = q.Push(8)
	fmt.Println(q)
	q = q.Push(18)
	fmt.Println(q)

	n, q := q.Pop()
	for n.IsSome() {
		fmt.Print(n.MustGet(), ", ")
		n, q = q.Pop()
	}
}
