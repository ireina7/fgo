package effect

import (
	"errors"
	"fmt"
	"testing"
)

func TestEffect(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("[exception]", r)
		}
	}()
	eff := NewDefaultEffect[int, int]()
	ans := eff.Do(func() {
		eff.Throw(errors.New("exception"))
		eff.Resume(7)
		fmt.Println("here")
	})
	fmt.Println("Got", ans)
	eff.Return(0)
}
