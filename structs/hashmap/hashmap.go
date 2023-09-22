package hashmap

type HashMap[K comparable, V any] map[K]V

func From[K comparable, V any](m map[K]V) HashMap[K, V] {
	return HashMap[K, V](m)
}
