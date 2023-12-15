package treemap

import (
	"fmt"
	"strings"
)

type color int

const (
	red   color = 0
	black color = 1
)

type TreeMap[K, V any] struct {
	root *Node[K, V]
}

type Node[K, V any] struct {
	key   K
	val   V
	color color
	Left  *Node[K, V]
	Right *Node[K, V]
}

func (tree TreeMap[K, V]) Debug() string {
	if tree.root == nil {
		return "nil"
	}
	ans := &strings.Builder{}
	ans.WriteString(fmt.Sprintf("%v", tree.root.key))
	return ans.String()
}
