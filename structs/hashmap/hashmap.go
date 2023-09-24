package hashmap

import "github.com/ireina7/fgo/structs/option"

type HashMap[K comparable, V any] map[K]V

func From[K comparable, V any](m map[K]V) HashMap[K, V] {
	return HashMap[K, V](m)
}

func (hm HashMap[K, V]) Get(key K) option.Option[V] {
	v, exist := hm[key]
	if !exist {
		return option.From[V](nil)
	}
	return option.From(&v)
}

func (hm HashMap[K, V]) Set(key K, v V) HashMap[K, V] {
	hm[key] = v
	return hm
}

func (hm HashMap[K, V]) Contains(key K) bool {
	_, exist := hm[key]
	return exist
}
