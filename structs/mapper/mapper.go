package mapper

type Mapper[K comparable, V any] map[K]V

func From[K comparable, V any](m map[K]V) Mapper[K, V] {
	return Mapper[K, V](m)
}
