package main

import (
	"fmt"

	"github.com/ireina7/fgo/structs/function"
)

func main() {
	fmt.Println("functional go!")
	x := function.Curry5(add)(1)(2)(3)(4)
	fmt.Println(x(5))
}

func add(a, b, c, d, e int) int {
	return a + b + c + d + e
}
