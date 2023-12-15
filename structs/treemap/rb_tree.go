package treemap

type TreeMap[K, V any] struct {
	root Node[K, V]
}

type Node[K, V any] struct {
	key   K
	val   V
	Left  *Node[K, V]
	Right *Node[K, V]
}

func (tree *TreeMap[K, V]) Debug() string {
	return ""
}
