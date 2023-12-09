package hashmap

import (
	"github.com/ireina7/fgo/structs/maybe"
	"github.com/ireina7/fgo/types"
)

type HashMapKind types.Kind

type HashMap[K comparable, V any] map[K]V

func (HashMap[K, V]) Kind(HashMapKind) {}
func (HashMap[K, V]) ElemType(V)       {}

func From[K comparable, V any](m map[K]V) HashMap[K, V] {
	return HashMap[K, V](m)
}

func (hm HashMap[K, V]) Get(key K) maybe.Maybe[V] {
	v, exist := hm[key]
	if !exist {
		return maybe.From[V](nil)
	}
	return maybe.From(&v)
}

func (hm HashMap[K, V]) Set(key K, v V) HashMap[K, V] {
	hm[key] = v
	return hm
}

func (hm HashMap[K, V]) Has(key K) bool {
	_, exist := hm[key]
	return exist
}

func (hm HashMap[K, V]) Delete(key K) {
	delete(hm, key)
}
